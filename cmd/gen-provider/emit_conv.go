// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package main

import (
	"fmt"
	"reflect"

	j "github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// emitExpand writes expand<Base>(ctx, *Model) (*proto, diag.Diagnostics).
func (g *generator) emitExpand(f *j.File, t reflect.Type, base string, plans []fieldPlan, _ bool) error {
	body := []j.Code{
		j.Var().Id("diags").Qual(pkgDiag, "Diagnostics"),
		j.If(j.Id("m").Op("==").Nil()).Block(j.Return(j.Nil(), j.Id("diags"))),
		j.Id("out").Op(":=").Op("&").Qual(t.PkgPath(), t.Name()).Values(),
	}
	for _, p := range plans {
		code, err := g.expandField(p)
		if err != nil {
			return fmt.Errorf("expand field %q: %w", p.Key, err)
		}
		body = append(body, code...)
	}
	body = append(body, j.Return(j.Id("out"), j.Id("diags")))

	f.Commentf("expand%s converts the Terraform model into a %s proto.", base, t.Name())
	f.Func().Id("expand"+base).
		Params(j.Id("ctx").Qual(pkgContext, "Context"), j.Id("m").Op("*").Id(base+"Model")).
		Params(j.Op("*").Qual(t.PkgPath(), t.Name()), j.Qual(pkgDiag, "Diagnostics")).
		Block(body...)
	return nil
}

// emitFlatten writes flatten<Base>(ctx, *proto, *Model) diag.Diagnostics.
func (g *generator) emitFlatten(f *j.File, t reflect.Type, base string, plans []fieldPlan, isResource bool) error {
	body := []j.Code{j.Var().Id("diags").Qual(pkgDiag, "Diagnostics")}
	if isResource {
		body = append(body, j.Id("m").Dot("ID").Op("=").Qual(pkgTypes, "StringValue").Call(j.Id("p").Dot("Fqn")))
	}
	for _, p := range plans {
		code, err := g.flattenField(p)
		if err != nil {
			return fmt.Errorf("flatten field %q: %w", p.Key, err)
		}
		body = append(body, code...)
	}
	body = append(body, j.Return(j.Id("diags")))

	f.Commentf("flatten%s writes a %s proto into the Terraform model.", base, t.Name())
	f.Func().Id("flatten"+base).
		Params(
			j.Id("ctx").Qual(pkgContext, "Context"),
			j.Id("p").Op("*").Qual(t.PkgPath(), t.Name()),
			j.Id("m").Op("*").Id(base+"Model"),
		).
		Params(j.Qual(pkgDiag, "Diagnostics")).
		Block(body...)
	return nil
}

func (g *generator) expandField(p fieldPlan) ([]j.Code, error) {
	if p.Oneof != nil {
		return g.expandOneof(p)
	}
	// Create-only secrets are output-only: never sent back to the server.
	if isCreateOnly(p.FD) {
		return nil, nil
	}
	fd := p.FD
	mDot := j.Id("m").Dot(p.GoName)
	outDot := j.Id("out").Dot(p.GoName)

	guard := j.Op("!").Add(mDot.Clone()).Dot("IsNull").Call().Op("&&").Op("!").Add(mDot.Clone()).Dot("IsUnknown").Call()

	switch {
	case fd.IsMap():
		if isMessageKind(fd.MapValue()) {
			nested := p.GoType.Elem().Elem() // map[string]*Msg -> Msg
			if err := g.enqueue(nested); err != nil {
				return nil, err
			}
			nb := g.baseName(nested)
			return []j.Code{
				j.If(mDot.Clone().Op("!=").Nil()).Block(
					outDot.Clone().Op("=").Make(j.Map(j.String()).Op("*").Qual(nested.PkgPath(), nested.Name()), j.Len(mDot.Clone())),
					j.For(j.List(j.Id("k"), j.Id("ev")).Op(":=").Range().Add(mDot.Clone())).Block(
						j.Id("ev").Op(":=").Id("ev"),
						j.List(j.Id("v"), j.Id("d")).Op(":=").Id("expand"+nb).Call(j.Id("ctx"), j.Op("&").Id("ev")),
						j.Id("diags").Dot("Append").Call(j.Id("d").Op("...")),
						outDot.Clone().Index(j.Id("k")).Op("=").Id("v"),
					),
				),
			}, nil
		}
		if fd.MapValue().Kind() != protoreflect.StringKind {
			return nil, fmt.Errorf("map with %s value is not yet supported", fd.MapValue().Kind())
		}
		return []j.Code{
			j.If(guard).Block(
				j.Id("diags").Dot("Append").Call(mDot.Clone().Dot("ElementsAs").Call(j.Id("ctx"), j.Op("&").Add(outDot.Clone()), j.False()).Op("...")),
			),
		}, nil

	case fd.IsList():
		if isMessageKind(fd) {
			nested := p.GoType.Elem().Elem() // []*Msg -> Msg
			if err := g.enqueue(nested); err != nil {
				return nil, err
			}
			nb := g.baseName(nested)
			return []j.Code{
				j.For(j.Id("i").Op(":=").Range().Add(mDot.Clone())).Block(
					j.List(j.Id("ev"), j.Id("d")).Op(":=").Id("expand"+nb).Call(j.Id("ctx"), j.Op("&").Add(mDot.Clone()).Index(j.Id("i"))),
					j.Id("diags").Dot("Append").Call(j.Id("d").Op("...")),
					j.If(j.Id("ev").Op("!=").Nil()).Block(
						outDot.Clone().Op("=").Append(outDot.Clone(), j.Id("ev")),
					),
				),
			}, nil
		}
		elem := p.GoType.Elem() // proto element Go type
		switch {
		case fd.Kind() == protoreflect.StringKind:
			return []j.Code{
				j.If(guard).Block(
					j.Id("diags").Dot("Append").Call(mDot.Clone().Dot("ElementsAs").Call(j.Id("ctx"), j.Op("&").Add(outDot.Clone()), j.False()).Op("...")),
				),
			}, nil
		case fd.Kind() == protoreflect.EnumKind:
			pkg, name := elem.PkgPath(), elem.Name()
			return []j.Code{
				j.If(guard).Block(
					j.Var().Id("tmp").Index().String(),
					j.Id("diags").Dot("Append").Call(mDot.Clone().Dot("ElementsAs").Call(j.Id("ctx"), j.Op("&").Id("tmp"), j.False()).Op("...")),
					outDot.Clone().Op("=").Make(j.Index().Qual(pkg, name), j.Len(j.Id("tmp"))),
					j.For(j.Id("i").Op(":=").Range().Id("tmp")).Block(
						outDot.Clone().Index(j.Id("i")).Op("=").Qual(pkg, name).Call(j.Qual(pkg, name+"_value").Index(j.Id("tmp").Index(j.Id("i")))),
					),
				),
			}, nil
		default:
			_, tmpGo, ok := listPrimitiveTypes(fd)
			if !ok {
				return nil, fmt.Errorf("list of %s is not yet supported", fd.Kind())
			}
			return []j.Code{
				j.If(guard).Block(
					j.Var().Id("tmp").Index().Add(tmpGo()),
					j.Id("diags").Dot("Append").Call(mDot.Clone().Dot("ElementsAs").Call(j.Id("ctx"), j.Op("&").Id("tmp"), j.False()).Op("...")),
					outDot.Clone().Op("=").Make(j.Index().Add(goTypeExpr(elem)), j.Len(j.Id("tmp"))),
					j.For(j.Id("i").Op(":=").Range().Id("tmp")).Block(
						outDot.Clone().Index(j.Id("i")).Op("=").Add(goTypeExpr(elem)).Parens(j.Id("tmp").Index(j.Id("i"))),
					),
				),
			}, nil
		}

	case isMessageKind(fd):
		if isWellKnownJSON(fd) {
			st := p.GoType.Elem() // structpb.Struct / Value / ListValue
			return []j.Code{
				j.If(j.Op("!").Add(mDot.Clone()).Dot("IsNull").Call().Op("&&").Op("!").Add(mDot.Clone()).Dot("IsUnknown").Call()).Block(
					j.Id("sv").Op(":=").Op("&").Qual(st.PkgPath(), st.Name()).Values(),
					j.If(
						j.Id("err").Op(":=").Qual(pkgProtoJSON, "Unmarshal").Call(j.Index().Byte().Parens(mDot.Clone().Dot("ValueString").Call()), j.Id("sv")),
						j.Id("err").Op("!=").Nil(),
					).Block(
						j.Id("diags").Dot("AddError").Call(j.Lit("Invalid JSON for "+p.Key), j.Id("err").Dot("Error").Call()),
					).Else().Block(
						outDot.Clone().Op("=").Id("sv"),
					),
				),
			}, nil
		}
		nested := p.GoType.Elem()
		if err := g.enqueue(nested); err != nil {
			return nil, err
		}
		nb := g.baseName(nested)
		return []j.Code{
			j.If(mDot.Clone().Op("!=").Nil()).Block(
				j.List(j.Id("v"), j.Id("d")).Op(":=").Id("expand"+nb).Call(j.Id("ctx"), mDot.Clone()),
				j.Id("diags").Dot("Append").Call(j.Id("d").Op("...")),
				outDot.Clone().Op("=").Id("v"),
			),
		}, nil

	default:
		return g.expandScalar(p, mDot, outDot)
	}
}

func (g *generator) expandScalar(p fieldPlan, mDot, outDot *j.Statement) ([]j.Code, error) {
	gt := p.GoType
	if p.OptionalScalar {
		gt = p.GoType.Elem() // *T -> T (for enum type naming)
	}
	expr, err := scalarToProto(p.FD, gt, mDot)
	if err != nil {
		return nil, err
	}
	if p.OptionalScalar {
		// proto3 optional: set a pointer only when the value is present.
		return []j.Code{
			j.If(j.Op("!").Add(mDot.Clone()).Dot("IsNull").Call().Op("&&").Op("!").Add(mDot.Clone()).Dot("IsUnknown").Call()).Block(
				j.Id("v").Op(":=").Add(expr),
				outDot.Clone().Op("=").Op("&").Id("v"),
			),
		}, nil
	}
	return []j.Code{outDot.Clone().Op("=").Add(expr)}, nil
}

// scalarToProto returns the Go expression converting a Terraform scalar accessor
// (src) into the proto field's native value.
func scalarToProto(fd protoreflect.FieldDescriptor, goType reflect.Type, src *j.Statement) (*j.Statement, error) {
	switch fd.Kind() {
	case protoreflect.StringKind:
		return src.Clone().Dot("ValueString").Call(), nil
	case protoreflect.BytesKind:
		return j.Index().Byte().Parens(src.Clone().Dot("ValueString").Call()), nil
	case protoreflect.BoolKind:
		return src.Clone().Dot("ValueBool").Call(), nil
	case protoreflect.EnumKind:
		pkg, name := goType.PkgPath(), goType.Name()
		return j.Qual(pkg, name).Call(j.Qual(pkg, name+"_value").Index(src.Clone().Dot("ValueString").Call())), nil
	case protoreflect.FloatKind:
		return j.Float32().Parens(src.Clone().Dot("ValueFloat64").Call()), nil
	case protoreflect.DoubleKind:
		return src.Clone().Dot("ValueFloat64").Call(), nil
	default:
		cast, ok := goIntCast(fd.Kind())
		if !ok {
			return nil, fmt.Errorf("unsupported scalar kind %s", fd.Kind())
		}
		return j.Id(cast).Parens(src.Clone().Dot("ValueInt64").Call()), nil
	}
}

func (g *generator) flattenField(p fieldPlan) ([]j.Code, error) {
	if p.Oneof != nil {
		return g.flattenOneof(p)
	}
	fd := p.FD
	mDot := j.Id("m").Dot(p.GoName)
	pDot := j.Id("p").Dot(p.GoName)

	// Create-only secrets are serialized to a JSON string (read/update preserve
	// the create-time value, set in the resource code, not here).
	if isCreateOnly(fd) {
		if fd.IsList() {
			return []j.Code{
				j.List(j.Id("v"), j.Id("d")).Op(":=").Id("protoMessagesToJSON").Call(pDot.Clone()),
				j.Id("diags").Dot("Append").Call(j.Id("d").Op("...")),
				mDot.Clone().Op("=").Id("v"),
			}, nil
		}
		return []j.Code{
			j.List(j.Id("v"), j.Id("d")).Op(":=").Id("protoMessageToJSON").Call(pDot.Clone()),
			j.Id("diags").Dot("Append").Call(j.Id("d").Op("...")),
			mDot.Clone().Op("=").Id("v"),
		}, nil
	}

	// INPUT_ONLY fields are never echoed back by the server; leave the model's
	// existing value (plan or prior state) untouched instead of overwriting it.
	if isInputOnly(fd) {
		return nil, nil
	}

	switch {
	case fd.IsMap():
		if isMessageKind(fd.MapValue()) {
			nested := p.GoType.Elem().Elem() // map[string]*Msg -> Msg
			if err := g.enqueue(nested); err != nil {
				return nil, err
			}
			nb := g.baseName(nested)
			return []j.Code{
				mDot.Clone().Op("=").Nil(),
				j.If(j.Len(pDot.Clone()).Op(">").Lit(0)).Block(
					mDot.Clone().Op("=").Make(j.Map(j.String()).Id(nb+"Model"), j.Len(pDot.Clone())),
					j.For(j.List(j.Id("k"), j.Id("pv")).Op(":=").Range().Add(pDot.Clone())).Block(
						j.Var().Id("em").Id(nb+"Model"),
						j.Id("diags").Dot("Append").Call(j.Id("flatten"+nb).Call(j.Id("ctx"), j.Id("pv"), j.Op("&").Id("em")).Op("...")),
						mDot.Clone().Index(j.Id("k")).Op("=").Id("em"),
					),
				),
			}, nil
		}
		if fd.MapValue().Kind() != protoreflect.StringKind {
			return nil, fmt.Errorf("map with %s value is not yet supported", fd.MapValue().Kind())
		}
		return []j.Code{j.Block(
			j.List(j.Id("mv"), j.Id("d")).Op(":=").Id("mapOrNull").Call(j.Id("ctx"), j.Qual(pkgTypes, "StringType"), pDot.Clone()),
			j.Id("diags").Dot("Append").Call(j.Id("d").Op("...")),
			mDot.Clone().Op("=").Id("mv"),
		)}, nil

	case fd.IsList():
		if isMessageKind(fd) {
			nested := p.GoType.Elem().Elem()
			if err := g.enqueue(nested); err != nil {
				return nil, err
			}
			nb := g.baseName(nested)
			return []j.Code{
				mDot.Clone().Op("=").Nil(),
				j.For(j.List(j.Id("_"), j.Id("e")).Op(":=").Range().Add(pDot.Clone())).Block(
					j.Var().Id("em").Id(nb+"Model"),
					j.Id("diags").Dot("Append").Call(j.Id("flatten"+nb).Call(j.Id("ctx"), j.Id("e"), j.Op("&").Id("em")).Op("...")),
					mDot.Clone().Op("=").Append(mDot.Clone(), j.Id("em")),
				),
			}, nil
		}
		switch {
		case fd.Kind() == protoreflect.StringKind:
			return []j.Code{j.Block(
				j.List(j.Id("lv"), j.Id("d")).Op(":=").Id("listOrNull").Call(j.Id("ctx"), j.Qual(pkgTypes, "StringType"), pDot.Clone()),
				j.Id("diags").Dot("Append").Call(j.Id("d").Op("...")),
				mDot.Clone().Op("=").Id("lv"),
			)}, nil
		case fd.Kind() == protoreflect.EnumKind:
			return []j.Code{j.Block(
				j.Id("tmp").Op(":=").Make(j.Index().String(), j.Len(pDot.Clone())),
				j.For(j.Id("i").Op(":=").Range().Add(pDot.Clone())).Block(
					j.Id("tmp").Index(j.Id("i")).Op("=").Add(pDot.Clone()).Index(j.Id("i")).Dot("String").Call(),
				),
				j.List(j.Id("lv"), j.Id("d")).Op(":=").Id("listOrNull").Call(j.Id("ctx"), j.Qual(pkgTypes, "StringType"), j.Id("tmp")),
				j.Id("diags").Dot("Append").Call(j.Id("d").Op("...")),
				mDot.Clone().Op("=").Id("lv"),
			)}, nil
		default:
			tfType, tmpGo, ok := listPrimitiveTypes(fd)
			if !ok {
				return nil, fmt.Errorf("list of %s is not yet supported", fd.Kind())
			}
			return []j.Code{j.Block(
				j.Id("tmp").Op(":=").Make(j.Index().Add(tmpGo()), j.Len(pDot.Clone())),
				j.For(j.Id("i").Op(":=").Range().Add(pDot.Clone())).Block(
					j.Id("tmp").Index(j.Id("i")).Op("=").Add(tmpGo()).Parens(j.Add(pDot.Clone()).Index(j.Id("i"))),
				),
				j.List(j.Id("lv"), j.Id("d")).Op(":=").Id("listOrNull").Call(j.Id("ctx"), tfType(), j.Id("tmp")),
				j.Id("diags").Dot("Append").Call(j.Id("d").Op("...")),
				mDot.Clone().Op("=").Id("lv"),
			)}, nil
		}

	case isMessageKind(fd):
		if isWellKnownJSON(fd) {
			return []j.Code{
				j.If(pDot.Clone().Op("!=").Nil()).Block(
					j.If(
						j.List(j.Id("b"), j.Id("err")).Op(":=").Qual(pkgProtoJSON, "Marshal").Call(pDot.Clone()),
						j.Id("err").Op("!=").Nil(),
					).Block(
						j.Id("diags").Dot("AddError").Call(j.Lit("Unable to encode "+p.Key), j.Id("err").Dot("Error").Call()),
					).Else().Block(
						mDot.Clone().Op("=").Qual(pkgJSONTypes, "NewNormalizedValue").Call(j.String().Parens(j.Id("b"))),
					),
				).Else().Block(
					mDot.Clone().Op("=").Qual(pkgJSONTypes, "NewNormalizedNull").Call(),
				),
			}, nil
		}
		nested := p.GoType.Elem()
		if err := g.enqueue(nested); err != nil {
			return nil, err
		}
		nb := g.baseName(nested)
		return []j.Code{
			j.If(pDot.Clone().Op("!=").Nil()).Block(
				j.Var().Id("nm").Id(nb+"Model"),
				j.Id("diags").Dot("Append").Call(j.Id("flatten"+nb).Call(j.Id("ctx"), pDot.Clone(), j.Op("&").Id("nm")).Op("...")),
				mDot.Clone().Op("=").Op("&").Id("nm"),
			).Else().Block(
				mDot.Clone().Op("=").Nil(),
			),
		}, nil

	default:
		return g.flattenScalar(p, mDot, pDot)
	}
}

func (g *generator) flattenScalar(p fieldPlan, mDot, pDot *j.Statement) ([]j.Code, error) {
	if p.OptionalScalar {
		// proto3 optional: deref the pointer when set, else null.
		val, err := scalarToModel(p.FD, j.Op("*").Add(pDot.Clone()))
		if err != nil {
			return nil, err
		}
		return []j.Code{
			j.If(pDot.Clone().Op("!=").Nil()).Block(
				mDot.Clone().Op("=").Add(val),
			).Else().Block(
				mDot.Clone().Op("=").Add(nullForKind(p.FD)),
			),
		}, nil
	}
	val, err := scalarToModel(p.FD, pDot)
	if err != nil {
		return nil, err
	}
	return []j.Code{mDot.Clone().Op("=").Add(val)}, nil
}

// scalarToModel returns the Go expression converting a proto scalar accessor
// (src) into the Terraform model value.
func scalarToModel(fd protoreflect.FieldDescriptor, src *j.Statement) (*j.Statement, error) {
	switch fd.Kind() {
	case protoreflect.StringKind:
		return j.Id("stringOrNull").Call(src.Clone()), nil
	case protoreflect.BytesKind:
		return j.Id("stringOrNull").Call(j.String().Parens(src.Clone())), nil
	case protoreflect.BoolKind:
		return j.Qual(pkgTypes, "BoolValue").Call(src.Clone()), nil
	case protoreflect.EnumKind:
		return j.Qual(pkgTypes, "StringValue").Call(src.Clone().Dot("String").Call()), nil
	case protoreflect.FloatKind:
		return j.Qual(pkgTypes, "Float64Value").Call(j.Float64().Parens(src.Clone())), nil
	case protoreflect.DoubleKind:
		return j.Qual(pkgTypes, "Float64Value").Call(src.Clone()), nil
	default:
		if _, ok := goIntCast(fd.Kind()); !ok {
			return nil, fmt.Errorf("unsupported scalar kind %s", fd.Kind())
		}
		return j.Qual(pkgTypes, "Int64Value").Call(j.Int64().Parens(src.Clone())), nil
	}
}

// listPrimitiveTypes returns, for a non-string, non-enum primitive list element,
// the Terraform element attr.Type and the Go type used for the intermediate slice
// that Terraform's ElementsAs/ListValueFrom understands. ok is false otherwise.
func listPrimitiveTypes(fd protoreflect.FieldDescriptor) (tfType func() *j.Statement, tmpGo func() *j.Statement, ok bool) {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		return func() *j.Statement { return j.Qual(pkgTypes, "BoolType") }, func() *j.Statement { return j.Bool() }, true
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return func() *j.Statement { return j.Qual(pkgTypes, "Float64Type") }, func() *j.Statement { return j.Float64() }, true
	}
	if _, isInt := goIntCast(fd.Kind()); isInt {
		return func() *j.Statement { return j.Qual(pkgTypes, "Int64Type") }, func() *j.Statement { return j.Int64() }, true
	}
	return nil, nil, false
}

