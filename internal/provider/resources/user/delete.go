package user

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	v2 "github.com/tetrateio/api/tsb/v2"
)

func (r *UserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var model UserModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.client.DeleteUser(ctx, &v2.DeleteUserRequest{Fqn: model.Id.ValueString()}); err != nil {
		resp.Diagnostics.AddError("Error deleting User", "DeleteUser request failed: "+err.Error())
		return
	}
}
