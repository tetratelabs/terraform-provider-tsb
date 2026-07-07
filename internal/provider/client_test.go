// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"strings"
	"testing"
)

// TestPerRPCCredentials verifies that a token takes precedence and that
// username/password falls back to HTTP Basic auth.
func TestPerRPCCredentials(t *testing.T) {
	ctx := context.Background()

	tokenMD, err := Config{Token: "abc", Username: "u", Password: "p"}.
		perRPCCredentials().GetRequestMetadata(ctx)
	if err != nil {
		t.Fatalf("token metadata: %v", err)
	}
	if got := tokenMD["x-tetrate-token"]; got != "abc" {
		t.Errorf("expected x-tetrate-token=abc, got %q (md=%v)", got, tokenMD)
	}
	if _, ok := tokenMD["authorization"]; ok {
		t.Error("token auth should not set an authorization header")
	}

	basicMD, err := Config{Username: "alice", Password: "secret"}.
		perRPCCredentials().GetRequestMetadata(ctx)
	if err != nil {
		t.Fatalf("basic metadata: %v", err)
	}
	if got := basicMD["authorization"]; !strings.HasPrefix(got, "Basic ") {
		t.Errorf("expected Basic authorization header, got %q (md=%v)", got, basicMD)
	}
}
