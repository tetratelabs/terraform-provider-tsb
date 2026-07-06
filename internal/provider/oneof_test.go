// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"testing"

	auth "github.com/tetrateio/tetrate/api/tsb/auth/v2"
)

// TestOneofRoundTrip verifies that the selected member of a proto oneof survives
// the proto -> model -> proto round trip, including the wrapper whose Go type
// carries a trailing underscore (Authentication_Rules_).
func TestOneofRoundTrip(t *testing.T) {
	ctx := context.Background()

	// Select the "rules" member (its wrapper type is Authentication_Rules_).
	in := &auth.Authentication{
		Authn: &auth.Authentication_Rules_{Rules: &auth.Authentication_Rules{}},
	}

	var m AuthAuthenticationModel
	if d := flattenAuthAuthentication(ctx, in, &m); d.HasError() {
		t.Fatalf("flatten: %v", d)
	}
	if m.Rules == nil {
		t.Fatal("rules member should be set after flatten")
	}
	if m.Jwt != nil {
		t.Error("jwt member should be nil when rules is selected")
	}

	out, d := expandAuthAuthentication(ctx, &m)
	if d.HasError() {
		t.Fatalf("expand: %v", d)
	}
	if _, ok := out.Authn.(*auth.Authentication_Rules_); !ok {
		t.Errorf("expected Authn to be *Authentication_Rules_, got %T", out.Authn)
	}
	if out.GetJwt() != nil {
		t.Error("jwt should not be set after round trip")
	}
}

// TestOneofUnset verifies that an unset oneof flattens to all-null members.
func TestOneofUnset(t *testing.T) {
	ctx := context.Background()
	var m AuthAuthenticationModel
	if d := flattenAuthAuthentication(ctx, &auth.Authentication{}, &m); d.HasError() {
		t.Fatalf("flatten: %v", d)
	}
	if m.Jwt != nil || m.Rules != nil {
		t.Errorf("unset oneof should leave members nil, got jwt=%v rules=%v", m.Jwt, m.Rules)
	}
}
