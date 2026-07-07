// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"encoding/json"
	"reflect"
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
