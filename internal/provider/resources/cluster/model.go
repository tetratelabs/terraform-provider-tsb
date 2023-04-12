package cluster

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type Tier1ClusterModel struct {
	Value types.Bool `tfsdk:"value"`
}
type LocalityModel struct {
	Region types.String `tfsdk:"region"`
}
type SelectorNamespacesServicesModel struct{}
type StatusNamespacesServicesWorkloadsProxyModel struct{}
type ProxyNamespacesServicesWorkloadsModel struct {
	EnvoyVersion        types.String                                `tfsdk:"envoy_version"`
	IstioVersion        types.String                                `tfsdk:"istio_version"`
	Status              StatusNamespacesServicesWorkloadsProxyModel `tfsdk:"status"`
	ControlPlaneAddress types.String                                `tfsdk:"control_plane_address"`
}
type WorkloadsNamespacesServicesModel struct {
	Proxy   ProxyNamespacesServicesWorkloadsModel `tfsdk:"proxy"`
	Address types.String                          `tfsdk:"address"`
	IsVm    types.Bool                            `tfsdk:"is_vm"`
	Name    types.String                          `tfsdk:"name"`
}
type PortsNamespacesServicesModel struct {
	KubernetesNodePort types.Number `tfsdk:"kubernetes_node_port"`
	Name               types.String `tfsdk:"name"`
	Number             types.Number `tfsdk:"number"`
}
type ServicesNamespacesModel struct {
	MeshExternal                types.Bool                          `tfsdk:"mesh_external"`
	Namespace                   types.String                        `tfsdk:"namespace"`
	NumHops                     types.Number                        `tfsdk:"num_hops"`
	Hostname                    types.String                        `tfsdk:"hostname"`
	KubernetesServiceFqdn       types.String                        `tfsdk:"kubernetes_service_fqdn"`
	Selector                    SelectorNamespacesServicesModel     `tfsdk:"selector"`
	CanonicalName               types.String                        `tfsdk:"canonical_name"`
	GatewayHost                 types.Bool                          `tfsdk:"gateway_host"`
	NumVmEndpoints              types.Number                        `tfsdk:"num_vm_endpoints"`
	KubernetesExternalAddresses types.List                          `tfsdk:"kubernetes_external_addresses"`
	Tier1GatewayHost            types.Bool                          `tfsdk:"tier1_gateway_host"`
	Workloads                   []*WorkloadsNamespacesServicesModel `tfsdk:"workloads"`
	Name                        types.String                        `tfsdk:"name"`
	Subsets                     types.List                          `tfsdk:"subsets"`
	Ports                       []*PortsNamespacesServicesModel     `tfsdk:"ports"`
	NumKubernetesEndpoints      types.Number                        `tfsdk:"num_kubernetes_endpoints"`
	SpiffeIds                   types.List                          `tfsdk:"spiffe_ids"`
	KubernetesServiceIp         types.String                        `tfsdk:"kubernetes_service_ip"`
}
type NamespacesModel struct {
	Name     types.String               `tfsdk:"name"`
	Services []*ServicesNamespacesModel `tfsdk:"services"`
}
type LabelsModel struct{}
type NamespaceScopeModel struct {
	Scope      types.String `tfsdk:"scope"`
	Exceptions types.List   `tfsdk:"exceptions"`
}
type AnnotationsInstallTemplateHelmOperatorServiceAccountModel struct{}
type ImagePullSecretsInstallTemplateHelmOperatorServiceAccountModel struct {
	Value   types.String `tfsdk:"value"`
	TypeUrl types.String `tfsdk:"type_url"`
}
type ServiceAccountInstallTemplateHelmOperatorModel struct {
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmOperatorServiceAccountModel `tfsdk:"image_pull_secrets"`
	PullPassword     types.String                                                      `tfsdk:"pull_password"`
	PullSecret       types.String                                                      `tfsdk:"pull_secret"`
	PullUsername     types.String                                                      `tfsdk:"pull_username"`
	Annotations      AnnotationsInstallTemplateHelmOperatorServiceAccountModel         `tfsdk:"annotations"`
}
type TolerationsInstallTemplateHelmOperatorDeploymentModel struct {
	Value   types.String `tfsdk:"value"`
	TypeUrl types.String `tfsdk:"type_url"`
}
type AffinityInstallTemplateHelmOperatorDeploymentModel struct {
	TypeUrl types.String `tfsdk:"type_url"`
	Value   types.String `tfsdk:"value"`
}
type AnnotationsInstallTemplateHelmOperatorDeploymentModel struct{}
type EnvInstallTemplateHelmOperatorDeploymentModel struct {
	TypeUrl types.String `tfsdk:"type_url"`
	Value   types.String `tfsdk:"value"`
}
type PodAnnotationsInstallTemplateHelmOperatorDeploymentModel struct{}
type StrategyInstallTemplateHelmOperatorDeploymentModel struct {
	Value   types.String `tfsdk:"value"`
	TypeUrl types.String `tfsdk:"type_url"`
}
type DeploymentInstallTemplateHelmOperatorModel struct {
	Strategy       StrategyInstallTemplateHelmOperatorDeploymentModel       `tfsdk:"strategy"`
	Tolerations    []*TolerationsInstallTemplateHelmOperatorDeploymentModel `tfsdk:"tolerations"`
	Affinity       AffinityInstallTemplateHelmOperatorDeploymentModel       `tfsdk:"affinity"`
	Annotations    AnnotationsInstallTemplateHelmOperatorDeploymentModel    `tfsdk:"annotations"`
	Env            []*EnvInstallTemplateHelmOperatorDeploymentModel         `tfsdk:"env"`
	PodAnnotations PodAnnotationsInstallTemplateHelmOperatorDeploymentModel `tfsdk:"pod_annotations"`
	ReplicaCount   types.Number                                             `tfsdk:"replica_count"`
}
type AnnotationsInstallTemplateHelmOperatorServiceModel struct{}
type ServiceInstallTemplateHelmOperatorModel struct {
	Annotations AnnotationsInstallTemplateHelmOperatorServiceModel `tfsdk:"annotations"`
}
type OperatorInstallTemplateHelmModel struct {
	ServiceAccount ServiceAccountInstallTemplateHelmOperatorModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmOperatorModel     `tfsdk:"deployment"`
	Service        ServiceInstallTemplateHelmOperatorModel        `tfsdk:"service"`
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
type ClusterServiceAccountInstallTemplateHelmSecretsModel struct {
	ClusterFqn types.String `tfsdk:"cluster_fqn"`
	EncodedJwk types.String `tfsdk:"encoded_jwk"`
	Jwk        types.String `tfsdk:"jwk"`
}
type SecretsInstallTemplateHelmModel struct {
	ClusterServiceAccount ClusterServiceAccountInstallTemplateHelmSecretsModel `tfsdk:"cluster_service_account"`
	Elasticsearch         ElasticsearchInstallTemplateHelmSecretsModel         `tfsdk:"elasticsearch"`
	Tsb                   TsbInstallTemplateHelmSecretsModel                   `tfsdk:"tsb"`
	Xcp                   XcpInstallTemplateHelmSecretsModel                   `tfsdk:"xcp"`
}
type SpecInstallTemplateHelmModel struct {
	Value   types.String `tfsdk:"value"`
	TypeUrl types.String `tfsdk:"type_url"`
}
type ImageInstallTemplateHelmModel struct {
	Tag      types.String `tfsdk:"tag"`
	Registry types.String `tfsdk:"registry"`
}
type HelmInstallTemplateModel struct {
	Operator OperatorInstallTemplateHelmModel `tfsdk:"operator"`
	Secrets  SecretsInstallTemplateHelmModel  `tfsdk:"secrets"`
	Spec     SpecInstallTemplateHelmModel     `tfsdk:"spec"`
	Image    ImageInstallTemplateHelmModel    `tfsdk:"image"`
}
type InstallTemplateModel struct {
	Helm    HelmInstallTemplateModel `tfsdk:"helm"`
	Message types.String             `tfsdk:"message"`
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
	Seconds types.Number `tfsdk:"seconds"`
	Nanos   types.Number `tfsdk:"nanos"`
}
type LastSyncTimeStateModel struct {
	Seconds types.Number `tfsdk:"seconds"`
	Nanos   types.Number `tfsdk:"nanos"`
}
type StateModel struct {
	IstioVersions types.List             `tfsdk:"istio_versions"`
	LastSyncTime  LastSyncTimeStateModel `tfsdk:"last_sync_time"`
	Provider      types.String           `tfsdk:"provider"`
	TsbCpVersion  types.String           `tfsdk:"tsb_cp_version"`
	XcpVersion    types.String           `tfsdk:"xcp_version"`
}
type ClusterModel struct {
	Parent          types.String         `tfsdk:"parent"`
	Description     types.String         `tfsdk:"description"`
	Tier1Cluster    Tier1ClusterModel    `tfsdk:"tier1_cluster"`
	TokenTtl        TokenTtlModel        `tfsdk:"token_ttl"`
	Id              types.String         `tfsdk:"id"`
	NamespaceScope  NamespaceScopeModel  `tfsdk:"namespace_scope"`
	InstallTemplate InstallTemplateModel `tfsdk:"install_template"`
	ServiceAccount  ServiceAccountModel  `tfsdk:"service_account"`
	TrustDomain     types.String         `tfsdk:"trust_domain"`
	Locality        LocalityModel        `tfsdk:"locality"`
	Namespaces      []*NamespacesModel   `tfsdk:"namespaces"`
	Labels          LabelsModel          `tfsdk:"labels"`
	State           StateModel           `tfsdk:"state"`
	Name            types.String         `tfsdk:"name"`
	Network         types.String         `tfsdk:"network"`
	DisplayName     types.String         `tfsdk:"display_name"`
}
