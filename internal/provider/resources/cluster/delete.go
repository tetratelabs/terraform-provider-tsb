package cluster

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	v2 "github.com/tetrateio/api/tsb/v2"
)

func (r *ClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var model ClusterModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if _, err := r.client.DeleteCluster(ctx, &v2.DeleteClusterRequest{Fqn: model.Id.ValueString()}); err != nil {
		resp.Diagnostics.AddError("Error deleting Cluster", "DeleteCluster request failed: "+err.Error())
		return
	}
}
