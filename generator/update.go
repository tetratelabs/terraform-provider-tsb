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
	j "github.com/dave/jennifer/jen"
	"github.com/samber/lo"
)

func genUpdate(r resource) *j.File {
	f := j.NewFile(r.lowerName)

	f.Func().
		Parens(j.Id("r").Op("*").Add(r.structId)).Id("Update").
		Params(j.Id("ctx").Qual("context", "Context"), j.Id("req").Qual(Resource, "UpdateRequest"), j.Id("resp").Op("*").Qual(Resource, "UpdateResponse")).
		Block(
			mergeCode(
				r.LoadPlanIntoModel(),

				// Get the latest etag
				j.Id("request").Op(":=").Op("&").Qual(r.PkgImportPath, "Get"+r.Name+"Request").Values(j.Dict{
					j.Id("Fqn"): j.Id("model").Dot("Id").Dot("ValueString").Call(),
				}),
				j.List(j.Id(r.lowerName), j.Id("err")).Op(":=").Id("r").Dot("client").Dot("Get"+r.Name).Call(
					j.Id("ctx"),
					j.Id("request"),
				),
				handleError("Error updating "+r.Name, "Get"+r.Name+"Request failed"),

				// Do the actual request
				j.Id("updateTo").Op(":=").Op("&").Qual(r.PkgImportPath, r.Name).Values(buildUpdateRequest(r)),
				j.List(j.Op("_"), j.Id("err")).Op("=").Id("r").Dot("client").Dot("Update"+r.Name).Call(
					j.Id("ctx"),
					j.Id("updateTo"),
				),
				handleError("Error updating "+r.Name, "Update"+r.Name+" failed"),

				saveState(),
			)...,
		)

	return f
}

func buildUpdateRequest(r resource) j.Dict {
	res := buildInnerResource(lo.OmitByKeys(r.Schema.Attributes, []string{"parent", "id", "name"}), []string{"model"}, r.PkgImportPath)

	res[j.Id("Fqn")] = j.Id("model").Dot("Id").Dot("ValueString").Call()
	res[j.Id("Etag")] = j.Id(r.lowerName).Dot("Etag")

	return res
}
