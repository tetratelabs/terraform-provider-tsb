package cluster

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type NamespaceScopeModel struct {
	Scope      types.String `tfsdk:"scope"`
	Exceptions types.List   `tfsdk:"exceptions"`
}
type LocalityModel struct {
	Region types.String `tfsdk:"region"`
}
type LastSyncTimeStateModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type StateModel struct {
	LastSyncTime  LastSyncTimeStateModel `tfsdk:"last_sync_time"`
	Provider      types.String           `tfsdk:"provider"`
	TsbCpVersion  types.String           `tfsdk:"tsb_cp_version"`
	XcpVersion    types.String           `tfsdk:"xcp_version"`
	IstioVersions types.List             `tfsdk:"istio_versions"`
}
type Tier1ClusterModel struct {
	Value types.Bool `tfsdk:"value"`
}
type Token_TTLModel struct {
	Seconds types.Int64 `tfsdk:"seconds"`
	Nanos   types.Int64 `tfsdk:"nanos"`
}
type ManagementPlaneInstallTemplateHelmSpecModel struct {
	SelfSigned  types.Bool   `tfsdk:"self_signed"`
	ClusterName types.String `tfsdk:"cluster_name"`
	Host        types.String `tfsdk:"host"`
	Port        types.Int64  `tfsdk:"port"`
}
type CustomGatewayInstallTemplateHelmSpecMeshExpansionModel struct {
	Port types.Int64  `tfsdk:"port"`
	Host types.String `tfsdk:"host"`
}
type ExpirationInstallTemplateHelmSpecMeshExpansionOnboardingTokenIssuerJwtModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type JwtInstallTemplateHelmSpecMeshExpansionOnboardingTokenIssuerModel struct {
	Expiration ExpirationInstallTemplateHelmSpecMeshExpansionOnboardingTokenIssuerJwtModel `tfsdk:"expiration"`
}
type TokenIssuerInstallTemplateHelmSpecMeshExpansionOnboardingModel struct {
	Jwt JwtInstallTemplateHelmSpecMeshExpansionOnboardingTokenIssuerModel `tfsdk:"jwt"`
}
type AttributesInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationJwtIssuersTokenFieldsModel struct {
	JsonPath types.String `tfsdk:"json_path"`
}
type TokenFieldsInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationJwtIssuersModel struct {
	Attributes AttributesInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationJwtIssuersTokenFieldsModel `tfsdk:"attributes"`
}
type IssuersInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationJwtModel struct {
	ShortName   types.String                                                                                    `tfsdk:"short_name"`
	TokenFields TokenFieldsInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationJwtIssuersModel `tfsdk:"token_fields"`
	Issuer      types.String                                                                                    `tfsdk:"issuer"`
	Jwks        types.String                                                                                    `tfsdk:"jwks"`
	JwksUri     types.String                                                                                    `tfsdk:"jwks_uri"`
}
type JwtInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationModel struct {
	Issuers []*IssuersInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationJwtModel `tfsdk:"issuers"`
}
type AuthenticationInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsModel struct {
	Jwt JwtInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationModel `tfsdk:"jwt"`
}
type PropagationDelayInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsDeregistrationModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type DeregistrationInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsModel struct {
	PropagationDelay PropagationDelayInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsDeregistrationModel `tfsdk:"propagation_delay"`
}
type WorkloadsInstallTemplateHelmSpecMeshExpansionOnboardingModel struct {
	Authentication AuthenticationInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsModel `tfsdk:"authentication"`
	Deregistration DeregistrationInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsModel `tfsdk:"deregistration"`
}
type EndpointInstallTemplateHelmSpecMeshExpansionOnboardingModel struct {
	Hosts      types.List   `tfsdk:"hosts"`
	SecretName types.String `tfsdk:"secret_name"`
}
type LocalRepositoryInstallTemplateHelmSpecMeshExpansionOnboardingModel struct{}
type OnboardingInstallTemplateHelmSpecMeshExpansionModel struct {
	LocalRepository LocalRepositoryInstallTemplateHelmSpecMeshExpansionOnboardingModel `tfsdk:"local_repository"`
	TokenIssuer     TokenIssuerInstallTemplateHelmSpecMeshExpansionOnboardingModel     `tfsdk:"token_issuer"`
	Uid             types.String                                                       `tfsdk:"uid"`
	Workloads       WorkloadsInstallTemplateHelmSpecMeshExpansionOnboardingModel       `tfsdk:"workloads"`
	Endpoint        EndpointInstallTemplateHelmSpecMeshExpansionOnboardingModel        `tfsdk:"endpoint"`
}
type MeshExpansionInstallTemplateHelmSpecModel struct {
	CustomGateway CustomGatewayInstallTemplateHelmSpecMeshExpansionModel `tfsdk:"custom_gateway"`
	Onboarding    OnboardingInstallTemplateHelmSpecMeshExpansionModel    `tfsdk:"onboarding"`
}
type DemoSettingsInstallTemplateHelmSpecMeshObservabilityModel struct {
	ApiEndpointMetricsEnabled types.Bool `tfsdk:"api_endpoint_metrics_enabled"`
}
type SettingsInstallTemplateHelmSpecMeshObservabilityModel struct {
	ApiEndpointMetricsEnabled types.Bool `tfsdk:"api_endpoint_metrics_enabled"`
}
type MeshObservabilityInstallTemplateHelmSpecModel struct {
	DemoSettings DemoSettingsInstallTemplateHelmSpecMeshObservabilityModel `tfsdk:"demo_settings"`
	Settings     SettingsInstallTemplateHelmSpecMeshObservabilityModel     `tfsdk:"settings"`
}
type ElasticInstallTemplateHelmSpecTelemetryStoreModel struct {
	Version    types.Int64  `tfsdk:"version"`
	Host       types.String `tfsdk:"host"`
	Port       types.Int64  `tfsdk:"port"`
	Protocol   types.String `tfsdk:"protocol"`
	SelfSigned types.Bool   `tfsdk:"self_signed"`
}
type TelemetryStoreInstallTemplateHelmSpecModel struct {
	Elastic ElasticInstallTemplateHelmSpecTelemetryStoreModel `tfsdk:"elastic"`
}
type PatchesInstallTemplateHelmSpecComponentsIstioPilotOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type PilotOverlaysInstallTemplateHelmSpecComponentsIstioModel struct {
	Patches    []*PatchesInstallTemplateHelmSpecComponentsIstioPilotOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                       `tfsdk:"api_version"`
	Kind       types.String                                                       `tfsdk:"kind"`
	Name       types.String                                                       `tfsdk:"name"`
}
type DefaultWorkloadCert_TTLInstallTemplateHelmSpecComponentsIstioModel struct {
	Seconds types.Int64 `tfsdk:"seconds"`
	Nanos   types.Int64 `tfsdk:"nanos"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsIstioKubeSpecServicePortsModel struct {
	Type   types.Int64                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
}
type PortsInstallTemplateHelmSpecComponentsIstioKubeSpecServiceModel struct {
	NodePort   types.Int64                                                               `tfsdk:"node_port"`
	Port       types.Int64                                                               `tfsdk:"port"`
	Protocol   types.String                                                              `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsIstioKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                              `tfsdk:"name"`
}
type ServiceInstallTemplateHelmSpecComponentsIstioKubeSpecModel struct {
	Type        types.String                                                       `tfsdk:"type"`
	Annotations types.Map                                                          `tfsdk:"annotations"`
	Labels      types.Map                                                          `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsIstioKubeSpecServiceModel `tfsdk:"ports"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsIstioKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsIstioKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsIstioKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type CNIInstallTemplateHelmSpecComponentsIstioKubeSpecModel struct {
	ConfigurationDirectory types.String `tfsdk:"configuration_directory"`
	ConfigurationFileName  types.String `tfsdk:"configuration_file_name"`
	Revision               types.String `tfsdk:"revision"`
	BinaryDirectory        types.String `tfsdk:"binary_directory"`
	Chained                types.Bool   `tfsdk:"chained"`
	ClusterRole            types.String `tfsdk:"cluster_role"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                                `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                      `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                  `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                    `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                      `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                  `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                          `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                    `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                                        `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                          `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                    `tfsdk:"topology_key"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
}
type FieldRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel struct {
	FieldPath  types.String `tfsdk:"field_path"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                            `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel struct {
	Resource      types.String                                                                                     `tfsdk:"resource"`
	ContainerName types.String                                                                                     `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                              `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                `tfsdk:"optional"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel struct {
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                   `tfsdk:"optional"`
	Key                  types.String                                                                                                 `tfsdk:"key"`
}
type ValueFromInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvModel struct {
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	Value     types.String                                                              `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                              `tfsdk:"name"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchLabels      types.Map                                                                                                    `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
}
type TargetInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                        `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                         `tfsdk:"type"`
}
type ObjectInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	Selector     SelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                  `tfsdk:"metric_name"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                             `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                  `tfsdk:"match_labels"`
}
type PodsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                      `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                          `tfsdk:"type"`
}
type TargetInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	AverageUtilization types.Int64                                                                                           `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                          `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                       `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                 `tfsdk:"type"`
}
type ResourceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	Name                     types.String                                                                                                `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchLabels      types.Map                                                                                                            `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
}
type ExternalInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                          `tfsdk:"metric_name"`
}
type MetricsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecModel struct {
	Object   ObjectInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                        `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
}
type HpaSpecInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	MinReplicas types.Int64                                                                    `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                    `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
}
type TolerationsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
}
type SysctlsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	FsGroupChangePolicy types.String                                                                                  `tfsdk:"fs_group_change_policy"`
	SupplementalGroups  types.List                                                                                    `tfsdk:"supplemental_groups"`
	RunAsUser           types.Int64                                                                                   `tfsdk:"run_as_user"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	RunAsNonRoot        types.Bool                                                                                    `tfsdk:"run_as_non_root"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	FsGroup             types.Int64                                                                                   `tfsdk:"fs_group"`
	RunAsGroup          types.Int64                                                                                   `tfsdk:"run_as_group"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
}
type ResourcesInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                            `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                      `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyModel struct {
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
}
type StrategyInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                       `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	ReadOnlyRootFilesystem   types.Bool                                                                                          `tfsdk:"read_only_root_filesystem"`
	RunAsUser                types.Int64                                                                                         `tfsdk:"run_as_user"`
	RunAsNonRoot             types.Bool                                                                                          `tfsdk:"run_as_non_root"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	ProcMount                types.String                                                                                        `tfsdk:"proc_mount"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	Privileged               types.Bool                                                                                          `tfsdk:"privileged"`
	AllowPrivilegeEscalation types.Bool                                                                                          `tfsdk:"allow_privilege_escalation"`
	RunAsGroup               types.Int64                                                                                         `tfsdk:"run_as_group"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
}
type DeploymentInstallTemplateHelmSpecComponentsIstioKubeSpecModel struct {
	ReplicaCount             types.Int64                                                                           `tfsdk:"replica_count"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel                `tfsdk:"resources"`
	PodAnnotations           types.Map                                                                             `tfsdk:"pod_annotations"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel                   `tfsdk:"env"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel           `tfsdk:"tolerations"`
}
type PatchesInstallTemplateHelmSpecComponentsIstioKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsIstioKubeSpecModel struct {
	Kind       types.String                                                          `tfsdk:"kind"`
	Name       types.String                                                          `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsIstioKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                          `tfsdk:"api_version"`
}
type KubeSpecInstallTemplateHelmSpecComponentsIstioModel struct {
	Deployment     DeploymentInstallTemplateHelmSpecComponentsIstioKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsIstioKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsIstioKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsIstioKubeSpecModel `tfsdk:"service_account"`
	CNI            CNIInstallTemplateHelmSpecComponentsIstioKubeSpecModel            `tfsdk:"cni"`
}
type PatchesInstallTemplateHelmSpecComponentsIstioBaseOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type BaseOverlaysInstallTemplateHelmSpecComponentsIstioModel struct {
	Name       types.String                                                      `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsIstioBaseOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                      `tfsdk:"api_version"`
	Kind       types.String                                                      `tfsdk:"kind"`
}
type MaxWorkloadCert_TTLInstallTemplateHelmSpecComponentsIstioModel struct {
	Seconds types.Int64 `tfsdk:"seconds"`
	Nanos   types.Int64 `tfsdk:"nanos"`
}
type PatchesInstallTemplateHelmSpecComponentsIstio_CNIOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type CNIOverlaysInstallTemplateHelmSpecComponentsIstioModel struct {
	Name       types.String                                                      `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsIstio_CNIOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                      `tfsdk:"api_version"`
	Kind       types.String                                                      `tfsdk:"kind"`
}
type IstioInstallTemplateHelmSpecComponentsModel struct {
	MaxWorkloadCert_TTL     MaxWorkloadCert_TTLInstallTemplateHelmSpecComponentsIstioModel     `tfsdk:"max_workload_cert_ttl"`
	TsbVersion              types.String                                                       `tfsdk:"tsb_version"`
	KubeSpec                KubeSpecInstallTemplateHelmSpecComponentsIstioModel                `tfsdk:"kube_spec"`
	BaseOverlays            []*BaseOverlaysInstallTemplateHelmSpecComponentsIstioModel         `tfsdk:"base_overlays"`
	TraceSamplingRate       types.Float64                                                      `tfsdk:"trace_sampling_rate"`
	TrustDomain             types.String                                                       `tfsdk:"trust_domain"`
	CniOverlays             []*CNIOverlaysInstallTemplateHelmSpecComponentsIstioModel          `tfsdk:"cni_overlays"`
	PilotOverlays           []*PilotOverlaysInstallTemplateHelmSpecComponentsIstioModel        `tfsdk:"pilot_overlays"`
	DefaultWorkloadCert_TTL DefaultWorkloadCert_TTLInstallTemplateHelmSpecComponentsIstioModel `tfsdk:"default_workload_cert_ttl"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsOapKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsOapKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsOapKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromModel struct {
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                 `tfsdk:"optional"`
	Key                  types.String                                                                                               `tfsdk:"key"`
}
type FieldRefInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromModel struct {
	FieldPath  types.String `tfsdk:"field_path"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	Type   types.Int64                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                                   `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                   `tfsdk:"resource"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                            `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                              `tfsdk:"optional"`
}
type ValueFromInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvModel struct {
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                            `tfsdk:"name"`
	Value     types.String                                                            `tfsdk:"value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                    `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                              `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                  `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                    `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                              `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                        `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                                  `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                    `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                      `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                        `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	TopologyKey   types.String                                                                                                                                  `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                    `tfsdk:"namespaces"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                              `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
}
type TolerationsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchLabels      types.Map                                                                                                `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
}
type PodsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                    `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                               `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                               `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	Type   types.Int64                                                                                        `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
}
type TargetInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	AverageUtilization types.Int64                                                                                         `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                        `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                     `tfsdk:"type"`
}
type ResourceInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                              `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                               `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                        `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                          `tfsdk:"match_labels"`
}
type ExternalInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                        `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                       `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchLabels      types.Map                                                                                                  `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
}
type TargetInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
}
type ObjectInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel struct {
	AverageValue AverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
}
type MetricsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecModel struct {
	External ExternalInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                      `tfsdk:"type"`
}
type HpaSpecInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                  `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                  `tfsdk:"max_replicas"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type SysctlsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	FsGroup             types.Int64                                                                                 `tfsdk:"fs_group"`
	SupplementalGroups  types.List                                                                                  `tfsdk:"supplemental_groups"`
	FsGroupChangePolicy types.String                                                                                `tfsdk:"fs_group_change_policy"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsNonRoot        types.Bool                                                                                  `tfsdk:"run_as_non_root"`
	RunAsUser           types.Int64                                                                                 `tfsdk:"run_as_user"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	RunAsGroup          types.Int64                                                                                 `tfsdk:"run_as_group"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel struct {
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	ReadOnlyRootFilesystem   types.Bool                                                                                        `tfsdk:"read_only_root_filesystem"`
	Privileged               types.Bool                                                                                        `tfsdk:"privileged"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	RunAsGroup               types.Int64                                                                                       `tfsdk:"run_as_group"`
	ProcMount                types.String                                                                                      `tfsdk:"proc_mount"`
	AllowPrivilegeEscalation types.Bool                                                                                        `tfsdk:"allow_privilege_escalation"`
	RunAsUser                types.Int64                                                                                       `tfsdk:"run_as_user"`
	RunAsNonRoot             types.Bool                                                                                        `tfsdk:"run_as_non_root"`
}
type ResourcesInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                          `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                    `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyModel struct {
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
}
type StrategyInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	Type          types.String                                                                     `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type DeploymentInstallTemplateHelmSpecComponentsOapKubeSpecModel struct {
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel                `tfsdk:"resources"`
	PodAnnotations           types.Map                                                                           `tfsdk:"pod_annotations"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel                   `tfsdk:"env"`
	ReplicaCount             types.Int64                                                                         `tfsdk:"replica_count"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
}
type PatchesInstallTemplateHelmSpecComponentsOapKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsOapKubeSpecModel struct {
	ApiVersion types.String                                                        `tfsdk:"api_version"`
	Kind       types.String                                                        `tfsdk:"kind"`
	Name       types.String                                                        `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsOapKubeSpecOverlaysModel `tfsdk:"patches"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsOapKubeSpecServicePortsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                   `tfsdk:"type"`
}
type PortsInstallTemplateHelmSpecComponentsOapKubeSpecServiceModel struct {
	Port       types.Int64                                                             `tfsdk:"port"`
	Protocol   types.String                                                            `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsOapKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                            `tfsdk:"name"`
	NodePort   types.Int64                                                             `tfsdk:"node_port"`
}
type ServiceInstallTemplateHelmSpecComponentsOapKubeSpecModel struct {
	Type        types.String                                                     `tfsdk:"type"`
	Annotations types.Map                                                        `tfsdk:"annotations"`
	Labels      types.Map                                                        `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsOapKubeSpecServiceModel `tfsdk:"ports"`
}
type KubeSpecInstallTemplateHelmSpecComponentsOapModel struct {
	Service        ServiceInstallTemplateHelmSpecComponentsOapKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsOapKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsOapKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsOapKubeSpecModel    `tfsdk:"overlays"`
}
type OapInstallTemplateHelmSpecComponentsModel struct {
	OnDemandEnvoyMetricsEnabled types.Bool                                        `tfsdk:"on_demand_envoy_metrics_enabled"`
	StorageIndexMergingEnabled  types.Bool                                        `tfsdk:"storage_index_merging_enabled"`
	StreamingLogEnabled         types.Bool                                        `tfsdk:"streaming_log_enabled"`
	KubeSpec                    KubeSpecInstallTemplateHelmSpecComponentsOapModel `tfsdk:"kube_spec"`
}
type PatchesInstallTemplateHelmSpecComponentsCollectorKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsCollectorKubeSpecModel struct {
	ApiVersion types.String                                                              `tfsdk:"api_version"`
	Kind       types.String                                                              `tfsdk:"kind"`
	Name       types.String                                                              `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsCollectorKubeSpecOverlaysModel `tfsdk:"patches"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsCollectorKubeSpecServicePortsModel struct {
	Type   types.Int64                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
}
type PortsInstallTemplateHelmSpecComponentsCollectorKubeSpecServiceModel struct {
	Protocol   types.String                                                                  `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsCollectorKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                  `tfsdk:"name"`
	NodePort   types.Int64                                                                   `tfsdk:"node_port"`
	Port       types.Int64                                                                   `tfsdk:"port"`
}
type ServiceInstallTemplateHelmSpecComponentsCollectorKubeSpecModel struct {
	Annotations types.Map                                                              `tfsdk:"annotations"`
	Labels      types.Map                                                              `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsCollectorKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                           `tfsdk:"type"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsCollectorKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsCollectorKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsCollectorKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type FieldRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	Type   types.Int64                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel struct {
	Resource      types.String                                                                                         `tfsdk:"resource"`
	ContainerName types.String                                                                                         `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                    `tfsdk:"optional"`
	Key                  types.String                                                                                                  `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                     `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                       `tfsdk:"optional"`
}
type ValueFromInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvModel struct {
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	Value     types.String                                                                  `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                  `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                          `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                           `tfsdk:"type"`
}
type TolerationsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                          `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityModel struct {
	Namespaces    types.List                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                    `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                          `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                    `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                        `tfsdk:"weight"`
}
type PodAffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                              `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	TopologyKey   types.String                                                                                                                                        `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                          `tfsdk:"namespaces"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                              `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                          `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                        `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                            `tfsdk:"weight"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                    `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	Privileged               types.Bool                                                                                              `tfsdk:"privileged"`
	RunAsUser                types.Int64                                                                                             `tfsdk:"run_as_user"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	AllowPrivilegeEscalation types.Bool                                                                                              `tfsdk:"allow_privilege_escalation"`
	RunAsGroup               types.Int64                                                                                             `tfsdk:"run_as_group"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	ProcMount                types.String                                                                                            `tfsdk:"proc_mount"`
	RunAsNonRoot             types.Bool                                                                                              `tfsdk:"run_as_non_root"`
	ReadOnlyRootFilesystem   types.Bool                                                                                              `tfsdk:"read_only_root_filesystem"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
}
type ResourcesInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchLabels      types.Map                                                                                                      `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
}
type PodsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	Selector           SelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                          `tfsdk:"metric_name"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	Type   types.Int64                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
}
type TargetInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	AverageUtilization types.Int64                                                                                               `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                              `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                     `tfsdk:"type"`
}
type ResourceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	Target                   TargetInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                    `tfsdk:"name"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                `tfsdk:"match_labels"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	Type   types.Int64                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
}
type ExternalInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                              `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                        `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                            `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Type   types.Int64                                                                                             `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
}
type ObjectInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName   types.String                                                                                      `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
}
type MetricsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecModel struct {
	Pods     PodsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                            `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
}
type HpaSpecInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	MinReplicas types.Int64                                                                        `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                        `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SysctlsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	FsGroupChangePolicy types.String                                                                                      `tfsdk:"fs_group_change_policy"`
	RunAsUser           types.Int64                                                                                       `tfsdk:"run_as_user"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	FsGroup             types.Int64                                                                                       `tfsdk:"fs_group"`
	RunAsNonRoot        types.Bool                                                                                        `tfsdk:"run_as_non_root"`
	RunAsGroup          types.Int64                                                                                       `tfsdk:"run_as_group"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	SupplementalGroups  types.List                                                                                        `tfsdk:"supplemental_groups"`
}
type DeploymentInstallTemplateHelmSpecComponentsCollectorKubeSpecModel struct {
	PodAnnotations           types.Map                                                                                 `tfsdk:"pod_annotations"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel                   `tfsdk:"env"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	ReplicaCount             types.Int64                                                                               `tfsdk:"replica_count"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel                `tfsdk:"resources"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel                 `tfsdk:"affinity"`
}
type KubeSpecInstallTemplateHelmSpecComponentsCollectorModel struct {
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsCollectorKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsCollectorKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsCollectorKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsCollectorKubeSpecModel     `tfsdk:"deployment"`
}
type CollectorInstallTemplateHelmSpecComponentsModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsCollectorModel `tfsdk:"kube_spec"`
}
type BatchWindowInstallTemplateHelmSpecComponentsGitopsModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type ManagementplaneRequestTimeoutInstallTemplateHelmSpecComponentsGitopsModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type ReconcileIntervalInstallTemplateHelmSpecComponentsGitopsModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type ReconcileRequestTimeoutInstallTemplateHelmSpecComponentsGitopsModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type WebhookTimeoutInstallTemplateHelmSpecComponentsGitopsModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type GitopsInstallTemplateHelmSpecComponentsModel struct {
	BatchWindow                   BatchWindowInstallTemplateHelmSpecComponentsGitopsModel                   `tfsdk:"batch_window"`
	Enabled                       types.Bool                                                                `tfsdk:"enabled"`
	ManagementplaneRequestTimeout ManagementplaneRequestTimeoutInstallTemplateHelmSpecComponentsGitopsModel `tfsdk:"managementplane_request_timeout"`
	ReconcileInterval             ReconcileIntervalInstallTemplateHelmSpecComponentsGitopsModel             `tfsdk:"reconcile_interval"`
	ReconcileRequestTimeout       ReconcileRequestTimeoutInstallTemplateHelmSpecComponentsGitopsModel       `tfsdk:"reconcile_request_timeout"`
	WebhookTimeout                WebhookTimeoutInstallTemplateHelmSpecComponentsGitopsModel                `tfsdk:"webhook_timeout"`
}
type ResourcesInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	Requests types.Map `tfsdk:"requests"`
	Limits   types.Map `tfsdk:"limits"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                          `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyRollingUpdateModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	Type          types.String                                                                           `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type TolerationsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                    `tfsdk:"optional"`
	Key                  types.String                                                                                                  `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                     `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                       `tfsdk:"optional"`
}
type FieldRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                                         `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                         `tfsdk:"resource"`
}
type ValueFromInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvModel struct {
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	Name      types.String                                                                  `tfsdk:"name"`
	Value     types.String                                                                  `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                              `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                          `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                        `tfsdk:"topology_key"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                              `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                          `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                        `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                                            `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                    `tfsdk:"weight"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                          `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                    `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                        `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                          `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityModel struct {
	Namespaces    types.List                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                    `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel struct {
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	AllowPrivilegeEscalation types.Bool                                                                                              `tfsdk:"allow_privilege_escalation"`
	ProcMount                types.String                                                                                            `tfsdk:"proc_mount"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	RunAsUser                types.Int64                                                                                             `tfsdk:"run_as_user"`
	Privileged               types.Bool                                                                                              `tfsdk:"privileged"`
	RunAsNonRoot             types.Bool                                                                                              `tfsdk:"run_as_non_root"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	ReadOnlyRootFilesystem   types.Bool                                                                                              `tfsdk:"read_only_root_filesystem"`
	RunAsGroup               types.Int64                                                                                             `tfsdk:"run_as_group"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
}
type ExternalInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                              `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Type   types.Int64                                                                                            `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                             `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                        `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
}
type ObjectInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel struct {
	Target       TargetInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                      `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                      `tfsdk:"match_labels"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
}
type PodsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                          `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                              `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type TargetInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type               types.String                                                                                              `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                               `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                           `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                     `tfsdk:"type"`
}
type ResourceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                    `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
}
type MetricsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecModel struct {
	External ExternalInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                            `tfsdk:"type"`
}
type HpaSpecInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                        `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                        `tfsdk:"max_replicas"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SysctlsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	SupplementalGroups  types.List                                                                                        `tfsdk:"supplemental_groups"`
	RunAsGroup          types.Int64                                                                                       `tfsdk:"run_as_group"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	FsGroup             types.Int64                                                                                       `tfsdk:"fs_group"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	FsGroupChangePolicy types.String                                                                                      `tfsdk:"fs_group_change_policy"`
	RunAsUser           types.Int64                                                                                       `tfsdk:"run_as_user"`
	RunAsNonRoot        types.Bool                                                                                        `tfsdk:"run_as_non_root"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
}
type DeploymentInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel struct {
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel                `tfsdk:"resources"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel                   `tfsdk:"env"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	ReplicaCount             types.Int64                                                                               `tfsdk:"replica_count"`
	PodAnnotations           types.Map                                                                                 `tfsdk:"pod_annotations"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel                 `tfsdk:"strategy"`
}
type PatchesInstallTemplateHelmSpecComponentsSatelliteKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel struct {
	Name       types.String                                                              `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsSatelliteKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                              `tfsdk:"api_version"`
	Kind       types.String                                                              `tfsdk:"kind"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsSatelliteKubeSpecServicePortsModel struct {
	Type   types.Int64                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
}
type PortsInstallTemplateHelmSpecComponentsSatelliteKubeSpecServiceModel struct {
	Port       types.Int64                                                                   `tfsdk:"port"`
	Protocol   types.String                                                                  `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsSatelliteKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                  `tfsdk:"name"`
	NodePort   types.Int64                                                                   `tfsdk:"node_port"`
}
type ServiceInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel struct {
	Ports       []*PortsInstallTemplateHelmSpecComponentsSatelliteKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                           `tfsdk:"type"`
	Annotations types.Map                                                              `tfsdk:"annotations"`
	Labels      types.Map                                                              `tfsdk:"labels"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsSatelliteKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsSatelliteKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type KubeSpecInstallTemplateHelmSpecComponentsSatelliteModel struct {
	Deployment     DeploymentInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel `tfsdk:"service_account"`
}
type SatelliteInstallTemplateHelmSpecComponentsModel struct {
	Enabled  types.Bool                                              `tfsdk:"enabled"`
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsSatelliteModel `tfsdk:"kube_spec"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsNgacKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsNgacKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsNgacKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type ResourcesInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	Requests types.Map `tfsdk:"requests"`
	Limits   types.Map `tfsdk:"limits"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type PreferenceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                               `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                     `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                               `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                 `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                   `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                     `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                 `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                               `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                         `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                     `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                   `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                       `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                         `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Namespaces    types.List                                                                                                                                     `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                   `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
}
type TolerationsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
}
type FieldRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                           `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                                    `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                    `tfsdk:"resource"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                             `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                               `tfsdk:"optional"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                  `tfsdk:"optional"`
	Key                  types.String                                                                                                `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
}
type ValueFromInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvModel struct {
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	Name      types.String                                                             `tfsdk:"name"`
	Value     types.String                                                             `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                      `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchLabels      types.Map                                                                                                           `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                         `tfsdk:"type"`
}
type ExternalInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                         `tfsdk:"metric_name"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                        `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                   `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                       `tfsdk:"type"`
}
type ObjectInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel struct {
	AverageValue AverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                 `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                 `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                            `tfsdk:"type"`
}
type PodsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                     `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	Type   types.Int64                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type TargetInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Value              ValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                          `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                         `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                      `tfsdk:"type"`
}
type ResourceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                               `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
}
type MetricsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecModel struct {
	Type     types.String                                                                       `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
}
type HpaSpecInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	MaxReplicas types.Int64                                                                   `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                   `tfsdk:"min_replicas"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	ReadOnlyRootFilesystem   types.Bool                                                                                         `tfsdk:"read_only_root_filesystem"`
	Privileged               types.Bool                                                                                         `tfsdk:"privileged"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	ProcMount                types.String                                                                                       `tfsdk:"proc_mount"`
	RunAsGroup               types.Int64                                                                                        `tfsdk:"run_as_group"`
	AllowPrivilegeEscalation types.Bool                                                                                         `tfsdk:"allow_privilege_escalation"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	RunAsUser                types.Int64                                                                                        `tfsdk:"run_as_user"`
	RunAsNonRoot             types.Bool                                                                                         `tfsdk:"run_as_non_root"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SysctlsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	RunAsGroup          types.Int64                                                                                  `tfsdk:"run_as_group"`
	RunAsNonRoot        types.Bool                                                                                   `tfsdk:"run_as_non_root"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	SupplementalGroups  types.List                                                                                   `tfsdk:"supplemental_groups"`
	FsGroupChangePolicy types.String                                                                                 `tfsdk:"fs_group_change_policy"`
	RunAsUser           types.Int64                                                                                  `tfsdk:"run_as_user"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	FsGroup             types.Int64                                                                                  `tfsdk:"fs_group"`
}
type DeploymentInstallTemplateHelmSpecComponentsNgacKubeSpecModel struct {
	Resources                ResourcesInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel                `tfsdk:"resources"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel                   `tfsdk:"env"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	ReplicaCount             types.Int64                                                                          `tfsdk:"replica_count"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	PodAnnotations           types.Map                                                                            `tfsdk:"pod_annotations"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
}
type PatchesInstallTemplateHelmSpecComponentsNgacKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsNgacKubeSpecModel struct {
	Patches    []*PatchesInstallTemplateHelmSpecComponentsNgacKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                         `tfsdk:"api_version"`
	Kind       types.String                                                         `tfsdk:"kind"`
	Name       types.String                                                         `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsNgacKubeSpecServicePortsModel struct {
	Type   types.Int64                                                                    `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
}
type PortsInstallTemplateHelmSpecComponentsNgacKubeSpecServiceModel struct {
	NodePort   types.Int64                                                              `tfsdk:"node_port"`
	Port       types.Int64                                                              `tfsdk:"port"`
	Protocol   types.String                                                             `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsNgacKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                             `tfsdk:"name"`
}
type ServiceInstallTemplateHelmSpecComponentsNgacKubeSpecModel struct {
	Ports       []*PortsInstallTemplateHelmSpecComponentsNgacKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                      `tfsdk:"type"`
	Annotations types.Map                                                         `tfsdk:"annotations"`
	Labels      types.Map                                                         `tfsdk:"labels"`
}
type KubeSpecInstallTemplateHelmSpecComponentsNgacModel struct {
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsNgacKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsNgacKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsNgacKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsNgacKubeSpecModel     `tfsdk:"deployment"`
}
type NgacInstallTemplateHelmSpecComponentsModel struct {
	KubeSpec  KubeSpecInstallTemplateHelmSpecComponentsNgacModel `tfsdk:"kube_spec"`
	LogLevels types.Map                                          `tfsdk:"log_levels"`
	Enabled   types.Bool                                         `tfsdk:"enabled"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsDefaultKubeSpecAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type AccountInstallTemplateHelmSpecComponentsDefaultKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsDefaultKubeSpecAccountModel `tfsdk:"image_pull_secrets"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                            `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                        `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                      `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                          `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                            `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	TopologyKey   types.String                                                                                                                                      `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                        `tfsdk:"namespaces"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type PreferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                                  `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                        `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                    `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                  `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                      `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                        `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityModel struct {
	TopologyKey   types.String                                                                                                                                  `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                    `tfsdk:"namespaces"`
}
type PodAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel struct {
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel struct {
	ProcMount                types.String                                                                                          `tfsdk:"proc_mount"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	ReadOnlyRootFilesystem   types.Bool                                                                                            `tfsdk:"read_only_root_filesystem"`
	RunAsNonRoot             types.Bool                                                                                            `tfsdk:"run_as_non_root"`
	RunAsGroup               types.Int64                                                                                           `tfsdk:"run_as_group"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	Privileged               types.Bool                                                                                            `tfsdk:"privileged"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	RunAsUser                types.Int64                                                                                           `tfsdk:"run_as_user"`
	AllowPrivilegeEscalation types.Bool                                                                                            `tfsdk:"allow_privilege_escalation"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                     `tfsdk:"optional"`
	Key                  types.String                                                                                                   `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
}
type FieldRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type IntValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                              `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                                       `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                       `tfsdk:"resource"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                  `tfsdk:"optional"`
	Key                  types.String                                                                                                `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
}
type ValueFromInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvModel struct {
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel struct {
	Value     types.String                                                                `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                `tfsdk:"name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel struct {
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SysctlsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel struct {
	FsGroup             types.Int64                                                                                     `tfsdk:"fs_group"`
	FsGroupChangePolicy types.String                                                                                    `tfsdk:"fs_group_change_policy"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsNonRoot        types.Bool                                                                                      `tfsdk:"run_as_non_root"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	RunAsUser           types.Int64                                                                                     `tfsdk:"run_as_user"`
	RunAsGroup          types.Int64                                                                                     `tfsdk:"run_as_group"`
	SupplementalGroups  types.List                                                                                      `tfsdk:"supplemental_groups"`
}
type IntValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                        `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                         `tfsdk:"type"`
}
type TolerationsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel struct {
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
}
type DeploymentInstallTemplateHelmSpecComponentsDefaultKubeSpecModel struct {
	Strategy                 StrategyInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel                   `tfsdk:"env"`
	PodAnnotations           types.Map                                                                               `tfsdk:"pod_annotations"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
}
type TolerationsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel struct {
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                 `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                             `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                           `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityModel struct {
	Weight          types.Int64                                                                                                                               `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                 `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                             `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                           `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                     `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityModel struct {
	TopologyKey   types.String                                                                                                                               `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                 `tfsdk:"namespaces"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                     `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                 `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                               `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                   `tfsdk:"weight"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                           `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel struct {
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityModel     `tfsdk:"pod_affinity"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel struct {
	AllowPrivilegeEscalation types.Bool                                                                                     `tfsdk:"allow_privilege_escalation"`
	ProcMount                types.String                                                                                   `tfsdk:"proc_mount"`
	RunAsNonRoot             types.Bool                                                                                     `tfsdk:"run_as_non_root"`
	RunAsUser                types.Int64                                                                                    `tfsdk:"run_as_user"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	ReadOnlyRootFilesystem   types.Bool                                                                                     `tfsdk:"read_only_root_filesystem"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsGroup               types.Int64                                                                                    `tfsdk:"run_as_group"`
	Privileged               types.Bool                                                                                     `tfsdk:"privileged"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel   `tfsdk:"capabilities"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel `tfsdk:"windows_options"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SysctlsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel struct {
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel `tfsdk:"se_linux_options"`
	FsGroup             types.Int64                                                                              `tfsdk:"fs_group"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel     `tfsdk:"sysctls"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel `tfsdk:"seccomp_profile"`
	FsGroupChangePolicy types.String                                                                             `tfsdk:"fs_group_change_policy"`
	RunAsUser           types.Int64                                                                              `tfsdk:"run_as_user"`
	SupplementalGroups  types.List                                                                               `tfsdk:"supplemental_groups"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel `tfsdk:"windows_options"`
	RunAsNonRoot        types.Bool                                                                               `tfsdk:"run_as_non_root"`
	RunAsGroup          types.Int64                                                                              `tfsdk:"run_as_group"`
}
type JobInstallTemplateHelmSpecComponentsDefaultKubeSpecModel struct {
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel           `tfsdk:"tolerations"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel                 `tfsdk:"affinity"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel `tfsdk:"container_security_context"`
	PodAnnotations           types.Map                                                                        `tfsdk:"pod_annotations"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel       `tfsdk:"pod_security_context"`
}
type ServiceInstallTemplateHelmSpecComponentsDefaultKubeSpecModel struct {
	Annotations types.Map `tfsdk:"annotations"`
}
type DefaultKubeSpecInstallTemplateHelmSpecComponentsModel struct {
	Account    AccountInstallTemplateHelmSpecComponentsDefaultKubeSpecModel    `tfsdk:"account"`
	Deployment DeploymentInstallTemplateHelmSpecComponentsDefaultKubeSpecModel `tfsdk:"deployment"`
	Job        JobInstallTemplateHelmSpecComponentsDefaultKubeSpecModel        `tfsdk:"job"`
	Service    ServiceInstallTemplateHelmSpecComponentsDefaultKubeSpecModel    `tfsdk:"service"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type ResourcesInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	Requests types.Map `tfsdk:"requests"`
	Limits   types.Map `tfsdk:"limits"`
}
type FieldRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                 `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                                          `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                          `tfsdk:"resource"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                   `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                     `tfsdk:"optional"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                      `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                        `tfsdk:"optional"`
}
type ValueFromInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvModel struct {
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                   `tfsdk:"name"`
	Value     types.String                                                                   `tfsdk:"value"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SysctlsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	RunAsNonRoot        types.Bool                                                                                         `tfsdk:"run_as_non_root"`
	FsGroup             types.Int64                                                                                        `tfsdk:"fs_group"`
	RunAsUser           types.Int64                                                                                        `tfsdk:"run_as_user"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsGroup          types.Int64                                                                                        `tfsdk:"run_as_group"`
	SupplementalGroups  types.List                                                                                         `tfsdk:"supplemental_groups"`
	FsGroupChangePolicy types.String                                                                                       `tfsdk:"fs_group_change_policy"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type PreferenceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                                     `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                           `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                       `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                     `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityModel struct {
	Weight          types.Int64                                                                                                                                         `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                           `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityModel struct {
	TopologyKey   types.String                                                                                                                                     `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                       `tfsdk:"namespaces"`
}
type PodAffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                               `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                           `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                         `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                                             `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                               `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	TopologyKey   types.String                                                                                                                                         `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                           `tfsdk:"namespaces"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyModel struct {
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
}
type StrategyInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                            `tfsdk:"type"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	RunAsGroup               types.Int64                                                                                              `tfsdk:"run_as_group"`
	AllowPrivilegeEscalation types.Bool                                                                                               `tfsdk:"allow_privilege_escalation"`
	ReadOnlyRootFilesystem   types.Bool                                                                                               `tfsdk:"read_only_root_filesystem"`
	Privileged               types.Bool                                                                                               `tfsdk:"privileged"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsUser                types.Int64                                                                                              `tfsdk:"run_as_user"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	RunAsNonRoot             types.Bool                                                                                               `tfsdk:"run_as_non_root"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	ProcMount                types.String                                                                                             `tfsdk:"proc_mount"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                              `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                         `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                             `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
}
type ObjectInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel struct {
	AverageValue AverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                       `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                  `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                       `tfsdk:"match_labels"`
}
type PodsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel struct {
	Selector           SelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                           `tfsdk:"metric_name"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                               `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type TargetInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type               types.String                                                                                               `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                            `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type ResourceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                     `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                               `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                 `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	Type   types.Int64                                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
}
type ExternalInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                               `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
}
type MetricsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecModel struct {
	Pods     PodsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                             `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
}
type HpaSpecInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	MinReplicas types.Int64                                                                         `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                         `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
}
type TolerationsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
}
type DeploymentInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel struct {
	Resources                ResourcesInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel                `tfsdk:"resources"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	PodAnnotations           types.Map                                                                                  `tfsdk:"pod_annotations"`
	ReplicaCount             types.Int64                                                                                `tfsdk:"replica_count"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel                   `tfsdk:"env"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel           `tfsdk:"tolerations"`
}
type PatchesInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel struct {
	Kind       types.String                                                               `tfsdk:"kind"`
	Name       types.String                                                               `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                               `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServicePortsModel struct {
	Type   types.Int64                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
}
type PortsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServiceModel struct {
	Port       types.Int64                                                                    `tfsdk:"port"`
	Protocol   types.String                                                                   `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                   `tfsdk:"name"`
	NodePort   types.Int64                                                                    `tfsdk:"node_port"`
}
type ServiceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel struct {
	Ports       []*PortsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                            `tfsdk:"type"`
	Annotations types.Map                                                               `tfsdk:"annotations"`
	Labels      types.Map                                                               `tfsdk:"labels"`
}
type KubeSpecInstallTemplateHelmSpecComponentsHpaAdapterModel struct {
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel        `tfsdk:"service"`
}
type HpaAdapterInstallTemplateHelmSpecComponentsModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsHpaAdapterModel `tfsdk:"kube_spec"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	Privileged               types.Bool                                                                                                    `tfsdk:"privileged"`
	RunAsNonRoot             types.Bool                                                                                                    `tfsdk:"run_as_non_root"`
	RunAsGroup               types.Int64                                                                                                   `tfsdk:"run_as_group"`
	ProcMount                types.String                                                                                                  `tfsdk:"proc_mount"`
	RunAsUser                types.Int64                                                                                                   `tfsdk:"run_as_user"`
	AllowPrivilegeEscalation types.Bool                                                                                                    `tfsdk:"allow_privilege_escalation"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                    `tfsdk:"read_only_root_filesystem"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                          `tfsdk:"optional"`
	Key                  types.String                                                                                                        `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromModel struct {
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                             `tfsdk:"optional"`
	Key                  types.String                                                                                                           `tfsdk:"key"`
}
type FieldRefInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromModel struct {
	FieldPath  types.String `tfsdk:"field_path"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                      `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                                               `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                               `tfsdk:"resource"`
}
type ValueFromInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvModel struct {
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	Value     types.String                                                                        `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                        `tfsdk:"name"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                    `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchLabels      types.Map                                                                                                                      `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
}
type ExternalInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                                    `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchLabels      types.Map                                                                                                              `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
}
type TargetInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                  `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Type   types.Int64                                                                                                   `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
}
type ObjectInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName   types.String                                                                                            `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                            `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                       `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
}
type PodsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                                `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                           `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                    `tfsdk:"type"`
}
type TargetInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type               types.String                                                                                                    `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                     `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type ResourceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel struct {
	Target                   TargetInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                          `tfsdk:"name"`
}
type MetricsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecModel struct {
	Pods     PodsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                                  `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
}
type HpaSpecInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                              `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                              `tfsdk:"max_replicas"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type SysctlsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	SupplementalGroups  types.List                                                                                              `tfsdk:"supplemental_groups"`
	FsGroupChangePolicy types.String                                                                                            `tfsdk:"fs_group_change_policy"`
	RunAsNonRoot        types.Bool                                                                                              `tfsdk:"run_as_non_root"`
	FsGroup             types.Int64                                                                                             `tfsdk:"fs_group"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsUser           types.Int64                                                                                             `tfsdk:"run_as_user"`
	RunAsGroup          types.Int64                                                                                             `tfsdk:"run_as_group"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                            `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                          `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                              `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                            `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                          `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                    `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                              `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                                                  `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                    `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Namespaces    types.List                                                                                                                                                `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                              `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                          `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
}
type TolerationsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                                 `tfsdk:"type"`
}
type ResourcesInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	Requests types.Map `tfsdk:"requests"`
	Limits   types.Map `tfsdk:"limits"`
}
type DeploymentInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel struct {
	Env                      []*EnvInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel                   `tfsdk:"env"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	ReplicaCount             types.Int64                                                                                     `tfsdk:"replica_count"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel                `tfsdk:"resources"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	PodAnnotations           types.Map                                                                                       `tfsdk:"pod_annotations"`
}
type PatchesInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel struct {
	Kind       types.String                                                                    `tfsdk:"kind"`
	Name       types.String                                                                    `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                    `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServicePortsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                               `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
}
type PortsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServiceModel struct {
	Name       types.String                                                                        `tfsdk:"name"`
	NodePort   types.Int64                                                                         `tfsdk:"node_port"`
	Port       types.Int64                                                                         `tfsdk:"port"`
	Protocol   types.String                                                                        `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServicePortsModel `tfsdk:"target_port"`
}
type ServiceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel struct {
	Annotations types.Map                                                                    `tfsdk:"annotations"`
	Labels      types.Map                                                                    `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                                 `tfsdk:"type"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type KubeSpecInstallTemplateHelmSpecComponentsRateLimitServerModel struct {
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel     `tfsdk:"deployment"`
}
type RedisInstallTemplateHelmSpecComponentsRateLimitServerBackendModel struct {
	Uri types.String `tfsdk:"uri"`
}
type BackendInstallTemplateHelmSpecComponentsRateLimitServerModel struct {
	Redis RedisInstallTemplateHelmSpecComponentsRateLimitServerBackendModel `tfsdk:"redis"`
}
type RateLimitServerInstallTemplateHelmSpecComponentsModel struct {
	Domain   types.String                                                  `tfsdk:"domain"`
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsRateLimitServerModel `tfsdk:"kube_spec"`
	Backend  BackendInstallTemplateHelmSpecComponentsRateLimitServerModel  `tfsdk:"backend"`
}
type CustomInstallTemplateHelmSpecComponentsInternalCertProviderModel struct {
	CsrSignerName      types.String `tfsdk:"csr_signer_name"`
	CaBundleSecretName types.String `tfsdk:"ca_bundle_secret_name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	RunAsNonRoot        types.Bool                                                                                                                                    `tfsdk:"run_as_non_root"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	FsGroup             types.Int64                                                                                                                                   `tfsdk:"fs_group"`
	RunAsGroup          types.Int64                                                                                                                                   `tfsdk:"run_as_group"`
	FsGroupChangePolicy types.String                                                                                                                                  `tfsdk:"fs_group_change_policy"`
	SupplementalGroups  types.List                                                                                                                                    `tfsdk:"supplemental_groups"`
	RunAsUser           types.Int64                                                                                                                                   `tfsdk:"run_as_user"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                                      `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                                  `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                                `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityModel struct {
	Weight          types.Int64                                                                                                                                                                                    `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                                                                      `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                  `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                                                          `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                                    `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                                                                                        `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                          `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Namespaces    types.List                                                                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                    `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                                                                `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                                                              `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                                                `tfsdk:"optional"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                                                                 `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                                                   `tfsdk:"optional"`
}
type FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                            `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                                                                                     `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                                                                     `tfsdk:"resource"`
}
type ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvModel struct {
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	Name      types.String                                                                                                              `tfsdk:"name"`
	Value     types.String                                                                                                              `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchLabels      types.Map                                                                                                                                                            `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                 `tfsdk:"type"`
}
type ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                                                          `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                    `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                        `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Type   types.Int64                                                                                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
}
type ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName   types.String                                                                                                                                  `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchLabels      types.Map                                                                                                                                                  `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                             `tfsdk:"type"`
}
type PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                                                      `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                       `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                 `tfsdk:"type"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                                                                          `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                                                           `tfsdk:"average_utilization"`
}
type ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                                                                `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
}
type MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecModel struct {
	Type     types.String                                                                                                                        `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
}
type HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	MaxReplicas types.Int64                                                                                                                    `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                                                                    `tfsdk:"min_replicas"`
}
type TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
}
type ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	RunAsNonRoot             types.Bool                                                                                                                                          `tfsdk:"run_as_non_root"`
	Privileged               types.Bool                                                                                                                                          `tfsdk:"privileged"`
	ProcMount                types.String                                                                                                                                        `tfsdk:"proc_mount"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                                                          `tfsdk:"read_only_root_filesystem"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	RunAsGroup               types.Int64                                                                                                                                         `tfsdk:"run_as_group"`
	RunAsUser                types.Int64                                                                                                                                         `tfsdk:"run_as_user"`
	AllowPrivilegeEscalation types.Bool                                                                                                                                          `tfsdk:"allow_privilege_escalation"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                                                                            `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyModel struct {
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
}
type StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	Type          types.String                                                                                                                       `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel struct {
	Affinity                 AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	ReplicaCount             types.Int64                                                                                                                           `tfsdk:"replica_count"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel                `tfsdk:"resources"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	PodAnnotations           types.Map                                                                                                                             `tfsdk:"pod_annotations"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel                   `tfsdk:"env"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel           `tfsdk:"tolerations"`
}
type PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel struct {
	Name       types.String                                                                                                          `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                                                          `tfsdk:"api_version"`
	Kind       types.String                                                                                                          `tfsdk:"kind"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServicePortsModel struct {
	Type   types.Int64                                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
}
type PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServiceModel struct {
	TargetPort TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                                                              `tfsdk:"name"`
	NodePort   types.Int64                                                                                                               `tfsdk:"node_port"`
	Port       types.Int64                                                                                                               `tfsdk:"port"`
	Protocol   types.String                                                                                                              `tfsdk:"protocol"`
}
type ServiceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel struct {
	Type        types.String                                                                                                       `tfsdk:"type"`
	Annotations types.Map                                                                                                          `tfsdk:"annotations"`
	Labels      types.Map                                                                                                          `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServiceModel `tfsdk:"ports"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type KubeSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecModel struct {
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel     `tfsdk:"deployment"`
}
type CertManagerWebhookSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecModel `tfsdk:"kube_spec"`
}
type ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                                                               `tfsdk:"optional"`
	Key                  types.String                                                                                                                                             `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                                                                `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                                                  `tfsdk:"optional"`
}
type FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	Type   types.Int64                                                                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                                                                                    `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                                                                    `tfsdk:"resource"`
}
type ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvModel struct {
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	Value     types.String                                                                                                             `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                                                             `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	AverageUtilization types.Int64                                                                                                                                          `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                                                                         `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
}
type ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                                                               `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                           `tfsdk:"match_labels"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	Type   types.Int64                                                                                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
}
type ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                                                                         `tfsdk:"metric_name"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                       `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Type   types.Int64                                                                                                                                        `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchLabels      types.Map                                                                                                                                                   `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
}
type ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	Selector     SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                                                                 `tfsdk:"metric_name"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                 `tfsdk:"match_labels"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                            `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
}
type PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                                                                     `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
}
type MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecModel struct {
	Type     types.String                                                                                                                       `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
}
type HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	MinReplicas types.Int64                                                                                                                   `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                                                                   `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	RunAsNonRoot        types.Bool                                                                                                                                   `tfsdk:"run_as_non_root"`
	FsGroupChangePolicy types.String                                                                                                                                 `tfsdk:"fs_group_change_policy"`
	RunAsUser           types.Int64                                                                                                                                  `tfsdk:"run_as_user"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsGroup          types.Int64                                                                                                                                  `tfsdk:"run_as_group"`
	FsGroup             types.Int64                                                                                                                                  `tfsdk:"fs_group"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	SupplementalGroups  types.List                                                                                                                                   `tfsdk:"supplemental_groups"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                           `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyModel struct {
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
}
type StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                                                                      `tfsdk:"type"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                                                                               `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                                     `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                                 `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                               `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                                                   `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                     `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityModel struct {
	TopologyKey   types.String                                                                                                                                                                               `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                 `tfsdk:"namespaces"`
}
type PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                                         `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                                     `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                                   `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                                                       `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                                                                         `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Namespaces    types.List                                                                                                                                                                                     `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                   `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	AllowPrivilegeEscalation types.Bool                                                                                                                                         `tfsdk:"allow_privilege_escalation"`
	ProcMount                types.String                                                                                                                                       `tfsdk:"proc_mount"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	Privileged               types.Bool                                                                                                                                         `tfsdk:"privileged"`
	RunAsUser                types.Int64                                                                                                                                        `tfsdk:"run_as_user"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsGroup               types.Int64                                                                                                                                        `tfsdk:"run_as_group"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                                                         `tfsdk:"read_only_root_filesystem"`
	RunAsNonRoot             types.Bool                                                                                                                                         `tfsdk:"run_as_non_root"`
}
type DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel struct {
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel                   `tfsdk:"env"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel                `tfsdk:"resources"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	ReplicaCount             types.Int64                                                                                                                          `tfsdk:"replica_count"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	PodAnnotations           types.Map                                                                                                                            `tfsdk:"pod_annotations"`
}
type PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel struct {
	ApiVersion types.String                                                                                                         `tfsdk:"api_version"`
	Kind       types.String                                                                                                         `tfsdk:"kind"`
	Name       types.String                                                                                                         `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecOverlaysModel `tfsdk:"patches"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServicePortsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                    `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
}
type PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServiceModel struct {
	NodePort   types.Int64                                                                                                              `tfsdk:"node_port"`
	Port       types.Int64                                                                                                              `tfsdk:"port"`
	Protocol   types.String                                                                                                             `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                                                             `tfsdk:"name"`
}
type ServiceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel struct {
	Labels      types.Map                                                                                                         `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                                                                      `tfsdk:"type"`
	Annotations types.Map                                                                                                         `tfsdk:"annotations"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type KubeSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorModel struct {
	Deployment     DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel `tfsdk:"service_account"`
}
type CertManagerCaInjectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorModel `tfsdk:"kube_spec"`
}
type ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsUser                types.Int64                                                                                                                                  `tfsdk:"run_as_user"`
	RunAsGroup               types.Int64                                                                                                                                  `tfsdk:"run_as_group"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	ProcMount                types.String                                                                                                                                 `tfsdk:"proc_mount"`
	AllowPrivilegeEscalation types.Bool                                                                                                                                   `tfsdk:"allow_privilege_escalation"`
	RunAsNonRoot             types.Bool                                                                                                                                   `tfsdk:"run_as_non_root"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                                                   `tfsdk:"read_only_root_filesystem"`
	Privileged               types.Bool                                                                                                                                   `tfsdk:"privileged"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	RunAsUser           types.Int64                                                                                                                            `tfsdk:"run_as_user"`
	RunAsNonRoot        types.Bool                                                                                                                             `tfsdk:"run_as_non_root"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	SupplementalGroups  types.List                                                                                                                             `tfsdk:"supplemental_groups"`
	FsGroupChangePolicy types.String                                                                                                                           `tfsdk:"fs_group_change_policy"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	FsGroup             types.Int64                                                                                                                            `tfsdk:"fs_group"`
	RunAsGroup          types.Int64                                                                                                                            `tfsdk:"run_as_group"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                               `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyModel struct {
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
}
type StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	Type          types.String                                                                                                                `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromModel struct {
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                                         `tfsdk:"optional"`
	Key                  types.String                                                                                                                                       `tfsdk:"key"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromModel struct {
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                                            `tfsdk:"optional"`
	Key                  types.String                                                                                                                                          `tfsdk:"key"`
}
type FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                     `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                                                                              `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                                                              `tfsdk:"resource"`
}
type ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvModel struct {
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                                                       `tfsdk:"name"`
	Value     types.String                                                                                                       `tfsdk:"value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                     `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	Type   types.Int64                                                                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	Type   types.Int64                                                                                                                                   `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
}
type ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                                                                   `tfsdk:"metric_name"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                  `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                             `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
}
type ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                                                           `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                      `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                           `tfsdk:"match_labels"`
}
type PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel struct {
	Selector           SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                                                               `tfsdk:"metric_name"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                   `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	AverageUtilization types.Int64                                                                                                                                    `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                                                                   `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                          `tfsdk:"type"`
}
type ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                                                         `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
}
type MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecModel struct {
	Type     types.String                                                                                                                 `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
}
type HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                                                             `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                                                             `tfsdk:"max_replicas"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                                                         `tfsdk:"weight"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                                               `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                                                                         `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                           `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                                             `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                               `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                           `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                         `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                                   `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                               `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                             `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                                                                                 `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                                                                   `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	TopologyKey   types.String                                                                                                                                                                             `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                               `tfsdk:"namespaces"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
}
type DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecModel struct {
	ReplicaCount             types.Int64                                                                                                                    `tfsdk:"replica_count"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel                `tfsdk:"resources"`
	PodAnnotations           types.Map                                                                                                                      `tfsdk:"pod_annotations"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel                   `tfsdk:"env"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
}
type PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecModel struct {
	Name       types.String                                                                                                   `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                                                   `tfsdk:"api_version"`
	Kind       types.String                                                                                                   `tfsdk:"kind"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServicePortsModel struct {
	Type   types.Int64                                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
}
type PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServiceModel struct {
	Protocol   types.String                                                                                                       `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                                                       `tfsdk:"name"`
	NodePort   types.Int64                                                                                                        `tfsdk:"node_port"`
	Port       types.Int64                                                                                                        `tfsdk:"port"`
}
type ServiceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecModel struct {
	Ports       []*PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                                                                `tfsdk:"type"`
	Annotations types.Map                                                                                                   `tfsdk:"annotations"`
	Labels      types.Map                                                                                                   `tfsdk:"labels"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type KubeSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecModel struct {
	Service        ServiceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecModel    `tfsdk:"overlays"`
}
type CertManagerSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecModel `tfsdk:"kube_spec"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServicePortsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                         `tfsdk:"type"`
}
type PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServiceModel struct {
	TargetPort TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                                                                  `tfsdk:"name"`
	NodePort   types.Int64                                                                                                                   `tfsdk:"node_port"`
	Port       types.Int64                                                                                                                   `tfsdk:"port"`
	Protocol   types.String                                                                                                                  `tfsdk:"protocol"`
}
type ServiceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel struct {
	Type        types.String                                                                                                           `tfsdk:"type"`
	Annotations types.Map                                                                                                              `tfsdk:"annotations"`
	Labels      types.Map                                                                                                              `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServiceModel `tfsdk:"ports"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                                                              `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                                                                          `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                                        `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                                                            `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                              `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                          `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                        `tfsdk:"topology_key"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                                                                    `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                                          `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                                    `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                                                        `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                                                                          `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                    `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
}
type TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                                                                       `tfsdk:"optional"`
	Key                  types.String                                                                                                                                                     `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
}
type FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel struct {
	FieldPath  types.String `tfsdk:"field_path"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                                                                                         `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                                                                         `tfsdk:"resource"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                                                                  `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                                                    `tfsdk:"optional"`
}
type ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvModel struct {
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	Name      types.String                                                                                                                  `tfsdk:"name"`
	Value     types.String                                                                                                                  `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	RunAsNonRoot             types.Bool                                                                                                                                              `tfsdk:"run_as_non_root"`
	RunAsGroup               types.Int64                                                                                                                                             `tfsdk:"run_as_group"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	ProcMount                types.String                                                                                                                                            `tfsdk:"proc_mount"`
	RunAsUser                types.Int64                                                                                                                                             `tfsdk:"run_as_user"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	Privileged               types.Bool                                                                                                                                              `tfsdk:"privileged"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                                                              `tfsdk:"read_only_root_filesystem"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	AllowPrivilegeEscalation types.Bool                                                                                                                                              `tfsdk:"allow_privilege_escalation"`
}
type ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                                                                           `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                 `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchLabels      types.Map                                                                                                                                                      `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
}
type PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                                                                          `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                           `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                     `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Value              ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                                                               `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                                                                              `tfsdk:"type"`
}
type ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                                                                    `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchLabels      types.Map                                                                                                                                                                `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                              `tfsdk:"type"`
}
type ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                                                              `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                            `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                             `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                        `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel struct {
	Selector     SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                                                                      `tfsdk:"metric_name"`
}
type MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecModel struct {
	Pods     PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                                                                            `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
}
type HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	MinReplicas types.Int64                                                                                                                        `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                                                                        `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	FsGroup             types.Int64                                                                                                                                       `tfsdk:"fs_group"`
	RunAsNonRoot        types.Bool                                                                                                                                        `tfsdk:"run_as_non_root"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SupplementalGroups  types.List                                                                                                                                        `tfsdk:"supplemental_groups"`
	RunAsUser           types.Int64                                                                                                                                       `tfsdk:"run_as_user"`
	RunAsGroup          types.Int64                                                                                                                                       `tfsdk:"run_as_group"`
	FsGroupChangePolicy types.String                                                                                                                                      `tfsdk:"fs_group_change_policy"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
}
type DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel struct {
	Affinity                 AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	PodAnnotations           types.Map                                                                                                                                 `tfsdk:"pod_annotations"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel                   `tfsdk:"env"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel                `tfsdk:"resources"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	ReplicaCount             types.Int64                                                                                                                               `tfsdk:"replica_count"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel `tfsdk:"container_security_context"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel struct {
	RunAsGroup          types.Int64                                                                                                                                `tfsdk:"run_as_group"`
	SupplementalGroups  types.List                                                                                                                                 `tfsdk:"supplemental_groups"`
	FsGroup             types.Int64                                                                                                                                `tfsdk:"fs_group"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel     `tfsdk:"sysctls"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel `tfsdk:"seccomp_profile"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsNonRoot        types.Bool                                                                                                                                 `tfsdk:"run_as_non_root"`
	RunAsUser           types.Int64                                                                                                                                `tfsdk:"run_as_user"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel `tfsdk:"windows_options"`
	FsGroupChangePolicy types.String                                                                                                                               `tfsdk:"fs_group_change_policy"`
}
type TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel struct {
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                                                                             `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                                   `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                                                                             `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                               `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                                                 `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                   `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityModel struct {
	TopologyKey   types.String                                                                                                                                                                             `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                               `tfsdk:"namespaces"`
}
type PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                       `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityModel struct {
	Namespaces    types.List                                                                                                                                                                                   `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                 `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                                       `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                                                                                 `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                                   `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                                                                                     `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityModel `tfsdk:"pod_anti_affinity"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel struct {
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel struct {
	ReadOnlyRootFilesystem   types.Bool                                                                                                                                       `tfsdk:"read_only_root_filesystem"`
	Privileged               types.Bool                                                                                                                                       `tfsdk:"privileged"`
	RunAsUser                types.Int64                                                                                                                                      `tfsdk:"run_as_user"`
	AllowPrivilegeEscalation types.Bool                                                                                                                                       `tfsdk:"allow_privilege_escalation"`
	RunAsNonRoot             types.Bool                                                                                                                                       `tfsdk:"run_as_non_root"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	ProcMount                types.String                                                                                                                                     `tfsdk:"proc_mount"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel   `tfsdk:"capabilities"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel `tfsdk:"windows_options"`
	RunAsGroup               types.Int64                                                                                                                                      `tfsdk:"run_as_group"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel `tfsdk:"se_linux_options"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromResourceFieldRefModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel struct {
	ContainerName types.String                                                                                                                                  `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                                                                  `tfsdk:"resource"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel struct {
	Key                  types.String                                                                                                                                           `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                                             `tfsdk:"optional"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel struct {
	Key                  types.String                                                                                                                                              `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                                                `tfsdk:"optional"`
}
type FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel struct {
	FieldPath  types.String `tfsdk:"field_path"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvModel struct {
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel         `tfsdk:"field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel struct {
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                                                           `tfsdk:"name"`
	Value     types.String                                                                                                           `tfsdk:"value"`
}
type JobInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel struct {
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel `tfsdk:"container_security_context"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel                   `tfsdk:"env"`
	PodAnnotations           types.Map                                                                                                                          `tfsdk:"pod_annotations"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel       `tfsdk:"pod_security_context"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel           `tfsdk:"tolerations"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel                 `tfsdk:"affinity"`
}
type PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel struct {
	Kind       types.String                                                                                                              `tfsdk:"kind"`
	Name       types.String                                                                                                              `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                                                              `tfsdk:"api_version"`
}
type KubeSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckModel struct {
	Job            JobInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel            `tfsdk:"job"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel     `tfsdk:"deployment"`
}
type CertManagerStartupapicheckInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckModel `tfsdk:"kube_spec"`
}
type CertManagerInstallTemplateHelmSpecComponentsInternalCertProviderModel struct {
	CertManagerWebhookSpec     CertManagerWebhookSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel     `tfsdk:"cert_manager_webhook_spec"`
	Managed                    types.String                                                                                    `tfsdk:"managed"`
	CertManagerCaInjector      CertManagerCaInjectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel      `tfsdk:"cert_manager_ca_injector"`
	CertManagerSpec            CertManagerSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel            `tfsdk:"cert_manager_spec"`
	CertManagerStartupapicheck CertManagerStartupapicheckInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel `tfsdk:"cert_manager_startupapicheck"`
}
type InternalCertProviderInstallTemplateHelmSpecComponentsModel struct {
	Custom      CustomInstallTemplateHelmSpecComponentsInternalCertProviderModel      `tfsdk:"custom"`
	CertManager CertManagerInstallTemplateHelmSpecComponentsInternalCertProviderModel `tfsdk:"cert_manager"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServicePortsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                    `tfsdk:"type"`
}
type PortsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServiceModel struct {
	Name       types.String                                                                             `tfsdk:"name"`
	NodePort   types.Int64                                                                              `tfsdk:"node_port"`
	Port       types.Int64                                                                              `tfsdk:"port"`
	Protocol   types.String                                                                             `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServicePortsModel `tfsdk:"target_port"`
}
type ServiceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecModel struct {
	Labels      types.Map                                                                         `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                                      `tfsdk:"type"`
	Annotations types.Map                                                                         `tfsdk:"annotations"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type SysctlsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	RunAsNonRoot        types.Bool                                                                                                   `tfsdk:"run_as_non_root"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	SupplementalGroups  types.List                                                                                                   `tfsdk:"supplemental_groups"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsGroup          types.Int64                                                                                                  `tfsdk:"run_as_group"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	FsGroupChangePolicy types.String                                                                                                 `tfsdk:"fs_group_change_policy"`
	FsGroup             types.Int64                                                                                                  `tfsdk:"fs_group"`
	RunAsUser           types.Int64                                                                                                  `tfsdk:"run_as_user"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                     `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                                 `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                               `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                   `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                     `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                 `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                               `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                         `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                                                   `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                     `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                       `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                         `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                     `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                   `tfsdk:"topology_key"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type PreferenceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                                               `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel struct {
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	ProcMount                types.String                                                                                                       `tfsdk:"proc_mount"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	RunAsGroup               types.Int64                                                                                                        `tfsdk:"run_as_group"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	AllowPrivilegeEscalation types.Bool                                                                                                         `tfsdk:"allow_privilege_escalation"`
	Privileged               types.Bool                                                                                                         `tfsdk:"privileged"`
	RunAsUser                types.Int64                                                                                                        `tfsdk:"run_as_user"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                         `tfsdk:"read_only_root_filesystem"`
	RunAsNonRoot             types.Bool                                                                                                         `tfsdk:"run_as_non_root"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                 `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                            `tfsdk:"type"`
}
type PodsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                     `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
}
type TargetInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                                         `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                          `tfsdk:"average_utilization"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                      `tfsdk:"type"`
}
type ResourceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel struct {
	Target                   TargetInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                               `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                         `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchLabels      types.Map                                                                                                                           `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type ExternalInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                                         `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                        `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchLabels      types.Map                                                                                                                   `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
}
type TargetInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                       `tfsdk:"type"`
}
type ObjectInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName   types.String                                                                                                 `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
}
type MetricsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecModel struct {
	Pods     PodsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                                       `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
}
type HpaSpecInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	MaxReplicas types.Int64                                                                                   `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                                   `tfsdk:"min_replicas"`
}
type TolerationsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
}
type ResourcesInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                                `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                  `tfsdk:"optional"`
}
type FieldRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                           `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                                                    `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                                    `tfsdk:"resource"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                             `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                               `tfsdk:"optional"`
}
type ValueFromInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvModel struct {
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	Value     types.String                                                                             `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                             `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                           `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	Type          types.String                                                                                      `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type DeploymentInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecModel struct {
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	PodAnnotations           types.Map                                                                                            `tfsdk:"pod_annotations"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel                `tfsdk:"resources"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel                   `tfsdk:"env"`
	ReplicaCount             types.Int64                                                                                          `tfsdk:"replica_count"`
}
type PatchesInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecModel struct {
	ApiVersion types.String                                                                         `tfsdk:"api_version"`
	Kind       types.String                                                                         `tfsdk:"kind"`
	Name       types.String                                                                         `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecOverlaysModel `tfsdk:"patches"`
}
type KubeSpecInstallTemplateHelmSpecComponentsOnboardingRepositoryModel struct {
	Service        ServiceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecModel    `tfsdk:"overlays"`
}
type RepositoryInstallTemplateHelmSpecComponentsOnboardingModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsOnboardingRepositoryModel `tfsdk:"kube_spec"`
}
type ResourcesInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                         `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                   `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyModel struct {
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
}
type StrategyInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                                    `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                   `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                               `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                             `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityModel struct {
	Weight          types.Int64                                                                                                                                                 `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                                   `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityModel struct {
	Namespaces    types.List                                                                                                                                               `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                             `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                       `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                                                 `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                   `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                                                     `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                                       `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                   `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                 `tfsdk:"topology_key"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                             `tfsdk:"weight"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	AllowPrivilegeEscalation types.Bool                                                                                                       `tfsdk:"allow_privilege_escalation"`
	Privileged               types.Bool                                                                                                       `tfsdk:"privileged"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsUser                types.Int64                                                                                                      `tfsdk:"run_as_user"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	ProcMount                types.String                                                                                                     `tfsdk:"proc_mount"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                       `tfsdk:"read_only_root_filesystem"`
	RunAsNonRoot             types.Bool                                                                                                       `tfsdk:"run_as_non_root"`
	RunAsGroup               types.Int64                                                                                                      `tfsdk:"run_as_group"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel struct {
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                             `tfsdk:"optional"`
	Key                  types.String                                                                                                           `tfsdk:"key"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                              `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                `tfsdk:"optional"`
}
type FieldRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                         `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                                                  `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                                  `tfsdk:"resource"`
}
type ValueFromInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvModel struct {
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                           `tfsdk:"name"`
	Value     types.String                                                                           `tfsdk:"value"`
}
type TargetInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Type   types.Int64                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                      `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                 `tfsdk:"match_labels"`
}
type ObjectInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel struct {
	AverageValue AverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                               `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchLabels      types.Map                                                                                                               `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
}
type PodsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                   `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	Type   types.Int64                                                                                                       `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                              `tfsdk:"type"`
}
type TargetInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	AverageUtilization types.Int64                                                                                                        `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                                       `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                                                    `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type ResourceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel struct {
	Name                     types.String                                                                                                             `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                         `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	Type   types.Int64                                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                       `tfsdk:"type"`
}
type ExternalInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                       `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
}
type MetricsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecModel struct {
	Type     types.String                                                                                     `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
}
type HpaSpecInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	MaxReplicas types.Int64                                                                                 `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                                 `tfsdk:"min_replicas"`
}
type TolerationsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type SysctlsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	FsGroup             types.Int64                                                                                                `tfsdk:"fs_group"`
	RunAsNonRoot        types.Bool                                                                                                 `tfsdk:"run_as_non_root"`
	FsGroupChangePolicy types.String                                                                                               `tfsdk:"fs_group_change_policy"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	SupplementalGroups  types.List                                                                                                 `tfsdk:"supplemental_groups"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsUser           types.Int64                                                                                                `tfsdk:"run_as_user"`
	RunAsGroup          types.Int64                                                                                                `tfsdk:"run_as_group"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
}
type DeploymentInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel struct {
	ReplicaCount             types.Int64                                                                                        `tfsdk:"replica_count"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel                   `tfsdk:"env"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel                `tfsdk:"resources"`
	PodAnnotations           types.Map                                                                                          `tfsdk:"pod_annotations"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel           `tfsdk:"tolerations"`
}
type PatchesInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel struct {
	Name       types.String                                                                       `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                       `tfsdk:"api_version"`
	Kind       types.String                                                                       `tfsdk:"kind"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServicePortsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                  `tfsdk:"type"`
}
type PortsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServiceModel struct {
	Name       types.String                                                                           `tfsdk:"name"`
	NodePort   types.Int64                                                                            `tfsdk:"node_port"`
	Port       types.Int64                                                                            `tfsdk:"port"`
	Protocol   types.String                                                                           `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServicePortsModel `tfsdk:"target_port"`
}
type ServiceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel struct {
	Labels      types.Map                                                                       `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                                    `tfsdk:"type"`
	Annotations types.Map                                                                       `tfsdk:"annotations"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type KubeSpecInstallTemplateHelmSpecComponentsOnboardingOperatorModel struct {
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel        `tfsdk:"service"`
}
type OperatorInstallTemplateHelmSpecComponentsOnboardingModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsOnboardingOperatorModel `tfsdk:"kube_spec"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type TolerationsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel struct {
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	ReadOnlyRootFilesystem   types.Bool                                                                                                            `tfsdk:"read_only_root_filesystem"`
	Privileged               types.Bool                                                                                                            `tfsdk:"privileged"`
	AllowPrivilegeEscalation types.Bool                                                                                                            `tfsdk:"allow_privilege_escalation"`
	RunAsNonRoot             types.Bool                                                                                                            `tfsdk:"run_as_non_root"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	ProcMount                types.String                                                                                                          `tfsdk:"proc_mount"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsUser                types.Int64                                                                                                           `tfsdk:"run_as_user"`
	RunAsGroup               types.Int64                                                                                                           `tfsdk:"run_as_group"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                              `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel struct {
	Divisor       DivisorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                                       `tfsdk:"resource"`
	ContainerName types.String                                                                                                       `tfsdk:"container_name"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                                  `tfsdk:"optional"`
	Key                  types.String                                                                                                                `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                                   `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                     `tfsdk:"optional"`
}
type FieldRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type ValueFromInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvModel struct {
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	Value     types.String                                                                                `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                                `tfsdk:"name"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                    `tfsdk:"match_labels"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	Type   types.Int64                                                                                                               `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
}
type PodsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                        `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                   `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                   `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                            `tfsdk:"type"`
}
type TargetInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                                            `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                             `tfsdk:"average_utilization"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
}
type ResourceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                                  `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                            `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                              `tfsdk:"match_labels"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                   `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type ExternalInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                                            `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                      `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                          `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                           `tfsdk:"type"`
}
type ObjectInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName   types.String                                                                                                    `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
}
type MetricsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecModel struct {
	Pods     PodsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                                          `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
}
type HpaSpecInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                                      `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                                      `tfsdk:"max_replicas"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                        `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                              `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	Type          types.String                                                                                         `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type SysctlsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	RunAsUser           types.Int64                                                                                                     `tfsdk:"run_as_user"`
	SupplementalGroups  types.List                                                                                                      `tfsdk:"supplemental_groups"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	RunAsNonRoot        types.Bool                                                                                                      `tfsdk:"run_as_non_root"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	FsGroup             types.Int64                                                                                                     `tfsdk:"fs_group"`
	FsGroupChangePolicy types.String                                                                                                    `tfsdk:"fs_group_change_policy"`
	RunAsGroup          types.Int64                                                                                                     `tfsdk:"run_as_group"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
}
type ResourcesInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	Requests types.Map `tfsdk:"requests"`
	Limits   types.Map `tfsdk:"limits"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type PreferenceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                                  `tfsdk:"weight"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                        `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                                                  `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                    `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                      `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                        `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityModel struct {
	Namespaces    types.List                                                                                                                                                    `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                  `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                            `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                        `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                      `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                          `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                            `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Namespaces    types.List                                                                                                                                                        `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                      `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
}
type DeploymentInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel struct {
	ReplicaCount             types.Int64                                                                                             `tfsdk:"replica_count"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel                `tfsdk:"resources"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	PodAnnotations           types.Map                                                                                               `tfsdk:"pod_annotations"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel                   `tfsdk:"env"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel                 `tfsdk:"strategy"`
}
type PatchesInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel struct {
	Patches    []*PatchesInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                            `tfsdk:"api_version"`
	Kind       types.String                                                                            `tfsdk:"kind"`
	Name       types.String                                                                            `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServicePortsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                       `tfsdk:"type"`
}
type PortsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServiceModel struct {
	TargetPort TargetPortInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                                `tfsdk:"name"`
	NodePort   types.Int64                                                                                 `tfsdk:"node_port"`
	Port       types.Int64                                                                                 `tfsdk:"port"`
	Protocol   types.String                                                                                `tfsdk:"protocol"`
}
type ServiceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel struct {
	Annotations types.Map                                                                            `tfsdk:"annotations"`
	Labels      types.Map                                                                            `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                                         `tfsdk:"type"`
}
type KubeSpecInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceModel struct {
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel        `tfsdk:"service"`
}
type InstanceInstallTemplateHelmSpecComponentsOnboardingPlaneModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceModel `tfsdk:"kube_spec"`
}
type PlaneInstallTemplateHelmSpecComponentsOnboardingModel struct {
	Instance InstanceInstallTemplateHelmSpecComponentsOnboardingPlaneModel `tfsdk:"instance"`
}
type OnboardingInstallTemplateHelmSpecComponentsModel struct {
	Repository RepositoryInstallTemplateHelmSpecComponentsOnboardingModel `tfsdk:"repository"`
	Operator   OperatorInstallTemplateHelmSpecComponentsOnboardingModel   `tfsdk:"operator"`
	Plane      PlaneInstallTemplateHelmSpecComponentsOnboardingModel      `tfsdk:"plane"`
}
type ConfigProtectionInstallTemplateHelmSpecComponentsXcpModel struct {
	EnableAuthorizedCreateUpdateDeleteOnXcpConfigs types.Bool `tfsdk:"enable_authorized_create_update_delete_on_xcp_configs"`
	EnableAuthorizedUpdateDeleteOnXcpConfigs       types.Bool `tfsdk:"enable_authorized_update_delete_on_xcp_configs"`
	AuthorizedUsers                                types.List `tfsdk:"authorized_users"`
}
type MaxWorkloadCert_TTLInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel struct {
	Seconds types.Int64 `tfsdk:"seconds"`
	Nanos   types.Int64 `tfsdk:"nanos"`
}
type PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioPilotOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type PilotOverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel struct {
	Kind       types.String                                                                                      `tfsdk:"kind"`
	Name       types.String                                                                                      `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioPilotOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                                      `tfsdk:"api_version"`
}
type PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioBaseOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type BaseOverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel struct {
	Patches    []*PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioBaseOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                                     `tfsdk:"api_version"`
	Kind       types.String                                                                                     `tfsdk:"kind"`
	Name       types.String                                                                                     `tfsdk:"name"`
}
type DefaultWorkloadCert_TTLInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel struct {
	Seconds types.Int64 `tfsdk:"seconds"`
	Nanos   types.Int64 `tfsdk:"nanos"`
}
type PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstio_CNIOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type CNIOverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel struct {
	Name       types.String                                                                                     `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstio_CNIOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                                     `tfsdk:"api_version"`
	Kind       types.String                                                                                     `tfsdk:"kind"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServicePortsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                    `tfsdk:"type"`
}
type PortsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServiceModel struct {
	Name       types.String                                                                                             `tfsdk:"name"`
	NodePort   types.Int64                                                                                              `tfsdk:"node_port"`
	Port       types.Int64                                                                                              `tfsdk:"port"`
	Protocol   types.String                                                                                             `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServicePortsModel `tfsdk:"target_port"`
}
type ServiceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel struct {
	Type        types.String                                                                                      `tfsdk:"type"`
	Annotations types.Map                                                                                         `tfsdk:"annotations"`
	Labels      types.Map                                                                                         `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServiceModel `tfsdk:"ports"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type CNIInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel struct {
	ClusterRole            types.String `tfsdk:"cluster_role"`
	ConfigurationDirectory types.String `tfsdk:"configuration_directory"`
	ConfigurationFileName  types.String `tfsdk:"configuration_file_name"`
	Revision               types.String `tfsdk:"revision"`
	BinaryDirectory        types.String `tfsdk:"binary_directory"`
	Chained                types.Bool   `tfsdk:"chained"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                         `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	TopologyKey   types.String                                                                                                                                                                   `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                     `tfsdk:"namespaces"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                         `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                     `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                   `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                                       `tfsdk:"weight"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                                                               `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                                                     `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                 `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                               `tfsdk:"topology_key"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                                     `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                                                               `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                 `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                                   `tfsdk:"weight"`
}
type PodAffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromModel struct {
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                                  `tfsdk:"optional"`
	Key                  types.String                                                                                                                                `tfsdk:"key"`
}
type FieldRefInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                           `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromModel struct {
	Resource      types.String                                                                                                                    `tfsdk:"resource"`
	ContainerName types.String                                                                                                                    `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                                               `tfsdk:"optional"`
	Key                  types.String                                                                                                                             `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
}
type ValueFromInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvModel struct {
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                                             `tfsdk:"name"`
	Value     types.String                                                                                             `tfsdk:"value"`
}
type ResourcesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type SysctlsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	FsGroup             types.Int64                                                                                                                  `tfsdk:"fs_group"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsUser           types.Int64                                                                                                                  `tfsdk:"run_as_user"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SupplementalGroups  types.List                                                                                                                   `tfsdk:"supplemental_groups"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	FsGroupChangePolicy types.String                                                                                                                 `tfsdk:"fs_group_change_policy"`
	RunAsNonRoot        types.Bool                                                                                                                   `tfsdk:"run_as_non_root"`
	RunAsGroup          types.Int64                                                                                                                  `tfsdk:"run_as_group"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchLabels      types.Map                                                                                                                                   `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
}
type TargetInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                       `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                        `tfsdk:"type"`
}
type ObjectInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName   types.String                                                                                                                 `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                 `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                            `tfsdk:"type"`
}
type PodsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                                     `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                `tfsdk:"type"`
}
type TargetInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type               types.String                                                                                                                         `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                                          `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
}
type ResourceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                                               `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                           `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
}
type ExternalInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                                         `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
}
type MetricsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecModel struct {
	Object   ObjectInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                                                       `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
}
type HpaSpecInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                                                   `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                                                   `tfsdk:"max_replicas"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                           `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                     `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyModel struct {
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
}
type StrategyInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                                                      `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	ReadOnlyRootFilesystem   types.Bool                                                                                                                         `tfsdk:"read_only_root_filesystem"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	ProcMount                types.String                                                                                                                       `tfsdk:"proc_mount"`
	RunAsGroup               types.Int64                                                                                                                        `tfsdk:"run_as_group"`
	AllowPrivilegeEscalation types.Bool                                                                                                                         `tfsdk:"allow_privilege_escalation"`
	RunAsNonRoot             types.Bool                                                                                                                         `tfsdk:"run_as_non_root"`
	Privileged               types.Bool                                                                                                                         `tfsdk:"privileged"`
	RunAsUser                types.Int64                                                                                                                        `tfsdk:"run_as_user"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
}
type TolerationsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
}
type DeploymentInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel struct {
	Env                      []*EnvInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel                   `tfsdk:"env"`
	PodAnnotations           types.Map                                                                                                            `tfsdk:"pod_annotations"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	ReplicaCount             types.Int64                                                                                                          `tfsdk:"replica_count"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel                `tfsdk:"resources"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel                 `tfsdk:"strategy"`
}
type PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel struct {
	Kind       types.String                                                                                         `tfsdk:"kind"`
	Name       types.String                                                                                         `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                                         `tfsdk:"api_version"`
}
type KubeSpecInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel struct {
	Deployment     DeploymentInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel `tfsdk:"service_account"`
	CNI            CNIInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel            `tfsdk:"cni"`
}
type IstioInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsModel struct {
	TsbVersion              types.String                                                                                      `tfsdk:"tsb_version"`
	PilotOverlays           []*PilotOverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel        `tfsdk:"pilot_overlays"`
	DefaultWorkloadCert_TTL DefaultWorkloadCert_TTLInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel `tfsdk:"default_workload_cert_ttl"`
	KubeSpec                KubeSpecInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel                `tfsdk:"kube_spec"`
	TrustDomain             types.String                                                                                      `tfsdk:"trust_domain"`
	TraceSamplingRate       types.Float64                                                                                     `tfsdk:"trace_sampling_rate"`
	MaxWorkloadCert_TTL     MaxWorkloadCert_TTLInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel     `tfsdk:"max_workload_cert_ttl"`
	BaseOverlays            []*BaseOverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel         `tfsdk:"base_overlays"`
	CniOverlays             []*CNIOverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel          `tfsdk:"cni_overlays"`
}
type RevisionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesModel struct {
	Name    types.String                                                               `tfsdk:"name"`
	Disable types.Bool                                                                 `tfsdk:"disable"`
	Istio   IstioInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsModel `tfsdk:"istio"`
}
type IsolationBoundariesInstallTemplateHelmSpecComponentsXcpModel struct {
	Revisions []*RevisionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesModel `tfsdk:"revisions"`
	Name      types.String                                                             `tfsdk:"name"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                              `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                    `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityModel struct {
	Namespaces    types.List                                                                                                                                `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                              `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                    `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                              `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                  `tfsdk:"weight"`
}
type PodAffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                        `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                    `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                  `tfsdk:"topology_key"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                        `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                    `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                  `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                                      `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                                   `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                   `tfsdk:"resource"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                            `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                              `tfsdk:"optional"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                               `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                 `tfsdk:"optional"`
}
type FieldRefInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromModel struct {
	FieldPath  types.String `tfsdk:"field_path"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type ValueFromInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvModel struct {
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	Value     types.String                                                            `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                            `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                               `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	Type   types.Int64                                                                                        `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchLabels      types.Map                                                                                                          `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
}
type ExternalInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                        `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Type   types.Int64                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                       `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                  `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type ObjectInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchLabels      types.Map                                                                                                `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                           `tfsdk:"type"`
}
type PodsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel struct {
	Selector           SelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                    `tfsdk:"metric_name"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                               `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                        `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	Type   types.Int64                                                                                               `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type TargetInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type               types.String                                                                                        `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                         `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
}
type ResourceInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                              `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
}
type MetricsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecModel struct {
	Resource ResourceInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                      `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
}
type HpaSpecInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                  `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                  `tfsdk:"max_replicas"`
}
type TolerationsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
}
type ResourcesInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SysctlsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	RunAsUser           types.Int64                                                                                 `tfsdk:"run_as_user"`
	FsGroupChangePolicy types.String                                                                                `tfsdk:"fs_group_change_policy"`
	RunAsNonRoot        types.Bool                                                                                  `tfsdk:"run_as_non_root"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	RunAsGroup          types.Int64                                                                                 `tfsdk:"run_as_group"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	FsGroup             types.Int64                                                                                 `tfsdk:"fs_group"`
	SupplementalGroups  types.List                                                                                  `tfsdk:"supplemental_groups"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	RunAsGroup               types.Int64                                                                                       `tfsdk:"run_as_group"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	ProcMount                types.String                                                                                      `tfsdk:"proc_mount"`
	AllowPrivilegeEscalation types.Bool                                                                                        `tfsdk:"allow_privilege_escalation"`
	ReadOnlyRootFilesystem   types.Bool                                                                                        `tfsdk:"read_only_root_filesystem"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	RunAsNonRoot             types.Bool                                                                                        `tfsdk:"run_as_non_root"`
	RunAsUser                types.Int64                                                                                       `tfsdk:"run_as_user"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	Privileged               types.Bool                                                                                        `tfsdk:"privileged"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                    `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	Type          types.String                                                                     `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type DeploymentInstallTemplateHelmSpecComponentsXcpKubeSpecModel struct {
	Strategy                 StrategyInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel                `tfsdk:"resources"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	ReplicaCount             types.Int64                                                                         `tfsdk:"replica_count"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel                   `tfsdk:"env"`
	PodAnnotations           types.Map                                                                           `tfsdk:"pod_annotations"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel           `tfsdk:"tolerations"`
}
type PatchesInstallTemplateHelmSpecComponentsXcpKubeSpecOverlaysModel struct {
	Path types.String `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsXcpKubeSpecModel struct {
	Kind       types.String                                                        `tfsdk:"kind"`
	Name       types.String                                                        `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsXcpKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                        `tfsdk:"api_version"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsXcpKubeSpecServicePortsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                   `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
}
type PortsInstallTemplateHelmSpecComponentsXcpKubeSpecServiceModel struct {
	NodePort   types.Int64                                                             `tfsdk:"node_port"`
	Port       types.Int64                                                             `tfsdk:"port"`
	Protocol   types.String                                                            `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsXcpKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                            `tfsdk:"name"`
}
type ServiceInstallTemplateHelmSpecComponentsXcpKubeSpecModel struct {
	Ports       []*PortsInstallTemplateHelmSpecComponentsXcpKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                     `tfsdk:"type"`
	Annotations types.Map                                                        `tfsdk:"annotations"`
	Labels      types.Map                                                        `tfsdk:"labels"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsXcpKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsXcpKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsXcpKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type KubeSpecInstallTemplateHelmSpecComponentsXcpModel struct {
	Deployment     DeploymentInstallTemplateHelmSpecComponentsXcpKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsXcpKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsXcpKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsXcpKubeSpecModel `tfsdk:"service_account"`
}
type XcpInstallTemplateHelmSpecComponentsModel struct {
	ConfigProtection                          ConfigProtectionInstallTemplateHelmSpecComponentsXcpModel       `tfsdk:"config_protection"`
	EnableHttpMeshInternalIdentityPropagation types.Bool                                                      `tfsdk:"enable_http_mesh_internal_identity_propagation"`
	IsolationBoundaries                       []*IsolationBoundariesInstallTemplateHelmSpecComponentsXcpModel `tfsdk:"isolation_boundaries"`
	KubeSpec                                  KubeSpecInstallTemplateHelmSpecComponentsXcpModel               `tfsdk:"kube_spec"`
	Revision                                  types.String                                                    `tfsdk:"revision"`
	CentralAuthMode                           types.String                                                    `tfsdk:"central_auth_mode"`
	CentralProvidedCaCert                     types.Bool                                                      `tfsdk:"central_provided_ca_cert"`
}
type ComponentsInstallTemplateHelmSpecModel struct {
	HpaAdapter           HpaAdapterInstallTemplateHelmSpecComponentsModel           `tfsdk:"hpa_adapter"`
	RateLimitServer      RateLimitServerInstallTemplateHelmSpecComponentsModel      `tfsdk:"rate_limit_server"`
	Onboarding           OnboardingInstallTemplateHelmSpecComponentsModel           `tfsdk:"onboarding"`
	Istio                IstioInstallTemplateHelmSpecComponentsModel                `tfsdk:"istio"`
	Oap                  OapInstallTemplateHelmSpecComponentsModel                  `tfsdk:"oap"`
	Satellite            SatelliteInstallTemplateHelmSpecComponentsModel            `tfsdk:"satellite"`
	Ngac                 NgacInstallTemplateHelmSpecComponentsModel                 `tfsdk:"ngac"`
	DefaultKubeSpec      DefaultKubeSpecInstallTemplateHelmSpecComponentsModel      `tfsdk:"default_kube_spec"`
	InternalCertProvider InternalCertProviderInstallTemplateHelmSpecComponentsModel `tfsdk:"internal_cert_provider"`
	Xcp                  XcpInstallTemplateHelmSpecComponentsModel                  `tfsdk:"xcp"`
	Collector            CollectorInstallTemplateHelmSpecComponentsModel            `tfsdk:"collector"`
	Gitops               GitopsInstallTemplateHelmSpecComponentsModel               `tfsdk:"gitops"`
}
type ImagePullSecretsInstallTemplateHelmSpecModel struct {
	Name types.String `tfsdk:"name"`
}
type SpecInstallTemplateHelmModel struct {
	MeshObservability MeshObservabilityInstallTemplateHelmSpecModel   `tfsdk:"mesh_observability"`
	TelemetryStore    TelemetryStoreInstallTemplateHelmSpecModel      `tfsdk:"telemetry_store"`
	Tier1Cluster      types.Bool                                      `tfsdk:"tier1_cluster"`
	Components        ComponentsInstallTemplateHelmSpecModel          `tfsdk:"components"`
	Hub               types.String                                    `tfsdk:"hub"`
	ImagePullSecrets  []*ImagePullSecretsInstallTemplateHelmSpecModel `tfsdk:"image_pull_secrets"`
	ManagementPlane   ManagementPlaneInstallTemplateHelmSpecModel     `tfsdk:"management_plane"`
	MeshExpansion     MeshExpansionInstallTemplateHelmSpecModel       `tfsdk:"mesh_expansion"`
}
type ImageInstallTemplateHelmModel struct {
	Tag      types.String `tfsdk:"tag"`
	Registry types.String `tfsdk:"registry"`
}
type ServiceAccountInstallTemplateHelmOperatorModel struct {
	Annotations      types.Map    `tfsdk:"annotations"`
	ImagePullSecrets types.List   `tfsdk:"image_pull_secrets"`
	PullPassword     types.String `tfsdk:"pull_password"`
	PullSecret       types.String `tfsdk:"pull_secret"`
	PullUsername     types.String `tfsdk:"pull_username"`
}
type FieldRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type IntValInstallTemplateHelmOperatorDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmOperatorDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmOperatorDeploymentEnvValueFromResourceFieldRefModel struct {
	Type   types.Int64                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmOperatorDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmOperatorDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
}
type ResourceFieldRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel struct {
	ContainerName types.String                                                                  `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmOperatorDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                  `tfsdk:"resource"`
}
type LocalObjectReferenceInstallTemplateHelmOperatorDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel struct {
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmOperatorDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                             `tfsdk:"optional"`
	Key                  types.String                                                                           `tfsdk:"key"`
}
type LocalObjectReferenceInstallTemplateHelmOperatorDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                              `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmOperatorDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                `tfsdk:"optional"`
}
type ValueFromInstallTemplateHelmOperatorDeploymentEnvModel struct {
	SecretKeyRef     SecretKeyRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
}
type EnvInstallTemplateHelmOperatorDeploymentModel struct {
	Name      types.String                                           `tfsdk:"name"`
	Value     types.String                                           `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmOperatorDeploymentEnvModel `tfsdk:"value_from"`
}
type IntValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                   `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
}
type RollingUpdateInstallTemplateHelmOperatorDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmOperatorDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmOperatorDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                    `tfsdk:"type"`
}
type TolerationsInstallTemplateHelmOperatorDeploymentModel struct {
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
}
type MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                       `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                 `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                   `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                     `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                       `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityModel struct {
	TopologyKey   types.String                                                                                                                 `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                   `tfsdk:"namespaces"`
}
type PodAntiAffinityInstallTemplateHelmOperatorDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type NodeSelectorTermsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type PreferenceInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                             `tfsdk:"weight"`
}
type NodeAffinityInstallTemplateHelmOperatorDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmOperatorDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                   `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmOperatorDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                               `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                             `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmOperatorDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmOperatorDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                 `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmOperatorDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                   `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAffinityModel struct {
	TopologyKey   types.String                                                                                                             `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmOperatorDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                               `tfsdk:"namespaces"`
}
type PodAffinityInstallTemplateHelmOperatorDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmOperatorDeploymentModel struct {
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmOperatorDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmOperatorDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmOperatorDeploymentAffinityModel     `tfsdk:"pod_affinity"`
}
type DeploymentInstallTemplateHelmOperatorModel struct {
	Annotations    types.Map                                                `tfsdk:"annotations"`
	Env            []*EnvInstallTemplateHelmOperatorDeploymentModel         `tfsdk:"env"`
	PodAnnotations types.Map                                                `tfsdk:"pod_annotations"`
	ReplicaCount   types.Int64                                              `tfsdk:"replica_count"`
	Strategy       StrategyInstallTemplateHelmOperatorDeploymentModel       `tfsdk:"strategy"`
	Tolerations    []*TolerationsInstallTemplateHelmOperatorDeploymentModel `tfsdk:"tolerations"`
	Affinity       AffinityInstallTemplateHelmOperatorDeploymentModel       `tfsdk:"affinity"`
}
type ServiceInstallTemplateHelmOperatorModel struct {
	Annotations types.Map `tfsdk:"annotations"`
}
type OperatorInstallTemplateHelmModel struct {
	Deployment     DeploymentInstallTemplateHelmOperatorModel     `tfsdk:"deployment"`
	Service        ServiceInstallTemplateHelmOperatorModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmOperatorModel `tfsdk:"service_account"`
}
type ClusterServiceAccountInstallTemplateHelmSecretsModel struct {
	JWK         types.String `tfsdk:"jwk"`
	Cluster_FQN types.String `tfsdk:"cluster_fqn"`
	Encoded_JWK types.String `tfsdk:"encoded_jwk"`
}
type ElasticsearchInstallTemplateHelmSecretsModel struct {
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
	Cacert   types.String `tfsdk:"cacert"`
}
type TsbInstallTemplateHelmSecretsModel struct {
	Cacert types.String `tfsdk:"cacert"`
}
type EdgeInstallTemplateHelmSecretsXcpModel struct {
	Token types.String `tfsdk:"token"`
	Cert  types.String `tfsdk:"cert"`
	Key   types.String `tfsdk:"key"`
}
type XcpInstallTemplateHelmSecretsModel struct {
	Edge              EdgeInstallTemplateHelmSecretsXcpModel `tfsdk:"edge"`
	Rootca            types.String                           `tfsdk:"rootca"`
	Rootcakey         types.String                           `tfsdk:"rootcakey"`
	AutoGenerateCerts types.Bool                             `tfsdk:"auto_generate_certs"`
}
type SecretsInstallTemplateHelmModel struct {
	ClusterServiceAccount ClusterServiceAccountInstallTemplateHelmSecretsModel `tfsdk:"cluster_service_account"`
	Elasticsearch         ElasticsearchInstallTemplateHelmSecretsModel         `tfsdk:"elasticsearch"`
	Tsb                   TsbInstallTemplateHelmSecretsModel                   `tfsdk:"tsb"`
	Xcp                   XcpInstallTemplateHelmSecretsModel                   `tfsdk:"xcp"`
}
type HelmInstallTemplateModel struct {
	Spec     SpecInstallTemplateHelmModel     `tfsdk:"spec"`
	Image    ImageInstallTemplateHelmModel    `tfsdk:"image"`
	Operator OperatorInstallTemplateHelmModel `tfsdk:"operator"`
	Secrets  SecretsInstallTemplateHelmModel  `tfsdk:"secrets"`
}
type InstallTemplateModel struct {
	Message types.String             `tfsdk:"message"`
	Helm    HelmInstallTemplateModel `tfsdk:"helm"`
}
type ProxyNamespacesServicesWorkloadsModel struct {
	EnvoyVersion        types.String `tfsdk:"envoy_version"`
	IstioVersion        types.String `tfsdk:"istio_version"`
	Status              types.Map    `tfsdk:"status"`
	ControlPlaneAddress types.String `tfsdk:"control_plane_address"`
}
type WorkloadsNamespacesServicesModel struct {
	Name    types.String                          `tfsdk:"name"`
	Proxy   ProxyNamespacesServicesWorkloadsModel `tfsdk:"proxy"`
	Address types.String                          `tfsdk:"address"`
	IsVm    types.Bool                            `tfsdk:"is_vm"`
}
type PortsNamespacesServicesModel struct {
	Name               types.String `tfsdk:"name"`
	Number             types.Int64  `tfsdk:"number"`
	KubernetesNodePort types.Int64  `tfsdk:"kubernetes_node_port"`
}
type ServicesNamespacesModel struct {
	KubernetesExternalAddresses types.List                          `tfsdk:"kubernetes_external_addresses"`
	Ports                       []*PortsNamespacesServicesModel     `tfsdk:"ports"`
	NumVmEndpoints              types.Int64                         `tfsdk:"num_vm_endpoints"`
	SpiffeIds                   types.List                          `tfsdk:"spiffe_ids"`
	CanonicalName               types.String                        `tfsdk:"canonical_name"`
	Workloads                   []*WorkloadsNamespacesServicesModel `tfsdk:"workloads"`
	Subsets                     types.List                          `tfsdk:"subsets"`
	Name                        types.String                        `tfsdk:"name"`
	GatewayHost                 types.Bool                          `tfsdk:"gateway_host"`
	NumKubernetesEndpoints      types.Int64                         `tfsdk:"num_kubernetes_endpoints"`
	Hostname                    types.String                        `tfsdk:"hostname"`
	Selector                    types.Map                           `tfsdk:"selector"`
	Namespace                   types.String                        `tfsdk:"namespace"`
	KubernetesServiceFqdn       types.String                        `tfsdk:"kubernetes_service_fqdn"`
	NumHops                     types.Int64                         `tfsdk:"num_hops"`
	MeshExternal                types.Bool                          `tfsdk:"mesh_external"`
	Tier1GatewayHost            types.Bool                          `tfsdk:"tier1_gateway_host"`
	KubernetesServiceIp         types.String                        `tfsdk:"kubernetes_service_ip"`
}
type NamespacesModel struct {
	Name     types.String               `tfsdk:"name"`
	Services []*ServicesNamespacesModel `tfsdk:"services"`
}
type KeysServiceAccountModel struct {
	DefaultToken types.String `tfsdk:"default_token"`
	Encoding     types.String `tfsdk:"encoding"`
	Id           types.String `tfsdk:"id"`
	PrivateKey   types.String `tfsdk:"private_key"`
	PublicKey    types.String `tfsdk:"public_key"`
}
type ServiceAccountModel struct {
	DisplayName types.String               `tfsdk:"display_name"`
	Keys        []*KeysServiceAccountModel `tfsdk:"keys"`
	Description types.String               `tfsdk:"description"`
}
type ClusterModel struct {
	Network         types.String         `tfsdk:"network"`
	Parent          types.String         `tfsdk:"parent"`
	Name            types.String         `tfsdk:"name"`
	Id              types.String         `tfsdk:"id"`
	InstallTemplate InstallTemplateModel `tfsdk:"install_template"`
	NamespaceScope  NamespaceScopeModel  `tfsdk:"namespace_scope"`
	Description     types.String         `tfsdk:"description"`
	Namespaces      []*NamespacesModel   `tfsdk:"namespaces"`
	TrustDomain     types.String         `tfsdk:"trust_domain"`
	Labels          types.Map            `tfsdk:"labels"`
	Locality        LocalityModel        `tfsdk:"locality"`
	State           StateModel           `tfsdk:"state"`
	Tier1Cluster    Tier1ClusterModel    `tfsdk:"tier1_cluster"`
	TokenTtl        Token_TTLModel       `tfsdk:"token_ttl"`
	ServiceAccount  ServiceAccountModel  `tfsdk:"service_account"`
	DisplayName     types.String         `tfsdk:"display_name"`
}
