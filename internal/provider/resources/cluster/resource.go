package cluster

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	v2 "github.com/tetrateio/api/tsb/v2"
	helpers "github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &ClusterResource{}
var _ resource.ResourceWithImportState = &ClusterResource{}

// Constructor
func NewResource() resource.Resource {
	return &ClusterResource{}
}

// Resource struct definition
type ClusterResource struct {
	client v2.ClustersClient
}

// Implement metadata function from Terraform resource interface
func (r *ClusterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

// Implement configure function from Terraform resource interface
func (r *ClusterResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	clients := helpers.BuildClientsResource(req, resp)
	if resp.Diagnostics.HasError() || clients == nil {
		return
	}
	r.client = clients.Cluster
}

// Implement schema function from Terraform resource interface
func (*ClusterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = v2.ClusterSchema()
}
