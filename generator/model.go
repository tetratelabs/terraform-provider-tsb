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
	"context"
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
	genStruct(f, r.modelId, r.Schema.Attributes, "")

	// f.Type().Add(r.modelId).Struct(fields...)

	return f
}

func genStruct(f *j.File, structName j.Code, attributes map[string]schema.Attribute, suffix string) j.Code {
	fields := lo.MapToSlice(
		lo.MapValues(attributes, func(attribute schema.Attribute, fieldName string) j.Code {
			// nestedModelName := j.Id(fieldName)
			fieldId := j.Id(snakeToCamel(fieldName))
			tag := j.Tag(map[string]string{"tfsdk": fieldName})

			switch attribute.(type) {
			// case schema.ListAttribute:
			// 	return fieldId.Add(j.Qual(Types, "List")).Add(tag)
			case schema.ListNestedAttribute:
				newStruct := genStruct(f, j.Id(snakeToCamel(fieldName+"_"+suffix+"_Model")), map[string]schema.Attribute{}, suffix+"_"+fieldName)
				return fieldId.Add(j.List().Add(newStruct)).Add(tag)
			case schema.ListAttribute:
				return fieldId.Qual(Types, "List").Add(tag)
			case schema.MapAttribute:
				newStruct := genStruct(f, j.Id(snakeToCamel(fieldName+"_"+suffix+"_Model")), map[string]schema.Attribute{}, suffix+"_"+fieldName)
				return fieldId.Add(newStruct).Add(tag)
			case schema.NestedAttribute:
				underlying := attribute.(schema.NestedAttribute).GetNestedObject().GetAttributes()
				// asdf := lo.MapEntries(underlying, func(k string, v fwschema.Attribute) (string, schema.Attribute) {
				// 	return k, v.(schema.Attribute)
				// })
				asdf := make(map[string]schema.Attribute)
				for k, v := range underlying {
					asdf[k] = v
				}

				newStruct := genStruct(f, j.Id(snakeToCamel(fieldName+"_"+suffix+"_Model")), asdf, suffix+"_"+fieldName)
				return fieldId.Add(newStruct).Add(tag)
			default:
				return fieldId.Add(attrToType(attribute)).Add(tag)
			}

		}),
		func(_ string, code j.Code) j.Code { return code },
	)

	f.Type().Add(structName).Struct(fields...)
	return structName
}

func isPrimitive(attribute schema.Attribute) bool {
	switch attribute.(type) {
	case schema.NestedAttribute, schema.MapAttribute, schema.ListAttribute:
		return false
	default:
		return true
	}
}

func snakeToCamel(s string) string {
	words := strings.Split(s, "_")
	caser := cases.Title(language.English)
	for i := range words {
		words[i] = caser.String(words[i])
	}
	return strings.Join(words, "")
}

// var attrTypeMap = map[]string{
// 	basetypes.StringType{}.String(): "String",
// 	// basetypes.
// }

// func asdf[T attr.Type](thing T) {

// }

func attrToType(attr schema.Attribute) j.Code {
	return j.Qual(Types, strings.Split(attr.GetType().TerraformType(context.Background()).String(), ".")[1])

	// return attrTypeMap[attr.GetType().String()]
}
