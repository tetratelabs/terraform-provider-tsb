// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// fixupSchema adapts a reused TSB schema (produced by protoc-gen-terraform) so it
// is valid and well-behaved for the Terraform Plugin Framework at serve time.
//
// Two adjustments are made, recursively, to every attribute:
//
//  1. Required attributes carrying a Default (the enum StaticString default the
//     schema generator emits) have the Default dropped — Required+Default is
//     contradictory and the framework rejects it.
//
//  2. Optional leaf attributes (scalars and primitive collections) are made
//     Optional+Computed. TSB is the authoritative store and returns a fully
//     materialized object, so a field the user leaves unset comes back populated
//     with the server's value (e.g. a bool defaulting to false, or an enum
//     default). Without Computed, Terraform fails with "Provider produced
//     inconsistent result after apply: was null, but now ..." for those
//     server-defaulted fields. Optional+Computed lets the provider return the
//     server value when unset, while still tracking drift on configured values.
//
//     Nested-object attributes (single/list/map/set nested) are deliberately NOT
//     made Computed: the generated models back them with plain Go structs, which
//     cannot represent the "unknown" plan value that Computed introduces. Their
//     leaf attributes are still fixed up via recursion.
func fixupSchema(s schema.Schema) schema.Schema {
	s.Attributes = fixupAttributes(s.Attributes)
	return s
}

func fixupAttributes(in map[string]schema.Attribute) map[string]schema.Attribute {
	out := make(map[string]schema.Attribute, len(in))
	for name, attr := range in {
		out[name] = fixupAttribute(attr)
	}
	return out
}

func fixupAttribute(attr schema.Attribute) schema.Attribute {
	switch a := attr.(type) {
	case schema.StringAttribute:
		if a.Required {
			a.Default = nil
		} else if a.Optional {
			a.Computed = true
		}
		return a
	case schema.BoolAttribute:
		if a.Optional {
			a.Computed = true
		}
		return a
	case schema.Int64Attribute:
		if a.Optional {
			a.Computed = true
		}
		return a
	case schema.Float64Attribute:
		if a.Optional {
			a.Computed = true
		}
		return a
	case schema.NumberAttribute:
		if a.Optional {
			a.Computed = true
		}
		return a
	case schema.ListAttribute:
		if a.Optional {
			a.Computed = true
		}
		return a
	case schema.MapAttribute:
		if a.Optional {
			a.Computed = true
		}
		return a
	case schema.SetAttribute:
		if a.Optional {
			a.Computed = true
		}
		return a
	case schema.ObjectAttribute:
		if a.Optional {
			a.Computed = true
		}
		return a
	case schema.SingleNestedAttribute:
		// Not made Computed: the model field is a *Model struct (can't be unknown).
		a.Attributes = fixupAttributes(a.Attributes)
		return a
	case schema.ListNestedAttribute:
		a.NestedObject.Attributes = fixupAttributes(a.NestedObject.Attributes)
		return a
	case schema.MapNestedAttribute:
		a.NestedObject.Attributes = fixupAttributes(a.NestedObject.Attributes)
		return a
	case schema.SetNestedAttribute:
		a.NestedObject.Attributes = fixupAttributes(a.NestedObject.Attributes)
		return a
	default:
		return attr
	}
}
