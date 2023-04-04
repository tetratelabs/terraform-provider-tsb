package tenant

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	v2 "github.com/tetrateio/api/tsb/v2"
	helpers "github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &TenantResource{}
var _ resource.ResourceWithImportState = &TenantResource{}

// Constructor
func NewResource() resource.Resource {
	return &TenantResource{}
}

// Resource struct definition
type TenantResource struct {
	client v2.TenantsClient
}

// Implement metadata function from Terraform resource interface
func (r *TenantResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant"
}

// Implement configure function from Terraform resource interface
func (r *TenantResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	clients := helpers.BuildClientsResource(req, resp)
	if resp.Diagnostics.HasError() || clients == nil {
		return
	}
	r.client = clients.Tenant
}

// Implement schema function from Terraform resource interface
func (*TenantResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = v2.TenantSchema()
}
