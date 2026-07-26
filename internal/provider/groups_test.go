// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import "testing"

// TestGroupsMatchResources asserts Groups() and Resources() are built from the
// same generator input; a mismatch means emitGroups and emitProvider silently
// diverged (e.g. one filtering differently) in cmd/gen-provider.
func TestGroupsMatchResources(t *testing.T) {
	groups := Groups()
	resources := Resources()
	if len(groups) != len(resources) {
		t.Fatalf("len(Groups()) = %d, len(Resources()) = %d, want equal", len(groups), len(resources))
	}
}