// goTypeExpr renders a non-pointer Go type, handling builtins (no import path).
func goTypeExpr(t reflect.Type) *j.Statement {
	if t.PkgPath() == "" {
		return j.Id(t.String())
	}
	return j.Qual(t.PkgPath(), t.Name())
}

// nullForKind returns the typed Terraform null for a scalar field kind.
func nullForKind(fd protoreflect.FieldDescriptor) *j.Statement {
	switch fd.Kind() {
	case protoreflect.BoolKind:
		return j.Qual(pkgTypes, "BoolNull").Call()
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return j.Qual(pkgTypes, "Float64Null").Call()
	}
	if _, ok := goIntCast(fd.Kind()); ok {
		return j.Qual(pkgTypes, "Int64Null").Call()
	}
	return j.Qual(pkgTypes, "StringNull").Call()
}

// expandOneof emits the conversion for a oneof member: build the member value and
// assign the corresponding wrapper to the parent's interface field.
func (g *generator) expandOneof(p fieldPlan) ([]j.Code, error) {
	info := p.Oneof
	fd := p.FD
	mDot := j.Id("m").Dot(p.GoName)
	wrapper := j.Qual(info.ParentType.PkgPath(), info.WrapperType.Name())
	assign := func(val j.Code) j.Code {
		return j.Id("out").Dot(info.OneofGoField).Op("=").Op("&").Add(wrapper.Clone()).Values(j.Dict{j.Id(info.WrapperField): val})
	}

	if isWellKnownJSON(fd) {
		st := p.GoType.Elem() // structpb.Struct / Value / ListValue
		return []j.Code{
			j.If(j.Op("!").Add(mDot.Clone()).Dot("IsNull").Call().Op("&&").Op("!").Add(mDot.Clone()).Dot("IsUnknown").Call()).Block(
				j.Id("sv").Op(":=").Op("&").Qual(st.PkgPath(), st.Name()).Values(),
				j.If(
					j.Id("err").Op(":=").Qual(pkgProtoJSON, "Unmarshal").Call(j.Index().Byte().Parens(mDot.Clone().Dot("ValueString").Call()), j.Id("sv")),
					j.Id("err").Op("!=").Nil(),
				).Block(
					j.Id("diags").Dot("AddError").Call(j.Lit("Invalid JSON for "+p.Key), j.Id("err").Dot("Error").Call()),
				).Else().Block(
					assign(j.Id("sv")),
				),
			),
		}, nil
	}
	if isMessageKind(fd) {
		nested := p.GoType.Elem()
		if err := g.enqueue(nested); err != nil {
			return nil, err
		}
		nb := g.baseName(nested)
		return []j.Code{
			j.If(mDot.Clone().Op("!=").Nil()).Block(
				j.List(j.Id("v"), j.Id("d")).Op(":=").Id("expand"+nb).Call(j.Id("ctx"), mDot.Clone()),
				j.Id("diags").Dot("Append").Call(j.Id("d").Op("...")),
				j.If(j.Id("v").Op("!=").Nil()).Block(assign(j.Id("v"))),
			),
		}, nil
	}

	expr, err := scalarToProto(fd, p.GoType, mDot)
	if err != nil {
		return nil, err
	}
	return []j.Code{
		j.If(j.Op("!").Add(mDot.Clone()).Dot("IsNull").Call().Op("&&").Op("!").Add(mDot.Clone()).Dot("IsUnknown").Call()).Block(
			assign(expr),
		),
	}, nil
}

