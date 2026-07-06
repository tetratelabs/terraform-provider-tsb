// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package main

import (
	"reflect"

	j "github.com/dave/jennifer/jen"
)

// emitResource writes the resource.Resource implementation for one resource.
func (g *generator) emitResource(f *j.File, r resourceInfo) error {
	// resBase identifies the resource (type, constructor, collection); modelBase
	// identifies its model/expand/flatten. They differ only for shared-message
	// resources (Istio direct mode), where many resources reuse one model.
	resBase := r.BaseName
	modelBase := g.baseName(r.MessageGoType)
	recv := "r"
	typeName := lowerFirst(resBase) + "Resource"
	model := func() *j.Statement { return j.Id(modelBase + "Model") }
	clientExpr := j.Qual(r.ServicePkg, r.ClientCtor).Call(j.Id(recv).Dot("client").Dot("conn"))

	// Create-only fields (server-set secrets) are captured from prior state before
	// flatten and restored after, so reads/updates never overwrite the create-time
	// value (the server stops returning it after creation).
	coFields, err := g.createOnlyFieldNames(r)
	if err != nil {
		return err
	}
	coCapture := func(src string) []j.Code {
		out := make([]j.Code, 0, len(coFields))
		for _, fn := range coFields {
			out = append(out, j.Id("co"+fn).Op(":=").Id(src).Dot(fn))
		}
		return out
	}
	coRestore := func(dst string) []j.Code {
		out := make([]j.Code, 0, len(coFields))
		for _, fn := range coFields {
			out = append(out, j.Id(dst).Dot(fn).Op("=").Id("co"+fn))
		}
		return out
	}

	f.Const().Id(lowerFirst(resBase) + "Collection").Op("=").Lit(r.Collection)

	// Interface assertions.
	f.Var().Defs(
		j.Op("_").Qual(pkgResource, "Resource").Op("=").Parens(j.Op("*").Id(typeName)).Parens(j.Nil()),
		j.Op("_").Qual(pkgResource, "ResourceWithConfigure").Op("=").Parens(j.Op("*").Id(typeName)).Parens(j.Nil()),
		j.Op("_").Qual(pkgResource, "ResourceWithImportState").Op("=").Parens(j.Op("*").Id(typeName)).Parens(j.Nil()),
	)

	f.Commentf("New%sResource constructs the tsb_%s resource.", resBase, r.TFName)
	f.Func().Id("New"+resBase+"Resource").Params().Qual(pkgResource, "Resource").Block(
		j.Return(j.Op("&").Id(typeName).Values()),
	)

	f.Type().Id(typeName).Struct(
		j.Id("client").Op("*").Id("client"),
	)

	// Metadata
	f.Func().Params(j.Id(recv).Op("*").Id(typeName)).Id("Metadata").
		Params(j.Id("_").Qual(pkgContext, "Context"), j.Id("req").Qual(pkgResource, "MetadataRequest"), j.Id("resp").Op("*").Qual(pkgResource, "MetadataResponse")).
		Block(j.Id("resp").Dot("TypeName").Op("=").Id("req").Dot("ProviderTypeName").Op("+").Lit("_" + r.TFName))

	// Schema
	f.Func().Params(j.Id(recv).Op("*").Id(typeName)).Id("Schema").
		Params(j.Id("_").Qual(pkgContext, "Context"), j.Id("_").Qual(pkgResource, "SchemaRequest"), j.Id("resp").Op("*").Qual(pkgResource, "SchemaResponse")).
		Block(j.Id("resp").Dot("Schema").Op("=").Id("fixupSchema").Call(j.Qual(r.MessageGoType.PkgPath(), r.MessageGoType.Name()+"Schema").Call()))

	// Configure
	f.Func().Params(j.Id(recv).Op("*").Id(typeName)).Id("Configure").
		Params(j.Id("_").Qual(pkgContext, "Context"), j.Id("req").Qual(pkgResource, "ConfigureRequest"), j.Id("resp").Op("*").Qual(pkgResource, "ConfigureResponse")).
		Block(
			j.If(j.Id("req").Dot("ProviderData").Op("==").Nil()).Block(j.Return()),
			j.List(j.Id("c"), j.Id("ok")).Op(":=").Id("req").Dot("ProviderData").Assert(j.Op("*").Id("client")),
			j.If(j.Op("!").Id("ok")).Block(
				j.Id("resp").Dot("Diagnostics").Dot("AddError").Call(
					j.Lit("Unexpected provider data"),
					j.Qual(pkgFmt, "Sprintf").Call(j.Lit("Expected *client, got %T. This is a provider bug."), j.Id("req").Dot("ProviderData")),
				),
				j.Return(),
			),
			j.Id(recv).Dot("client").Op("=").Id("c"),
		)

	// Create
	createReqFields := j.Dict{}
	if r.CreateHasParent {
		createReqFields[j.Id("Parent")] = j.Id("plan").Dot("Parent").Dot("ValueString").Call()
	}
	if r.CreateHasName {
		createReqFields[j.Id("Name")] = j.Id("plan").Dot("Name").Dot("ValueString").Call()
	}
	createReqFields[j.Id(r.CreateMsgField)] = j.Id("msg")
	f.Func().Params(j.Id(recv).Op("*").Id(typeName)).Id("Create").
		Params(j.Id("ctx").Qual(pkgContext, "Context"), j.Id("req").Qual(pkgResource, "CreateRequest"), j.Id("resp").Op("*").Qual(pkgResource, "CreateResponse")).
		Block(
			j.Var().Id("plan").Add(model()),
			appendDiags(j.Id("req").Dot("Plan").Dot("Get").Call(j.Id("ctx"), j.Op("&").Id("plan"))),
			returnIfError(),
			j.List(j.Id("msg"), j.Id("d")).Op(":=").Id("expand"+modelBase).Call(j.Id("ctx"), j.Op("&").Id("plan")),
			j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("d").Op("...")),
			returnIfError(),
			j.Id("c").Op(":=").Add(clientExpr),
			j.List(j.Id("created"), j.Id("err")).Op(":=").Id("c").Dot(r.CreateMethod).Call(
				j.Id("ctx"),
				j.Op("&").Add(msgReqQual(r, r.CreateReqGoType)).Values(createReqFields),
			),
			ifErrAddError("Unable to create "+resBase),
			j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("flatten"+modelBase).Call(j.Id("ctx"), j.Id("created"), j.Op("&").Id("plan")).Op("...")),
			j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("resp").Dot("State").Dot("Set").Call(j.Id("ctx"), j.Op("&").Id("plan")).Op("...")),
		)

	// Read
	readBody := []j.Code{
		j.Var().Id("state").Add(model()),
		appendDiags(j.Id("req").Dot("State").Dot("Get").Call(j.Id("ctx"), j.Op("&").Id("state"))),
		returnIfError(),
	}
	readBody = append(readBody, coCapture("state")...)
	readBody = append(readBody,
		j.Id("c").Op(":=").Add(clientExpr),
		j.List(j.Id("got"), j.Id("err")).Op(":=").Id("c").Dot(r.GetMethod).Call(
			j.Id("ctx"),
			j.Op("&").Add(msgReqQual(r, r.GetReqGoType)).Values(j.Dict{j.Id("Fqn"): j.Id("state").Dot("ID").Dot("ValueString").Call()}),
		),
		j.If(j.Id("err").Op("!=").Nil()).Block(
			j.If(j.Qual(pkgStatus, "Code").Call(j.Id("err")).Op("==").Qual(pkgCodes, "NotFound")).Block(
				j.Id("resp").Dot("State").Dot("RemoveResource").Call(j.Id("ctx")),
				j.Return(),
			),
			j.Id("resp").Dot("Diagnostics").Dot("AddError").Call(j.Lit("Unable to read "+resBase), j.Id("err").Dot("Error").Call()),
			j.Return(),
		),
		j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("flatten"+modelBase).Call(j.Id("ctx"), j.Id("got"), j.Op("&").Id("state")).Op("...")),
	)
	readBody = append(readBody, coRestore("state")...)
	readBody = append(readBody,
		j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("resp").Dot("State").Dot("Set").Call(j.Id("ctx"), j.Op("&").Id("state")).Op("...")),
	)
	f.Func().Params(j.Id(recv).Op("*").Id(typeName)).Id("Read").
		Params(j.Id("ctx").Qual(pkgContext, "Context"), j.Id("req").Qual(pkgResource, "ReadRequest"), j.Id("resp").Op("*").Qual(pkgResource, "ReadResponse")).
		Block(readBody...)

	// Update
	updateBody := []j.Code{
		j.Var().Id("plan").Add(model()),
		appendDiags(j.Id("req").Dot("Plan").Dot("Get").Call(j.Id("ctx"), j.Op("&").Id("plan"))),
		returnIfError(),
	}
	updateBody = append(updateBody, coCapture("plan")...)
	updateBody = append(updateBody,
		j.List(j.Id("msg"), j.Id("d")).Op(":=").Id("expand"+modelBase).Call(j.Id("ctx"), j.Op("&").Id("plan")),
		j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("d").Op("...")),
		returnIfError(),
		j.Id("c").Op(":=").Add(clientExpr),
		j.Comment("Update takes the full resource including the server-managed etag, which the"),
		j.Comment("Terraform model intentionally drops; fetch the current one first."),
		j.List(j.Id("current"), j.Id("err")).Op(":=").Id("c").Dot(r.GetMethod).Call(
			j.Id("ctx"),
			j.Op("&").Add(msgReqQual(r, r.GetReqGoType)).Values(j.Dict{j.Id("Fqn"): j.Id("plan").Dot("ID").Dot("ValueString").Call()}),
		),
		ifErrAddError("Unable to read "+resBase+" before update"),
		j.Id("msg").Dot("Fqn").Op("=").Id("current").Dot("Fqn"),
		j.Id("msg").Dot("Etag").Op("=").Id("current").Dot("Etag"),
		j.List(j.Id("updated"), j.Id("err2")).Op(":=").Id("c").Dot(r.UpdateMethod).Call(j.Id("ctx"), j.Id("msg")),
		j.If(j.Id("err2").Op("!=").Nil()).Block(
			j.Id("resp").Dot("Diagnostics").Dot("AddError").Call(j.Lit("Unable to update "+resBase), j.Id("err2").Dot("Error").Call()),
			j.Return(),
		),
		j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("flatten"+modelBase).Call(j.Id("ctx"), j.Id("updated"), j.Op("&").Id("plan")).Op("...")),
	)
	updateBody = append(updateBody, coRestore("plan")...)
	updateBody = append(updateBody,
		j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("resp").Dot("State").Dot("Set").Call(j.Id("ctx"), j.Op("&").Id("plan")).Op("...")),
	)
	f.Func().Params(j.Id(recv).Op("*").Id(typeName)).Id("Update").
		Params(j.Id("ctx").Qual(pkgContext, "Context"), j.Id("req").Qual(pkgResource, "UpdateRequest"), j.Id("resp").Op("*").Qual(pkgResource, "UpdateResponse")).
		Block(updateBody...)

	// Delete
	delDict := j.Dict{j.Id("Fqn"): j.Id("state").Dot("ID").Dot("ValueString").Call()}
	f.Func().Params(j.Id(recv).Op("*").Id(typeName)).Id("Delete").
		Params(j.Id("ctx").Qual(pkgContext, "Context"), j.Id("req").Qual(pkgResource, "DeleteRequest"), j.Id("resp").Op("*").Qual(pkgResource, "DeleteResponse")).
		Block(
			j.Var().Id("state").Add(model()),
			appendDiags(j.Id("req").Dot("State").Dot("Get").Call(j.Id("ctx"), j.Op("&").Id("state"))),
			returnIfError(),
			j.Id("c").Op(":=").Add(clientExpr),
			j.List(j.Id("_"), j.Id("err")).Op(":=").Id("c").Dot(r.DeleteMethod).Call(
				j.Id("ctx"),
				j.Op("&").Add(msgReqQual(r, r.DeleteReqGoType)).Values(delDict),
			),
			j.If(j.Id("err").Op("!=").Nil()).Block(
				j.If(j.Qual(pkgStatus, "Code").Call(j.Id("err")).Op("==").Qual(pkgCodes, "NotFound")).Block(j.Return()),
				j.Id("resp").Dot("Diagnostics").Dot("AddError").Call(j.Lit("Unable to delete "+resBase), j.Id("err").Dot("Error").Call()),
				j.Return(),
			),
		)

	// ImportState
	f.Func().Params(j.Id(recv).Op("*").Id(typeName)).Id("ImportState").
		Params(j.Id("ctx").Qual(pkgContext, "Context"), j.Id("req").Qual(pkgResource, "ImportStateRequest"), j.Id("resp").Op("*").Qual(pkgResource, "ImportStateResponse")).
		Block(
			j.Id("fqn").Op(":=").Id("req").Dot("ID"),
			j.List(j.Id("parent"), j.Id("name")).Op(":=").Id("fqnParentName").Call(j.Id("fqn"), j.Id(lowerFirst(resBase)+"Collection")),
			j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("resp").Dot("State").Dot("SetAttribute").Call(j.Id("ctx"), j.Qual(pkgPath, "Root").Call(j.Lit("id")), j.Id("fqn")).Op("...")),
			j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("resp").Dot("State").Dot("SetAttribute").Call(j.Id("ctx"), j.Qual(pkgPath, "Root").Call(j.Lit("parent")), j.Id("parent")).Op("...")),
			j.Id("resp").Dot("Diagnostics").Dot("Append").Call(j.Id("resp").Dot("State").Dot("SetAttribute").Call(j.Id("ctx"), j.Qual(pkgPath, "Root").Call(j.Lit("name")), j.Id("name")).Op("...")),
		)

	return nil
}

