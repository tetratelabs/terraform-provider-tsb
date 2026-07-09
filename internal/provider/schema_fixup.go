// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/defaults"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
)

// enumZeroValueDefault is the sentinel a proto enum's zero value renders as
// (e.g. SourceType_INVALID = 0 -> "INVALID"). protoc-gen-terraform falls back
// to this name as a Computed attribute's Default whenever the proto field is
// missing a `+protoc-gen-terraform:enumdefault:N` override tag — it is never a
// deliberately chosen default. See defaultIsUntrustworthy.
const enumZeroValueDefault = "INVALID"

// fixupSchema adapts a reused TSB schema (produced by protoc-gen-terraform) so it
// is valid and well-behaved for the Terraform Plugin Framework at serve time.
//
// Adjustments are made, recursively, to every attribute:
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
//
//  3. Computed-only (not Optional/Required) string attributes whose Default
//     resolves to the proto enum zero-value sentinel have the Default dropped
//     in favor of UseStateForUnknown. Such a Default is protoc-gen-terraform's
//     "no real default configured" fallback, not a value TSB actually persists
//     (e.g. Team.source_type defaults to "INVALID", but CreateTeam/UpdateTeam
//     always coerce an unset source type to LOCAL server-side) — keeping it
//     forces a known planned value that the server's real response then
//     contradicts, producing "Provider produced inconsistent result after
//     apply" on every create. Leaving the planned value unknown lets it be
//     filled in from whatever the server actually returns.
func fixupSchema(s schema.Schema) schema.Schema {
	s.Attributes = fixupAttributes(s.Attributes)
	return s
}

// defaultIsUntrustworthy reports whether d resolves to the proto enum
// zero-value sentinel (see enumZeroValueDefault).
func defaultIsUntrustworthy(d defaults.String) bool {
	var resp defaults.StringResponse
	d.DefaultString(context.Background(), defaults.StringRequest{}, &resp)
	return resp.PlanValue.ValueString() == enumZeroValueDefault
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
		switch {
		case a.Required:
			a.Default = nil
		case a.Optional:
			a.Computed = true
		case a.Computed && a.Default != nil && defaultIsUntrustworthy(a.Default):
			a.Default = nil
			a.PlanModifiers = append(a.PlanModifiers, stringplanmodifier.UseStateForUnknown())
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
