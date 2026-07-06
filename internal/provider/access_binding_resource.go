// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	rbac "github.com/tetrateio/tetrate/api/tsb/rbac/v2"
)

// tsb_access_binding manages the RBAC AccessBindings (AccessPolicy) of a TSB
// resource. Unlike the generated resources this is hand-written, because the
// Policy API is Get/Set-based rather than the standard CRUD quartet: an access
// policy implicitly exists for every resource and cannot be created or deleted,
// only read and replaced.
//
// This resource is authoritative for the target resource's bindings: it owns the
// entire `allow` list. Delete clears the bindings (SetPolicy with an empty list)
// since the policy object itself cannot be removed.

var (
	_ resource.Resource                = (*accessBindingResource)(nil)
	_ resource.ResourceWithConfigure   = (*accessBindingResource)(nil)
	_ resource.ResourceWithImportState = (*accessBindingResource)(nil)
)

// NewAccessBindingResource constructs the tsb_access_binding resource.
func NewAccessBindingResource() resource.Resource {
	return &accessBindingResource{}
}

type accessBindingResource struct {
	client *client
}

type accessBindingModel struct {
	ID          types.String             `tfsdk:"id"`
	FQN         types.String             `tfsdk:"fqn"`
	Description types.String             `tfsdk:"description"`
	Allow       []accessBindingRuleModel `tfsdk:"allow"`
}

type accessBindingRuleModel struct {
	Role     types.String                `tfsdk:"role"`
	Subjects []accessBindingSubjectModel `tfsdk:"subjects"`
}

type accessBindingSubjectModel struct {
	User           types.String `tfsdk:"user"`
	Team           types.String `tfsdk:"team"`
	ServiceAccount types.String `tfsdk:"service_account"`
}

func (r *accessBindingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_access_binding"
}

func (r *accessBindingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manages the RBAC access bindings (AccessPolicy) of a TSB resource. " +
			"Authoritative: it owns the complete set of bindings for the target resource; " +
			"destroying it clears them (the policy object itself always exists).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Fully-qualified name of the resource whose bindings are managed.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"fqn": schema.StringAttribute{
				Required:      true,
				Description:   "Fully-qualified name of the TSB resource these bindings apply to (e.g. organizations/myorg/tenants/mytenant).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Description: "A description of the access policy.",
			},
			"allow": schema.ListNestedAttribute{
				Required:    true,
				Description: "The set of role-to-subject bindings granted on the resource.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"role": schema.StringAttribute{
							Required:    true,
							Description: "FQN of the role being granted.",
						},
						"subjects": schema.ListNestedAttribute{
							Required:    true,
							Description: "The subjects the role is granted to. Set exactly one of user, team, or service_account per subject.",
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"user":            schema.StringAttribute{Optional: true, Description: "User the role is granted to."},
									"team":            schema.StringAttribute{Optional: true, Description: "Team the role is granted to."},
									"service_account": schema.StringAttribute{Optional: true, Description: "Service account the role is granted to."},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *accessBindingResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected provider data", fmt.Sprintf("Expected *client, got %T. This is a provider bug.", req.ProviderData))
		return
	}
	r.client = c
}

func (r *accessBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan accessBindingModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.set(ctx, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.ID = plan.FQN
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *accessBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state accessBindingModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	policy, err := rbac.NewPolicyClient(r.client.conn).GetPolicy(ctx, &rbac.GetPolicyRequest{Fqn: state.FQN.ValueString()})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Unable to read access bindings", err.Error())
		return
	}

	state.ID = state.FQN
	state.Description = stringOrNull(policy.Description)
	state.Allow = flattenBindings(policy.Allow)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *accessBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan accessBindingModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	r.set(ctx, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.ID = plan.FQN
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *accessBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state accessBindingModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	// The policy cannot be deleted, only emptied: clear the managed bindings.
	_, err := rbac.NewPolicyClient(r.client.conn).SetPolicy(ctx, &rbac.AccessPolicy{Fqn: state.FQN.ValueString()})
	if err != nil && status.Code(err) != codes.NotFound {
		resp.Diagnostics.AddError("Unable to clear access bindings", err.Error())
	}
}

func (r *accessBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("fqn"), req.ID)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}

// set replaces the target resource's policy with the planned bindings.
func (r *accessBindingResource) set(ctx context.Context, plan *accessBindingModel, diags *diag.Diagnostics) {
	policy := &rbac.AccessPolicy{
		Fqn:         plan.FQN.ValueString(),
		Description: plan.Description.ValueString(),
		Allow:       expandBindings(plan.Allow),
	}
	if _, err := rbac.NewPolicyClient(r.client.conn).SetPolicy(ctx, policy); err != nil {
		diags.AddError("Unable to set access bindings", err.Error())
	}
}

func expandBindings(in []accessBindingRuleModel) []*rbac.Binding {
	if len(in) == 0 {
		return nil
	}
	out := make([]*rbac.Binding, 0, len(in))
	for _, b := range in {
		out = append(out, &rbac.Binding{
			Role:     b.Role.ValueString(),
			Subjects: expandSubjects(b.Subjects),
		})
	}
	return out
}

func expandSubjects(in []accessBindingSubjectModel) []*rbac.Subject {
	if len(in) == 0 {
		return nil
	}
	out := make([]*rbac.Subject, 0, len(in))
	for _, s := range in {
		switch {
		case !s.User.IsNull() && !s.User.IsUnknown():
			out = append(out, &rbac.Subject{Sub: &rbac.Subject_User{User: s.User.ValueString()}})
		case !s.Team.IsNull() && !s.Team.IsUnknown():
			out = append(out, &rbac.Subject{Sub: &rbac.Subject_Team{Team: s.Team.ValueString()}})
		case !s.ServiceAccount.IsNull() && !s.ServiceAccount.IsUnknown():
			out = append(out, &rbac.Subject{Sub: &rbac.Subject_ServiceAccount{ServiceAccount: s.ServiceAccount.ValueString()}})
		}
	}
	return out
}

func flattenBindings(in []*rbac.Binding) []accessBindingRuleModel {
	if len(in) == 0 {
		return nil
	}
	out := make([]accessBindingRuleModel, 0, len(in))
	for _, b := range in {
		out = append(out, accessBindingRuleModel{
			Role:     types.StringValue(b.GetRole()),
			Subjects: flattenSubjects(b.GetSubjects()),
		})
	}
	return out
}

func flattenSubjects(in []*rbac.Subject) []accessBindingSubjectModel {
	if len(in) == 0 {
		return nil
	}
	out := make([]accessBindingSubjectModel, 0, len(in))
	for _, s := range in {
		m := accessBindingSubjectModel{
			User:           types.StringNull(),
			Team:           types.StringNull(),
			ServiceAccount: types.StringNull(),
		}
		switch v := s.GetSub().(type) {
		case *rbac.Subject_User:
			m.User = types.StringValue(v.User)
		case *rbac.Subject_Team:
			m.Team = types.StringValue(v.Team)
		case *rbac.Subject_ServiceAccount:
			m.ServiceAccount = types.StringValue(v.ServiceAccount)
		}
		out = append(out, m)
	}
	return out
}
