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
	model.InstallTemplate = InstallTemplateModel{
		Helm: HelmInstallTemplateModel{
			Image: ImageHelmInstallTemplateModel{
				Registry: types.StringValue(cluster.Helm.Image.Registry[int32(cluster.Helm.Image.Registry)]),
				Tag:      types.StringValue(cluster.Helm.Image.Tag[int32(cluster.Helm.Image.Tag)]),
			},
			Operator: HelmOperatorDeploymentModel{
				Deployment: DeploymentOperatorHelmServiceAccountModel{
					Affinity: NodeAffinityAffinityDeploymentOperatorPodAntiAffinityModel{
						NodeAffinity: HelmOperatorDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
							PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
							RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityDeploymentOperatorHelmServiceAccountModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
						},
						PodAffinity: PodAffinityNodeAffinityAffinityDeploymentOperatorHelmModel{
							PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
							RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
						},
						PodAntiAffinity: PodAntiAffinityOperatorDeploymentAffinityNodeAffinityPodAffinityModel{
							PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
							RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
						},
					},
					Annotations: func() basetypes.MapValue {
						r, diag := types.MapValue(ctx, AnnotationsDeploymentOperatorHelmServiceAccountModel.ElementType(ctx), cluster.Helm.Operator.Deployment.Annotations)
						resp.Diagnostics.Append(diag...)
						return r
					}(),
					Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
					PodAnnotations: func() basetypes.MapValue {
						r, diag := types.MapValue(ctx, PodAnnotationsDeploymentOperatorHelmServiceAccountModel.ElementType(ctx), cluster.Helm.Operator.Deployment.PodAnnotations)
						resp.Diagnostics.Append(diag...)
						return r
					}(),
					ReplicaCount: types.Int64Value(cluster.Helm.Operator.Deployment.ReplicaCount),
					Strategy: HelmOperatorDeploymentStrategyRollingUpdateModel{
						RollingUpdate: RollingUpdateStrategyDeploymentOperatorHelmMaxUnavailableModel{
							MaxSurge: MaxSurgeRollingUpdateStrategyDeploymentOperatorHelmStrValModel{
								IntVal: IntValMaxSurgeRollingUpdateStrategyDeploymentOperatorHelmServiceAccountModel{Value: types.Int64Value(cluster.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
								StrVal: StrValHelmOperatorDeploymentStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
								Type:   types.Int64Value(cluster.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
							},
							MaxUnavailable: MaxUnavailableHelmOperatorDeploymentStrategyRollingUpdateStrValModel{
								IntVal: IntValMaxUnavailableHelmOperatorDeploymentStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
								StrVal: StrValRollingUpdateStrategyDeploymentOperatorHelmMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
								Type:   types.Int64Value(cluster.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
							},
						},
						Type: types.StringValue(cluster.Helm.Operator.Deployment.Strategy.Type[int32(cluster.Helm.Operator.Deployment.Strategy.Type)]),
					},
					Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
				},
				Service: ServiceOperatorHelmInstallTemplateModel{Annotations: func() basetypes.MapValue {
					r, diag := types.MapValue(ctx, AnnotationsServiceOperatorHelmInstallTemplateModel.ElementType(ctx), cluster.Helm.Operator.Service.Annotations)
					resp.Diagnostics.Append(diag...)
					return r
				}()},
				ServiceAccount: ServiceAccountHelmOperatorServiceModel{
					Annotations: func() basetypes.MapValue {
						r, diag := types.MapValue(ctx, AnnotationsServiceAccountHelmOperatorServiceModel.ElementType(ctx), cluster.Helm.Operator.ServiceAccount.Annotations)
						resp.Diagnostics.Append(diag...)
						return r
					}(),
					ImagePullSecrets: func() basetypes.ListValue {
						r, diag := types.ListValue(ctx, ImagePullSecretsServiceAccountHelmOperatorServiceModel.ElementType(ctx), cluster.Helm.Operator.ServiceAccount.ImagePullSecrets)
						resp.Diagnostics.Append(diag...)
					}(),
					PullPassword: types.StringValue(cluster.Helm.Operator.ServiceAccount.PullPassword[int32(cluster.Helm.Operator.ServiceAccount.PullPassword)]),
					PullSecret:   types.StringValue(cluster.Helm.Operator.ServiceAccount.PullSecret[int32(cluster.Helm.Operator.ServiceAccount.PullSecret)]),
					PullUsername: types.StringValue(cluster.Helm.Operator.ServiceAccount.PullUsername[int32(cluster.Helm.Operator.ServiceAccount.PullUsername)]),
				},
			},
			Secrets: SecretsHelmClusterServiceAccountModel{
				ClusterServiceAccount: ClusterServiceAccountHelmSecretsXcpModel{
					Cluster_FQN: types.StringValue(cluster.Helm.Secrets.ClusterServiceAccount.Cluster_FQN[int32(cluster.Helm.Secrets.ClusterServiceAccount.Cluster_FQN)]),
					Encoded_JWK: types.StringValue(cluster.Helm.Secrets.ClusterServiceAccount.Encoded_JWK[int32(cluster.Helm.Secrets.ClusterServiceAccount.Encoded_JWK)]),
					JWK:         types.StringValue(cluster.Helm.Secrets.ClusterServiceAccount.JWK[int32(cluster.Helm.Secrets.ClusterServiceAccount.JWK)]),
				},
				Elasticsearch: ElasticsearchSecretsHelmInstallTemplateModel{
					Cacert:   types.StringValue(cluster.Helm.Secrets.Elasticsearch.Cacert[int32(cluster.Helm.Secrets.Elasticsearch.Cacert)]),
					Password: types.StringValue(cluster.Helm.Secrets.Elasticsearch.Password[int32(cluster.Helm.Secrets.Elasticsearch.Password)]),
					Username: types.StringValue(cluster.Helm.Secrets.Elasticsearch.Username[int32(cluster.Helm.Secrets.Elasticsearch.Username)]),
				},
				Tsb: TsbHelmSecretsElasticsearchModel{Cacert: types.StringValue(cluster.Helm.Secrets.Tsb.Cacert[int32(cluster.Helm.Secrets.Tsb.Cacert)])},
				Xcp: XcpSecretsHelmTsbModel{
					AutoGenerateCerts: "primitive / BoolKind",
					Edge: EdgeXcpSecretsHelmTsbModel{
						Cert:  types.StringValue(cluster.Helm.Secrets.Xcp.Edge.Cert[int32(cluster.Helm.Secrets.Xcp.Edge.Cert)]),
						Key:   types.StringValue(cluster.Helm.Secrets.Xcp.Edge.Key[int32(cluster.Helm.Secrets.Xcp.Edge.Key)]),
						Token: types.StringValue(cluster.Helm.Secrets.Xcp.Edge.Token[int32(cluster.Helm.Secrets.Xcp.Edge.Token)]),
					},
					Rootca:    types.StringValue(cluster.Helm.Secrets.Xcp.Rootca[int32(cluster.Helm.Secrets.Xcp.Rootca)]),
					Rootcakey: types.StringValue(cluster.Helm.Secrets.Xcp.Rootcakey[int32(cluster.Helm.Secrets.Xcp.Rootcakey)]),
				},
			},
			Spec: HelmSpecTelemetryStoreModel{
				Components: ComponentsSpecHelmInstallTemplateModel{
					Collector: SpecComponentsCollectorKubeSpecDeploymentModel{KubeSpec: DeploymentKubeSpecCollectorComponentsSpecHelmModel{
						Deployment: HelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel{
							Affinity: AffinityDeploymentKubeSpecCollectorComponentsSpecHelmStrategyModel{
								NodeAffinity: HelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
									PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
									RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityDeploymentKubeSpecCollectorComponentsSpecHelmStrategyModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
								},
								PodAffinity: PodAffinityAffinityDeploymentKubeSpecCollectorComponentsSpecHelmStrategyModel{
									PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
									RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
								},
								PodAntiAffinity: PodAntiAffinityAffinityDeploymentKubeSpecCollectorComponentsSpecHelmStrategyModel{
									PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
									RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
								},
							},
							ContainerSecurityContext: ContainerSecurityContextDeploymentKubeSpecCollectorComponentsSpecHelmPodSecurityContextModel{
								AllowPrivilegeEscalation: "primitive* / BoolKind",
								Capabilities: CapabilitiesContainerSecurityContextDeploymentKubeSpecCollectorComponentsSpecHelmDropModel{
									Add: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, AddCapabilitiesContainerSecurityContextDeploymentKubeSpecCollectorComponentsSpecHelmPodSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
										resp.Diagnostics.Append(diag...)
									}(),
									Drop: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, DropHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
										resp.Diagnostics.Append(diag...)
									}(),
								},
								Privileged:             "primitive* / BoolKind",
								ProcMount:              types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
								ReadOnlyRootFilesystem: "primitive* / BoolKind",
								RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
								RunAsNonRoot:           "primitive* / BoolKind",
								RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
								SeLinuxOptions: SeLinuxOptionsContainerSecurityContextDeploymentKubeSpecCollectorComponentsSpecHelmPodSecurityContextModel{
									Level: types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
									Role:  types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
									Type:  types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
									User:  types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
								},
								SeccompProfile: SeccompProfileContainerSecurityContextDeploymentKubeSpecCollectorComponentsSpecHelmPodSecurityContextModel{
									LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
									Type:             types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
								},
								WindowsOptions: WindowsOptionsContainerSecurityContextDeploymentKubeSpecCollectorComponentsSpecHelmPodSecurityContextModel{
									GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
									GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
									RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
								},
							},
							Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
							HpaSpec: HpaSpecHelmSpecComponentsCollectorKubeSpecDeploymentAffinityModel{
								MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.HpaSpec.MaxReplicas),
								Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
								MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.HpaSpec.MinReplicas),
							},
							PodAnnotations: func() basetypes.MapValue {
								r, diag := types.MapValue(ctx, PodAnnotationsDeploymentKubeSpecCollectorComponentsSpecHelmHpaSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodAnnotations)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							PodSecurityContext: PodSecurityContextHelmSpecComponentsCollectorKubeSpecDeploymentPodAnnotationsModel{
								FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.FsGroup),
								FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
								RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
								RunAsNonRoot:        "primitive* / BoolKind",
								RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
								SeLinuxOptions: SeLinuxOptionsPodSecurityContextHelmSpecComponentsCollectorKubeSpecDeploymentPodAnnotationsModel{
									Level: types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
									Role:  types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
									Type:  types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
									User:  types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
								},
								SeccompProfile: SeccompProfilePodSecurityContextHelmSpecComponentsCollectorKubeSpecDeploymentPodAnnotationsModel{
									LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
									Type:             types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
								},
								SupplementalGroups: func() basetypes.ListValue {
									r, diag := types.ListValue(ctx, SupplementalGroupsPodSecurityContextHelmSpecComponentsCollectorKubeSpecDeploymentPodAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
									resp.Diagnostics.Append(diag...)
								}(),
								Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
								WindowsOptions: WindowsOptionsPodSecurityContextHelmSpecComponentsCollectorKubeSpecDeploymentPodAnnotationsModel{
									GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
									GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
									RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
								},
							},
							ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.ReplicaCount),
							Resources: ResourcesDeploymentKubeSpecCollectorComponentsSpecHelmServiceAccountModel{
								Limits: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, LimitsResourcesDeploymentKubeSpecCollectorComponentsSpecHelmServiceAccountModel.ElementType(ctx), cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.Resources.Limits)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Requests: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, RequestsResourcesDeploymentKubeSpecCollectorComponentsSpecHelmServiceAccountModel.ElementType(ctx), cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.Resources.Requests)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
							},
							Strategy: StrategyHelmSpecComponentsCollectorKubeSpecDeploymentResourcesModel{
								RollingUpdate: RollingUpdateStrategyHelmSpecComponentsCollectorKubeSpecDeploymentMaxUnavailableModel{
									MaxSurge: MaxSurgeRollingUpdateStrategyHelmSpecComponentsCollectorKubeSpecDeploymentStrValModel{
										IntVal: IntValMaxSurgeRollingUpdateStrategyHelmSpecComponentsCollectorKubeSpecDeploymentResourcesModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
										StrVal: StrValDeploymentKubeSpecCollectorComponentsSpecHelmStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
										Type:   types.Int64Value(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
									},
									MaxUnavailable: MaxUnavailableDeploymentKubeSpecCollectorComponentsSpecHelmStrategyRollingUpdateIntValModel{
										IntVal: IntValRollingUpdateStrategyHelmSpecComponentsCollectorKubeSpecDeploymentMaxUnavailableStrValModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
										StrVal: StrValMaxUnavailableDeploymentKubeSpecCollectorComponentsSpecHelmStrategyRollingUpdateMaxSurgeModel{Value: types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
										Type:   types.Int64Value(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
									},
								},
								Type: types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.Type)]),
							},
							Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
						},
						Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
						Service: ServiceKubeSpecCollectorComponentsSpecHelmLabelsModel{
							Annotations: func() basetypes.MapValue {
								r, diag := types.MapValue(ctx, AnnotationsServiceKubeSpecCollectorComponentsSpecHelmInstallTemplateModel.ElementType(ctx), cluster.Helm.Spec.Components.Collector.KubeSpec.Service.Annotations)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							Labels: func() basetypes.MapValue {
								r, diag := types.MapValue(ctx, LabelsHelmSpecComponentsCollectorKubeSpecServiceAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.Collector.KubeSpec.Service.Labels)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
							Type:  types.StringValue(cluster.Helm.Spec.Components.Collector.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.Collector.KubeSpec.Service.Type)]),
						},
						ServiceAccount: ServiceAccountHelmSpecComponentsCollectorKubeSpecServiceModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
					}},
					DefaultKubeSpec: ServiceDeploymentHelmSpecComponentsModel{
						Account: AccountDefaultKubeSpecComponentsSpecHelmInstallTemplateModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
						Deployment: StrategyComponentsSpecHelmDeploymentAffinityModel{
							Affinity: DefaultKubeSpecComponentsSpecHelmDeploymentAffinityPodAntiAffinityModel{
								NodeAffinity: NodeAffinityAffinityDeploymentHelmSpecComponentsDefaultKubeSpecAccountModel{
									PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
									RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityDeploymentHelmSpecComponentsDefaultKubeSpecAccountModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
								},
								PodAffinity: PodAffinityDefaultKubeSpecComponentsSpecHelmDeploymentAffinityNodeAffinityModel{
									PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
									RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
								},
								PodAntiAffinity: PodAntiAffinityAffinityDeploymentHelmSpecComponentsDefaultKubeSpecPodAffinityModel{
									PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
									RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
								},
							},
							ContainerSecurityContext: ContainerSecurityContextAffinityDeploymentHelmSpecComponentsWindowsOptionsModel{
								AllowPrivilegeEscalation: "primitive* / BoolKind",
								Capabilities: CapabilitiesComponentsSpecHelmDeploymentAffinityContainerSecurityContextSeLinuxOptionsModel{
									Add: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, AddCapabilitiesComponentsSpecHelmDeploymentAffinityContainerSecurityContextSeLinuxOptionsModel.ElementType(ctx), cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
										resp.Diagnostics.Append(diag...)
									}(),
									Drop: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, DropCapabilitiesComponentsSpecHelmDeploymentAffinityContainerSecurityContextSeLinuxOptionsModel.ElementType(ctx), cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
										resp.Diagnostics.Append(diag...)
									}(),
								},
								Privileged:             "primitive* / BoolKind",
								ProcMount:              types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
								ReadOnlyRootFilesystem: "primitive* / BoolKind",
								RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
								RunAsNonRoot:           "primitive* / BoolKind",
								RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
								SeLinuxOptions: SeLinuxOptionsContainerSecurityContextAffinityDeploymentHelmSpecComponentsDefaultKubeSpecModel{
									Level: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
									Role:  types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
									Type:  types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
									User:  types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
								},
								SeccompProfile: SeccompProfileContainerSecurityContextAffinityDeploymentHelmSpecComponentsCapabilitiesModel{
									LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
									Type:             types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
								},
								WindowsOptions: WindowsOptionsComponentsSpecHelmDeploymentAffinityContainerSecurityContextSeccompProfileModel{
									GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
									GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
									RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
								},
							},
							Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
							PodAnnotations: func() basetypes.MapValue {
								r, diag := types.MapValue(ctx, PodAnnotationsComponentsSpecHelmDeploymentAffinityContainerSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodAnnotations)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							PodSecurityContext: PodSecurityContextAffinityDeploymentHelmSpecComponentsWindowsOptionsModel{
								FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.FsGroup),
								FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
								RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.RunAsGroup),
								RunAsNonRoot:        "primitive* / BoolKind",
								RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.RunAsUser),
								SeLinuxOptions: SeLinuxOptionsPodSecurityContextAffinityDeploymentHelmSpecComponentsSupplementalGroupsModel{
									Level: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
									Role:  types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
									Type:  types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
									User:  types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
								},
								SeccompProfile: SeccompProfilePodSecurityContextAffinityDeploymentHelmSpecComponentsPodAnnotationsModel{
									LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
									Type:             types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
								},
								SupplementalGroups: func() basetypes.ListValue {
									r, diag := types.ListValue(ctx, SupplementalGroupsComponentsSpecHelmDeploymentAffinityPodSecurityContextSeccompProfileModel.ElementType(ctx), cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
									resp.Diagnostics.Append(diag...)
								}(),
								Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
								WindowsOptions: WindowsOptionsComponentsSpecHelmDeploymentAffinityPodSecurityContextSeLinuxOptionsModel{
									GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
									GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
									RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
								},
							},
							Strategy: AffinityDeploymentHelmSpecComponentsStrategyRollingUpdateModel{
								RollingUpdate: RollingUpdateStrategyComponentsSpecHelmDeploymentAffinityPodSecurityContextModel{
									MaxSurge: MaxSurgeRollingUpdateStrategyComponentsSpecHelmDeploymentAffinityStrValModel{
										IntVal: IntValMaxSurgeRollingUpdateStrategyComponentsSpecHelmDeploymentAffinityPodSecurityContextModel{Value: types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
										StrVal: StrValAffinityDeploymentHelmSpecComponentsStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
										Type:   types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
									},
									MaxUnavailable: MaxUnavailableRollingUpdateStrategyComponentsSpecHelmDeploymentAffinityIntValModel{
										IntVal: IntValAffinityDeploymentHelmSpecComponentsStrategyRollingUpdateMaxUnavailableStrValModel{Value: types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
										StrVal: StrValMaxUnavailableRollingUpdateStrategyComponentsSpecHelmDeploymentAffinityPodSecurityContextModel{Value: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
										Type:   types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
									},
								},
								Type: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.Type)]),
							},
							Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
						},
						Job: AffinityComponentsSpecHelmDeploymentJobModel{
							Affinity: JobDeploymentHelmSpecComponentsAffinityPodAntiAffinityModel{
								NodeAffinity: NodeAffinityAffinityComponentsSpecHelmDeploymentJobPodSecurityContextModel{
									PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
									RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityComponentsSpecHelmDeploymentJobPodSecurityContextModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
								},
								PodAffinity: PodAffinityJobDeploymentHelmSpecComponentsAffinityNodeAffinityModel{
									PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
									RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
								},
								PodAntiAffinity: PodAntiAffinityAffinityComponentsSpecHelmDeploymentJobPodAffinityModel{
									PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
									RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
								},
							},
							ContainerSecurityContext: ContainerSecurityContextJobDeploymentHelmSpecComponentsWindowsOptionsModel{
								AllowPrivilegeEscalation: "primitive* / BoolKind",
								Capabilities: CapabilitiesContainerSecurityContextJobDeploymentHelmSpecComponentsStrategyModel{
									Add: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, AddCapabilitiesContainerSecurityContextJobDeploymentHelmSpecComponentsStrategyModel.ElementType(ctx), cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.Capabilities.Add)
										resp.Diagnostics.Append(diag...)
									}(),
									Drop: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, DropCapabilitiesContainerSecurityContextJobDeploymentHelmSpecComponentsStrategyModel.ElementType(ctx), cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.Capabilities.Drop)
										resp.Diagnostics.Append(diag...)
									}(),
								},
								Privileged:             "primitive* / BoolKind",
								ProcMount:              types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.ProcMount)]),
								ReadOnlyRootFilesystem: "primitive* / BoolKind",
								RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.RunAsGroup),
								RunAsNonRoot:           "primitive* / BoolKind",
								RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.RunAsUser),
								SeLinuxOptions: SeLinuxOptionsComponentsSpecHelmDeploymentJobContainerSecurityContextCapabilitiesModel{
									Level: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Level)]),
									Role:  types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Role)]),
									Type:  types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Type)]),
									User:  types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.User)]),
								},
								SeccompProfile: SeccompProfileContainerSecurityContextJobDeploymentHelmSpecComponentsSeLinuxOptionsModel{
									LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
									Type:             types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeccompProfile.Type)]),
								},
								WindowsOptions: WindowsOptionsComponentsSpecHelmDeploymentJobContainerSecurityContextSeccompProfileModel{
									GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
									GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
									RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
								},
							},
							PodAnnotations: func() basetypes.MapValue {
								r, diag := types.MapValue(ctx, PodAnnotationsComponentsSpecHelmDeploymentJobContainerSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodAnnotations)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							PodSecurityContext: PodSecurityContextJobDeploymentHelmSpecComponentsSeLinuxOptionsModel{
								FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.FsGroup),
								FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.FsGroupChangePolicy)]),
								RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.RunAsGroup),
								RunAsNonRoot:        "primitive* / BoolKind",
								RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.RunAsUser),
								SeLinuxOptions: SeLinuxOptionsComponentsSpecHelmDeploymentJobPodSecurityContextSupplementalGroupsModel{
									Level: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeLinuxOptions.Level)]),
									Role:  types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeLinuxOptions.Role)]),
									Type:  types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeLinuxOptions.Type)]),
									User:  types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeLinuxOptions.User)]),
								},
								SeccompProfile: SeccompProfileComponentsSpecHelmDeploymentJobPodSecurityContextWindowsOptionsModel{
									LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
									Type:             types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeccompProfile.Type)]),
								},
								SupplementalGroups: func() basetypes.ListValue {
									r, diag := types.ListValue(ctx, SupplementalGroupsPodSecurityContextJobDeploymentHelmSpecComponentsSeccompProfileModel.ElementType(ctx), cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SupplementalGroups)
									resp.Diagnostics.Append(diag...)
								}(),
								Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
								WindowsOptions: WindowsOptionsPodSecurityContextJobDeploymentHelmSpecComponentsPodAnnotationsModel{
									GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
									GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
									RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.WindowsOptions.RunAsUserName)]),
								},
							},
							Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
						},
						Service: ComponentsSpecHelmDeploymentServiceAnnotationsModel{Annotations: func() basetypes.MapValue {
							r, diag := types.MapValue(ctx, AnnotationsServiceDeploymentHelmSpecComponentsAffinityModel.ElementType(ctx), cluster.Helm.Spec.Components.DefaultKubeSpec.Service.Annotations)
							resp.Diagnostics.Append(diag...)
							return r
						}()},
					},
					Gitops: HelmSpecComponentsGitopsReconcileIntervalModel{
						BatchWindow: BatchWindowGitopsComponentsSpecHelmWebhookTimeoutModel{
							Nanos:   types.Int64Value(cluster.Helm.Spec.Components.Gitops.BatchWindow.Nanos),
							Seconds: types.Int64Value(cluster.Helm.Spec.Components.Gitops.BatchWindow.Seconds),
						},
						Enabled: "primitive / BoolKind",
						ManagementplaneRequestTimeout: ManagementplaneRequestTimeoutHelmSpecComponentsGitopsBatchWindowModel{
							Nanos:   types.Int64Value(cluster.Helm.Spec.Components.Gitops.ManagementplaneRequestTimeout.Nanos),
							Seconds: types.Int64Value(cluster.Helm.Spec.Components.Gitops.ManagementplaneRequestTimeout.Seconds),
						},
						ReconcileInterval: ReconcileIntervalGitopsComponentsSpecHelmManagementplaneRequestTimeoutModel{
							Nanos:   types.Int64Value(cluster.Helm.Spec.Components.Gitops.ReconcileInterval.Nanos),
							Seconds: types.Int64Value(cluster.Helm.Spec.Components.Gitops.ReconcileInterval.Seconds),
						},
						ReconcileRequestTimeout: ReconcileRequestTimeoutGitopsComponentsSpecHelmInstallTemplateModel{
							Nanos:   types.Int64Value(cluster.Helm.Spec.Components.Gitops.ReconcileRequestTimeout.Nanos),
							Seconds: types.Int64Value(cluster.Helm.Spec.Components.Gitops.ReconcileRequestTimeout.Seconds),
						},
						WebhookTimeout: WebhookTimeoutHelmSpecComponentsGitopsReconcileRequestTimeoutModel{
							Nanos:   types.Int64Value(cluster.Helm.Spec.Components.Gitops.WebhookTimeout.Nanos),
							Seconds: types.Int64Value(cluster.Helm.Spec.Components.Gitops.WebhookTimeout.Seconds),
						},
					},
					HpaAdapter: SpecComponentsHpaAdapterKubeSpecDeploymentModel{KubeSpec: DeploymentKubeSpecHpaAdapterComponentsSpecServiceAccountModel{
						Deployment: HelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel{
							Affinity: AffinityHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyModel{
								NodeAffinity: DeploymentKubeSpecHpaAdapterComponentsSpecHelmAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
									PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
									RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
								},
								PodAffinity: PodAffinityAffinityHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyModel{
									PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
									RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
								},
								PodAntiAffinity: PodAntiAffinityAffinityHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyModel{
									PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
									RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
								},
							},
							ContainerSecurityContext: ContainerSecurityContextDeploymentKubeSpecHpaAdapterComponentsSpecHelmInstallTemplateModel{
								AllowPrivilegeEscalation: "primitive* / BoolKind",
								Capabilities: CapabilitiesContainerSecurityContextDeploymentKubeSpecHpaAdapterComponentsSpecHelmDropModel{
									Add: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, AddCapabilitiesContainerSecurityContextDeploymentKubeSpecHpaAdapterComponentsSpecHelmInstallTemplateModel.ElementType(ctx), cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
										resp.Diagnostics.Append(diag...)
									}(),
									Drop: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, DropHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
										resp.Diagnostics.Append(diag...)
									}(),
								},
								Privileged:             "primitive* / BoolKind",
								ProcMount:              types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
								ReadOnlyRootFilesystem: "primitive* / BoolKind",
								RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
								RunAsNonRoot:           "primitive* / BoolKind",
								RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
								SeLinuxOptions: SeLinuxOptionsContainerSecurityContextDeploymentKubeSpecHpaAdapterComponentsSpecHelmInstallTemplateModel{
									Level: types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
									Role:  types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
									Type:  types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
									User:  types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
								},
								SeccompProfile: SeccompProfileContainerSecurityContextDeploymentKubeSpecHpaAdapterComponentsSpecHelmInstallTemplateModel{
									LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
									Type:             types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
								},
								WindowsOptions: WindowsOptionsContainerSecurityContextDeploymentKubeSpecHpaAdapterComponentsSpecHelmInstallTemplateModel{
									GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
									GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
									RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
								},
							},
							Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
							HpaSpec: HpaSpecDeploymentKubeSpecHpaAdapterComponentsSpecHelmAffinityModel{
								MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.HpaSpec.MaxReplicas),
								Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
								MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.HpaSpec.MinReplicas),
							},
							PodAnnotations: func() basetypes.MapValue {
								r, diag := types.MapValue(ctx, PodAnnotationsHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodAnnotations)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							PodSecurityContext: PodSecurityContextDeploymentKubeSpecHpaAdapterComponentsSpecHelmPodAnnotationsModel{
								FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.FsGroup),
								FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
								RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
								RunAsNonRoot:        "primitive* / BoolKind",
								RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
								SeLinuxOptions: SeLinuxOptionsPodSecurityContextDeploymentKubeSpecHpaAdapterComponentsSpecHelmPodAnnotationsModel{
									Level: types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
									Role:  types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
									Type:  types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
									User:  types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
								},
								SeccompProfile: SeccompProfilePodSecurityContextDeploymentKubeSpecHpaAdapterComponentsSpecHelmPodAnnotationsModel{
									LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
									Type:             types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
								},
								SupplementalGroups: func() basetypes.ListValue {
									r, diag := types.ListValue(ctx, SupplementalGroupsPodSecurityContextDeploymentKubeSpecHpaAdapterComponentsSpecHelmPodAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
									resp.Diagnostics.Append(diag...)
								}(),
								Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
								WindowsOptions: WindowsOptionsPodSecurityContextDeploymentKubeSpecHpaAdapterComponentsSpecHelmPodAnnotationsModel{
									GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
									GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
									RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
								},
							},
							ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ReplicaCount),
							Resources: ResourcesHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel{
								Limits: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, LimitsResourcesHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Resources.Limits)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Requests: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, RequestsResourcesHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Resources.Requests)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
							},
							Strategy: StrategyDeploymentKubeSpecHpaAdapterComponentsSpecHelmResourcesModel{
								RollingUpdate: RollingUpdateStrategyDeploymentKubeSpecHpaAdapterComponentsSpecHelmMaxUnavailableModel{
									MaxSurge: MaxSurgeRollingUpdateStrategyDeploymentKubeSpecHpaAdapterComponentsSpecHelmStrValModel{
										IntVal: IntValMaxSurgeRollingUpdateStrategyDeploymentKubeSpecHpaAdapterComponentsSpecHelmResourcesModel{Value: types.Int64Value(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
										StrVal: StrValHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
										Type:   types.Int64Value(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
									},
									MaxUnavailable: MaxUnavailableHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateStrValModel{
										IntVal: IntValMaxUnavailableHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
										StrVal: StrValRollingUpdateStrategyDeploymentKubeSpecHpaAdapterComponentsSpecHelmMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
										Type:   types.Int64Value(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
									},
								},
								Type: types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.Type)]),
							},
							Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
						},
						Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
						Service: ServiceDeploymentKubeSpecHpaAdapterComponentsSpecLabelsModel{
							Annotations: func() basetypes.MapValue {
								r, diag := types.MapValue(ctx, AnnotationsServiceDeploymentKubeSpecHpaAdapterComponentsSpecHelmModel.ElementType(ctx), cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Service.Annotations)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							Labels: func() basetypes.MapValue {
								r, diag := types.MapValue(ctx, LabelsSpecComponentsHpaAdapterKubeSpecDeploymentServiceAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Service.Labels)
								resp.Diagnostics.Append(diag...)
								return r
							}(),
							Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
							Type:  types.StringValue(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.HpaAdapter.KubeSpec.Service.Type)]),
						},
						ServiceAccount: ServiceAccountSpecComponentsHpaAdapterKubeSpecDeploymentServiceModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
					}},
					InternalCertProvider: CertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerCustomModel{
						CertManager: CertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderModel{
							CertManagerCaInjector: HelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel{KubeSpec: KubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmInstallTemplateModel{
								Deployment: CertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmAffinityStrategyResourcesModel{
									Affinity: NodeAffinityAffinityHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel{
										NodeAffinity: DeploymentKubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodAntiAffinityModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
										},
										PodAffinity: PodAffinityAffinityHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
										},
										PodAntiAffinity: PodAntiAffinityDeploymentKubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmAffinityPodAffinityModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
										},
									},
									ContainerSecurityContext: ContainerSecurityContextStrategyAffinityHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorSeccompProfileModel{
										AllowPrivilegeEscalation: "primitive* / BoolKind",
										Capabilities: CapabilitiesCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmAffinityStrategyContainerSecurityContextDropModel{
											Add: func() basetypes.ListValue {
												r, diag := types.ListValue(ctx, AddCapabilitiesCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmAffinityStrategyContainerSecurityContextWindowsOptionsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
												resp.Diagnostics.Append(diag...)
											}(),
											Drop: func() basetypes.ListValue {
												r, diag := types.ListValue(ctx, DropContainerSecurityContextStrategyAffinityHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
												resp.Diagnostics.Append(diag...)
											}(),
										},
										Privileged:             "primitive* / BoolKind",
										ProcMount:              types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
										ReadOnlyRootFilesystem: "primitive* / BoolKind",
										RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
										RunAsNonRoot:           "primitive* / BoolKind",
										RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
										SeLinuxOptions: SeLinuxOptionsContainerSecurityContextStrategyAffinityHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCapabilitiesModel{
											Level: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
											Role:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
											Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
											User:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
										},
										SeccompProfile: SeccompProfileCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmAffinityStrategyContainerSecurityContextSeLinuxOptionsModel{
											LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
											Type:             types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
										},
										WindowsOptions: WindowsOptionsContainerSecurityContextStrategyAffinityHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel{
											GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
											GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
											RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
										},
									},
									Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
									HpaSpec: HpaSpecKubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmAffinityNodeAffinityModel{
										MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.HpaSpec.MaxReplicas),
										Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
										MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.HpaSpec.MinReplicas),
									},
									PodAnnotations: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, PodAnnotationsCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmAffinityStrategyContainerSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodAnnotations)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									PodSecurityContext: PodSecurityContextDeploymentKubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmWindowsOptionsModel{
										FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.FsGroup),
										FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
										RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
										RunAsNonRoot:        "primitive* / BoolKind",
										RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
										SeLinuxOptions: SeLinuxOptionsHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextSupplementalGroupsModel{
											Level: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
											Role:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
											Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
											User:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
										},
										SeccompProfile: SeccompProfilePodSecurityContextDeploymentKubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmSeLinuxOptionsModel{
											LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
											Type:             types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
										},
										SupplementalGroups: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, SupplementalGroupsPodSecurityContextDeploymentKubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmInstallTemplateModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
											resp.Diagnostics.Append(diag...)
										}(),
										Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
										WindowsOptions: WindowsOptionsHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextSeccompProfileModel{
											GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
											GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
											RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
										},
									},
									ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ReplicaCount),
									Resources: ResourcesStrategyAffinityHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorRequestsModel{
										Limits: func() basetypes.MapValue {
											r, diag := types.MapValue(ctx, LimitsResourcesStrategyAffinityHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorPodAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Resources.Limits)
											resp.Diagnostics.Append(diag...)
											return r
										}(),
										Requests: func() basetypes.MapValue {
											r, diag := types.MapValue(ctx, RequestsCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmAffinityStrategyResourcesLimitsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Resources.Requests)
											resp.Diagnostics.Append(diag...)
											return r
										}(),
									},
									Strategy: KubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmAffinityStrategyRollingUpdateModel{
										RollingUpdate: RollingUpdateStrategyAffinityHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecMaxUnavailableModel{
											MaxSurge: MaxSurgeRollingUpdateStrategyAffinityHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecStrValModel{
												IntVal: IntValMaxSurgeRollingUpdateStrategyAffinityHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecHpaSpecModel{Value: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
												StrVal: StrValKubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmAffinityStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
												Type:   types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
											},
											MaxUnavailable: MaxUnavailableKubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmAffinityStrategyRollingUpdateStrValModel{
												IntVal: IntValMaxUnavailableKubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmAffinityStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
												StrVal: StrValRollingUpdateStrategyAffinityHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
												Type:   types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
											},
										},
										Type: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.Type)]),
									},
									Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
								},
								Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
								Service: ServiceKubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmLabelsModel{
									Annotations: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, AnnotationsServiceKubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmInstallTemplateModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Service.Annotations)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Labels: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, LabelsHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServiceAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Service.Labels)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
									Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Service.Type)]),
								},
								ServiceAccount: ServiceAccountKubeSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmInstallTemplateModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
							}},
							CertManagerSpec: SpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecModel{KubeSpec: KubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmModel{
								Deployment: CertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecAffinityStrategyHpaSpecModel{
									Affinity: NodeAffinityAffinitySpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecPodAntiAffinityModel{
										NodeAffinity: DeploymentKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinitySpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecDeploymentResourcesModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
										},
										PodAffinity: PodAffinityNodeAffinityAffinitySpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecDeploymentModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
										},
										PodAntiAffinity: PodAntiAffinityKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecAffinityNodeAffinityPodAffinityModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
										},
									},
									ContainerSecurityContext: ContainerSecurityContextDeploymentKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecWindowsOptionsModel{
										AllowPrivilegeEscalation: "primitive* / BoolKind",
										Capabilities: CapabilitiesContainerSecurityContextDeploymentKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecDropModel{
											Add: func() basetypes.ListValue {
												r, diag := types.ListValue(ctx, AddCapabilitiesContainerSecurityContextDeploymentKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecSeccompProfileModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
												resp.Diagnostics.Append(diag...)
											}(),
											Drop: func() basetypes.ListValue {
												r, diag := types.ListValue(ctx, DropSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecDeploymentContainerSecurityContextCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
												resp.Diagnostics.Append(diag...)
											}(),
										},
										Privileged:             "primitive* / BoolKind",
										ProcMount:              types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
										ReadOnlyRootFilesystem: "primitive* / BoolKind",
										RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
										RunAsNonRoot:           "primitive* / BoolKind",
										RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
										SeLinuxOptions: SeLinuxOptionsContainerSecurityContextDeploymentKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmModel{
											Level: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
											Role:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
											Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
											User:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
										},
										SeccompProfile: SeccompProfileSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecDeploymentContainerSecurityContextSeLinuxOptionsModel{
											LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
											Type:             types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
										},
										WindowsOptions: WindowsOptionsSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecDeploymentContainerSecurityContextCapabilitiesModel{
											GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
											GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
											RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
										},
									},
									Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
									HpaSpec: HpaSpecStrategyAffinitySpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecModel{
										MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.HpaSpec.MaxReplicas),
										Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
										MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.HpaSpec.MinReplicas),
									},
									PodAnnotations: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, PodAnnotationsSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodAnnotations)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									PodSecurityContext: PodSecurityContextKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecAffinitySupplementalGroupsModel{
										FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.FsGroup),
										FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
										RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
										RunAsNonRoot:        "primitive* / BoolKind",
										RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
										SeLinuxOptions: SeLinuxOptionsPodSecurityContextKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecAffinityWindowsOptionsModel{
											Level: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
											Role:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
											Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
											User:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
										},
										SeccompProfile: SeccompProfilePodSecurityContextKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecAffinityNodeAffinityModel{
											LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
											Type:             types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
										},
										SupplementalGroups: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, SupplementalGroupsAffinitySpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecPodSecurityContextSeLinuxOptionsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
											resp.Diagnostics.Append(diag...)
										}(),
										Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
										WindowsOptions: WindowsOptionsAffinitySpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecPodSecurityContextSeccompProfileModel{
											GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
											GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
											RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
										},
									},
									ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ReplicaCount),
									Resources: ResourcesDeploymentKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecRequestsModel{
										Limits: func() basetypes.MapValue {
											r, diag := types.MapValue(ctx, LimitsResourcesDeploymentKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecPodAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Resources.Limits)
											resp.Diagnostics.Append(diag...)
											return r
										}(),
										Requests: func() basetypes.MapValue {
											r, diag := types.MapValue(ctx, RequestsSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecDeploymentResourcesLimitsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Resources.Requests)
											resp.Diagnostics.Append(diag...)
											return r
										}(),
									},
									Strategy: KubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecAffinityStrategyRollingUpdateModel{
										RollingUpdate: RollingUpdateStrategyAffinitySpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecMaxUnavailableModel{
											MaxSurge: MaxSurgeRollingUpdateStrategyAffinitySpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecStrValModel{
												IntVal: IntValMaxSurgeRollingUpdateStrategyAffinitySpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecPodSecurityContextModel{Value: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
												StrVal: StrValKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecAffinityStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
												Type:   types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
											},
											MaxUnavailable: MaxUnavailableKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecAffinityStrategyRollingUpdateStrValModel{
												IntVal: IntValMaxUnavailableKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecAffinityStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
												StrVal: StrValRollingUpdateStrategyAffinitySpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
												Type:   types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
											},
										},
										Type: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.Type)]),
									},
									Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
								},
								Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
								Service: ServiceKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecLabelsModel{
									Annotations: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, AnnotationsServiceKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Service.Annotations)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Labels: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, LabelsSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecKubeSpecServiceAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Service.Labels)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
									Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Service.Type)]),
								},
								ServiceAccount: ServiceAccountKubeSpecCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecHelmModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
							}},
							CertManagerStartupapicheck: ComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecModel{KubeSpec: KubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecModel{
								Deployment: DeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrategyPodAnnotationsModel{
									Affinity: NodeAffinityAffinityDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderPodAntiAffinityModel{
										NodeAffinity: ComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
										},
										PodAffinity: PodAffinityNodeAffinityAffinityDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
										},
										PodAntiAffinity: PodAntiAffinityInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityPodAffinityModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
										},
									},
									ContainerSecurityContext: ContainerSecurityContextStrategyInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentWindowsOptionsModel{
										AllowPrivilegeEscalation: "primitive* / BoolKind",
										Capabilities: CapabilitiesContainerSecurityContextStrategyInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentDropModel{
											Add: func() basetypes.ListValue {
												r, diag := types.ListValue(ctx, AddCapabilitiesContainerSecurityContextStrategyInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentSeccompProfileModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
												resp.Diagnostics.Append(diag...)
											}(),
											Drop: func() basetypes.ListValue {
												r, diag := types.ListValue(ctx, DropDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrategyContainerSecurityContextCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
												resp.Diagnostics.Append(diag...)
											}(),
										},
										Privileged:             "primitive* / BoolKind",
										ProcMount:              types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
										ReadOnlyRootFilesystem: "primitive* / BoolKind",
										RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
										RunAsNonRoot:           "primitive* / BoolKind",
										RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
										SeLinuxOptions: SeLinuxOptionsContainerSecurityContextStrategyInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentResourcesModel{
											Level: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
											Role:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
											Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
											User:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
										},
										SeccompProfile: SeccompProfileDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrategyContainerSecurityContextSeLinuxOptionsModel{
											LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
											Type:             types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
										},
										WindowsOptions: WindowsOptionsDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrategyContainerSecurityContextCapabilitiesModel{
											GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
											GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
											RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
										},
									},
									Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
									HpaSpec: HpaSpecDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrategyContainerSecurityContextModel{
										MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.HpaSpec.MaxReplicas),
										Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
										MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.HpaSpec.MinReplicas),
									},
									PodAnnotations: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, PodAnnotationsStrategyInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentHpaSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodAnnotations)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									PodSecurityContext: PodSecurityContextStrategyInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentSupplementalGroupsModel{
										FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.FsGroup),
										FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
										RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
										RunAsNonRoot:        "primitive* / BoolKind",
										RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
										SeLinuxOptions: SeLinuxOptionsPodSecurityContextStrategyInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentAffinityModel{
											Level: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
											Role:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
											Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
											User:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
										},
										SeccompProfile: SeccompProfileDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrategyPodSecurityContextSeLinuxOptionsModel{
											LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
											Type:             types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
										},
										SupplementalGroups: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, SupplementalGroupsDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrategyPodSecurityContextWindowsOptionsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
											resp.Diagnostics.Append(diag...)
										}(),
										Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
										WindowsOptions: WindowsOptionsPodSecurityContextStrategyInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentSeccompProfileModel{
											GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
											GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
											RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
										},
									},
									ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ReplicaCount),
									Resources: ResourcesDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrategyRequestsModel{
										Limits: func() basetypes.MapValue {
											r, diag := types.MapValue(ctx, LimitsResourcesDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrategyPodSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Resources.Limits)
											resp.Diagnostics.Append(diag...)
											return r
										}(),
										Requests: func() basetypes.MapValue {
											r, diag := types.MapValue(ctx, RequestsStrategyInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentResourcesLimitsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Resources.Requests)
											resp.Diagnostics.Append(diag...)
											return r
										}(),
									},
									Strategy: AffinityDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrategyRollingUpdateModel{
										RollingUpdate: RollingUpdateStrategyInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentAffinityMaxUnavailableModel{
											MaxSurge: MaxSurgeRollingUpdateStrategyInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentAffinityStrValModel{
												IntVal: IntValMaxSurgeRollingUpdateStrategyInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityModel{Value: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
												StrVal: StrValAffinityDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
												Type:   types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
											},
											MaxUnavailable: MaxUnavailableAffinityDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrategyRollingUpdateStrValModel{
												IntVal: IntValMaxUnavailableAffinityDeploymentKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
												StrVal: StrValRollingUpdateStrategyInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecDeploymentAffinityMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
												Type:   types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
											},
										},
										Type: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.Type)]),
									},
									Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
								},
								Job: KubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsAffinityNodeAffinityModel{
									Affinity: NodeAffinityAffinityComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecPodAntiAffinityModel{
										NodeAffinity: JobKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
										},
										PodAffinity: PodAffinityNodeAffinityAffinityComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecJobModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
										},
										PodAntiAffinity: PodAntiAffinityKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsAffinityNodeAffinityPodAffinityModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
										},
									},
									ContainerSecurityContext: ContainerSecurityContextJobKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsWindowsOptionsModel{
										AllowPrivilegeEscalation: "primitive* / BoolKind",
										Capabilities: CapabilitiesComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecJobContainerSecurityContextDropModel{
											Add: func() basetypes.ListValue {
												r, diag := types.ListValue(ctx, AddCapabilitiesComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecJobContainerSecurityContextSeLinuxOptionsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.Capabilities.Add)
												resp.Diagnostics.Append(diag...)
											}(),
											Drop: func() basetypes.ListValue {
												r, diag := types.ListValue(ctx, DropContainerSecurityContextJobKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.Capabilities.Drop)
												resp.Diagnostics.Append(diag...)
											}(),
										},
										Privileged:             "primitive* / BoolKind",
										ProcMount:              types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.ProcMount)]),
										ReadOnlyRootFilesystem: "primitive* / BoolKind",
										RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.RunAsGroup),
										RunAsNonRoot:           "primitive* / BoolKind",
										RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.RunAsUser),
										SeLinuxOptions: SeLinuxOptionsContainerSecurityContextJobKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecModel{
											Level: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Level)]),
											Role:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Role)]),
											Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Type)]),
											User:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.User)]),
										},
										SeccompProfile: SeccompProfileContainerSecurityContextJobKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsCapabilitiesModel{
											LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
											Type:             types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeccompProfile.Type)]),
										},
										WindowsOptions: WindowsOptionsComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecJobContainerSecurityContextSeccompProfileModel{
											GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
											GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
											RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
										},
									},
									Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
									PodAnnotations: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, PodAnnotationsComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodAnnotations)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									PodSecurityContext: PodSecurityContextJobKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSupplementalGroupsModel{
										FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.FsGroup),
										FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.FsGroupChangePolicy)]),
										RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.RunAsGroup),
										RunAsNonRoot:        "primitive* / BoolKind",
										RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.RunAsUser),
										SeLinuxOptions: SeLinuxOptionsComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecJobPodSecurityContextWindowsOptionsModel{
											Level: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeLinuxOptions.Level)]),
											Role:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeLinuxOptions.Role)]),
											Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeLinuxOptions.Type)]),
											User:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeLinuxOptions.User)]),
										},
										SeccompProfile: SeccompProfilePodSecurityContextJobKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSeLinuxOptionsModel{
											LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
											Type:             types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeccompProfile.Type)]),
										},
										SupplementalGroups: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, SupplementalGroupsComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecJobPodSecurityContextSeccompProfileModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SupplementalGroups)
											resp.Diagnostics.Append(diag...)
										}(),
										Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
										WindowsOptions: WindowsOptionsPodSecurityContextJobKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsPodAnnotationsModel{
											GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
											GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
											RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.WindowsOptions.RunAsUserName)]),
										},
									},
									Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
								},
								Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
								Service: ServiceKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsLabelsModel{
									Annotations: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, AnnotationsServiceKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Service.Annotations)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Labels: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, LabelsComponentsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckKubeSpecServiceAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Service.Labels)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
									Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Service.Type)]),
								},
								ServiceAccount: ServiceAccountKubeSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsSpecModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
							}},
							CertManagerWebhookSpec: InternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecModel{KubeSpec: KubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsModel{
								Deployment: AffinityStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorContainerSecurityContextModel{
									Affinity: NodeAffinityAffinityStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerModel{
										NodeAffinity: CertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategyAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerPodAntiAffinityModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
										},
										PodAffinity: PodAffinityAffinityStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerPodSecurityContextModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
										},
										PodAntiAffinity: PodAntiAffinityCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategyAffinityPodAffinityModel{
											PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
											RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
										},
									},
									ContainerSecurityContext: ContainerSecurityContextCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategyAffinityCapabilitiesModel{
										AllowPrivilegeEscalation: "primitive* / BoolKind",
										Capabilities: CapabilitiesAffinityStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorContainerSecurityContextAddModel{
											Add: func() basetypes.ListValue {
												r, diag := types.ListValue(ctx, AddContainerSecurityContextCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategyAffinityCapabilitiesDropModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
												resp.Diagnostics.Append(diag...)
											}(),
											Drop: func() basetypes.ListValue {
												r, diag := types.ListValue(ctx, DropCapabilitiesAffinityStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorContainerSecurityContextSeccompProfileModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
												resp.Diagnostics.Append(diag...)
											}(),
										},
										Privileged:             "primitive* / BoolKind",
										ProcMount:              types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
										ReadOnlyRootFilesystem: "primitive* / BoolKind",
										RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
										RunAsNonRoot:           "primitive* / BoolKind",
										RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
										SeLinuxOptions: SeLinuxOptionsAffinityStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorContainerSecurityContextWindowsOptionsModel{
											Level: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
											Role:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
											Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
											User:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
										},
										SeccompProfile: SeccompProfileContainerSecurityContextCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategyAffinitySeLinuxOptionsModel{
											LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
											Type:             types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
										},
										WindowsOptions: WindowsOptionsContainerSecurityContextCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategyAffinityNodeAffinityModel{
											GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
											GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
											RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
										},
									},
									Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
									HpaSpec: HpaSpecStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderModel{
										MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.HpaSpec.MaxReplicas),
										Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
										MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.HpaSpec.MinReplicas),
									},
									PodAnnotations: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, PodAnnotationsCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategyHpaSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodAnnotations)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									PodSecurityContext: PodSecurityContextCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategySeLinuxOptionsModel{
										FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.FsGroup),
										FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
										RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
										RunAsNonRoot:        "primitive* / BoolKind",
										RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
										SeLinuxOptions: SeLinuxOptionsStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerPodSecurityContextWindowsOptionsModel{
											Level: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
											Role:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
											Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
											User:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
										},
										SeccompProfile: SeccompProfilePodSecurityContextCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategyResourcesModel{
											LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
											Type:             types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
										},
										SupplementalGroups: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, SupplementalGroupsStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerPodSecurityContextSeccompProfileModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
											resp.Diagnostics.Append(diag...)
										}(),
										Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
										WindowsOptions: WindowsOptionsPodSecurityContextCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategySupplementalGroupsModel{
											GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
											GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
											RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
										},
									},
									ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ReplicaCount),
									Resources: ResourcesStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerRequestsModel{
										Limits: func() basetypes.MapValue {
											r, diag := types.MapValue(ctx, LimitsResourcesStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerPodAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Resources.Limits)
											resp.Diagnostics.Append(diag...)
											return r
										}(),
										Requests: func() basetypes.MapValue {
											r, diag := types.MapValue(ctx, RequestsCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategyResourcesLimitsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Resources.Requests)
											resp.Diagnostics.Append(diag...)
											return r
										}(),
									},
									Strategy: InternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateModel{
										RollingUpdate: RollingUpdateStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderMaxUnavailableModel{
											MaxSurge: MaxSurgeRollingUpdateStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderStrValModel{
												IntVal: IntValMaxSurgeRollingUpdateStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsModel{Value: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
												StrVal: StrValInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
												Type:   types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
											},
											MaxUnavailable: MaxUnavailableInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateStrValModel{
												IntVal: IntValMaxUnavailableInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
												StrVal: StrValRollingUpdateStrategyDeploymentKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
												Type:   types.Int64Value(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
											},
										},
										Type: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.Type)]),
									},
									Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
								},
								Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
								Service: ServiceKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderLabelsModel{
									Annotations: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, AnnotationsServiceKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Service.Annotations)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Labels: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, LabelsInternalCertProviderCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecKubeSpecServiceAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Service.Labels)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
									Type:  types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Service.Type)]),
								},
								ServiceAccount: ServiceAccountKubeSpecCertManagerWebhookSpecCertManagerStartupapicheckCertManagerSpecCertManagerCaInjectorCertManagerInternalCertProviderComponentsModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
							}},
							Managed: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.Managed[int32(cluster.Helm.Spec.Components.InternalCertProvider.CertManager.Managed)]),
						},
						Custom: CustomCertManagerCertManagerCaInjectorCertManagerSpecCertManagerStartupapicheckCertManagerWebhookSpecModel{
							CaBundleSecretName: types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.Custom.CaBundleSecretName[int32(cluster.Helm.Spec.Components.InternalCertProvider.Custom.CaBundleSecretName)]),
							CsrSignerName:      types.StringValue(cluster.Helm.Spec.Components.InternalCertProvider.Custom.CsrSignerName[int32(cluster.Helm.Spec.Components.InternalCertProvider.Custom.CsrSignerName)]),
						},
					},
					Istio: IstioComponentsSpecHelmDeploymentModel{
						BaseOverlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
						CniOverlays:  "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
						DefaultWorkloadCert_TTL: DefaultWorkloadCert_TTLHelmSpecComponentsIstioMaxWorkloadCert_TTLModel{
							Nanos:   types.Int64Value(cluster.Helm.Spec.Components.Istio.DefaultWorkloadCert_TTL.Nanos),
							Seconds: types.Int64Value(cluster.Helm.Spec.Components.Istio.DefaultWorkloadCert_TTL.Seconds),
						},
						KubeSpec: DeploymentHelmSpecComponentsIstioServiceAccountModel{
							CNI: CNIKubeSpecIstioComponentsSpecHelmDefaultWorkloadCert_TTLModel{
								BinaryDirectory:        types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.CNI.BinaryDirectory[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.CNI.BinaryDirectory)]),
								Chained:                "primitive / BoolKind",
								ClusterRole:            types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.CNI.ClusterRole[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.CNI.ClusterRole)]),
								ConfigurationDirectory: types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.CNI.ConfigurationDirectory[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.CNI.ConfigurationDirectory)]),
								ConfigurationFileName:  types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.CNI.ConfigurationFileName[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.CNI.ConfigurationFileName)]),
								Revision:               types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.CNI.Revision[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.CNI.Revision)]),
							},
							Deployment: KubeSpecIstioComponentsSpecHelmDeploymentAffinityModel{
								Affinity: AffinityDeploymentHelmSpecComponentsIstioKubeSpecStrategyModel{
									NodeAffinity: KubeSpecIstioComponentsSpecHelmDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityDeploymentHelmSpecComponentsIstioKubeSpecStrategyModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
									},
									PodAffinity: PodAffinityAffinityDeploymentHelmSpecComponentsIstioKubeSpecStrategyModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
									PodAntiAffinity: PodAntiAffinityAffinityDeploymentHelmSpecComponentsIstioKubeSpecStrategyModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
								},
								ContainerSecurityContext: ContainerSecurityContextDeploymentHelmSpecComponentsIstioKubeSpec_CNIModel{
									AllowPrivilegeEscalation: "primitive* / BoolKind",
									Capabilities: CapabilitiesContainerSecurityContextDeploymentHelmSpecComponentsIstioKubeSpecDropModel{
										Add: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, AddCapabilitiesContainerSecurityContextDeploymentHelmSpecComponentsIstioKubeSpec_CNIModel.ElementType(ctx), cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
											resp.Diagnostics.Append(diag...)
										}(),
										Drop: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, DropKubeSpecIstioComponentsSpecHelmDeploymentContainerSecurityContextCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
											resp.Diagnostics.Append(diag...)
										}(),
									},
									Privileged:             "primitive* / BoolKind",
									ProcMount:              types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
									ReadOnlyRootFilesystem: "primitive* / BoolKind",
									RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
									RunAsNonRoot:           "primitive* / BoolKind",
									RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsContainerSecurityContextDeploymentHelmSpecComponentsIstioKubeSpec_CNIModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfileContainerSecurityContextDeploymentHelmSpecComponentsIstioKubeSpec_CNIModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
									},
									WindowsOptions: WindowsOptionsContainerSecurityContextDeploymentHelmSpecComponentsIstioKubeSpec_CNIModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
								HpaSpec: HpaSpecKubeSpecIstioComponentsSpecHelmDeploymentContainerSecurityContextModel{
									MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.HpaSpec.MaxReplicas),
									Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
									MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.HpaSpec.MinReplicas),
								},
								PodAnnotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, PodAnnotationsDeploymentHelmSpecComponentsIstioKubeSpecHpaSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodAnnotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								PodSecurityContext: PodSecurityContextKubeSpecIstioComponentsSpecHelmDeploymentPodAnnotationsModel{
									FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.FsGroup),
									FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
									RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
									RunAsNonRoot:        "primitive* / BoolKind",
									RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsPodSecurityContextKubeSpecIstioComponentsSpecHelmDeploymentPodAnnotationsModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfilePodSecurityContextKubeSpecIstioComponentsSpecHelmDeploymentPodAnnotationsModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
									},
									SupplementalGroups: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, SupplementalGroupsPodSecurityContextKubeSpecIstioComponentsSpecHelmDeploymentPodAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
										resp.Diagnostics.Append(diag...)
									}(),
									Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
									WindowsOptions: WindowsOptionsPodSecurityContextKubeSpecIstioComponentsSpecHelmDeploymentPodAnnotationsModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.ReplicaCount),
								Resources: ResourcesDeploymentHelmSpecComponentsIstioKubeSpecPodSecurityContextModel{
									Limits: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, LimitsResourcesDeploymentHelmSpecComponentsIstioKubeSpecPodSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.Resources.Limits)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Requests: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, RequestsResourcesDeploymentHelmSpecComponentsIstioKubeSpecPodSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.Resources.Requests)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
								},
								Strategy: StrategyKubeSpecIstioComponentsSpecHelmDeploymentResourcesModel{
									RollingUpdate: RollingUpdateStrategyKubeSpecIstioComponentsSpecHelmDeploymentMaxUnavailableModel{
										MaxSurge: MaxSurgeRollingUpdateStrategyKubeSpecIstioComponentsSpecHelmDeploymentStrValModel{
											IntVal: IntValMaxSurgeRollingUpdateStrategyKubeSpecIstioComponentsSpecHelmDeploymentResourcesModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
											StrVal: StrValDeploymentHelmSpecComponentsIstioKubeSpecStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
										},
										MaxUnavailable: MaxUnavailableDeploymentHelmSpecComponentsIstioKubeSpecStrategyRollingUpdateStrValModel{
											IntVal: IntValMaxUnavailableDeploymentHelmSpecComponentsIstioKubeSpecStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
											StrVal: StrValRollingUpdateStrategyKubeSpecIstioComponentsSpecHelmDeploymentMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
										},
									},
									Type: types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.Type)]),
								},
								Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
							},
							Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
							Service: ServiceDeploymentHelmSpecComponentsIstioAnnotationsModel{
								Annotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, AnnotationsIstioComponentsSpecHelmDeploymentServiceLabelsModel.ElementType(ctx), cluster.Helm.Spec.Components.Istio.KubeSpec.Service.Annotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Labels: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, LabelsServiceDeploymentHelmSpecComponentsIstioKubeSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.Istio.KubeSpec.Service.Labels)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
								Type:  types.StringValue(cluster.Helm.Spec.Components.Istio.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.Istio.KubeSpec.Service.Type)]),
							},
							ServiceAccount: ServiceAccountIstioComponentsSpecHelmDeploymentServiceModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
						},
						MaxWorkloadCert_TTL: MaxWorkloadCert_TTLIstioComponentsSpecHelmInstallTemplateModel{
							Nanos:   types.Int64Value(cluster.Helm.Spec.Components.Istio.MaxWorkloadCert_TTL.Nanos),
							Seconds: types.Int64Value(cluster.Helm.Spec.Components.Istio.MaxWorkloadCert_TTL.Seconds),
						},
						PilotOverlays:     "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
						TraceSamplingRate: types.Float64Value(cluster.Helm.Spec.Components.Istio.TraceSamplingRate),
						TrustDomain:       types.StringValue(cluster.Helm.Spec.Components.Istio.TrustDomain[int32(cluster.Helm.Spec.Components.Istio.TrustDomain)]),
						TsbVersion:        types.StringValue(cluster.Helm.Spec.Components.Istio.TsbVersion[int32(cluster.Helm.Spec.Components.Istio.TsbVersion)]),
					},
					Ngac: ComponentsSpecHelmDeploymentLogLevelsModel{
						Enabled: "primitive / BoolKind",
						KubeSpec: NgacComponentsSpecHelmDeploymentServiceModel{
							Deployment: KubeSpecNgacComponentsSpecHelmDeploymentStrategyModel{
								Affinity: AffinityKubeSpecNgacComponentsSpecHelmDeploymentHpaSpecModel{
									NodeAffinity: DeploymentHelmSpecComponentsNgacKubeSpecAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityKubeSpecNgacComponentsSpecHelmDeploymentHpaSpecModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
									},
									PodAffinity: PodAffinityAffinityKubeSpecNgacComponentsSpecHelmDeploymentHpaSpecModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
									PodAntiAffinity: PodAntiAffinityAffinityKubeSpecNgacComponentsSpecHelmDeploymentHpaSpecModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
								},
								ContainerSecurityContext: ContainerSecurityContextDeploymentHelmSpecComponentsNgacKubeSpecAffinityModel{
									AllowPrivilegeEscalation: "primitive* / BoolKind",
									Capabilities: CapabilitiesContainerSecurityContextDeploymentHelmSpecComponentsNgacKubeSpecDropModel{
										Add: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, AddCapabilitiesContainerSecurityContextDeploymentHelmSpecComponentsNgacKubeSpecAffinityModel.ElementType(ctx), cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
											resp.Diagnostics.Append(diag...)
										}(),
										Drop: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, DropKubeSpecNgacComponentsSpecHelmDeploymentContainerSecurityContextCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
											resp.Diagnostics.Append(diag...)
										}(),
									},
									Privileged:             "primitive* / BoolKind",
									ProcMount:              types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
									ReadOnlyRootFilesystem: "primitive* / BoolKind",
									RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
									RunAsNonRoot:           "primitive* / BoolKind",
									RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsContainerSecurityContextDeploymentHelmSpecComponentsNgacKubeSpecAffinityModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfileContainerSecurityContextDeploymentHelmSpecComponentsNgacKubeSpecAffinityModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
									},
									WindowsOptions: WindowsOptionsContainerSecurityContextDeploymentHelmSpecComponentsNgacKubeSpecAffinityModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
								HpaSpec: HpaSpecDeploymentHelmSpecComponentsNgacKubeSpecServiceAccountModel{
									MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.HpaSpec.MaxReplicas),
									Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
									MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.HpaSpec.MinReplicas),
								},
								PodAnnotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, PodAnnotationsKubeSpecNgacComponentsSpecHelmDeploymentContainerSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodAnnotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								PodSecurityContext: PodSecurityContextDeploymentHelmSpecComponentsNgacKubeSpecPodAnnotationsModel{
									FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.FsGroup),
									FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
									RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
									RunAsNonRoot:        "primitive* / BoolKind",
									RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsPodSecurityContextDeploymentHelmSpecComponentsNgacKubeSpecPodAnnotationsModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfilePodSecurityContextDeploymentHelmSpecComponentsNgacKubeSpecPodAnnotationsModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
									},
									SupplementalGroups: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, SupplementalGroupsPodSecurityContextDeploymentHelmSpecComponentsNgacKubeSpecPodAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
										resp.Diagnostics.Append(diag...)
									}(),
									Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
									WindowsOptions: WindowsOptionsPodSecurityContextDeploymentHelmSpecComponentsNgacKubeSpecPodAnnotationsModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ReplicaCount),
								Resources: ResourcesKubeSpecNgacComponentsSpecHelmDeploymentPodSecurityContextModel{
									Limits: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, LimitsResourcesKubeSpecNgacComponentsSpecHelmDeploymentPodSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Resources.Limits)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Requests: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, RequestsResourcesKubeSpecNgacComponentsSpecHelmDeploymentPodSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Resources.Requests)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
								},
								Strategy: StrategyDeploymentHelmSpecComponentsNgacKubeSpecResourcesModel{
									RollingUpdate: RollingUpdateStrategyDeploymentHelmSpecComponentsNgacKubeSpecMaxUnavailableModel{
										MaxSurge: MaxSurgeRollingUpdateStrategyDeploymentHelmSpecComponentsNgacKubeSpecIntValModel{
											IntVal: IntValKubeSpecNgacComponentsSpecHelmDeploymentStrategyRollingUpdateMaxSurgeStrValModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
											StrVal: StrValMaxSurgeRollingUpdateStrategyDeploymentHelmSpecComponentsNgacKubeSpecResourcesModel{Value: types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
										},
										MaxUnavailable: MaxUnavailableKubeSpecNgacComponentsSpecHelmDeploymentStrategyRollingUpdateStrValModel{
											IntVal: IntValMaxUnavailableKubeSpecNgacComponentsSpecHelmDeploymentStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
											StrVal: StrValRollingUpdateStrategyDeploymentHelmSpecComponentsNgacKubeSpecMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
										},
									},
									Type: types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.Type)]),
								},
								Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
							},
							Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
							Service: ServiceDeploymentHelmSpecComponentsNgacLabelsModel{
								Annotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, AnnotationsServiceDeploymentHelmSpecComponentsNgacKubeSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.Ngac.KubeSpec.Service.Annotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Labels: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, LabelsNgacComponentsSpecHelmDeploymentServiceAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.Ngac.KubeSpec.Service.Labels)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
								Type:  types.StringValue(cluster.Helm.Spec.Components.Ngac.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.Ngac.KubeSpec.Service.Type)]),
							},
							ServiceAccount: ServiceAccountKubeSpecNgacComponentsSpecHelmInstallTemplateModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
						},
						LogLevels: func() basetypes.MapValue {
							r, diag := types.MapValue(ctx, LogLevelsDeploymentHelmSpecComponentsNgacModel.ElementType(ctx), cluster.Helm.Spec.Components.Ngac.LogLevels)
							resp.Diagnostics.Append(diag...)
							return r
						}(),
					},
					Oap: SpecComponentsOapKubeSpecDeploymentModel{
						KubeSpec: DeploymentKubeSpecOapComponentsSpecHelmModel{
							Deployment: HelmSpecComponentsOapKubeSpecDeploymentResourcesModel{
								Affinity: AffinityDeploymentKubeSpecOapComponentsSpecHelmStrategyModel{
									NodeAffinity: HelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityDeploymentKubeSpecOapComponentsSpecHelmStrategyModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
									},
									PodAffinity: PodAffinityAffinityDeploymentKubeSpecOapComponentsSpecHelmStrategyModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
									PodAntiAffinity: PodAntiAffinityAffinityDeploymentKubeSpecOapComponentsSpecHelmStrategyModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
								},
								ContainerSecurityContext: ContainerSecurityContextHelmSpecComponentsOapKubeSpecDeploymentAffinityModel{
									AllowPrivilegeEscalation: "primitive* / BoolKind",
									Capabilities: CapabilitiesContainerSecurityContextHelmSpecComponentsOapKubeSpecDeploymentDropModel{
										Add: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, AddCapabilitiesContainerSecurityContextHelmSpecComponentsOapKubeSpecDeploymentAffinityModel.ElementType(ctx), cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
											resp.Diagnostics.Append(diag...)
										}(),
										Drop: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, DropDeploymentKubeSpecOapComponentsSpecHelmContainerSecurityContextCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
											resp.Diagnostics.Append(diag...)
										}(),
									},
									Privileged:             "primitive* / BoolKind",
									ProcMount:              types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
									ReadOnlyRootFilesystem: "primitive* / BoolKind",
									RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
									RunAsNonRoot:           "primitive* / BoolKind",
									RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsContainerSecurityContextHelmSpecComponentsOapKubeSpecDeploymentAffinityModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfileContainerSecurityContextHelmSpecComponentsOapKubeSpecDeploymentAffinityModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
									},
									WindowsOptions: WindowsOptionsContainerSecurityContextHelmSpecComponentsOapKubeSpecDeploymentAffinityModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
								HpaSpec: HpaSpecDeploymentKubeSpecOapComponentsSpecHelmServiceAccountModel{
									MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.HpaSpec.MaxReplicas),
									Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
									MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.HpaSpec.MinReplicas),
								},
								PodAnnotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, PodAnnotationsHelmSpecComponentsOapKubeSpecDeploymentHpaSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodAnnotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								PodSecurityContext: PodSecurityContextDeploymentKubeSpecOapComponentsSpecHelmPodAnnotationsModel{
									FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.FsGroup),
									FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
									RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
									RunAsNonRoot:        "primitive* / BoolKind",
									RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsPodSecurityContextDeploymentKubeSpecOapComponentsSpecHelmPodAnnotationsModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfilePodSecurityContextDeploymentKubeSpecOapComponentsSpecHelmPodAnnotationsModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
									},
									SupplementalGroups: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, SupplementalGroupsPodSecurityContextDeploymentKubeSpecOapComponentsSpecHelmPodAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
										resp.Diagnostics.Append(diag...)
									}(),
									Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
									WindowsOptions: WindowsOptionsPodSecurityContextDeploymentKubeSpecOapComponentsSpecHelmPodAnnotationsModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.ReplicaCount),
								Resources: ResourcesDeploymentKubeSpecOapComponentsSpecHelmContainerSecurityContextModel{
									Limits: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, LimitsResourcesDeploymentKubeSpecOapComponentsSpecHelmContainerSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.Resources.Limits)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Requests: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, RequestsResourcesDeploymentKubeSpecOapComponentsSpecHelmContainerSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.Resources.Requests)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
								},
								Strategy: StrategyHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel{
									RollingUpdate: RollingUpdateStrategyHelmSpecComponentsOapKubeSpecDeploymentMaxUnavailableModel{
										MaxSurge: MaxSurgeRollingUpdateStrategyHelmSpecComponentsOapKubeSpecDeploymentStrValModel{
											IntVal: IntValMaxSurgeRollingUpdateStrategyHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
											StrVal: StrValDeploymentKubeSpecOapComponentsSpecHelmStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
										},
										MaxUnavailable: MaxUnavailableDeploymentKubeSpecOapComponentsSpecHelmStrategyRollingUpdateStrValModel{
											IntVal: IntValMaxUnavailableDeploymentKubeSpecOapComponentsSpecHelmStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
											StrVal: StrValRollingUpdateStrategyHelmSpecComponentsOapKubeSpecDeploymentMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
										},
									},
									Type: types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.Type)]),
								},
								Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
							},
							Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
							Service: ServiceKubeSpecOapComponentsSpecHelmLabelsModel{
								Annotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, AnnotationsServiceKubeSpecOapComponentsSpecHelmInstallTemplateModel.ElementType(ctx), cluster.Helm.Spec.Components.Oap.KubeSpec.Service.Annotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Labels: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, LabelsHelmSpecComponentsOapKubeSpecServiceAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.Oap.KubeSpec.Service.Labels)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
								Type:  types.StringValue(cluster.Helm.Spec.Components.Oap.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.Oap.KubeSpec.Service.Type)]),
							},
							ServiceAccount: ServiceAccountHelmSpecComponentsOapKubeSpecServiceModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
						},
						OnDemandEnvoyMetricsEnabled: "primitive / BoolKind",
						StorageIndexMergingEnabled:  "primitive / BoolKind",
						StreamingLogEnabled:         "primitive / BoolKind",
					},
					Onboarding: SpecComponentsOnboardingRepositoryKubeSpecModel{
						Operator: KubeSpecOperatorOnboardingComponentsSpecHelmModel{KubeSpec: HelmSpecComponentsOnboardingOperatorKubeSpecServiceAccountModel{
							Deployment: DeploymentKubeSpecOperatorOnboardingComponentsSpecHelmInstallTemplateModel{
								Affinity: NodeAffinityAffinityDeploymentKubeSpecOperatorOnboardingComponentsSpecPodAntiAffinityModel{
									NodeAffinity: HelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmInstallTemplateModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
									},
									PodAffinity: PodAffinityNodeAffinityAffinityDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
									PodAntiAffinity: PodAntiAffinitySpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityPodAffinityModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
								},
								ContainerSecurityContext: ContainerSecurityContextDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmCapabilitiesModel{
									AllowPrivilegeEscalation: "primitive* / BoolKind",
									Capabilities: CapabilitiesHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextDropModel{
										Add: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, AddCapabilitiesHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextWindowsOptionsModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
											resp.Diagnostics.Append(diag...)
										}(),
										Drop: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, DropContainerSecurityContextDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
											resp.Diagnostics.Append(diag...)
										}(),
									},
									Privileged:             "primitive* / BoolKind",
									ProcMount:              types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
									ReadOnlyRootFilesystem: "primitive* / BoolKind",
									RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
									RunAsNonRoot:           "primitive* / BoolKind",
									RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsContainerSecurityContextDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmInstallTemplateModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfileHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextSeLinuxOptionsModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
									},
									WindowsOptions: WindowsOptionsContainerSecurityContextDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmSeccompProfileModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
								HpaSpec: HpaSpecDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmInstallTemplateModel{
									MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.HpaSpec.MaxReplicas),
									Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
									MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.HpaSpec.MinReplicas),
								},
								PodAnnotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, PodAnnotationsDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmInstallTemplateModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodAnnotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								PodSecurityContext: PodSecurityContextDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmSupplementalGroupsModel{
									FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.FsGroup),
									FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
									RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
									RunAsNonRoot:        "primitive* / BoolKind",
									RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsPodSecurityContextDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmWindowsOptionsModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfilePodSecurityContextDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmInstallTemplateModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
									},
									SupplementalGroups: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, SupplementalGroupsHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextSeLinuxOptionsModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
										resp.Diagnostics.Append(diag...)
									}(),
									Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
									WindowsOptions: WindowsOptionsHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextSeccompProfileModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ReplicaCount),
								Resources: ResourcesDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmRequestsModel{
									Limits: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, LimitsResourcesDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmInstallTemplateModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Resources.Limits)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Requests: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, RequestsHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentResourcesLimitsModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Resources.Requests)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
								},
								Strategy: HelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateModel{
									RollingUpdate: RollingUpdateStrategyDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmMaxUnavailableModel{
										MaxSurge: MaxSurgeRollingUpdateStrategyDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmStrValModel{
											IntVal: IntValMaxSurgeRollingUpdateStrategyDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmInstallTemplateModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
											StrVal: StrValHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
										},
										MaxUnavailable: MaxUnavailableHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateStrValModel{
											IntVal: IntValMaxUnavailableHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
											StrVal: StrValRollingUpdateStrategyDeploymentKubeSpecOperatorOnboardingComponentsSpecHelmMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
										},
									},
									Type: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.Type)]),
								},
								Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
							},
							Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
							Service: ServiceHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel{
								Annotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, AnnotationsServiceHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Service.Annotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Labels: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, LabelsServiceHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Service.Labels)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
								Type:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Service.Type)]),
							},
							ServiceAccount: ServiceAccountKubeSpecOperatorOnboardingComponentsSpecHelmServiceModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
						}},
						Plane: InstancePlaneSpecComponentsOnboardingOperatorModel{Instance: OperatorOnboardingComponentsSpecPlaneInstanceKubeSpecModel{KubeSpec: KubeSpecInstancePlaneSpecComponentsOnboardingOperatorKubeSpecModel{
							Deployment: AffinityStrategyDeploymentKubeSpecInstancePlaneSpecComponentsHpaSpecModel{
								Affinity: NodeAffinityAffinityStrategyDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingModel{
									NodeAffinity: OnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentStrategyAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityStrategyDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingPodAntiAffinityModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
									},
									PodAffinity: PodAffinityAffinityStrategyDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingOperatorModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
									PodAntiAffinity: PodAntiAffinityOnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentStrategyAffinityPodAffinityModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
								},
								ContainerSecurityContext: ContainerSecurityContextDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingOperatorSeLinuxOptionsModel{
									AllowPrivilegeEscalation: "primitive* / BoolKind",
									Capabilities: CapabilitiesContainerSecurityContextDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingOperatorDropModel{
										Add: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, AddCapabilitiesContainerSecurityContextDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingOperatorKubeSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
											resp.Diagnostics.Append(diag...)
										}(),
										Drop: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, DropOperatorOnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentContainerSecurityContextCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
											resp.Diagnostics.Append(diag...)
										}(),
									},
									Privileged:             "primitive* / BoolKind",
									ProcMount:              types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
									ReadOnlyRootFilesystem: "primitive* / BoolKind",
									RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
									RunAsNonRoot:           "primitive* / BoolKind",
									RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsOperatorOnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentContainerSecurityContextWindowsOptionsModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfileOperatorOnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentContainerSecurityContextCapabilitiesModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
									},
									WindowsOptions: WindowsOptionsContainerSecurityContextDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingOperatorSeccompProfileModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
								HpaSpec: HpaSpecComponentsSpecPlaneInstanceKubeSpecDeploymentStrategyAffinityNodeAffinityModel{
									MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.HpaSpec.MaxReplicas),
									Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
									MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.HpaSpec.MinReplicas),
								},
								PodAnnotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, PodAnnotationsOperatorOnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodAnnotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								PodSecurityContext: PodSecurityContextDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingOperatorSupplementalGroupsModel{
									FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.FsGroup),
									FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
									RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
									RunAsNonRoot:        "primitive* / BoolKind",
									RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsOperatorOnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentPodSecurityContextWindowsOptionsModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfilePodSecurityContextDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingOperatorSeLinuxOptionsModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
									},
									SupplementalGroups: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, SupplementalGroupsOperatorOnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentPodSecurityContextSeccompProfileModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
										resp.Diagnostics.Append(diag...)
									}(),
									Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
									WindowsOptions: WindowsOptionsPodSecurityContextDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingOperatorPodAnnotationsModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ReplicaCount),
								Resources: ResourcesOperatorOnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentRequestsModel{
									Limits: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, LimitsResourcesOperatorOnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentPodSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Resources.Limits)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Requests: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, RequestsDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingOperatorResourcesLimitsModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Resources.Requests)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
								},
								Strategy: OperatorOnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateModel{
									RollingUpdate: RollingUpdateStrategyDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingOperatorMaxUnavailableModel{
										MaxSurge: MaxSurgeRollingUpdateStrategyDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingOperatorStrValModel{
											IntVal: IntValMaxSurgeRollingUpdateStrategyDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingOperatorResourcesModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
											StrVal: StrValOperatorOnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
										},
										MaxUnavailable: MaxUnavailableOperatorOnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateStrValModel{
											IntVal: IntValMaxUnavailableOperatorOnboardingComponentsSpecPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
											StrVal: StrValRollingUpdateStrategyDeploymentKubeSpecInstancePlaneSpecComponentsOnboardingOperatorMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
										},
									},
									Type: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.Type)]),
								},
								Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
							},
							Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
							Service: ServiceKubeSpecInstancePlaneSpecComponentsOnboardingOperatorAnnotationsModel{
								Annotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, AnnotationsOperatorOnboardingComponentsSpecPlaneInstanceKubeSpecServiceLabelsModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Service.Annotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Labels: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, LabelsServiceKubeSpecInstancePlaneSpecComponentsOnboardingOperatorKubeSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Service.Labels)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
								Type:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Service.Type)]),
							},
							ServiceAccount: ServiceAccountKubeSpecInstancePlaneSpecComponentsOnboardingOperatorKubeSpecModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
						}}},
						Repository: KubeSpecRepositoryOnboardingComponentsSpecPlaneModel{KubeSpec: PlaneSpecComponentsOnboardingRepositoryKubeSpecServiceAccountModel{
							Deployment: DeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneInstanceModel{
								Affinity: NodeAffinityAffinityDeploymentKubeSpecRepositoryOnboardingComponentsSpecPodAntiAffinityModel{
									NodeAffinity: PlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneInstanceModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
									},
									PodAffinity: PodAffinityNodeAffinityAffinityDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
									PodAntiAffinity: PodAntiAffinitySpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityPodAffinityModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
								},
								ContainerSecurityContext: ContainerSecurityContextDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneSeccompProfileModel{
									AllowPrivilegeEscalation: "primitive* / BoolKind",
									Capabilities: CapabilitiesContainerSecurityContextDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneDropModel{
										Add: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, AddCapabilitiesContainerSecurityContextDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneWindowsOptionsModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
											resp.Diagnostics.Append(diag...)
										}(),
										Drop: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, DropPlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
											resp.Diagnostics.Append(diag...)
										}(),
									},
									Privileged:             "primitive* / BoolKind",
									ProcMount:              types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
									ReadOnlyRootFilesystem: "primitive* / BoolKind",
									RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
									RunAsNonRoot:           "primitive* / BoolKind",
									RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsContainerSecurityContextDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneInstanceModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfilePlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextCapabilitiesModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
									},
									WindowsOptions: WindowsOptionsPlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextSeLinuxOptionsModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
								HpaSpec: HpaSpecDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneInstanceModel{
									MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.HpaSpec.MaxReplicas),
									Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
									MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.HpaSpec.MinReplicas),
								},
								PodAnnotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, PodAnnotationsDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneInstanceModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodAnnotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								PodSecurityContext: PodSecurityContextDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneWindowsOptionsModel{
									FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.FsGroup),
									FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
									RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
									RunAsNonRoot:        "primitive* / BoolKind",
									RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsPodSecurityContextDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneInstanceModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfilePodSecurityContextDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneSupplementalGroupsModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
									},
									SupplementalGroups: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, SupplementalGroupsPlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextSeLinuxOptionsModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
										resp.Diagnostics.Append(diag...)
									}(),
									Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
									WindowsOptions: WindowsOptionsPlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextSeccompProfileModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ReplicaCount),
								Resources: ResourcesDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneLimitsModel{
									Limits: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, LimitsPlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentResourcesRequestsModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Resources.Limits)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Requests: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, RequestsResourcesDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneInstanceModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Resources.Requests)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
								},
								Strategy: PlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateModel{
									RollingUpdate: RollingUpdateStrategyDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneMaxUnavailableModel{
										MaxSurge: MaxSurgeRollingUpdateStrategyDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneStrValModel{
											IntVal: IntValMaxSurgeRollingUpdateStrategyDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneInstanceModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
											StrVal: StrValPlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
										},
										MaxUnavailable: MaxUnavailablePlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateStrValModel{
											IntVal: IntValMaxUnavailablePlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
											StrVal: StrValRollingUpdateStrategyDeploymentKubeSpecRepositoryOnboardingComponentsSpecPlaneMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
										},
									},
									Type: types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.Type)]),
								},
								Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
							},
							Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
							Service: ServicePlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel{
								Annotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, AnnotationsServicePlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Service.Annotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Labels: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, LabelsServicePlaneSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel.ElementType(ctx), cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Service.Labels)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
								Type:  types.StringValue(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Service.Type)]),
							},
							ServiceAccount: ServiceAccountKubeSpecRepositoryOnboardingComponentsSpecPlaneServiceModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
						}},
					},
					RateLimitServer: DeploymentSpecComponentsRateLimitServerBackendModel{
						Backend: HelmSpecComponentsRateLimitServerBackendRedisModel{Redis: RedisBackendRateLimitServerComponentsSpecHelmInstallTemplateModel{Uri: types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.Backend.Redis.Uri[int32(cluster.Helm.Spec.Components.RateLimitServer.Backend.Redis.Uri)])}},
						Domain:  types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.Domain[int32(cluster.Helm.Spec.Components.RateLimitServer.Domain)]),
						KubeSpec: BackendRateLimitServerComponentsSpecDeploymentServiceModel{
							Deployment: KubeSpecBackendRateLimitServerComponentsSpecDeploymentAffinityModel{
								Affinity: AffinityDeploymentSpecComponentsRateLimitServerBackendKubeSpecPodSecurityContextModel{
									NodeAffinity: KubeSpecBackendRateLimitServerComponentsSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityDeploymentSpecComponentsRateLimitServerBackendKubeSpecPodSecurityContextModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
									},
									PodAffinity: PodAffinityAffinityDeploymentSpecComponentsRateLimitServerBackendKubeSpecPodSecurityContextModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
									PodAntiAffinity: PodAntiAffinityAffinityDeploymentSpecComponentsRateLimitServerBackendKubeSpecPodSecurityContextModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
								},
								ContainerSecurityContext: ContainerSecurityContextKubeSpecBackendRateLimitServerComponentsSpecDeploymentStrategyModel{
									AllowPrivilegeEscalation: "primitive* / BoolKind",
									Capabilities: CapabilitiesContainerSecurityContextKubeSpecBackendRateLimitServerComponentsSpecDeploymentDropModel{
										Add: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, AddCapabilitiesContainerSecurityContextKubeSpecBackendRateLimitServerComponentsSpecDeploymentStrategyModel.ElementType(ctx), cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
											resp.Diagnostics.Append(diag...)
										}(),
										Drop: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, DropDeploymentSpecComponentsRateLimitServerBackendKubeSpecContainerSecurityContextCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
											resp.Diagnostics.Append(diag...)
										}(),
									},
									Privileged:             "primitive* / BoolKind",
									ProcMount:              types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
									ReadOnlyRootFilesystem: "primitive* / BoolKind",
									RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
									RunAsNonRoot:           "primitive* / BoolKind",
									RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsContainerSecurityContextKubeSpecBackendRateLimitServerComponentsSpecDeploymentStrategyModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfileContainerSecurityContextKubeSpecBackendRateLimitServerComponentsSpecDeploymentStrategyModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
									},
									WindowsOptions: WindowsOptionsContainerSecurityContextKubeSpecBackendRateLimitServerComponentsSpecDeploymentStrategyModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
								HpaSpec: HpaSpecDeploymentSpecComponentsRateLimitServerBackendKubeSpecServiceAccountModel{
									MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.HpaSpec.MaxReplicas),
									Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
									MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.HpaSpec.MinReplicas),
								},
								PodAnnotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, PodAnnotationsDeploymentSpecComponentsRateLimitServerBackendKubeSpecContainerSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodAnnotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								PodSecurityContext: PodSecurityContextKubeSpecBackendRateLimitServerComponentsSpecDeploymentPodAnnotationsModel{
									FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.FsGroup),
									FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
									RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
									RunAsNonRoot:        "primitive* / BoolKind",
									RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsPodSecurityContextKubeSpecBackendRateLimitServerComponentsSpecDeploymentPodAnnotationsModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfilePodSecurityContextKubeSpecBackendRateLimitServerComponentsSpecDeploymentPodAnnotationsModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
									},
									SupplementalGroups: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, SupplementalGroupsPodSecurityContextKubeSpecBackendRateLimitServerComponentsSpecDeploymentPodAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
										resp.Diagnostics.Append(diag...)
									}(),
									Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
									WindowsOptions: WindowsOptionsPodSecurityContextKubeSpecBackendRateLimitServerComponentsSpecDeploymentPodAnnotationsModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ReplicaCount),
								Resources: ResourcesKubeSpecBackendRateLimitServerComponentsSpecDeploymentHpaSpecModel{
									Limits: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, LimitsResourcesKubeSpecBackendRateLimitServerComponentsSpecDeploymentHpaSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Resources.Limits)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Requests: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, RequestsResourcesKubeSpecBackendRateLimitServerComponentsSpecDeploymentHpaSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Resources.Requests)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
								},
								Strategy: StrategyDeploymentSpecComponentsRateLimitServerBackendKubeSpecResourcesModel{
									RollingUpdate: RollingUpdateStrategyDeploymentSpecComponentsRateLimitServerBackendKubeSpecMaxUnavailableModel{
										MaxSurge: MaxSurgeRollingUpdateStrategyDeploymentSpecComponentsRateLimitServerBackendKubeSpecStrValModel{
											IntVal: IntValMaxSurgeRollingUpdateStrategyDeploymentSpecComponentsRateLimitServerBackendKubeSpecResourcesModel{Value: types.Int64Value(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
											StrVal: StrValKubeSpecBackendRateLimitServerComponentsSpecDeploymentStrategyRollingUpdateMaxSurgeIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
										},
										MaxUnavailable: MaxUnavailableKubeSpecBackendRateLimitServerComponentsSpecDeploymentStrategyRollingUpdateStrValModel{
											IntVal: IntValMaxUnavailableKubeSpecBackendRateLimitServerComponentsSpecDeploymentStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
											StrVal: StrValRollingUpdateStrategyDeploymentSpecComponentsRateLimitServerBackendKubeSpecMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
										},
									},
									Type: types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.Type)]),
								},
								Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
							},
							Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
							Service: ServiceDeploymentSpecComponentsRateLimitServerBackendLabelsModel{
								Annotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, AnnotationsServiceDeploymentSpecComponentsRateLimitServerBackendKubeSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Service.Annotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Labels: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, LabelsBackendRateLimitServerComponentsSpecDeploymentServiceAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Service.Labels)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
								Type:  types.StringValue(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.RateLimitServer.KubeSpec.Service.Type)]),
							},
							ServiceAccount: ServiceAccountKubeSpecBackendRateLimitServerComponentsSpecHelmModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
						},
					},
					Satellite: SpecComponentsSatelliteKubeSpecDeploymentModel{
						Enabled: "primitive / BoolKind",
						KubeSpec: DeploymentKubeSpecSatelliteComponentsSpecServiceAccountModel{
							Deployment: HelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel{
								Affinity: AffinityDeploymentKubeSpecSatelliteComponentsSpecHelmPodSecurityContextModel{
									NodeAffinity: HelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityDeploymentKubeSpecSatelliteComponentsSpecHelmPodSecurityContextModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
									},
									PodAffinity: PodAffinityAffinityDeploymentKubeSpecSatelliteComponentsSpecHelmPodSecurityContextModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
									PodAntiAffinity: PodAntiAffinityAffinityDeploymentKubeSpecSatelliteComponentsSpecHelmPodSecurityContextModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
								},
								ContainerSecurityContext: ContainerSecurityContextDeploymentKubeSpecSatelliteComponentsSpecHelmStrategyModel{
									AllowPrivilegeEscalation: "primitive* / BoolKind",
									Capabilities: CapabilitiesContainerSecurityContextDeploymentKubeSpecSatelliteComponentsSpecHelmDropModel{
										Add: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, AddCapabilitiesContainerSecurityContextDeploymentKubeSpecSatelliteComponentsSpecHelmStrategyModel.ElementType(ctx), cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
											resp.Diagnostics.Append(diag...)
										}(),
										Drop: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, DropHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
											resp.Diagnostics.Append(diag...)
										}(),
									},
									Privileged:             "primitive* / BoolKind",
									ProcMount:              types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
									ReadOnlyRootFilesystem: "primitive* / BoolKind",
									RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
									RunAsNonRoot:           "primitive* / BoolKind",
									RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsContainerSecurityContextDeploymentKubeSpecSatelliteComponentsSpecHelmStrategyModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfileContainerSecurityContextDeploymentKubeSpecSatelliteComponentsSpecHelmStrategyModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
									},
									WindowsOptions: WindowsOptionsContainerSecurityContextDeploymentKubeSpecSatelliteComponentsSpecHelmStrategyModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
								HpaSpec: HpaSpecHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityModel{
									MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.HpaSpec.MaxReplicas),
									Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
									MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.HpaSpec.MinReplicas),
								},
								PodAnnotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, PodAnnotationsDeploymentKubeSpecSatelliteComponentsSpecHelmInstallTemplateModel.ElementType(ctx), cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodAnnotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								PodSecurityContext: PodSecurityContextHelmSpecComponentsSatelliteKubeSpecDeploymentPodAnnotationsModel{
									FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.FsGroup),
									FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
									RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
									RunAsNonRoot:        "primitive* / BoolKind",
									RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsPodSecurityContextHelmSpecComponentsSatelliteKubeSpecDeploymentPodAnnotationsModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfilePodSecurityContextHelmSpecComponentsSatelliteKubeSpecDeploymentPodAnnotationsModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
									},
									SupplementalGroups: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, SupplementalGroupsPodSecurityContextHelmSpecComponentsSatelliteKubeSpecDeploymentPodAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
										resp.Diagnostics.Append(diag...)
									}(),
									Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
									WindowsOptions: WindowsOptionsPodSecurityContextHelmSpecComponentsSatelliteKubeSpecDeploymentPodAnnotationsModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ReplicaCount),
								Resources: ResourcesDeploymentKubeSpecSatelliteComponentsSpecHelmHpaSpecModel{
									Limits: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, LimitsResourcesDeploymentKubeSpecSatelliteComponentsSpecHelmHpaSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Resources.Limits)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Requests: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, RequestsResourcesDeploymentKubeSpecSatelliteComponentsSpecHelmHpaSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Resources.Requests)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
								},
								Strategy: StrategyHelmSpecComponentsSatelliteKubeSpecDeploymentResourcesModel{
									RollingUpdate: RollingUpdateStrategyHelmSpecComponentsSatelliteKubeSpecDeploymentMaxUnavailableModel{
										MaxSurge: MaxSurgeRollingUpdateStrategyHelmSpecComponentsSatelliteKubeSpecDeploymentIntValModel{
											IntVal: IntValDeploymentKubeSpecSatelliteComponentsSpecHelmStrategyRollingUpdateMaxSurgeStrValModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
											StrVal: StrValMaxSurgeRollingUpdateStrategyHelmSpecComponentsSatelliteKubeSpecDeploymentResourcesModel{Value: types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
										},
										MaxUnavailable: MaxUnavailableDeploymentKubeSpecSatelliteComponentsSpecHelmStrategyRollingUpdateStrValModel{
											IntVal: IntValMaxUnavailableDeploymentKubeSpecSatelliteComponentsSpecHelmStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
											StrVal: StrValRollingUpdateStrategyHelmSpecComponentsSatelliteKubeSpecDeploymentMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
										},
									},
									Type: types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.Type)]),
								},
								Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
							},
							Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
							Service: ServiceDeploymentKubeSpecSatelliteComponentsSpecLabelsModel{
								Annotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, AnnotationsServiceDeploymentKubeSpecSatelliteComponentsSpecHelmModel.ElementType(ctx), cluster.Helm.Spec.Components.Satellite.KubeSpec.Service.Annotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Labels: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, LabelsSpecComponentsSatelliteKubeSpecDeploymentServiceAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.Satellite.KubeSpec.Service.Labels)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
								Type:  types.StringValue(cluster.Helm.Spec.Components.Satellite.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.Satellite.KubeSpec.Service.Type)]),
							},
							ServiceAccount: ServiceAccountSpecComponentsSatelliteKubeSpecDeploymentServiceModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
						},
					},
					Xcp: ConfigProtectionSpecComponentsXcpKubeSpecModel{
						CentralAuthMode:       types.StringValue(cluster.Helm.Spec.Components.Xcp.CentralAuthMode[int32(cluster.Helm.Spec.Components.Xcp.CentralAuthMode)]),
						CentralProvidedCaCert: "primitive / BoolKind",
						ConfigProtection: KubeSpecXcpComponentsSpecConfigProtectionAuthorizedUsersModel{
							AuthorizedUsers: func() basetypes.ListValue {
								r, diag := types.ListValue(ctx, AuthorizedUsersConfigProtectionSpecComponentsXcpKubeSpecDeploymentModel.ElementType(ctx), cluster.Helm.Spec.Components.Xcp.ConfigProtection.AuthorizedUsers)
								resp.Diagnostics.Append(diag...)
							}(),
							EnableAuthorizedCreateUpdateDeleteOnXcpConfigs: "primitive / BoolKind",
							EnableAuthorizedUpdateDeleteOnXcpConfigs:       "primitive / BoolKind",
						},
						EnableHttpMeshInternalIdentityPropagation: "primitive / BoolKind",
						IsolationBoundaries:                       "github.com/tetrateio/api/install/controlplane/v1alpha1 / IsolationBoundary",
						KubeSpec: DeploymentKubeSpecXcpComponentsSpecServiceAccountModel{
							Deployment: HelmSpecComponentsXcpKubeSpecDeploymentPodAnnotationsModel{
								Affinity: AffinityHelmSpecComponentsXcpKubeSpecDeploymentStrategyModel{
									NodeAffinity: DeploymentKubeSpecXcpComponentsSpecHelmAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / PreferredSchedulingTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  RequiredDuringSchedulingIgnoredDuringExecutionNodeAffinityAffinityHelmSpecComponentsXcpKubeSpecDeploymentStrategyModel{NodeSelectorTerms: "github.com/tetrateio/api/install/kubernetes / NodeSelectorTerm"},
									},
									PodAffinity: PodAffinityAffinityHelmSpecComponentsXcpKubeSpecDeploymentStrategyModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
									PodAntiAffinity: PodAntiAffinityAffinityHelmSpecComponentsXcpKubeSpecDeploymentStrategyModel{
										PreferredDuringSchedulingIgnoredDuringExecution: "github.com/tetrateio/api/install/kubernetes / WeightedPodAffinityTerm",
										RequiredDuringSchedulingIgnoredDuringExecution:  "github.com/tetrateio/api/install/kubernetes / PodAffinityTerm",
									},
								},
								ContainerSecurityContext: ContainerSecurityContextHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecModel{
									AllowPrivilegeEscalation: "primitive* / BoolKind",
									Capabilities: CapabilitiesContainerSecurityContextHelmSpecComponentsXcpKubeSpecDeploymentDropModel{
										Add: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, AddCapabilitiesContainerSecurityContextHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecModel.ElementType(ctx), cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add)
											resp.Diagnostics.Append(diag...)
										}(),
										Drop: func() basetypes.ListValue {
											r, diag := types.ListValue(ctx, DropDeploymentKubeSpecXcpComponentsSpecHelmContainerSecurityContextCapabilitiesAddModel.ElementType(ctx), cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop)
											resp.Diagnostics.Append(diag...)
										}(),
									},
									Privileged:             "primitive* / BoolKind",
									ProcMount:              types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.ProcMount[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.ProcMount)]),
									ReadOnlyRootFilesystem: "primitive* / BoolKind",
									RunAsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup),
									RunAsNonRoot:           "primitive* / BoolKind",
									RunAsUser:              types.Int64Value(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsContainerSecurityContextHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfileContainerSecurityContextHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type)]),
									},
									WindowsOptions: WindowsOptionsContainerSecurityContextHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								Env: "github.com/tetrateio/api/install/kubernetes / EnvVar",
								HpaSpec: HpaSpecDeploymentKubeSpecXcpComponentsSpecHelmAffinityModel{
									MaxReplicas: types.Int64Value(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.HpaSpec.MaxReplicas),
									Metrics:     "github.com/tetrateio/api/install/kubernetes / MetricSpec",
									MinReplicas: types.Int64Value(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.HpaSpec.MinReplicas),
								},
								PodAnnotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, PodAnnotationsDeploymentKubeSpecXcpComponentsSpecHelmContainerSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodAnnotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								PodSecurityContext: PodSecurityContextDeploymentKubeSpecXcpComponentsSpecHelmInstallTemplateModel{
									FsGroup:             types.Int64Value(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.FsGroup),
									FsGroupChangePolicy: types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy)]),
									RunAsGroup:          types.Int64Value(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.RunAsGroup),
									RunAsNonRoot:        "primitive* / BoolKind",
									RunAsUser:           types.Int64Value(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.RunAsUser),
									SeLinuxOptions: SeLinuxOptionsPodSecurityContextDeploymentKubeSpecXcpComponentsSpecHelmInstallTemplateModel{
										Level: types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level)]),
										Role:  types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role)]),
										Type:  types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type)]),
										User:  types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User)]),
									},
									SeccompProfile: SeccompProfilePodSecurityContextDeploymentKubeSpecXcpComponentsSpecHelmInstallTemplateModel{
										LocalhostProfile: types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile)]),
										Type:             types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type)]),
									},
									SupplementalGroups: func() basetypes.ListValue {
										r, diag := types.ListValue(ctx, SupplementalGroupsPodSecurityContextDeploymentKubeSpecXcpComponentsSpecHelmInstallTemplateModel.ElementType(ctx), cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups)
										resp.Diagnostics.Append(diag...)
									}(),
									Sysctls: "github.com/tetrateio/api/install/kubernetes / Sysctl",
									WindowsOptions: WindowsOptionsPodSecurityContextDeploymentKubeSpecXcpComponentsSpecHelmInstallTemplateModel{
										GmsaCredentialSpec:     types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec)]),
										GmsaCredentialSpecName: types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName)]),
										RunAsUserName:          types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName)]),
									},
								},
								ReplicaCount: types.Int64Value(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ReplicaCount),
								Resources: ResourcesHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel{
									Limits: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, LimitsResourcesHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Resources.Limits)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
									Requests: func() basetypes.MapValue {
										r, diag := types.MapValue(ctx, RequestsResourcesHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel.ElementType(ctx), cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Resources.Requests)
										resp.Diagnostics.Append(diag...)
										return r
									}(),
								},
								Strategy: StrategyDeploymentKubeSpecXcpComponentsSpecHelmResourcesModel{
									RollingUpdate: RollingUpdateStrategyDeploymentKubeSpecXcpComponentsSpecHelmMaxUnavailableModel{
										MaxSurge: MaxSurgeRollingUpdateStrategyDeploymentKubeSpecXcpComponentsSpecHelmIntValModel{
											IntVal: IntValHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxSurgeStrValModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value)},
											StrVal: StrValMaxSurgeRollingUpdateStrategyDeploymentKubeSpecXcpComponentsSpecHelmResourcesModel{Value: types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type),
										},
										MaxUnavailable: MaxUnavailableHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateStrValModel{
											IntVal: IntValMaxUnavailableHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel{Value: types.Int64Value(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value)},
											StrVal: StrValRollingUpdateStrategyDeploymentKubeSpecXcpComponentsSpecHelmMaxUnavailableIntValModel{Value: types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value)])},
											Type:   types.Int64Value(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type),
										},
									},
									Type: types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.Type[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.Type)]),
								},
								Tolerations: "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1 / Toleration",
							},
							Overlays: "istio.io/api/operator/v1alpha1 / K8SObjectOverlay",
							Service: ServiceDeploymentKubeSpecXcpComponentsSpecLabelsModel{
								Annotations: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, AnnotationsServiceDeploymentKubeSpecXcpComponentsSpecHelmModel.ElementType(ctx), cluster.Helm.Spec.Components.Xcp.KubeSpec.Service.Annotations)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Labels: func() basetypes.MapValue {
									r, diag := types.MapValue(ctx, LabelsSpecComponentsXcpKubeSpecDeploymentServiceAnnotationsModel.ElementType(ctx), cluster.Helm.Spec.Components.Xcp.KubeSpec.Service.Labels)
									resp.Diagnostics.Append(diag...)
									return r
								}(),
								Ports: "github.com/tetrateio/api/install/kubernetes / ServicePort",
								Type:  types.StringValue(cluster.Helm.Spec.Components.Xcp.KubeSpec.Service.Type[int32(cluster.Helm.Spec.Components.Xcp.KubeSpec.Service.Type)]),
							},
							ServiceAccount: ServiceAccountSpecComponentsXcpKubeSpecDeploymentServiceModel{ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference"},
						},
						Revision: types.StringValue(cluster.Helm.Spec.Components.Xcp.Revision[int32(cluster.Helm.Spec.Components.Xcp.Revision)]),
					},
				},
				Hub:              types.StringValue(cluster.Helm.Spec.Hub[int32(cluster.Helm.Spec.Hub)]),
				ImagePullSecrets: "github.com/tetrateio/api/install/kubernetes / LocalObjectReference",
				ManagementPlane: ManagementPlaneHelmSpecComponentsModel{
					ClusterName: types.StringValue(cluster.Helm.Spec.ManagementPlane.ClusterName[int32(cluster.Helm.Spec.ManagementPlane.ClusterName)]),
					Host:        types.StringValue(cluster.Helm.Spec.ManagementPlane.Host[int32(cluster.Helm.Spec.ManagementPlane.Host)]),
					Port:        types.Int64Value(cluster.Helm.Spec.ManagementPlane.Port),
					SelfSigned:  "primitive / BoolKind",
				},
				MeshExpansion: MeshExpansionSpecHelmManagementPlaneModel{
					CustomGateway: CustomGatewayMeshExpansionSpecHelmManagementPlaneModel{
						Host: types.StringValue(cluster.Helm.Spec.MeshExpansion.CustomGateway.Host[int32(cluster.Helm.Spec.MeshExpansion.CustomGateway.Host)]),
						Port: types.Int64Value(cluster.Helm.Spec.MeshExpansion.CustomGateway.Port),
					},
					Onboarding: MeshExpansionOnboardingWorkloadsAuthenticationDeregistrationModel{
						Endpoint: HelmSpecMeshExpansionOnboardingEndpointHostsModel{
							Hosts: func() basetypes.ListValue {
								r, diag := types.ListValue(ctx, HostsEndpointOnboardingMeshExpansionSpecHelmManagementPlaneModel.ElementType(ctx), cluster.Helm.Spec.MeshExpansion.Onboarding.Endpoint.Hosts)
								resp.Diagnostics.Append(diag...)
							}(),
							SecretName: types.StringValue(cluster.Helm.Spec.MeshExpansion.Onboarding.Endpoint.SecretName[int32(cluster.Helm.Spec.MeshExpansion.Onboarding.Endpoint.SecretName)]),
						},
						LocalRepository: LocalRepositoryEndpointOnboardingMeshExpansionSpecHelmModel{},
						TokenIssuer: JwtTokenIssuerSpecMeshExpansionOnboardingEndpointModel{Jwt: EndpointOnboardingMeshExpansionSpecTokenIssuerJwtExpirationModel{Expiration: ExpirationJwtTokenIssuerSpecMeshExpansionOnboardingEndpointLocalRepositoryModel{
							Nanos:   types.Int64Value(cluster.Helm.Spec.MeshExpansion.Onboarding.TokenIssuer.Jwt.Expiration.Nanos),
							Seconds: types.Int64Value(cluster.Helm.Spec.MeshExpansion.Onboarding.TokenIssuer.Jwt.Expiration.Seconds),
						}}},
						Uid: types.StringValue(cluster.Helm.Spec.MeshExpansion.Onboarding.Uid[int32(cluster.Helm.Spec.MeshExpansion.Onboarding.Uid)]),
						Workloads: DeregistrationAuthenticationWorkloadsOnboardingMeshExpansionSpecModel{
							Authentication: TokenIssuerSpecMeshExpansionOnboardingWorkloadsAuthenticationJwtModel{Jwt: JwtAuthenticationWorkloadsOnboardingMeshExpansionSpecTokenIssuerJwtModel{Issuers: "github.com/tetrateio/api/onboarding/config/install/v1alpha1 / JwtIssuer"}},
							Deregistration: SpecMeshExpansionOnboardingWorkloadsAuthenticationDeregistrationPropagationDelayModel{PropagationDelay: PropagationDelayDeregistrationAuthenticationWorkloadsOnboardingMeshExpansionSpecTokenIssuerModel{
								Nanos:   types.Int64Value(cluster.Helm.Spec.MeshExpansion.Onboarding.Workloads.Deregistration.PropagationDelay.Nanos),
								Seconds: types.Int64Value(cluster.Helm.Spec.MeshExpansion.Onboarding.Workloads.Deregistration.PropagationDelay.Seconds),
							}},
						},
					},
				},
				MeshObservability: MeshObservabilityHelmSpecMeshExpansionModel{
					DemoSettings: DemoSettingsMeshObservabilityHelmSpecMeshExpansionModel{ApiEndpointMetricsEnabled: "primitive / BoolKind"},
					Settings:     SettingsMeshObservabilityHelmSpecMeshExpansionModel{ApiEndpointMetricsEnabled: "primitive / BoolKind"},
				},
				TelemetryStore: TelemetryStoreSpecHelmMeshObservabilityModel{Elastic: ElasticTelemetryStoreSpecHelmMeshObservabilityModel{
					Host:       types.StringValue(cluster.Helm.Spec.TelemetryStore.Elastic.Host[int32(cluster.Helm.Spec.TelemetryStore.Elastic.Host)]),
					Port:       types.Int64Value(cluster.Helm.Spec.TelemetryStore.Elastic.Port),
					Protocol:   types.StringValue(cluster.Helm.Spec.TelemetryStore.Elastic.Protocol[int32(cluster.Helm.Spec.TelemetryStore.Elastic.Protocol)]),
					SelfSigned: "primitive / BoolKind",
					Version:    types.Int64Value(cluster.Helm.Spec.TelemetryStore.Elastic.Version),
				}},
				Tier1Cluster: "primitive / BoolKind",
			},
		},
		Message: types.StringValue(cluster.Message[int32(cluster.Message)]),
	}
	model.DisplayName = types.StringValue(cluster[int32(cluster)])
	model.TrustDomain = types.StringValue(cluster[int32(cluster)])
	model.Description = types.StringValue(cluster[int32(cluster)])
	model.NamespaceScope = NamespaceScopeModel{
		Exceptions: func() basetypes.ListValue {
			r, diag := types.ListValue(ctx, ExceptionsNamespaceScopeModel.ElementType(ctx), cluster.Exceptions)
			resp.Diagnostics.Append(diag...)
		}(),
		Scope: types.StringValue(cluster.Scope[int32(cluster.Scope)]),
	}
	model.Namespaces = "github.com/tetrateio/api/tsb/v2 / Namespace"
	model.Locality = LocalityModel{Region: types.StringValue(cluster.Region[int32(cluster.Region)])}
	model.Tier1Cluster = Tier1ClusterModel{Value: "primitive / BoolKind"}
	model.Labels = func() basetypes.MapValue {
		r, diag := types.MapValue(ctx, LabelsModel.ElementType(ctx), cluster)
		resp.Diagnostics.Append(diag...)
		return r
	}()
	model.ServiceAccount = ServiceAccountModel{
		Description: types.StringValue(cluster.Description[int32(cluster.Description)]),
		DisplayName: types.StringValue(cluster.DisplayName[int32(cluster.DisplayName)]),
		Keys:        "github.com/tetrateio/api/tsb/v2 / ServiceAccount_KeyPair",
	}
	model.State = StateModel{
		IstioVersions: func() basetypes.ListValue {
			r, diag := types.ListValue(ctx, IstioVersionsStateModel.ElementType(ctx), cluster.IstioVersions)
			resp.Diagnostics.Append(diag...)
		}(),
		LastSyncTime: LastSyncTimeStateModel{
			Nanos:   types.Int64Value(cluster.LastSyncTime.Nanos),
			Seconds: types.Int64Value(cluster.LastSyncTime.Seconds),
		},
		Provider:     types.StringValue(cluster.Provider[int32(cluster.Provider)]),
		TsbCpVersion: types.StringValue(cluster.TsbCpVersion[int32(cluster.TsbCpVersion)]),
		XcpVersion:   types.StringValue(cluster.XcpVersion[int32(cluster.XcpVersion)]),
	}
	model.Network = types.StringValue(cluster[int32(cluster)])
	model.TokenTtl = Token_TTLModel{
		Nanos:   types.Int64Value(cluster.Nanos),
		Seconds: types.Int64Value(cluster.Seconds),
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
