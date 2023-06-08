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
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func genModel(r resource) *j.File {
	f := j.NewFile(r.lowerName)

	f.Comment("// tfsdk typed model definition")
	genStruct(f, r.modelId, r.Schema.Attributes, []string{})

	// f.Type().Add(r.modelId).Struct(fields...)

	return f
}

func genStruct(f *j.File, structName string, attributes map[string]schema.Attribute, prefix []string) string {
	fields := lo.MapToSlice(
		lo.MapValues(attributes, func(attribute schema.Attribute, fieldName string) j.Code {
			fieldId := j.Id(snakeToCamel(fieldName))
			tag := j.Tag(map[string]string{"tfsdk": fieldName})

			switch attribute.(type) {
			case schema.ListNestedAttribute:
				tmp := make(map[string]schema.Attribute)
				underlying := attribute.(schema.NestedAttribute).GetNestedObject().GetAttributes()
				for k, v := range underlying {
					tmp[k] = v
				}
				newStruct := genStruct(f, snakeSliceToCamel(append(prefix, fieldName, "Model")), tmp, append(prefix, fieldName))
				// return fieldId.Add(j.Op("[]*").Id(newStruct)).Add(tag)
				return fieldId.Qual(Types, "List").Add(tag).Comment(newStruct)
			case schema.ListAttribute:
				return fieldId.Qual(Types, "List").Add(tag)
			case schema.MapAttribute:
				return fieldId.Qual(Types, "Map").Add(tag)
			case schema.NestedAttribute:
				underlying := attribute.(schema.NestedAttribute).GetNestedObject().GetAttributes()
				tmp := make(map[string]schema.Attribute)
				for k, v := range underlying {
					tmp[k] = v
				}

				newStruct := genStruct(f, snakeSliceToCamel(append(prefix, fieldName, "Model")), tmp, append(prefix, fieldName))
				return fieldId.Id(newStruct).Add(tag)
			default:
				return fieldId.Add(attrToType(attribute)).Add(tag)
			}

		}),
		func(_ string, code j.Code) j.Code { return code },
	)

	f.Commentf("%v", strings.Join(prefix, "/"))
	f.Type().Id(structName).Struct(fields...)
	return structName
}

func snakeToCamel(s string) string {
	// exceptions
	switch s {
	case "cni_overlays":
		return "CniOverlays"
	case "token_ttl":
		return "TokenTtl"
	}
	words := strings.Split(s, "_")
	caser := cases.Title(language.English)
	caps := cases.Upper(language.English)
	for i := range words {
		switch words[i] {
		// Acronyms...
		case "jwk", "fqn", "ttl", "cni":
			words[i] = caps.String(words[i])
			if i > 0 {
				words[i] = "_" + words[i]
			}
		default:
			words[i] = caser.String(words[i])
		}
	}
	return strings.Join(words, "")
}

func snakeSliceToCamel(s []string) string {
	return strings.Join(lo.Map(s, func(x string, _ int) string { return snakeToCamel(x) }), "_")
}

func attrToType(attr schema.Attribute) j.Code {
	var attrType string
	switch attr.(type) {
	case schema.BoolAttribute:
		attrType = "Bool"
	case schema.StringAttribute:
		attrType = "String"
	case schema.Float64Attribute:
		attrType = "Float64"
	case schema.Int64Attribute:
		attrType = "Int64"
	default:
		panic(fmt.Sprintf("Unknown primitive attribute type %v\n", attr))
	}
	return j.Qual(Types, attrType)
}
