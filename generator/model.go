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
	"github.com/samber/lo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func genModel(r resource) *j.File {
	f := j.NewFile(r.lowerName)

	fields := lo.MapToSlice(
		lo.MapValues(r.Schema.Attributes, func(attribute schema.Attribute, name string) j.Code {
			return j.Id(snakeToCamel(name)).Qual(Types, "String").Tag(map[string]string{"tfsdk": name})
		}),
		func(_ string, code j.Code) j.Code { return code },
	)

	f.Comment("// tfsdk typed model definition")
	f.Type().Add(r.modelId).Struct(fields...)

	return f
}

func snakeToCamel(s string) string {
	words := strings.Split(s, "_")
	caser := cases.Title(language.English)
	for i := range words {
		words[i] = caser.String(words[i])
	}
	return strings.Join(words, "")
}
