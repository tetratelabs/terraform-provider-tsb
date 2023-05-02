package cluster

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type Tier1ClusterModel struct {
	Value types.Bool `tfsdk:"value"`
}
type LocalityModel struct {
	Region types.String `tfsdk:"region"`
}
type KeysServiceAccountModel struct {
	PrivateKey   types.String `tfsdk:"private_key"`
	PublicKey    types.String `tfsdk:"public_key"`
	DefaultToken types.String `tfsdk:"default_token"`
	Encoding     types.String `tfsdk:"encoding"`
	Id           types.String `tfsdk:"id"`
}
type ServiceAccountModel struct {
	DisplayName types.String               `tfsdk:"display_name"`
	Keys        []*KeysServiceAccountModel `tfsdk:"keys"`
	Description types.String               `tfsdk:"description"`
}
type TokenTtlModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type NamespaceScopeModel struct {
	Exceptions types.List   `tfsdk:"exceptions"`
	Scope      types.String `tfsdk:"scope"`
}
type PortsNamespacesServicesModel struct {
	Name               types.String `tfsdk:"name"`
	Number             types.Int64  `tfsdk:"number"`
	KubernetesNodePort types.Int64  `tfsdk:"kubernetes_node_port"`
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
type ServicesNamespacesModel struct {
	Name                        types.String                        `tfsdk:"name"`
	CanonicalName               types.String                        `tfsdk:"canonical_name"`
	KubernetesExternalAddresses types.List                          `tfsdk:"kubernetes_external_addresses"`
	Ports                       []*PortsNamespacesServicesModel     `tfsdk:"ports"`
	Tier1GatewayHost            types.Bool                          `tfsdk:"tier1_gateway_host"`
	Workloads                   []*WorkloadsNamespacesServicesModel `tfsdk:"workloads"`
	GatewayHost                 types.Bool                          `tfsdk:"gateway_host"`
	Namespace                   types.String                        `tfsdk:"namespace"`
	KubernetesServiceIp         types.String                        `tfsdk:"kubernetes_service_ip"`
	Hostname                    types.String                        `tfsdk:"hostname"`
	NumHops                     types.Int64                         `tfsdk:"num_hops"`
	SpiffeIds                   types.List                          `tfsdk:"spiffe_ids"`
	NumVmEndpoints              types.Int64                         `tfsdk:"num_vm_endpoints"`
	Subsets                     types.List                          `tfsdk:"subsets"`
	NumKubernetesEndpoints      types.Int64                         `tfsdk:"num_kubernetes_endpoints"`
	Selector                    types.Map                           `tfsdk:"selector"`
	KubernetesServiceFqdn       types.String                        `tfsdk:"kubernetes_service_fqdn"`
	MeshExternal                types.Bool                          `tfsdk:"mesh_external"`
}
type NamespacesModel struct {
	Name     types.String               `tfsdk:"name"`
	Services []*ServicesNamespacesModel `tfsdk:"services"`
}
type LastSyncTimeStateModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type StateModel struct {
	TsbCpVersion  types.String           `tfsdk:"tsb_cp_version"`
	XcpVersion    types.String           `tfsdk:"xcp_version"`
	IstioVersions types.List             `tfsdk:"istio_versions"`
	LastSyncTime  LastSyncTimeStateModel `tfsdk:"last_sync_time"`
	Provider      types.String           `tfsdk:"provider"`
}
type EdgeInstallTemplateHelmSecretsXcpModel struct {
	Cert  types.String `tfsdk:"cert"`
	Key   types.String `tfsdk:"key"`
	Token types.String `tfsdk:"token"`
}
type XcpInstallTemplateHelmSecretsModel struct {
	Rootcakey         types.String                           `tfsdk:"rootcakey"`
	AutoGenerateCerts types.Bool                             `tfsdk:"auto_generate_certs"`
	Edge              EdgeInstallTemplateHelmSecretsXcpModel `tfsdk:"edge"`
	Rootca            types.String                           `tfsdk:"rootca"`
}
type ClusterServiceAccountInstallTemplateHelmSecretsModel struct {
	Encoded_JWK types.String `tfsdk:"encoded_jwk"`
	JWK         types.String `tfsdk:"jwk"`
	Cluster_FQN types.String `tfsdk:"cluster_fqn"`
}
type ElasticsearchInstallTemplateHelmSecretsModel struct {
	Cacert   types.String `tfsdk:"cacert"`
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}
type TsbInstallTemplateHelmSecretsModel struct {
	Cacert types.String `tfsdk:"cacert"`
}
type SecretsInstallTemplateHelmModel struct {
	Tsb                   TsbInstallTemplateHelmSecretsModel                   `tfsdk:"tsb"`
	Xcp                   XcpInstallTemplateHelmSecretsModel                   `tfsdk:"xcp"`
	ClusterServiceAccount ClusterServiceAccountInstallTemplateHelmSecretsModel `tfsdk:"cluster_service_account"`
	Elasticsearch         ElasticsearchInstallTemplateHelmSecretsModel         `tfsdk:"elasticsearch"`
}
type ManagementplaneRequestTimeoutInstallTemplateHelmSpecComponentsGitopsModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type ReconcileIntervalInstallTemplateHelmSpecComponentsGitopsModel struct {
	Seconds types.Int64 `tfsdk:"seconds"`
	Nanos   types.Int64 `tfsdk:"nanos"`
}
type ReconcileRequestTimeoutInstallTemplateHelmSpecComponentsGitopsModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type WebhookTimeoutInstallTemplateHelmSpecComponentsGitopsModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type BatchWindowInstallTemplateHelmSpecComponentsGitopsModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type GitopsInstallTemplateHelmSpecComponentsModel struct {
	ReconcileRequestTimeout       ReconcileRequestTimeoutInstallTemplateHelmSpecComponentsGitopsModel       `tfsdk:"reconcile_request_timeout"`
	WebhookTimeout                WebhookTimeoutInstallTemplateHelmSpecComponentsGitopsModel                `tfsdk:"webhook_timeout"`
	BatchWindow                   BatchWindowInstallTemplateHelmSpecComponentsGitopsModel                   `tfsdk:"batch_window"`
	Enabled                       types.Bool                                                                `tfsdk:"enabled"`
	ManagementplaneRequestTimeout ManagementplaneRequestTimeoutInstallTemplateHelmSpecComponentsGitopsModel `tfsdk:"managementplane_request_timeout"`
	ReconcileInterval             ReconcileIntervalInstallTemplateHelmSpecComponentsGitopsModel             `tfsdk:"reconcile_interval"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	ProcMount                types.String                                                                                             `tfsdk:"proc_mount"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	RunAsGroup               types.Int64                                                                                              `tfsdk:"run_as_group"`
	RunAsUser                types.Int64                                                                                              `tfsdk:"run_as_user"`
	Privileged               types.Bool                                                                                               `tfsdk:"privileged"`
	RunAsNonRoot             types.Bool                                                                                               `tfsdk:"run_as_non_root"`
	ReadOnlyRootFilesystem   types.Bool                                                                                               `tfsdk:"read_only_root_filesystem"`
	AllowPrivilegeEscalation types.Bool                                                                                               `tfsdk:"allow_privilege_escalation"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type SysctlsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	RunAsGroup          types.Int64                                                                                        `tfsdk:"run_as_group"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	FsGroup             types.Int64                                                                                        `tfsdk:"fs_group"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	SupplementalGroups  types.List                                                                                         `tfsdk:"supplemental_groups"`
	FsGroupChangePolicy types.String                                                                                       `tfsdk:"fs_group_change_policy"`
	RunAsNonRoot        types.Bool                                                                                         `tfsdk:"run_as_non_root"`
	RunAsUser           types.Int64                                                                                        `tfsdk:"run_as_user"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
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
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                 `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                            `tfsdk:"type"`
}
type ResourcesInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	Requests types.Map `tfsdk:"requests"`
	Limits   types.Map `tfsdk:"limits"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                     `tfsdk:"optional"`
	Key                  types.String                                                                                                   `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                      `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                        `tfsdk:"optional"`
}
type FieldRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel struct {
	FieldPath  types.String `tfsdk:"field_path"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	Type   types.Int64                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel struct {
	Resource      types.String                                                                                          `tfsdk:"resource"`
	ContainerName types.String                                                                                          `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
}
type ValueFromInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvModel struct {
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                   `tfsdk:"name"`
	Value     types.String                                                                   `tfsdk:"value"`
}
type TolerationsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                       `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	Type   types.Int64                                                                                                  `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
}
type PodsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel struct {
	Selector           SelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                           `tfsdk:"metric_name"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                      `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	Type   types.Int64                                                                                               `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
}
type TargetInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Value              ValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                               `tfsdk:"type"`
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
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                      `tfsdk:"type"`
}
type ResourceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel struct {
	Name                     types.String                                                                                                     `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                               `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchLabels      types.Map                                                                                                                 `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
}
type ExternalInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                               `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                         `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                             `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
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
type ObjectInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel struct {
	Target       TargetInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                       `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
}
type MetricsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecModel struct {
	Type     types.String                                                                             `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
}
type HpaSpecInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                         `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                         `tfsdk:"max_replicas"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                           `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                       `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                     `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
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
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                           `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                       `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                     `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                               `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                           `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                         `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                             `tfsdk:"weight"`
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
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                                     `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel struct {
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
}
type DeploymentInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel struct {
	Affinity                 AffinityInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	PodAnnotations           types.Map                                                                                  `tfsdk:"pod_annotations"`
	ReplicaCount             types.Int64                                                                                `tfsdk:"replica_count"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel                `tfsdk:"resources"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecDeploymentModel                   `tfsdk:"env"`
}
type ValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecOverlaysPatchesModel struct {
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
}
type PatchesInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecOverlaysModel struct {
	Value ValueInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecOverlaysPatchesModel `tfsdk:"value"`
	Path  types.String                                                                 `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel struct {
	Name       types.String                                                               `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                               `tfsdk:"api_version"`
	Kind       types.String                                                               `tfsdk:"kind"`
}
type StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServicePortsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
}
type PortsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServiceModel struct {
	TargetPort TargetPortInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                   `tfsdk:"name"`
	NodePort   types.Int64                                                                    `tfsdk:"node_port"`
	Port       types.Int64                                                                    `tfsdk:"port"`
	Protocol   types.String                                                                   `tfsdk:"protocol"`
}
type ServiceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel struct {
	Labels      types.Map                                                               `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                            `tfsdk:"type"`
	Annotations types.Map                                                               `tfsdk:"annotations"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type KubeSpecInstallTemplateHelmSpecComponentsHpaAdapterModel struct {
	Deployment     DeploymentInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsHpaAdapterKubeSpecModel `tfsdk:"service_account"`
}
type HpaAdapterInstallTemplateHelmSpecComponentsModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsHpaAdapterModel `tfsdk:"kube_spec"`
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
	Key                  types.String                                                                                                           `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                             `tfsdk:"optional"`
}
type FieldRefInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromModel struct {
	FieldPath  types.String `tfsdk:"field_path"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromModel struct {
	Divisor       DivisorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                               `tfsdk:"resource"`
	ContainerName types.String                                                                                               `tfsdk:"container_name"`
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
type ResourcesInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
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
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                      `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	Type          types.String                                                                                 `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel struct {
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	RunAsUser                types.Int64                                                                                                   `tfsdk:"run_as_user"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                    `tfsdk:"read_only_root_filesystem"`
	ProcMount                types.String                                                                                                  `tfsdk:"proc_mount"`
	RunAsNonRoot             types.Bool                                                                                                    `tfsdk:"run_as_non_root"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	Privileged               types.Bool                                                                                                    `tfsdk:"privileged"`
	RunAsGroup               types.Int64                                                                                                   `tfsdk:"run_as_group"`
	AllowPrivilegeEscalation types.Bool                                                                                                    `tfsdk:"allow_privilege_escalation"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
}
type TolerationsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
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
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                                    `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                     `tfsdk:"average_utilization"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                           `tfsdk:"type"`
}
type ResourceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel struct {
	Name                     types.String                                                                                                          `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
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
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
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
	Type   types.Int64                                                                                                    `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
}
type ExternalInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                                    `tfsdk:"metric_name"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Type   types.Int64                                                                                                  `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
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
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                              `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
}
type ObjectInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                            `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                            `tfsdk:"match_labels"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                       `tfsdk:"type"`
}
type PodsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
}
type MetricsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecModel struct {
	Type     types.String                                                                                  `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
}
type HpaSpecInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                              `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                              `tfsdk:"max_replicas"`
}
type SysctlsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	RunAsUser           types.Int64                                                                                             `tfsdk:"run_as_user"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsGroup          types.Int64                                                                                             `tfsdk:"run_as_group"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	FsGroup             types.Int64                                                                                             `tfsdk:"fs_group"`
	FsGroupChangePolicy types.String                                                                                            `tfsdk:"fs_group_change_policy"`
	SupplementalGroups  types.List                                                                                              `tfsdk:"supplemental_groups"`
	RunAsNonRoot        types.Bool                                                                                              `tfsdk:"run_as_non_root"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type PreferenceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                                          `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
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
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                            `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                          `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityModel struct {
	Weight          types.Int64                                                                                                                                              `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type PodAffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                    `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                                `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                              `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                  `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                                    `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                              `tfsdk:"topology_key"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
}
type DeploymentInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel struct {
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	ReplicaCount             types.Int64                                                                                     `tfsdk:"replica_count"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel                   `tfsdk:"env"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel                `tfsdk:"resources"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	PodAnnotations           types.Map                                                                                       `tfsdk:"pod_annotations"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecDeploymentModel                 `tfsdk:"affinity"`
}
type ValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecOverlaysPatchesModel struct {
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
}
type PatchesInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecOverlaysModel struct {
	Path  types.String                                                                      `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecOverlaysPatchesModel `tfsdk:"value"`
}
type OverlaysInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel struct {
	ApiVersion types.String                                                                    `tfsdk:"api_version"`
	Kind       types.String                                                                    `tfsdk:"kind"`
	Name       types.String                                                                    `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecOverlaysModel `tfsdk:"patches"`
}
type StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServicePortsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                               `tfsdk:"type"`
}
type PortsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServiceModel struct {
	Port       types.Int64                                                                         `tfsdk:"port"`
	Protocol   types.String                                                                        `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                        `tfsdk:"name"`
	NodePort   types.Int64                                                                         `tfsdk:"node_port"`
}
type ServiceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel struct {
	Labels      types.Map                                                                    `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                                 `tfsdk:"type"`
	Annotations types.Map                                                                    `tfsdk:"annotations"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type KubeSpecInstallTemplateHelmSpecComponentsRateLimitServerModel struct {
	Deployment     DeploymentInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsRateLimitServerKubeSpecModel `tfsdk:"service_account"`
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
type ConfigProtectionInstallTemplateHelmSpecComponentsXcpModel struct {
	AuthorizedUsers                                types.List `tfsdk:"authorized_users"`
	EnableAuthorizedCreateUpdateDeleteOnXcpConfigs types.Bool `tfsdk:"enable_authorized_create_update_delete_on_xcp_configs"`
	EnableAuthorizedUpdateDeleteOnXcpConfigs       types.Bool `tfsdk:"enable_authorized_update_delete_on_xcp_configs"`
}
type DefaultWorkloadCertTtlInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type ValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioBaseOverlaysPatchesModel struct {
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
}
type PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioBaseOverlaysModel struct {
	Value ValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioBaseOverlaysPatchesModel `tfsdk:"value"`
	Path  types.String                                                                                       `tfsdk:"path"`
}
type BaseOverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel struct {
	ApiVersion types.String                                                                                     `tfsdk:"api_version"`
	Kind       types.String                                                                                     `tfsdk:"kind"`
	Name       types.String                                                                                     `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioBaseOverlaysModel `tfsdk:"patches"`
}
type ResourcesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type SysctlsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SupplementalGroups  types.List                                                                                                                   `tfsdk:"supplemental_groups"`
	FsGroupChangePolicy types.String                                                                                                                 `tfsdk:"fs_group_change_policy"`
	RunAsNonRoot        types.Bool                                                                                                                   `tfsdk:"run_as_non_root"`
	FsGroup             types.Int64                                                                                                                  `tfsdk:"fs_group"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsGroup          types.Int64                                                                                                                  `tfsdk:"run_as_group"`
	RunAsUser           types.Int64                                                                                                                  `tfsdk:"run_as_user"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                                               `tfsdk:"optional"`
	Key                  types.String                                                                                                                             `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
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
type ValueFromInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvModel struct {
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	Value     types.String                                                                                             `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                                             `tfsdk:"name"`
}
type TolerationsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                                     `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                 `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                               `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityModel struct {
	Weight          types.Int64                                                                                                                                                                   `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                     `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                 `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                               `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
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
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                         `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                     `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                   `tfsdk:"topology_key"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type PreferenceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                                                               `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                                                      `tfsdk:"type"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	AllowPrivilegeEscalation types.Bool                                                                                                                         `tfsdk:"allow_privilege_escalation"`
	RunAsGroup               types.Int64                                                                                                                        `tfsdk:"run_as_group"`
	RunAsUser                types.Int64                                                                                                                        `tfsdk:"run_as_user"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                                         `tfsdk:"read_only_root_filesystem"`
	RunAsNonRoot             types.Bool                                                                                                                         `tfsdk:"run_as_non_root"`
	Privileged               types.Bool                                                                                                                         `tfsdk:"privileged"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	ProcMount                types.String                                                                                                                       `tfsdk:"proc_mount"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                           `tfsdk:"match_labels"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                         `tfsdk:"type"`
}
type ExternalInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                                         `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
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
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                   `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                       `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
}
type ObjectInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	AverageValue AverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                                                 `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
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
	Type   types.Int64                                                                                                                            `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
}
type PodsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                                     `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	Type   types.Int64                                                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                         `tfsdk:"type"`
}
type TargetInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Value              ValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                                          `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                                                         `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                      `tfsdk:"type"`
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
type ResourceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	Name                     types.String                                                                                                                               `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
}
type MetricsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecModel struct {
	Type     types.String                                                                                                       `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
}
type HpaSpecInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel struct {
	MinReplicas types.Int64                                                                                                   `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                                                   `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
}
type DeploymentInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel struct {
	Env                      []*EnvInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel                   `tfsdk:"env"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel                `tfsdk:"resources"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	ReplicaCount             types.Int64                                                                                                          `tfsdk:"replica_count"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	PodAnnotations           types.Map                                                                                                            `tfsdk:"pod_annotations"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecDeploymentModel                 `tfsdk:"strategy"`
}
type ValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecOverlaysPatchesModel struct {
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
}
type PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecOverlaysModel struct {
	Path  types.String                                                                                           `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecOverlaysPatchesModel `tfsdk:"value"`
}
type OverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel struct {
	ApiVersion types.String                                                                                         `tfsdk:"api_version"`
	Kind       types.String                                                                                         `tfsdk:"kind"`
	Name       types.String                                                                                         `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecOverlaysModel `tfsdk:"patches"`
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
	TargetPort TargetPortInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                                             `tfsdk:"name"`
	NodePort   types.Int64                                                                                              `tfsdk:"node_port"`
	Port       types.Int64                                                                                              `tfsdk:"port"`
	Protocol   types.String                                                                                             `tfsdk:"protocol"`
}
type ServiceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel struct {
	Ports       []*PortsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                                                      `tfsdk:"type"`
	Annotations types.Map                                                                                         `tfsdk:"annotations"`
	Labels      types.Map                                                                                         `tfsdk:"labels"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type CniInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel struct {
	BinaryDirectory        types.String `tfsdk:"binary_directory"`
	Chained                types.Bool   `tfsdk:"chained"`
	ClusterRole            types.String `tfsdk:"cluster_role"`
	ConfigurationDirectory types.String `tfsdk:"configuration_directory"`
	ConfigurationFileName  types.String `tfsdk:"configuration_file_name"`
	Revision               types.String `tfsdk:"revision"`
}
type KubeSpecInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel struct {
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel `tfsdk:"service_account"`
	Cni            CniInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel            `tfsdk:"cni"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioKubeSpecModel        `tfsdk:"service"`
}
type ValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioCniOverlaysPatchesModel struct {
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
}
type PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioCniOverlaysModel struct {
	Path  types.String                                                                                      `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioCniOverlaysPatchesModel `tfsdk:"value"`
}
type CniOverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel struct {
	ApiVersion types.String                                                                                    `tfsdk:"api_version"`
	Kind       types.String                                                                                    `tfsdk:"kind"`
	Name       types.String                                                                                    `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioCniOverlaysModel `tfsdk:"patches"`
}
type MaxWorkloadCertTtlInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel struct {
	Seconds types.Int64 `tfsdk:"seconds"`
	Nanos   types.Int64 `tfsdk:"nanos"`
}
type ValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioPilotOverlaysPatchesModel struct {
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
}
type PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioPilotOverlaysModel struct {
	Value ValueInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioPilotOverlaysPatchesModel `tfsdk:"value"`
	Path  types.String                                                                                        `tfsdk:"path"`
}
type PilotOverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel struct {
	Kind       types.String                                                                                      `tfsdk:"kind"`
	Name       types.String                                                                                      `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioPilotOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                                      `tfsdk:"api_version"`
}
type IstioInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsModel struct {
	TsbVersion             types.String                                                                                     `tfsdk:"tsb_version"`
	TraceSamplingRate      types.Float64                                                                                    `tfsdk:"trace_sampling_rate"`
	TrustDomain            types.String                                                                                     `tfsdk:"trust_domain"`
	DefaultWorkloadCertTtl DefaultWorkloadCertTtlInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel `tfsdk:"default_workload_cert_ttl"`
	BaseOverlays           []*BaseOverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel        `tfsdk:"base_overlays"`
	MaxWorkloadCertTtl     MaxWorkloadCertTtlInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel     `tfsdk:"max_workload_cert_ttl"`
	KubeSpec               KubeSpecInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel               `tfsdk:"kube_spec"`
	CniOverlays            []*CniOverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel         `tfsdk:"cni_overlays"`
	PilotOverlays          []*PilotOverlaysInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsIstioModel       `tfsdk:"pilot_overlays"`
}
type RevisionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesModel struct {
	Disable types.Bool                                                                 `tfsdk:"disable"`
	Istio   IstioInstallTemplateHelmSpecComponentsXcpIsolationBoundariesRevisionsModel `tfsdk:"istio"`
	Name    types.String                                                               `tfsdk:"name"`
}
type IsolationBoundariesInstallTemplateHelmSpecComponentsXcpModel struct {
	Revisions []*RevisionsInstallTemplateHelmSpecComponentsXcpIsolationBoundariesModel `tfsdk:"revisions"`
	Name      types.String                                                             `tfsdk:"name"`
}
type ValueInstallTemplateHelmSpecComponentsXcpKubeSpecOverlaysPatchesModel struct {
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
}
type PatchesInstallTemplateHelmSpecComponentsXcpKubeSpecOverlaysModel struct {
	Path  types.String                                                          `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsXcpKubeSpecOverlaysPatchesModel `tfsdk:"value"`
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
	Type   types.Int64                                                                   `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
}
type PortsInstallTemplateHelmSpecComponentsXcpKubeSpecServiceModel struct {
	Protocol   types.String                                                            `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsXcpKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                            `tfsdk:"name"`
	NodePort   types.Int64                                                             `tfsdk:"node_port"`
	Port       types.Int64                                                             `tfsdk:"port"`
}
type ServiceInstallTemplateHelmSpecComponentsXcpKubeSpecModel struct {
	Annotations types.Map                                                        `tfsdk:"annotations"`
	Labels      types.Map                                                        `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsXcpKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                     `tfsdk:"type"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsXcpKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsXcpKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsXcpKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type TolerationsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchLabels      types.Map                                                                                                          `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                               `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                        `tfsdk:"type"`
}
type ExternalInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                        `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                       `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
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
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
}
type ObjectInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName   types.String                                                                                `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	Type   types.Int64                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                `tfsdk:"match_labels"`
}
type PodsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                    `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                               `tfsdk:"type"`
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
type TargetInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	AverageUtilization types.Int64                                                                                         `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                        `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                               `tfsdk:"type"`
}
type ResourceInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel struct {
	Target                   TargetInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                              `tfsdk:"name"`
}
type MetricsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecModel struct {
	External ExternalInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                      `tfsdk:"type"`
}
type HpaSpecInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	MinReplicas types.Int64                                                                  `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                  `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                        `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
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
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                      `tfsdk:"weight"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type PreferenceInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                              `tfsdk:"weight"`
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
type NodeAffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                    `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                              `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                  `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                    `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityModel struct {
	TopologyKey   types.String                                                                                                                              `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                `tfsdk:"namespaces"`
}
type PodAffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                          `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromModel struct {
	Resource      types.String                                                                                   `tfsdk:"resource"`
	ContainerName types.String                                                                                   `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
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
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                 `tfsdk:"optional"`
	Key                  types.String                                                                                               `tfsdk:"key"`
}
type FieldRefInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
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
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                    `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
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
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                     `tfsdk:"type"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type SysctlsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	RunAsUser           types.Int64                                                                                 `tfsdk:"run_as_user"`
	RunAsGroup          types.Int64                                                                                 `tfsdk:"run_as_group"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	FsGroup             types.Int64                                                                                 `tfsdk:"fs_group"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	SupplementalGroups  types.List                                                                                  `tfsdk:"supplemental_groups"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsNonRoot        types.Bool                                                                                  `tfsdk:"run_as_non_root"`
	FsGroupChangePolicy types.String                                                                                `tfsdk:"fs_group_change_policy"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	ProcMount                types.String                                                                                      `tfsdk:"proc_mount"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	AllowPrivilegeEscalation types.Bool                                                                                        `tfsdk:"allow_privilege_escalation"`
	Privileged               types.Bool                                                                                        `tfsdk:"privileged"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	ReadOnlyRootFilesystem   types.Bool                                                                                        `tfsdk:"read_only_root_filesystem"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsUser                types.Int64                                                                                       `tfsdk:"run_as_user"`
	RunAsGroup               types.Int64                                                                                       `tfsdk:"run_as_group"`
	RunAsNonRoot             types.Bool                                                                                        `tfsdk:"run_as_non_root"`
}
type ResourcesInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type DeploymentInstallTemplateHelmSpecComponentsXcpKubeSpecModel struct {
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel                `tfsdk:"resources"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	PodAnnotations           types.Map                                                                           `tfsdk:"pod_annotations"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel                   `tfsdk:"env"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsXcpKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	ReplicaCount             types.Int64                                                                         `tfsdk:"replica_count"`
}
type KubeSpecInstallTemplateHelmSpecComponentsXcpModel struct {
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsXcpKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsXcpKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsXcpKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsXcpKubeSpecModel     `tfsdk:"deployment"`
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
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
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
	Optional             types.Bool                                                                                              `tfsdk:"optional"`
	Key                  types.String                                                                                            `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
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
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type ValueFromInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvModel struct {
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	Name      types.String                                                            `tfsdk:"name"`
	Value     types.String                                                            `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	Type   types.Int64                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                `tfsdk:"match_labels"`
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
type AverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                               `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                        `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
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
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                               `tfsdk:"type"`
}
type ResourceInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel struct {
	Target                   TargetInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                              `tfsdk:"name"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
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
type ExternalInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                        `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
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
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Type   types.Int64                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Type   types.Int64                                                                                       `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
}
type ObjectInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel struct {
	Target       TargetInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
}
type MetricsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecModel struct {
	Pods     PodsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                      `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
}
type HpaSpecInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	MinReplicas types.Int64                                                                  `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                  `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
}
type ResourcesInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	Requests types.Map `tfsdk:"requests"`
	Limits   types.Map `tfsdk:"limits"`
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
type StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                    `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyModel struct {
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
}
type StrategyInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	Type          types.String                                                                     `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type TolerationsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
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
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                    `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                              `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                  `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                    `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                              `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
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
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                    `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                  `tfsdk:"topology_key"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	AllowPrivilegeEscalation types.Bool                                                                                        `tfsdk:"allow_privilege_escalation"`
	Privileged               types.Bool                                                                                        `tfsdk:"privileged"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	ProcMount                types.String                                                                                      `tfsdk:"proc_mount"`
	ReadOnlyRootFilesystem   types.Bool                                                                                        `tfsdk:"read_only_root_filesystem"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	RunAsNonRoot             types.Bool                                                                                        `tfsdk:"run_as_non_root"`
	RunAsGroup               types.Int64                                                                                       `tfsdk:"run_as_group"`
	RunAsUser                types.Int64                                                                                       `tfsdk:"run_as_user"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type SysctlsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel struct {
	RunAsGroup          types.Int64                                                                                 `tfsdk:"run_as_group"`
	FsGroup             types.Int64                                                                                 `tfsdk:"fs_group"`
	RunAsNonRoot        types.Bool                                                                                  `tfsdk:"run_as_non_root"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	FsGroupChangePolicy types.String                                                                                `tfsdk:"fs_group_change_policy"`
	RunAsUser           types.Int64                                                                                 `tfsdk:"run_as_user"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	SupplementalGroups  types.List                                                                                  `tfsdk:"supplemental_groups"`
}
type DeploymentInstallTemplateHelmSpecComponentsOapKubeSpecModel struct {
	Env                      []*EnvInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel                   `tfsdk:"env"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	ReplicaCount             types.Int64                                                                         `tfsdk:"replica_count"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel                `tfsdk:"resources"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsOapKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	PodAnnotations           types.Map                                                                           `tfsdk:"pod_annotations"`
}
type ValueInstallTemplateHelmSpecComponentsOapKubeSpecOverlaysPatchesModel struct {
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
}
type PatchesInstallTemplateHelmSpecComponentsOapKubeSpecOverlaysModel struct {
	Path  types.String                                                          `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsOapKubeSpecOverlaysPatchesModel `tfsdk:"value"`
}
type OverlaysInstallTemplateHelmSpecComponentsOapKubeSpecModel struct {
	Kind       types.String                                                        `tfsdk:"kind"`
	Name       types.String                                                        `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsOapKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                        `tfsdk:"api_version"`
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
type ImagePullSecretsInstallTemplateHelmSpecComponentsOapKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsOapKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsOapKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
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
type ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
}
type ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                                                               `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                                                                `tfsdk:"type"`
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
type PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel struct {
	Selector           SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                                                               `tfsdk:"metric_name"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                   `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
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
	Type               types.String                                                                                                                                   `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                                                    `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                `tfsdk:"type"`
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
	Name                     types.String                                                                                                                                         `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
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
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                          `tfsdk:"type"`
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
	Type   types.Int64                                                                                                                                  `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                             `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                 `tfsdk:"type"`
}
type ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel struct {
	AverageValue AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                                                           `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
}
type MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecModel struct {
	External ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                                                                 `tfsdk:"type"`
}
type HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                                                             `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                                                             `tfsdk:"max_replicas"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsGroup          types.Int64                                                                                                                            `tfsdk:"run_as_group"`
	RunAsUser           types.Int64                                                                                                                            `tfsdk:"run_as_user"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	FsGroup             types.Int64                                                                                                                            `tfsdk:"fs_group"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SupplementalGroups  types.List                                                                                                                             `tfsdk:"supplemental_groups"`
	FsGroupChangePolicy types.String                                                                                                                           `tfsdk:"fs_group_change_policy"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsNonRoot        types.Bool                                                                                                                             `tfsdk:"run_as_non_root"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	RunAsUser                types.Int64                                                                                                                                  `tfsdk:"run_as_user"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	AllowPrivilegeEscalation types.Bool                                                                                                                                   `tfsdk:"allow_privilege_escalation"`
	RunAsNonRoot             types.Bool                                                                                                                                   `tfsdk:"run_as_non_root"`
	Privileged               types.Bool                                                                                                                                   `tfsdk:"privileged"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                                                   `tfsdk:"read_only_root_filesystem"`
	RunAsGroup               types.Int64                                                                                                                                  `tfsdk:"run_as_group"`
	ProcMount                types.String                                                                                                                                 `tfsdk:"proc_mount"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
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
	TopologyKey   types.String                                                                                                                                                                         `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                           `tfsdk:"namespaces"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                               `tfsdk:"match_labels"`
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
type PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
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
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                   `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                               `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                             `tfsdk:"topology_key"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                                                         `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel struct {
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                                                         `tfsdk:"optional"`
	Key                  types.String                                                                                                                                       `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                                                            `tfsdk:"optional"`
	Key                  types.String                                                                                                                                          `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
}
type FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
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
	Name      types.String                                                                                                       `tfsdk:"name"`
	Value     types.String                                                                                                       `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
}
type DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecModel struct {
	ReplicaCount             types.Int64                                                                                                                    `tfsdk:"replica_count"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel                `tfsdk:"resources"`
	PodAnnotations           types.Map                                                                                                                      `tfsdk:"pod_annotations"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel                   `tfsdk:"env"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecDeploymentModel                 `tfsdk:"affinity"`
}
type ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecOverlaysPatchesModel struct {
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
}
type PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecOverlaysModel struct {
	Path  types.String                                                                                                     `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecOverlaysPatchesModel `tfsdk:"value"`
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
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                              `tfsdk:"type"`
}
type PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServiceModel struct {
	Port       types.Int64                                                                                                        `tfsdk:"port"`
	Protocol   types.String                                                                                                       `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                                                       `tfsdk:"name"`
	NodePort   types.Int64                                                                                                        `tfsdk:"node_port"`
}
type ServiceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecModel struct {
	Type        types.String                                                                                                `tfsdk:"type"`
	Annotations types.Map                                                                                                   `tfsdk:"annotations"`
	Labels      types.Map                                                                                                   `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerSpecKubeSpecServiceModel `tfsdk:"ports"`
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
type ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecOverlaysPatchesModel struct {
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
}
type PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecOverlaysModel struct {
	Path  types.String                                                                                                                `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecOverlaysPatchesModel `tfsdk:"value"`
}
type OverlaysInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel struct {
	Name       types.String                                                                                                              `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                                                              `tfsdk:"api_version"`
	Kind       types.String                                                                                                              `tfsdk:"kind"`
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
	Protocol   types.String                                                                                                                  `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                                                                  `tfsdk:"name"`
	NodePort   types.Int64                                                                                                                   `tfsdk:"node_port"`
	Port       types.Int64                                                                                                                   `tfsdk:"port"`
}
type ServiceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel struct {
	Annotations types.Map                                                                                                              `tfsdk:"annotations"`
	Labels      types.Map                                                                                                              `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                                                                           `tfsdk:"type"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
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
	SupplementalGroups  types.List                                                                                                                                        `tfsdk:"supplemental_groups"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	FsGroupChangePolicy types.String                                                                                                                                      `tfsdk:"fs_group_change_policy"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsUser           types.Int64                                                                                                                                       `tfsdk:"run_as_user"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsGroup          types.Int64                                                                                                                                       `tfsdk:"run_as_group"`
	RunAsNonRoot        types.Bool                                                                                                                                        `tfsdk:"run_as_non_root"`
	FsGroup             types.Int64                                                                                                                                       `tfsdk:"fs_group"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
}
type ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	AllowPrivilegeEscalation types.Bool                                                                                                                                              `tfsdk:"allow_privilege_escalation"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	Privileged               types.Bool                                                                                                                                              `tfsdk:"privileged"`
	ProcMount                types.String                                                                                                                                            `tfsdk:"proc_mount"`
	RunAsUser                types.Int64                                                                                                                                             `tfsdk:"run_as_user"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsNonRoot             types.Bool                                                                                                                                              `tfsdk:"run_as_non_root"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                                                              `tfsdk:"read_only_root_filesystem"`
	RunAsGroup               types.Int64                                                                                                                                             `tfsdk:"run_as_group"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
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
	Type          types.String                                                                                                                           `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                                                                                    `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
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
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
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
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                                                          `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
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
type PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                              `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Namespaces    types.List                                                                                                                                                                                          `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                        `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                                              `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                                                                                        `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                                          `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                                                                                            `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
}
type FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
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
	Optional             types.Bool                                                                                                                                                    `tfsdk:"optional"`
	Key                  types.String                                                                                                                                                  `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                                                                     `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                                                       `tfsdk:"optional"`
}
type ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvModel struct {
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	Name      types.String                                                                                                                  `tfsdk:"name"`
	Value     types.String                                                                                                                  `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                      `tfsdk:"match_labels"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
}
type PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                                                          `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
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
	Type   types.Int64                                                                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	AverageUtilization types.Int64                                                                                                                                               `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                                                                              `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                                                                    `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                     `tfsdk:"type"`
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
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                                                                              `tfsdk:"metric_name"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                        `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Type   types.Int64                                                                                                                                            `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
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
type ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel struct {
	Target       TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                                                                      `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
}
type MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecModel struct {
	Pods     PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                                                                            `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
}
type HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel struct {
	MaxReplicas types.Int64                                                                                                                        `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                                                                        `tfsdk:"min_replicas"`
}
type DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel struct {
	PodAnnotations           types.Map                                                                                                                                 `tfsdk:"pod_annotations"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel                   `tfsdk:"env"`
	ReplicaCount             types.Int64                                                                                                                               `tfsdk:"replica_count"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel                `tfsdk:"resources"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
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
	Namespaces    types.List                                                                                                                                                                                               `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                             `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                                                 `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
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
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                                                       `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                                                                   `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                                 `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                                                                                     `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
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
	TopologyKey   types.String                                                                                                                                                                                 `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                   `tfsdk:"namespaces"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                                                             `tfsdk:"weight"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel struct {
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobAffinityModel    `tfsdk:"node_affinity"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel struct {
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel `tfsdk:"windows_options"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel   `tfsdk:"capabilities"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	Privileged               types.Bool                                                                                                                                       `tfsdk:"privileged"`
	RunAsGroup               types.Int64                                                                                                                                      `tfsdk:"run_as_group"`
	RunAsNonRoot             types.Bool                                                                                                                                       `tfsdk:"run_as_non_root"`
	RunAsUser                types.Int64                                                                                                                                      `tfsdk:"run_as_user"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobContainerSecurityContextModel `tfsdk:"se_linux_options"`
	ProcMount                types.String                                                                                                                                     `tfsdk:"proc_mount"`
	AllowPrivilegeEscalation types.Bool                                                                                                                                       `tfsdk:"allow_privilege_escalation"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                                                       `tfsdk:"read_only_root_filesystem"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel struct {
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                                             `tfsdk:"optional"`
	Key                  types.String                                                                                                                                           `tfsdk:"key"`
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
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromResourceFieldRefModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                         `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel struct {
	Divisor       DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                                                                  `tfsdk:"resource"`
	ContainerName types.String                                                                                                                                  `tfsdk:"container_name"`
}
type ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvModel struct {
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvValueFromModel `tfsdk:"resource_field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel struct {
	Value     types.String                                                                                                           `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                                                           `tfsdk:"name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel struct {
	RunAsUser           types.Int64                                                                                                                                `tfsdk:"run_as_user"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel     `tfsdk:"sysctls"`
	RunAsGroup          types.Int64                                                                                                                                `tfsdk:"run_as_group"`
	RunAsNonRoot        types.Bool                                                                                                                                 `tfsdk:"run_as_non_root"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel `tfsdk:"windows_options"`
	FsGroup             types.Int64                                                                                                                                `tfsdk:"fs_group"`
	FsGroupChangePolicy types.String                                                                                                                               `tfsdk:"fs_group_change_policy"`
	SupplementalGroups  types.List                                                                                                                                 `tfsdk:"supplemental_groups"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobPodSecurityContextModel `tfsdk:"seccomp_profile"`
}
type TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel struct {
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
}
type JobInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecModel struct {
	Affinity                 AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel                 `tfsdk:"affinity"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel `tfsdk:"container_security_context"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel                   `tfsdk:"env"`
	PodAnnotations           types.Map                                                                                                                          `tfsdk:"pod_annotations"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel       `tfsdk:"pod_security_context"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerStartupapicheckKubeSpecJobModel           `tfsdk:"tolerations"`
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
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServicePortsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
}
type PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServiceModel struct {
	NodePort   types.Int64                                                                                                               `tfsdk:"node_port"`
	Port       types.Int64                                                                                                               `tfsdk:"port"`
	Protocol   types.String                                                                                                              `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                                                              `tfsdk:"name"`
}
type ServiceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel struct {
	Labels      types.Map                                                                                                          `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                                                                       `tfsdk:"type"`
	Annotations types.Map                                                                                                          `tfsdk:"annotations"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	Type   types.Int64                                                                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
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
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                                                                          `tfsdk:"metric_name"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                    `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
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
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
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
	Type   types.Int64                                                                                                                                             `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
}
type PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                                                      `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Value              ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                                                           `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                                                                          `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                       `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel struct {
	Name                     types.String                                                                                                                                                `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
}
type MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecModel struct {
	External ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                                                                        `tfsdk:"type"`
}
type HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	MinReplicas types.Int64                                                                                                                    `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                                                                    `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
}
type ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	Requests types.Map `tfsdk:"requests"`
	Limits   types.Map `tfsdk:"limits"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
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
type NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
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
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                      `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityModel struct {
	Namespaces    types.List                                                                                                                                                                                  `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                                          `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                                                    `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                                                        `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                          `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	TopologyKey   types.String                                                                                                                                                                                    `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                      `tfsdk:"namespaces"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	RunAsUser           types.Int64                                                                                                                                   `tfsdk:"run_as_user"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SupplementalGroups  types.List                                                                                                                                    `tfsdk:"supplemental_groups"`
	FsGroupChangePolicy types.String                                                                                                                                  `tfsdk:"fs_group_change_policy"`
	RunAsGroup          types.Int64                                                                                                                                   `tfsdk:"run_as_group"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	FsGroup             types.Int64                                                                                                                                   `tfsdk:"fs_group"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsNonRoot        types.Bool                                                                                                                                    `tfsdk:"run_as_non_root"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	Type   types.Int64                                                                                                                                            `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel struct {
	Resource      types.String                                                                                                                                     `tfsdk:"resource"`
	ContainerName types.String                                                                                                                                     `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel struct {
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                                                `tfsdk:"optional"`
	Key                  types.String                                                                                                                                              `tfsdk:"key"`
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
	FieldPath  types.String `tfsdk:"field_path"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvModel struct {
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                                                              `tfsdk:"name"`
	Value     types.String                                                                                                              `tfsdk:"value"`
}
type TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                            `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                                                                       `tfsdk:"type"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel struct {
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	RunAsNonRoot             types.Bool                                                                                                                                          `tfsdk:"run_as_non_root"`
	RunAsUser                types.Int64                                                                                                                                         `tfsdk:"run_as_user"`
	ProcMount                types.String                                                                                                                                        `tfsdk:"proc_mount"`
	AllowPrivilegeEscalation types.Bool                                                                                                                                          `tfsdk:"allow_privilege_escalation"`
	Privileged               types.Bool                                                                                                                                          `tfsdk:"privileged"`
	RunAsGroup               types.Int64                                                                                                                                         `tfsdk:"run_as_group"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                                                          `tfsdk:"read_only_root_filesystem"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
}
type DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel struct {
	ReplicaCount             types.Int64                                                                                                                           `tfsdk:"replica_count"`
	PodAnnotations           types.Map                                                                                                                             `tfsdk:"pod_annotations"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel                `tfsdk:"resources"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel                   `tfsdk:"env"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecDeploymentModel `tfsdk:"container_security_context"`
}
type ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecOverlaysPatchesModel struct {
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
}
type PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecOverlaysModel struct {
	Value ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecOverlaysPatchesModel `tfsdk:"value"`
	Path  types.String                                                                                                            `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel struct {
	Patches    []*PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                                                          `tfsdk:"api_version"`
	Kind       types.String                                                                                                          `tfsdk:"kind"`
	Name       types.String                                                                                                          `tfsdk:"name"`
}
type KubeSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecModel struct {
	Service        ServiceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecKubeSpecModel    `tfsdk:"overlays"`
}
type CertManagerWebhookSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerWebhookSpecModel `tfsdk:"kube_spec"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServicePortsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                    `tfsdk:"type"`
}
type PortsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServiceModel struct {
	Protocol   types.String                                                                                                             `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                                                             `tfsdk:"name"`
	NodePort   types.Int64                                                                                                              `tfsdk:"node_port"`
	Port       types.Int64                                                                                                              `tfsdk:"port"`
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
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                 `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                            `tfsdk:"type"`
}
type PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                                                     `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
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
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	Name                     types.String                                                                                                                                               `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	Type   types.Int64                                                                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                         `tfsdk:"type"`
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
type ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                                                                         `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
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
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                   `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                       `tfsdk:"type"`
}
type ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                                                                 `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
}
type MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecModel struct {
	Pods     PodsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                                                                       `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
}
type HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	MaxReplicas types.Int64                                                                                                                   `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                                                                   `tfsdk:"min_replicas"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                                                                               `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
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
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                                                     `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                                                                               `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                                 `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityModel struct {
	Weight          types.Int64                                                                                                                                                                                   `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
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
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
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
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                                         `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	TopologyKey   types.String                                                                                                                                                                                   `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                                     `tfsdk:"namespaces"`
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
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
}
type IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                     `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                                                                      `tfsdk:"type"`
}
type TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                                                                  `tfsdk:"optional"`
	Key                  types.String                                                                                                                                                `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
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
	IntVal IntValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                                           `tfsdk:"type"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel struct {
	Divisor       DivisorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                                                                    `tfsdk:"resource"`
	ContainerName types.String                                                                                                                                    `tfsdk:"container_name"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                                                             `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                                               `tfsdk:"optional"`
}
type ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvModel struct {
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	Name      types.String                                                                                                             `tfsdk:"name"`
	Value     types.String                                                                                                             `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	RunAsNonRoot             types.Bool                                                                                                                                         `tfsdk:"run_as_non_root"`
	RunAsGroup               types.Int64                                                                                                                                        `tfsdk:"run_as_group"`
	ProcMount                types.String                                                                                                                                       `tfsdk:"proc_mount"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                                                         `tfsdk:"read_only_root_filesystem"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	AllowPrivilegeEscalation types.Bool                                                                                                                                         `tfsdk:"allow_privilege_escalation"`
	Privileged               types.Bool                                                                                                                                         `tfsdk:"privileged"`
	RunAsUser                types.Int64                                                                                                                                        `tfsdk:"run_as_user"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
}
type ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel struct {
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsUser           types.Int64                                                                                                                                  `tfsdk:"run_as_user"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	FsGroup             types.Int64                                                                                                                                  `tfsdk:"fs_group"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	RunAsNonRoot        types.Bool                                                                                                                                   `tfsdk:"run_as_non_root"`
	RunAsGroup          types.Int64                                                                                                                                  `tfsdk:"run_as_group"`
	SupplementalGroups  types.List                                                                                                                                   `tfsdk:"supplemental_groups"`
	FsGroupChangePolicy types.String                                                                                                                                 `tfsdk:"fs_group_change_policy"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
}
type DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel struct {
	Resources                ResourcesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel                `tfsdk:"resources"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel                   `tfsdk:"env"`
	ReplicaCount             types.Int64                                                                                                                          `tfsdk:"replica_count"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	PodAnnotations           types.Map                                                                                                                            `tfsdk:"pod_annotations"`
}
type ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecOverlaysPatchesModel struct {
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
}
type PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecOverlaysModel struct {
	Value ValueInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecOverlaysPatchesModel `tfsdk:"value"`
	Path  types.String                                                                                                           `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel struct {
	Patches    []*PatchesInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                                                         `tfsdk:"api_version"`
	Kind       types.String                                                                                                         `tfsdk:"kind"`
	Name       types.String                                                                                                         `tfsdk:"name"`
}
type KubeSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorModel struct {
	Service        ServiceInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorKubeSpecModel    `tfsdk:"overlays"`
}
type CertManagerCaInjectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerCertManagerCaInjectorModel `tfsdk:"kube_spec"`
}
type CertManagerInstallTemplateHelmSpecComponentsInternalCertProviderModel struct {
	CertManagerSpec            CertManagerSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel            `tfsdk:"cert_manager_spec"`
	CertManagerStartupapicheck CertManagerStartupapicheckInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel `tfsdk:"cert_manager_startupapicheck"`
	CertManagerWebhookSpec     CertManagerWebhookSpecInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel     `tfsdk:"cert_manager_webhook_spec"`
	Managed                    types.String                                                                                    `tfsdk:"managed"`
	CertManagerCaInjector      CertManagerCaInjectorInstallTemplateHelmSpecComponentsInternalCertProviderCertManagerModel      `tfsdk:"cert_manager_ca_injector"`
}
type CustomInstallTemplateHelmSpecComponentsInternalCertProviderModel struct {
	CaBundleSecretName types.String `tfsdk:"ca_bundle_secret_name"`
	CsrSignerName      types.String `tfsdk:"csr_signer_name"`
}
type InternalCertProviderInstallTemplateHelmSpecComponentsModel struct {
	CertManager CertManagerInstallTemplateHelmSpecComponentsInternalCertProviderModel `tfsdk:"cert_manager"`
	Custom      CustomInstallTemplateHelmSpecComponentsInternalCertProviderModel      `tfsdk:"custom"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsDefaultKubeSpecAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type AccountInstallTemplateHelmSpecComponentsDefaultKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsDefaultKubeSpecAccountModel `tfsdk:"image_pull_secrets"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                                  `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
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
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
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
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                            `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                        `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                      `tfsdk:"topology_key"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                            `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
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
type PodAntiAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel struct {
	ProcMount                types.String                                                                                          `tfsdk:"proc_mount"`
	ReadOnlyRootFilesystem   types.Bool                                                                                            `tfsdk:"read_only_root_filesystem"`
	RunAsNonRoot             types.Bool                                                                                            `tfsdk:"run_as_non_root"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsUser                types.Int64                                                                                           `tfsdk:"run_as_user"`
	AllowPrivilegeEscalation types.Bool                                                                                            `tfsdk:"allow_privilege_escalation"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	RunAsGroup               types.Int64                                                                                           `tfsdk:"run_as_group"`
	Privileged               types.Bool                                                                                            `tfsdk:"privileged"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                  `tfsdk:"optional"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                   `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                     `tfsdk:"optional"`
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
	Divisor       DivisorInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                       `tfsdk:"resource"`
	ContainerName types.String                                                                                       `tfsdk:"container_name"`
}
type ValueFromInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvModel struct {
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel struct {
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                `tfsdk:"name"`
	Value     types.String                                                                `tfsdk:"value"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type SysctlsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel struct {
	FsGroupChangePolicy types.String                                                                                    `tfsdk:"fs_group_change_policy"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SupplementalGroups  types.List                                                                                      `tfsdk:"supplemental_groups"`
	RunAsGroup          types.Int64                                                                                     `tfsdk:"run_as_group"`
	RunAsNonRoot        types.Bool                                                                                      `tfsdk:"run_as_non_root"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	FsGroup             types.Int64                                                                                     `tfsdk:"fs_group"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsUser           types.Int64                                                                                     `tfsdk:"run_as_user"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
}
type IntValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                              `tfsdk:"type"`
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
type RollingUpdateInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyModel struct {
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
}
type StrategyInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                         `tfsdk:"type"`
}
type TolerationsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel struct {
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
}
type DeploymentInstallTemplateHelmSpecComponentsDefaultKubeSpecModel struct {
	Affinity                 AffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel                   `tfsdk:"env"`
	PodAnnotations           types.Map                                                                               `tfsdk:"pod_annotations"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsDefaultKubeSpecDeploymentModel           `tfsdk:"tolerations"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type SysctlsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel struct {
	FsGroupChangePolicy types.String                                                                             `tfsdk:"fs_group_change_policy"`
	RunAsGroup          types.Int64                                                                              `tfsdk:"run_as_group"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsUser           types.Int64                                                                              `tfsdk:"run_as_user"`
	FsGroup             types.Int64                                                                              `tfsdk:"fs_group"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel     `tfsdk:"sysctls"`
	SupplementalGroups  types.List                                                                               `tfsdk:"supplemental_groups"`
	RunAsNonRoot        types.Bool                                                                               `tfsdk:"run_as_non_root"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobPodSecurityContextModel `tfsdk:"windows_options"`
}
type TolerationsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel struct {
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
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
type MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type PreferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                           `tfsdk:"weight"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                 `tfsdk:"match_labels"`
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
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                 `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityModel struct {
	Namespaces    types.List                                                                                                                             `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                           `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
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
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                     `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                               `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                 `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                                   `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel struct {
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobAffinityModel     `tfsdk:"pod_affinity"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel struct {
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel struct {
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsUser                types.Int64                                                                                    `tfsdk:"run_as_user"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel `tfsdk:"windows_options"`
	ReadOnlyRootFilesystem   types.Bool                                                                                     `tfsdk:"read_only_root_filesystem"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel   `tfsdk:"capabilities"`
	RunAsGroup               types.Int64                                                                                    `tfsdk:"run_as_group"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsDefaultKubeSpecJobContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	AllowPrivilegeEscalation types.Bool                                                                                     `tfsdk:"allow_privilege_escalation"`
	RunAsNonRoot             types.Bool                                                                                     `tfsdk:"run_as_non_root"`
	Privileged               types.Bool                                                                                     `tfsdk:"privileged"`
	ProcMount                types.String                                                                                   `tfsdk:"proc_mount"`
}
type JobInstallTemplateHelmSpecComponentsDefaultKubeSpecModel struct {
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel `tfsdk:"container_security_context"`
	PodAnnotations           types.Map                                                                        `tfsdk:"pod_annotations"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel       `tfsdk:"pod_security_context"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel           `tfsdk:"tolerations"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsDefaultKubeSpecJobModel                 `tfsdk:"affinity"`
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
type ValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecOverlaysPatchesModel struct {
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
}
type PatchesInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecOverlaysModel struct {
	Path  types.String                                                                              `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecOverlaysPatchesModel `tfsdk:"value"`
}
type OverlaysInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel struct {
	Name       types.String                                                                            `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                            `tfsdk:"api_version"`
	Kind       types.String                                                                            `tfsdk:"kind"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServicePortsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                       `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
}
type PortsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServiceModel struct {
	Port       types.Int64                                                                                 `tfsdk:"port"`
	Protocol   types.String                                                                                `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                                `tfsdk:"name"`
	NodePort   types.Int64                                                                                 `tfsdk:"node_port"`
}
type ServiceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel struct {
	Labels      types.Map                                                                            `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                                         `tfsdk:"type"`
	Annotations types.Map                                                                            `tfsdk:"annotations"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                                  `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                        `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                    `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                  `tfsdk:"topology_key"`
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
	TopologyKey   types.String                                                                                                                                                  `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                    `tfsdk:"namespaces"`
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
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                        `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                      `tfsdk:"topology_key"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	AllowPrivilegeEscalation types.Bool                                                                                                            `tfsdk:"allow_privilege_escalation"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	ReadOnlyRootFilesystem   types.Bool                                                                                                            `tfsdk:"read_only_root_filesystem"`
	RunAsGroup               types.Int64                                                                                                           `tfsdk:"run_as_group"`
	RunAsNonRoot             types.Bool                                                                                                            `tfsdk:"run_as_non_root"`
	RunAsUser                types.Int64                                                                                                           `tfsdk:"run_as_user"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	Privileged               types.Bool                                                                                                            `tfsdk:"privileged"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	ProcMount                types.String                                                                                                          `tfsdk:"proc_mount"`
}
type SysctlsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	FsGroupChangePolicy types.String                                                                                                    `tfsdk:"fs_group_change_policy"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	FsGroup             types.Int64                                                                                                     `tfsdk:"fs_group"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	RunAsUser           types.Int64                                                                                                     `tfsdk:"run_as_user"`
	SupplementalGroups  types.List                                                                                                      `tfsdk:"supplemental_groups"`
	RunAsNonRoot        types.Bool                                                                                                      `tfsdk:"run_as_non_root"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsGroup          types.Int64                                                                                                     `tfsdk:"run_as_group"`
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
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel struct {
	Resource      types.String                                                                                                       `tfsdk:"resource"`
	ContainerName types.String                                                                                                       `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                                `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                                  `tfsdk:"optional"`
}
type ValueFromInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvModel struct {
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	Value     types.String                                                                                `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                                `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateModel struct {
	Type   types.Int64                                                                                                        `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
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
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
}
type StrategyInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	Type          types.String                                                                                         `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type TolerationsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                            `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	Type   types.Int64                                                                                                                   `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type TargetInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type               types.String                                                                                                            `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                             `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
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
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                                                   `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type ResourceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsModel struct {
	Name                     types.String                                                                                                                  `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                              `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                   `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                            `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
}
type ExternalInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                                            `tfsdk:"metric_name"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchLabels      types.Map                                                                                                                      `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
}
type TargetInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
}
type ObjectInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName   types.String                                                                                                    `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchLabels      types.Map                                                                                                                    `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                               `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
}
type PodsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                        `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
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
type ResourcesInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type DeploymentInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel struct {
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	PodAnnotations           types.Map                                                                                               `tfsdk:"pod_annotations"`
	ReplicaCount             types.Int64                                                                                             `tfsdk:"replica_count"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel                   `tfsdk:"env"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel                `tfsdk:"resources"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecDeploymentModel                 `tfsdk:"affinity"`
}
type KubeSpecInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceModel struct {
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceKubeSpecModel     `tfsdk:"deployment"`
}
type InstanceInstallTemplateHelmSpecComponentsOnboardingPlaneModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsOnboardingPlaneInstanceModel `tfsdk:"kube_spec"`
}
type PlaneInstallTemplateHelmSpecComponentsOnboardingModel struct {
	Instance InstanceInstallTemplateHelmSpecComponentsOnboardingPlaneModel `tfsdk:"instance"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
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
type SeccompProfileInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	RunAsNonRoot        types.Bool                                                                                                   `tfsdk:"run_as_non_root"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	RunAsGroup          types.Int64                                                                                                  `tfsdk:"run_as_group"`
	FsGroup             types.Int64                                                                                                  `tfsdk:"fs_group"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsUser           types.Int64                                                                                                  `tfsdk:"run_as_user"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	SupplementalGroups  types.List                                                                                                   `tfsdk:"supplemental_groups"`
	FsGroupChangePolicy types.String                                                                                                 `tfsdk:"fs_group_change_policy"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                               `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
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
	Weight          types.Int64                                                                                                                                                   `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
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
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
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
	Weight          types.Int64                                                                                                                                                       `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                                         `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Namespaces    types.List                                                                                                                                                     `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                   `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
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
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                            `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
}
type PodsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                                     `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                         `tfsdk:"type"`
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
type TargetInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Value              ValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                                          `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                                         `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
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
type ResourceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel struct {
	Name                     types.String                                                                                                               `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
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
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                           `tfsdk:"match_labels"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                                `tfsdk:"type"`
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
	Type   types.Int64                                                                                                        `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                   `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                       `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
}
type ObjectInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel struct {
	AverageValue AverageValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                                 `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
}
type MetricsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecModel struct {
	Type     types.String                                                                                       `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
}
type HpaSpecInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	MaxReplicas types.Int64                                                                                   `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                                   `tfsdk:"min_replicas"`
}
type ResourcesInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	Requests types.Map `tfsdk:"requests"`
	Limits   types.Map `tfsdk:"limits"`
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
	Type   types.Int64                                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	Type          types.String                                                                                      `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type TolerationsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	ReadOnlyRootFilesystem   types.Bool                                                                                                         `tfsdk:"read_only_root_filesystem"`
	RunAsNonRoot             types.Bool                                                                                                         `tfsdk:"run_as_non_root"`
	AllowPrivilegeEscalation types.Bool                                                                                                         `tfsdk:"allow_privilege_escalation"`
	ProcMount                types.String                                                                                                       `tfsdk:"proc_mount"`
	RunAsUser                types.Int64                                                                                                        `tfsdk:"run_as_user"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsGroup               types.Int64                                                                                                        `tfsdk:"run_as_group"`
	Privileged               types.Bool                                                                                                         `tfsdk:"privileged"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
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
	Optional             types.Bool                                                                                                               `tfsdk:"optional"`
	Key                  types.String                                                                                                             `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                                  `tfsdk:"optional"`
	Key                  types.String                                                                                                                `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
}
type FieldRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type ValueFromInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvModel struct {
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel struct {
	Name      types.String                                                                             `tfsdk:"name"`
	Value     types.String                                                                             `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
}
type DeploymentInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecModel struct {
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	PodAnnotations           types.Map                                                                                            `tfsdk:"pod_annotations"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel                   `tfsdk:"env"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel                `tfsdk:"resources"`
	ReplicaCount             types.Int64                                                                                          `tfsdk:"replica_count"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecDeploymentModel                 `tfsdk:"affinity"`
}
type ValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecOverlaysPatchesModel struct {
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
}
type PatchesInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecOverlaysModel struct {
	Path  types.String                                                                           `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecOverlaysPatchesModel `tfsdk:"value"`
}
type OverlaysInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecModel struct {
	Patches    []*PatchesInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                         `tfsdk:"api_version"`
	Kind       types.String                                                                         `tfsdk:"kind"`
	Name       types.String                                                                         `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServicePortsModel struct {
	Type   types.Int64                                                                                    `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
}
type PortsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServiceModel struct {
	Name       types.String                                                                             `tfsdk:"name"`
	NodePort   types.Int64                                                                              `tfsdk:"node_port"`
	Port       types.Int64                                                                              `tfsdk:"port"`
	Protocol   types.String                                                                             `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServicePortsModel `tfsdk:"target_port"`
}
type ServiceInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecModel struct {
	Annotations types.Map                                                                         `tfsdk:"annotations"`
	Labels      types.Map                                                                         `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                                      `tfsdk:"type"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsOnboardingRepositoryKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
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
type ImagePullSecretsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel struct {
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	ReadOnlyRootFilesystem   types.Bool                                                                                                       `tfsdk:"read_only_root_filesystem"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	RunAsGroup               types.Int64                                                                                                      `tfsdk:"run_as_group"`
	AllowPrivilegeEscalation types.Bool                                                                                                       `tfsdk:"allow_privilege_escalation"`
	RunAsUser                types.Int64                                                                                                      `tfsdk:"run_as_user"`
	RunAsNonRoot             types.Bool                                                                                                       `tfsdk:"run_as_non_root"`
	ProcMount                types.String                                                                                                     `tfsdk:"proc_mount"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	Privileged               types.Bool                                                                                                       `tfsdk:"privileged"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel struct {
	Type             types.String `tfsdk:"type"`
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type SysctlsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	FsGroup             types.Int64                                                                                                `tfsdk:"fs_group"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	FsGroupChangePolicy types.String                                                                                               `tfsdk:"fs_group_change_policy"`
	RunAsUser           types.Int64                                                                                                `tfsdk:"run_as_user"`
	RunAsNonRoot        types.Bool                                                                                                 `tfsdk:"run_as_non_root"`
	SupplementalGroups  types.List                                                                                                 `tfsdk:"supplemental_groups"`
	RunAsGroup          types.Int64                                                                                                `tfsdk:"run_as_group"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
}
type ResourcesInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	Type   types.Int64                                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
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
type PodsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                                   `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	Type   types.Int64                                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
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
	Type   types.Int64                                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
}
type ResourceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel struct {
	Target                   TargetInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                             `tfsdk:"name"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
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
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
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
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                                       `tfsdk:"metric_name"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Type   types.Int64                                                                                                      `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
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
type TargetInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
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
type ObjectInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName   types.String                                                                                               `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
}
type MetricsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecModel struct {
	Resource ResourceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                                     `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
}
type HpaSpecInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                                 `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                                 `tfsdk:"max_replicas"`
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
type LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                           `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                             `tfsdk:"optional"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                                `tfsdk:"optional"`
	Key                  types.String                                                                                                              `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
}
type FieldRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type ValueFromInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvModel struct {
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	Value     types.String                                                                           `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                           `tfsdk:"name"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                             `tfsdk:"weight"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                   `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityModel struct {
	Namespaces    types.List                                                                                                                                               `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                             `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                                   `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	Namespaces    types.List                                                                                                                                                               `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                             `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                 `tfsdk:"weight"`
}
type PodAffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                       `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                   `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                 `tfsdk:"topology_key"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                                       `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                                   `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                                 `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                                     `tfsdk:"weight"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
}
type IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                   `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
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
type RollingUpdateInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                                    `tfsdk:"type"`
}
type TolerationsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel struct {
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
}
type DeploymentInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel struct {
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel                `tfsdk:"resources"`
	ReplicaCount             types.Int64                                                                                        `tfsdk:"replica_count"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	PodAnnotations           types.Map                                                                                          `tfsdk:"pod_annotations"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel                   `tfsdk:"env"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
}
type ValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecOverlaysPatchesModel struct {
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
}
type PatchesInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecOverlaysModel struct {
	Path  types.String                                                                         `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecOverlaysPatchesModel `tfsdk:"value"`
}
type OverlaysInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel struct {
	Kind       types.String                                                                       `tfsdk:"kind"`
	Name       types.String                                                                       `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                                       `tfsdk:"api_version"`
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
	NodePort   types.Int64                                                                            `tfsdk:"node_port"`
	Port       types.Int64                                                                            `tfsdk:"port"`
	Protocol   types.String                                                                           `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                           `tfsdk:"name"`
}
type ServiceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel struct {
	Labels      types.Map                                                                       `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                                    `tfsdk:"type"`
	Annotations types.Map                                                                       `tfsdk:"annotations"`
}
type KubeSpecInstallTemplateHelmSpecComponentsOnboardingOperatorModel struct {
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsOnboardingOperatorKubeSpecModel     `tfsdk:"deployment"`
}
type OperatorInstallTemplateHelmSpecComponentsOnboardingModel struct {
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsOnboardingOperatorModel `tfsdk:"kube_spec"`
}
type OnboardingInstallTemplateHelmSpecComponentsModel struct {
	Plane      PlaneInstallTemplateHelmSpecComponentsOnboardingModel      `tfsdk:"plane"`
	Repository RepositoryInstallTemplateHelmSpecComponentsOnboardingModel `tfsdk:"repository"`
	Operator   OperatorInstallTemplateHelmSpecComponentsOnboardingModel   `tfsdk:"operator"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsSatelliteKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsSatelliteKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	AllowPrivilegeEscalation types.Bool                                                                                              `tfsdk:"allow_privilege_escalation"`
	RunAsGroup               types.Int64                                                                                             `tfsdk:"run_as_group"`
	Privileged               types.Bool                                                                                              `tfsdk:"privileged"`
	ProcMount                types.String                                                                                            `tfsdk:"proc_mount"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	RunAsNonRoot             types.Bool                                                                                              `tfsdk:"run_as_non_root"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	ReadOnlyRootFilesystem   types.Bool                                                                                              `tfsdk:"read_only_root_filesystem"`
	RunAsUser                types.Int64                                                                                             `tfsdk:"run_as_user"`
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
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                           `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                          `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                                    `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                      `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                        `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                          `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                    `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                              `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                          `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                        `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                                            `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                              `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	TopologyKey   types.String                                                                                                                                        `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                          `tfsdk:"namespaces"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
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
type AffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
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
	Divisor       DivisorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                         `tfsdk:"resource"`
	ContainerName types.String                                                                                         `tfsdk:"container_name"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                  `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                    `tfsdk:"optional"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                       `tfsdk:"optional"`
	Key                  types.String                                                                                                     `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
}
type ValueFromInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvModel struct {
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	Value     types.String                                                                  `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                                  `tfsdk:"name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type SysctlsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel struct {
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	FsGroupChangePolicy types.String                                                                                      `tfsdk:"fs_group_change_policy"`
	FsGroup             types.Int64                                                                                       `tfsdk:"fs_group"`
	RunAsGroup          types.Int64                                                                                       `tfsdk:"run_as_group"`
	RunAsUser           types.Int64                                                                                       `tfsdk:"run_as_user"`
	RunAsNonRoot        types.Bool                                                                                        `tfsdk:"run_as_non_root"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SupplementalGroups  types.List                                                                                        `tfsdk:"supplemental_groups"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
}
type ResourcesInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	Requests types.Map `tfsdk:"requests"`
	Limits   types.Map `tfsdk:"limits"`
}
type TolerationsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchLabels      types.Map                                                                                                      `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
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
	Selector           SelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                          `tfsdk:"metric_name"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                     `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                     `tfsdk:"type"`
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
type TargetInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                              `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                               `tfsdk:"average_utilization"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
}
type ResourceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                    `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
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
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	Type   types.Int64                                                                                              `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
}
type ExternalInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                              `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                            `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Type   types.Int64                                                                                             `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                        `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
}
type ObjectInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                      `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
}
type MetricsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecModel struct {
	Resource ResourceInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                            `tfsdk:"type"`
	External ExternalInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
}
type HpaSpecInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel struct {
	MinReplicas types.Int64                                                                        `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                        `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
}
type DeploymentInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel struct {
	Env                      []*EnvInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel                   `tfsdk:"env"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	ReplicaCount             types.Int64                                                                               `tfsdk:"replica_count"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	PodAnnotations           types.Map                                                                                 `tfsdk:"pod_annotations"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsSatelliteKubeSpecDeploymentModel                `tfsdk:"resources"`
}
type ValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecOverlaysPatchesModel struct {
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
}
type PatchesInstallTemplateHelmSpecComponentsSatelliteKubeSpecOverlaysModel struct {
	Path  types.String                                                                `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsSatelliteKubeSpecOverlaysPatchesModel `tfsdk:"value"`
}
type OverlaysInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel struct {
	Kind       types.String                                                              `tfsdk:"kind"`
	Name       types.String                                                              `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsSatelliteKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                              `tfsdk:"api_version"`
}
type StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsSatelliteKubeSpecServicePortsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsSatelliteKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsSatelliteKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                         `tfsdk:"type"`
}
type PortsInstallTemplateHelmSpecComponentsSatelliteKubeSpecServiceModel struct {
	TargetPort TargetPortInstallTemplateHelmSpecComponentsSatelliteKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                  `tfsdk:"name"`
	NodePort   types.Int64                                                                   `tfsdk:"node_port"`
	Port       types.Int64                                                                   `tfsdk:"port"`
	Protocol   types.String                                                                  `tfsdk:"protocol"`
}
type ServiceInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel struct {
	Labels      types.Map                                                              `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsSatelliteKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                           `tfsdk:"type"`
	Annotations types.Map                                                              `tfsdk:"annotations"`
}
type KubeSpecInstallTemplateHelmSpecComponentsSatelliteModel struct {
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsSatelliteKubeSpecModel        `tfsdk:"service"`
}
type SatelliteInstallTemplateHelmSpecComponentsModel struct {
	Enabled  types.Bool                                              `tfsdk:"enabled"`
	KubeSpec KubeSpecInstallTemplateHelmSpecComponentsSatelliteModel `tfsdk:"kube_spec"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsCollectorKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsCollectorKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsCollectorKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel struct {
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	RunAsNonRoot             types.Bool                                                                                              `tfsdk:"run_as_non_root"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsUser                types.Int64                                                                                             `tfsdk:"run_as_user"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	Privileged               types.Bool                                                                                              `tfsdk:"privileged"`
	ProcMount                types.String                                                                                            `tfsdk:"proc_mount"`
	RunAsGroup               types.Int64                                                                                             `tfsdk:"run_as_group"`
	AllowPrivilegeEscalation types.Bool                                                                                              `tfsdk:"allow_privilege_escalation"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	ReadOnlyRootFilesystem   types.Bool                                                                                              `tfsdk:"read_only_root_filesystem"`
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
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                             `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectAverageValueModel `tfsdk:"int_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type SelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchLabels      types.Map                                                                                                        `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
}
type TargetInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type ObjectInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                      `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                 `tfsdk:"type"`
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
type PodsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                          `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                     `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                              `tfsdk:"type"`
}
type TargetInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type               types.String                                                                                              `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                               `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                           `tfsdk:"type"`
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
	Name                     types.String                                                                                                    `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MetricSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	MatchLabels      types.Map                                                                                                                `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalMetricSelectorModel `tfsdk:"match_expressions"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
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
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                              `tfsdk:"type"`
}
type ExternalInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                              `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
}
type MetricsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecModel struct {
	External ExternalInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                            `tfsdk:"type"`
}
type HpaSpecInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                        `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                        `tfsdk:"max_replicas"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel struct {
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                       `tfsdk:"optional"`
	Key                  types.String                                                                                                     `tfsdk:"key"`
}
type FieldRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel struct {
	Divisor       DivisorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
	Resource      types.String                                                                                         `tfsdk:"resource"`
	ContainerName types.String                                                                                         `tfsdk:"container_name"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel struct {
	Key                  types.String                                                                                                  `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                    `tfsdk:"optional"`
}
type ValueFromInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvModel struct {
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	Name      types.String                                                                  `tfsdk:"name"`
	Value     types.String                                                                  `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
}
type ResourcesInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	Requests types.Map `tfsdk:"requests"`
	Limits   types.Map `tfsdk:"limits"`
}
type TolerationsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SysctlsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	RunAsUser           types.Int64                                                                                       `tfsdk:"run_as_user"`
	FsGroupChangePolicy types.String                                                                                      `tfsdk:"fs_group_change_policy"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	RunAsNonRoot        types.Bool                                                                                        `tfsdk:"run_as_non_root"`
	SupplementalGroups  types.List                                                                                        `tfsdk:"supplemental_groups"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	FsGroup             types.Int64                                                                                       `tfsdk:"fs_group"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	RunAsGroup          types.Int64                                                                                       `tfsdk:"run_as_group"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type PreferenceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                    `tfsdk:"weight"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
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
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                                          `tfsdk:"match_labels"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                    `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityModel struct {
	Weight          types.Int64                                                                                                                                        `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type PodAffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                              `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	TopologyKey   types.String                                                                                                                                                        `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                          `tfsdk:"namespaces"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                            `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                              `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	Namespaces    types.List                                                                                                                                          `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                        `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
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
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyModel struct {
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
}
type StrategyInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel struct {
	Type          types.String                                                                           `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type DeploymentInstallTemplateHelmSpecComponentsCollectorKubeSpecModel struct {
	ReplicaCount             types.Int64                                                                               `tfsdk:"replica_count"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel                   `tfsdk:"env"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel                `tfsdk:"resources"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	PodAnnotations           types.Map                                                                                 `tfsdk:"pod_annotations"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsCollectorKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
}
type ValueInstallTemplateHelmSpecComponentsCollectorKubeSpecOverlaysPatchesModel struct {
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
}
type PatchesInstallTemplateHelmSpecComponentsCollectorKubeSpecOverlaysModel struct {
	Path  types.String                                                                `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsCollectorKubeSpecOverlaysPatchesModel `tfsdk:"value"`
}
type OverlaysInstallTemplateHelmSpecComponentsCollectorKubeSpecModel struct {
	ApiVersion types.String                                                              `tfsdk:"api_version"`
	Kind       types.String                                                              `tfsdk:"kind"`
	Name       types.String                                                              `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsCollectorKubeSpecOverlaysModel `tfsdk:"patches"`
}
type StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsCollectorKubeSpecServicePortsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsCollectorKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsCollectorKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
}
type PortsInstallTemplateHelmSpecComponentsCollectorKubeSpecServiceModel struct {
	Protocol   types.String                                                                  `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsCollectorKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                                  `tfsdk:"name"`
	NodePort   types.Int64                                                                   `tfsdk:"node_port"`
	Port       types.Int64                                                                   `tfsdk:"port"`
}
type ServiceInstallTemplateHelmSpecComponentsCollectorKubeSpecModel struct {
	Labels      types.Map                                                              `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsCollectorKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                           `tfsdk:"type"`
	Annotations types.Map                                                              `tfsdk:"annotations"`
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
type ValueInstallTemplateHelmSpecComponentsIstioPilotOverlaysPatchesModel struct {
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
}
type PatchesInstallTemplateHelmSpecComponentsIstioPilotOverlaysModel struct {
	Value ValueInstallTemplateHelmSpecComponentsIstioPilotOverlaysPatchesModel `tfsdk:"value"`
	Path  types.String                                                         `tfsdk:"path"`
}
type PilotOverlaysInstallTemplateHelmSpecComponentsIstioModel struct {
	Name       types.String                                                       `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsIstioPilotOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                       `tfsdk:"api_version"`
	Kind       types.String                                                       `tfsdk:"kind"`
}
type ValueInstallTemplateHelmSpecComponentsIstioCniOverlaysPatchesModel struct {
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
}
type PatchesInstallTemplateHelmSpecComponentsIstioCniOverlaysModel struct {
	Path  types.String                                                       `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsIstioCniOverlaysPatchesModel `tfsdk:"value"`
}
type CniOverlaysInstallTemplateHelmSpecComponentsIstioModel struct {
	Patches    []*PatchesInstallTemplateHelmSpecComponentsIstioCniOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                     `tfsdk:"api_version"`
	Kind       types.String                                                     `tfsdk:"kind"`
	Name       types.String                                                     `tfsdk:"name"`
}
type MaxWorkloadCertTtlInstallTemplateHelmSpecComponentsIstioModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type ValueInstallTemplateHelmSpecComponentsIstioBaseOverlaysPatchesModel struct {
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
}
type PatchesInstallTemplateHelmSpecComponentsIstioBaseOverlaysModel struct {
	Path  types.String                                                        `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsIstioBaseOverlaysPatchesModel `tfsdk:"value"`
}
type BaseOverlaysInstallTemplateHelmSpecComponentsIstioModel struct {
	ApiVersion types.String                                                      `tfsdk:"api_version"`
	Kind       types.String                                                      `tfsdk:"kind"`
	Name       types.String                                                      `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsIstioBaseOverlaysModel `tfsdk:"patches"`
}
type DefaultWorkloadCertTtlInstallTemplateHelmSpecComponentsIstioModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type ValueInstallTemplateHelmSpecComponentsIstioKubeSpecOverlaysPatchesModel struct {
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
}
type PatchesInstallTemplateHelmSpecComponentsIstioKubeSpecOverlaysModel struct {
	Value ValueInstallTemplateHelmSpecComponentsIstioKubeSpecOverlaysPatchesModel `tfsdk:"value"`
	Path  types.String                                                            `tfsdk:"path"`
}
type OverlaysInstallTemplateHelmSpecComponentsIstioKubeSpecModel struct {
	ApiVersion types.String                                                          `tfsdk:"api_version"`
	Kind       types.String                                                          `tfsdk:"kind"`
	Name       types.String                                                          `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsIstioKubeSpecOverlaysModel `tfsdk:"patches"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsIstioKubeSpecServicePortsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
}
type PortsInstallTemplateHelmSpecComponentsIstioKubeSpecServiceModel struct {
	Port       types.Int64                                                               `tfsdk:"port"`
	Protocol   types.String                                                              `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsIstioKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                              `tfsdk:"name"`
	NodePort   types.Int64                                                               `tfsdk:"node_port"`
}
type ServiceInstallTemplateHelmSpecComponentsIstioKubeSpecModel struct {
	Annotations types.Map                                                          `tfsdk:"annotations"`
	Labels      types.Map                                                          `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsIstioKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                       `tfsdk:"type"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsIstioKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsIstioKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsIstioKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type CniInstallTemplateHelmSpecComponentsIstioKubeSpecModel struct {
	Chained                types.Bool   `tfsdk:"chained"`
	ClusterRole            types.String `tfsdk:"cluster_role"`
	ConfigurationDirectory types.String `tfsdk:"configuration_directory"`
	ConfigurationFileName  types.String `tfsdk:"configuration_file_name"`
	Revision               types.String `tfsdk:"revision"`
	BinaryDirectory        types.String `tfsdk:"binary_directory"`
}
type TolerationsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel struct {
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                                   `tfsdk:"optional"`
	Key                  types.String                                                                                                 `tfsdk:"key"`
}
type FieldRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	FieldPath  types.String `tfsdk:"field_path"`
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
type ValueFromInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvModel struct {
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
}
type EnvInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	Name      types.String                                                              `tfsdk:"name"`
	Value     types.String                                                              `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SysctlsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	Value types.String `tfsdk:"value"`
	Name  types.String `tfsdk:"name"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	RunAsGroup          types.Int64                                                                                   `tfsdk:"run_as_group"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	FsGroup             types.Int64                                                                                   `tfsdk:"fs_group"`
	RunAsUser           types.Int64                                                                                   `tfsdk:"run_as_user"`
	RunAsNonRoot        types.Bool                                                                                    `tfsdk:"run_as_non_root"`
	SupplementalGroups  types.List                                                                                    `tfsdk:"supplemental_groups"`
	FsGroupChangePolicy types.String                                                                                  `tfsdk:"fs_group_change_policy"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	Type   types.Int64                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
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
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                 `tfsdk:"type"`
}
type ExternalInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                          `tfsdk:"metric_name"`
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                    `tfsdk:"match_labels"`
}
type TargetInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                        `tfsdk:"type"`
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
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                  `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
	Target       TargetInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                  `tfsdk:"match_labels"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                             `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
}
type PodsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricName         types.String                                                                                      `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
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
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type ValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	Type   types.Int64                                                                                          `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetValueModel `tfsdk:"str_val"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                 `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type TargetInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Value              ValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                           `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
	Type               types.String                                                                                          `tfsdk:"type"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type   types.Int64                                                                                                       `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
}
type ResourceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel struct {
	Target                   TargetInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
	Name                     types.String                                                                                                `tfsdk:"name"`
}
type MetricsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecModel struct {
	External ExternalInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                        `tfsdk:"type"`
}
type HpaSpecInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	MinReplicas types.Int64                                                                    `tfsdk:"min_replicas"`
	MaxReplicas types.Int64                                                                    `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel struct {
	Add  types.List `tfsdk:"add"`
	Drop types.List `tfsdk:"drop"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	ProcMount                types.String                                                                                        `tfsdk:"proc_mount"`
	RunAsUser                types.Int64                                                                                         `tfsdk:"run_as_user"`
	Privileged               types.Bool                                                                                          `tfsdk:"privileged"`
	ReadOnlyRootFilesystem   types.Bool                                                                                          `tfsdk:"read_only_root_filesystem"`
	RunAsNonRoot             types.Bool                                                                                          `tfsdk:"run_as_non_root"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
	AllowPrivilegeEscalation types.Bool                                                                                          `tfsdk:"allow_privilege_escalation"`
	RunAsGroup               types.Int64                                                                                         `tfsdk:"run_as_group"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
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
type StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                            `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
	Type          types.String                                                                       `tfsdk:"type"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                          `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                      `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                    `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	PodAffinityTerm PodAffinityTermInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
	Weight          types.Int64                                                                                                                                        `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                                          `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel struct {
	TopologyKey   types.String                                                                                                                                    `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                      `tfsdk:"namespaces"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type PreferenceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Preference PreferenceInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
	Weight     types.Int64                                                                                                                                `tfsdk:"weight"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
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
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                                      `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityModel struct {
	Namespaces    types.List                                                                                                                                  `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
}
type ResourcesInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type DeploymentInstallTemplateHelmSpecComponentsIstioKubeSpecModel struct {
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	ReplicaCount             types.Int64                                                                           `tfsdk:"replica_count"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	PodAnnotations           types.Map                                                                             `tfsdk:"pod_annotations"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel                `tfsdk:"resources"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel                   `tfsdk:"env"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsIstioKubeSpecDeploymentModel                 `tfsdk:"affinity"`
}
type KubeSpecInstallTemplateHelmSpecComponentsIstioModel struct {
	Service        ServiceInstallTemplateHelmSpecComponentsIstioKubeSpecModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsIstioKubeSpecModel `tfsdk:"service_account"`
	Cni            CniInstallTemplateHelmSpecComponentsIstioKubeSpecModel            `tfsdk:"cni"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsIstioKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsIstioKubeSpecModel    `tfsdk:"overlays"`
}
type IstioInstallTemplateHelmSpecComponentsModel struct {
	TraceSamplingRate      types.Float64                                                     `tfsdk:"trace_sampling_rate"`
	CniOverlays            []*CniOverlaysInstallTemplateHelmSpecComponentsIstioModel         `tfsdk:"cni_overlays"`
	MaxWorkloadCertTtl     MaxWorkloadCertTtlInstallTemplateHelmSpecComponentsIstioModel     `tfsdk:"max_workload_cert_ttl"`
	BaseOverlays           []*BaseOverlaysInstallTemplateHelmSpecComponentsIstioModel        `tfsdk:"base_overlays"`
	DefaultWorkloadCertTtl DefaultWorkloadCertTtlInstallTemplateHelmSpecComponentsIstioModel `tfsdk:"default_workload_cert_ttl"`
	PilotOverlays          []*PilotOverlaysInstallTemplateHelmSpecComponentsIstioModel       `tfsdk:"pilot_overlays"`
	TrustDomain            types.String                                                      `tfsdk:"trust_domain"`
	KubeSpec               KubeSpecInstallTemplateHelmSpecComponentsIstioModel               `tfsdk:"kube_spec"`
	TsbVersion             types.String                                                      `tfsdk:"tsb_version"`
}
type ImagePullSecretsInstallTemplateHelmSpecComponentsNgacKubeSpecServiceAccountModel struct {
	Name types.String `tfsdk:"name"`
}
type ServiceAccountInstallTemplateHelmSpecComponentsNgacKubeSpecModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmSpecComponentsNgacKubeSpecServiceAccountModel `tfsdk:"image_pull_secrets"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                                               `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchFieldsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type NodeSelectorTermsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
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
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                 `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                               `tfsdk:"topology_key"`
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
	MatchLabels      types.Map                                                                                                                                                     `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityModel struct {
	Namespaces    types.List                                                                                                                                 `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                               `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
}
type PodAffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel struct {
	MatchLabels      types.Map                                                                                                                                                                         `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel `tfsdk:"match_expressions"`
}
type PodAffinityTermInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                                     `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                                   `tfsdk:"topology_key"`
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
	TopologyKey   types.String                                                                                                                                   `tfsdk:"topology_key"`
	LabelSelector LabelSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                     `tfsdk:"namespaces"`
}
type PodAntiAffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityModel struct {
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel struct {
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
	Optional             types.Bool                                                                                               `tfsdk:"optional"`
	Key                  types.String                                                                                             `tfsdk:"key"`
}
type LocalObjectReferenceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                                  `tfsdk:"optional"`
	Key                  types.String                                                                                                `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
}
type FieldRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel struct {
	FieldPath  types.String `tfsdk:"field_path"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type DivisorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromResourceFieldRefModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                           `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
}
type ResourceFieldRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel struct {
	Resource      types.String                                                                                    `tfsdk:"resource"`
	ContainerName types.String                                                                                    `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
}
type ValueFromInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvModel struct {
	SecretKeyRef     SecretKeyRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
	FieldRef         FieldRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
}
type EnvInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	Value     types.String                                                             `tfsdk:"value"`
	ValueFrom ValueFromInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentEnvModel `tfsdk:"value_from"`
	Name      types.String                                                             `tfsdk:"name"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SysctlsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel struct {
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type PodSecurityContextInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	Sysctls             []*SysctlsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel     `tfsdk:"sysctls"`
	RunAsNonRoot        types.Bool                                                                                   `tfsdk:"run_as_non_root"`
	RunAsUser           types.Int64                                                                                  `tfsdk:"run_as_user"`
	FsGroupChangePolicy types.String                                                                                 `tfsdk:"fs_group_change_policy"`
	SeccompProfile      SeccompProfileInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel `tfsdk:"seccomp_profile"`
	FsGroup             types.Int64                                                                                  `tfsdk:"fs_group"`
	RunAsGroup          types.Int64                                                                                  `tfsdk:"run_as_group"`
	SupplementalGroups  types.List                                                                                   `tfsdk:"supplemental_groups"`
	SeLinuxOptions      SeLinuxOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel `tfsdk:"se_linux_options"`
	WindowsOptions      WindowsOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentPodSecurityContextModel `tfsdk:"windows_options"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                           `tfsdk:"type"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type MaxSurgeInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                     `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
}
type RollingUpdateInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyModel struct {
	MaxUnavailable MaxUnavailableInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
	MaxSurge       MaxSurgeInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
}
type StrategyInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	Type          types.String                                                                      `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type ResourcesInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	Limits   types.Map `tfsdk:"limits"`
	Requests types.Map `tfsdk:"requests"`
}
type TolerationsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	Effect            types.String `tfsdk:"effect"`
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
}
type WindowsOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel struct {
	GmsaCredentialSpec     types.String `tfsdk:"gmsa_credential_spec"`
	GmsaCredentialSpecName types.String `tfsdk:"gmsa_credential_spec_name"`
	RunAsUserName          types.String `tfsdk:"run_as_user_name"`
}
type SeccompProfileInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel struct {
	LocalhostProfile types.String `tfsdk:"localhost_profile"`
	Type             types.String `tfsdk:"type"`
}
type CapabilitiesInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel struct {
	Drop types.List `tfsdk:"drop"`
	Add  types.List `tfsdk:"add"`
}
type SeLinuxOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel struct {
	Role  types.String `tfsdk:"role"`
	Type  types.String `tfsdk:"type"`
	User  types.String `tfsdk:"user"`
	Level types.String `tfsdk:"level"`
}
type ContainerSecurityContextInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	SeLinuxOptions           SeLinuxOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"se_linux_options"`
	RunAsGroup               types.Int64                                                                                        `tfsdk:"run_as_group"`
	ProcMount                types.String                                                                                       `tfsdk:"proc_mount"`
	ReadOnlyRootFilesystem   types.Bool                                                                                         `tfsdk:"read_only_root_filesystem"`
	Privileged               types.Bool                                                                                         `tfsdk:"privileged"`
	RunAsNonRoot             types.Bool                                                                                         `tfsdk:"run_as_non_root"`
	RunAsUser                types.Int64                                                                                        `tfsdk:"run_as_user"`
	Capabilities             CapabilitiesInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel   `tfsdk:"capabilities"`
	AllowPrivilegeEscalation types.Bool                                                                                         `tfsdk:"allow_privilege_escalation"`
	WindowsOptions           WindowsOptionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"windows_options"`
	SeccompProfile           SeccompProfileInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentContainerSecurityContextModel `tfsdk:"seccomp_profile"`
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
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetAverageValueModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalModel struct {
	Type   types.Int64                                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalTargetValueModel `tfsdk:"str_val"`
}
type ExternalInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel struct {
	MetricSelector     MetricSelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalModel     `tfsdk:"metric_selector"`
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalModel `tfsdk:"target_average_value"`
	TargetValue        TargetValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsExternalModel        `tfsdk:"target_value"`
	MetricName         types.String                                                                                         `tfsdk:"metric_name"`
}
type TargetInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	Kind       types.String `tfsdk:"kind"`
	Name       types.String `tfsdk:"name"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                       `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectTargetValueModel `tfsdk:"int_val"`
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
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type SelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel struct {
	MatchLabels      types.Map                                                                                                   `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectSelectorModel `tfsdk:"match_expressions"`
}
type ObjectInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel struct {
	Target       TargetInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel       `tfsdk:"target"`
	TargetValue  TargetValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel  `tfsdk:"target_value"`
	AverageValue AverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel `tfsdk:"average_value"`
	MetricName   types.String                                                                                 `tfsdk:"metric_name"`
	Selector     SelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsObjectModel     `tfsdk:"selector"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetAverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                            `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsTargetAverageValueModel `tfsdk:"int_val"`
}
type MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type SelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                 `tfsdk:"match_labels"`
}
type PodsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel struct {
	TargetAverageValue TargetAverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsModel `tfsdk:"target_average_value"`
	MetricName         types.String                                                                                     `tfsdk:"metric_name"`
	Selector           SelectorInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsPodsModel           `tfsdk:"selector"`
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
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type AverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageValueModel `tfsdk:"int_val"`
}
type TargetInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	Type               types.String                                                                                         `tfsdk:"type"`
	Value              ValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetModel        `tfsdk:"value"`
	AverageUtilization types.Int64                                                                                          `tfsdk:"average_utilization"`
	AverageValue       AverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetModel `tfsdk:"average_value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.String `tfsdk:"value"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type TargetAverageUtilizationInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceModel struct {
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceTargetAverageUtilizationModel `tfsdk:"str_val"`
	Type   types.Int64                                                                                                      `tfsdk:"type"`
}
type ResourceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel struct {
	Name                     types.String                                                                                               `tfsdk:"name"`
	Target                   TargetInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceModel                   `tfsdk:"target"`
	TargetAverageUtilization TargetAverageUtilizationInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceModel `tfsdk:"target_average_utilization"`
	TargetAverageValue       TargetAverageValueInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsResourceModel       `tfsdk:"target_average_value"`
}
type MetricsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecModel struct {
	External ExternalInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"external"`
	Object   ObjectInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel   `tfsdk:"object"`
	Pods     PodsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel     `tfsdk:"pods"`
	Resource ResourceInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecMetricsModel `tfsdk:"resource"`
	Type     types.String                                                                       `tfsdk:"type"`
}
type HpaSpecInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel struct {
	MaxReplicas types.Int64                                                                   `tfsdk:"max_replicas"`
	Metrics     []*MetricsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentHpaSpecModel `tfsdk:"metrics"`
	MinReplicas types.Int64                                                                   `tfsdk:"min_replicas"`
}
type DeploymentInstallTemplateHelmSpecComponentsNgacKubeSpecModel struct {
	ReplicaCount             types.Int64                                                                          `tfsdk:"replica_count"`
	Strategy                 StrategyInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel                 `tfsdk:"strategy"`
	PodAnnotations           types.Map                                                                            `tfsdk:"pod_annotations"`
	Affinity                 AffinityInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel                 `tfsdk:"affinity"`
	Env                      []*EnvInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel                   `tfsdk:"env"`
	PodSecurityContext       PodSecurityContextInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel       `tfsdk:"pod_security_context"`
	Resources                ResourcesInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel                `tfsdk:"resources"`
	Tolerations              []*TolerationsInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel           `tfsdk:"tolerations"`
	ContainerSecurityContext ContainerSecurityContextInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel `tfsdk:"container_security_context"`
	HpaSpec                  HpaSpecInstallTemplateHelmSpecComponentsNgacKubeSpecDeploymentModel                  `tfsdk:"hpa_spec"`
}
type ValueInstallTemplateHelmSpecComponentsNgacKubeSpecOverlaysPatchesModel struct {
	NullValue   types.String  `tfsdk:"null_value"`
	NumberValue types.Float64 `tfsdk:"number_value"`
	StringValue types.String  `tfsdk:"string_value"`
	StructValue types.Map     `tfsdk:"struct_value"`
	BoolValue   types.Bool    `tfsdk:"bool_value"`
	ListValue   types.Map     `tfsdk:"list_value"`
}
type PatchesInstallTemplateHelmSpecComponentsNgacKubeSpecOverlaysModel struct {
	Path  types.String                                                           `tfsdk:"path"`
	Value ValueInstallTemplateHelmSpecComponentsNgacKubeSpecOverlaysPatchesModel `tfsdk:"value"`
}
type OverlaysInstallTemplateHelmSpecComponentsNgacKubeSpecModel struct {
	Name       types.String                                                         `tfsdk:"name"`
	Patches    []*PatchesInstallTemplateHelmSpecComponentsNgacKubeSpecOverlaysModel `tfsdk:"patches"`
	ApiVersion types.String                                                         `tfsdk:"api_version"`
	Kind       types.String                                                         `tfsdk:"kind"`
}
type IntValInstallTemplateHelmSpecComponentsNgacKubeSpecServicePortsTargetPortModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmSpecComponentsNgacKubeSpecServicePortsTargetPortModel struct {
	Value types.String `tfsdk:"value"`
}
type TargetPortInstallTemplateHelmSpecComponentsNgacKubeSpecServicePortsModel struct {
	StrVal StrValInstallTemplateHelmSpecComponentsNgacKubeSpecServicePortsTargetPortModel `tfsdk:"str_val"`
	Type   types.Int64                                                                    `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmSpecComponentsNgacKubeSpecServicePortsTargetPortModel `tfsdk:"int_val"`
}
type PortsInstallTemplateHelmSpecComponentsNgacKubeSpecServiceModel struct {
	NodePort   types.Int64                                                              `tfsdk:"node_port"`
	Port       types.Int64                                                              `tfsdk:"port"`
	Protocol   types.String                                                             `tfsdk:"protocol"`
	TargetPort TargetPortInstallTemplateHelmSpecComponentsNgacKubeSpecServicePortsModel `tfsdk:"target_port"`
	Name       types.String                                                             `tfsdk:"name"`
}
type ServiceInstallTemplateHelmSpecComponentsNgacKubeSpecModel struct {
	Annotations types.Map                                                         `tfsdk:"annotations"`
	Labels      types.Map                                                         `tfsdk:"labels"`
	Ports       []*PortsInstallTemplateHelmSpecComponentsNgacKubeSpecServiceModel `tfsdk:"ports"`
	Type        types.String                                                      `tfsdk:"type"`
}
type KubeSpecInstallTemplateHelmSpecComponentsNgacModel struct {
	ServiceAccount ServiceAccountInstallTemplateHelmSpecComponentsNgacKubeSpecModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmSpecComponentsNgacKubeSpecModel     `tfsdk:"deployment"`
	Overlays       []*OverlaysInstallTemplateHelmSpecComponentsNgacKubeSpecModel    `tfsdk:"overlays"`
	Service        ServiceInstallTemplateHelmSpecComponentsNgacKubeSpecModel        `tfsdk:"service"`
}
type NgacInstallTemplateHelmSpecComponentsModel struct {
	KubeSpec  KubeSpecInstallTemplateHelmSpecComponentsNgacModel `tfsdk:"kube_spec"`
	LogLevels types.Map                                          `tfsdk:"log_levels"`
	Enabled   types.Bool                                         `tfsdk:"enabled"`
}
type ComponentsInstallTemplateHelmSpecModel struct {
	HpaAdapter           HpaAdapterInstallTemplateHelmSpecComponentsModel           `tfsdk:"hpa_adapter"`
	RateLimitServer      RateLimitServerInstallTemplateHelmSpecComponentsModel      `tfsdk:"rate_limit_server"`
	Collector            CollectorInstallTemplateHelmSpecComponentsModel            `tfsdk:"collector"`
	Gitops               GitopsInstallTemplateHelmSpecComponentsModel               `tfsdk:"gitops"`
	Xcp                  XcpInstallTemplateHelmSpecComponentsModel                  `tfsdk:"xcp"`
	Oap                  OapInstallTemplateHelmSpecComponentsModel                  `tfsdk:"oap"`
	InternalCertProvider InternalCertProviderInstallTemplateHelmSpecComponentsModel `tfsdk:"internal_cert_provider"`
	DefaultKubeSpec      DefaultKubeSpecInstallTemplateHelmSpecComponentsModel      `tfsdk:"default_kube_spec"`
	Onboarding           OnboardingInstallTemplateHelmSpecComponentsModel           `tfsdk:"onboarding"`
	Satellite            SatelliteInstallTemplateHelmSpecComponentsModel            `tfsdk:"satellite"`
	Istio                IstioInstallTemplateHelmSpecComponentsModel                `tfsdk:"istio"`
	Ngac                 NgacInstallTemplateHelmSpecComponentsModel                 `tfsdk:"ngac"`
}
type ImagePullSecretsInstallTemplateHelmSpecModel struct {
	Name types.String `tfsdk:"name"`
}
type ManagementPlaneInstallTemplateHelmSpecModel struct {
	ClusterName types.String `tfsdk:"cluster_name"`
	Host        types.String `tfsdk:"host"`
	Port        types.Int64  `tfsdk:"port"`
	SelfSigned  types.Bool   `tfsdk:"self_signed"`
}
type EndpointInstallTemplateHelmSpecMeshExpansionOnboardingModel struct {
	Hosts      types.List   `tfsdk:"hosts"`
	SecretName types.String `tfsdk:"secret_name"`
}
type LocalRepositoryInstallTemplateHelmSpecMeshExpansionOnboardingModel struct{}
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
type PropagationDelayInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsDeregistrationModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type DeregistrationInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsModel struct {
	PropagationDelay PropagationDelayInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsDeregistrationModel `tfsdk:"propagation_delay"`
}
type AttributesInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationJwtIssuersTokenFieldsModel struct {
	JsonPath types.String `tfsdk:"json_path"`
}
type TokenFieldsInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationJwtIssuersModel struct {
	Attributes AttributesInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationJwtIssuersTokenFieldsModel `tfsdk:"attributes"`
}
type IssuersInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationJwtModel struct {
	JwksUri     types.String                                                                                    `tfsdk:"jwks_uri"`
	ShortName   types.String                                                                                    `tfsdk:"short_name"`
	TokenFields TokenFieldsInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationJwtIssuersModel `tfsdk:"token_fields"`
	Issuer      types.String                                                                                    `tfsdk:"issuer"`
	Jwks        types.String                                                                                    `tfsdk:"jwks"`
}
type JwtInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationModel struct {
	Issuers []*IssuersInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationJwtModel `tfsdk:"issuers"`
}
type AuthenticationInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsModel struct {
	Jwt JwtInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsAuthenticationModel `tfsdk:"jwt"`
}
type WorkloadsInstallTemplateHelmSpecMeshExpansionOnboardingModel struct {
	Deregistration DeregistrationInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsModel `tfsdk:"deregistration"`
	Authentication AuthenticationInstallTemplateHelmSpecMeshExpansionOnboardingWorkloadsModel `tfsdk:"authentication"`
}
type OnboardingInstallTemplateHelmSpecMeshExpansionModel struct {
	Uid             types.String                                                       `tfsdk:"uid"`
	Workloads       WorkloadsInstallTemplateHelmSpecMeshExpansionOnboardingModel       `tfsdk:"workloads"`
	Endpoint        EndpointInstallTemplateHelmSpecMeshExpansionOnboardingModel        `tfsdk:"endpoint"`
	LocalRepository LocalRepositoryInstallTemplateHelmSpecMeshExpansionOnboardingModel `tfsdk:"local_repository"`
	TokenIssuer     TokenIssuerInstallTemplateHelmSpecMeshExpansionOnboardingModel     `tfsdk:"token_issuer"`
}
type CustomGatewayInstallTemplateHelmSpecMeshExpansionModel struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}
type MeshExpansionInstallTemplateHelmSpecModel struct {
	Onboarding    OnboardingInstallTemplateHelmSpecMeshExpansionModel    `tfsdk:"onboarding"`
	CustomGateway CustomGatewayInstallTemplateHelmSpecMeshExpansionModel `tfsdk:"custom_gateway"`
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
	Host       types.String `tfsdk:"host"`
	Port       types.Int64  `tfsdk:"port"`
	Protocol   types.String `tfsdk:"protocol"`
	SelfSigned types.Bool   `tfsdk:"self_signed"`
	Version    types.Int64  `tfsdk:"version"`
}
type TelemetryStoreInstallTemplateHelmSpecModel struct {
	Elastic ElasticInstallTemplateHelmSpecTelemetryStoreModel `tfsdk:"elastic"`
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
	Registry types.String `tfsdk:"registry"`
	Tag      types.String `tfsdk:"tag"`
}
type FieldRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel struct {
	FieldPath  types.String `tfsdk:"field_path"`
	ApiVersion types.String `tfsdk:"api_version"`
}
type IntValInstallTemplateHelmOperatorDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmOperatorDeploymentEnvValueFromResourceFieldRefDivisorModel struct {
	Value types.String `tfsdk:"value"`
}
type DivisorInstallTemplateHelmOperatorDeploymentEnvValueFromResourceFieldRefModel struct {
	StrVal StrValInstallTemplateHelmOperatorDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"str_val"`
	Type   types.Int64                                                                         `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmOperatorDeploymentEnvValueFromResourceFieldRefDivisorModel `tfsdk:"int_val"`
}
type ResourceFieldRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel struct {
	Resource      types.String                                                                  `tfsdk:"resource"`
	ContainerName types.String                                                                  `tfsdk:"container_name"`
	Divisor       DivisorInstallTemplateHelmOperatorDeploymentEnvValueFromResourceFieldRefModel `tfsdk:"divisor"`
}
type LocalObjectReferenceInstallTemplateHelmOperatorDeploymentEnvValueFromSecretKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type SecretKeyRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                             `tfsdk:"optional"`
	Key                  types.String                                                                           `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmOperatorDeploymentEnvValueFromSecretKeyRefModel `tfsdk:"local_object_reference"`
}
type LocalObjectReferenceInstallTemplateHelmOperatorDeploymentEnvValueFromConfigMapKeyRefModel struct {
	Name types.String `tfsdk:"name"`
}
type ConfigMapKeyRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel struct {
	Optional             types.Bool                                                                                `tfsdk:"optional"`
	Key                  types.String                                                                              `tfsdk:"key"`
	LocalObjectReference LocalObjectReferenceInstallTemplateHelmOperatorDeploymentEnvValueFromConfigMapKeyRefModel `tfsdk:"local_object_reference"`
}
type ValueFromInstallTemplateHelmOperatorDeploymentEnvModel struct {
	FieldRef         FieldRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel         `tfsdk:"field_ref"`
	ResourceFieldRef ResourceFieldRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel `tfsdk:"resource_field_ref"`
	SecretKeyRef     SecretKeyRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel     `tfsdk:"secret_key_ref"`
	ConfigMapKeyRef  ConfigMapKeyRefInstallTemplateHelmOperatorDeploymentEnvValueFromModel  `tfsdk:"config_map_key_ref"`
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
	StrVal StrValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"str_val"`
	Type   types.Int64                                                                   `tfsdk:"type"`
	IntVal IntValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxSurgeModel `tfsdk:"int_val"`
}
type IntValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.Int64 `tfsdk:"value"`
}
type StrValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxUnavailableModel struct {
	Value types.String `tfsdk:"value"`
}
type MaxUnavailableInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateModel struct {
	IntVal IntValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"int_val"`
	StrVal StrValInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateMaxUnavailableModel `tfsdk:"str_val"`
	Type   types.Int64                                                                         `tfsdk:"type"`
}
type RollingUpdateInstallTemplateHelmOperatorDeploymentStrategyModel struct {
	MaxSurge       MaxSurgeInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateModel       `tfsdk:"max_surge"`
	MaxUnavailable MaxUnavailableInstallTemplateHelmOperatorDeploymentStrategyRollingUpdateModel `tfsdk:"max_unavailable"`
}
type StrategyInstallTemplateHelmOperatorDeploymentModel struct {
	Type          types.String                                                    `tfsdk:"type"`
	RollingUpdate RollingUpdateInstallTemplateHelmOperatorDeploymentStrategyModel `tfsdk:"rolling_update"`
}
type TolerationsInstallTemplateHelmOperatorDeploymentModel struct {
	Key               types.String `tfsdk:"key"`
	Operator          types.String `tfsdk:"operator"`
	TolerationSeconds types.Int64  `tfsdk:"toleration_seconds"`
	Value             types.String `tfsdk:"value"`
	Effect            types.String `tfsdk:"effect"`
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
	LabelSelector LabelSelectorInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                                   `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                                 `tfsdk:"topology_key"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityModel struct {
	Weight          types.Int64                                                                                                                     `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
}
type LabelSelectorInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchLabels      types.Map                                                                                                                                       `tfsdk:"match_labels"`
	MatchExpressions []*MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                                   `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                                 `tfsdk:"topology_key"`
}
type PodAntiAffinityInstallTemplateHelmOperatorDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAntiAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchFieldsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type PreferenceInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionPreferenceModel      `tfsdk:"match_fields"`
}
type PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityModel struct {
	Weight     types.Int64                                                                                                             `tfsdk:"weight"`
	Preference PreferenceInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"preference"`
}
type MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
	Key      types.String `tfsdk:"key"`
}
type MatchFieldsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type NodeSelectorTermsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel `tfsdk:"match_expressions"`
	MatchFields      []*MatchFieldsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionNodeSelectorTermsModel      `tfsdk:"match_fields"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityModel struct {
	NodeSelectorTerms []*NodeSelectorTermsInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"node_selector_terms"`
}
type NodeAffinityInstallTemplateHelmOperatorDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityNodeAffinityModel     `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionPodAffinityTermLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
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
	Weight          types.Int64                                                                                                                 `tfsdk:"weight"`
	PodAffinityTerm PodAffinityTermInstallTemplateHelmOperatorDeploymentAffinityPodAffinityPreferredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"pod_affinity_term"`
}
type MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel struct {
	Key      types.String `tfsdk:"key"`
	Operator types.String `tfsdk:"operator"`
	Values   types.List   `tfsdk:"values"`
}
type LabelSelectorInstallTemplateHelmOperatorDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel struct {
	MatchExpressions []*MatchExpressionsInstallTemplateHelmOperatorDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionLabelSelectorModel `tfsdk:"match_expressions"`
	MatchLabels      types.Map                                                                                                                                   `tfsdk:"match_labels"`
}
type RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAffinityModel struct {
	LabelSelector LabelSelectorInstallTemplateHelmOperatorDeploymentAffinityPodAffinityRequiredDuringSchedulingIgnoredDuringExecutionModel `tfsdk:"label_selector"`
	Namespaces    types.List                                                                                                               `tfsdk:"namespaces"`
	TopologyKey   types.String                                                                                                             `tfsdk:"topology_key"`
}
type PodAffinityInstallTemplateHelmOperatorDeploymentAffinityModel struct {
	PreferredDuringSchedulingIgnoredDuringExecution []*PreferredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAffinityModel `tfsdk:"preferred_during_scheduling_ignored_during_execution"`
	RequiredDuringSchedulingIgnoredDuringExecution  []*RequiredDuringSchedulingIgnoredDuringExecutionInstallTemplateHelmOperatorDeploymentAffinityPodAffinityModel  `tfsdk:"required_during_scheduling_ignored_during_execution"`
}
type AffinityInstallTemplateHelmOperatorDeploymentModel struct {
	NodeAffinity    NodeAffinityInstallTemplateHelmOperatorDeploymentAffinityModel    `tfsdk:"node_affinity"`
	PodAffinity     PodAffinityInstallTemplateHelmOperatorDeploymentAffinityModel     `tfsdk:"pod_affinity"`
	PodAntiAffinity PodAntiAffinityInstallTemplateHelmOperatorDeploymentAffinityModel `tfsdk:"pod_anti_affinity"`
}
type DeploymentInstallTemplateHelmOperatorModel struct {
	PodAnnotations types.Map                                                `tfsdk:"pod_annotations"`
	ReplicaCount   types.Int64                                              `tfsdk:"replica_count"`
	Strategy       StrategyInstallTemplateHelmOperatorDeploymentModel       `tfsdk:"strategy"`
	Tolerations    []*TolerationsInstallTemplateHelmOperatorDeploymentModel `tfsdk:"tolerations"`
	Affinity       AffinityInstallTemplateHelmOperatorDeploymentModel       `tfsdk:"affinity"`
	Annotations    types.Map                                                `tfsdk:"annotations"`
	Env            []*EnvInstallTemplateHelmOperatorDeploymentModel         `tfsdk:"env"`
}
type ServiceInstallTemplateHelmOperatorModel struct {
	Annotations types.Map `tfsdk:"annotations"`
}
type ServiceAccountInstallTemplateHelmOperatorModel struct {
	PullSecret       types.String `tfsdk:"pull_secret"`
	PullUsername     types.String `tfsdk:"pull_username"`
	Annotations      types.Map    `tfsdk:"annotations"`
	ImagePullSecrets types.List   `tfsdk:"image_pull_secrets"`
	PullPassword     types.String `tfsdk:"pull_password"`
}
type OperatorInstallTemplateHelmModel struct {
	Deployment     DeploymentInstallTemplateHelmOperatorModel     `tfsdk:"deployment"`
	Service        ServiceInstallTemplateHelmOperatorModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmOperatorModel `tfsdk:"service_account"`
}
type HelmInstallTemplateModel struct {
	Operator OperatorInstallTemplateHelmModel `tfsdk:"operator"`
	Secrets  SecretsInstallTemplateHelmModel  `tfsdk:"secrets"`
	Spec     SpecInstallTemplateHelmModel     `tfsdk:"spec"`
	Image    ImageInstallTemplateHelmModel    `tfsdk:"image"`
}
type InstallTemplateModel struct {
	Message types.String             `tfsdk:"message"`
	Helm    HelmInstallTemplateModel `tfsdk:"helm"`
}
type ClusterModel struct {
	NamespaceScope  NamespaceScopeModel  `tfsdk:"namespace_scope"`
	InstallTemplate InstallTemplateModel `tfsdk:"install_template"`
	Tier1Cluster    Tier1ClusterModel    `tfsdk:"tier1_cluster"`
	TrustDomain     types.String         `tfsdk:"trust_domain"`
	Labels          types.Map            `tfsdk:"labels"`
	Id              types.String         `tfsdk:"id"`
	Name            types.String         `tfsdk:"name"`
	Locality        LocalityModel        `tfsdk:"locality"`
	TokenTtl        TokenTtlModel        `tfsdk:"token_ttl"`
	State           StateModel           `tfsdk:"state"`
	DisplayName     types.String         `tfsdk:"display_name"`
	ServiceAccount  ServiceAccountModel  `tfsdk:"service_account"`
	Namespaces      []*NamespacesModel   `tfsdk:"namespaces"`
	Network         types.String         `tfsdk:"network"`
	Parent          types.String         `tfsdk:"parent"`
	Description     types.String         `tfsdk:"description"`
}
