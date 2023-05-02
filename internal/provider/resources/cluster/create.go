package cluster

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	v11 "github.com/tetrateio/api/common-protos/k8s.io/api/core/v1"
	v1 "github.com/tetrateio/api/common-protos/k8s.io/apimachinery/pkg/apis/meta/v1"
	common "github.com/tetrateio/api/install/common"
	v1alpha13 "github.com/tetrateio/api/install/controlplane/v1alpha1"
	v1alpha11 "github.com/tetrateio/api/install/helm/common/v1alpha1"
	v1alpha1 "github.com/tetrateio/api/install/helm/controlplane/v1alpha1"
	kubernetes "github.com/tetrateio/api/install/kubernetes"
	v1alpha14 "github.com/tetrateio/api/onboarding/config/install/v1alpha1"
	v2 "github.com/tetrateio/api/tsb/v2"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	v1alpha12 "istio.io/api/operator/v1alpha1"
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
			InstallTemplate: &v2.Cluster_InstallTemplate{
				Helm: &v1alpha1.Values{
					Image: &v1alpha11.Image{
						Registry: model.InstallTemplate.Helm.Image.Registry.ValueString(),
						Tag:      model.InstallTemplate.Helm.Image.Tag.ValueString(),
					},
					Operator: &v1alpha11.Operator{
						Deployment: &v1alpha11.Operator_Deployment{
							Affinity: &kubernetes.Affinity{
								NodeAffinity: &kubernetes.NodeAffinity{
									PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
										a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Operator.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
										for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Operator.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
											a[i] = &kubernetes.PreferredSchedulingTerm{
												Preference: &kubernetes.NodeSelectorTerm{
													MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
														a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
														for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
															a[i] = &kubernetes.NodeSelectorRequirement{
																Key:      match_expressions.Key.ValueString(),
																Operator: match_expressions.Operator.ValueString(),
																Values: func() []string {
																	tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																	resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															}
														}
														return a
													}(),
													MatchFields: func() []*kubernetes.NodeSelectorRequirement {
														a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
														for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
															a[i] = &kubernetes.NodeSelectorRequirement{
																Key:      match_fields.Key.ValueString(),
																Operator: match_fields.Operator.ValueString(),
																Values: func() []string {
																	tmp := make([]string, 0, len(match_fields.Values.Elements()))
																	resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															}
														}
														return a
													}(),
												},
												Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
											}
										}
										return a
									}(),
									RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
										a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Operator.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
										for i, node_selector_terms := range model.InstallTemplate.Helm.Operator.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
											a[i] = &kubernetes.NodeSelectorTerm{
												MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
													a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
													for i, match_expressions := range node_selector_terms.MatchExpressions {
														a[i] = &kubernetes.NodeSelectorRequirement{
															Key:      match_expressions.Key.ValueString(),
															Operator: match_expressions.Operator.ValueString(),
															Values: func() []string {
																tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														}
													}
													return a
												}(),
												MatchFields: func() []*kubernetes.NodeSelectorRequirement {
													a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
													for i, match_fields := range node_selector_terms.MatchFields {
														a[i] = &kubernetes.NodeSelectorRequirement{
															Key:      match_fields.Key.ValueString(),
															Operator: match_fields.Operator.ValueString(),
															Values: func() []string {
																tmp := make([]string, 0, len(match_fields.Values.Elements()))
																resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														}
													}
													return a
												}(),
											}
										}
										return a
									}()},
								},
								PodAffinity: &kubernetes.PodAffinity{
									PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
										a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Operator.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
										for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Operator.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
											a[i] = &kubernetes.WeightedPodAffinityTerm{
												PodAffinityTerm: &kubernetes.PodAffinityTerm{
													LabelSelector: &v1.LabelSelector{
														MatchExpressions: func() []*v1.LabelSelectorRequirement {
															a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
															for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																a[i] = &v1.LabelSelectorRequirement{
																	Key:      ptrify(match_expressions.Key.ValueString()),
																	Operator: ptrify(match_expressions.Operator.ValueString()),
																	Values: func() []string {
																		tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																		resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																}
															}
															return a
														}(),
														MatchLabels: func() map[string]string {
															tmp := make(map[string]string)
															resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
															return tmp
														}(),
													},
													Namespaces: func() []string {
														tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
														resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
													TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
												},
												Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
											}
										}
										return a
									}(),
									RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
										a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Operator.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
										for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Operator.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
											a[i] = &kubernetes.PodAffinityTerm{
												LabelSelector: &v1.LabelSelector{
													MatchExpressions: func() []*v1.LabelSelectorRequirement {
														a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
														for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
															a[i] = &v1.LabelSelectorRequirement{
																Key:      ptrify(match_expressions.Key.ValueString()),
																Operator: ptrify(match_expressions.Operator.ValueString()),
																Values: func() []string {
																	tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																	resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															}
														}
														return a
													}(),
													MatchLabels: func() map[string]string {
														tmp := make(map[string]string)
														resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
												},
												Namespaces: func() []string {
													tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
													resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
											}
										}
										return a
									}(),
								},
								PodAntiAffinity: &kubernetes.PodAntiAffinity{
									PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
										a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Operator.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
										for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Operator.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
											a[i] = &kubernetes.WeightedPodAffinityTerm{
												PodAffinityTerm: &kubernetes.PodAffinityTerm{
													LabelSelector: &v1.LabelSelector{
														MatchExpressions: func() []*v1.LabelSelectorRequirement {
															a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
															for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																a[i] = &v1.LabelSelectorRequirement{
																	Key:      ptrify(match_expressions.Key.ValueString()),
																	Operator: ptrify(match_expressions.Operator.ValueString()),
																	Values: func() []string {
																		tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																		resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																}
															}
															return a
														}(),
														MatchLabels: func() map[string]string {
															tmp := make(map[string]string)
															resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
															return tmp
														}(),
													},
													Namespaces: func() []string {
														tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
														resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
													TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
												},
												Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
											}
										}
										return a
									}(),
									RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
										a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Operator.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
										for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Operator.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
											a[i] = &kubernetes.PodAffinityTerm{
												LabelSelector: &v1.LabelSelector{
													MatchExpressions: func() []*v1.LabelSelectorRequirement {
														a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
														for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
															a[i] = &v1.LabelSelectorRequirement{
																Key:      ptrify(match_expressions.Key.ValueString()),
																Operator: ptrify(match_expressions.Operator.ValueString()),
																Values: func() []string {
																	tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																	resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															}
														}
														return a
													}(),
													MatchLabels: func() map[string]string {
														tmp := make(map[string]string)
														resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
												},
												Namespaces: func() []string {
													tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
													resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
											}
										}
										return a
									}(),
								},
							},
							Annotations: func() map[string]string {
								tmp := make(map[string]string)
								resp.Diagnostics.Append(model.InstallTemplate.Helm.Operator.Deployment.Annotations.ElementsAs(ctx, &tmp, false)...)
								return tmp
							}(),
							Env: func() []*kubernetes.EnvVar {
								a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Operator.Deployment.Env))
								for i, env := range model.InstallTemplate.Helm.Operator.Deployment.Env {
									a[i] = &kubernetes.EnvVar{
										Name:  env.Name.ValueString(),
										Value: env.Value.ValueString(),
										ValueFrom: &kubernetes.EnvVarSource{
											ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
												Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
												LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
												Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
											},
											FieldRef: &kubernetes.ObjectFieldSelector{
												ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
												FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
											},
											ResourceFieldRef: &kubernetes.ResourceFieldSelector{
												ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
												Divisor: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
													Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
												},
												Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
											},
											SecretKeyRef: &kubernetes.SecretKeySelector{
												Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
												LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
												Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
											},
										},
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
							Strategy: &kubernetes.DeploymentStrategy{
								RollingUpdate: &kubernetes.RollingUpdateDeployment{
									MaxSurge: &v1alpha12.IntOrString{
										IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
										StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
										Type:   model.InstallTemplate.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
									},
									MaxUnavailable: &v1alpha12.IntOrString{
										IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
										StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
										Type:   model.InstallTemplate.Helm.Operator.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
									},
								},
								Type: model.InstallTemplate.Helm.Operator.Deployment.Strategy.Type.ValueString(),
							},
							Tolerations: func() []*v11.Toleration {
								a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Operator.Deployment.Tolerations))
								for i, tolerations := range model.InstallTemplate.Helm.Operator.Deployment.Tolerations {
									a[i] = &v11.Toleration{
										Effect:            ptrify(tolerations.Effect.ValueString()),
										Key:               ptrify(tolerations.Key.ValueString()),
										Operator:          ptrify(tolerations.Operator.ValueString()),
										TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
										Value:             ptrify(tolerations.Value.ValueString()),
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
							ImagePullSecrets: func() []string {
								tmp := make([]string, 0, len(model.InstallTemplate.Helm.Operator.ServiceAccount.ImagePullSecrets.Elements()))
								resp.Diagnostics.Append(model.InstallTemplate.Helm.Operator.ServiceAccount.ImagePullSecrets.ElementsAs(ctx, &tmp, false)...)
								return tmp
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
					Spec: &v1alpha13.ControlPlaneSpec{
						Components: &v1alpha13.ControlPlaneComponentSet{
							Collector: &v1alpha13.OpenTelemetryCollector{KubeSpec: &kubernetes.KubernetesComponentSpec{
								Deployment: &kubernetes.Deployment{
									Affinity: &kubernetes.Affinity{
										NodeAffinity: &kubernetes.NodeAffinity{
											PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
												a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
												for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.PreferredSchedulingTerm{
														Preference: &kubernetes.NodeSelectorTerm{
															MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_expressions.Key.ValueString(),
																		Operator: match_expressions.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_fields.Key.ValueString(),
																		Operator: match_fields.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_fields.Values.Elements()))
																			resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
														},
														Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
													}
												}
												return a
											}(),
											RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
												a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
												for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
													a[i] = &kubernetes.NodeSelectorTerm{
														MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
															a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
															for i, match_expressions := range node_selector_terms.MatchExpressions {
																a[i] = &kubernetes.NodeSelectorRequirement{
																	Key:      match_expressions.Key.ValueString(),
																	Operator: match_expressions.Operator.ValueString(),
																	Values: func() []string {
																		tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																		resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																}
															}
															return a
														}(),
														MatchFields: func() []*kubernetes.NodeSelectorRequirement {
															a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
															for i, match_fields := range node_selector_terms.MatchFields {
																a[i] = &kubernetes.NodeSelectorRequirement{
																	Key:      match_fields.Key.ValueString(),
																	Operator: match_fields.Operator.ValueString(),
																	Values: func() []string {
																		tmp := make([]string, 0, len(match_fields.Values.Elements()))
																		resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																}
															}
															return a
														}(),
													}
												}
												return a
											}()},
										},
										PodAffinity: &kubernetes.PodAffinity{
											PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
												a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
												for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.WeightedPodAffinityTerm{
														PodAffinityTerm: &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
														},
														Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
													}
												}
												return a
											}(),
											RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
												a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
												for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.PodAffinityTerm{
														LabelSelector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														Namespaces: func() []string {
															tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
															resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
															return tmp
														}(),
														TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
													}
												}
												return a
											}(),
										},
										PodAntiAffinity: &kubernetes.PodAntiAffinity{
											PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
												a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
												for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.WeightedPodAffinityTerm{
														PodAffinityTerm: &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
														},
														Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
													}
												}
												return a
											}(),
											RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
												a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
												for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.PodAffinityTerm{
														LabelSelector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														Namespaces: func() []string {
															tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
															resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
															return tmp
														}(),
														TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
													}
												}
												return a
											}(),
										},
									},
									ContainerSecurityContext: &kubernetes.SecurityContext{
										AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
										Capabilities: &kubernetes.Capabilities{
											Add: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Drop: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
										},
										Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
										ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
										ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
										RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
										RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
										RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
										SeLinuxOptions: &kubernetes.SELinuxOptions{
											Level: model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
											Role:  model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
											Type:  model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
											User:  model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
										},
										SeccompProfile: &kubernetes.SeccompProfile{
											LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
											Type:             model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
										},
										WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
											GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
											GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
											RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
										},
									},
									Env: func() []*kubernetes.EnvVar {
										a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Env))
										for i, env := range model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Env {
											a[i] = &kubernetes.EnvVar{
												Name:  env.Name.ValueString(),
												Value: env.Value.ValueString(),
												ValueFrom: &kubernetes.EnvVarSource{
													ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
														Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
														LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
														Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
													},
													FieldRef: &kubernetes.ObjectFieldSelector{
														ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
														FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
													},
													ResourceFieldRef: &kubernetes.ResourceFieldSelector{
														ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
														Divisor: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
															Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
														},
														Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
													},
													SecretKeyRef: &kubernetes.SecretKeySelector{
														Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
														LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
														Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
													},
												},
											}
										}
										return a
									}(),
									HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
										MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
										Metrics: func() []*kubernetes.MetricSpec {
											a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.HpaSpec.Metrics))
											for i, metrics := range model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.HpaSpec.Metrics {
												a[i] = &kubernetes.MetricSpec{
													External: &kubernetes.ExternalMetricSource{
														MetricName: metrics.External.MetricName.ValueString(),
														MetricSelector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														TargetAverageValue: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
															Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
														},
														TargetValue: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
															Type:   metrics.External.TargetValue.Type.ValueInt64(),
														},
													},
													Object: &kubernetes.ObjectMetricSource{
														AverageValue: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
															Type:   metrics.Object.AverageValue.Type.ValueInt64(),
														},
														MetricName: metrics.Object.MetricName.ValueString(),
														Selector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														Target: &kubernetes.CrossVersionObjectReference{
															ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
															Kind:       metrics.Object.Target.Kind.ValueString(),
															Name:       metrics.Object.Target.Name.ValueString(),
														},
														TargetValue: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
															Type:   metrics.Object.TargetValue.Type.ValueInt64(),
														},
													},
													Pods: &kubernetes.PodsMetricSource{
														MetricName: metrics.Pods.MetricName.ValueString(),
														Selector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														TargetAverageValue: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
															Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
														},
													},
													Resource: &kubernetes.ResourceMetricSource{
														Name: metrics.Resource.Name.ValueString(),
														Target: &kubernetes.MetricTarget{
															AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
															AverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
															},
															Type: metrics.Resource.Target.Type.ValueString(),
															Value: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
															},
														},
														TargetAverageUtilization: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
															Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
														},
														TargetAverageValue: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
															Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
														},
													},
													Type: metrics.Type.ValueString(),
												}
											}
											return a
										}(),
										MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
									},
									PodAnnotations: func() map[string]string {
										tmp := make(map[string]string)
										resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									PodSecurityContext: &kubernetes.PodSecurityContext{
										FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
										FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
										RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
										RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
										RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
										SeLinuxOptions: &kubernetes.SELinuxOptions{
											Level: model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
											Role:  model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
											Type:  model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
											User:  model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
										},
										SeccompProfile: &kubernetes.SeccompProfile{
											LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
											Type:             model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
										},
										SupplementalGroups: func() []string {
											tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Sysctls: func() []*kubernetes.Sysctl {
											a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.Sysctls))
											for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.Sysctls {
												a[i] = &kubernetes.Sysctl{
													Name:  sysctls.Name.ValueString(),
													Value: sysctls.Value.ValueString(),
												}
											}
											return a
										}(),
										WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
											GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
											GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
											RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
										},
									},
									ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
									Resources: &kubernetes.Resources{
										Limits: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Requests: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
									},
									Strategy: &kubernetes.DeploymentStrategy{
										RollingUpdate: &kubernetes.RollingUpdateDeployment{
											MaxSurge: &v1alpha12.IntOrString{
												IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
												StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
												Type:   model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
											},
											MaxUnavailable: &v1alpha12.IntOrString{
												IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
												StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
												Type:   model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
											},
										},
										Type: model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Strategy.Type.ValueString(),
									},
									Tolerations: func() []*v11.Toleration {
										a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Tolerations))
										for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Deployment.Tolerations {
											a[i] = &v11.Toleration{
												Effect:            ptrify(tolerations.Effect.ValueString()),
												Key:               ptrify(tolerations.Key.ValueString()),
												Operator:          ptrify(tolerations.Operator.ValueString()),
												TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
												Value:             ptrify(tolerations.Value.ValueString()),
											}
										}
										return a
									}(),
								},
								Overlays: func() []*v1alpha12.K8SObjectOverlay {
									a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Overlays))
									for i, overlays := range model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Overlays {
										a[i] = &v1alpha12.K8SObjectOverlay{
											ApiVersion: overlays.ApiVersion.ValueString(),
											Kind:       overlays.Kind.ValueString(),
											Name:       overlays.Name.ValueString(),
											Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
												a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
												for i, patches := range overlays.Patches {
													a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
														Path: patches.Path.ValueString(),
														Value: &structpb.Value{
															BoolValue: patches.Value.BoolValue.ValueBool(),
															ListValue: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
															NumberValue: patches.Value.NumberValue.ValueFloat64(),
															StringValue: patches.Value.StringValue.ValueString(),
															StructValue: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
								Service: &kubernetes.Service{
									Annotations: func() map[string]string {
										tmp := make(map[string]string)
										resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									Labels: func() map[string]string {
										tmp := make(map[string]string)
										resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									Ports: func() []*kubernetes.ServicePort {
										a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Service.Ports))
										for i, ports := range model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Service.Ports {
											a[i] = &kubernetes.ServicePort{
												Name:     ports.Name.ValueString(),
												NodePort: int32(ports.NodePort.ValueInt64()),
												Port:     int32(ports.Port.ValueInt64()),
												Protocol: ports.Protocol.ValueString(),
												TargetPort: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
													Type:   ports.TargetPort.Type.ValueInt64(),
												},
											}
										}
										return a
									}(),
									Type: model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.Service.Type.ValueString(),
								},
								ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
									a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.ServiceAccount.ImagePullSecrets))
									for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.Collector.KubeSpec.ServiceAccount.ImagePullSecrets {
										a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
									}
									return a
								}()},
							}},
							DefaultKubeSpec: &kubernetes.KubernetesSpec{
								Account: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
									a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Account.ImagePullSecrets))
									for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Account.ImagePullSecrets {
										a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
									}
									return a
								}()},
								Deployment: &kubernetes.GlobalDeployment{
									Affinity: &kubernetes.Affinity{
										NodeAffinity: &kubernetes.NodeAffinity{
											PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
												a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
												for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.PreferredSchedulingTerm{
														Preference: &kubernetes.NodeSelectorTerm{
															MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_expressions.Key.ValueString(),
																		Operator: match_expressions.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_fields.Key.ValueString(),
																		Operator: match_fields.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_fields.Values.Elements()))
																			resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
														},
														Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
													}
												}
												return a
											}(),
											RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
												a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
												for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
													a[i] = &kubernetes.NodeSelectorTerm{
														MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
															a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
															for i, match_expressions := range node_selector_terms.MatchExpressions {
																a[i] = &kubernetes.NodeSelectorRequirement{
																	Key:      match_expressions.Key.ValueString(),
																	Operator: match_expressions.Operator.ValueString(),
																	Values: func() []string {
																		tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																		resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																}
															}
															return a
														}(),
														MatchFields: func() []*kubernetes.NodeSelectorRequirement {
															a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
															for i, match_fields := range node_selector_terms.MatchFields {
																a[i] = &kubernetes.NodeSelectorRequirement{
																	Key:      match_fields.Key.ValueString(),
																	Operator: match_fields.Operator.ValueString(),
																	Values: func() []string {
																		tmp := make([]string, 0, len(match_fields.Values.Elements()))
																		resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																}
															}
															return a
														}(),
													}
												}
												return a
											}()},
										},
										PodAffinity: &kubernetes.PodAffinity{
											PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
												a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
												for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.WeightedPodAffinityTerm{
														PodAffinityTerm: &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
														},
														Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
													}
												}
												return a
											}(),
											RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
												a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
												for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.PodAffinityTerm{
														LabelSelector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														Namespaces: func() []string {
															tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
															resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
															return tmp
														}(),
														TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
													}
												}
												return a
											}(),
										},
										PodAntiAffinity: &kubernetes.PodAntiAffinity{
											PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
												a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
												for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.WeightedPodAffinityTerm{
														PodAffinityTerm: &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
														},
														Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
													}
												}
												return a
											}(),
											RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
												a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
												for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.PodAffinityTerm{
														LabelSelector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														Namespaces: func() []string {
															tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
															resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
															return tmp
														}(),
														TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
													}
												}
												return a
											}(),
										},
									},
									ContainerSecurityContext: &kubernetes.SecurityContext{
										AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
										Capabilities: &kubernetes.Capabilities{
											Add: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Drop: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
										},
										Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
										ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
										ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
										RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
										RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
										RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
										SeLinuxOptions: &kubernetes.SELinuxOptions{
											Level: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
											Role:  model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
											Type:  model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
											User:  model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
										},
										SeccompProfile: &kubernetes.SeccompProfile{
											LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
											Type:             model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
										},
										WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
											GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
											GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
											RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
										},
									},
									Env: func() []*kubernetes.EnvVar {
										a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Env))
										for i, env := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Env {
											a[i] = &kubernetes.EnvVar{
												Name:  env.Name.ValueString(),
												Value: env.Value.ValueString(),
												ValueFrom: &kubernetes.EnvVarSource{
													ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
														Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
														LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
														Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
													},
													FieldRef: &kubernetes.ObjectFieldSelector{
														ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
														FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
													},
													ResourceFieldRef: &kubernetes.ResourceFieldSelector{
														ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
														Divisor: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
															Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
														},
														Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
													},
													SecretKeyRef: &kubernetes.SecretKeySelector{
														Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
														LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
														Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
													},
												},
											}
										}
										return a
									}(),
									PodAnnotations: func() map[string]string {
										tmp := make(map[string]string)
										resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									PodSecurityContext: &kubernetes.PodSecurityContext{
										FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
										FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
										RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
										RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
										RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
										SeLinuxOptions: &kubernetes.SELinuxOptions{
											Level: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
											Role:  model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
											Type:  model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
											User:  model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
										},
										SeccompProfile: &kubernetes.SeccompProfile{
											LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
											Type:             model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
										},
										SupplementalGroups: func() []string {
											tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Sysctls: func() []*kubernetes.Sysctl {
											a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.Sysctls))
											for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.Sysctls {
												a[i] = &kubernetes.Sysctl{
													Name:  sysctls.Name.ValueString(),
													Value: sysctls.Value.ValueString(),
												}
											}
											return a
										}(),
										WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
											GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
											GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
											RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
										},
									},
									Strategy: &kubernetes.DeploymentStrategy{
										RollingUpdate: &kubernetes.RollingUpdateDeployment{
											MaxSurge: &v1alpha12.IntOrString{
												IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
												StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
												Type:   model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
											},
											MaxUnavailable: &v1alpha12.IntOrString{
												IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
												StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
												Type:   model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
											},
										},
										Type: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Strategy.Type.ValueString(),
									},
									Tolerations: func() []*v11.Toleration {
										a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Tolerations))
										for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Deployment.Tolerations {
											a[i] = &v11.Toleration{
												Effect:            ptrify(tolerations.Effect.ValueString()),
												Key:               ptrify(tolerations.Key.ValueString()),
												Operator:          ptrify(tolerations.Operator.ValueString()),
												TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
												Value:             ptrify(tolerations.Value.ValueString()),
											}
										}
										return a
									}(),
								},
								Job: &kubernetes.GlobalJob{
									Affinity: &kubernetes.Affinity{
										NodeAffinity: &kubernetes.NodeAffinity{
											PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
												a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
												for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.PreferredSchedulingTerm{
														Preference: &kubernetes.NodeSelectorTerm{
															MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_expressions.Key.ValueString(),
																		Operator: match_expressions.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_fields.Key.ValueString(),
																		Operator: match_fields.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_fields.Values.Elements()))
																			resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
														},
														Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
													}
												}
												return a
											}(),
											RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
												a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
												for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
													a[i] = &kubernetes.NodeSelectorTerm{
														MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
															a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
															for i, match_expressions := range node_selector_terms.MatchExpressions {
																a[i] = &kubernetes.NodeSelectorRequirement{
																	Key:      match_expressions.Key.ValueString(),
																	Operator: match_expressions.Operator.ValueString(),
																	Values: func() []string {
																		tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																		resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																}
															}
															return a
														}(),
														MatchFields: func() []*kubernetes.NodeSelectorRequirement {
															a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
															for i, match_fields := range node_selector_terms.MatchFields {
																a[i] = &kubernetes.NodeSelectorRequirement{
																	Key:      match_fields.Key.ValueString(),
																	Operator: match_fields.Operator.ValueString(),
																	Values: func() []string {
																		tmp := make([]string, 0, len(match_fields.Values.Elements()))
																		resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																}
															}
															return a
														}(),
													}
												}
												return a
											}()},
										},
										PodAffinity: &kubernetes.PodAffinity{
											PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
												a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
												for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.WeightedPodAffinityTerm{
														PodAffinityTerm: &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
														},
														Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
													}
												}
												return a
											}(),
											RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
												a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
												for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.PodAffinityTerm{
														LabelSelector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														Namespaces: func() []string {
															tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
															resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
															return tmp
														}(),
														TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
													}
												}
												return a
											}(),
										},
										PodAntiAffinity: &kubernetes.PodAntiAffinity{
											PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
												a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
												for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.WeightedPodAffinityTerm{
														PodAffinityTerm: &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
														},
														Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
													}
												}
												return a
											}(),
											RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
												a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
												for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.PodAffinityTerm{
														LabelSelector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														Namespaces: func() []string {
															tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
															resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
															return tmp
														}(),
														TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
													}
												}
												return a
											}(),
										},
									},
									ContainerSecurityContext: &kubernetes.SecurityContext{
										AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
										Capabilities: &kubernetes.Capabilities{
											Add: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.Capabilities.Add.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Drop: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.Capabilities.Drop.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
										},
										Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.Privileged.ValueBool()),
										ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.ProcMount.ValueString()),
										ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
										RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.RunAsGroup.ValueInt64())),
										RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
										RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.RunAsUser.ValueInt64())),
										SeLinuxOptions: &kubernetes.SELinuxOptions{
											Level: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
											Role:  model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
											Type:  model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
											User:  model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
										},
										SeccompProfile: &kubernetes.SeccompProfile{
											LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
											Type:             model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
										},
										WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
											GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
											GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
											RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
										},
									},
									PodAnnotations: func() map[string]string {
										tmp := make(map[string]string)
										resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									PodSecurityContext: &kubernetes.PodSecurityContext{
										FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.FsGroup.ValueInt64())),
										FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.FsGroupChangePolicy.ValueString()),
										RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.RunAsGroup.ValueInt64())),
										RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.RunAsNonRoot.ValueBool()),
										RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.RunAsUser.ValueInt64())),
										SeLinuxOptions: &kubernetes.SELinuxOptions{
											Level: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
											Role:  model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
											Type:  model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
											User:  model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeLinuxOptions.User.ValueString(),
										},
										SeccompProfile: &kubernetes.SeccompProfile{
											LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
											Type:             model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SeccompProfile.Type.ValueString(),
										},
										SupplementalGroups: func() []string {
											tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SupplementalGroups.Elements()))
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Sysctls: func() []*kubernetes.Sysctl {
											a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.Sysctls))
											for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.Sysctls {
												a[i] = &kubernetes.Sysctl{
													Name:  sysctls.Name.ValueString(),
													Value: sysctls.Value.ValueString(),
												}
											}
											return a
										}(),
										WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
											GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
											GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
											RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
										},
									},
									Tolerations: func() []*v11.Toleration {
										a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Tolerations))
										for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Job.Tolerations {
											a[i] = &v11.Toleration{
												Effect:            ptrify(tolerations.Effect.ValueString()),
												Key:               ptrify(tolerations.Key.ValueString()),
												Operator:          ptrify(tolerations.Operator.ValueString()),
												TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
												Value:             ptrify(tolerations.Value.ValueString()),
											}
										}
										return a
									}(),
								},
								Service: &kubernetes.GlobalService{Annotations: func() map[string]string {
									tmp := make(map[string]string)
									resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.DefaultKubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
									return tmp
								}()},
							},
							Gitops: &common.GitOps{
								BatchWindow: &durationpb.Duration{
									Nanos:   int32(model.InstallTemplate.Helm.Spec.Components.Gitops.BatchWindow.Nanos.ValueInt64()),
									Seconds: model.InstallTemplate.Helm.Spec.Components.Gitops.BatchWindow.Seconds.ValueInt64(),
								},
								Enabled: model.InstallTemplate.Helm.Spec.Components.Gitops.Enabled.ValueBool(),
								ManagementplaneRequestTimeout: &durationpb.Duration{
									Nanos:   int32(model.InstallTemplate.Helm.Spec.Components.Gitops.ManagementplaneRequestTimeout.Nanos.ValueInt64()),
									Seconds: model.InstallTemplate.Helm.Spec.Components.Gitops.ManagementplaneRequestTimeout.Seconds.ValueInt64(),
								},
								ReconcileInterval: &durationpb.Duration{
									Nanos:   int32(model.InstallTemplate.Helm.Spec.Components.Gitops.ReconcileInterval.Nanos.ValueInt64()),
									Seconds: model.InstallTemplate.Helm.Spec.Components.Gitops.ReconcileInterval.Seconds.ValueInt64(),
								},
								ReconcileRequestTimeout: &durationpb.Duration{
									Nanos:   int32(model.InstallTemplate.Helm.Spec.Components.Gitops.ReconcileRequestTimeout.Nanos.ValueInt64()),
									Seconds: model.InstallTemplate.Helm.Spec.Components.Gitops.ReconcileRequestTimeout.Seconds.ValueInt64(),
								},
								WebhookTimeout: &durationpb.Duration{
									Nanos:   int32(model.InstallTemplate.Helm.Spec.Components.Gitops.WebhookTimeout.Nanos.ValueInt64()),
									Seconds: model.InstallTemplate.Helm.Spec.Components.Gitops.WebhookTimeout.Seconds.ValueInt64(),
								},
							},
							HpaAdapter: &v1alpha13.HpaAdapter{KubeSpec: &kubernetes.KubernetesComponentSpec{
								Deployment: &kubernetes.Deployment{
									Affinity: &kubernetes.Affinity{
										NodeAffinity: &kubernetes.NodeAffinity{
											PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
												a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
												for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.PreferredSchedulingTerm{
														Preference: &kubernetes.NodeSelectorTerm{
															MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_expressions.Key.ValueString(),
																		Operator: match_expressions.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_fields.Key.ValueString(),
																		Operator: match_fields.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_fields.Values.Elements()))
																			resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
														},
														Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
													}
												}
												return a
											}(),
											RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
												a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
												for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
													a[i] = &kubernetes.NodeSelectorTerm{
														MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
															a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
															for i, match_expressions := range node_selector_terms.MatchExpressions {
																a[i] = &kubernetes.NodeSelectorRequirement{
																	Key:      match_expressions.Key.ValueString(),
																	Operator: match_expressions.Operator.ValueString(),
																	Values: func() []string {
																		tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																		resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																}
															}
															return a
														}(),
														MatchFields: func() []*kubernetes.NodeSelectorRequirement {
															a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
															for i, match_fields := range node_selector_terms.MatchFields {
																a[i] = &kubernetes.NodeSelectorRequirement{
																	Key:      match_fields.Key.ValueString(),
																	Operator: match_fields.Operator.ValueString(),
																	Values: func() []string {
																		tmp := make([]string, 0, len(match_fields.Values.Elements()))
																		resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																}
															}
															return a
														}(),
													}
												}
												return a
											}()},
										},
										PodAffinity: &kubernetes.PodAffinity{
											PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
												a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
												for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.WeightedPodAffinityTerm{
														PodAffinityTerm: &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
														},
														Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
													}
												}
												return a
											}(),
											RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
												a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
												for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.PodAffinityTerm{
														LabelSelector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														Namespaces: func() []string {
															tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
															resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
															return tmp
														}(),
														TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
													}
												}
												return a
											}(),
										},
										PodAntiAffinity: &kubernetes.PodAntiAffinity{
											PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
												a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
												for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.WeightedPodAffinityTerm{
														PodAffinityTerm: &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
														},
														Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
													}
												}
												return a
											}(),
											RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
												a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
												for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
													a[i] = &kubernetes.PodAffinityTerm{
														LabelSelector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														Namespaces: func() []string {
															tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
															resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
															return tmp
														}(),
														TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
													}
												}
												return a
											}(),
										},
									},
									ContainerSecurityContext: &kubernetes.SecurityContext{
										AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
										Capabilities: &kubernetes.Capabilities{
											Add: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Drop: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
										},
										Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
										ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
										ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
										RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
										RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
										RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
										SeLinuxOptions: &kubernetes.SELinuxOptions{
											Level: model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
											Role:  model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
											Type:  model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
											User:  model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
										},
										SeccompProfile: &kubernetes.SeccompProfile{
											LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
											Type:             model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
										},
										WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
											GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
											GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
											RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
										},
									},
									Env: func() []*kubernetes.EnvVar {
										a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Env))
										for i, env := range model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Env {
											a[i] = &kubernetes.EnvVar{
												Name:  env.Name.ValueString(),
												Value: env.Value.ValueString(),
												ValueFrom: &kubernetes.EnvVarSource{
													ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
														Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
														LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
														Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
													},
													FieldRef: &kubernetes.ObjectFieldSelector{
														ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
														FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
													},
													ResourceFieldRef: &kubernetes.ResourceFieldSelector{
														ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
														Divisor: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
															Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
														},
														Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
													},
													SecretKeyRef: &kubernetes.SecretKeySelector{
														Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
														LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
														Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
													},
												},
											}
										}
										return a
									}(),
									HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
										MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
										Metrics: func() []*kubernetes.MetricSpec {
											a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.HpaSpec.Metrics))
											for i, metrics := range model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.HpaSpec.Metrics {
												a[i] = &kubernetes.MetricSpec{
													External: &kubernetes.ExternalMetricSource{
														MetricName: metrics.External.MetricName.ValueString(),
														MetricSelector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														TargetAverageValue: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
															Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
														},
														TargetValue: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
															Type:   metrics.External.TargetValue.Type.ValueInt64(),
														},
													},
													Object: &kubernetes.ObjectMetricSource{
														AverageValue: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
															Type:   metrics.Object.AverageValue.Type.ValueInt64(),
														},
														MetricName: metrics.Object.MetricName.ValueString(),
														Selector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														Target: &kubernetes.CrossVersionObjectReference{
															ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
															Kind:       metrics.Object.Target.Kind.ValueString(),
															Name:       metrics.Object.Target.Name.ValueString(),
														},
														TargetValue: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
															Type:   metrics.Object.TargetValue.Type.ValueInt64(),
														},
													},
													Pods: &kubernetes.PodsMetricSource{
														MetricName: metrics.Pods.MetricName.ValueString(),
														Selector: &v1.LabelSelector{
															MatchExpressions: func() []*v1.LabelSelectorRequirement {
																a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																	a[i] = &v1.LabelSelectorRequirement{
																		Key:      ptrify(match_expressions.Key.ValueString()),
																		Operator: ptrify(match_expressions.Operator.ValueString()),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchLabels: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
														},
														TargetAverageValue: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
															Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
														},
													},
													Resource: &kubernetes.ResourceMetricSource{
														Name: metrics.Resource.Name.ValueString(),
														Target: &kubernetes.MetricTarget{
															AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
															AverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
															},
															Type: metrics.Resource.Target.Type.ValueString(),
															Value: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
															},
														},
														TargetAverageUtilization: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
															Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
														},
														TargetAverageValue: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
															Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
														},
													},
													Type: metrics.Type.ValueString(),
												}
											}
											return a
										}(),
										MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
									},
									PodAnnotations: func() map[string]string {
										tmp := make(map[string]string)
										resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									PodSecurityContext: &kubernetes.PodSecurityContext{
										FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
										FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
										RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
										RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
										RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
										SeLinuxOptions: &kubernetes.SELinuxOptions{
											Level: model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
											Role:  model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
											Type:  model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
											User:  model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
										},
										SeccompProfile: &kubernetes.SeccompProfile{
											LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
											Type:             model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
										},
										SupplementalGroups: func() []string {
											tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Sysctls: func() []*kubernetes.Sysctl {
											a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.Sysctls))
											for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.Sysctls {
												a[i] = &kubernetes.Sysctl{
													Name:  sysctls.Name.ValueString(),
													Value: sysctls.Value.ValueString(),
												}
											}
											return a
										}(),
										WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
											GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
											GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
											RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
										},
									},
									ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
									Resources: &kubernetes.Resources{
										Limits: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Requests: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
									},
									Strategy: &kubernetes.DeploymentStrategy{
										RollingUpdate: &kubernetes.RollingUpdateDeployment{
											MaxSurge: &v1alpha12.IntOrString{
												IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
												StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
												Type:   model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
											},
											MaxUnavailable: &v1alpha12.IntOrString{
												IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
												StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
												Type:   model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
											},
										},
										Type: model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Strategy.Type.ValueString(),
									},
									Tolerations: func() []*v11.Toleration {
										a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Tolerations))
										for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Deployment.Tolerations {
											a[i] = &v11.Toleration{
												Effect:            ptrify(tolerations.Effect.ValueString()),
												Key:               ptrify(tolerations.Key.ValueString()),
												Operator:          ptrify(tolerations.Operator.ValueString()),
												TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
												Value:             ptrify(tolerations.Value.ValueString()),
											}
										}
										return a
									}(),
								},
								Overlays: func() []*v1alpha12.K8SObjectOverlay {
									a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Overlays))
									for i, overlays := range model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Overlays {
										a[i] = &v1alpha12.K8SObjectOverlay{
											ApiVersion: overlays.ApiVersion.ValueString(),
											Kind:       overlays.Kind.ValueString(),
											Name:       overlays.Name.ValueString(),
											Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
												a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
												for i, patches := range overlays.Patches {
													a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
														Path: patches.Path.ValueString(),
														Value: &structpb.Value{
															BoolValue: patches.Value.BoolValue.ValueBool(),
															ListValue: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
															NumberValue: patches.Value.NumberValue.ValueFloat64(),
															StringValue: patches.Value.StringValue.ValueString(),
															StructValue: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
								Service: &kubernetes.Service{
									Annotations: func() map[string]string {
										tmp := make(map[string]string)
										resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									Labels: func() map[string]string {
										tmp := make(map[string]string)
										resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									Ports: func() []*kubernetes.ServicePort {
										a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Service.Ports))
										for i, ports := range model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Service.Ports {
											a[i] = &kubernetes.ServicePort{
												Name:     ports.Name.ValueString(),
												NodePort: int32(ports.NodePort.ValueInt64()),
												Port:     int32(ports.Port.ValueInt64()),
												Protocol: ports.Protocol.ValueString(),
												TargetPort: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
													Type:   ports.TargetPort.Type.ValueInt64(),
												},
											}
										}
										return a
									}(),
									Type: model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.Service.Type.ValueString(),
								},
								ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
									a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.ServiceAccount.ImagePullSecrets))
									for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.HpaAdapter.KubeSpec.ServiceAccount.ImagePullSecrets {
										a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
									}
									return a
								}()},
							}},
							InternalCertProvider: &common.InternalCertProvider{
								CertManager: &common.CertManagerSettings{
									CertManagerCaInjector: &common.CertManagerSettings_CertManagerCAInjector{KubeSpec: &kubernetes.KubernetesComponentSpec{
										Deployment: &kubernetes.Deployment{
											Affinity: &kubernetes.Affinity{
												NodeAffinity: &kubernetes.NodeAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
														a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PreferredSchedulingTerm{
																Preference: &kubernetes.NodeSelectorTerm{
																	MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																		a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																			a[i] = &kubernetes.NodeSelectorRequirement{
																				Key:      match_expressions.Key.ValueString(),
																				Operator: match_expressions.Operator.ValueString(),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																		a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																		for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																			a[i] = &kubernetes.NodeSelectorRequirement{
																				Key:      match_fields.Key.ValueString(),
																				Operator: match_fields.Operator.ValueString(),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_fields.Values.Elements()))
																					resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
														a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
														for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
															a[i] = &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																	for i, match_expressions := range node_selector_terms.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																	for i, match_fields := range node_selector_terms.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															}
														}
														return a
													}()},
												},
												PodAffinity: &kubernetes.PodAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
														a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.WeightedPodAffinityTerm{
																PodAffinityTerm: &kubernetes.PodAffinityTerm{
																	LabelSelector: &v1.LabelSelector{
																		MatchExpressions: func() []*v1.LabelSelectorRequirement {
																			a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																			for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																				a[i] = &v1.LabelSelectorRequirement{
																					Key:      ptrify(match_expressions.Key.ValueString()),
																					Operator: ptrify(match_expressions.Operator.ValueString()),
																					Values: func() []string {
																						tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																						resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																						return tmp
																					}(),
																				}
																			}
																			return a
																		}(),
																		MatchLabels: func() map[string]string {
																			tmp := make(map[string]string)
																			resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	},
																	Namespaces: func() []string {
																		tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
														a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
														for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																		for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
															}
														}
														return a
													}(),
												},
												PodAntiAffinity: &kubernetes.PodAntiAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
														a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.WeightedPodAffinityTerm{
																PodAffinityTerm: &kubernetes.PodAffinityTerm{
																	LabelSelector: &v1.LabelSelector{
																		MatchExpressions: func() []*v1.LabelSelectorRequirement {
																			a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																			for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																				a[i] = &v1.LabelSelectorRequirement{
																					Key:      ptrify(match_expressions.Key.ValueString()),
																					Operator: ptrify(match_expressions.Operator.ValueString()),
																					Values: func() []string {
																						tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																						resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																						return tmp
																					}(),
																				}
																			}
																			return a
																		}(),
																		MatchLabels: func() map[string]string {
																			tmp := make(map[string]string)
																			resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	},
																	Namespaces: func() []string {
																		tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
														a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
														for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																		for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
															}
														}
														return a
													}(),
												},
											},
											ContainerSecurityContext: &kubernetes.SecurityContext{
												AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
												Capabilities: &kubernetes.Capabilities{
													Add: func() []string {
														tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
														resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
													Drop: func() []string {
														tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
														resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
												},
												Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
												ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
												ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
												RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
												RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
												RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
												SeLinuxOptions: &kubernetes.SELinuxOptions{
													Level: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
													Role:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
													Type:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
													User:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
												},
												SeccompProfile: &kubernetes.SeccompProfile{
													LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
													Type:             model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
												},
												WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
													GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
													GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
													RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
												},
											},
											Env: func() []*kubernetes.EnvVar {
												a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Env))
												for i, env := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Env {
													a[i] = &kubernetes.EnvVar{
														Name:  env.Name.ValueString(),
														Value: env.Value.ValueString(),
														ValueFrom: &kubernetes.EnvVarSource{
															ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
																Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
																LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
																Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
															},
															FieldRef: &kubernetes.ObjectFieldSelector{
																ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
																FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
															},
															ResourceFieldRef: &kubernetes.ResourceFieldSelector{
																ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
																Divisor: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																	Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
																},
																Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
															},
															SecretKeyRef: &kubernetes.SecretKeySelector{
																Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
																LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
																Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
															},
														},
													}
												}
												return a
											}(),
											HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
												MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
												Metrics: func() []*kubernetes.MetricSpec {
													a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.HpaSpec.Metrics))
													for i, metrics := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.HpaSpec.Metrics {
														a[i] = &kubernetes.MetricSpec{
															External: &kubernetes.ExternalMetricSource{
																MetricName: metrics.External.MetricName.ValueString(),
																MetricSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																		for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																TargetAverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
																},
																TargetValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																	Type:   metrics.External.TargetValue.Type.ValueInt64(),
																},
															},
															Object: &kubernetes.ObjectMetricSource{
																AverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Object.AverageValue.Type.ValueInt64(),
																},
																MetricName: metrics.Object.MetricName.ValueString(),
																Selector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																		for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Target: &kubernetes.CrossVersionObjectReference{
																	ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																	Kind:       metrics.Object.Target.Kind.ValueString(),
																	Name:       metrics.Object.Target.Name.ValueString(),
																},
																TargetValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																	Type:   metrics.Object.TargetValue.Type.ValueInt64(),
																},
															},
															Pods: &kubernetes.PodsMetricSource{
																MetricName: metrics.Pods.MetricName.ValueString(),
																Selector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																		for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																TargetAverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
																},
															},
															Resource: &kubernetes.ResourceMetricSource{
																Name: metrics.Resource.Name.ValueString(),
																Target: &kubernetes.MetricTarget{
																	AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																	AverageValue: &v1alpha12.IntOrString{
																		IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																		StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																		Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																	},
																	Type: metrics.Resource.Target.Type.ValueString(),
																	Value: &v1alpha12.IntOrString{
																		IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																		StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																		Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																	},
																},
																TargetAverageUtilization: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
																},
																TargetAverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
																},
															},
															Type: metrics.Type.ValueString(),
														}
													}
													return a
												}(),
												MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
											},
											PodAnnotations: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											PodSecurityContext: &kubernetes.PodSecurityContext{
												FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
												FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
												RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
												RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
												RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
												SeLinuxOptions: &kubernetes.SELinuxOptions{
													Level: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
													Role:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
													Type:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
													User:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
												},
												SeccompProfile: &kubernetes.SeccompProfile{
													LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
													Type:             model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
												},
												SupplementalGroups: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Sysctls: func() []*kubernetes.Sysctl {
													a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.Sysctls))
													for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.Sysctls {
														a[i] = &kubernetes.Sysctl{
															Name:  sysctls.Name.ValueString(),
															Value: sysctls.Value.ValueString(),
														}
													}
													return a
												}(),
												WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
													GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
													GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
													RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
												},
											},
											ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
											Resources: &kubernetes.Resources{
												Limits: func() map[string]string {
													tmp := make(map[string]string)
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Requests: func() map[string]string {
													tmp := make(map[string]string)
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
											},
											Strategy: &kubernetes.DeploymentStrategy{
												RollingUpdate: &kubernetes.RollingUpdateDeployment{
													MaxSurge: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
														Type:   model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
													},
													MaxUnavailable: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
														Type:   model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
													},
												},
												Type: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Strategy.Type.ValueString(),
											},
											Tolerations: func() []*v11.Toleration {
												a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Tolerations))
												for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Deployment.Tolerations {
													a[i] = &v11.Toleration{
														Effect:            ptrify(tolerations.Effect.ValueString()),
														Key:               ptrify(tolerations.Key.ValueString()),
														Operator:          ptrify(tolerations.Operator.ValueString()),
														TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
														Value:             ptrify(tolerations.Value.ValueString()),
													}
												}
												return a
											}(),
										},
										Overlays: func() []*v1alpha12.K8SObjectOverlay {
											a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Overlays))
											for i, overlays := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Overlays {
												a[i] = &v1alpha12.K8SObjectOverlay{
													ApiVersion: overlays.ApiVersion.ValueString(),
													Kind:       overlays.Kind.ValueString(),
													Name:       overlays.Name.ValueString(),
													Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
														a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
														for i, patches := range overlays.Patches {
															a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
																Path: patches.Path.ValueString(),
																Value: &structpb.Value{
																	BoolValue: patches.Value.BoolValue.ValueBool(),
																	ListValue: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																	NumberValue: patches.Value.NumberValue.ValueFloat64(),
																	StringValue: patches.Value.StringValue.ValueString(),
																	StructValue: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
										Service: &kubernetes.Service{
											Annotations: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Labels: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Ports: func() []*kubernetes.ServicePort {
												a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Service.Ports))
												for i, ports := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Service.Ports {
													a[i] = &kubernetes.ServicePort{
														Name:     ports.Name.ValueString(),
														NodePort: int32(ports.NodePort.ValueInt64()),
														Port:     int32(ports.Port.ValueInt64()),
														Protocol: ports.Protocol.ValueString(),
														TargetPort: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
															Type:   ports.TargetPort.Type.ValueInt64(),
														},
													}
												}
												return a
											}(),
											Type: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.Service.Type.ValueString(),
										},
										ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
											a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.ServiceAccount.ImagePullSecrets))
											for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerCaInjector.KubeSpec.ServiceAccount.ImagePullSecrets {
												a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
											}
											return a
										}()},
									}},
									CertManagerSpec: &common.CertManagerSettings_CertManagerSpec{KubeSpec: &kubernetes.KubernetesComponentSpec{
										Deployment: &kubernetes.Deployment{
											Affinity: &kubernetes.Affinity{
												NodeAffinity: &kubernetes.NodeAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
														a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PreferredSchedulingTerm{
																Preference: &kubernetes.NodeSelectorTerm{
																	MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																		a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																			a[i] = &kubernetes.NodeSelectorRequirement{
																				Key:      match_expressions.Key.ValueString(),
																				Operator: match_expressions.Operator.ValueString(),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																		a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																		for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																			a[i] = &kubernetes.NodeSelectorRequirement{
																				Key:      match_fields.Key.ValueString(),
																				Operator: match_fields.Operator.ValueString(),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_fields.Values.Elements()))
																					resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
														a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
														for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
															a[i] = &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																	for i, match_expressions := range node_selector_terms.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																	for i, match_fields := range node_selector_terms.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															}
														}
														return a
													}()},
												},
												PodAffinity: &kubernetes.PodAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
														a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.WeightedPodAffinityTerm{
																PodAffinityTerm: &kubernetes.PodAffinityTerm{
																	LabelSelector: &v1.LabelSelector{
																		MatchExpressions: func() []*v1.LabelSelectorRequirement {
																			a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																			for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																				a[i] = &v1.LabelSelectorRequirement{
																					Key:      ptrify(match_expressions.Key.ValueString()),
																					Operator: ptrify(match_expressions.Operator.ValueString()),
																					Values: func() []string {
																						tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																						resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																						return tmp
																					}(),
																				}
																			}
																			return a
																		}(),
																		MatchLabels: func() map[string]string {
																			tmp := make(map[string]string)
																			resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	},
																	Namespaces: func() []string {
																		tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
														a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
														for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																		for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
															}
														}
														return a
													}(),
												},
												PodAntiAffinity: &kubernetes.PodAntiAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
														a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.WeightedPodAffinityTerm{
																PodAffinityTerm: &kubernetes.PodAffinityTerm{
																	LabelSelector: &v1.LabelSelector{
																		MatchExpressions: func() []*v1.LabelSelectorRequirement {
																			a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																			for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																				a[i] = &v1.LabelSelectorRequirement{
																					Key:      ptrify(match_expressions.Key.ValueString()),
																					Operator: ptrify(match_expressions.Operator.ValueString()),
																					Values: func() []string {
																						tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																						resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																						return tmp
																					}(),
																				}
																			}
																			return a
																		}(),
																		MatchLabels: func() map[string]string {
																			tmp := make(map[string]string)
																			resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	},
																	Namespaces: func() []string {
																		tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
														a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
														for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																		for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
															}
														}
														return a
													}(),
												},
											},
											ContainerSecurityContext: &kubernetes.SecurityContext{
												AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
												Capabilities: &kubernetes.Capabilities{
													Add: func() []string {
														tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
														resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
													Drop: func() []string {
														tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
														resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
												},
												Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
												ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
												ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
												RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
												RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
												RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
												SeLinuxOptions: &kubernetes.SELinuxOptions{
													Level: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
													Role:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
													Type:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
													User:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
												},
												SeccompProfile: &kubernetes.SeccompProfile{
													LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
													Type:             model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
												},
												WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
													GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
													GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
													RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
												},
											},
											Env: func() []*kubernetes.EnvVar {
												a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Env))
												for i, env := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Env {
													a[i] = &kubernetes.EnvVar{
														Name:  env.Name.ValueString(),
														Value: env.Value.ValueString(),
														ValueFrom: &kubernetes.EnvVarSource{
															ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
																Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
																LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
																Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
															},
															FieldRef: &kubernetes.ObjectFieldSelector{
																ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
																FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
															},
															ResourceFieldRef: &kubernetes.ResourceFieldSelector{
																ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
																Divisor: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																	Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
																},
																Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
															},
															SecretKeyRef: &kubernetes.SecretKeySelector{
																Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
																LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
																Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
															},
														},
													}
												}
												return a
											}(),
											HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
												MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
												Metrics: func() []*kubernetes.MetricSpec {
													a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.HpaSpec.Metrics))
													for i, metrics := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.HpaSpec.Metrics {
														a[i] = &kubernetes.MetricSpec{
															External: &kubernetes.ExternalMetricSource{
																MetricName: metrics.External.MetricName.ValueString(),
																MetricSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																		for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																TargetAverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
																},
																TargetValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																	Type:   metrics.External.TargetValue.Type.ValueInt64(),
																},
															},
															Object: &kubernetes.ObjectMetricSource{
																AverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Object.AverageValue.Type.ValueInt64(),
																},
																MetricName: metrics.Object.MetricName.ValueString(),
																Selector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																		for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Target: &kubernetes.CrossVersionObjectReference{
																	ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																	Kind:       metrics.Object.Target.Kind.ValueString(),
																	Name:       metrics.Object.Target.Name.ValueString(),
																},
																TargetValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																	Type:   metrics.Object.TargetValue.Type.ValueInt64(),
																},
															},
															Pods: &kubernetes.PodsMetricSource{
																MetricName: metrics.Pods.MetricName.ValueString(),
																Selector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																		for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																TargetAverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
																},
															},
															Resource: &kubernetes.ResourceMetricSource{
																Name: metrics.Resource.Name.ValueString(),
																Target: &kubernetes.MetricTarget{
																	AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																	AverageValue: &v1alpha12.IntOrString{
																		IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																		StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																		Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																	},
																	Type: metrics.Resource.Target.Type.ValueString(),
																	Value: &v1alpha12.IntOrString{
																		IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																		StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																		Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																	},
																},
																TargetAverageUtilization: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
																},
																TargetAverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
																},
															},
															Type: metrics.Type.ValueString(),
														}
													}
													return a
												}(),
												MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
											},
											PodAnnotations: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											PodSecurityContext: &kubernetes.PodSecurityContext{
												FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
												FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
												RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
												RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
												RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
												SeLinuxOptions: &kubernetes.SELinuxOptions{
													Level: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
													Role:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
													Type:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
													User:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
												},
												SeccompProfile: &kubernetes.SeccompProfile{
													LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
													Type:             model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
												},
												SupplementalGroups: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Sysctls: func() []*kubernetes.Sysctl {
													a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.Sysctls))
													for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.Sysctls {
														a[i] = &kubernetes.Sysctl{
															Name:  sysctls.Name.ValueString(),
															Value: sysctls.Value.ValueString(),
														}
													}
													return a
												}(),
												WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
													GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
													GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
													RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
												},
											},
											ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
											Resources: &kubernetes.Resources{
												Limits: func() map[string]string {
													tmp := make(map[string]string)
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Requests: func() map[string]string {
													tmp := make(map[string]string)
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
											},
											Strategy: &kubernetes.DeploymentStrategy{
												RollingUpdate: &kubernetes.RollingUpdateDeployment{
													MaxSurge: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
														Type:   model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
													},
													MaxUnavailable: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
														Type:   model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
													},
												},
												Type: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Strategy.Type.ValueString(),
											},
											Tolerations: func() []*v11.Toleration {
												a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Tolerations))
												for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Deployment.Tolerations {
													a[i] = &v11.Toleration{
														Effect:            ptrify(tolerations.Effect.ValueString()),
														Key:               ptrify(tolerations.Key.ValueString()),
														Operator:          ptrify(tolerations.Operator.ValueString()),
														TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
														Value:             ptrify(tolerations.Value.ValueString()),
													}
												}
												return a
											}(),
										},
										Overlays: func() []*v1alpha12.K8SObjectOverlay {
											a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Overlays))
											for i, overlays := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Overlays {
												a[i] = &v1alpha12.K8SObjectOverlay{
													ApiVersion: overlays.ApiVersion.ValueString(),
													Kind:       overlays.Kind.ValueString(),
													Name:       overlays.Name.ValueString(),
													Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
														a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
														for i, patches := range overlays.Patches {
															a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
																Path: patches.Path.ValueString(),
																Value: &structpb.Value{
																	BoolValue: patches.Value.BoolValue.ValueBool(),
																	ListValue: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																	NumberValue: patches.Value.NumberValue.ValueFloat64(),
																	StringValue: patches.Value.StringValue.ValueString(),
																	StructValue: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
										Service: &kubernetes.Service{
											Annotations: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Labels: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Ports: func() []*kubernetes.ServicePort {
												a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Service.Ports))
												for i, ports := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Service.Ports {
													a[i] = &kubernetes.ServicePort{
														Name:     ports.Name.ValueString(),
														NodePort: int32(ports.NodePort.ValueInt64()),
														Port:     int32(ports.Port.ValueInt64()),
														Protocol: ports.Protocol.ValueString(),
														TargetPort: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
															Type:   ports.TargetPort.Type.ValueInt64(),
														},
													}
												}
												return a
											}(),
											Type: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.Service.Type.ValueString(),
										},
										ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
											a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.ServiceAccount.ImagePullSecrets))
											for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerSpec.KubeSpec.ServiceAccount.ImagePullSecrets {
												a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
											}
											return a
										}()},
									}},
									CertManagerStartupapicheck: &common.CertManagerSettings_CertManagerStartupAPICheck{KubeSpec: &kubernetes.KubernetesJobComponentSpec{
										Deployment: &kubernetes.Deployment{
											Affinity: &kubernetes.Affinity{
												NodeAffinity: &kubernetes.NodeAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
														a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PreferredSchedulingTerm{
																Preference: &kubernetes.NodeSelectorTerm{
																	MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																		a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																			a[i] = &kubernetes.NodeSelectorRequirement{
																				Key:      match_expressions.Key.ValueString(),
																				Operator: match_expressions.Operator.ValueString(),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																		a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																		for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																			a[i] = &kubernetes.NodeSelectorRequirement{
																				Key:      match_fields.Key.ValueString(),
																				Operator: match_fields.Operator.ValueString(),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_fields.Values.Elements()))
																					resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
														a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
														for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
															a[i] = &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																	for i, match_expressions := range node_selector_terms.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																	for i, match_fields := range node_selector_terms.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															}
														}
														return a
													}()},
												},
												PodAffinity: &kubernetes.PodAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
														a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.WeightedPodAffinityTerm{
																PodAffinityTerm: &kubernetes.PodAffinityTerm{
																	LabelSelector: &v1.LabelSelector{
																		MatchExpressions: func() []*v1.LabelSelectorRequirement {
																			a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																			for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																				a[i] = &v1.LabelSelectorRequirement{
																					Key:      ptrify(match_expressions.Key.ValueString()),
																					Operator: ptrify(match_expressions.Operator.ValueString()),
																					Values: func() []string {
																						tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																						resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																						return tmp
																					}(),
																				}
																			}
																			return a
																		}(),
																		MatchLabels: func() map[string]string {
																			tmp := make(map[string]string)
																			resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	},
																	Namespaces: func() []string {
																		tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
														a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
														for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																		for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
															}
														}
														return a
													}(),
												},
												PodAntiAffinity: &kubernetes.PodAntiAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
														a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.WeightedPodAffinityTerm{
																PodAffinityTerm: &kubernetes.PodAffinityTerm{
																	LabelSelector: &v1.LabelSelector{
																		MatchExpressions: func() []*v1.LabelSelectorRequirement {
																			a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																			for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																				a[i] = &v1.LabelSelectorRequirement{
																					Key:      ptrify(match_expressions.Key.ValueString()),
																					Operator: ptrify(match_expressions.Operator.ValueString()),
																					Values: func() []string {
																						tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																						resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																						return tmp
																					}(),
																				}
																			}
																			return a
																		}(),
																		MatchLabels: func() map[string]string {
																			tmp := make(map[string]string)
																			resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	},
																	Namespaces: func() []string {
																		tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
														a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
														for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																		for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
															}
														}
														return a
													}(),
												},
											},
											ContainerSecurityContext: &kubernetes.SecurityContext{
												AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
												Capabilities: &kubernetes.Capabilities{
													Add: func() []string {
														tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
														resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
													Drop: func() []string {
														tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
														resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
												},
												Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
												ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
												ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
												RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
												RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
												RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
												SeLinuxOptions: &kubernetes.SELinuxOptions{
													Level: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
													Role:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
													Type:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
													User:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
												},
												SeccompProfile: &kubernetes.SeccompProfile{
													LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
													Type:             model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
												},
												WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
													GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
													GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
													RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
												},
											},
											Env: func() []*kubernetes.EnvVar {
												a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Env))
												for i, env := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Env {
													a[i] = &kubernetes.EnvVar{
														Name:  env.Name.ValueString(),
														Value: env.Value.ValueString(),
														ValueFrom: &kubernetes.EnvVarSource{
															ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
																Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
																LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
																Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
															},
															FieldRef: &kubernetes.ObjectFieldSelector{
																ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
																FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
															},
															ResourceFieldRef: &kubernetes.ResourceFieldSelector{
																ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
																Divisor: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																	Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
																},
																Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
															},
															SecretKeyRef: &kubernetes.SecretKeySelector{
																Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
																LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
																Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
															},
														},
													}
												}
												return a
											}(),
											HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
												MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
												Metrics: func() []*kubernetes.MetricSpec {
													a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.HpaSpec.Metrics))
													for i, metrics := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.HpaSpec.Metrics {
														a[i] = &kubernetes.MetricSpec{
															External: &kubernetes.ExternalMetricSource{
																MetricName: metrics.External.MetricName.ValueString(),
																MetricSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																		for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																TargetAverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
																},
																TargetValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																	Type:   metrics.External.TargetValue.Type.ValueInt64(),
																},
															},
															Object: &kubernetes.ObjectMetricSource{
																AverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Object.AverageValue.Type.ValueInt64(),
																},
																MetricName: metrics.Object.MetricName.ValueString(),
																Selector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																		for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Target: &kubernetes.CrossVersionObjectReference{
																	ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																	Kind:       metrics.Object.Target.Kind.ValueString(),
																	Name:       metrics.Object.Target.Name.ValueString(),
																},
																TargetValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																	Type:   metrics.Object.TargetValue.Type.ValueInt64(),
																},
															},
															Pods: &kubernetes.PodsMetricSource{
																MetricName: metrics.Pods.MetricName.ValueString(),
																Selector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																		for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																TargetAverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
																},
															},
															Resource: &kubernetes.ResourceMetricSource{
																Name: metrics.Resource.Name.ValueString(),
																Target: &kubernetes.MetricTarget{
																	AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																	AverageValue: &v1alpha12.IntOrString{
																		IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																		StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																		Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																	},
																	Type: metrics.Resource.Target.Type.ValueString(),
																	Value: &v1alpha12.IntOrString{
																		IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																		StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																		Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																	},
																},
																TargetAverageUtilization: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
																},
																TargetAverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
																},
															},
															Type: metrics.Type.ValueString(),
														}
													}
													return a
												}(),
												MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
											},
											PodAnnotations: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											PodSecurityContext: &kubernetes.PodSecurityContext{
												FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
												FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
												RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
												RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
												RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
												SeLinuxOptions: &kubernetes.SELinuxOptions{
													Level: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
													Role:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
													Type:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
													User:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
												},
												SeccompProfile: &kubernetes.SeccompProfile{
													LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
													Type:             model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
												},
												SupplementalGroups: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Sysctls: func() []*kubernetes.Sysctl {
													a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.Sysctls))
													for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.Sysctls {
														a[i] = &kubernetes.Sysctl{
															Name:  sysctls.Name.ValueString(),
															Value: sysctls.Value.ValueString(),
														}
													}
													return a
												}(),
												WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
													GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
													GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
													RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
												},
											},
											ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
											Resources: &kubernetes.Resources{
												Limits: func() map[string]string {
													tmp := make(map[string]string)
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Requests: func() map[string]string {
													tmp := make(map[string]string)
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
											},
											Strategy: &kubernetes.DeploymentStrategy{
												RollingUpdate: &kubernetes.RollingUpdateDeployment{
													MaxSurge: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
														Type:   model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
													},
													MaxUnavailable: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
														Type:   model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
													},
												},
												Type: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Strategy.Type.ValueString(),
											},
											Tolerations: func() []*v11.Toleration {
												a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Tolerations))
												for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Deployment.Tolerations {
													a[i] = &v11.Toleration{
														Effect:            ptrify(tolerations.Effect.ValueString()),
														Key:               ptrify(tolerations.Key.ValueString()),
														Operator:          ptrify(tolerations.Operator.ValueString()),
														TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
														Value:             ptrify(tolerations.Value.ValueString()),
													}
												}
												return a
											}(),
										},
										Job: &kubernetes.Job{
											Affinity: &kubernetes.Affinity{
												NodeAffinity: &kubernetes.NodeAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
														a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PreferredSchedulingTerm{
																Preference: &kubernetes.NodeSelectorTerm{
																	MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																		a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																			a[i] = &kubernetes.NodeSelectorRequirement{
																				Key:      match_expressions.Key.ValueString(),
																				Operator: match_expressions.Operator.ValueString(),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																		a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																		for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																			a[i] = &kubernetes.NodeSelectorRequirement{
																				Key:      match_fields.Key.ValueString(),
																				Operator: match_fields.Operator.ValueString(),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_fields.Values.Elements()))
																					resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
														a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
														for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
															a[i] = &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																	for i, match_expressions := range node_selector_terms.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																	for i, match_fields := range node_selector_terms.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															}
														}
														return a
													}()},
												},
												PodAffinity: &kubernetes.PodAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
														a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.WeightedPodAffinityTerm{
																PodAffinityTerm: &kubernetes.PodAffinityTerm{
																	LabelSelector: &v1.LabelSelector{
																		MatchExpressions: func() []*v1.LabelSelectorRequirement {
																			a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																			for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																				a[i] = &v1.LabelSelectorRequirement{
																					Key:      ptrify(match_expressions.Key.ValueString()),
																					Operator: ptrify(match_expressions.Operator.ValueString()),
																					Values: func() []string {
																						tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																						resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																						return tmp
																					}(),
																				}
																			}
																			return a
																		}(),
																		MatchLabels: func() map[string]string {
																			tmp := make(map[string]string)
																			resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	},
																	Namespaces: func() []string {
																		tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
														a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
														for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																		for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
															}
														}
														return a
													}(),
												},
												PodAntiAffinity: &kubernetes.PodAntiAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
														a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.WeightedPodAffinityTerm{
																PodAffinityTerm: &kubernetes.PodAffinityTerm{
																	LabelSelector: &v1.LabelSelector{
																		MatchExpressions: func() []*v1.LabelSelectorRequirement {
																			a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																			for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																				a[i] = &v1.LabelSelectorRequirement{
																					Key:      ptrify(match_expressions.Key.ValueString()),
																					Operator: ptrify(match_expressions.Operator.ValueString()),
																					Values: func() []string {
																						tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																						resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																						return tmp
																					}(),
																				}
																			}
																			return a
																		}(),
																		MatchLabels: func() map[string]string {
																			tmp := make(map[string]string)
																			resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	},
																	Namespaces: func() []string {
																		tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
														a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
														for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																		for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
															}
														}
														return a
													}(),
												},
											},
											ContainerSecurityContext: &kubernetes.SecurityContext{
												AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
												Capabilities: &kubernetes.Capabilities{
													Add: func() []string {
														tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.Capabilities.Add.Elements()))
														resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
													Drop: func() []string {
														tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.Capabilities.Drop.Elements()))
														resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
												},
												Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.Privileged.ValueBool()),
												ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.ProcMount.ValueString()),
												ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
												RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.RunAsGroup.ValueInt64())),
												RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
												RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.RunAsUser.ValueInt64())),
												SeLinuxOptions: &kubernetes.SELinuxOptions{
													Level: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
													Role:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
													Type:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
													User:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
												},
												SeccompProfile: &kubernetes.SeccompProfile{
													LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
													Type:             model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
												},
												WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
													GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
													GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
													RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
												},
											},
											Env: func() []*kubernetes.EnvVar {
												a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Env))
												for i, env := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Env {
													a[i] = &kubernetes.EnvVar{
														Name:  env.Name.ValueString(),
														Value: env.Value.ValueString(),
														ValueFrom: &kubernetes.EnvVarSource{
															ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
																Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
																LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
																Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
															},
															FieldRef: &kubernetes.ObjectFieldSelector{
																ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
																FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
															},
															ResourceFieldRef: &kubernetes.ResourceFieldSelector{
																ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
																Divisor: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																	Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
																},
																Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
															},
															SecretKeyRef: &kubernetes.SecretKeySelector{
																Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
																LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
																Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
															},
														},
													}
												}
												return a
											}(),
											PodAnnotations: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											PodSecurityContext: &kubernetes.PodSecurityContext{
												FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.FsGroup.ValueInt64())),
												FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.FsGroupChangePolicy.ValueString()),
												RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.RunAsGroup.ValueInt64())),
												RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.RunAsNonRoot.ValueBool()),
												RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.RunAsUser.ValueInt64())),
												SeLinuxOptions: &kubernetes.SELinuxOptions{
													Level: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
													Role:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
													Type:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
													User:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeLinuxOptions.User.ValueString(),
												},
												SeccompProfile: &kubernetes.SeccompProfile{
													LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
													Type:             model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SeccompProfile.Type.ValueString(),
												},
												SupplementalGroups: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SupplementalGroups.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Sysctls: func() []*kubernetes.Sysctl {
													a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.Sysctls))
													for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.Sysctls {
														a[i] = &kubernetes.Sysctl{
															Name:  sysctls.Name.ValueString(),
															Value: sysctls.Value.ValueString(),
														}
													}
													return a
												}(),
												WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
													GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
													GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
													RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
												},
											},
											Tolerations: func() []*v11.Toleration {
												a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Tolerations))
												for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Job.Tolerations {
													a[i] = &v11.Toleration{
														Effect:            ptrify(tolerations.Effect.ValueString()),
														Key:               ptrify(tolerations.Key.ValueString()),
														Operator:          ptrify(tolerations.Operator.ValueString()),
														TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
														Value:             ptrify(tolerations.Value.ValueString()),
													}
												}
												return a
											}(),
										},
										Overlays: func() []*v1alpha12.K8SObjectOverlay {
											a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Overlays))
											for i, overlays := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Overlays {
												a[i] = &v1alpha12.K8SObjectOverlay{
													ApiVersion: overlays.ApiVersion.ValueString(),
													Kind:       overlays.Kind.ValueString(),
													Name:       overlays.Name.ValueString(),
													Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
														a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
														for i, patches := range overlays.Patches {
															a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
																Path: patches.Path.ValueString(),
																Value: &structpb.Value{
																	BoolValue: patches.Value.BoolValue.ValueBool(),
																	ListValue: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																	NumberValue: patches.Value.NumberValue.ValueFloat64(),
																	StringValue: patches.Value.StringValue.ValueString(),
																	StructValue: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
										Service: &kubernetes.Service{
											Annotations: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Labels: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Ports: func() []*kubernetes.ServicePort {
												a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Service.Ports))
												for i, ports := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Service.Ports {
													a[i] = &kubernetes.ServicePort{
														Name:     ports.Name.ValueString(),
														NodePort: int32(ports.NodePort.ValueInt64()),
														Port:     int32(ports.Port.ValueInt64()),
														Protocol: ports.Protocol.ValueString(),
														TargetPort: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
															Type:   ports.TargetPort.Type.ValueInt64(),
														},
													}
												}
												return a
											}(),
											Type: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.Service.Type.ValueString(),
										},
										ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
											a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.ServiceAccount.ImagePullSecrets))
											for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerStartupapicheck.KubeSpec.ServiceAccount.ImagePullSecrets {
												a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
											}
											return a
										}()},
									}},
									CertManagerWebhookSpec: &common.CertManagerSettings_CertManagerWebhookSpec{KubeSpec: &kubernetes.KubernetesComponentSpec{
										Deployment: &kubernetes.Deployment{
											Affinity: &kubernetes.Affinity{
												NodeAffinity: &kubernetes.NodeAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
														a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PreferredSchedulingTerm{
																Preference: &kubernetes.NodeSelectorTerm{
																	MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																		a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																			a[i] = &kubernetes.NodeSelectorRequirement{
																				Key:      match_expressions.Key.ValueString(),
																				Operator: match_expressions.Operator.ValueString(),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																		a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																		for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																			a[i] = &kubernetes.NodeSelectorRequirement{
																				Key:      match_fields.Key.ValueString(),
																				Operator: match_fields.Operator.ValueString(),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_fields.Values.Elements()))
																					resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
														a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
														for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
															a[i] = &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																	for i, match_expressions := range node_selector_terms.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																	for i, match_fields := range node_selector_terms.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															}
														}
														return a
													}()},
												},
												PodAffinity: &kubernetes.PodAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
														a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.WeightedPodAffinityTerm{
																PodAffinityTerm: &kubernetes.PodAffinityTerm{
																	LabelSelector: &v1.LabelSelector{
																		MatchExpressions: func() []*v1.LabelSelectorRequirement {
																			a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																			for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																				a[i] = &v1.LabelSelectorRequirement{
																					Key:      ptrify(match_expressions.Key.ValueString()),
																					Operator: ptrify(match_expressions.Operator.ValueString()),
																					Values: func() []string {
																						tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																						resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																						return tmp
																					}(),
																				}
																			}
																			return a
																		}(),
																		MatchLabels: func() map[string]string {
																			tmp := make(map[string]string)
																			resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	},
																	Namespaces: func() []string {
																		tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
														a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
														for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																		for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
															}
														}
														return a
													}(),
												},
												PodAntiAffinity: &kubernetes.PodAntiAffinity{
													PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
														a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
														for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.WeightedPodAffinityTerm{
																PodAffinityTerm: &kubernetes.PodAffinityTerm{
																	LabelSelector: &v1.LabelSelector{
																		MatchExpressions: func() []*v1.LabelSelectorRequirement {
																			a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																			for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																				a[i] = &v1.LabelSelectorRequirement{
																					Key:      ptrify(match_expressions.Key.ValueString()),
																					Operator: ptrify(match_expressions.Operator.ValueString()),
																					Values: func() []string {
																						tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																						resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																						return tmp
																					}(),
																				}
																			}
																			return a
																		}(),
																		MatchLabels: func() map[string]string {
																			tmp := make(map[string]string)
																			resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	},
																	Namespaces: func() []string {
																		tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
																},
																Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
															}
														}
														return a
													}(),
													RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
														a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
														for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
															a[i] = &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																		for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
															}
														}
														return a
													}(),
												},
											},
											ContainerSecurityContext: &kubernetes.SecurityContext{
												AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
												Capabilities: &kubernetes.Capabilities{
													Add: func() []string {
														tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
														resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
													Drop: func() []string {
														tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
														resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
														return tmp
													}(),
												},
												Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
												ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
												ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
												RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
												RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
												RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
												SeLinuxOptions: &kubernetes.SELinuxOptions{
													Level: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
													Role:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
													Type:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
													User:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
												},
												SeccompProfile: &kubernetes.SeccompProfile{
													LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
													Type:             model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
												},
												WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
													GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
													GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
													RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
												},
											},
											Env: func() []*kubernetes.EnvVar {
												a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Env))
												for i, env := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Env {
													a[i] = &kubernetes.EnvVar{
														Name:  env.Name.ValueString(),
														Value: env.Value.ValueString(),
														ValueFrom: &kubernetes.EnvVarSource{
															ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
																Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
																LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
																Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
															},
															FieldRef: &kubernetes.ObjectFieldSelector{
																ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
																FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
															},
															ResourceFieldRef: &kubernetes.ResourceFieldSelector{
																ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
																Divisor: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																	Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
																},
																Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
															},
															SecretKeyRef: &kubernetes.SecretKeySelector{
																Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
																LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
																Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
															},
														},
													}
												}
												return a
											}(),
											HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
												MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
												Metrics: func() []*kubernetes.MetricSpec {
													a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.HpaSpec.Metrics))
													for i, metrics := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.HpaSpec.Metrics {
														a[i] = &kubernetes.MetricSpec{
															External: &kubernetes.ExternalMetricSource{
																MetricName: metrics.External.MetricName.ValueString(),
																MetricSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																		for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																TargetAverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
																},
																TargetValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																	Type:   metrics.External.TargetValue.Type.ValueInt64(),
																},
															},
															Object: &kubernetes.ObjectMetricSource{
																AverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Object.AverageValue.Type.ValueInt64(),
																},
																MetricName: metrics.Object.MetricName.ValueString(),
																Selector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																		for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Target: &kubernetes.CrossVersionObjectReference{
																	ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																	Kind:       metrics.Object.Target.Kind.ValueString(),
																	Name:       metrics.Object.Target.Name.ValueString(),
																},
																TargetValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																	Type:   metrics.Object.TargetValue.Type.ValueInt64(),
																},
															},
															Pods: &kubernetes.PodsMetricSource{
																MetricName: metrics.Pods.MetricName.ValueString(),
																Selector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																		for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																TargetAverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
																},
															},
															Resource: &kubernetes.ResourceMetricSource{
																Name: metrics.Resource.Name.ValueString(),
																Target: &kubernetes.MetricTarget{
																	AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																	AverageValue: &v1alpha12.IntOrString{
																		IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																		StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																		Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																	},
																	Type: metrics.Resource.Target.Type.ValueString(),
																	Value: &v1alpha12.IntOrString{
																		IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																		StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																		Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																	},
																},
																TargetAverageUtilization: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
																},
																TargetAverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
																},
															},
															Type: metrics.Type.ValueString(),
														}
													}
													return a
												}(),
												MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
											},
											PodAnnotations: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											PodSecurityContext: &kubernetes.PodSecurityContext{
												FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
												FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
												RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
												RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
												RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
												SeLinuxOptions: &kubernetes.SELinuxOptions{
													Level: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
													Role:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
													Type:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
													User:  model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
												},
												SeccompProfile: &kubernetes.SeccompProfile{
													LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
													Type:             model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
												},
												SupplementalGroups: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Sysctls: func() []*kubernetes.Sysctl {
													a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.Sysctls))
													for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.Sysctls {
														a[i] = &kubernetes.Sysctl{
															Name:  sysctls.Name.ValueString(),
															Value: sysctls.Value.ValueString(),
														}
													}
													return a
												}(),
												WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
													GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
													GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
													RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
												},
											},
											ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
											Resources: &kubernetes.Resources{
												Limits: func() map[string]string {
													tmp := make(map[string]string)
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Requests: func() map[string]string {
													tmp := make(map[string]string)
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
											},
											Strategy: &kubernetes.DeploymentStrategy{
												RollingUpdate: &kubernetes.RollingUpdateDeployment{
													MaxSurge: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
														Type:   model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
													},
													MaxUnavailable: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
														Type:   model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
													},
												},
												Type: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Strategy.Type.ValueString(),
											},
											Tolerations: func() []*v11.Toleration {
												a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Tolerations))
												for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Deployment.Tolerations {
													a[i] = &v11.Toleration{
														Effect:            ptrify(tolerations.Effect.ValueString()),
														Key:               ptrify(tolerations.Key.ValueString()),
														Operator:          ptrify(tolerations.Operator.ValueString()),
														TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
														Value:             ptrify(tolerations.Value.ValueString()),
													}
												}
												return a
											}(),
										},
										Overlays: func() []*v1alpha12.K8SObjectOverlay {
											a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Overlays))
											for i, overlays := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Overlays {
												a[i] = &v1alpha12.K8SObjectOverlay{
													ApiVersion: overlays.ApiVersion.ValueString(),
													Kind:       overlays.Kind.ValueString(),
													Name:       overlays.Name.ValueString(),
													Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
														a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
														for i, patches := range overlays.Patches {
															a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
																Path: patches.Path.ValueString(),
																Value: &structpb.Value{
																	BoolValue: patches.Value.BoolValue.ValueBool(),
																	ListValue: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																	NumberValue: patches.Value.NumberValue.ValueFloat64(),
																	StringValue: patches.Value.StringValue.ValueString(),
																	StructValue: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
										Service: &kubernetes.Service{
											Annotations: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Labels: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Ports: func() []*kubernetes.ServicePort {
												a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Service.Ports))
												for i, ports := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Service.Ports {
													a[i] = &kubernetes.ServicePort{
														Name:     ports.Name.ValueString(),
														NodePort: int32(ports.NodePort.ValueInt64()),
														Port:     int32(ports.Port.ValueInt64()),
														Protocol: ports.Protocol.ValueString(),
														TargetPort: &v1alpha12.IntOrString{
															IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
															StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
															Type:   ports.TargetPort.Type.ValueInt64(),
														},
													}
												}
												return a
											}(),
											Type: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.Service.Type.ValueString(),
										},
										ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
											a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.ServiceAccount.ImagePullSecrets))
											for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.CertManagerWebhookSpec.KubeSpec.ServiceAccount.ImagePullSecrets {
												a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
											}
											return a
										}()},
									}},
									Managed: common.CertManagerSettings_Managed(common.CertManagerSettings_Managed_value[model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.CertManager.Managed.ValueString()]),
								},
								Custom: &common.CustomCertProviderSettings{
									CaBundleSecretName: model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.Custom.CaBundleSecretName.ValueString(),
									CsrSignerName:      model.InstallTemplate.Helm.Spec.Components.InternalCertProvider.Custom.CsrSignerName.ValueString(),
								},
							},
							Istio: &v1alpha13.Istio{
								BaseOverlays: func() []*v1alpha12.K8SObjectOverlay {
									a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.BaseOverlays))
									for i, base_overlays := range model.InstallTemplate.Helm.Spec.Components.Istio.BaseOverlays {
										a[i] = &v1alpha12.K8SObjectOverlay{
											ApiVersion: base_overlays.ApiVersion.ValueString(),
											Kind:       base_overlays.Kind.ValueString(),
											Name:       base_overlays.Name.ValueString(),
											Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
												a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(base_overlays.Patches))
												for i, patches := range base_overlays.Patches {
													a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
														Path: patches.Path.ValueString(),
														Value: &structpb.Value{
															BoolValue: patches.Value.BoolValue.ValueBool(),
															ListValue: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
															NumberValue: patches.Value.NumberValue.ValueFloat64(),
															StringValue: patches.Value.StringValue.ValueString(),
															StructValue: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
								CniOverlays: func() []*v1alpha12.K8SObjectOverlay {
									a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.CniOverlays))
									for i, cni_overlays := range model.InstallTemplate.Helm.Spec.Components.Istio.CniOverlays {
										a[i] = &v1alpha12.K8SObjectOverlay{
											ApiVersion: cni_overlays.ApiVersion.ValueString(),
											Kind:       cni_overlays.Kind.ValueString(),
											Name:       cni_overlays.Name.ValueString(),
											Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
												a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(cni_overlays.Patches))
												for i, patches := range cni_overlays.Patches {
													a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
														Path: patches.Path.ValueString(),
														Value: &structpb.Value{
															BoolValue: patches.Value.BoolValue.ValueBool(),
															ListValue: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
															NumberValue: patches.Value.NumberValue.ValueFloat64(),
															StringValue: patches.Value.StringValue.ValueString(),
															StructValue: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
								DefaultWorkloadCertTtl: &durationpb.Duration{
									Nanos:   int32(model.InstallTemplate.Helm.Spec.Components.Istio.DefaultWorkloadCertTtl.Nanos.ValueInt64()),
									Seconds: model.InstallTemplate.Helm.Spec.Components.Istio.DefaultWorkloadCertTtl.Seconds.ValueInt64(),
								},
								KubeSpec: &kubernetes.KubernetesIstioComponentSpec{
									Cni: &kubernetes.CNI{
										BinaryDirectory:        model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Cni.BinaryDirectory.ValueString(),
										Chained:                model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Cni.Chained.ValueBool(),
										ClusterRole:            model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Cni.ClusterRole.ValueString(),
										ConfigurationDirectory: model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Cni.ConfigurationDirectory.ValueString(),
										ConfigurationFileName:  model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Cni.ConfigurationFileName.ValueString(),
										Revision:               model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Cni.Revision.ValueString(),
									},
									Deployment: &kubernetes.Deployment{
										Affinity: &kubernetes.Affinity{
											NodeAffinity: &kubernetes.NodeAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
													a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PreferredSchedulingTerm{
															Preference: &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																	for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
													a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
													for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
														a[i] = &kubernetes.NodeSelectorTerm{
															MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																for i, match_expressions := range node_selector_terms.MatchExpressions {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_expressions.Key.ValueString(),
																		Operator: match_expressions.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																for i, match_fields := range node_selector_terms.MatchFields {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_fields.Key.ValueString(),
																		Operator: match_fields.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_fields.Values.Elements()))
																			resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
														}
													}
													return a
												}()},
											},
											PodAffinity: &kubernetes.PodAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
											PodAntiAffinity: &kubernetes.PodAntiAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
										},
										ContainerSecurityContext: &kubernetes.SecurityContext{
											AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
											Capabilities: &kubernetes.Capabilities{
												Add: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Drop: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
											},
											Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
											ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
											ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
											RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
											},
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										Env: func() []*kubernetes.EnvVar {
											a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Env))
											for i, env := range model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Env {
												a[i] = &kubernetes.EnvVar{
													Name:  env.Name.ValueString(),
													Value: env.Value.ValueString(),
													ValueFrom: &kubernetes.EnvVarSource{
														ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
															Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
														},
														FieldRef: &kubernetes.ObjectFieldSelector{
															ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
															FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
														},
														ResourceFieldRef: &kubernetes.ResourceFieldSelector{
															ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
															Divisor: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
															},
															Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
														},
														SecretKeyRef: &kubernetes.SecretKeySelector{
															Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
														},
													},
												}
											}
											return a
										}(),
										HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
											MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
											Metrics: func() []*kubernetes.MetricSpec {
												a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.HpaSpec.Metrics))
												for i, metrics := range model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.HpaSpec.Metrics {
													a[i] = &kubernetes.MetricSpec{
														External: &kubernetes.ExternalMetricSource{
															MetricName: metrics.External.MetricName.ValueString(),
															MetricSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																	for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetValue.Type.ValueInt64(),
															},
														},
														Object: &kubernetes.ObjectMetricSource{
															AverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.AverageValue.Type.ValueInt64(),
															},
															MetricName: metrics.Object.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Target: &kubernetes.CrossVersionObjectReference{
																ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																Kind:       metrics.Object.Target.Kind.ValueString(),
																Name:       metrics.Object.Target.Name.ValueString(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.TargetValue.Type.ValueInt64(),
															},
														},
														Pods: &kubernetes.PodsMetricSource{
															MetricName: metrics.Pods.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Resource: &kubernetes.ResourceMetricSource{
															Name: metrics.Resource.Name.ValueString(),
															Target: &kubernetes.MetricTarget{
																AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																AverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																},
																Type: metrics.Resource.Target.Type.ValueString(),
																Value: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																},
															},
															TargetAverageUtilization: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Type: metrics.Type.ValueString(),
													}
												}
												return a
											}(),
											MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
										},
										PodAnnotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										PodSecurityContext: &kubernetes.PodSecurityContext{
											FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
											FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
											RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
											},
											SupplementalGroups: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Sysctls: func() []*kubernetes.Sysctl {
												a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.Sysctls))
												for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.Sysctls {
													a[i] = &kubernetes.Sysctl{
														Name:  sysctls.Name.ValueString(),
														Value: sysctls.Value.ValueString(),
													}
												}
												return a
											}(),
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
										Resources: &kubernetes.Resources{
											Limits: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Requests: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
										},
										Strategy: &kubernetes.DeploymentStrategy{
											RollingUpdate: &kubernetes.RollingUpdateDeployment{
												MaxSurge: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
												},
												MaxUnavailable: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
												},
											},
											Type: model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Strategy.Type.ValueString(),
										},
										Tolerations: func() []*v11.Toleration {
											a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Tolerations))
											for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Deployment.Tolerations {
												a[i] = &v11.Toleration{
													Effect:            ptrify(tolerations.Effect.ValueString()),
													Key:               ptrify(tolerations.Key.ValueString()),
													Operator:          ptrify(tolerations.Operator.ValueString()),
													TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
													Value:             ptrify(tolerations.Value.ValueString()),
												}
											}
											return a
										}(),
									},
									Overlays: func() []*v1alpha12.K8SObjectOverlay {
										a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Overlays))
										for i, overlays := range model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Overlays {
											a[i] = &v1alpha12.K8SObjectOverlay{
												ApiVersion: overlays.ApiVersion.ValueString(),
												Kind:       overlays.Kind.ValueString(),
												Name:       overlays.Name.ValueString(),
												Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
													a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
													for i, patches := range overlays.Patches {
														a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
															Path: patches.Path.ValueString(),
															Value: &structpb.Value{
																BoolValue: patches.Value.BoolValue.ValueBool(),
																ListValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																NumberValue: patches.Value.NumberValue.ValueFloat64(),
																StringValue: patches.Value.StringValue.ValueString(),
																StructValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
									Service: &kubernetes.Service{
										Annotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Labels: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Ports: func() []*kubernetes.ServicePort {
											a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Service.Ports))
											for i, ports := range model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Service.Ports {
												a[i] = &kubernetes.ServicePort{
													Name:     ports.Name.ValueString(),
													NodePort: int32(ports.NodePort.ValueInt64()),
													Port:     int32(ports.Port.ValueInt64()),
													Protocol: ports.Protocol.ValueString(),
													TargetPort: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
														Type:   ports.TargetPort.Type.ValueInt64(),
													},
												}
											}
											return a
										}(),
										Type: model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.Service.Type.ValueString(),
									},
									ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
										a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.ServiceAccount.ImagePullSecrets))
										for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.Istio.KubeSpec.ServiceAccount.ImagePullSecrets {
											a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
										}
										return a
									}()},
								},
								MaxWorkloadCertTtl: &durationpb.Duration{
									Nanos:   int32(model.InstallTemplate.Helm.Spec.Components.Istio.MaxWorkloadCertTtl.Nanos.ValueInt64()),
									Seconds: model.InstallTemplate.Helm.Spec.Components.Istio.MaxWorkloadCertTtl.Seconds.ValueInt64(),
								},
								PilotOverlays: func() []*v1alpha12.K8SObjectOverlay {
									a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.Istio.PilotOverlays))
									for i, pilot_overlays := range model.InstallTemplate.Helm.Spec.Components.Istio.PilotOverlays {
										a[i] = &v1alpha12.K8SObjectOverlay{
											ApiVersion: pilot_overlays.ApiVersion.ValueString(),
											Kind:       pilot_overlays.Kind.ValueString(),
											Name:       pilot_overlays.Name.ValueString(),
											Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
												a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(pilot_overlays.Patches))
												for i, patches := range pilot_overlays.Patches {
													a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
														Path: patches.Path.ValueString(),
														Value: &structpb.Value{
															BoolValue: patches.Value.BoolValue.ValueBool(),
															ListValue: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
															NumberValue: patches.Value.NumberValue.ValueFloat64(),
															StringValue: patches.Value.StringValue.ValueString(),
															StructValue: func() map[string]string {
																tmp := make(map[string]string)
																resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
								TraceSamplingRate: model.InstallTemplate.Helm.Spec.Components.Istio.TraceSamplingRate.ValueFloat64(),
								TrustDomain:       model.InstallTemplate.Helm.Spec.Components.Istio.TrustDomain.ValueString(),
								TsbVersion:        model.InstallTemplate.Helm.Spec.Components.Istio.TsbVersion.ValueString(),
							},
							Ngac: &v1alpha13.NGAC{
								Enabled: model.InstallTemplate.Helm.Spec.Components.Ngac.Enabled.ValueBool(),
								KubeSpec: &kubernetes.KubernetesComponentSpec{
									Deployment: &kubernetes.Deployment{
										Affinity: &kubernetes.Affinity{
											NodeAffinity: &kubernetes.NodeAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
													a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PreferredSchedulingTerm{
															Preference: &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																	for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
													a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
													for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
														a[i] = &kubernetes.NodeSelectorTerm{
															MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																for i, match_expressions := range node_selector_terms.MatchExpressions {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_expressions.Key.ValueString(),
																		Operator: match_expressions.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																for i, match_fields := range node_selector_terms.MatchFields {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_fields.Key.ValueString(),
																		Operator: match_fields.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_fields.Values.Elements()))
																			resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
														}
													}
													return a
												}()},
											},
											PodAffinity: &kubernetes.PodAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
											PodAntiAffinity: &kubernetes.PodAntiAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
										},
										ContainerSecurityContext: &kubernetes.SecurityContext{
											AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
											Capabilities: &kubernetes.Capabilities{
												Add: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Drop: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
											},
											Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
											ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
											ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
											RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
											},
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										Env: func() []*kubernetes.EnvVar {
											a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Env))
											for i, env := range model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Env {
												a[i] = &kubernetes.EnvVar{
													Name:  env.Name.ValueString(),
													Value: env.Value.ValueString(),
													ValueFrom: &kubernetes.EnvVarSource{
														ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
															Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
														},
														FieldRef: &kubernetes.ObjectFieldSelector{
															ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
															FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
														},
														ResourceFieldRef: &kubernetes.ResourceFieldSelector{
															ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
															Divisor: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
															},
															Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
														},
														SecretKeyRef: &kubernetes.SecretKeySelector{
															Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
														},
													},
												}
											}
											return a
										}(),
										HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
											MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
											Metrics: func() []*kubernetes.MetricSpec {
												a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.HpaSpec.Metrics))
												for i, metrics := range model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.HpaSpec.Metrics {
													a[i] = &kubernetes.MetricSpec{
														External: &kubernetes.ExternalMetricSource{
															MetricName: metrics.External.MetricName.ValueString(),
															MetricSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																	for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetValue.Type.ValueInt64(),
															},
														},
														Object: &kubernetes.ObjectMetricSource{
															AverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.AverageValue.Type.ValueInt64(),
															},
															MetricName: metrics.Object.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Target: &kubernetes.CrossVersionObjectReference{
																ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																Kind:       metrics.Object.Target.Kind.ValueString(),
																Name:       metrics.Object.Target.Name.ValueString(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.TargetValue.Type.ValueInt64(),
															},
														},
														Pods: &kubernetes.PodsMetricSource{
															MetricName: metrics.Pods.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Resource: &kubernetes.ResourceMetricSource{
															Name: metrics.Resource.Name.ValueString(),
															Target: &kubernetes.MetricTarget{
																AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																AverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																},
																Type: metrics.Resource.Target.Type.ValueString(),
																Value: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																},
															},
															TargetAverageUtilization: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Type: metrics.Type.ValueString(),
													}
												}
												return a
											}(),
											MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
										},
										PodAnnotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										PodSecurityContext: &kubernetes.PodSecurityContext{
											FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
											FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
											RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
											},
											SupplementalGroups: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Sysctls: func() []*kubernetes.Sysctl {
												a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.Sysctls))
												for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.Sysctls {
													a[i] = &kubernetes.Sysctl{
														Name:  sysctls.Name.ValueString(),
														Value: sysctls.Value.ValueString(),
													}
												}
												return a
											}(),
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
										Resources: &kubernetes.Resources{
											Limits: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Requests: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
										},
										Strategy: &kubernetes.DeploymentStrategy{
											RollingUpdate: &kubernetes.RollingUpdateDeployment{
												MaxSurge: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
												},
												MaxUnavailable: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
												},
											},
											Type: model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Strategy.Type.ValueString(),
										},
										Tolerations: func() []*v11.Toleration {
											a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Tolerations))
											for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Deployment.Tolerations {
												a[i] = &v11.Toleration{
													Effect:            ptrify(tolerations.Effect.ValueString()),
													Key:               ptrify(tolerations.Key.ValueString()),
													Operator:          ptrify(tolerations.Operator.ValueString()),
													TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
													Value:             ptrify(tolerations.Value.ValueString()),
												}
											}
											return a
										}(),
									},
									Overlays: func() []*v1alpha12.K8SObjectOverlay {
										a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Overlays))
										for i, overlays := range model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Overlays {
											a[i] = &v1alpha12.K8SObjectOverlay{
												ApiVersion: overlays.ApiVersion.ValueString(),
												Kind:       overlays.Kind.ValueString(),
												Name:       overlays.Name.ValueString(),
												Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
													a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
													for i, patches := range overlays.Patches {
														a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
															Path: patches.Path.ValueString(),
															Value: &structpb.Value{
																BoolValue: patches.Value.BoolValue.ValueBool(),
																ListValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																NumberValue: patches.Value.NumberValue.ValueFloat64(),
																StringValue: patches.Value.StringValue.ValueString(),
																StructValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
									Service: &kubernetes.Service{
										Annotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Labels: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Ports: func() []*kubernetes.ServicePort {
											a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Service.Ports))
											for i, ports := range model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Service.Ports {
												a[i] = &kubernetes.ServicePort{
													Name:     ports.Name.ValueString(),
													NodePort: int32(ports.NodePort.ValueInt64()),
													Port:     int32(ports.Port.ValueInt64()),
													Protocol: ports.Protocol.ValueString(),
													TargetPort: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
														Type:   ports.TargetPort.Type.ValueInt64(),
													},
												}
											}
											return a
										}(),
										Type: model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.Service.Type.ValueString(),
									},
									ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
										a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.ServiceAccount.ImagePullSecrets))
										for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.Ngac.KubeSpec.ServiceAccount.ImagePullSecrets {
											a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
										}
										return a
									}()},
								},
								LogLevels: func() map[string]string {
									tmp := make(map[string]string)
									resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Ngac.LogLevels.ElementsAs(ctx, &tmp, false)...)
									return tmp
								}(),
							},
							Oap: &v1alpha13.Oap{
								KubeSpec: &kubernetes.KubernetesComponentSpec{
									Deployment: &kubernetes.Deployment{
										Affinity: &kubernetes.Affinity{
											NodeAffinity: &kubernetes.NodeAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
													a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PreferredSchedulingTerm{
															Preference: &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																	for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
													a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
													for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
														a[i] = &kubernetes.NodeSelectorTerm{
															MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																for i, match_expressions := range node_selector_terms.MatchExpressions {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_expressions.Key.ValueString(),
																		Operator: match_expressions.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																for i, match_fields := range node_selector_terms.MatchFields {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_fields.Key.ValueString(),
																		Operator: match_fields.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_fields.Values.Elements()))
																			resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
														}
													}
													return a
												}()},
											},
											PodAffinity: &kubernetes.PodAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
											PodAntiAffinity: &kubernetes.PodAntiAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
										},
										ContainerSecurityContext: &kubernetes.SecurityContext{
											AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
											Capabilities: &kubernetes.Capabilities{
												Add: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Drop: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
											},
											Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
											ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
											ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
											RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
											},
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										Env: func() []*kubernetes.EnvVar {
											a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Env))
											for i, env := range model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Env {
												a[i] = &kubernetes.EnvVar{
													Name:  env.Name.ValueString(),
													Value: env.Value.ValueString(),
													ValueFrom: &kubernetes.EnvVarSource{
														ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
															Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
														},
														FieldRef: &kubernetes.ObjectFieldSelector{
															ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
															FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
														},
														ResourceFieldRef: &kubernetes.ResourceFieldSelector{
															ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
															Divisor: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
															},
															Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
														},
														SecretKeyRef: &kubernetes.SecretKeySelector{
															Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
														},
													},
												}
											}
											return a
										}(),
										HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
											MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
											Metrics: func() []*kubernetes.MetricSpec {
												a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.HpaSpec.Metrics))
												for i, metrics := range model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.HpaSpec.Metrics {
													a[i] = &kubernetes.MetricSpec{
														External: &kubernetes.ExternalMetricSource{
															MetricName: metrics.External.MetricName.ValueString(),
															MetricSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																	for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetValue.Type.ValueInt64(),
															},
														},
														Object: &kubernetes.ObjectMetricSource{
															AverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.AverageValue.Type.ValueInt64(),
															},
															MetricName: metrics.Object.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Target: &kubernetes.CrossVersionObjectReference{
																ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																Kind:       metrics.Object.Target.Kind.ValueString(),
																Name:       metrics.Object.Target.Name.ValueString(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.TargetValue.Type.ValueInt64(),
															},
														},
														Pods: &kubernetes.PodsMetricSource{
															MetricName: metrics.Pods.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Resource: &kubernetes.ResourceMetricSource{
															Name: metrics.Resource.Name.ValueString(),
															Target: &kubernetes.MetricTarget{
																AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																AverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																},
																Type: metrics.Resource.Target.Type.ValueString(),
																Value: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																},
															},
															TargetAverageUtilization: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Type: metrics.Type.ValueString(),
													}
												}
												return a
											}(),
											MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
										},
										PodAnnotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										PodSecurityContext: &kubernetes.PodSecurityContext{
											FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
											FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
											RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
											},
											SupplementalGroups: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Sysctls: func() []*kubernetes.Sysctl {
												a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.Sysctls))
												for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.Sysctls {
													a[i] = &kubernetes.Sysctl{
														Name:  sysctls.Name.ValueString(),
														Value: sysctls.Value.ValueString(),
													}
												}
												return a
											}(),
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
										Resources: &kubernetes.Resources{
											Limits: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Requests: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
										},
										Strategy: &kubernetes.DeploymentStrategy{
											RollingUpdate: &kubernetes.RollingUpdateDeployment{
												MaxSurge: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
												},
												MaxUnavailable: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
												},
											},
											Type: model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Strategy.Type.ValueString(),
										},
										Tolerations: func() []*v11.Toleration {
											a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Tolerations))
											for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Deployment.Tolerations {
												a[i] = &v11.Toleration{
													Effect:            ptrify(tolerations.Effect.ValueString()),
													Key:               ptrify(tolerations.Key.ValueString()),
													Operator:          ptrify(tolerations.Operator.ValueString()),
													TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
													Value:             ptrify(tolerations.Value.ValueString()),
												}
											}
											return a
										}(),
									},
									Overlays: func() []*v1alpha12.K8SObjectOverlay {
										a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Overlays))
										for i, overlays := range model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Overlays {
											a[i] = &v1alpha12.K8SObjectOverlay{
												ApiVersion: overlays.ApiVersion.ValueString(),
												Kind:       overlays.Kind.ValueString(),
												Name:       overlays.Name.ValueString(),
												Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
													a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
													for i, patches := range overlays.Patches {
														a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
															Path: patches.Path.ValueString(),
															Value: &structpb.Value{
																BoolValue: patches.Value.BoolValue.ValueBool(),
																ListValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																NumberValue: patches.Value.NumberValue.ValueFloat64(),
																StringValue: patches.Value.StringValue.ValueString(),
																StructValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
									Service: &kubernetes.Service{
										Annotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Labels: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Ports: func() []*kubernetes.ServicePort {
											a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Service.Ports))
											for i, ports := range model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Service.Ports {
												a[i] = &kubernetes.ServicePort{
													Name:     ports.Name.ValueString(),
													NodePort: int32(ports.NodePort.ValueInt64()),
													Port:     int32(ports.Port.ValueInt64()),
													Protocol: ports.Protocol.ValueString(),
													TargetPort: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
														Type:   ports.TargetPort.Type.ValueInt64(),
													},
												}
											}
											return a
										}(),
										Type: model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.Service.Type.ValueString(),
									},
									ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
										a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.ServiceAccount.ImagePullSecrets))
										for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.Oap.KubeSpec.ServiceAccount.ImagePullSecrets {
											a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
										}
										return a
									}()},
								},
								OnDemandEnvoyMetricsEnabled: model.InstallTemplate.Helm.Spec.Components.Oap.OnDemandEnvoyMetricsEnabled.ValueBool(),
								StorageIndexMergingEnabled:  model.InstallTemplate.Helm.Spec.Components.Oap.StorageIndexMergingEnabled.ValueBool(),
								StreamingLogEnabled:         model.InstallTemplate.Helm.Spec.Components.Oap.StreamingLogEnabled.ValueBool(),
							},
							Onboarding: &v1alpha13.Onboarding{
								Operator: &v1alpha13.OnboardingOperator{KubeSpec: &kubernetes.KubernetesComponentSpec{
									Deployment: &kubernetes.Deployment{
										Affinity: &kubernetes.Affinity{
											NodeAffinity: &kubernetes.NodeAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
													a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PreferredSchedulingTerm{
															Preference: &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																	for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
													a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
													for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
														a[i] = &kubernetes.NodeSelectorTerm{
															MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																for i, match_expressions := range node_selector_terms.MatchExpressions {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_expressions.Key.ValueString(),
																		Operator: match_expressions.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																for i, match_fields := range node_selector_terms.MatchFields {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_fields.Key.ValueString(),
																		Operator: match_fields.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_fields.Values.Elements()))
																			resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
														}
													}
													return a
												}()},
											},
											PodAffinity: &kubernetes.PodAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
											PodAntiAffinity: &kubernetes.PodAntiAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
										},
										ContainerSecurityContext: &kubernetes.SecurityContext{
											AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
											Capabilities: &kubernetes.Capabilities{
												Add: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Drop: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
											},
											Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
											ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
											ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
											RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
											},
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										Env: func() []*kubernetes.EnvVar {
											a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Env))
											for i, env := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Env {
												a[i] = &kubernetes.EnvVar{
													Name:  env.Name.ValueString(),
													Value: env.Value.ValueString(),
													ValueFrom: &kubernetes.EnvVarSource{
														ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
															Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
														},
														FieldRef: &kubernetes.ObjectFieldSelector{
															ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
															FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
														},
														ResourceFieldRef: &kubernetes.ResourceFieldSelector{
															ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
															Divisor: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
															},
															Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
														},
														SecretKeyRef: &kubernetes.SecretKeySelector{
															Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
														},
													},
												}
											}
											return a
										}(),
										HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
											MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
											Metrics: func() []*kubernetes.MetricSpec {
												a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.HpaSpec.Metrics))
												for i, metrics := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.HpaSpec.Metrics {
													a[i] = &kubernetes.MetricSpec{
														External: &kubernetes.ExternalMetricSource{
															MetricName: metrics.External.MetricName.ValueString(),
															MetricSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																	for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetValue.Type.ValueInt64(),
															},
														},
														Object: &kubernetes.ObjectMetricSource{
															AverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.AverageValue.Type.ValueInt64(),
															},
															MetricName: metrics.Object.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Target: &kubernetes.CrossVersionObjectReference{
																ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																Kind:       metrics.Object.Target.Kind.ValueString(),
																Name:       metrics.Object.Target.Name.ValueString(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.TargetValue.Type.ValueInt64(),
															},
														},
														Pods: &kubernetes.PodsMetricSource{
															MetricName: metrics.Pods.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Resource: &kubernetes.ResourceMetricSource{
															Name: metrics.Resource.Name.ValueString(),
															Target: &kubernetes.MetricTarget{
																AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																AverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																},
																Type: metrics.Resource.Target.Type.ValueString(),
																Value: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																},
															},
															TargetAverageUtilization: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Type: metrics.Type.ValueString(),
													}
												}
												return a
											}(),
											MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
										},
										PodAnnotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										PodSecurityContext: &kubernetes.PodSecurityContext{
											FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
											FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
											RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
											},
											SupplementalGroups: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Sysctls: func() []*kubernetes.Sysctl {
												a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.Sysctls))
												for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.Sysctls {
													a[i] = &kubernetes.Sysctl{
														Name:  sysctls.Name.ValueString(),
														Value: sysctls.Value.ValueString(),
													}
												}
												return a
											}(),
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
										Resources: &kubernetes.Resources{
											Limits: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Requests: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
										},
										Strategy: &kubernetes.DeploymentStrategy{
											RollingUpdate: &kubernetes.RollingUpdateDeployment{
												MaxSurge: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
												},
												MaxUnavailable: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
												},
											},
											Type: model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Strategy.Type.ValueString(),
										},
										Tolerations: func() []*v11.Toleration {
											a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Tolerations))
											for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Deployment.Tolerations {
												a[i] = &v11.Toleration{
													Effect:            ptrify(tolerations.Effect.ValueString()),
													Key:               ptrify(tolerations.Key.ValueString()),
													Operator:          ptrify(tolerations.Operator.ValueString()),
													TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
													Value:             ptrify(tolerations.Value.ValueString()),
												}
											}
											return a
										}(),
									},
									Overlays: func() []*v1alpha12.K8SObjectOverlay {
										a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Overlays))
										for i, overlays := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Overlays {
											a[i] = &v1alpha12.K8SObjectOverlay{
												ApiVersion: overlays.ApiVersion.ValueString(),
												Kind:       overlays.Kind.ValueString(),
												Name:       overlays.Name.ValueString(),
												Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
													a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
													for i, patches := range overlays.Patches {
														a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
															Path: patches.Path.ValueString(),
															Value: &structpb.Value{
																BoolValue: patches.Value.BoolValue.ValueBool(),
																ListValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																NumberValue: patches.Value.NumberValue.ValueFloat64(),
																StringValue: patches.Value.StringValue.ValueString(),
																StructValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
									Service: &kubernetes.Service{
										Annotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Labels: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Ports: func() []*kubernetes.ServicePort {
											a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Service.Ports))
											for i, ports := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Service.Ports {
												a[i] = &kubernetes.ServicePort{
													Name:     ports.Name.ValueString(),
													NodePort: int32(ports.NodePort.ValueInt64()),
													Port:     int32(ports.Port.ValueInt64()),
													Protocol: ports.Protocol.ValueString(),
													TargetPort: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
														Type:   ports.TargetPort.Type.ValueInt64(),
													},
												}
											}
											return a
										}(),
										Type: model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.Service.Type.ValueString(),
									},
									ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
										a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.ServiceAccount.ImagePullSecrets))
										for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Operator.KubeSpec.ServiceAccount.ImagePullSecrets {
											a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
										}
										return a
									}()},
								}},
								Plane: &v1alpha13.OnboardingPlane{Instance: &v1alpha14.OnboardingPlaneInstance{KubeSpec: &kubernetes.KubernetesComponentSpec{
									Deployment: &kubernetes.Deployment{
										Affinity: &kubernetes.Affinity{
											NodeAffinity: &kubernetes.NodeAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
													a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PreferredSchedulingTerm{
															Preference: &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																	for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
													a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
													for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
														a[i] = &kubernetes.NodeSelectorTerm{
															MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																for i, match_expressions := range node_selector_terms.MatchExpressions {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_expressions.Key.ValueString(),
																		Operator: match_expressions.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																for i, match_fields := range node_selector_terms.MatchFields {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_fields.Key.ValueString(),
																		Operator: match_fields.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_fields.Values.Elements()))
																			resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
														}
													}
													return a
												}()},
											},
											PodAffinity: &kubernetes.PodAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
											PodAntiAffinity: &kubernetes.PodAntiAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
										},
										ContainerSecurityContext: &kubernetes.SecurityContext{
											AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
											Capabilities: &kubernetes.Capabilities{
												Add: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Drop: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
											},
											Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
											ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
											ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
											RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
											},
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										Env: func() []*kubernetes.EnvVar {
											a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Env))
											for i, env := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Env {
												a[i] = &kubernetes.EnvVar{
													Name:  env.Name.ValueString(),
													Value: env.Value.ValueString(),
													ValueFrom: &kubernetes.EnvVarSource{
														ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
															Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
														},
														FieldRef: &kubernetes.ObjectFieldSelector{
															ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
															FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
														},
														ResourceFieldRef: &kubernetes.ResourceFieldSelector{
															ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
															Divisor: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
															},
															Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
														},
														SecretKeyRef: &kubernetes.SecretKeySelector{
															Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
														},
													},
												}
											}
											return a
										}(),
										HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
											MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
											Metrics: func() []*kubernetes.MetricSpec {
												a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.HpaSpec.Metrics))
												for i, metrics := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.HpaSpec.Metrics {
													a[i] = &kubernetes.MetricSpec{
														External: &kubernetes.ExternalMetricSource{
															MetricName: metrics.External.MetricName.ValueString(),
															MetricSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																	for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetValue.Type.ValueInt64(),
															},
														},
														Object: &kubernetes.ObjectMetricSource{
															AverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.AverageValue.Type.ValueInt64(),
															},
															MetricName: metrics.Object.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Target: &kubernetes.CrossVersionObjectReference{
																ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																Kind:       metrics.Object.Target.Kind.ValueString(),
																Name:       metrics.Object.Target.Name.ValueString(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.TargetValue.Type.ValueInt64(),
															},
														},
														Pods: &kubernetes.PodsMetricSource{
															MetricName: metrics.Pods.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Resource: &kubernetes.ResourceMetricSource{
															Name: metrics.Resource.Name.ValueString(),
															Target: &kubernetes.MetricTarget{
																AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																AverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																},
																Type: metrics.Resource.Target.Type.ValueString(),
																Value: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																},
															},
															TargetAverageUtilization: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Type: metrics.Type.ValueString(),
													}
												}
												return a
											}(),
											MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
										},
										PodAnnotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										PodSecurityContext: &kubernetes.PodSecurityContext{
											FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
											FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
											RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
											},
											SupplementalGroups: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Sysctls: func() []*kubernetes.Sysctl {
												a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.Sysctls))
												for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.Sysctls {
													a[i] = &kubernetes.Sysctl{
														Name:  sysctls.Name.ValueString(),
														Value: sysctls.Value.ValueString(),
													}
												}
												return a
											}(),
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
										Resources: &kubernetes.Resources{
											Limits: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Requests: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
										},
										Strategy: &kubernetes.DeploymentStrategy{
											RollingUpdate: &kubernetes.RollingUpdateDeployment{
												MaxSurge: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
												},
												MaxUnavailable: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
												},
											},
											Type: model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Strategy.Type.ValueString(),
										},
										Tolerations: func() []*v11.Toleration {
											a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Tolerations))
											for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Deployment.Tolerations {
												a[i] = &v11.Toleration{
													Effect:            ptrify(tolerations.Effect.ValueString()),
													Key:               ptrify(tolerations.Key.ValueString()),
													Operator:          ptrify(tolerations.Operator.ValueString()),
													TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
													Value:             ptrify(tolerations.Value.ValueString()),
												}
											}
											return a
										}(),
									},
									Overlays: func() []*v1alpha12.K8SObjectOverlay {
										a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Overlays))
										for i, overlays := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Overlays {
											a[i] = &v1alpha12.K8SObjectOverlay{
												ApiVersion: overlays.ApiVersion.ValueString(),
												Kind:       overlays.Kind.ValueString(),
												Name:       overlays.Name.ValueString(),
												Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
													a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
													for i, patches := range overlays.Patches {
														a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
															Path: patches.Path.ValueString(),
															Value: &structpb.Value{
																BoolValue: patches.Value.BoolValue.ValueBool(),
																ListValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																NumberValue: patches.Value.NumberValue.ValueFloat64(),
																StringValue: patches.Value.StringValue.ValueString(),
																StructValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
									Service: &kubernetes.Service{
										Annotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Labels: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Ports: func() []*kubernetes.ServicePort {
											a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Service.Ports))
											for i, ports := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Service.Ports {
												a[i] = &kubernetes.ServicePort{
													Name:     ports.Name.ValueString(),
													NodePort: int32(ports.NodePort.ValueInt64()),
													Port:     int32(ports.Port.ValueInt64()),
													Protocol: ports.Protocol.ValueString(),
													TargetPort: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
														Type:   ports.TargetPort.Type.ValueInt64(),
													},
												}
											}
											return a
										}(),
										Type: model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.Service.Type.ValueString(),
									},
									ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
										a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.ServiceAccount.ImagePullSecrets))
										for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Plane.Instance.KubeSpec.ServiceAccount.ImagePullSecrets {
											a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
										}
										return a
									}()},
								}}},
								Repository: &v1alpha13.OnboardingRepository{KubeSpec: &kubernetes.KubernetesComponentSpec{
									Deployment: &kubernetes.Deployment{
										Affinity: &kubernetes.Affinity{
											NodeAffinity: &kubernetes.NodeAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
													a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PreferredSchedulingTerm{
															Preference: &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																	for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
													a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
													for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
														a[i] = &kubernetes.NodeSelectorTerm{
															MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																for i, match_expressions := range node_selector_terms.MatchExpressions {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_expressions.Key.ValueString(),
																		Operator: match_expressions.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																for i, match_fields := range node_selector_terms.MatchFields {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_fields.Key.ValueString(),
																		Operator: match_fields.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_fields.Values.Elements()))
																			resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
														}
													}
													return a
												}()},
											},
											PodAffinity: &kubernetes.PodAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
											PodAntiAffinity: &kubernetes.PodAntiAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
										},
										ContainerSecurityContext: &kubernetes.SecurityContext{
											AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
											Capabilities: &kubernetes.Capabilities{
												Add: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Drop: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
											},
											Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
											ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
											ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
											RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
											},
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										Env: func() []*kubernetes.EnvVar {
											a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Env))
											for i, env := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Env {
												a[i] = &kubernetes.EnvVar{
													Name:  env.Name.ValueString(),
													Value: env.Value.ValueString(),
													ValueFrom: &kubernetes.EnvVarSource{
														ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
															Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
														},
														FieldRef: &kubernetes.ObjectFieldSelector{
															ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
															FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
														},
														ResourceFieldRef: &kubernetes.ResourceFieldSelector{
															ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
															Divisor: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
															},
															Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
														},
														SecretKeyRef: &kubernetes.SecretKeySelector{
															Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
														},
													},
												}
											}
											return a
										}(),
										HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
											MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
											Metrics: func() []*kubernetes.MetricSpec {
												a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.HpaSpec.Metrics))
												for i, metrics := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.HpaSpec.Metrics {
													a[i] = &kubernetes.MetricSpec{
														External: &kubernetes.ExternalMetricSource{
															MetricName: metrics.External.MetricName.ValueString(),
															MetricSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																	for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetValue.Type.ValueInt64(),
															},
														},
														Object: &kubernetes.ObjectMetricSource{
															AverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.AverageValue.Type.ValueInt64(),
															},
															MetricName: metrics.Object.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Target: &kubernetes.CrossVersionObjectReference{
																ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																Kind:       metrics.Object.Target.Kind.ValueString(),
																Name:       metrics.Object.Target.Name.ValueString(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.TargetValue.Type.ValueInt64(),
															},
														},
														Pods: &kubernetes.PodsMetricSource{
															MetricName: metrics.Pods.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Resource: &kubernetes.ResourceMetricSource{
															Name: metrics.Resource.Name.ValueString(),
															Target: &kubernetes.MetricTarget{
																AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																AverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																},
																Type: metrics.Resource.Target.Type.ValueString(),
																Value: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																},
															},
															TargetAverageUtilization: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Type: metrics.Type.ValueString(),
													}
												}
												return a
											}(),
											MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
										},
										PodAnnotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										PodSecurityContext: &kubernetes.PodSecurityContext{
											FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
											FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
											RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
											},
											SupplementalGroups: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Sysctls: func() []*kubernetes.Sysctl {
												a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.Sysctls))
												for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.Sysctls {
													a[i] = &kubernetes.Sysctl{
														Name:  sysctls.Name.ValueString(),
														Value: sysctls.Value.ValueString(),
													}
												}
												return a
											}(),
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
										Resources: &kubernetes.Resources{
											Limits: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Requests: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
										},
										Strategy: &kubernetes.DeploymentStrategy{
											RollingUpdate: &kubernetes.RollingUpdateDeployment{
												MaxSurge: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
												},
												MaxUnavailable: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
												},
											},
											Type: model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Strategy.Type.ValueString(),
										},
										Tolerations: func() []*v11.Toleration {
											a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Tolerations))
											for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Deployment.Tolerations {
												a[i] = &v11.Toleration{
													Effect:            ptrify(tolerations.Effect.ValueString()),
													Key:               ptrify(tolerations.Key.ValueString()),
													Operator:          ptrify(tolerations.Operator.ValueString()),
													TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
													Value:             ptrify(tolerations.Value.ValueString()),
												}
											}
											return a
										}(),
									},
									Overlays: func() []*v1alpha12.K8SObjectOverlay {
										a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Overlays))
										for i, overlays := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Overlays {
											a[i] = &v1alpha12.K8SObjectOverlay{
												ApiVersion: overlays.ApiVersion.ValueString(),
												Kind:       overlays.Kind.ValueString(),
												Name:       overlays.Name.ValueString(),
												Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
													a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
													for i, patches := range overlays.Patches {
														a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
															Path: patches.Path.ValueString(),
															Value: &structpb.Value{
																BoolValue: patches.Value.BoolValue.ValueBool(),
																ListValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																NumberValue: patches.Value.NumberValue.ValueFloat64(),
																StringValue: patches.Value.StringValue.ValueString(),
																StructValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
									Service: &kubernetes.Service{
										Annotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Labels: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Ports: func() []*kubernetes.ServicePort {
											a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Service.Ports))
											for i, ports := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Service.Ports {
												a[i] = &kubernetes.ServicePort{
													Name:     ports.Name.ValueString(),
													NodePort: int32(ports.NodePort.ValueInt64()),
													Port:     int32(ports.Port.ValueInt64()),
													Protocol: ports.Protocol.ValueString(),
													TargetPort: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
														Type:   ports.TargetPort.Type.ValueInt64(),
													},
												}
											}
											return a
										}(),
										Type: model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.Service.Type.ValueString(),
									},
									ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
										a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.ServiceAccount.ImagePullSecrets))
										for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.Onboarding.Repository.KubeSpec.ServiceAccount.ImagePullSecrets {
											a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
										}
										return a
									}()},
								}},
							},
							RateLimitServer: &v1alpha13.RateLimitServer{
								Backend: &v1alpha13.RateLimitServer_Backend{Redis: &v1alpha13.RateLimitServer_Backend_RedisSettings{Uri: model.InstallTemplate.Helm.Spec.Components.RateLimitServer.Backend.Redis.Uri.ValueString()}},
								Domain:  model.InstallTemplate.Helm.Spec.Components.RateLimitServer.Domain.ValueString(),
								KubeSpec: &kubernetes.KubernetesComponentSpec{
									Deployment: &kubernetes.Deployment{
										Affinity: &kubernetes.Affinity{
											NodeAffinity: &kubernetes.NodeAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
													a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PreferredSchedulingTerm{
															Preference: &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																	for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
													a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
													for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
														a[i] = &kubernetes.NodeSelectorTerm{
															MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																for i, match_expressions := range node_selector_terms.MatchExpressions {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_expressions.Key.ValueString(),
																		Operator: match_expressions.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																for i, match_fields := range node_selector_terms.MatchFields {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_fields.Key.ValueString(),
																		Operator: match_fields.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_fields.Values.Elements()))
																			resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
														}
													}
													return a
												}()},
											},
											PodAffinity: &kubernetes.PodAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
											PodAntiAffinity: &kubernetes.PodAntiAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
										},
										ContainerSecurityContext: &kubernetes.SecurityContext{
											AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
											Capabilities: &kubernetes.Capabilities{
												Add: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Drop: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
											},
											Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
											ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
											ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
											RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
											},
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										Env: func() []*kubernetes.EnvVar {
											a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Env))
											for i, env := range model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Env {
												a[i] = &kubernetes.EnvVar{
													Name:  env.Name.ValueString(),
													Value: env.Value.ValueString(),
													ValueFrom: &kubernetes.EnvVarSource{
														ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
															Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
														},
														FieldRef: &kubernetes.ObjectFieldSelector{
															ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
															FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
														},
														ResourceFieldRef: &kubernetes.ResourceFieldSelector{
															ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
															Divisor: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
															},
															Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
														},
														SecretKeyRef: &kubernetes.SecretKeySelector{
															Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
														},
													},
												}
											}
											return a
										}(),
										HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
											MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
											Metrics: func() []*kubernetes.MetricSpec {
												a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.HpaSpec.Metrics))
												for i, metrics := range model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.HpaSpec.Metrics {
													a[i] = &kubernetes.MetricSpec{
														External: &kubernetes.ExternalMetricSource{
															MetricName: metrics.External.MetricName.ValueString(),
															MetricSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																	for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetValue.Type.ValueInt64(),
															},
														},
														Object: &kubernetes.ObjectMetricSource{
															AverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.AverageValue.Type.ValueInt64(),
															},
															MetricName: metrics.Object.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Target: &kubernetes.CrossVersionObjectReference{
																ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																Kind:       metrics.Object.Target.Kind.ValueString(),
																Name:       metrics.Object.Target.Name.ValueString(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.TargetValue.Type.ValueInt64(),
															},
														},
														Pods: &kubernetes.PodsMetricSource{
															MetricName: metrics.Pods.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Resource: &kubernetes.ResourceMetricSource{
															Name: metrics.Resource.Name.ValueString(),
															Target: &kubernetes.MetricTarget{
																AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																AverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																},
																Type: metrics.Resource.Target.Type.ValueString(),
																Value: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																},
															},
															TargetAverageUtilization: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Type: metrics.Type.ValueString(),
													}
												}
												return a
											}(),
											MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
										},
										PodAnnotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										PodSecurityContext: &kubernetes.PodSecurityContext{
											FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
											FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
											RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
											},
											SupplementalGroups: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Sysctls: func() []*kubernetes.Sysctl {
												a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.Sysctls))
												for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.Sysctls {
													a[i] = &kubernetes.Sysctl{
														Name:  sysctls.Name.ValueString(),
														Value: sysctls.Value.ValueString(),
													}
												}
												return a
											}(),
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
										Resources: &kubernetes.Resources{
											Limits: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Requests: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
										},
										Strategy: &kubernetes.DeploymentStrategy{
											RollingUpdate: &kubernetes.RollingUpdateDeployment{
												MaxSurge: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
												},
												MaxUnavailable: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
												},
											},
											Type: model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Strategy.Type.ValueString(),
										},
										Tolerations: func() []*v11.Toleration {
											a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Tolerations))
											for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Deployment.Tolerations {
												a[i] = &v11.Toleration{
													Effect:            ptrify(tolerations.Effect.ValueString()),
													Key:               ptrify(tolerations.Key.ValueString()),
													Operator:          ptrify(tolerations.Operator.ValueString()),
													TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
													Value:             ptrify(tolerations.Value.ValueString()),
												}
											}
											return a
										}(),
									},
									Overlays: func() []*v1alpha12.K8SObjectOverlay {
										a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Overlays))
										for i, overlays := range model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Overlays {
											a[i] = &v1alpha12.K8SObjectOverlay{
												ApiVersion: overlays.ApiVersion.ValueString(),
												Kind:       overlays.Kind.ValueString(),
												Name:       overlays.Name.ValueString(),
												Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
													a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
													for i, patches := range overlays.Patches {
														a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
															Path: patches.Path.ValueString(),
															Value: &structpb.Value{
																BoolValue: patches.Value.BoolValue.ValueBool(),
																ListValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																NumberValue: patches.Value.NumberValue.ValueFloat64(),
																StringValue: patches.Value.StringValue.ValueString(),
																StructValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
									Service: &kubernetes.Service{
										Annotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Labels: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Ports: func() []*kubernetes.ServicePort {
											a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Service.Ports))
											for i, ports := range model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Service.Ports {
												a[i] = &kubernetes.ServicePort{
													Name:     ports.Name.ValueString(),
													NodePort: int32(ports.NodePort.ValueInt64()),
													Port:     int32(ports.Port.ValueInt64()),
													Protocol: ports.Protocol.ValueString(),
													TargetPort: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
														Type:   ports.TargetPort.Type.ValueInt64(),
													},
												}
											}
											return a
										}(),
										Type: model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.Service.Type.ValueString(),
									},
									ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
										a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.ServiceAccount.ImagePullSecrets))
										for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.RateLimitServer.KubeSpec.ServiceAccount.ImagePullSecrets {
											a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
										}
										return a
									}()},
								},
							},
							Satellite: &v1alpha13.Satellite{
								Enabled: model.InstallTemplate.Helm.Spec.Components.Satellite.Enabled.ValueBool(),
								KubeSpec: &kubernetes.KubernetesComponentSpec{
									Deployment: &kubernetes.Deployment{
										Affinity: &kubernetes.Affinity{
											NodeAffinity: &kubernetes.NodeAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
													a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PreferredSchedulingTerm{
															Preference: &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																	for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
													a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
													for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
														a[i] = &kubernetes.NodeSelectorTerm{
															MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																for i, match_expressions := range node_selector_terms.MatchExpressions {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_expressions.Key.ValueString(),
																		Operator: match_expressions.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																for i, match_fields := range node_selector_terms.MatchFields {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_fields.Key.ValueString(),
																		Operator: match_fields.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_fields.Values.Elements()))
																			resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
														}
													}
													return a
												}()},
											},
											PodAffinity: &kubernetes.PodAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
											PodAntiAffinity: &kubernetes.PodAntiAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
										},
										ContainerSecurityContext: &kubernetes.SecurityContext{
											AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
											Capabilities: &kubernetes.Capabilities{
												Add: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Drop: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
											},
											Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
											ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
											ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
											RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
											},
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										Env: func() []*kubernetes.EnvVar {
											a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Env))
											for i, env := range model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Env {
												a[i] = &kubernetes.EnvVar{
													Name:  env.Name.ValueString(),
													Value: env.Value.ValueString(),
													ValueFrom: &kubernetes.EnvVarSource{
														ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
															Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
														},
														FieldRef: &kubernetes.ObjectFieldSelector{
															ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
															FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
														},
														ResourceFieldRef: &kubernetes.ResourceFieldSelector{
															ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
															Divisor: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
															},
															Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
														},
														SecretKeyRef: &kubernetes.SecretKeySelector{
															Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
														},
													},
												}
											}
											return a
										}(),
										HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
											MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
											Metrics: func() []*kubernetes.MetricSpec {
												a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.HpaSpec.Metrics))
												for i, metrics := range model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.HpaSpec.Metrics {
													a[i] = &kubernetes.MetricSpec{
														External: &kubernetes.ExternalMetricSource{
															MetricName: metrics.External.MetricName.ValueString(),
															MetricSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																	for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetValue.Type.ValueInt64(),
															},
														},
														Object: &kubernetes.ObjectMetricSource{
															AverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.AverageValue.Type.ValueInt64(),
															},
															MetricName: metrics.Object.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Target: &kubernetes.CrossVersionObjectReference{
																ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																Kind:       metrics.Object.Target.Kind.ValueString(),
																Name:       metrics.Object.Target.Name.ValueString(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.TargetValue.Type.ValueInt64(),
															},
														},
														Pods: &kubernetes.PodsMetricSource{
															MetricName: metrics.Pods.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Resource: &kubernetes.ResourceMetricSource{
															Name: metrics.Resource.Name.ValueString(),
															Target: &kubernetes.MetricTarget{
																AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																AverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																},
																Type: metrics.Resource.Target.Type.ValueString(),
																Value: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																},
															},
															TargetAverageUtilization: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Type: metrics.Type.ValueString(),
													}
												}
												return a
											}(),
											MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
										},
										PodAnnotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										PodSecurityContext: &kubernetes.PodSecurityContext{
											FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
											FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
											RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
											},
											SupplementalGroups: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Sysctls: func() []*kubernetes.Sysctl {
												a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.Sysctls))
												for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.Sysctls {
													a[i] = &kubernetes.Sysctl{
														Name:  sysctls.Name.ValueString(),
														Value: sysctls.Value.ValueString(),
													}
												}
												return a
											}(),
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
										Resources: &kubernetes.Resources{
											Limits: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Requests: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
										},
										Strategy: &kubernetes.DeploymentStrategy{
											RollingUpdate: &kubernetes.RollingUpdateDeployment{
												MaxSurge: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
												},
												MaxUnavailable: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
												},
											},
											Type: model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Strategy.Type.ValueString(),
										},
										Tolerations: func() []*v11.Toleration {
											a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Tolerations))
											for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Deployment.Tolerations {
												a[i] = &v11.Toleration{
													Effect:            ptrify(tolerations.Effect.ValueString()),
													Key:               ptrify(tolerations.Key.ValueString()),
													Operator:          ptrify(tolerations.Operator.ValueString()),
													TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
													Value:             ptrify(tolerations.Value.ValueString()),
												}
											}
											return a
										}(),
									},
									Overlays: func() []*v1alpha12.K8SObjectOverlay {
										a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Overlays))
										for i, overlays := range model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Overlays {
											a[i] = &v1alpha12.K8SObjectOverlay{
												ApiVersion: overlays.ApiVersion.ValueString(),
												Kind:       overlays.Kind.ValueString(),
												Name:       overlays.Name.ValueString(),
												Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
													a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
													for i, patches := range overlays.Patches {
														a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
															Path: patches.Path.ValueString(),
															Value: &structpb.Value{
																BoolValue: patches.Value.BoolValue.ValueBool(),
																ListValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																NumberValue: patches.Value.NumberValue.ValueFloat64(),
																StringValue: patches.Value.StringValue.ValueString(),
																StructValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
									Service: &kubernetes.Service{
										Annotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Labels: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Ports: func() []*kubernetes.ServicePort {
											a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Service.Ports))
											for i, ports := range model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Service.Ports {
												a[i] = &kubernetes.ServicePort{
													Name:     ports.Name.ValueString(),
													NodePort: int32(ports.NodePort.ValueInt64()),
													Port:     int32(ports.Port.ValueInt64()),
													Protocol: ports.Protocol.ValueString(),
													TargetPort: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
														Type:   ports.TargetPort.Type.ValueInt64(),
													},
												}
											}
											return a
										}(),
										Type: model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.Service.Type.ValueString(),
									},
									ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
										a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.ServiceAccount.ImagePullSecrets))
										for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.Satellite.KubeSpec.ServiceAccount.ImagePullSecrets {
											a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
										}
										return a
									}()},
								},
							},
							Xcp: &v1alpha13.XCP{
								CentralAuthMode:       v1alpha13.XCP_CentralAuthMode(v1alpha13.XCP_CentralAuthMode_value[model.InstallTemplate.Helm.Spec.Components.Xcp.CentralAuthMode.ValueString()]),
								CentralProvidedCaCert: model.InstallTemplate.Helm.Spec.Components.Xcp.CentralProvidedCaCert.ValueBool(),
								ConfigProtection: &common.ConfigProtection{
									AuthorizedUsers: func() []string {
										tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.ConfigProtection.AuthorizedUsers.Elements()))
										resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Xcp.ConfigProtection.AuthorizedUsers.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									EnableAuthorizedCreateUpdateDeleteOnXcpConfigs: model.InstallTemplate.Helm.Spec.Components.Xcp.ConfigProtection.EnableAuthorizedCreateUpdateDeleteOnXcpConfigs.ValueBool(),
									EnableAuthorizedUpdateDeleteOnXcpConfigs:       model.InstallTemplate.Helm.Spec.Components.Xcp.ConfigProtection.EnableAuthorizedUpdateDeleteOnXcpConfigs.ValueBool(),
								},
								EnableHttpMeshInternalIdentityPropagation: model.InstallTemplate.Helm.Spec.Components.Xcp.EnableHttpMeshInternalIdentityPropagation.ValueBool(),
								IsolationBoundaries: func() []*v1alpha13.IsolationBoundary {
									a := make([]*v1alpha13.IsolationBoundary, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.IsolationBoundaries))
									for i, isolation_boundaries := range model.InstallTemplate.Helm.Spec.Components.Xcp.IsolationBoundaries {
										a[i] = &v1alpha13.IsolationBoundary{
											Name: isolation_boundaries.Name.ValueString(),
											Revisions: func() []*v1alpha13.IstioRevision {
												a := make([]*v1alpha13.IstioRevision, 0, len(isolation_boundaries.Revisions))
												for i, revisions := range isolation_boundaries.Revisions {
													a[i] = &v1alpha13.IstioRevision{
														Disable: revisions.Disable.ValueBool(),
														Istio: &v1alpha13.Istio{
															BaseOverlays: func() []*v1alpha12.K8SObjectOverlay {
																a := make([]*v1alpha12.K8SObjectOverlay, 0, len(revisions.Istio.BaseOverlays))
																for i, base_overlays := range revisions.Istio.BaseOverlays {
																	a[i] = &v1alpha12.K8SObjectOverlay{
																		ApiVersion: base_overlays.ApiVersion.ValueString(),
																		Kind:       base_overlays.Kind.ValueString(),
																		Name:       base_overlays.Name.ValueString(),
																		Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
																			a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(base_overlays.Patches))
																			for i, patches := range base_overlays.Patches {
																				a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
																					Path: patches.Path.ValueString(),
																					Value: &structpb.Value{
																						BoolValue: patches.Value.BoolValue.ValueBool(),
																						ListValue: func() map[string]string {
																							tmp := make(map[string]string)
																							resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																							return tmp
																						}(),
																						NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																						NumberValue: patches.Value.NumberValue.ValueFloat64(),
																						StringValue: patches.Value.StringValue.ValueString(),
																						StructValue: func() map[string]string {
																							tmp := make(map[string]string)
																							resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
															CniOverlays: func() []*v1alpha12.K8SObjectOverlay {
																a := make([]*v1alpha12.K8SObjectOverlay, 0, len(revisions.Istio.CniOverlays))
																for i, cni_overlays := range revisions.Istio.CniOverlays {
																	a[i] = &v1alpha12.K8SObjectOverlay{
																		ApiVersion: cni_overlays.ApiVersion.ValueString(),
																		Kind:       cni_overlays.Kind.ValueString(),
																		Name:       cni_overlays.Name.ValueString(),
																		Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
																			a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(cni_overlays.Patches))
																			for i, patches := range cni_overlays.Patches {
																				a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
																					Path: patches.Path.ValueString(),
																					Value: &structpb.Value{
																						BoolValue: patches.Value.BoolValue.ValueBool(),
																						ListValue: func() map[string]string {
																							tmp := make(map[string]string)
																							resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																							return tmp
																						}(),
																						NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																						NumberValue: patches.Value.NumberValue.ValueFloat64(),
																						StringValue: patches.Value.StringValue.ValueString(),
																						StructValue: func() map[string]string {
																							tmp := make(map[string]string)
																							resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
															DefaultWorkloadCertTtl: &durationpb.Duration{
																Nanos:   int32(revisions.Istio.DefaultWorkloadCertTtl.Nanos.ValueInt64()),
																Seconds: revisions.Istio.DefaultWorkloadCertTtl.Seconds.ValueInt64(),
															},
															KubeSpec: &kubernetes.KubernetesIstioComponentSpec{
																Cni: &kubernetes.CNI{
																	BinaryDirectory:        revisions.Istio.KubeSpec.Cni.BinaryDirectory.ValueString(),
																	Chained:                revisions.Istio.KubeSpec.Cni.Chained.ValueBool(),
																	ClusterRole:            revisions.Istio.KubeSpec.Cni.ClusterRole.ValueString(),
																	ConfigurationDirectory: revisions.Istio.KubeSpec.Cni.ConfigurationDirectory.ValueString(),
																	ConfigurationFileName:  revisions.Istio.KubeSpec.Cni.ConfigurationFileName.ValueString(),
																	Revision:               revisions.Istio.KubeSpec.Cni.Revision.ValueString(),
																},
																Deployment: &kubernetes.Deployment{
																	Affinity: &kubernetes.Affinity{
																		NodeAffinity: &kubernetes.NodeAffinity{
																			PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
																				a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(revisions.Istio.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
																				for i, preferred_during_scheduling_ignored_during_execution := range revisions.Istio.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
																					a[i] = &kubernetes.PreferredSchedulingTerm{
																						Preference: &kubernetes.NodeSelectorTerm{
																							MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																								a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																								for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																									a[i] = &kubernetes.NodeSelectorRequirement{
																										Key:      match_expressions.Key.ValueString(),
																										Operator: match_expressions.Operator.ValueString(),
																										Values: func() []string {
																											tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																											resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																											return tmp
																										}(),
																									}
																								}
																								return a
																							}(),
																							MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																								a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																								for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																									a[i] = &kubernetes.NodeSelectorRequirement{
																										Key:      match_fields.Key.ValueString(),
																										Operator: match_fields.Operator.ValueString(),
																										Values: func() []string {
																											tmp := make([]string, 0, len(match_fields.Values.Elements()))
																											resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																											return tmp
																										}(),
																									}
																								}
																								return a
																							}(),
																						},
																						Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
																					}
																				}
																				return a
																			}(),
																			RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
																				a := make([]*kubernetes.NodeSelectorTerm, 0, len(revisions.Istio.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
																				for i, node_selector_terms := range revisions.Istio.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
																					a[i] = &kubernetes.NodeSelectorTerm{
																						MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																							a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																							for i, match_expressions := range node_selector_terms.MatchExpressions {
																								a[i] = &kubernetes.NodeSelectorRequirement{
																									Key:      match_expressions.Key.ValueString(),
																									Operator: match_expressions.Operator.ValueString(),
																									Values: func() []string {
																										tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																										resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																										return tmp
																									}(),
																								}
																							}
																							return a
																						}(),
																						MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																							a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																							for i, match_fields := range node_selector_terms.MatchFields {
																								a[i] = &kubernetes.NodeSelectorRequirement{
																									Key:      match_fields.Key.ValueString(),
																									Operator: match_fields.Operator.ValueString(),
																									Values: func() []string {
																										tmp := make([]string, 0, len(match_fields.Values.Elements()))
																										resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																										return tmp
																									}(),
																								}
																							}
																							return a
																						}(),
																					}
																				}
																				return a
																			}()},
																		},
																		PodAffinity: &kubernetes.PodAffinity{
																			PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
																				a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(revisions.Istio.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
																				for i, preferred_during_scheduling_ignored_during_execution := range revisions.Istio.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
																					a[i] = &kubernetes.WeightedPodAffinityTerm{
																						PodAffinityTerm: &kubernetes.PodAffinityTerm{
																							LabelSelector: &v1.LabelSelector{
																								MatchExpressions: func() []*v1.LabelSelectorRequirement {
																									a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																									for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																										a[i] = &v1.LabelSelectorRequirement{
																											Key:      ptrify(match_expressions.Key.ValueString()),
																											Operator: ptrify(match_expressions.Operator.ValueString()),
																											Values: func() []string {
																												tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																												resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																												return tmp
																											}(),
																										}
																									}
																									return a
																								}(),
																								MatchLabels: func() map[string]string {
																									tmp := make(map[string]string)
																									resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																									return tmp
																								}(),
																							},
																							Namespaces: func() []string {
																								tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																								resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																								return tmp
																							}(),
																							TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
																						},
																						Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
																					}
																				}
																				return a
																			}(),
																			RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
																				a := make([]*kubernetes.PodAffinityTerm, 0, len(revisions.Istio.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
																				for i, required_during_scheduling_ignored_during_execution := range revisions.Istio.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
																					a[i] = &kubernetes.PodAffinityTerm{
																						LabelSelector: &v1.LabelSelector{
																							MatchExpressions: func() []*v1.LabelSelectorRequirement {
																								a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																								for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																									a[i] = &v1.LabelSelectorRequirement{
																										Key:      ptrify(match_expressions.Key.ValueString()),
																										Operator: ptrify(match_expressions.Operator.ValueString()),
																										Values: func() []string {
																											tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																											resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																											return tmp
																										}(),
																									}
																								}
																								return a
																							}(),
																							MatchLabels: func() map[string]string {
																								tmp := make(map[string]string)
																								resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																								return tmp
																							}(),
																						},
																						Namespaces: func() []string {
																							tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																							resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																							return tmp
																						}(),
																						TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
																					}
																				}
																				return a
																			}(),
																		},
																		PodAntiAffinity: &kubernetes.PodAntiAffinity{
																			PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
																				a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(revisions.Istio.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
																				for i, preferred_during_scheduling_ignored_during_execution := range revisions.Istio.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
																					a[i] = &kubernetes.WeightedPodAffinityTerm{
																						PodAffinityTerm: &kubernetes.PodAffinityTerm{
																							LabelSelector: &v1.LabelSelector{
																								MatchExpressions: func() []*v1.LabelSelectorRequirement {
																									a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																									for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																										a[i] = &v1.LabelSelectorRequirement{
																											Key:      ptrify(match_expressions.Key.ValueString()),
																											Operator: ptrify(match_expressions.Operator.ValueString()),
																											Values: func() []string {
																												tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																												resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																												return tmp
																											}(),
																										}
																									}
																									return a
																								}(),
																								MatchLabels: func() map[string]string {
																									tmp := make(map[string]string)
																									resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																									return tmp
																								}(),
																							},
																							Namespaces: func() []string {
																								tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																								resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																								return tmp
																							}(),
																							TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
																						},
																						Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
																					}
																				}
																				return a
																			}(),
																			RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
																				a := make([]*kubernetes.PodAffinityTerm, 0, len(revisions.Istio.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
																				for i, required_during_scheduling_ignored_during_execution := range revisions.Istio.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
																					a[i] = &kubernetes.PodAffinityTerm{
																						LabelSelector: &v1.LabelSelector{
																							MatchExpressions: func() []*v1.LabelSelectorRequirement {
																								a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																								for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																									a[i] = &v1.LabelSelectorRequirement{
																										Key:      ptrify(match_expressions.Key.ValueString()),
																										Operator: ptrify(match_expressions.Operator.ValueString()),
																										Values: func() []string {
																											tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																											resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																											return tmp
																										}(),
																									}
																								}
																								return a
																							}(),
																							MatchLabels: func() map[string]string {
																								tmp := make(map[string]string)
																								resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																								return tmp
																							}(),
																						},
																						Namespaces: func() []string {
																							tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																							resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																							return tmp
																						}(),
																						TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
																					}
																				}
																				return a
																			}(),
																		},
																	},
																	ContainerSecurityContext: &kubernetes.SecurityContext{
																		AllowPrivilegeEscalation: ptrify(revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
																		Capabilities: &kubernetes.Capabilities{
																			Add: func() []string {
																				tmp := make([]string, 0, len(revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
																				resp.Diagnostics.Append(revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																			Drop: func() []string {
																				tmp := make([]string, 0, len(revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
																				resp.Diagnostics.Append(revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		},
																		Privileged:             ptrify(revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
																		ProcMount:              ptrify(revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
																		ReadOnlyRootFilesystem: ptrify(revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
																		RunAsGroup:             ptrify(uint32(revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
																		RunAsNonRoot:           ptrify(revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
																		RunAsUser:              ptrify(uint32(revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
																		SeLinuxOptions: &kubernetes.SELinuxOptions{
																			Level: revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
																			Role:  revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
																			Type:  revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
																			User:  revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
																		},
																		SeccompProfile: &kubernetes.SeccompProfile{
																			LocalhostProfile: revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
																			Type:             revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
																		},
																		WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
																			GmsaCredentialSpec:     revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
																			GmsaCredentialSpecName: revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
																			RunAsUserName:          revisions.Istio.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
																		},
																	},
																	Env: func() []*kubernetes.EnvVar {
																		a := make([]*kubernetes.EnvVar, 0, len(revisions.Istio.KubeSpec.Deployment.Env))
																		for i, env := range revisions.Istio.KubeSpec.Deployment.Env {
																			a[i] = &kubernetes.EnvVar{
																				Name:  env.Name.ValueString(),
																				Value: env.Value.ValueString(),
																				ValueFrom: &kubernetes.EnvVarSource{
																					ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
																						Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
																						LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
																						Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
																					},
																					FieldRef: &kubernetes.ObjectFieldSelector{
																						ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
																						FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
																					},
																					ResourceFieldRef: &kubernetes.ResourceFieldSelector{
																						ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
																						Divisor: &v1alpha12.IntOrString{
																							IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																							StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																							Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
																						},
																						Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
																					},
																					SecretKeyRef: &kubernetes.SecretKeySelector{
																						Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
																						LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
																						Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
																					},
																				},
																			}
																		}
																		return a
																	}(),
																	HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
																		MaxReplicas: int32(revisions.Istio.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
																		Metrics: func() []*kubernetes.MetricSpec {
																			a := make([]*kubernetes.MetricSpec, 0, len(revisions.Istio.KubeSpec.Deployment.HpaSpec.Metrics))
																			for i, metrics := range revisions.Istio.KubeSpec.Deployment.HpaSpec.Metrics {
																				a[i] = &kubernetes.MetricSpec{
																					External: &kubernetes.ExternalMetricSource{
																						MetricName: metrics.External.MetricName.ValueString(),
																						MetricSelector: &v1.LabelSelector{
																							MatchExpressions: func() []*v1.LabelSelectorRequirement {
																								a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																								for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																									a[i] = &v1.LabelSelectorRequirement{
																										Key:      ptrify(match_expressions.Key.ValueString()),
																										Operator: ptrify(match_expressions.Operator.ValueString()),
																										Values: func() []string {
																											tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																											resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																											return tmp
																										}(),
																									}
																								}
																								return a
																							}(),
																							MatchLabels: func() map[string]string {
																								tmp := make(map[string]string)
																								resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																								return tmp
																							}(),
																						},
																						TargetAverageValue: &v1alpha12.IntOrString{
																							IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																							StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																							Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
																						},
																						TargetValue: &v1alpha12.IntOrString{
																							IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																							StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																							Type:   metrics.External.TargetValue.Type.ValueInt64(),
																						},
																					},
																					Object: &kubernetes.ObjectMetricSource{
																						AverageValue: &v1alpha12.IntOrString{
																							IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																							StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																							Type:   metrics.Object.AverageValue.Type.ValueInt64(),
																						},
																						MetricName: metrics.Object.MetricName.ValueString(),
																						Selector: &v1.LabelSelector{
																							MatchExpressions: func() []*v1.LabelSelectorRequirement {
																								a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																								for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																									a[i] = &v1.LabelSelectorRequirement{
																										Key:      ptrify(match_expressions.Key.ValueString()),
																										Operator: ptrify(match_expressions.Operator.ValueString()),
																										Values: func() []string {
																											tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																											resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																											return tmp
																										}(),
																									}
																								}
																								return a
																							}(),
																							MatchLabels: func() map[string]string {
																								tmp := make(map[string]string)
																								resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																								return tmp
																							}(),
																						},
																						Target: &kubernetes.CrossVersionObjectReference{
																							ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																							Kind:       metrics.Object.Target.Kind.ValueString(),
																							Name:       metrics.Object.Target.Name.ValueString(),
																						},
																						TargetValue: &v1alpha12.IntOrString{
																							IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																							StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																							Type:   metrics.Object.TargetValue.Type.ValueInt64(),
																						},
																					},
																					Pods: &kubernetes.PodsMetricSource{
																						MetricName: metrics.Pods.MetricName.ValueString(),
																						Selector: &v1.LabelSelector{
																							MatchExpressions: func() []*v1.LabelSelectorRequirement {
																								a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																								for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																									a[i] = &v1.LabelSelectorRequirement{
																										Key:      ptrify(match_expressions.Key.ValueString()),
																										Operator: ptrify(match_expressions.Operator.ValueString()),
																										Values: func() []string {
																											tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																											resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																											return tmp
																										}(),
																									}
																								}
																								return a
																							}(),
																							MatchLabels: func() map[string]string {
																								tmp := make(map[string]string)
																								resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																								return tmp
																							}(),
																						},
																						TargetAverageValue: &v1alpha12.IntOrString{
																							IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																							StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																							Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
																						},
																					},
																					Resource: &kubernetes.ResourceMetricSource{
																						Name: metrics.Resource.Name.ValueString(),
																						Target: &kubernetes.MetricTarget{
																							AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																							AverageValue: &v1alpha12.IntOrString{
																								IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																								StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																								Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																							},
																							Type: metrics.Resource.Target.Type.ValueString(),
																							Value: &v1alpha12.IntOrString{
																								IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																								StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																								Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																							},
																						},
																						TargetAverageUtilization: &v1alpha12.IntOrString{
																							IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																							StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																							Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
																						},
																						TargetAverageValue: &v1alpha12.IntOrString{
																							IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																							StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																							Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
																						},
																					},
																					Type: metrics.Type.ValueString(),
																				}
																			}
																			return a
																		}(),
																		MinReplicas: int32(revisions.Istio.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
																	},
																	PodAnnotations: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(revisions.Istio.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	PodSecurityContext: &kubernetes.PodSecurityContext{
																		FsGroup:             ptrify(uint32(revisions.Istio.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
																		FsGroupChangePolicy: ptrify(revisions.Istio.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
																		RunAsGroup:          ptrify(uint32(revisions.Istio.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
																		RunAsNonRoot:        ptrify(revisions.Istio.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
																		RunAsUser:           ptrify(uint32(revisions.Istio.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
																		SeLinuxOptions: &kubernetes.SELinuxOptions{
																			Level: revisions.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
																			Role:  revisions.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
																			Type:  revisions.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
																			User:  revisions.Istio.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
																		},
																		SeccompProfile: &kubernetes.SeccompProfile{
																			LocalhostProfile: revisions.Istio.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
																			Type:             revisions.Istio.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
																		},
																		SupplementalGroups: func() []string {
																			tmp := make([]string, 0, len(revisions.Istio.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
																			resp.Diagnostics.Append(revisions.Istio.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																		Sysctls: func() []*kubernetes.Sysctl {
																			a := make([]*kubernetes.Sysctl, 0, len(revisions.Istio.KubeSpec.Deployment.PodSecurityContext.Sysctls))
																			for i, sysctls := range revisions.Istio.KubeSpec.Deployment.PodSecurityContext.Sysctls {
																				a[i] = &kubernetes.Sysctl{
																					Name:  sysctls.Name.ValueString(),
																					Value: sysctls.Value.ValueString(),
																				}
																			}
																			return a
																		}(),
																		WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
																			GmsaCredentialSpec:     revisions.Istio.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
																			GmsaCredentialSpecName: revisions.Istio.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
																			RunAsUserName:          revisions.Istio.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
																		},
																	},
																	ReplicaCount: uint32(revisions.Istio.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
																	Resources: &kubernetes.Resources{
																		Limits: func() map[string]string {
																			tmp := make(map[string]string)
																			resp.Diagnostics.Append(revisions.Istio.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																		Requests: func() map[string]string {
																			tmp := make(map[string]string)
																			resp.Diagnostics.Append(revisions.Istio.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	},
																	Strategy: &kubernetes.DeploymentStrategy{
																		RollingUpdate: &kubernetes.RollingUpdateDeployment{
																			MaxSurge: &v1alpha12.IntOrString{
																				IntVal: &wrapperspb.Int32Value{Value: int32(revisions.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
																				StrVal: &wrapperspb.StringValue{Value: revisions.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
																				Type:   revisions.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
																			},
																			MaxUnavailable: &v1alpha12.IntOrString{
																				IntVal: &wrapperspb.Int32Value{Value: int32(revisions.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
																				StrVal: &wrapperspb.StringValue{Value: revisions.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
																				Type:   revisions.Istio.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
																			},
																		},
																		Type: revisions.Istio.KubeSpec.Deployment.Strategy.Type.ValueString(),
																	},
																	Tolerations: func() []*v11.Toleration {
																		a := make([]*v11.Toleration, 0, len(revisions.Istio.KubeSpec.Deployment.Tolerations))
																		for i, tolerations := range revisions.Istio.KubeSpec.Deployment.Tolerations {
																			a[i] = &v11.Toleration{
																				Effect:            ptrify(tolerations.Effect.ValueString()),
																				Key:               ptrify(tolerations.Key.ValueString()),
																				Operator:          ptrify(tolerations.Operator.ValueString()),
																				TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
																				Value:             ptrify(tolerations.Value.ValueString()),
																			}
																		}
																		return a
																	}(),
																},
																Overlays: func() []*v1alpha12.K8SObjectOverlay {
																	a := make([]*v1alpha12.K8SObjectOverlay, 0, len(revisions.Istio.KubeSpec.Overlays))
																	for i, overlays := range revisions.Istio.KubeSpec.Overlays {
																		a[i] = &v1alpha12.K8SObjectOverlay{
																			ApiVersion: overlays.ApiVersion.ValueString(),
																			Kind:       overlays.Kind.ValueString(),
																			Name:       overlays.Name.ValueString(),
																			Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
																				a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
																				for i, patches := range overlays.Patches {
																					a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
																						Path: patches.Path.ValueString(),
																						Value: &structpb.Value{
																							BoolValue: patches.Value.BoolValue.ValueBool(),
																							ListValue: func() map[string]string {
																								tmp := make(map[string]string)
																								resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																								return tmp
																							}(),
																							NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																							NumberValue: patches.Value.NumberValue.ValueFloat64(),
																							StringValue: patches.Value.StringValue.ValueString(),
																							StructValue: func() map[string]string {
																								tmp := make(map[string]string)
																								resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
																Service: &kubernetes.Service{
																	Annotations: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(revisions.Istio.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	Labels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(revisions.Istio.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																	Ports: func() []*kubernetes.ServicePort {
																		a := make([]*kubernetes.ServicePort, 0, len(revisions.Istio.KubeSpec.Service.Ports))
																		for i, ports := range revisions.Istio.KubeSpec.Service.Ports {
																			a[i] = &kubernetes.ServicePort{
																				Name:     ports.Name.ValueString(),
																				NodePort: int32(ports.NodePort.ValueInt64()),
																				Port:     int32(ports.Port.ValueInt64()),
																				Protocol: ports.Protocol.ValueString(),
																				TargetPort: &v1alpha12.IntOrString{
																					IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
																					StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
																					Type:   ports.TargetPort.Type.ValueInt64(),
																				},
																			}
																		}
																		return a
																	}(),
																	Type: revisions.Istio.KubeSpec.Service.Type.ValueString(),
																},
																ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
																	a := make([]*kubernetes.LocalObjectReference, 0, len(revisions.Istio.KubeSpec.ServiceAccount.ImagePullSecrets))
																	for i, image_pull_secrets := range revisions.Istio.KubeSpec.ServiceAccount.ImagePullSecrets {
																		a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
																	}
																	return a
																}()},
															},
															MaxWorkloadCertTtl: &durationpb.Duration{
																Nanos:   int32(revisions.Istio.MaxWorkloadCertTtl.Nanos.ValueInt64()),
																Seconds: revisions.Istio.MaxWorkloadCertTtl.Seconds.ValueInt64(),
															},
															PilotOverlays: func() []*v1alpha12.K8SObjectOverlay {
																a := make([]*v1alpha12.K8SObjectOverlay, 0, len(revisions.Istio.PilotOverlays))
																for i, pilot_overlays := range revisions.Istio.PilotOverlays {
																	a[i] = &v1alpha12.K8SObjectOverlay{
																		ApiVersion: pilot_overlays.ApiVersion.ValueString(),
																		Kind:       pilot_overlays.Kind.ValueString(),
																		Name:       pilot_overlays.Name.ValueString(),
																		Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
																			a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(pilot_overlays.Patches))
																			for i, patches := range pilot_overlays.Patches {
																				a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
																					Path: patches.Path.ValueString(),
																					Value: &structpb.Value{
																						BoolValue: patches.Value.BoolValue.ValueBool(),
																						ListValue: func() map[string]string {
																							tmp := make(map[string]string)
																							resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																							return tmp
																						}(),
																						NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																						NumberValue: patches.Value.NumberValue.ValueFloat64(),
																						StringValue: patches.Value.StringValue.ValueString(),
																						StructValue: func() map[string]string {
																							tmp := make(map[string]string)
																							resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
															TraceSamplingRate: revisions.Istio.TraceSamplingRate.ValueFloat64(),
															TrustDomain:       revisions.Istio.TrustDomain.ValueString(),
															TsbVersion:        revisions.Istio.TsbVersion.ValueString(),
														},
														Name: revisions.Name.ValueString(),
													}
												}
												return a
											}(),
										}
									}
									return a
								}(),
								KubeSpec: &kubernetes.KubernetesComponentSpec{
									Deployment: &kubernetes.Deployment{
										Affinity: &kubernetes.Affinity{
											NodeAffinity: &kubernetes.NodeAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PreferredSchedulingTerm {
													a := make([]*kubernetes.PreferredSchedulingTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Affinity.NodeAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PreferredSchedulingTerm{
															Preference: &kubernetes.NodeSelectorTerm{
																MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions))
																	for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.Preference.MatchExpressions {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_expressions.Key.ValueString(),
																			Operator: match_expressions.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																	a := make([]*kubernetes.NodeSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.Preference.MatchFields))
																	for i, match_fields := range preferred_during_scheduling_ignored_during_execution.Preference.MatchFields {
																		a[i] = &kubernetes.NodeSelectorRequirement{
																			Key:      match_fields.Key.ValueString(),
																			Operator: match_fields.Operator.ValueString(),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_fields.Values.Elements()))
																				resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: &kubernetes.NodeSelector{NodeSelectorTerms: func() []*kubernetes.NodeSelectorTerm {
													a := make([]*kubernetes.NodeSelectorTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms))
													for i, node_selector_terms := range model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms {
														a[i] = &kubernetes.NodeSelectorTerm{
															MatchExpressions: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchExpressions))
																for i, match_expressions := range node_selector_terms.MatchExpressions {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_expressions.Key.ValueString(),
																		Operator: match_expressions.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																			resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
															MatchFields: func() []*kubernetes.NodeSelectorRequirement {
																a := make([]*kubernetes.NodeSelectorRequirement, 0, len(node_selector_terms.MatchFields))
																for i, match_fields := range node_selector_terms.MatchFields {
																	a[i] = &kubernetes.NodeSelectorRequirement{
																		Key:      match_fields.Key.ValueString(),
																		Operator: match_fields.Operator.ValueString(),
																		Values: func() []string {
																			tmp := make([]string, 0, len(match_fields.Values.Elements()))
																			resp.Diagnostics.Append(match_fields.Values.ElementsAs(ctx, &tmp, false)...)
																			return tmp
																		}(),
																	}
																}
																return a
															}(),
														}
													}
													return a
												}()},
											},
											PodAffinity: &kubernetes.PodAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Affinity.PodAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
											PodAntiAffinity: &kubernetes.PodAntiAffinity{
												PreferredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.WeightedPodAffinityTerm {
													a := make([]*kubernetes.WeightedPodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution))
													for i, preferred_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Affinity.PodAntiAffinity.PreferredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.WeightedPodAffinityTerm{
															PodAffinityTerm: &kubernetes.PodAffinityTerm{
																LabelSelector: &v1.LabelSelector{
																	MatchExpressions: func() []*v1.LabelSelectorRequirement {
																		a := make([]*v1.LabelSelectorRequirement, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions))
																		for i, match_expressions := range preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchExpressions {
																			a[i] = &v1.LabelSelectorRequirement{
																				Key:      ptrify(match_expressions.Key.ValueString()),
																				Operator: ptrify(match_expressions.Operator.ValueString()),
																				Values: func() []string {
																					tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																					resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																					return tmp
																				}(),
																			}
																		}
																		return a
																	}(),
																	MatchLabels: func() map[string]string {
																		tmp := make(map[string]string)
																		resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																		return tmp
																	}(),
																},
																Namespaces: func() []string {
																	tmp := make([]string, 0, len(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.Elements()))
																	resp.Diagnostics.Append(preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.Namespaces.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																TopologyKey: preferred_during_scheduling_ignored_during_execution.PodAffinityTerm.TopologyKey.ValueString(),
															},
															Weight: int32(preferred_during_scheduling_ignored_during_execution.Weight.ValueInt64()),
														}
													}
													return a
												}(),
												RequiredDuringSchedulingIgnoredDuringExecution: func() []*kubernetes.PodAffinityTerm {
													a := make([]*kubernetes.PodAffinityTerm, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution))
													for i, required_during_scheduling_ignored_during_execution := range model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution {
														a[i] = &kubernetes.PodAffinityTerm{
															LabelSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions))
																	for i, match_expressions := range required_during_scheduling_ignored_during_execution.LabelSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.LabelSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Namespaces: func() []string {
																tmp := make([]string, 0, len(required_during_scheduling_ignored_during_execution.Namespaces.Elements()))
																resp.Diagnostics.Append(required_during_scheduling_ignored_during_execution.Namespaces.ElementsAs(ctx, &tmp, false)...)
																return tmp
															}(),
															TopologyKey: required_during_scheduling_ignored_during_execution.TopologyKey.ValueString(),
														}
													}
													return a
												}(),
											},
										},
										ContainerSecurityContext: &kubernetes.SecurityContext{
											AllowPrivilegeEscalation: ptrify(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.AllowPrivilegeEscalation.ValueBool()),
											Capabilities: &kubernetes.Capabilities{
												Add: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Add.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
												Drop: func() []string {
													tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.Elements()))
													resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.Capabilities.Drop.ElementsAs(ctx, &tmp, false)...)
													return tmp
												}(),
											},
											Privileged:             ptrify(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.Privileged.ValueBool()),
											ProcMount:              ptrify(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.ProcMount.ValueString()),
											ReadOnlyRootFilesystem: ptrify(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.ReadOnlyRootFilesystem.ValueBool()),
											RunAsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:           ptrify(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:              ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.SeccompProfile.Type.ValueString(),
											},
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ContainerSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										Env: func() []*kubernetes.EnvVar {
											a := make([]*kubernetes.EnvVar, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Env))
											for i, env := range model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Env {
												a[i] = &kubernetes.EnvVar{
													Name:  env.Name.ValueString(),
													Value: env.Value.ValueString(),
													ValueFrom: &kubernetes.EnvVarSource{
														ConfigMapKeyRef: &kubernetes.ConfigMapKeySelector{
															Key:                  env.ValueFrom.ConfigMapKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.ConfigMapKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.ConfigMapKeyRef.Optional.ValueBool(),
														},
														FieldRef: &kubernetes.ObjectFieldSelector{
															ApiVersion: env.ValueFrom.FieldRef.ApiVersion.ValueString(),
															FieldPath:  env.ValueFrom.FieldRef.FieldPath.ValueString(),
														},
														ResourceFieldRef: &kubernetes.ResourceFieldSelector{
															ContainerName: env.ValueFrom.ResourceFieldRef.ContainerName.ValueString(),
															Divisor: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(env.ValueFrom.ResourceFieldRef.Divisor.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: env.ValueFrom.ResourceFieldRef.Divisor.StrVal.Value.ValueString()},
																Type:   env.ValueFrom.ResourceFieldRef.Divisor.Type.ValueInt64(),
															},
															Resource: env.ValueFrom.ResourceFieldRef.Resource.ValueString(),
														},
														SecretKeyRef: &kubernetes.SecretKeySelector{
															Key:                  env.ValueFrom.SecretKeyRef.Key.ValueString(),
															LocalObjectReference: &kubernetes.LocalObjectReference{Name: env.ValueFrom.SecretKeyRef.LocalObjectReference.Name.ValueString()},
															Optional:             env.ValueFrom.SecretKeyRef.Optional.ValueBool(),
														},
													},
												}
											}
											return a
										}(),
										HpaSpec: &kubernetes.HorizontalPodAutoscalerSpec{
											MaxReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.HpaSpec.MaxReplicas.ValueInt64()),
											Metrics: func() []*kubernetes.MetricSpec {
												a := make([]*kubernetes.MetricSpec, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.HpaSpec.Metrics))
												for i, metrics := range model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.HpaSpec.Metrics {
													a[i] = &kubernetes.MetricSpec{
														External: &kubernetes.ExternalMetricSource{
															MetricName: metrics.External.MetricName.ValueString(),
															MetricSelector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.External.MetricSelector.MatchExpressions))
																	for i, match_expressions := range metrics.External.MetricSelector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.External.MetricSelector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetAverageValue.Type.ValueInt64(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.External.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.External.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.External.TargetValue.Type.ValueInt64(),
															},
														},
														Object: &kubernetes.ObjectMetricSource{
															AverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.AverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.AverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.AverageValue.Type.ValueInt64(),
															},
															MetricName: metrics.Object.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Object.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Object.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Object.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															Target: &kubernetes.CrossVersionObjectReference{
																ApiVersion: metrics.Object.Target.ApiVersion.ValueString(),
																Kind:       metrics.Object.Target.Kind.ValueString(),
																Name:       metrics.Object.Target.Name.ValueString(),
															},
															TargetValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Object.TargetValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Object.TargetValue.StrVal.Value.ValueString()},
																Type:   metrics.Object.TargetValue.Type.ValueInt64(),
															},
														},
														Pods: &kubernetes.PodsMetricSource{
															MetricName: metrics.Pods.MetricName.ValueString(),
															Selector: &v1.LabelSelector{
																MatchExpressions: func() []*v1.LabelSelectorRequirement {
																	a := make([]*v1.LabelSelectorRequirement, 0, len(metrics.Pods.Selector.MatchExpressions))
																	for i, match_expressions := range metrics.Pods.Selector.MatchExpressions {
																		a[i] = &v1.LabelSelectorRequirement{
																			Key:      ptrify(match_expressions.Key.ValueString()),
																			Operator: ptrify(match_expressions.Operator.ValueString()),
																			Values: func() []string {
																				tmp := make([]string, 0, len(match_expressions.Values.Elements()))
																				resp.Diagnostics.Append(match_expressions.Values.ElementsAs(ctx, &tmp, false)...)
																				return tmp
																			}(),
																		}
																	}
																	return a
																}(),
																MatchLabels: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(metrics.Pods.Selector.MatchLabels.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Pods.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Pods.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Pods.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Resource: &kubernetes.ResourceMetricSource{
															Name: metrics.Resource.Name.ValueString(),
															Target: &kubernetes.MetricTarget{
																AverageUtilization: int32(metrics.Resource.Target.AverageUtilization.ValueInt64()),
																AverageValue: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.AverageValue.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.AverageValue.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.AverageValue.Type.ValueInt64(),
																},
																Type: metrics.Resource.Target.Type.ValueString(),
																Value: &v1alpha12.IntOrString{
																	IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.Target.Value.IntVal.Value.ValueInt64())},
																	StrVal: &wrapperspb.StringValue{Value: metrics.Resource.Target.Value.StrVal.Value.ValueString()},
																	Type:   metrics.Resource.Target.Value.Type.ValueInt64(),
																},
															},
															TargetAverageUtilization: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageUtilization.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageUtilization.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageUtilization.Type.ValueInt64(),
															},
															TargetAverageValue: &v1alpha12.IntOrString{
																IntVal: &wrapperspb.Int32Value{Value: int32(metrics.Resource.TargetAverageValue.IntVal.Value.ValueInt64())},
																StrVal: &wrapperspb.StringValue{Value: metrics.Resource.TargetAverageValue.StrVal.Value.ValueString()},
																Type:   metrics.Resource.TargetAverageValue.Type.ValueInt64(),
															},
														},
														Type: metrics.Type.ValueString(),
													}
												}
												return a
											}(),
											MinReplicas: int32(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.HpaSpec.MinReplicas.ValueInt64()),
										},
										PodAnnotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodAnnotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										PodSecurityContext: &kubernetes.PodSecurityContext{
											FsGroup:             ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.FsGroup.ValueInt64())),
											FsGroupChangePolicy: ptrify(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.FsGroupChangePolicy.ValueString()),
											RunAsGroup:          ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.RunAsGroup.ValueInt64())),
											RunAsNonRoot:        ptrify(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.RunAsNonRoot.ValueBool()),
											RunAsUser:           ptrify(uint32(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.RunAsUser.ValueInt64())),
											SeLinuxOptions: &kubernetes.SELinuxOptions{
												Level: model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Level.ValueString(),
												Role:  model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Role.ValueString(),
												Type:  model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.Type.ValueString(),
												User:  model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeLinuxOptions.User.ValueString(),
											},
											SeccompProfile: &kubernetes.SeccompProfile{
												LocalhostProfile: model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.LocalhostProfile.ValueString(),
												Type:             model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SeccompProfile.Type.ValueString(),
											},
											SupplementalGroups: func() []string {
												tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.Elements()))
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.SupplementalGroups.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Sysctls: func() []*kubernetes.Sysctl {
												a := make([]*kubernetes.Sysctl, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.Sysctls))
												for i, sysctls := range model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.Sysctls {
													a[i] = &kubernetes.Sysctl{
														Name:  sysctls.Name.ValueString(),
														Value: sysctls.Value.ValueString(),
													}
												}
												return a
											}(),
											WindowsOptions: &kubernetes.WindowsSecurityContextOptions{
												GmsaCredentialSpec:     model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpec.ValueString(),
												GmsaCredentialSpecName: model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.GmsaCredentialSpecName.ValueString(),
												RunAsUserName:          model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.PodSecurityContext.WindowsOptions.RunAsUserName.ValueString(),
											},
										},
										ReplicaCount: uint32(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.ReplicaCount.ValueInt64()),
										Resources: &kubernetes.Resources{
											Limits: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Resources.Limits.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
											Requests: func() map[string]string {
												tmp := make(map[string]string)
												resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Resources.Requests.ElementsAs(ctx, &tmp, false)...)
												return tmp
											}(),
										},
										Strategy: &kubernetes.DeploymentStrategy{
											RollingUpdate: &kubernetes.RollingUpdateDeployment{
												MaxSurge: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxSurge.Type.ValueInt64(),
												},
												MaxUnavailable: &v1alpha12.IntOrString{
													IntVal: &wrapperspb.Int32Value{Value: int32(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.IntVal.Value.ValueInt64())},
													StrVal: &wrapperspb.StringValue{Value: model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.StrVal.Value.ValueString()},
													Type:   model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.RollingUpdate.MaxUnavailable.Type.ValueInt64(),
												},
											},
											Type: model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Strategy.Type.ValueString(),
										},
										Tolerations: func() []*v11.Toleration {
											a := make([]*v11.Toleration, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Tolerations))
											for i, tolerations := range model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Deployment.Tolerations {
												a[i] = &v11.Toleration{
													Effect:            ptrify(tolerations.Effect.ValueString()),
													Key:               ptrify(tolerations.Key.ValueString()),
													Operator:          ptrify(tolerations.Operator.ValueString()),
													TolerationSeconds: ptrify(tolerations.TolerationSeconds.ValueInt64()),
													Value:             ptrify(tolerations.Value.ValueString()),
												}
											}
											return a
										}(),
									},
									Overlays: func() []*v1alpha12.K8SObjectOverlay {
										a := make([]*v1alpha12.K8SObjectOverlay, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Overlays))
										for i, overlays := range model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Overlays {
											a[i] = &v1alpha12.K8SObjectOverlay{
												ApiVersion: overlays.ApiVersion.ValueString(),
												Kind:       overlays.Kind.ValueString(),
												Name:       overlays.Name.ValueString(),
												Patches: func() []*v1alpha12.K8SObjectOverlay_PathValue {
													a := make([]*v1alpha12.K8SObjectOverlay_PathValue, 0, len(overlays.Patches))
													for i, patches := range overlays.Patches {
														a[i] = &v1alpha12.K8SObjectOverlay_PathValue{
															Path: patches.Path.ValueString(),
															Value: &structpb.Value{
																BoolValue: patches.Value.BoolValue.ValueBool(),
																ListValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.ListValue.ElementsAs(ctx, &tmp, false)...)
																	return tmp
																}(),
																NullValue:   structpb.NullValue(structpb.NullValue_value[patches.Value.NullValue.ValueString()]),
																NumberValue: patches.Value.NumberValue.ValueFloat64(),
																StringValue: patches.Value.StringValue.ValueString(),
																StructValue: func() map[string]string {
																	tmp := make(map[string]string)
																	resp.Diagnostics.Append(patches.Value.StructValue.ElementsAs(ctx, &tmp, false)...)
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
									Service: &kubernetes.Service{
										Annotations: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Service.Annotations.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Labels: func() map[string]string {
											tmp := make(map[string]string)
											resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Service.Labels.ElementsAs(ctx, &tmp, false)...)
											return tmp
										}(),
										Ports: func() []*kubernetes.ServicePort {
											a := make([]*kubernetes.ServicePort, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Service.Ports))
											for i, ports := range model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Service.Ports {
												a[i] = &kubernetes.ServicePort{
													Name:     ports.Name.ValueString(),
													NodePort: int32(ports.NodePort.ValueInt64()),
													Port:     int32(ports.Port.ValueInt64()),
													Protocol: ports.Protocol.ValueString(),
													TargetPort: &v1alpha12.IntOrString{
														IntVal: &wrapperspb.Int32Value{Value: int32(ports.TargetPort.IntVal.Value.ValueInt64())},
														StrVal: &wrapperspb.StringValue{Value: ports.TargetPort.StrVal.Value.ValueString()},
														Type:   ports.TargetPort.Type.ValueInt64(),
													},
												}
											}
											return a
										}(),
										Type: model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.Service.Type.ValueString(),
									},
									ServiceAccount: &kubernetes.ServiceAccount{ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
										a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.ServiceAccount.ImagePullSecrets))
										for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.Components.Xcp.KubeSpec.ServiceAccount.ImagePullSecrets {
											a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
										}
										return a
									}()},
								},
								Revision: model.InstallTemplate.Helm.Spec.Components.Xcp.Revision.ValueString(),
							},
						},
						Hub: model.InstallTemplate.Helm.Spec.Hub.ValueString(),
						ImagePullSecrets: func() []*kubernetes.LocalObjectReference {
							a := make([]*kubernetes.LocalObjectReference, 0, len(model.InstallTemplate.Helm.Spec.ImagePullSecrets))
							for i, image_pull_secrets := range model.InstallTemplate.Helm.Spec.ImagePullSecrets {
								a[i] = &kubernetes.LocalObjectReference{Name: image_pull_secrets.Name.ValueString()}
							}
							return a
						}(),
						ManagementPlane: &v1alpha13.ManagementPlaneSettings{
							ClusterName: model.InstallTemplate.Helm.Spec.ManagementPlane.ClusterName.ValueString(),
							Host:        model.InstallTemplate.Helm.Spec.ManagementPlane.Host.ValueString(),
							Port:        int32(model.InstallTemplate.Helm.Spec.ManagementPlane.Port.ValueInt64()),
							SelfSigned:  model.InstallTemplate.Helm.Spec.ManagementPlane.SelfSigned.ValueBool(),
						},
						MeshExpansion: &v1alpha13.MeshExpansionSettings{
							CustomGateway: &v1alpha13.MeshExpansionSettings_Gateway{
								Host: model.InstallTemplate.Helm.Spec.MeshExpansion.CustomGateway.Host.ValueString(),
								Port: int32(model.InstallTemplate.Helm.Spec.MeshExpansion.CustomGateway.Port.ValueInt64()),
							},
							Onboarding: &v1alpha13.MeshExpansionSettings_OnboardingPlane{
								Endpoint: &v1alpha13.MeshExpansionSettings_OnboardingPlane_Endpoint{
									Hosts: func() []string {
										tmp := make([]string, 0, len(model.InstallTemplate.Helm.Spec.MeshExpansion.Onboarding.Endpoint.Hosts.Elements()))
										resp.Diagnostics.Append(model.InstallTemplate.Helm.Spec.MeshExpansion.Onboarding.Endpoint.Hosts.ElementsAs(ctx, &tmp, false)...)
										return tmp
									}(),
									SecretName: model.InstallTemplate.Helm.Spec.MeshExpansion.Onboarding.Endpoint.SecretName.ValueString(),
								},
								LocalRepository: &v1alpha13.MeshExpansionSettings_OnboardingPlane_LocalRepository{},
								TokenIssuer: &v1alpha13.MeshExpansionSettings_OnboardingPlane_TokenIssuer{Jwt: &v1alpha13.MeshExpansionSettings_OnboardingPlane_TokenIssuer_JwtTokenIssuer{Expiration: &durationpb.Duration{
									Nanos:   int32(model.InstallTemplate.Helm.Spec.MeshExpansion.Onboarding.TokenIssuer.Jwt.Expiration.Nanos.ValueInt64()),
									Seconds: model.InstallTemplate.Helm.Spec.MeshExpansion.Onboarding.TokenIssuer.Jwt.Expiration.Seconds.ValueInt64(),
								}}},
								Uid: model.InstallTemplate.Helm.Spec.MeshExpansion.Onboarding.Uid.ValueString(),
								Workloads: &v1alpha14.WorkloadConfiguration{
									Authentication: &v1alpha14.WorkloadAuthenticationConfiguration{Jwt: &v1alpha14.JwtAuthenticationConfiguration{Issuers: func() []*v1alpha14.JwtIssuer {
										a := make([]*v1alpha14.JwtIssuer, 0, len(model.InstallTemplate.Helm.Spec.MeshExpansion.Onboarding.Workloads.Authentication.Jwt.Issuers))
										for i, issuers := range model.InstallTemplate.Helm.Spec.MeshExpansion.Onboarding.Workloads.Authentication.Jwt.Issuers {
											a[i] = &v1alpha14.JwtIssuer{
												Issuer:      issuers.Issuer.ValueString(),
												Jwks:        issuers.Jwks.ValueString(),
												JwksUri:     issuers.JwksUri.ValueString(),
												ShortName:   issuers.ShortName.ValueString(),
												TokenFields: &v1alpha14.JwtTokenFields{Attributes: &v1alpha14.JwtTokenField{JsonPath: issuers.TokenFields.Attributes.JsonPath.ValueString()}},
											}
										}
										return a
									}()}},
									Deregistration: &v1alpha14.WorkloadDeregistrationConfiguration{PropagationDelay: &durationpb.Duration{
										Nanos:   int32(model.InstallTemplate.Helm.Spec.MeshExpansion.Onboarding.Workloads.Deregistration.PropagationDelay.Nanos.ValueInt64()),
										Seconds: model.InstallTemplate.Helm.Spec.MeshExpansion.Onboarding.Workloads.Deregistration.PropagationDelay.Seconds.ValueInt64(),
									}},
								},
							},
						},
						MeshObservability: &v1alpha13.ControlPlaneSpec_MeshObservability{
							DemoSettings: &common.MeshObservabilitySettings{ApiEndpointMetricsEnabled: model.InstallTemplate.Helm.Spec.MeshObservability.DemoSettings.ApiEndpointMetricsEnabled.ValueBool()},
							Settings:     &common.MeshObservabilitySettings{ApiEndpointMetricsEnabled: model.InstallTemplate.Helm.Spec.MeshObservability.Settings.ApiEndpointMetricsEnabled.ValueBool()},
						},
						TelemetryStore: &v1alpha13.ControlPlaneSpec_TelemetryStore{Elastic: &v1alpha13.ElasticSearchSettings{
							Host:       model.InstallTemplate.Helm.Spec.TelemetryStore.Elastic.Host.ValueString(),
							Port:       int32(model.InstallTemplate.Helm.Spec.TelemetryStore.Elastic.Port.ValueInt64()),
							Protocol:   v1alpha13.ElasticSearchSettings_Protocol(v1alpha13.ElasticSearchSettings_Protocol_value[model.InstallTemplate.Helm.Spec.TelemetryStore.Elastic.Protocol.ValueString()]),
							SelfSigned: model.InstallTemplate.Helm.Spec.TelemetryStore.Elastic.SelfSigned.ValueBool(),
							Version:    int32(model.InstallTemplate.Helm.Spec.TelemetryStore.Elastic.Version.ValueInt64()),
						}},
						Tier1Cluster: model.InstallTemplate.Helm.Spec.Tier1Cluster.ValueBool(),
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
