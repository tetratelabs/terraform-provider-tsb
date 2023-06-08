package serviceaccount

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	v2 "github.com/tetrateio/api/tsb/v2"
)

func (r *ServiceAccountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var model ServiceAccountModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.client.DeleteServiceAccount(ctx, &v2.DeleteServiceAccountRequest{Fqn: model.Id.ValueString()}); err != nil {
		resp.Diagnostics.AddError("Error deleting ServiceAccount", "DeleteServiceAccount request failed: "+err.Error())
		return
	}
}
