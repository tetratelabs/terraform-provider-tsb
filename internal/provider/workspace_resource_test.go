// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"reflect"
	"sort"
	"testing"

	rschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"google.golang.org/protobuf/types/known/wrapperspb"

	typesv2 "github.com/tetrateio/tetrate/api/tsb/types/v2"
	v2 "github.com/tetrateio/tetrate/api/tsb/v2"
)

// tfsdkTags returns the set of tfsdk attribute names declared on a model struct.
func tfsdkTags(t reflect.Type) []string {
	var out []string
	for i := 0; i < t.NumField(); i++ {
		if tag, ok := t.Field(i).Tag.Lookup("tfsdk"); ok {
			out = append(out, tag)
		}
	}
	sort.Strings(out)
	return out
}

// TestWorkspaceModelMatchesSchema asserts the top-level WorkspaceModel tfsdk
// tags line up exactly with v2.WorkspaceSchema(); drift would panic at runtime.
func TestWorkspaceModelMatchesSchema(t *testing.T) {
	schema := v2.WorkspaceSchema()

	want := make([]string, 0, len(schema.Attributes))
	for name := range schema.Attributes {
		want = append(want, name)
	}
	sort.Strings(want)

	got := tfsdkTags(reflect.TypeOf(WorkspaceModel{}))

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("WorkspaceModel tfsdk tags do not match schema attributes\n schema: %v\n model:  %v", want, got)
	}
}

// TestNamespaceSelectorMatchesSchema checks the nested attribute parity.
func TestNamespaceSelectorMatchesSchema(t *testing.T) {
	schema := v2.WorkspaceSchema()
	nested, ok := schema.Attributes["namespace_selector"].(rschema.SingleNestedAttribute)
	if !ok {
		t.Fatalf("namespace_selector is not a SingleNestedAttribute: %T", schema.Attributes["namespace_selector"])
	}
	want := make([]string, 0, len(nested.Attributes))
	for name := range nested.Attributes {
		want = append(want, name)
	}
	sort.Strings(want)

	got := tfsdkTags(reflect.TypeOf(NamespaceSelectorModel{}))
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("NamespaceSelectorModel tfsdk tags do not match schema\n schema: %v\n model:  %v", want, got)
	}
}

// TestWorkspaceRoundTrip flattens a proto into the model and expands it back,
// asserting that the user-managed fields survive the round trip.
func TestWorkspaceRoundTrip(t *testing.T) {
	ctx := context.Background()
	in := &v2.Workspace{
		Fqn:                       "organizations/o/tenants/t/workspaces/w",
		DisplayName:               "My Workspace",
		Description:               "desc",
		IsolationBoundary:         "global",
		SecurityDomain:            "sd",
		DeletionProtectionEnabled: true,
		Profiles:                  []string{"p1", "p2"},
		NamespaceSelector:         &typesv2.NamespaceSelector{Names: []string{"*/ns1"}},
		Privileged:                &wrapperspb.BoolValue{Value: true},
		ConfigGenerationMetadata: &typesv2.ConfigGenerationMetadata{
			Labels:      map[string]string{"k": "v"},
			Annotations: map[string]string{"a": "b"},
		},
	}

	var m WorkspaceModel
	if d := flattenWorkspace(ctx, in, &m); d.HasError() {
		t.Fatalf("flatten failed: %v", d)
	}
	if got := m.ID.ValueString(); got != in.Fqn {
		t.Errorf("id = %q, want %q", got, in.Fqn)
	}

	out, d := expandWorkspace(ctx, &m)
	if d.HasError() {
		t.Fatalf("expand failed: %v", d)
	}

	if out.DisplayName != in.DisplayName ||
		out.Description != in.Description ||
		out.IsolationBoundary != in.IsolationBoundary ||
		out.SecurityDomain != in.SecurityDomain ||
		out.DeletionProtectionEnabled != in.DeletionProtectionEnabled {
		t.Errorf("scalar round-trip mismatch:\n in:  %+v\n out: %+v", in, out)
	}
	if !reflect.DeepEqual(out.Profiles, in.Profiles) {
		t.Errorf("profiles round-trip mismatch: got %v want %v", out.Profiles, in.Profiles)
	}
	if out.NamespaceSelector == nil || !reflect.DeepEqual(out.NamespaceSelector.Names, in.NamespaceSelector.Names) {
		t.Errorf("namespace_selector round-trip mismatch: got %v", out.NamespaceSelector)
	}
	if out.Privileged == nil || out.Privileged.GetValue() != true {
		t.Errorf("privileged round-trip mismatch: got %v", out.Privileged)
	}
	if out.ConfigGenerationMetadata == nil ||
		!reflect.DeepEqual(out.ConfigGenerationMetadata.Labels, in.ConfigGenerationMetadata.Labels) ||
		!reflect.DeepEqual(out.ConfigGenerationMetadata.Annotations, in.ConfigGenerationMetadata.Annotations) {
		t.Errorf("config_generation_metadata round-trip mismatch: got %v", out.ConfigGenerationMetadata)
	}
}
