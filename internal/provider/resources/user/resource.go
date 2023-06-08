package user

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	v2 "github.com/tetrateio/api/tsb/v2"
	helpers "github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &UserResource{}
var _ resource.ResourceWithImportState = &UserResource{}

// Constructor
func NewResource() resource.Resource {
	return &UserResource{}
}

// Resource struct definition
type UserResource struct {
	client v2.TeamsClient
}

// Implement metadata function from Terraform resource interface
func (r *UserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

// Implement configure function from Terraform resource interface
func (r *UserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	clients := helpers.BuildClientsResource(req, resp)
	if resp.Diagnostics.HasError() || clients == nil {
		return
	}
	r.client = clients.User
}

// Implement schema function from Terraform resource interface
func (*UserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = v2.UserSchema()
}
