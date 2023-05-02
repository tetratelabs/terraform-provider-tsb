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
		res = append(res, j.Id("model").Dot(snakeToCamel(k)).Op("=").Add(modelField([]string{k}, []string{"cluster"}, v)))
	}
	return res
}

func modelField(modelPath []string, apiPath []string, attr schema.Attribute) j.Code {

	//modelPath := j.Id(fullyQualified[0]).Add(lo.Map(fullyQualified[1:], func(s string, _ int) j.Code { return j.Dot(s) })...)
	p2c := func(path []string) j.Code {
		return j.Id(path[0]).Add(lo.Map(path[1:], func(s string, _ int) j.Code {
			return j.Dot(s)
		})...)
	}
	switch t := attr.(type) {
	case schema.StringAttribute:
		if t.CustomType == nil {
			return j.Qual(Types, "StringValue").Call(j.Id("rLowerName").Dot(snakeToCamel(apiPath[0])))
		}
		// If it's a string but has a custom type it's an enum.
		// This code works if the enum is in the same package as the resource, but we may need to pass
		// more information to the struct at schema generation time in the future.
		return j.Qual(Types, "StringValue").Call(j.Qual("pkgimportpath", snakeToCamel(apiPath[0])+"_name").Index(j.Int32().Call(j.Id("rLowerName").Dot(snakeToCamel(apiPath[0])))))
	case schema.ListNestedAttribute:
		return j.Func().Call().Op("[]*").Add(p2c(modelPath)).Block(
			j.Id("tmp").Op(":=").Make(j.Op("[]*").Add(p2c(modelPath)), j.Len(p2c(apiPath))),
			j.For(j.List(j.Id("_"), j.Id("n")).Op(":=").Range().Add(p2c(apiPath))).Block(
				j.Id("tmp").Op("=").Append(j.Id("tmp"), j.Lit("magicalRecursion")),
			),
			j.Return(j.Id("tmp")),
		).Call()
	case schema.MapAttribute:
		return j.Func().Call().Qual("basetypes", "MapValue").Block(
			j.List(j.Id("r"), j.Id("diag")).Op(":=").Qual(Types, "MapValue").Call(
				j.Id("ctx"),
				j.Id(modelPath[0]).Dot("ElementType").Call(j.Id("ctx")),
				j.Id(apiPath[0]),
			),
			j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("diag").Op("...")),
			j.Return(j.Id("r")),
		).Call()
	default:
		// fmt.Printf("%v isn't supported as a resource value\n", t.GetType().String())
		return j.Nil()
	}
}

// func buildInnerResource(s map[string]schema.Attribute, prefixes []string, pkgImportPath string) j.Dict {
// 	res := j.Dict{}
// 	for k, v := range s {
// 		fullyQualified := append(prefixes, snakeToCamel(k))
// 		modelPath := j.Id(fullyQualified[0]).Add(lo.Map(fullyQualified[1:], func(s string, _ int) j.Code { return j.Dot(s) })...)
// 		desc := strings.Split(v.GetDescription(), "!!")
// 		gotype := desc[len(desc)-1]
// 		goimport := desc[len(desc)-2]

// 		switch t := v.(type) {
// 		// TODO: handle non-string fields
// 		case schema.StringAttribute:
// 			if gotype == "BytesKind" {
// 				res[j.Id(snakeToCamel(k))] = j.Op("[]").Byte().Call(modelPath.Dot("ValueString").Call())
// 				continue
// 			}
// 			if t.CustomType == nil {
// 				res[j.Id(snakeToCamel(k))] = modelPath.Dot("ValueString").Call()
// 				continue
// 			}
// 			// If it's a string but has a custom type it's an enum.
// 			// This code works if the enum is in the same package as the resource, but we may need to pass
// 			// more information to the struct at schema generation time in the future.
// 			if goimport == "primitive" {
// 				goimport = pkgImportPath
// 			}
// 			res[j.Id(snakeToCamel(k))] = j.Qual(goimport, gotype).Call(
// 				j.Qual(goimport, gotype+"_value").Index(modelPath.Dot("ValueString").Call()),
// 			)
// 		case schema.BoolAttribute:
// 			res[j.Id(snakeToCamel(k))] = modelPath.Dot("ValueBool").Call()
// 		case schema.Int64Attribute:
// 			var enclosing *j.Statement
// 			switch gotype {
// 			case "Int32Kind":
// 				enclosing = j.Int32()
// 			case "Int64Kind":
// 				res[j.Id(snakeToCamel(k))] = modelPath.Dot("ValueInt64").Call()
// 				continue
// 			case "Uint32Kind":
// 				enclosing = j.Uint32()
// 			case "Uint64Kind":
// 				enclosing = j.Uint64()
// 			default:
// 				fmt.Printf("%s unknown int kind", gotype)
// 				enclosing = j.Int64()
// 			}
// 			res[j.Id(snakeToCamel(k))] = enclosing.Call(modelPath.Dot("ValueInt64").Call())

