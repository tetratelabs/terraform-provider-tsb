// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// stringOrNull maps an empty proto string to a null Terraform value so optional
// fields the user never set stay null in state, and non-empty values round-trip.
func stringOrNull(s string) types.String {
	if s == "" {
		return types.StringNull()
	}
	return types.StringValue(s)
}

// expandEnum converts a Terraform string into the proto enum number registered
// under that name in values (an <Enum>_value map), appending an error diagnostic
// when the name is not declared.
//
// A null or unknown value maps to the enum's zero value with no diagnostic.
// fixupSchema makes optional leaf attributes Optional+Computed, so an attribute
// the user never set arrives null (or unknown, at create) — that is how "unset"
// legitimately reaches the wire as the proto default.
func expandEnum(v types.String, values map[string]int32, attrName string, diags *diag.Diagnostics) int32 {
	if v.IsNull() || v.IsUnknown() {
		return 0
	}
	return expandEnumString(v.ValueString(), values, attrName, diags)
}

// expandEnumString is expandEnum for a value already unwrapped to a string, as in
// the repeated-enum path where elements arrive through ElementsAs.
//
// An unrecognised name is an error rather than a silent fallback to the zero
// value. Besides catching typos, this covers version skew: protobuf-go renders an
// enum number the compiled descriptor does not know as its decimal, so a newer
// server's value round-trips through state as e.g. "7". Sending the zero value
// instead would silently overwrite the server's real configuration.
func expandEnumString(s string, values map[string]int32, attrName string, diags *diag.Diagnostics) int32 {
	if s == "" {
		return 0
	}
	if n, ok := values[s]; ok {
		return n
	}
	diags.AddError(
		"Invalid value for "+attrName,
		fmt.Sprintf("%q is not a valid %s value. Valid values are: %s.",
			s, attrName, strings.Join(enumNames(values), ", ")),
	)
	return 0
}

// enumNames returns the declared names of an <Enum>_value map ordered by enum
// number, so diagnostics list them in proto declaration order rather than in Go's
// randomized map order.
func enumNames(values map[string]int32) []string {
	names := make([]string, 0, len(values))
	for n := range values {
		names = append(names, n)
	}
	sort.Slice(names, func(i, j int) bool { return values[names[i]] < values[names[j]] })
	return names
}

// listOrNull builds a Terraform list from a Go slice, returning a null list for
// nil/empty input.
func listOrNull[T any](ctx context.Context, elemType attr.Type, in []T) (types.List, diag.Diagnostics) {
	if len(in) == 0 {
		return types.ListNull(elemType), nil
	}
	return types.ListValueFrom(ctx, elemType, in)
}

// mapOrNull builds a Terraform map from a Go map, returning a null map for
// nil/empty input.
func mapOrNull[V any](ctx context.Context, elemType attr.Type, in map[string]V) (types.Map, diag.Diagnostics) {
	if len(in) == 0 {
		return types.MapNull(elemType), nil
	}
	return types.MapValueFrom(ctx, elemType, in)
}

// protoMessageToJSON serializes a single proto message to a normalized JSON
// string (null when the message is nil). Used for create-only secret fields.
func protoMessageToJSON(m proto.Message) (jsontypes.Normalized, diag.Diagnostics) {
	var diags diag.Diagnostics
	if m == nil || reflect.ValueOf(m).IsNil() {
		return jsontypes.NewNormalizedNull(), diags
	}
	b, err := protojson.Marshal(m)
	if err != nil {
		diags.AddError("Unable to encode value", err.Error())
		return jsontypes.NewNormalizedNull(), diags
	}
	return jsontypes.NewNormalizedValue(string(b)), diags
}

// protoMessagesToJSON serializes a slice of proto messages to a normalized JSON
// array string (null when the slice is empty). Used for repeated create-only
// secret fields such as ServiceAccount.keys.
func protoMessagesToJSON[T proto.Message](msgs []T) (jsontypes.Normalized, diag.Diagnostics) {
	var diags diag.Diagnostics
	if len(msgs) == 0 {
		return jsontypes.NewNormalizedNull(), diags
	}
	parts := make([]json.RawMessage, len(msgs))
	for i, m := range msgs {
		b, err := protojson.Marshal(m)
		if err != nil {
			diags.AddError("Unable to encode value", err.Error())
			return jsontypes.NewNormalizedNull(), diags
		}
		parts[i] = b
	}
	b, err := json.Marshal(parts)
	if err != nil {
		diags.AddError("Unable to encode value", err.Error())
		return jsontypes.NewNormalizedNull(), diags
	}
	return jsontypes.NewNormalizedValue(string(b)), diags
}

// fqnParentName splits a fully-qualified name into its parent prefix and short
// name given the resource's collection segment (e.g. "workspaces" in
// organizations/o/tenants/t/workspaces/w -> parent organizations/o/tenants/t,
// name w). Used on import, where only the FQN is known.
func fqnParentName(fqn, collection string) (parent, name string) {
	marker := "/" + collection + "/"
	idx := strings.LastIndex(fqn, marker)
	if idx < 0 {
		// Fall back to the last path segment as the name.
		if i := strings.LastIndex(fqn, "/"); i >= 0 {
			return fqn[:i], fqn[i+1:]
		}
		return "", fqn
	}
	return fqn[:idx], fqn[idx+len(marker):]
}
