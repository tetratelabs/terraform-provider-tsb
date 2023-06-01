package cluster

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	v2 "github.com/tetrateio/api/tsb/v2"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func ptrify[T any](v T) *T {
	return &v
}
func (r *ClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model ClusterModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request := &v2.CreateClusterRequest{
		Cluster: &v2.Cluster{
			Description: model.Description.ValueString(),
			DisplayName: model.DisplayName.ValueString(),
			Labels: func() map[string]string {
				tmp := make(map[string]string)
				resp.Diagnostics.Append(model.Labels.ElementsAs(ctx, &tmp, false)...)
				return tmp
			}(),
			Locality: &v2.Locality{Region: model.Locality.Region.ValueString()},
			NamespaceScope: &v2.NamespaceScoping{
				Exceptions: func() []string {
					tmp := make([]string, 0, len(model.NamespaceScope.Exceptions.Elements()))
					resp.Diagnostics.Append(model.NamespaceScope.Exceptions.ElementsAs(ctx, &tmp, false)...)
					return tmp
				}(),
				Scope: v2.NamespaceScoping_Scope(v2.NamespaceScoping_Scope_value[model.NamespaceScope.Scope.ValueString()]),
			},
			Namespaces: func() []*v2.Namespace {
				a := make([]*v2.Namespace, 0, len(model.Namespaces))
				for i, namespaces := range model.Namespaces {
					a[i] = &v2.Namespace{
						Name: namespaces.Name.ValueString(),
						Services: func() []*v2.Service {
							a := make([]*v2.Service, 0, len(namespaces.Services))
							for i, services := range namespaces.Services {
								a[i] = &v2.Service{
									CanonicalName: services.CanonicalName.ValueString(),
									GatewayHost:   services.GatewayHost.ValueBool(),
									Hostname:      services.Hostname.ValueString(),
									KubernetesExternalAddresses: func() []string {
										tmp := make([]string, 0, len(services.KubernetesExternalAddresses.Elements()))
										resp.Diagnostics.Append(services.KubernetesExternalAddresses.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									KubernetesServiceFqdn:  services.KubernetesServiceFqdn.ValueString(),
									KubernetesServiceIp:    services.KubernetesServiceIp.ValueString(),
									MeshExternal:           services.MeshExternal.ValueBool(),
									Name:                   services.Name.ValueString(),
									Namespace:              services.Namespace.ValueString(),
									NumHops:                uint32(services.NumHops.ValueInt64()),
									NumKubernetesEndpoints: uint32(services.NumKubernetesEndpoints.ValueInt64()),
									NumVmEndpoints:         uint32(services.NumVmEndpoints.ValueInt64()),
									Ports: func() []*v2.Service_Port {
										a := make([]*v2.Service_Port, 0, len(services.Ports))
										for i, ports := range services.Ports {
											a[i] = &v2.Service_Port{
												KubernetesNodePort: uint32(ports.KubernetesNodePort.ValueInt64()),
												Name:               ports.Name.ValueString(),
												Number:             uint32(ports.Number.ValueInt64()),
											}
										}
										return a
									}(),
									Selector: func() map[string]string {
										tmp := make(map[string]string)
										resp.Diagnostics.Append(services.Selector.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									SpiffeIds: func() []string {
										tmp := make([]string, 0, len(services.SpiffeIds.Elements()))
										resp.Diagnostics.Append(services.SpiffeIds.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									Subsets: func() []string {
										tmp := make([]string, 0, len(services.Subsets.Elements()))
										resp.Diagnostics.Append(services.Subsets.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									Tier1GatewayHost: services.Tier1GatewayHost.ValueBool(),
									Workloads: func() []*v2.Workload {
										a := make([]*v2.Workload, 0, len(services.Workloads))
										for i, workloads := range services.Workloads {
											a[i] = &v2.Workload{
												Address: workloads.Address.ValueString(),
												IsVm:    workloads.IsVm.ValueBool(),
												Name:    workloads.Name.ValueString(),
												Proxy: &v2.Workload_Proxy{
													ControlPlaneAddress: workloads.Proxy.ControlPlaneAddress.ValueString(),
													EnvoyVersion:        workloads.Proxy.EnvoyVersion.ValueString(),
													IstioVersion:        workloads.Proxy.IstioVersion.ValueString(),
													Status: func() map[string]string {
														tmp := make(map[string]string)
														resp.Diagnostics.Append(workloads.Proxy.Status.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
												},
											}
										}
										return a
									}(),
								}
							}
							return a
						}(),
					}
				}
				return a
			}(),
			Network: model.Network.ValueString(),
			ServiceAccount: &v2.ServiceAccount{
				Description: model.ServiceAccount.Description.ValueString(),
				DisplayName: model.ServiceAccount.DisplayName.ValueString(),
				Keys: func() []*v2.ServiceAccount_KeyPair {
					a := make([]*v2.ServiceAccount_KeyPair, 0, len(model.ServiceAccount.Keys))
					for i, keys := range model.ServiceAccount.Keys {
						a[i] = &v2.ServiceAccount_KeyPair{
							DefaultToken: keys.DefaultToken.ValueString(),
							Encoding:     v2.ServiceAccount_KeyPair_Encoding(v2.ServiceAccount_KeyPair_Encoding_value[keys.Encoding.ValueString()]),
							Id:           keys.Id.ValueString(),
							PrivateKey:   keys.PrivateKey.ValueString(),
							PublicKey:    keys.PublicKey.ValueString(),
						}
					}
					return a
				}(),
			},
			State: &v2.Cluster_State{
				DiscoveredLocality: &v2.Locality{Region: model.State.DiscoveredLocality.Region.ValueString()},
				IstioVersions: func() []string {
					tmp := make([]string, 0, len(model.State.IstioVersions.Elements()))
					resp.Diagnostics.Append(model.State.IstioVersions.ElementsAs(ctx, &tmp, false)...)
					return tmp
				}(),
				LastSyncTime: &timestamppb.Timestamp{
					Nanos:   int32(model.State.LastSyncTime.Nanos.ValueInt64()),
					Seconds: model.State.LastSyncTime.Seconds.ValueInt64(),
				},
				Provider:     model.State.Provider.ValueString(),
				TsbCpVersion: model.State.TsbCpVersion.ValueString(),
				XcpVersion:   model.State.XcpVersion.ValueString(),
			},
			Tier1Cluster: &wrapperspb.BoolValue{Value: model.Tier1Cluster.Value.ValueBool()},
			TokenTtl: &durationpb.Duration{
				Nanos:   int32(model.TokenTtl.Nanos.ValueInt64()),
				Seconds: model.TokenTtl.Seconds.ValueInt64(),
			},
			TrustDomain: model.TrustDomain.ValueString(),
		},
		Name:   model.Name.ValueString(),
		Parent: model.Parent.ValueString(),
	}
	cluster, err := r.client.CreateCluster(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Cluster", "CreateCluster request failed: "+err.Error())
		return
	}
	model.Id = types.StringValue(cluster.Fqn)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
