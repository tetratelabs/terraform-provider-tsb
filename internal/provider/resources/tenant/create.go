package tenant

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	v21 "github.com/tetrateio/api/tsb/types/v2"
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
			ConfigGenerationMetadata: &v21.ConfigGenerationMetadata{
				Annotations: func() map[string]string {
					tmp := make(map[string]string)
					resp.Diagnostics.Append(model.ConfigGenerationMetadata.Annotations.ElementsAs(ctx, &tmp, false)...)
					return tmp
				}(),
				Labels: func() map[string]string {
					tmp := make(map[string]string)
					resp.Diagnostics.Append(model.ConfigGenerationMetadata.Labels.ElementsAs(ctx, &tmp, false)...)
					return tmp
				}(),
			},
			DeletionProtectionEnabled: model.DeletionProtectionEnabled.ValueBool(),
			Description:               model.Description.ValueString(),
			DisplayName:               model.DisplayName.ValueString(),
			SecurityDomain:            model.SecurityDomain.ValueString(),
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
