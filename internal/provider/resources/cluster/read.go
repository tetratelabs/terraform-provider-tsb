package cluster

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	v2 "github.com/tetrateio/api/tsb/v2"
	api "github.com/tetrateio/tetrate/pkg/api"
	helpers "github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

func (r *ClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var model ClusterModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request := &v2.GetClusterRequest{Fqn: model.Id.ValueString()}
	cluster, err := r.client.GetCluster(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Cluster", "GetClusterRequest failed: "+err.Error())
		return
	}
	meta, err := helpers.FromFQN(api.ClusterKind, model.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error readingCluster", "FQN parsing failed: "+err.Error())
		return
	}
	model.Id = types.StringValue(cluster.Fqn)
	model.Name = types.StringValue(meta.Name)
	model.Parent = types.StringValue(helpers.ParentFQN(api.ClusterKind, meta))
	model.Namespaces = nil
	model.TokenTtl = nil
	model.DisplayName = types.StringValue(cluster.DisplayName)
	model.InstallTemplate = nil
	model.ServiceAccount = nil
	model.State = nil
	model.TrustDomain = types.StringValue(cluster.TrustDomain)
	model.Locality = nil
	model.NamespaceScope = nil
	model.Description = types.StringValue(cluster.Description)
	model.Tier1Cluster = nil
	model.Labels = nil
	model.Network = types.StringValue(cluster.Network)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
