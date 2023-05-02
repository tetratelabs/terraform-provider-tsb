package cluster

import (
	basetypes "basetypes"
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	v2 "github.com/tetrateio/api/tsb/v2"
	api "github.com/tetrateio/tetrate/pkg/api"
	helpers "github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
	pkgimportpath "pkgimportpath"
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
	model.Tier1Cluster = nil
	model.Locality = nil
	model.ServiceAccount = nil
	model.TokenTtl = nil
	model.Description = types.StringValue(pkgimportpath.Cluster_name[int32(rLowerName.Cluster)])
	model.Network = types.StringValue(pkgimportpath.Cluster_name[int32(rLowerName.Cluster)])
	model.State = nil
	model.DisplayName = types.StringValue(pkgimportpath.Cluster_name[int32(rLowerName.Cluster)])
	model.NamespaceScope = nil
	model.Namespaces = func() []*namespaces {
		tmp := make([]*namespaces, len(cluster))
		for _, n := range cluster {
			tmp = append(tmp, "magicalRecursion")
		}
		return tmp
	}()
	model.InstallTemplate = nil
	model.Labels = func() basetypes.MapValue {
		r, diag := types.MapValue(ctx, labels.ElementType(ctx), cluster)
		resp.Diagnostics.Append(diag...)
		return r
	}()
	model.TrustDomain = types.StringValue(pkgimportpath.Cluster_name[int32(rLowerName.Cluster)])
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
