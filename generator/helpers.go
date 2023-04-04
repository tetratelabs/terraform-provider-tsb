// Copyright 2023 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"strings"

	j "github.com/dave/jennifer/jen"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func NewResource(opts ...resourceOption) *resource {
	r := &resource{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

type resourceOption func(*resource)

func WithName(name string) resourceOption {
	return func(r *resource) {
		r.Name = name
		r.lowerName = strings.ToLower(name)
		r.structId = j.Id(r.Name + "Resource")
		r.modelId = j.Id(r.Name + "Model")
	}
}

func WithPkgImportPath(path string) resourceOption {
	return func(r *resource) {
		r.PkgImportPath = path
	}
}

func WithClient(client string) resourceOption {
	return func(r *resource) {
		r.Client = client
	}
}

func WithSchema(s schema.Schema) resourceOption {
	return func(r *resource) {
		r.Schema = s
	}
}

type resource struct {
	Name          string
	PkgImportPath string
	Client        string
	Schema        schema.Schema

	structId  j.Code
	modelId   j.Code
	lowerName string
}

func (r resource) LoadPlanIntoModel() j.Statement {
	return []j.Code{
		j.Var().Id("model").Add(r.modelId),
		j.Id("resp").Dot("Diagnostics").Dot("Append").Call(
			j.Id("req").Dot("Plan").Dot("Get").Call(j.Id("ctx"), j.Op("&").Id("model")).Op("..."),
		),
		j.If(j.Id("resp").Dot("Diagnostics").Dot("HasError").Call()).Block(j.Return()),
	}
}

func (r resource) LoadStateIntoModel() j.Statement {
	return []j.Code{
		j.Var().Id("model").Add(r.modelId),
		j.Id("resp").Dot("Diagnostics").Dot("Append").Call(
			j.Id("req").Dot("State").Dot("Get").Call(j.Id("ctx"), j.Op("&").Id("model")).Op("..."),
		),
		j.If(j.Id("resp").Dot("Diagnostics").Dot("HasError").Call()).Block(j.Return()),
	}
}

func handleError(topLevelMessage, secondLevelMessage string) j.Code {
	return j.If(j.Id("err").Op("!=").Nil()).Block(
		j.Id("resp").Dot("Diagnostics").Dot("AddError").Call(
			j.Lit(topLevelMessage),
			j.Lit(secondLevelMessage+": ").Op("+").Id("err").Dot("Error").Call(),
		),
		j.Return(),
	)
}

func saveState() j.Code {
	return j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("resp").Dot("State").Dot("Set").Call(j.Id("ctx"), j.Op("&").Id("model")).Op("..."))
}

// Gross, but go doesn't support type guards and I don't want to do several levels of append inline
func mergeCode(code ...interface{}) j.Statement {
	res := []j.Code{}
	for _, c := range code {
		switch v := c.(type) {
		case j.Code:
			res = append(res, v)
		case []j.Code:
			res = append(res, v...)
		case j.Statement:
			res = append(res, v...)
		}
	}
	return res
}
