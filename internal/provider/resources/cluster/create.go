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
								Value:   model.InstallTemplate.Helm.Operator.Deployment.Affinity.Value.ValueString(),
							},
							Env: func() []*anypb.Any {
								a := make([]*anypb.Any, 0)
								for _, env := range model.InstallTemplate.Helm.Operator.Deployment.Env {
									a = append(a, &anypb.Any{
										TypeUrl: env.TypeUrl.ValueString(),
										Value:   env.Value.ValueString(),
									})
								}
								return a
							}(),
							Strategy: &anypb.Any{
								TypeUrl: model.InstallTemplate.Helm.Operator.Deployment.Strategy.TypeUrl.ValueString(),
								Value:   model.InstallTemplate.Helm.Operator.Deployment.Strategy.Value.ValueString(),
							},
							Tolerations: func() []*anypb.Any {
								a := make([]*anypb.Any, 0)
								for _, tolerations := range model.InstallTemplate.Helm.Operator.Deployment.Tolerations {
									a = append(a, &anypb.Any{
										TypeUrl: tolerations.TypeUrl.ValueString(),
										Value:   tolerations.Value.ValueString(),
									})
								}
								return a
							}(),
						},
						Service: &v1alpha11.Operator_Service{},
						ServiceAccount: &v1alpha11.Operator_ServiceAccount{
							ImagePullSecrets: func() []*anypb.Any {
								a := make([]*anypb.Any, 0)
								for _, image_pull_secrets := range model.InstallTemplate.Helm.Operator.ServiceAccount.ImagePullSecrets {
									a = append(a, &anypb.Any{
										TypeUrl: image_pull_secrets.TypeUrl.ValueString(),
										Value:   image_pull_secrets.Value.ValueString(),
									})
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
							ClusterFqn: model.InstallTemplate.Helm.Secrets.ClusterServiceAccount.ClusterFqn.ValueString(),
							EncodedJwk: model.InstallTemplate.Helm.Secrets.ClusterServiceAccount.EncodedJwk.ValueString(),
							Jwk:        model.InstallTemplate.Helm.Secrets.ClusterServiceAccount.Jwk.ValueString(),
						},
						Elasticsearch: &v1alpha1.Secrets_ElasticSearch{
							Cacert:   model.InstallTemplate.Helm.Secrets.Elasticsearch.Cacert.ValueString(),
							Password: model.InstallTemplate.Helm.Secrets.Elasticsearch.Password.ValueString(),
							Username: model.InstallTemplate.Helm.Secrets.Elasticsearch.Username.ValueString(),
						},
						Tsb: &v1alpha1.Secrets_TSB{Cacert: model.InstallTemplate.Helm.Secrets.Tsb.Cacert.ValueString()},
						Xcp: &v1alpha1.Secrets_XCP{
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
						Value:   model.InstallTemplate.Helm.Spec.Value.ValueString(),
					},
				},
				Message: model.InstallTemplate.Message.ValueString(),
			},
			Locality:       &v2.Cluster_Locality{Region: model.Locality.Region.ValueString()},
			NamespaceScope: &v2.NamespaceScoping{Scope: v2.Scope(v2.Scope_value[model.NamespaceScope.Scope.ValueString()])},
			Namespaces: func() []*v2.Namespace {
				a := make([]*v2.Namespace, 0)
				for _, namespaces := range model.Namespaces {
					a = append(a, &v2.Namespace{
						Name: namespaces.Name.ValueString(),
						Services: func() []*v2.Service {
							a := make([]*v2.Service, 0)
							for _, services := range namespaces.Services {
								a = append(a, &v2.Service{
									CanonicalName:         services.CanonicalName.ValueString(),
									Hostname:              services.Hostname.ValueString(),
									KubernetesServiceFqdn: services.KubernetesServiceFqdn.ValueString(),
									KubernetesServiceIp:   services.KubernetesServiceIp.ValueString(),
									Name:                  services.Name.ValueString(),
									Namespace:             services.Namespace.ValueString(),
									Ports: func() []*v2.Service_Port {
										a := make([]*v2.Service_Port, 0)
										for _, ports := range services.Ports {
											a = append(a, &v2.Service_Port{Name: ports.Name.ValueString()})
										}
										return a
									}(),
									Workloads: func() []*v2.Workload {
										a := make([]*v2.Workload, 0)
										for _, workloads := range services.Workloads {
											a = append(a, &v2.Workload{
												Address: workloads.Address.ValueString(),
												Name:    workloads.Name.ValueString(),
												Proxy: &v2.Workload_Proxy{
													ControlPlaneAddress: workloads.Proxy.ControlPlaneAddress.ValueString(),
													EnvoyVersion:        workloads.Proxy.EnvoyVersion.ValueString(),
													IstioVersion:        workloads.Proxy.IstioVersion.ValueString(),
												},
											})
										}
										return a
									}(),
								})
							}
							return a
						}(),
					})
				}
				return a
			}(),
			Network: model.Network.ValueString(),
			ServiceAccount: &v2.ServiceAccount{
				Description: model.ServiceAccount.Description.ValueString(),
				DisplayName: model.ServiceAccount.DisplayName.ValueString(),
				Keys: func() []*v2.ServiceAccount_KeyPair {
					a := make([]*v2.ServiceAccount_KeyPair, 0)
					for _, keys := range model.ServiceAccount.Keys {
						a = append(a, &v2.ServiceAccount_KeyPair{
							DefaultToken: keys.DefaultToken.ValueString(),
							Encoding:     v2.Encoding(v2.Encoding_value[keys.Encoding.ValueString()]),
							Id:           keys.Id.ValueString(),
							PrivateKey:   keys.PrivateKey.ValueString(),
							PublicKey:    keys.PublicKey.ValueString(),
						})
					}
					return a
				}(),
			},
			State: &v2.Cluster_State{
				LastSyncTime: &timestamppb.Timestamp{},
				Provider:     model.State.Provider.ValueString(),
				TsbCpVersion: model.State.TsbCpVersion.ValueString(),
				XcpVersion:   model.State.XcpVersion.ValueString(),
			},
			Tier1Cluster: &wrapperspb.BoolValue{},
			TokenTtl:     &durationpb.Duration{},
			TrustDomain:  model.TrustDomain.ValueString(),
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
