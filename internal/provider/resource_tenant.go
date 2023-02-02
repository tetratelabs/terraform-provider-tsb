// Copyright 2023 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tsbv2 "github.com/tetrateio/api/tsb/v2"

	"github.com/tetrateio/tetrate/pkg/api"
	"github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &TenantResource{}
var _ resource.ResourceWithImportState = &TenantResource{}

func NewTenantResource() resource.Resource {
	return &TenantResource{}
}

// TenantResource defines the resource implementation.
type TenantResource struct {
	client tsbv2.TenantsClient
}

// Schema implements resource.Resource
func (*TenantResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	s, d := tsbv2.GenSchemaCreateTenantRequest(ctx)
	resp.Diagnostics.Append(d...)
	resp.Schema = s
}

func (r *TenantResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant"
}

func (r *TenantResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	clients := helpers.BuildClientsResource(ctx, req, resp)
	if resp.Diagnostics.HasError() || clients == nil {
		return
	}
	r.client = clients.Tenant
}

func (r *TenantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	data := tsbv2.NewCreateTenantRequestModelFromPlan(ctx, req.Plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	tenant, err := r.client.CreateTenant(ctx, data.ToGo(ctx))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Tenant",
			"Could not create tenant: "+err.Error(),
		)
		return
	}

	data.LoadFromResult(ctx, tenant.Fqn, &tsbv2.CreateTenantRequest{
		Parent: data.Parent.ValueString(),
		Name:   data.Name.ValueString(),
		Tenant: tenant,
	})

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TenantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	data := tsbv2.NewCreateTenantRequestModelFromState(ctx, req.State, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	fqn := data.Id.ValueString()
	meta, err := helpers.FromFQN(api.TenantKind, fqn)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Tenant", "Could not parse fqn: "+err.Error())
		return
	}

	tenant, err := r.client.GetTenant(ctx, &tsbv2.GetTenantRequest{Fqn: fqn})
	if err != nil {
		resp.Diagnostics.AddError("Error reading Tenant", "GetTenant request failed: "+err.Error())
		return
	}

	data.LoadFromResult(ctx, fqn, &tsbv2.CreateTenantRequest{
		Parent: helpers.ParentFQN(api.TenantKind, meta),
		Name:   meta.Name,
		Tenant: tenant,
	})

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TenantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	data := tsbv2.NewCreateTenantRequestModelFromPlan(ctx, req.Plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// We need to run a get to retreive the latest etag and set fqn as its not going to be in the plan (only state)
	fqn := data.Id.ValueString()
	tenant, err := r.client.GetTenant(ctx, &tsbv2.GetTenantRequest{Fqn: fqn})
	if err != nil {
		resp.Diagnostics.AddError("Error updating Tenant", "GetTenant request failed: "+err.Error())
		return
	}
	data.Tenant.Etag = types.StringValue(tenant.Etag)
	data.Tenant.Fqn = types.StringValue(tenant.Fqn)

	// Do the actual update
	_, err = r.client.UpdateTenant(ctx, data.Tenant.ToGo(ctx))
	if err != nil {
		resp.Diagnostics.AddError("Error updating Tenant", "UpdateTenant request failed: "+err.Error())
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TenantResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	data := tsbv2.NewCreateTenantRequestModelFromState(ctx, req.State, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.client.DeleteTenant(ctx, &tsbv2.DeleteTenantRequest{Fqn: data.Id.ValueString()}); err != nil {
		resp.Diagnostics.AddError("Error deleting Tenant", "DeleteTenant request failed: "+err.Error())
		return
	}
}

func (r *TenantResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
