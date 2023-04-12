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
	"strings"

	j "github.com/dave/jennifer/jen"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/samber/lo"
)

func genCreate(r resource) *j.File {
	f := j.NewFile(r.lowerName)

	f.Func().
		Parens(j.Id("r").Op("*").Add(r.structId)).Id("Create").
		Params(j.Id("ctx").Qual("context", "Context"), j.Id("req").Qual(Resource, "CreateRequest"), j.Id("resp").Op("*").Qual(Resource, "CreateResponse")).
		Block(
			mergeCode(
				r.LoadPlanIntoModel(),

				// Do the actual request
				j.Id("request").Op(":=").Op("&").Qual(r.PkgImportPath, "Create"+r.Name+"Request").Values(buildCreateRequest(r)),
				j.List(j.Id(r.lowerName), j.Id("err")).Op(":=").Id("r").Dot("client").Dot("Create"+r.Name).Call(
					j.Id("ctx"),
					j.Id("request"),
				),

				handleError("Error creating "+r.Name, "Create"+r.Name+" request failed"),

				// Set ID in the model and then save it to state
				j.Id("model").Dot("Id").Op("=").Qual(Types, "StringValue").Call(j.Id(r.lowerName).Dot("Fqn")),
				saveState(),
			)...,
		)

	return f
}

func buildCreateRequest(r resource) j.Dict {
	res := j.Dict{
		j.Id("Name"):   j.Id("model").Dot("Name").Dot("ValueString").Call(),
		j.Id("Parent"): j.Id("model").Dot("Parent").Dot("ValueString").Call(),
		j.Id(r.Name): j.Op("&").Qual(r.PkgImportPath, r.Name).Values(
			// Filter out the top level or terraform only attributes
			buildInnerResource(lo.OmitByKeys(r.Schema.Attributes, []string{"parent", "id", "name"}), []string{"model"}, r.PkgImportPath),
		),
	}
	return res
}

// func buildResource(s map[string]schema.Attribute, prefixes []string, r resource) j.Dict {
// 	res := j.Dict{}
// 	for k, v := range s {
// 		fullyQualified := append(prefixes, snakeToCamel(k))
// 		modelPath := j.Id("model").Add(lo.Map(fullyQualified, func(s string, _ int) j.Code { return j.Dot(s) })...)

// 		switch t := v.(type) {
// 		// TODO: handle non-string fields
// 		case schema.StringAttribute:
// 			if t.CustomType == nil {
// 				res[j.Id(snakeToCamel(k))] = modelPath.Dot("ValueString").Call()
// 				continue
// 			}
// 			// If it's a string but has a custom type it's an enum.
// 			// This code works if the enum is in the same package as the resource, but we may need to pass
// 			// more information to the struct at schema generation time in the future.
// 			res[j.Id(snakeToCamel(k))] = j.Qual(r.PkgImportPath, snakeToCamel(k)).Call(
// 				j.Qual(r.PkgImportPath, snakeToCamel(k)+"_value").Index(modelPath.Dot("ValueString").Call()),
// 			)
// 		case schema.NestedAttribute:
// 			// underlying := t.(schema.NestedAttribute).GetNestedObject().GetAttributes()
// 			desc := strings.Split(t.GetDescription(), "!!")
// 			gotype := desc[len(desc)-1]
// 			res[j.Id(snakeToCamel(k))] = j.Op("&").Qual(r.PkgImportPath, gotype).Block()
// 			continue
// 		default:
// 			// panic(fmt.Sprintf("%v isn't supported as a resource value", t))
// 			fmt.Printf("%v isn't supported as a resource value", t)
// 		}
// 	}
// 	return res
// }

func buildInnerResource(s map[string]schema.Attribute, prefixes []string, pkgImportPath string) j.Dict {
	res := j.Dict{}
	for k, v := range s {
		fullyQualified := append(prefixes, snakeToCamel(k))
		modelPath := j.Id(fullyQualified[0]).Add(lo.Map(fullyQualified[1:], func(s string, _ int) j.Code { return j.Dot(s) })...)

		switch t := v.(type) {
		// TODO: handle non-string fields
		case schema.StringAttribute:
			if t.CustomType == nil {
				res[j.Id(snakeToCamel(k))] = modelPath.Dot("ValueString").Call()
				continue
			}
			// If it's a string but has a custom type it's an enum.
			// This code works if the enum is in the same package as the resource, but we may need to pass
			// more information to the struct at schema generation time in the future.
			res[j.Id(snakeToCamel(k))] = j.Qual(pkgImportPath, snakeToCamel(k)).Call(
				j.Qual(pkgImportPath, snakeToCamel(k)+"_value").Index(modelPath.Dot("ValueString").Call()),
			)
		case schema.ListNestedAttribute:
			underlying := make(map[string]schema.Attribute)
			for k, v := range t.GetNestedObject().GetAttributes() {
				underlying[k] = v
			}
			desc := strings.Split(t.GetDescription(), "!!")
			gotype := desc[len(desc)-1]
			goimport := desc[len(desc)-2]

			// res[j.Id(snakeToCamel(k))] = j.Qual(SamberLo, "Map").Call(modelPath, j.Func().Call(
			// 	j.Id(k).Interface(), j.Id("index").Int(),
			// ).Op("*").Qual(goimport, gotype).Block(j.Return(
			// 	j.Op("&").Qual(goimport, gotype).Values(),
			// )))
			res[j.Id(snakeToCamel(k))] = j.Func().Call().Op("[]*").Qual(goimport, gotype).Block(
				j.Id("a").Op(":=").Make(j.Op("[]*").Qual(goimport, gotype), j.Lit(0)),
				j.For(j.Id("_").Op(",").Id(k).Op(":=").Range().Add(modelPath)).Block(
					j.Id("a").Op("=").Append(j.Id("a"), j.Op("&").Qual(goimport, gotype).Values(
						buildInnerResource(underlying, []string{k}, pkgImportPath),
					)),
				),
				j.Return(j.Id("a")),
			).Call()
			// j.Op("&").Qual(goimport, gotype).Values(
			// 	buildInnerResource(underlying, []string{k}, pkgImportPath),
			// ),
			// )
			continue

		case schema.NestedAttribute:
			underlying := make(map[string]schema.Attribute)
			for k, v := range t.GetNestedObject().GetAttributes() {
				underlying[k] = v
			}
			desc := strings.Split(t.GetDescription(), "!!")
			gotype := desc[len(desc)-1]
			goimport := desc[len(desc)-2]

			res[j.Id(snakeToCamel(k))] = j.Op("&").Qual(goimport, gotype).Values(
				buildInnerResource(underlying, append(prefixes, snakeToCamel(k)), pkgImportPath),
			)
			continue
		default:
			// panic(fmt.Sprintf("%v isn't supported as a resource value", t))
			fmt.Printf("%v isn't supported as a resource value", t)
		}

	}
	return res
}
