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
var _ resource.Resource = &TeamResource{}
var _ resource.ResourceWithImportState = &TeamResource{}

func NewTeamResource() resource.Resource {
	return &TeamResource{}
}

// TeamResource defines the resource implementation.
type TeamResource struct {
	client tsbv2.TeamsClient
}

// Schema implements resource.Resource
func (*TeamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	s, d := tsbv2.GenSchemaCreateTeamRequest(ctx)
	resp.Diagnostics.Append(d...)
	resp.Schema = s
}

func (r *TeamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_team"
}

func (r *TeamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	clients := helpers.BuildClientsResource(ctx, req, resp)
	if resp.Diagnostics.HasError() || clients == nil {
		return
	}
	r.client = clients.Team
}

func (r *TeamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	data := tsbv2.NewCreateTeamRequestModelFromPlan(ctx, req.Plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Override source type
	data.Team.SourceType = types.Int64Value(int64(tsbv2.SourceType_MANUAL))

	team, err := r.client.CreateTeam(ctx, data.ToGo(ctx))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating Team",
			"Could not create team: "+err.Error(),
		)
		return
	}

	data.LoadFromResult(ctx, team.Fqn, &tsbv2.CreateTeamRequest{
		Parent: data.Parent.ValueString(),
		Name:   data.Name.ValueString(),
		Team:   team,
	})
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TeamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	data := tsbv2.NewCreateTeamRequestModelFromState(ctx, req.State, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	fqn := data.Id.ValueString()
	meta, err := helpers.FromFQN(api.TeamKind, fqn)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Team", "Could not parse fqn: "+err.Error())
		return
	}

	team, err := r.client.GetTeam(ctx, &tsbv2.GetTeamRequest{Fqn: fqn})
	if err != nil {
		resp.Diagnostics.AddError("Error reading Team", "GetTeam request failed: "+err.Error())
		return
	}

	data.LoadFromResult(ctx, fqn, &tsbv2.CreateTeamRequest{
		Parent: helpers.ParentFQN(api.TenantKind, meta),
		Name:   meta.Name,
		Team:   team,
	})

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TeamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	data := tsbv2.NewCreateTeamRequestModelFromPlan(ctx, req.Plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// We need to run a get to retreive the latest etag and set fqn + source_type as its not going to be in the plan (only state)
	fqn := data.Id.ValueString()
	team, err := r.client.GetTeam(ctx, &tsbv2.GetTeamRequest{Fqn: fqn})
	if err != nil {
		resp.Diagnostics.AddError("Error updating Team", "GetTeam request failed: "+err.Error())
		return
	}
	data.Team.Etag = types.StringValue(team.Etag)
	data.Team.Fqn = types.StringValue(team.Fqn)
	data.Team.SourceType = types.Int64Value(int64(tsbv2.SourceType_MANUAL))

	// Do the actual update
	_, err = r.client.UpdateTeam(ctx, data.Team.ToGo(ctx))
	if err != nil {
		resp.Diagnostics.AddError("Error updating Team", "UpdateTeam request failed: "+err.Error())
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TeamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	data := tsbv2.NewCreateTeamRequestModelFromState(ctx, req.State, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.client.DeleteTeam(ctx, &tsbv2.DeleteTeamRequest{Fqn: data.Id.ValueString()}); err != nil {
		resp.Diagnostics.AddError("Error deleting Team", "DeleteTeam request failed: "+err.Error())
		return
	}
}

func (r *TeamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
