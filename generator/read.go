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
		res = append(res, j.Id("model").Dot(snakeToCamel(k)).Op("=").Add(modelField(k, v, r)))
	}
	return res
}

func modelField(name string, attr schema.Attribute, r resource) j.Code {
	switch t := attr.(type) {
	case schema.StringAttribute:
		if t.CustomType == nil {
			return j.Qual(Types, "StringValue").Call(j.Id(r.lowerName).Dot(snakeToCamel(name)))
		}
		// If it's a string but has a custom type it's an enum.
		// This code works if the enum is in the same package as the resource, but we may need to pass
		// more information to the struct at schema generation time in the future.
		return j.Qual(Types, "StringValue").Call(j.Qual(r.PkgImportPath, snakeToCamel(name)+"_name").Index(j.Int32().Call(j.Id(r.lowerName).Dot(snakeToCamel(name)))))
	case schema.ListNestedAttribute:
		return nil
	default:
		panic(fmt.Sprintf("%v isn't supported as a resource value", t))
	}
}
