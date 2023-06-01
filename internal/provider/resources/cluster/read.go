package cluster

import (
	basetypes "basetypes"
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
	model.TrustDomain = types.StringValue(cluster.TrustDomain)
	model.TokenTtl = TokenTtl_Model{
		Nanos:   types.Int64Value(int64(cluster.TokenTtl.Nanos)),
		Seconds: types.Int64Value(int64(cluster.TokenTtl.Seconds)),
	}
	model.NamespaceScope = NamespaceScope_Model{
		Exceptions: func() basetypes.ListValue {
			r, diag := types.ListValueFrom(ctx, model.NamespaceScope.Exceptions.ElementType(ctx), cluster.NamespaceScope.Exceptions)
			resp.Diagnostics.Append(diag...)
			return r
		}(),
		Scope: types.StringValue(v2.NamespaceScoping_Scope_name[int32(cluster.NamespaceScope.Scope)]),
	}
	model.Locality = Locality_Model{Region: types.StringValue(cluster.Locality.Region)}
	model.ServiceAccount = ServiceAccount_Model{
		Description: types.StringValue(cluster.ServiceAccount.Description),
		DisplayName: types.StringValue(cluster.ServiceAccount.Displayname),
		Keys: func() []*ServiceAccount_Keys_Model {
			size := len(cluster.ServiceAccount.Keys)
			tmp := make([]*ServiceAccount_Keys_Model, size, size)
			for i, v := range cluster.ServiceAccount.Keys {
				tmp[i] = &ServiceAccount_Keys_Model{
					DefaultToken: types.StringValue(cluster.ServiceAccount.Keys.Defaulttoken),
					Encoding:     types.StringValue(v2.ServiceAccount_KeyPair_Encoding_name[int32(cluster.ServiceAccount.Keys.Encoding)]),
					Id:           types.StringValue(cluster.ServiceAccount.Keys.Id),
					PrivateKey:   types.StringValue(cluster.ServiceAccount.Keys.Privatekey),
					PublicKey:    types.StringValue(cluster.ServiceAccount.Keys.Publickey),
				}
			}
			// asdfasd
			return tmp
		}(),
	}
	model.State = State_Model{
		DiscoveredLocality: State_DiscoveredLocality_Model{Region: types.StringValue(cluster.State.Discoveredlocality.Region)},
		IstioVersions: func() basetypes.ListValue {
			r, diag := types.ListValueFrom(ctx, model.State.IstioVersions.ElementType(ctx), cluster.State.Istioversions)
			resp.Diagnostics.Append(diag...)
			return r
		}(),
		LastSyncTime: State_LastSyncTime_Model{
			Nanos:   types.Int64Value(int64(cluster.State.Lastsynctime.Nanos)),
			Seconds: types.Int64Value(int64(cluster.State.Lastsynctime.Seconds)),
		},
		Provider:     types.StringValue(cluster.State.Provider),
		TsbCpVersion: types.StringValue(cluster.State.Tsbcpversion),
		XcpVersion:   types.StringValue(cluster.State.Xcpversion),
	}
	model.Labels = func() basetypes.MapValue {
		r, diag := types.MapValueFrom(ctx, model.Labels.ElementType(ctx), cluster.Labels)
		resp.Diagnostics.Append(diag...)
		return r
	}()
	model.Network = types.StringValue(cluster.Network)
	model.DisplayName = types.StringValue(cluster.DisplayName)
	model.Tier1Cluster = Tier1Cluster_Model{Value: types.BoolValue(cluster.Tier1Cluster.Value)}
	model.Description = types.StringValue(cluster.Description)
	model.Namespaces = func() []*Namespaces_Model {
		size := len(cluster.Namespaces)
		tmp := make([]*Namespaces_Model, size, size)
		for i, v := range cluster.Namespaces {
			tmp[i] = &Namespaces_Model{
				Name: types.StringValue(cluster.Namespaces.Name),
				Services: func() []*Namespaces_Services_Model {
					size := len(cluster.Namespaces.Services)
					tmp := make([]*Namespaces_Services_Model, size, size)
					for i, v := range cluster.Namespaces.Services {
						tmp[i] = &Namespaces_Services_Model{
							CanonicalName: types.StringValue(cluster.Namespaces.Services.Canonicalname),
							GatewayHost:   types.BoolValue(cluster.Namespaces.Services.Gatewayhost),
							Hostname:      types.StringValue(cluster.Namespaces.Services.Hostname),
							KubernetesExternalAddresses: func() basetypes.ListValue {
								r, diag := types.ListValueFrom(ctx, model.Namespaces.Services.KubernetesExternalAddresses.ElementType(ctx), cluster.Namespaces.Services.Kubernetesexternaladdresses)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							KubernetesServiceFqdn:  types.StringValue(cluster.Namespaces.Services.Kubernetesservicefqdn),
							KubernetesServiceIp:    types.StringValue(cluster.Namespaces.Services.Kubernetesserviceip),
							MeshExternal:           types.BoolValue(cluster.Namespaces.Services.Meshexternal),
							Name:                   types.StringValue(cluster.Namespaces.Services.Name),
							Namespace:              types.StringValue(cluster.Namespaces.Services.Namespace),
							NumHops:                types.Int64Value(int64(cluster.Namespaces.Services.Numhops)),
							NumKubernetesEndpoints: types.Int64Value(int64(cluster.Namespaces.Services.Numkubernetesendpoints)),
							NumVmEndpoints:         types.Int64Value(int64(cluster.Namespaces.Services.Numvmendpoints)),
							Ports: func() []*Namespaces_Services_Ports_Model {
								size := len(cluster.Namespaces.Services.Ports)
								tmp := make([]*Namespaces_Services_Ports_Model, size, size)
								for i, v := range cluster.Namespaces.Services.Ports {
									tmp[i] = &Namespaces_Services_Ports_Model{
										KubernetesNodePort: types.Int64Value(int64(cluster.Namespaces.Services.Ports.Kubernetesnodeport)),
										Name:               types.StringValue(cluster.Namespaces.Services.Ports.Name),
										Number:             types.Int64Value(int64(cluster.Namespaces.Services.Ports.Number)),
									}
								}
								// asdfasd
								return tmp
							}(),
							Selector: func() basetypes.MapValue {
								r, diag := types.MapValueFrom(ctx, model.Namespaces.Services.Selector.ElementType(ctx), cluster.Namespaces.Services.Selector)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							SpiffeIds: func() basetypes.ListValue {
								r, diag := types.ListValueFrom(ctx, model.Namespaces.Services.SpiffeIds.ElementType(ctx), cluster.Namespaces.Services.Spiffeids)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							Subsets: func() basetypes.ListValue {
								r, diag := types.ListValueFrom(ctx, model.Namespaces.Services.Subsets.ElementType(ctx), cluster.Namespaces.Services.Subsets)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							Tier1GatewayHost: types.BoolValue(cluster.Namespaces.Services.Tier1gatewayhost),
							Workloads: func() []*Namespaces_Services_Workloads_Model {
								size := len(cluster.Namespaces.Services.Workloads)
								tmp := make([]*Namespaces_Services_Workloads_Model, size, size)
								for i, v := range cluster.Namespaces.Services.Workloads {
									tmp[i] = &Namespaces_Services_Workloads_Model{
										Address: types.StringValue(cluster.Namespaces.Services.Workloads.Address),
										IsVm:    types.BoolValue(cluster.Namespaces.Services.Workloads.Isvm),
										Name:    types.StringValue(cluster.Namespaces.Services.Workloads.Name),
										Proxy: Namespaces_Services_Workloads_Proxy_Model{
											ControlPlaneAddress: types.StringValue(cluster.Namespaces.Services.Workloads.Proxy.Controlplaneaddress),
											EnvoyVersion:        types.StringValue(cluster.Namespaces.Services.Workloads.Proxy.Envoyversion),
											IstioVersion:        types.StringValue(cluster.Namespaces.Services.Workloads.Proxy.Istioversion),
											Status: func() basetypes.MapValue {
												r, diag := types.MapValueFrom(ctx, model.Namespaces.Services.Workloads.Proxy.Status.ElementType(ctx), cluster.Namespaces.Services.Workloads.Proxy.Status)
												resp.Diagnostics.Append(diag...)
												return r
											}(),
										},
									}
								}
								// asdfasd
								return tmp
							}(),
						}
					}
					// asdfasd
					return tmp
				}(),
			}
		}
		// asdfasd
		return tmp
	}()
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
