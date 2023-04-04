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
