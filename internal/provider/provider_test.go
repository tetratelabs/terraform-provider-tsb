// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// TestProviderSchema drives the whole provider (every registered resource) through
// the framework's GetProviderSchema, which structurally validates each schema.
// This catches schema-level problems the name-only parity test cannot.
func TestProviderSchema(t *testing.T) {
	ctx := context.Background()
	srv := providerserver.NewProtocol6(New()())()

	resp, err := srv.GetProviderSchema(ctx, &tfprotov6.GetProviderSchemaRequest{})
	if err != nil {
		t.Fatalf("GetProviderSchema: %v", err)
	}
	for _, d := range resp.Diagnostics {
		if d.Severity == tfprotov6.DiagnosticSeverityError {
			t.Errorf("schema diagnostic: %s: %s", d.Summary, d.Detail)
		}
	}
	if len(resp.ResourceSchemas) == 0 {
		t.Fatal("provider exposes no resources")
	}
	if _, ok := resp.ResourceSchemas["tsb_workspace"]; !ok {
		t.Errorf("tsb_workspace not registered; got %d resources", len(resp.ResourceSchemas))
	}
}
