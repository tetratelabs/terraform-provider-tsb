package tenant

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	v2 "github.com/tetrateio/api/tsb/v2"
)

func (r *TenantResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var model TenantModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.client.DeleteTenant(ctx, &v2.DeleteTenantRequest{Fqn: model.Id.ValueString()}); err != nil {
		resp.Diagnostics.AddError("Error deleting Tenant", "DeleteTenant request failed: "+err.Error())
		return
	}
}