// flattenOneof emits the conversion for a oneof member: type-assert the parent's
// interface field to this member's wrapper, otherwise leave the attribute null.
func (g *generator) flattenOneof(p fieldPlan) ([]j.Code, error) {
	info := p.Oneof
	fd := p.FD
	mDot := j.Id("m").Dot(p.GoName)
	wrapper := j.Qual(info.ParentType.PkgPath(), info.WrapperType.Name())
	wAccess := j.Id("w").Dot(info.WrapperField)
	assertInit := j.List(j.Id("w"), j.Id("ok")).Op(":=").Id("p").Dot(info.OneofGoField).Assert(j.Op("*").Add(wrapper.Clone()))

	if isWellKnownJSON(fd) {
		return []j.Code{
			j.If(assertInit, j.Id("ok")).Block(
				j.If(
					j.List(j.Id("b"), j.Id("err")).Op(":=").Qual(pkgProtoJSON, "Marshal").Call(wAccess.Clone()),
					j.Id("err").Op("!=").Nil(),
				).Block(
					j.Id("diags").Dot("AddError").Call(j.Lit("Unable to encode "+p.Key), j.Id("err").Dot("Error").Call()),
				).Else().Block(
					mDot.Clone().Op("=").Qual(pkgJSONTypes, "NewNormalizedValue").Call(j.String().Parens(j.Id("b"))),
				),
			).Else().Block(
				mDot.Clone().Op("=").Qual(pkgJSONTypes, "NewNormalizedNull").Call(),
			),
		}, nil
	}
	if isMessageKind(fd) {
		nested := p.GoType.Elem()
		if err := g.enqueue(nested); err != nil {
			return nil, err
		}
		nb := g.baseName(nested)
		return []j.Code{
			j.If(assertInit, j.Id("ok").Op("&&").Add(wAccess.Clone()).Op("!=").Nil()).Block(
				j.Var().Id("nm").Id(nb+"Model"),
				j.Id("diags").Dot("Append").Call(j.Id("flatten"+nb).Call(j.Id("ctx"), wAccess.Clone(), j.Op("&").Id("nm")).Op("...")),
				mDot.Clone().Op("=").Op("&").Id("nm"),
			).Else().Block(
				mDot.Clone().Op("=").Nil(),
			),
		}, nil
	}

	val, err := scalarToModel(fd, wAccess)
	if err != nil {
		return nil, err
	}
	return []j.Code{
		j.If(assertInit, j.Id("ok")).Block(
			mDot.Clone().Op("=").Add(val),
		).Else().Block(
			mDot.Clone().Op("=").Add(nullForKind(fd)),
		),
	}, nil
}
