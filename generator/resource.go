package main

import (
	j "github.com/dave/jennifer/jen"
)

func genResource(r resource) *j.File {
	f := j.NewFile(r.lowerName)

	f.Comment("// Ensure provider defined types fully satisfy framework interfaces")
	f.Var().Id("_").Qual(Resource, "Resource").Op("=").Op("&").Add(r.structId).Values()
	f.Var().Id("_").Qual(Resource, "ResourceWithImportState").Op("=").Op("&").Add(r.structId).Values()

	f.Comment("// Constructor")
	f.Func().Id("NewResource").Params().Params(j.Qual(Resource, "Resource")).Block(j.Return().Op("&").Add(r.structId).Block())

	f.Comment("// Resource struct definition")
	f.Type().Add(r.structId).Struct(j.Id("client").Qual(r.PkgImportPath, r.Client))

	f.Comment("// Implement metadata function from Terraform resource interface")
	f.Func().
		Parens(j.Id("r").Op("*").Add(r.structId)).Id("Metadata").
		Params(j.Op("_").Qual("context", "Context"), j.Id("req").Qual(Resource, "MetadataRequest"), j.Id("resp").Op("*").Qual(Resource, "MetadataResponse")).
		Block(
			j.Id("resp").Dot("TypeName").Op("=").Id("req").Dot("ProviderTypeName").Op("+").Lit("_" + r.lowerName),
		)

	f.Comment("// Implement configure function from Terraform resource interface")
	f.Func().
		Parens(j.Id("r").Op("*").Add(r.structId)).Id("Configure").
		Params(j.Op("_").Qual("context", "Context"), j.Id("req").Qual(Resource, "ConfigureRequest"), j.Id("resp").Op("*").Qual(Resource, "ConfigureResponse")).
		Block(
			j.Id("clients").Op(":=").Qual(Helpers, "BuildClientsResource").Call(j.Id("req"), j.Id("resp")),
			j.If(j.Id("resp").Dot("Diagnostics").Dot("HasError").Call().Op("||").Id("clients").Op("==").Nil().Block(j.Return())),
			j.Id("r").Dot("client").Op("=").Id("clients").Dot(r.Name),
		)

	f.Comment("// Implement schema function from Terraform resource interface")
	f.Func().
		Parens(j.Op("*").Add(r.structId)).Id("Schema").
		Params(j.Op("_").Qual("context", "Context"), j.Op("_").Qual(Resource, "SchemaRequest"), j.Id("resp").Op("*").Qual(Resource, "SchemaResponse")).
		Block(
			j.Id("resp").Dot("Schema").Op("=").Qual(r.PkgImportPath, r.Name+"Schema").Call(),
		)

	return f
}
