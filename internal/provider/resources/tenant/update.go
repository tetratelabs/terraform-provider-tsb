package tenant

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	v21 "github.com/tetrateio/api/tsb/types/v2"
	v2 "github.com/tetrateio/api/tsb/v2"
)

func (r *TenantResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var model TenantModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request := &v2.GetTenantRequest{Fqn: model.Id.ValueString()}
	tenant, err := r.client.GetTenant(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Tenant", "GetTenantRequest failed: "+err.Error())
		return
	}
	updateTo := &v2.Tenant{
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
		Etag:                      tenant.Etag,
		Fqn:                       model.Id.ValueString(),
		SecurityDomain:            model.SecurityDomain.ValueString(),
	}
	_, err = r.client.UpdateTenant(ctx, updateTo)
	if err != nil {
		resp.Diagnostics.AddError("Error updating Tenant", "UpdateTenant failed: "+err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
