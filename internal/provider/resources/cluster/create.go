package cluster

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	v2 "github.com/tetrateio/api/tsb/v2"
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
			InstallTemplate: &v2.testPrefix_InstallTemplate{
				Helm: &v2.testPrefix_Helm{
					Image: &v2.testPrefix_Image{
						Registry: model.InstallTemplate.Helm.Image.Registry.ValueString(),
						Tag:      model.InstallTemplate.Helm.Image.Tag.ValueString(),
					},
					Operator: &v2.testPrefix_Operator{
						Deployment: &v2.testPrefix_Deployment{
							Affinity: &v2.testPrefix_Affinity{
								TypeUrl: model.InstallTemplate.Helm.Operator.Deployment.Affinity.TypeUrl.ValueString(),
								Value:   model.InstallTemplate.Helm.Operator.Deployment.Affinity.Value.ValueString(),
							},
							Env: &v2.testPrefix_Env{
								TypeUrl: model.InstallTemplate.Helm.Operator.Deployment.Env.TypeUrl.ValueString(),
								Value:   model.InstallTemplate.Helm.Operator.Deployment.Env.Value.ValueString(),
							},
							Strategy: &v2.testPrefix_Strategy{
								TypeUrl: model.InstallTemplate.Helm.Operator.Deployment.Strategy.TypeUrl.ValueString(),
								Value:   model.InstallTemplate.Helm.Operator.Deployment.Strategy.Value.ValueString(),
							},
							Tolerations: &v2.testPrefix_Tolerations{
								TypeUrl: model.InstallTemplate.Helm.Operator.Deployment.Tolerations.TypeUrl.ValueString(),
								Value:   model.InstallTemplate.Helm.Operator.Deployment.Tolerations.Value.ValueString(),
							},
						},
						Service: &v2.testPrefix_Service{},
						ServiceAccount: &v2.testPrefix_ServiceAccount{
							ImagePullSecrets: &v2.testPrefix_ImagePullSecrets{
								TypeUrl: model.InstallTemplate.Helm.Operator.ServiceAccount.ImagePullSecrets.TypeUrl.ValueString(),
								Value:   model.InstallTemplate.Helm.Operator.ServiceAccount.ImagePullSecrets.Value.ValueString(),
							},
							PullPassword: model.InstallTemplate.Helm.Operator.ServiceAccount.PullPassword.ValueString(),
							PullSecret:   model.InstallTemplate.Helm.Operator.ServiceAccount.PullSecret.ValueString(),
							PullUsername: model.InstallTemplate.Helm.Operator.ServiceAccount.PullUsername.ValueString(),
						},
					},
					Secrets: &v2.testPrefix_Secrets{
						ClusterServiceAccount: &v2.testPrefix_ClusterServiceAccount{
							ClusterFqn: model.InstallTemplate.Helm.Secrets.ClusterServiceAccount.ClusterFqn.ValueString(),
							EncodedJwk: model.InstallTemplate.Helm.Secrets.ClusterServiceAccount.EncodedJwk.ValueString(),
							Jwk:        model.InstallTemplate.Helm.Secrets.ClusterServiceAccount.Jwk.ValueString(),
						},
						Elasticsearch: &v2.testPrefix_Elasticsearch{
							Cacert:   model.InstallTemplate.Helm.Secrets.Elasticsearch.Cacert.ValueString(),
							Password: model.InstallTemplate.Helm.Secrets.Elasticsearch.Password.ValueString(),
							Username: model.InstallTemplate.Helm.Secrets.Elasticsearch.Username.ValueString(),
						},
						Tsb: &v2.testPrefix_Tsb{Cacert: model.InstallTemplate.Helm.Secrets.Tsb.Cacert.ValueString()},
						Xcp: &v2.testPrefix_Xcp{
							Edge: &v2.testPrefix_Edge{
								Cert:  model.InstallTemplate.Helm.Secrets.Xcp.Edge.Cert.ValueString(),
								Key:   model.InstallTemplate.Helm.Secrets.Xcp.Edge.Key.ValueString(),
								Token: model.InstallTemplate.Helm.Secrets.Xcp.Edge.Token.ValueString(),
							},
							Rootca:    model.InstallTemplate.Helm.Secrets.Xcp.Rootca.ValueString(),
							Rootcakey: model.InstallTemplate.Helm.Secrets.Xcp.Rootcakey.ValueString(),
						},
					},
					Spec: &v2.testPrefix_Spec{
						TypeUrl: model.InstallTemplate.Helm.Spec.TypeUrl.ValueString(),
						Value:   model.InstallTemplate.Helm.Spec.Value.ValueString(),
					},
				},
				Message: model.InstallTemplate.Message.ValueString(),
			},
			Locality:       &v2.testPrefix_Locality{Region: model.Locality.Region.ValueString()},
			NamespaceScope: &v2.testPrefix_NamespaceScope{Scope: v2.Scope(v2.Scope_value[model.NamespaceScope.Scope.ValueString()])},
			Namespaces: &v2.testPrefix_Namespaces{
				Name: model.Namespaces.Name.ValueString(),
				Services: &v2.testPrefix_Services{
					CanonicalName:         model.Namespaces.Services.CanonicalName.ValueString(),
					Hostname:              model.Namespaces.Services.Hostname.ValueString(),
					KubernetesServiceFqdn: model.Namespaces.Services.KubernetesServiceFqdn.ValueString(),
					KubernetesServiceIp:   model.Namespaces.Services.KubernetesServiceIp.ValueString(),
					Name:                  model.Namespaces.Services.Name.ValueString(),
					Namespace:             model.Namespaces.Services.Namespace.ValueString(),
					Ports:                 &v2.testPrefix_Ports{Name: model.Namespaces.Services.Ports.Name.ValueString()},
					Workloads: &v2.testPrefix_Workloads{
						Address: model.Namespaces.Services.Workloads.Address.ValueString(),
						Name:    model.Namespaces.Services.Workloads.Name.ValueString(),
						Proxy: &v2.testPrefix_Proxy{
							ControlPlaneAddress: model.Namespaces.Services.Workloads.Proxy.ControlPlaneAddress.ValueString(),
							EnvoyVersion:        model.Namespaces.Services.Workloads.Proxy.EnvoyVersion.ValueString(),
							IstioVersion:        model.Namespaces.Services.Workloads.Proxy.IstioVersion.ValueString(),
						},
					},
				},
			},
			Network: model.Network.ValueString(),
			ServiceAccount: &v2.testPrefix_ServiceAccount{
				Description: model.ServiceAccount.Description.ValueString(),
				DisplayName: model.ServiceAccount.DisplayName.ValueString(),
				Keys: &v2.testPrefix_Keys{
					DefaultToken: model.ServiceAccount.Keys.DefaultToken.ValueString(),
					Encoding:     v2.Encoding(v2.Encoding_value[model.ServiceAccount.Keys.Encoding.ValueString()]),
					Id:           model.ServiceAccount.Keys.Id.ValueString(),
					PrivateKey:   model.ServiceAccount.Keys.PrivateKey.ValueString(),
					PublicKey:    model.ServiceAccount.Keys.PublicKey.ValueString(),
				},
			},
			State: &v2.testPrefix_State{
				LastSyncTime: &v2.testPrefix_LastSyncTime{},
				Provider:     model.State.Provider.ValueString(),
				TsbCpVersion: model.State.TsbCpVersion.ValueString(),
				XcpVersion:   model.State.XcpVersion.ValueString(),
			},
			Tier1Cluster: &v2.testPrefix_Tier1Cluster{},
			TokenTtl:     &v2.testPrefix_TokenTtl{},
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
