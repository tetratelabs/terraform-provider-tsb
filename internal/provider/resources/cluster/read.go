package cluster

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	basetypes "github.com/hashicorp/terraform-plugin-framework/types/basetypes"
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
	model.ServiceAccount = ServiceAccount_Model{
		Description: types.StringValue(cluster.ServiceAccount.Description),
		DisplayName: types.StringValue(cluster.ServiceAccount.DisplayName),
		Keys: func() []*ServiceAccount_Keys_Model {
			size := len(cluster.ServiceAccount.Keys)
			tmp := make([]*ServiceAccount_Keys_Model, size, size)
			for i, keys := range cluster.ServiceAccount.Keys {
				tmp[i] = &ServiceAccount_Keys_Model{
					DefaultToken: types.StringValue(keys.DefaultToken),
					Encoding:     types.StringValue(v2.ServiceAccount_KeyPair_Encoding_name[int32(keys.Encoding)]),
					Id:           types.StringValue(keys.Id),
					PrivateKey:   types.StringValue(keys.PrivateKey),
					PublicKey:    types.StringValue(keys.PublicKey),
				}
			}
			return tmp
		}(),
	}
	model.TrustDomain = types.StringValue(cluster.TrustDomain)
	model.Locality = Locality_Model{Region: types.StringValue(cluster.Locality.Region)}
	model.TokenTtl = TokenTtl_Model{
		Nanos:   types.Int64Value(int64(cluster.TokenTtl.Nanos)),
		Seconds: types.Int64Value(int64(cluster.TokenTtl.Seconds)),
	}
	model.State = State_Model{
		DiscoveredLocality: State_DiscoveredLocality_Model{Region: types.StringValue(cluster.State.DiscoveredLocality.Region)},
		IstioVersions: func() basetypes.ListValue {
			r, diag := types.ListValueFrom(ctx, basetypes.StringType{}, cluster.State.IstioVersions)
			resp.Diagnostics.Append(diag...)
			return r
		}(),
		LastSyncTime: State_LastSyncTime_Model{
			Nanos:   types.Int64Value(int64(cluster.State.LastSyncTime.Nanos)),
			Seconds: types.Int64Value(int64(cluster.State.LastSyncTime.Seconds)),
		},
		Provider:     types.StringValue(cluster.State.Provider),
		TsbCpVersion: types.StringValue(cluster.State.TsbCpVersion),
		XcpVersion:   types.StringValue(cluster.State.XcpVersion),
	}
	model.Description = types.StringValue(cluster.Description)
	model.DisplayName = types.StringValue(cluster.DisplayName)
	model.Network = types.StringValue(cluster.Network)
	model.Namespaces = func() []*Namespaces_Model {
		size := len(cluster.Namespaces)
		tmp := make([]*Namespaces_Model, size, size)
		for i, namespaces := range cluster.Namespaces {
			tmp[i] = &Namespaces_Model{
				Name: types.StringValue(namespaces.Name),
				Services: func() []*Namespaces_Services_Model {
					size := len(namespaces.Services)
					tmp := make([]*Namespaces_Services_Model, size, size)
					for i, services := range namespaces.Services {
						tmp[i] = &Namespaces_Services_Model{
							CanonicalName: types.StringValue(services.CanonicalName),
							GatewayHost:   types.BoolValue(services.GatewayHost),
							Hostname:      types.StringValue(services.Hostname),
							KubernetesExternalAddresses: func() basetypes.ListValue {
								r, diag := types.ListValueFrom(ctx, basetypes.StringType{}, services.KubernetesExternalAddresses)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							KubernetesServiceFqdn:  types.StringValue(services.KubernetesServiceFqdn),
							KubernetesServiceIp:    types.StringValue(services.KubernetesServiceIp),
							MeshExternal:           types.BoolValue(services.MeshExternal),
							Name:                   types.StringValue(services.Name),
							Namespace:              types.StringValue(services.Namespace),
							NumHops:                types.Int64Value(int64(services.NumHops)),
							NumKubernetesEndpoints: types.Int64Value(int64(services.NumKubernetesEndpoints)),
							NumVmEndpoints:         types.Int64Value(int64(services.NumVmEndpoints)),
							Ports: func() []*Namespaces_Services_Ports_Model {
								size := len(services.Ports)
								tmp := make([]*Namespaces_Services_Ports_Model, size, size)
								for i, ports := range services.Ports {
									tmp[i] = &Namespaces_Services_Ports_Model{
										KubernetesNodePort: types.Int64Value(int64(ports.KubernetesNodePort)),
										Name:               types.StringValue(ports.Name),
										Number:             types.Int64Value(int64(ports.Number)),
									}
								}
								return tmp
							}(),
							Selector: func() basetypes.MapValue {
								r, diag := types.MapValueFrom(ctx, basetypes.StringType{}, services.Selector)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							SpiffeIds: func() basetypes.ListValue {
								r, diag := types.ListValueFrom(ctx, basetypes.StringType{}, services.SpiffeIds)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							Subsets: func() basetypes.ListValue {
								r, diag := types.ListValueFrom(ctx, basetypes.StringType{}, services.Subsets)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							Tier1GatewayHost: types.BoolValue(services.Tier1GatewayHost),
							Workloads: func() []*Namespaces_Services_Workloads_Model {
								size := len(services.Workloads)
								tmp := make([]*Namespaces_Services_Workloads_Model, size, size)
								for i, workloads := range services.Workloads {
									tmp[i] = &Namespaces_Services_Workloads_Model{
										Address: types.StringValue(workloads.Address),
										IsVm:    types.BoolValue(workloads.IsVm),
										Name:    types.StringValue(workloads.Name),
										Proxy: Namespaces_Services_Workloads_Proxy_Model{
											ControlPlaneAddress: types.StringValue(workloads.Proxy.ControlPlaneAddress),
											EnvoyVersion:        types.StringValue(workloads.Proxy.EnvoyVersion),
											IstioVersion:        types.StringValue(workloads.Proxy.IstioVersion),
											Status: func() basetypes.MapValue {
												r, diag := types.MapValueFrom(ctx, basetypes.StringType{}, workloads.Proxy.Status)
												resp.Diagnostics.Append(diag...)
												return r
											}(),
										},
									}
								}
								return tmp
							}(),
						}
					}
					return tmp
				}(),
			}
		}
		return tmp
	}()
	model.Labels = func() basetypes.MapValue {
		r, diag := types.MapValueFrom(ctx, basetypes.StringType{}, cluster.Labels)
		resp.Diagnostics.Append(diag...)
		return r
	}()
	model.NamespaceScope = NamespaceScope_Model{
		Exceptions: func() basetypes.ListValue {
			r, diag := types.ListValueFrom(ctx, basetypes.StringType{}, cluster.NamespaceScope.Exceptions)
			resp.Diagnostics.Append(diag...)
			return r
		}(),
		Scope: types.StringValue(v2.NamespaceScoping_Scope_name[int32(cluster.NamespaceScope.Scope)]),
	}
	model.Tier1Cluster = Tier1Cluster_Model{Value: types.BoolValue(cluster.Tier1Cluster.Value)}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
