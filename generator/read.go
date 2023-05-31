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
	"fmt"

	j "github.com/dave/jennifer/jen"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/samber/lo"
)

func genRead(r resource) *j.File {
	f := j.NewFile(r.lowerName)

	f.Func().
		Parens(j.Id("r").Op("*").Add(r.structId)).Id("Read").
		Params(j.Id("ctx").Qual("context", "Context"), j.Id("req").Qual(Resource, "ReadRequest"), j.Id("resp").Op("*").Qual(Resource, "ReadResponse")).
		Block(
			mergeCode(
				r.LoadStateIntoModel(),

				// Do the actual request
				j.Id("request").Op(":=").Op("&").Qual(r.PkgImportPath, "Get"+r.Name+"Request").Values(j.Dict{
					j.Id("Fqn"): j.Id("model").Dot("Id").Dot("ValueString").Call(),
				}),
				j.List(j.Id(r.lowerName), j.Id("err")).Op(":=").Id("r").Dot("client").Dot("Get"+r.Name).Call(
					j.Id("ctx"),
					j.Id("request"),
				),
				handleError("Error reading "+r.Name, "Get"+r.Name+"Request failed"),

				// Get meta object from response
				j.List(j.Id("meta"), j.Id("err")).Op(":=").Qual(Helpers, "FromFQN").Call(
					j.Qual(PkgAPI, r.Name+"Kind"),
					j.Id("model").Dot("Id").Dot("ValueString").Call(),
				),
				handleError("Error reading"+r.Name, "FQN parsing failed"),

				populateModel(r, lo.OmitByKeys(r.Schema.Attributes, []string{"parent", "id", "name"})),

				saveState(),
			)...,
		)

	return f
}

func populateModel(r resource, s map[string]schema.Attribute) []j.Code {
	// Handle common variables manually
	res := []j.Code{
		j.Id("model").Dot("Id").Op("=").Qual(Types, "StringValue").Call(j.Id(r.lowerName).Dot("Fqn")),
		j.Id("model").Dot("Name").Op("=").Qual(Types, "StringValue").Call(j.Id("meta").Dot("Name")),
		j.Id("model").Dot("Parent").Op("=").Qual(Types, "StringValue").Call(j.Qual(Helpers, "ParentFQN").Call(j.Qual(PkgAPI, r.Name+"Kind"), j.Id("meta"))),
	}

	// Handle remining recursively
	for k, v := range s {
		res = append(res, j.Id("model").Dot(snakeToCamel(k)).Op("=").Add(modelField(r.PkgImportPath, []string{k}, []string{r.lowerName, k}, v)))
	}
	return res
}

