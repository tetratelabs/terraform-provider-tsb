package main

import (
	j "github.com/dave/jennifer/jen"
)

func genImport(r resource) *j.File {
	f := j.NewFile(r.lowerName)

	f.Func().
		Parens(j.Id("r").Op("*").Add(r.structId)).Id("ImportState").
		Params(j.Id("ctx").Qual("context", "Context"), j.Id("req").Qual(Resource, "ImportStateRequest"), j.Id("resp").Op("*").Qual(Resource, "ImportStateResponse")).
		Block(
			j.Qual(Resource, "ImportStatePassthroughID").Call(j.Id("ctx"), j.Qual(Path, "Root").Call(j.Lit("id")), j.Id("req"), j.Id("resp")),
		)

	return f
}
