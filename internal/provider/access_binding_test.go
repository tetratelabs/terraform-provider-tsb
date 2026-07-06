// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	rbac "github.com/tetrateio/tetrate/api/tsb/rbac/v2"
)

// TestAccessBindingRoundTrip verifies bindings/subjects (incl. the subject oneof)
// survive the model -> proto -> model conversion.
func TestAccessBindingRoundTrip(t *testing.T) {
	in := []accessBindingRuleModel{
		{
			Role: types.StringValue("organizations/o/roles/admin"),
			Subjects: []accessBindingSubjectModel{
				{User: types.StringValue("alice"), Team: types.StringNull(), ServiceAccount: types.StringNull()},
				{User: types.StringNull(), Team: types.StringValue("platform"), ServiceAccount: types.StringNull()},
				{User: types.StringNull(), Team: types.StringNull(), ServiceAccount: types.StringValue("ci")},
			},
		},
	}

	proto := expandBindings(in)
	if len(proto) != 1 || proto[0].Role != "organizations/o/roles/admin" || len(proto[0].Subjects) != 3 {
		t.Fatalf("expand mismatch: %+v", proto)
	}
	if proto[0].Subjects[0].GetUser() != "alice" ||
		proto[0].Subjects[1].GetTeam() != "platform" ||
		proto[0].Subjects[2].GetServiceAccount() != "ci" {
		t.Errorf("subject oneof not expanded correctly: %+v", proto[0].Subjects)
	}

	out := flattenBindings(proto)
	if !reflect.DeepEqual(in, out) {
		t.Errorf("round-trip mismatch:\n in:  %+v\n out: %+v", in, out)
	}
}

// TestAccessBindingEmpty verifies empty bindings flatten to nil (delete clears).
func TestAccessBindingEmpty(t *testing.T) {
	if got := expandBindings(nil); got != nil {
		t.Errorf("expand(nil) = %v, want nil", got)
	}
	if got := flattenBindings([]*rbac.Binding{}); got != nil {
		t.Errorf("flatten(empty) = %v, want nil", got)
	}
}