// createOnlyFieldNames returns the model (Go) field names of a resource's
// top-level create-only fields, which the provider preserves across read/update.
func (g *generator) createOnlyFieldNames(r resourceInfo) ([]string, error) {
	plans, err := g.fieldPlans(r.MessageGoType, r.MessageDesc)
	if err != nil {
		return nil, err
	}
	var names []string
	for _, p := range plans {
		if isCreateOnly(p.FD) {
			names = append(names, p.GoName)
		}
	}
	return names, nil
}

// msgReqQual qualifies a request type by its Go import path and name.
func msgReqQual(_ resourceInfo, reqType reflect.Type) *j.Statement {
	return j.Qual(reqType.PkgPath(), reqType.Name())
}

func appendDiags(call *j.Statement) j.Code {
	return j.Id("resp").Dot("Diagnostics").Dot("Append").Call(call.Op("..."))
}

func returnIfError() j.Code {
	return j.If(j.Id("resp").Dot("Diagnostics").Dot("HasError").Call()).Block(j.Return())
}

func ifErrAddError(msg string) j.Code {
	return j.If(j.Id("err").Op("!=").Nil()).Block(
		j.Id("resp").Dot("Diagnostics").Dot("AddError").Call(j.Lit(msg), j.Id("err").Dot("Error").Call()),
		j.Return(),
	)
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	return string(s[0]|0x20) + s[1:]
}
