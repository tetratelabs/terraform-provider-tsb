package cluster

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	v1alpha11 "github.com/tetrateio/api/install/helm/common/v1alpha1"
	v1alpha1 "github.com/tetrateio/api/install/helm/controlplane/v1alpha1"
	v2 "github.com/tetrateio/api/tsb/v2"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

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
			InstallTemplate: &v2.Cluster_InstallTemplate{
				Helm: &v1alpha1.Values{
					Image: &v1alpha11.Image{
						Registry: model.InstallTemplate.Helm.Image.Registry.ValueString(),
						Tag:      model.InstallTemplate.Helm.Image.Tag.ValueString(),
					},
					Operator: &v1alpha11.Operator{
						Deployment: &v1alpha11.Operator_Deployment{
							Affinity: &anypb.Any{
								TypeUrl: model.InstallTemplate.Helm.Operator.Deployment.Affinity.TypeUrl.ValueString(),
								Value:   []byte(model.InstallTemplate.Helm.Operator.Deployment.Affinity.Value.ValueString()),
							},
							Annotations: func() map[string]string {
								tmp := make(map[string]string)
								resp.Diagnostics.Append(model.InstallTemplate.Helm.Operator.Deployment.Annotations.ElementsAs(ctx, &tmp, false)...)
								return tmp
							}(),
							Env: func() []*anypb.Any {
								a := make([]*anypb.Any, len(model.InstallTemplate.Helm.Operator.Deployment.Env))
								for i, env := range model.InstallTemplate.Helm.Operator.Deployment.Env {
									a[i] = &anypb.Any{
										TypeUrl: env.TypeUrl.ValueString(),
										Value:   []byte(env.Value.ValueString()),
									}
								}
								return a
							}(),
							PodAnnotations: func() map[string]string {
								tmp := make(map[string]string)
								resp.Diagnostics.Append(model.InstallTemplate.Helm.Operator.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
								return tmp
							}(),
							ReplicaCount: int32(model.InstallTemplate.Helm.Operator.Deployment.ReplicaCount.ValueInt64()),
							Strategy: &anypb.Any{
								TypeUrl: model.InstallTemplate.Helm.Operator.Deployment.Strategy.TypeUrl.ValueString(),
								Value:   []byte(model.InstallTemplate.Helm.Operator.Deployment.Strategy.Value.ValueString()),
							},
							Tolerations: func() []*anypb.Any {
								a := make([]*anypb.Any, len(model.InstallTemplate.Helm.Operator.Deployment.Tolerations))
								for i, tolerations := range model.InstallTemplate.Helm.Operator.Deployment.Tolerations {
									a[i] = &anypb.Any{
										TypeUrl: tolerations.TypeUrl.ValueString(),
										Value:   []byte(tolerations.Value.ValueString()),
									}
								}
								return a
							}(),
						},
						Service: &v1alpha11.Operator_Service{Annotations: func() map[string]string {
							tmp := make(map[string]string)
							resp.Diagnostics.Append(model.InstallTemplate.Helm.Operator.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
							return tmp
						}()},
						ServiceAccount: &v1alpha11.Operator_ServiceAccount{
							Annotations: func() map[string]string {
								tmp := make(map[string]string)
								resp.Diagnostics.Append(model.InstallTemplate.Helm.Operator.ServiceAccount.Annotations.ElementsAs(ctx, &tmp, false)...)
								return tmp
							}(),
							ImagePullSecrets: func() []*anypb.Any {
								a := make([]*anypb.Any, len(model.InstallTemplate.Helm.Operator.ServiceAccount.ImagePullSecrets))
								for i, image_pull_secrets := range model.InstallTemplate.Helm.Operator.ServiceAccount.ImagePullSecrets {
									a[i] = &anypb.Any{
										TypeUrl: image_pull_secrets.TypeUrl.ValueString(),
										Value:   []byte(image_pull_secrets.Value.ValueString()),
									}
								}
								return a
							}(),
							PullPassword: model.InstallTemplate.Helm.Operator.ServiceAccount.PullPassword.ValueString(),
							PullSecret:   model.InstallTemplate.Helm.Operator.ServiceAccount.PullSecret.ValueString(),
							PullUsername: model.InstallTemplate.Helm.Operator.ServiceAccount.PullUsername.ValueString(),
						},
					},
					Secrets: &v1alpha1.Secrets{
						ClusterServiceAccount: &v1alpha1.Secrets_ClusterServiceAccount{
							Cluster_FQN: model.InstallTemplate.Helm.Secrets.ClusterServiceAccount.Cluster_FQN.ValueString(),
							Encoded_JWK: model.InstallTemplate.Helm.Secrets.ClusterServiceAccount.Encoded_JWK.ValueString(),
							JWK:         model.InstallTemplate.Helm.Secrets.ClusterServiceAccount.JWK.ValueString(),
						},
						Elasticsearch: &v1alpha1.Secrets_ElasticSearch{
							Cacert:   model.InstallTemplate.Helm.Secrets.Elasticsearch.Cacert.ValueString(),
							Password: model.InstallTemplate.Helm.Secrets.Elasticsearch.Password.ValueString(),
							Username: model.InstallTemplate.Helm.Secrets.Elasticsearch.Username.ValueString(),
						},
						Tsb: &v1alpha1.Secrets_TSB{Cacert: model.InstallTemplate.Helm.Secrets.Tsb.Cacert.ValueString()},
						Xcp: &v1alpha1.Secrets_XCP{
							AutoGenerateCerts: model.InstallTemplate.Helm.Secrets.Xcp.AutoGenerateCerts.ValueBool(),
							Edge: &v1alpha1.Secrets_XCP_Edge{
								Cert:  model.InstallTemplate.Helm.Secrets.Xcp.Edge.Cert.ValueString(),
								Key:   model.InstallTemplate.Helm.Secrets.Xcp.Edge.Key.ValueString(),
								Token: model.InstallTemplate.Helm.Secrets.Xcp.Edge.Token.ValueString(),
							},
							Rootca:    model.InstallTemplate.Helm.Secrets.Xcp.Rootca.ValueString(),
							Rootcakey: model.InstallTemplate.Helm.Secrets.Xcp.Rootcakey.ValueString(),
						},
					},
					Spec: &anypb.Any{
						TypeUrl: model.InstallTemplate.Helm.Spec.TypeUrl.ValueString(),
						Value:   []byte(model.InstallTemplate.Helm.Spec.Value.ValueString()),
					},
				},
				Message: model.InstallTemplate.Message.ValueString(),
			},
			Labels: func() map[string]string {
				tmp := make(map[string]string)
				resp.Diagnostics.Append(model.Labels.ElementsAs(ctx, &tmp, false)...)
				return tmp
			}(),
			Locality: &v2.Cluster_Locality{Region: model.Locality.Region.ValueString()},
			NamespaceScope: &v2.NamespaceScoping{
				Exceptions: func() []string {
					tmp := make([]string, len(model.NamespaceScope.Exceptions.Elements()))
					resp.Diagnostics.Append(model.NamespaceScope.Exceptions.ElementsAs(ctx, &tmp, false)...)
					return tmp
				}(),
				Scope: v2.Scope(v2.Scope_value[model.NamespaceScope.Scope.ValueString()]),
			},
			Namespaces: func() []*v2.Namespace {
				a := make([]*v2.Namespace, len(model.Namespaces))
				for i, namespaces := range model.Namespaces {
					a[i] = &v2.Namespace{
						Name: namespaces.Name.ValueString(),
						Services: func() []*v2.Service {
							a := make([]*v2.Service, len(namespaces.Services))
							for i, services := range namespaces.Services {
								a[i] = &v2.Service{
									CanonicalName: services.CanonicalName.ValueString(),
									GatewayHost:   services.GatewayHost.ValueBool(),
									Hostname:      services.Hostname.ValueString(),
									KubernetesExternalAddresses: func() []string {
										tmp := make([]string, len(services.KubernetesExternalAddresses.Elements()))
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
										a := make([]*v2.Service_Port, len(services.Ports))
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
										tmp := make([]string, len(services.SpiffeIds.Elements()))
										resp.Diagnostics.Append(services.SpiffeIds.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									Subsets: func() []string {
										tmp := make([]string, len(services.Subsets.Elements()))
										resp.Diagnostics.Append(services.Subsets.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									Tier1GatewayHost: services.Tier1GatewayHost.ValueBool(),
									Workloads: func() []*v2.Workload {
										a := make([]*v2.Workload, len(services.Workloads))
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
					a := make([]*v2.ServiceAccount_KeyPair, len(model.ServiceAccount.Keys))
					for i, keys := range model.ServiceAccount.Keys {
						a[i] = &v2.ServiceAccount_KeyPair{
							DefaultToken: keys.DefaultToken.ValueString(),
							Encoding:     v2.Encoding(v2.Encoding_value[keys.Encoding.ValueString()]),
							Id:           keys.Id.ValueString(),
							PrivateKey:   keys.PrivateKey.ValueString(),
							PublicKey:    keys.PublicKey.ValueString(),
						}
					}
					return a
				}(),
			},
			State: &v2.Cluster_State{
				IstioVersions: func() []string {
					tmp := make([]string, len(model.State.IstioVersions.Elements()))
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
