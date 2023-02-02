package provider

import (
	"context"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tsbv2 "github.com/tetrateio/api/tsb/v2"

	"github.com/tetrateio/tetrate/pkg/api"
	"github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &ServiceAccount{}
var _ resource.ResourceWithImportState = &ServiceAccount{}

func NewServiceAccount() resource.Resource {
	return &ServiceAccount{}
}

// ServiceAccount defines the resource implementation.
type ServiceAccount struct {
	client tsbv2.TeamsClient
}

// Schema implements resource.Resource
func (*ServiceAccount) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	s, d := tsbv2.GenSchemaCreateServiceAccountRequest(ctx)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Schema = s
}

func (r *ServiceAccount) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_account"
}

func (r *ServiceAccount) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	clients := helpers.BuildClientsResource(ctx, req, resp)
	if resp.Diagnostics.HasError() || clients == nil {
		return
	}
	r.client = clients.Team
}

func (r *ServiceAccount) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	spew.Dump(req.Plan)

	data := tsbv2.NewCreateServiceAccountRequestModelFromPlan(ctx, req.Plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	spew.Dump(data)

	serviceAccount, err := r.client.CreateServiceAccount(ctx, data.ToGo(ctx))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating ServiceAccount",
			"Could not create serviceAccount: "+err.Error(),
		)
		return
	}

	data.LoadFromResult(ctx, serviceAccount.Fqn, &tsbv2.CreateServiceAccountRequest{
		Parent:         data.Parent.ValueString(),
		Name:           data.Name.ValueString(),
		ServiceAccount: serviceAccount,
	})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceAccount) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	data := tsbv2.NewCreateServiceAccountRequestModelFromState(ctx, req.State, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	fqn := data.Id.ValueString()
	meta, err := helpers.FromFQN(api.ServiceAccountKind, fqn)
	if err != nil {
		resp.Diagnostics.AddError("Error reading ServiceAccount", "Could not parse fqn: "+err.Error())
		return
	}

	serviceAccount, err := r.client.GetServiceAccount(ctx, &tsbv2.GetServiceAccountRequest{Fqn: fqn})
	if err != nil {
		resp.Diagnostics.AddError("Error reading ServiceAccount", "GetServiceAccount request failed: "+err.Error())
		return
	}

	data.LoadFromResult(ctx, fqn, &tsbv2.CreateServiceAccountRequest{
		Parent:         helpers.ParentFQN(api.TenantKind, meta),
		Name:           meta.Name,
		ServiceAccount: serviceAccount,
	})

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceAccount) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	data := tsbv2.NewCreateServiceAccountRequestModelFromPlan(ctx, req.Plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// We need to run a get to retreive the latest etag and set fqn as its not going to be in the plan (only state)
	fqn := data.Id.ValueString()
	serviceAccount, err := r.client.GetServiceAccount(ctx, &tsbv2.GetServiceAccountRequest{Fqn: fqn})
	if err != nil {
		resp.Diagnostics.AddError("Error updating ServiceAccount", "GetServiceAccount request failed: "+err.Error())
		return
	}
	data.ServiceAccount.Etag = types.StringValue(serviceAccount.Etag)
	data.ServiceAccount.Fqn = types.StringValue(serviceAccount.Fqn)

	// Do the actual update
	_, err = r.client.UpdateServiceAccount(ctx, data.ServiceAccount.ToGo(ctx))
	if err != nil {
		resp.Diagnostics.AddError("Error updating ServiceAccount", "UpdateServiceAccount request failed: "+err.Error())
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceAccount) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	data := tsbv2.NewCreateServiceAccountRequestModelFromState(ctx, req.State, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.client.DeleteServiceAccount(ctx, &tsbv2.DeleteServiceAccountRequest{Fqn: data.Id.ValueString()}); err != nil {
		resp.Diagnostics.AddError("Error deleting ServiceAccount", "DeleteServiceAccount request failed: "+err.Error())
		return
	}
}

func (r *ServiceAccount) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