// 		case schema.MapAttribute:
// 			// tmp := make(map[string]string)
// 			// diag.Append(data.Headers.ElementsAs(ctx, &tmp, false)...)
// 			//TODO handle different map types
// 			res[j.Id(snakeToCamel(k))] = j.Func().Call().Map(j.String()).String().Block(
// 				j.Id("tmp").Op(":=").Make(j.Map(j.String()).String()),
// 				j.Id("resp").Dot("Diagnostics").Dot("Append").Call(modelPath.Dot("ElementsAs").Call(j.Id("ctx"), j.Op("&").Id("tmp"), j.False()).Op("...")),
// 				j.Return(j.Id("tmp")),
// 			).Call()
// 		case schema.MapNestedAttribute:
// 			// TODO: handle inner object
// 			res[j.Id(snakeToCamel(k))] = j.Func().Call().Map(j.String()).String().Block(
// 				j.Id("tmp").Op(":=").Make(j.Map(j.String()).String()),
// 				j.Id("resp").Dot("Diagnostics").Dot("Append").Call(modelPath.Dot("ElementsAs").Call(j.Id("ctx"), j.Op("&").Id("tmp"), j.False()).Op("...")),
// 				j.Return(j.Id("tmp")),
// 			).Call()

// 		case schema.ListAttribute:
// 			// TODO: handle different list types
// 			listtype := "string"
// 			res[j.Id(snakeToCamel(k))] = j.Func().Call().Op("[]").Id(listtype).Block(
// 				j.Id("tmp").Op(":=").Make(j.Op("[]").Id(listtype), j.Len(j.Add(modelPath.Clone()).Dot("Elements").Call())),
// 				j.Id("resp").Dot("Diagnostics").Dot("Append").Call(modelPath.Clone().Dot("ElementsAs").Call(j.Id("ctx"), j.Op("&").Id("tmp"), j.False()).Op("...")),
// 				j.Return(j.Id("tmp")),
// 			).Call()

// 		case schema.ListNestedAttribute:
// 			underlying := make(map[string]schema.Attribute)
// 			for k, v := range t.GetNestedObject().GetAttributes() {
// 				underlying[k] = v
// 			}

// 			// res[j.Id(snakeToCamel(k))] = j.Qual(SamberLo, "Map").Call(modelPath, j.Func().Call(
// 			// 	j.Id(k).Interface(), j.Id("index").Int(),
// 			// ).Op("*").Qual(goimport, gotype).Block(j.Return(
// 			// 	j.Op("&").Qual(goimport, gotype).Values(),
// 			// )))
// 			res[j.Id(snakeToCamel(k))] = j.Func().Call().Op("[]*").Qual(goimport, gotype).Block(
// 				j.Id("a").Op(":=").Make(j.Op("[]*").Qual(goimport, gotype), j.Len(modelPath.Clone())),
// 				j.For(j.Id("i").Op(",").Id(k).Op(":=").Range().Add(modelPath.Clone())).Block(
// 					j.Id("a").Index(j.Id("i")).Op("=").Op("&").Qual(goimport, gotype).Values(
// 						buildInnerResource(underlying, []string{k}, pkgImportPath),
// 					),
// 				),
// 				j.Return(j.Id("a")),
// 			).Call()
// 			// j.Op("&").Qual(goimport, gotype).Values(
// 			// 	buildInnerResource(underlying, []string{k}, pkgImportPath),
// 			// ),
// 			// )
// 			continue

// 		case schema.NestedAttribute:
// 			underlying := make(map[string]schema.Attribute)
// 			for k, v := range t.GetNestedObject().GetAttributes() {
// 				underlying[k] = v
// 			}
// 			// desc := strings.Split(t.GetDescription(), "!!")
// 			// gotype := desc[len(desc)-1]
// 			// goimport := desc[len(desc)-2]

// 			res[j.Id(snakeToCamel(k))] = j.Op("&").Qual(goimport, gotype).Values(
// 				buildInnerResource(underlying, append(prefixes, snakeToCamel(k)), pkgImportPath),
// 			)
// 			continue
// 		default:
// 			// panic(fmt.Sprintf("%v isn't supported as a resource value", t))
// 			fmt.Printf("%v isn't supported as a resource value\n", t.GetType().String())
// 		}

// 	}
// 	return res
// }
