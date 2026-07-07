// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"strings"
	"testing"

	v2 "github.com/tetrateio/tetrate/api/tsb/v2"
)

// TestServiceAccountKeysCreateOnly verifies that the create-time service account
// keys (including the private key) are flattened into the JSON-string attribute,
// and that expand never sends them back (they are output-only).
func TestServiceAccountKeysCreateOnly(t *testing.T) {
	ctx := context.Background()

	in := &v2.ServiceAccount{
		Fqn: "organizations/o/serviceaccounts/sa",
		Keys: []*v2.ServiceAccount_KeyPair{
			{Id: "k1", PublicKey: "PUB", PrivateKey: "PRIVATE-SECRET", DefaultToken: "TOKEN"},
		},
	}

	var m ServiceAccountModel
	if d := flattenServiceAccount(ctx, in, &m); d.HasError() {
		t.Fatalf("flatten: %v", d)
	}
	if m.Keys.IsNull() {
		t.Fatal("keys should be populated after create flatten")
	}
	js := m.Keys.ValueString()
	for _, want := range []string{"PRIVATE-SECRET", "TOKEN", "PUB"} {
		if !strings.Contains(js, want) {
			t.Errorf("keys JSON missing %q: %s", want, js)
		}
	}

	// expand must not populate the output-only keys.
	out, d := expandServiceAccount(ctx, &m)
	if d.HasError() {
		t.Fatalf("expand: %v", d)
	}
	if len(out.Keys) != 0 {
		t.Errorf("expand should not send output-only keys, got %d", len(out.Keys))
	}
}

// TestServiceAccountNoKeys verifies an empty key set flattens to null.
func TestServiceAccountNoKeys(t *testing.T) {
	ctx := context.Background()
	var m ServiceAccountModel
	if d := flattenServiceAccount(ctx, &v2.ServiceAccount{}, &m); d.HasError() {
		t.Fatalf("flatten: %v", d)
	}
	if !m.Keys.IsNull() {
		t.Errorf("keys should be null when none returned, got %q", m.Keys.ValueString())
	}
}