func modelField(pkgImportPath string, modelPath []string, apiPath []string, attr schema.Attribute) *j.Statement {
	a2c := func() *j.Statement {
		return j.Id(apiPath[0]).Add(lo.Map(apiPath[1:], func(s string, _ int) j.Code {
			return j.Dot(snakeToCamel(s))
		})...)
	}
	m2c := func() *j.Statement {
		return j.Id(snakeSliceToCamel(modelPath) + "_Model")
	}

	typeInfo, err := getGoTypes(attr.GetType())
	if err != nil {
		return j.Nil().Commentf("failed to get type for %v. %v", attr, err)
	}
	goType := typeInfo.GetGoType()
	goImport := typeInfo.GetGoImport()

	switch t := attr.(type) {
	case schema.StringAttribute:
		if goType == "BytesKind" {
			return j.Qual(Types, "StringValue").Call(a2c())
		}
		if goType == "StringKind" {
			return j.Qual(Types, "StringValue").Call(a2c())
		}

		if goImport == "primitive" {
			// return j.Qual(Types, "StringValue").Call(j.Id("rLowerName").Dot(snakeToCamel(apiPath[0])))
			// return j.Qual(Types, "StringValue").Call(a2c())
			goImport = pkgImportPath
		}
		// If it's a string but has a custom type it's an enum.
		// This code works if the enum is in the same package as the resource, but we may need to pass
		// more information to the struct at schema generation time in the future.
		// return j.Qual(Types, "StringValue").Call(j.Qual(goImport, snakeToCamel(modelPath[len(modelPath)-1])+"_name").Index(j.Int32().Call(a2c())))
		return j.Qual(Types, "StringValue").Call(j.Qual(goImport, goType+"_name").Index(j.Int32().Call(a2c())))
	case schema.Int64Attribute:
		return j.Qual(Types, "Int64Value").Call(j.Int64().Call(a2c()))
	case schema.Float64Attribute:
		return j.Qual(Types, "Float64Value").Call(a2c())
	case schema.BoolAttribute:
		return j.Qual(Types, "BoolValue").Call(a2c())
	case schema.ListNestedAttribute:
		inner := j.Dict{}
		for k, v := range t.NestedObject.Attributes {
			inner[j.Id(snakeToCamel(k))] = modelField(pkgImportPath, append(modelPath, k), append(apiPath, snakeToCamel(k)), v) //j.Lit(fmt.Sprintf("modelField(append(%v, x), append(%v, x), %v)", modelPath, apiPath, v))

		}
		return j.Func().Call().Op("[]*").Add(m2c()).Block(
			j.Id("size").Op(":=").Len(a2c()),
			j.Id("tmp").Op(":=").Make(j.Op("[]*").Add(m2c()), j.Id("size"), j.Id("size")),
			j.For(j.List(j.Id("i"), j.Id("v")).Op(":=").Range().Add(a2c())).Block(
				j.Id("tmp").Index(j.Id("i")).Op("=").Op("&").Add(m2c()).Values(inner),
			),
			j.Comment("asdfasd"),
			j.Return(j.Id("tmp")),
		).Call()
	// 	return j.Lit(fmt.Sprintf("%v / %v", typeInfo.GetGoImport(), typeInfo.GetGoType()))
	case schema.ListAttribute:
		return j.Func().Call().Qual("basetypes", "ListValue").Block(
			j.List(j.Id("r"), j.Id("diag")).Op(":=").Qual(Types, "ListValueFrom").Call(
				j.Id("ctx"),
				j.Id("model").Add(lo.Map(modelPath, func(s string, _ int) j.Code {
					return j.Dot(snakeToCamel(s))
				})...).Dot("ElementType").Call(j.Id("ctx")),
				a2c(),
			),
			j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("diag").Op("...")),
			j.Return(j.Id("r")),
		).Call()

		// return j.Func().Call().Op("[]*").Add(a2c()).Block(
		// 	j.Id("tmp").Op(":=").Make(j.Op("[]*").Add(p2c(modelPath)), j.Len(p2c(apiPath))),
		// 	j.For(j.List(j.Id("_"), j.Id("n")).Op(":=").Range().Add(p2c(apiPath))).Block(
		// 		j.Id("tmp").Op("=").Append(j.Id("tmp"), j.Lit("magicalRecursion")),
		// 	),
		// 	j.Return(j.Id("tmp")),
		// ).Call()
	// case schema.MapNestedAttribute:
	// 	return j.Lit(fmt.Sprintf("%v / %v", typeInfo.GetGoImport(), typeInfo.GetGoType()))
	case schema.MapAttribute:
		return j.Func().Call().Qual("basetypes", "MapValue").Block(
			// j.Comment("TODO: populate map with attr.Valueified values"),
			// j.Id("tmp").Op(":=").Make(j.Map(j.String()).Qual("attr", "Value")),
			j.List(j.Id("r"), j.Id("diag")).Op(":=").Qual(Types, "MapValueFrom").Call(
				j.Id("ctx"),
				j.Id("model").Add(lo.Map(modelPath, func(s string, _ int) j.Code {
					return j.Dot(snakeToCamel(s))
				})...).Dot("ElementType").Call(j.Id("ctx")),
				a2c(),
			),
			j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("diag").Op("...")),
			j.Return(j.Id("r")),
		).Call()
	case schema.SingleNestedAttribute:
		inner := j.Dict{}
		for k, v := range t.Attributes {
			inner[j.Id(snakeToCamel(k))] = modelField(pkgImportPath, append(modelPath, k), append(apiPath, snakeToCamel(k)), v) //j.Lit(fmt.Sprintf("modelField(append(%v, x), append(%v, x), %v)", modelPath, apiPath, v))
		}

		return m2c().Values(inner)

	default:
		// fmt.Printf("%v isn't supported as a resource value\n", t.GetType().String())
		return j.Lit(fmt.Sprintf("%T: %v / %v", t, typeInfo.GetGoImport(), typeInfo.GetGoType()))
	}
}
