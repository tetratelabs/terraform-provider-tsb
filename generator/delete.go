package main

import (
	"github.com/dave/jennifer/jen"
	j "github.com/dave/jennifer/jen"
)

func genDelete(r resource) *j.File {
	f := j.NewFile(r.lowerName)

	f.Func().
		Parens(j.Id("r").Op("*").Add(r.structId)).Id("Delete").
		Params(j.Id("ctx").Qual("context", "Context"), j.Id("req").Qual(Resource, "DeleteRequest"), j.Id("resp").Op("*").Qual(Resource, "DeleteResponse")).
		Block(
			j.Var().Id("model").Add(r.modelId),
			j.Id("resp").Dot("Diagnostics").Dot("Append").Call(
				j.Id("req").Dot("State").Dot("Get").Call(j.Id("ctx"), jen.Op("&").Id("model")).Op("..."),
			),
			j.If(j.Id("resp").Dot("Diagnostics").Dot("HasError").Call()).Block(j.Return()),
			j.If(
				j.List(j.Id("_"), j.Id("err")).Op(":=").Id("r").Dot("client").Dot("Delete"+r.Name).Call(
					j.Id("ctx"),
					j.Op("&").Qual(r.PkgImportPath, "Delete"+r.Name+"Request").Values(
						j.Dict{j.Id("Fqn"): jen.Id("model").Dot("Id").Dot("ValueString").Call()},
					),
				),
				j.Err().Op("!=").Nil(),
			).Block(
				j.Id("resp").Dot("Diagnostics").Dot("AddError").Call(
					j.Lit("Error deleting "+r.Name),
					j.Lit("Delete"+r.Name+" request failed: ").Op("+").Err().Dot("Error").Call(),
				),
				j.Return(),
			),
		)

	return f
}
