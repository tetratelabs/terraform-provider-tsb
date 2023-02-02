package helpers

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	tsbv2 "github.com/tetrateio/api/tsb/v2"
	"google.golang.org/grpc"
)

type Clients struct {
	Organization tsbv2.OrganizationsClient
	Tenant       tsbv2.TenantsClient
	Team         tsbv2.TeamsClient
}

func NewClients(cc *grpc.ClientConn) *Clients {
	return &Clients{
		Organization: tsbv2.NewOrganizationsClient(cc),
		Tenant:       tsbv2.NewTenantsClient(cc),
		Team:         tsbv2.NewTeamsClient(cc),
	}
}

// Can't use generics for this yet...
// See https://github.com/golang/go/issues/48522
func BuildClientsDatasource(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) *Clients {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		// resp.Diagnostics.AddError("Provider not configured", "Expected req.ProviderData to be populated, but was nil")
		return nil
	}

	clients, ok := req.ProviderData.(*Clients)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *clients, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return nil
	}
	return clients
}

// Can't use generics for this yet...
// See https://github.com/golang/go/issues/48522
func BuildClientsResource(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) *Clients {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		// resp.Diagnostics.AddError("Provider not configured", "Expectect req.ProviderData to be populated, but was nil")
		return nil
	}

	clients, ok := req.ProviderData.(*Clients)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *clients, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return nil
	}
	return clients
}
