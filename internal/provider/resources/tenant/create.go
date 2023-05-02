package tenant

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	v2 "github.com/tetrateio/api/tsb/v2"
)

func ptrify[T any](v T) *T {
	return &v
}
func (r *TenantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model TenantModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request := &v2.CreateTenantRequest{
		Name:   model.Name.ValueString(),
		Parent: model.Parent.ValueString(),
		Tenant: &v2.Tenant{
			Description:    model.Description.ValueString(),
			DisplayName:    model.DisplayName.ValueString(),
			SecurityDomain: model.SecurityDomain.ValueString(),
		},
	}
	tenant, err := r.client.CreateTenant(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Tenant", "CreateTenant request failed: "+err.Error())
		return
	}
	model.Id = types.StringValue(tenant.Fqn)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
