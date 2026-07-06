// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"encoding/json"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"

	extension "github.com/tetrateio/tetrate/api/tsb/extension/v2"
)

// TestStructpbRoundTrip verifies that a google.protobuf.Struct field survives the
// proto -> JSON-string model -> proto round trip (WasmExtension.config).
func TestStructpbRoundTrip(t *testing.T) {
	ctx := context.Background()

	cfg, err := structpb.NewStruct(map[string]any{
		"headersToAdd": []any{map[string]any{"key": "x", "value": "y"}},
		"enabled":      true,
		"retries":      float64(3),
	})
	if err != nil {
		t.Fatalf("NewStruct: %v", err)
	}
	in := &extension.WasmExtension{Fqn: "organizations/o/extensions/e", Config: cfg}

	var m ExtensionWasmExtensionModel
	if d := flattenExtensionWasmExtension(ctx, in, &m); d.HasError() {
		t.Fatalf("flatten: %v", d)
	}

	if m.Config.IsNull() {
		t.Fatal("config should not be null after flatten")
	}
	// The flattened value must be valid JSON.
	var decoded map[string]any
	if err := json.Unmarshal([]byte(m.Config.ValueString()), &decoded); err != nil {
		t.Fatalf("flattened config is not valid JSON: %v (%q)", err, m.Config.ValueString())
	}

	out, d := expandExtensionWasmExtension(ctx, &m)
	if d.HasError() {
		t.Fatalf("expand: %v", d)
	}
	if !proto.Equal(in.Config, out.Config) {
		t.Errorf("config round-trip mismatch:\n in:  %v\n out: %v", in.Config, out.Config)
	}
}

// TestStructpbNull verifies nil structpb flattens to a null JSON string.
func TestStructpbNull(t *testing.T) {
	ctx := context.Background()
	var m ExtensionWasmExtensionModel
	if d := flattenExtensionWasmExtension(ctx, &extension.WasmExtension{}, &m); d.HasError() {
		t.Fatalf("flatten: %v", d)
	}
	if !m.Config.IsNull() {
		t.Errorf("config should be null for a nil Struct, got %q", m.Config.ValueString())
	}
}
