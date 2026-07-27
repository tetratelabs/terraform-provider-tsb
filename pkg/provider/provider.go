// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

// Package provider re-exports the TSB provider constructor from
// internal/provider so external modules (e.g. the pulumi-tsb bridge) can
// construct it without depending on an internal package.
package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/provider"

	internal "github.com/tetratelabs/terraform-provider-tsb/internal/provider"
)

// New returns a constructor for the TSB provider, suitable for
// providerserver.Serve or for bridging into another provider framework.
func New() func() provider.Provider {
	return internal.New()
}

// Groups maps each generated resource's Terraform type suffix to its TSB API
// group (see internal/provider/groups_gen.go for the exact contract and the
// hand-written-resource caveat).
func Groups() map[string]string {
	return internal.Groups()
}
