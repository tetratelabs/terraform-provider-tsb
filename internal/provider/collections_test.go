// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"testing"

	k8s "github.com/tetrateio/tetrate/api/install/kubernetes"
	gateway "github.com/tetrateio/tetrate/api/tsb/gateway/v2"
	rbac "github.com/tetrateio/tetrate/api/tsb/rbac/v2"
)

// TestMapMessageRoundTrip exercises a map<string, message> whose value message
// itself contains a oneof (StringMatch.match_type).
func TestMapMessageRoundTrip(t *testing.T) {
	ctx := context.Background()
	in := &gateway.HttpMatchCondition{
		Headers: map[string]*gateway.StringMatch{
			"x-request-id": {MatchType: &gateway.StringMatch_Exact{Exact: "abc"}},
			"x-env":        {MatchType: &gateway.StringMatch_Prefix{Prefix: "prod-"}},
		},
	}

	var m HttpMatchConditionModel
	if d := flattenHttpMatchCondition(ctx, in, &m); d.HasError() {
		t.Fatalf("flatten: %v", d)
	}
	if len(m.Headers) != 2 {
		t.Fatalf("expected 2 headers, got %d", len(m.Headers))
	}
	if got := m.Headers["x-request-id"].Exact.ValueString(); got != "abc" {
		t.Errorf("x-request-id exact = %q, want abc", got)
	}

	out, d := expandHttpMatchCondition(ctx, &m)
	if d.HasError() {
		t.Fatalf("expand: %v", d)
	}
	if got := out.Headers["x-request-id"].GetExact(); got != "abc" {
		t.Errorf("round-trip x-request-id exact = %q, want abc", got)
	}
	if got := out.Headers["x-env"].GetPrefix(); got != "prod-" {
		t.Errorf("round-trip x-env prefix = %q, want prod-", got)
	}
}

// TestNumericListRoundTrip exercises a repeated uint32 (PodSecurityContext.supplemental_groups).
func TestNumericListRoundTrip(t *testing.T) {
	ctx := context.Background()
	in := &k8s.PodSecurityContext{SupplementalGroups: []uint32{1, 2, 65534}}

	var m PodSecurityContextModel
	if d := flattenPodSecurityContext(ctx, in, &m); d.HasError() {
		t.Fatalf("flatten: %v", d)
	}
	out, d := expandPodSecurityContext(ctx, &m)
	if d.HasError() {
		t.Fatalf("expand: %v", d)
	}
	if len(out.SupplementalGroups) != 3 || out.SupplementalGroups[2] != 65534 {
		t.Errorf("supplemental_groups round-trip mismatch: %v", out.SupplementalGroups)
	}
}

// TestEnumListRoundTrip exercises a repeated enum (Role_Rule.permissions).
func TestEnumListRoundTrip(t *testing.T) {
	ctx := context.Background()
	in := &rbac.Role_Rule{Permissions: []rbac.Permission{rbac.Permission_READ, rbac.Permission_WRITE}}

	var m Role_RuleModel
	if d := flattenRole_Rule(ctx, in, &m); d.HasError() {
		t.Fatalf("flatten: %v", d)
	}
	out, d := expandRole_Rule(ctx, &m)
	if d.HasError() {
		t.Fatalf("expand: %v", d)
	}
	if len(out.Permissions) != 2 ||
		out.Permissions[0] != rbac.Permission_READ ||
		out.Permissions[1] != rbac.Permission_WRITE {
		t.Errorf("permissions round-trip mismatch: %v", out.Permissions)
	}
}
