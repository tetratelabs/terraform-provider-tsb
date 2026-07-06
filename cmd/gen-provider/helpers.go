// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package main

import (
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// descriptorFor returns the proto message descriptor for a Go message struct type.
func descriptorFor(t reflect.Type) (protoreflect.MessageDescriptor, error) {
	v := reflect.New(t).Interface()
	pm, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("%s does not implement proto.Message", t.Name())
	}
	return pm.ProtoReflect().Descriptor(), nil
}

// protoStructFields maps a message's proto field names to their Go struct fields,
// by parsing the `protobuf:"...,name=<proto-name>,..."` tag.
func protoStructFields(t reflect.Type) map[string]reflect.StructField {
	out := map[string]reflect.StructField{}
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		tag := sf.Tag.Get("protobuf")
		if tag == "" {
			continue
		}
		if name := protoTagName(tag); name != "" {
			out[name] = sf
		}
	}
	return out
}

// protoTagName extracts the name=... value from a protobuf struct tag.
func protoTagName(tag string) string {
	for _, part := range strings.Split(tag, ",") {
		if rest, ok := strings.CutPrefix(part, "name="); ok {
			return rest
		}
	}
	return ""
}

// oneofMember resolves the Go representation of a proto oneof member field by
// reflection: it instantiates the parent message, activates the member through
// protoreflect, and reads back the concrete wrapper type assigned to the parent's
// interface field. This is robust to protoc-gen-go's trailing-underscore
// disambiguation of wrapper type names (e.g. Authentication_Rules_). It returns
// the resolved oneofInfo and the Go type of the member value.
func oneofMember(parent reflect.Type, fd protoreflect.FieldDescriptor) (oneofInfo, reflect.Type, error) {
	oneofName := string(fd.ContainingOneof().Name())

	var goField string
	for i := 0; i < parent.NumField(); i++ {
		if parent.Field(i).Tag.Get("protobuf_oneof") == oneofName {
			goField = parent.Field(i).Name
			break
		}
	}
	if goField == "" {
		return oneofInfo{}, nil, fmt.Errorf("no Go oneof field tagged %q on %s", oneofName, parent.Name())
	}

	msgPtr := reflect.New(parent)
	pm, ok := msgPtr.Interface().(proto.Message)
	if !ok {
		return oneofInfo{}, nil, fmt.Errorf("%s does not implement proto.Message", parent.Name())
	}
	refl := pm.ProtoReflect()

	// Activate this member so the Go interface field holds its wrapper.
	var v protoreflect.Value
	if fd.Message() != nil {
		v = refl.NewField(fd)
	} else {
		v = fd.Default()
	}
	refl.Set(fd, v)

	iface := msgPtr.Elem().FieldByName(goField)
	if iface.IsNil() {
		return oneofInfo{}, nil, fmt.Errorf("oneof member %q did not populate %s.%s", fd.Name(), parent.Name(), goField)
	}
	wrapperType := iface.Elem().Type().Elem() // *Wrapper -> Wrapper
	if wrapperType.NumField() == 0 {
		return oneofInfo{}, nil, fmt.Errorf("oneof wrapper %s has no fields", wrapperType.Name())
	}
	wf := wrapperType.Field(0)

	return oneofInfo{
		OneofGoField: goField,
		WrapperType:  wrapperType,
		WrapperField: wf.Name,
		ParentType:   parent,
	}, wf.Type, nil
}

// goIntCast returns the Go integer type name for an integer-kinded field, used to
// cast a types.Int64 value back to the proto field's concrete type.
func goIntCast(kind protoreflect.Kind) (string, bool) {
	switch kind {
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return "int32", true
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return "int64", true
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		return "uint32", true
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		return "uint64", true
	default:
		return "", false
	}
}
