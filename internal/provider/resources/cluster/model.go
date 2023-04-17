package cluster

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type NamespaceScopeModel struct {
	Exceptions types.List   `tfsdk:"exceptions"`
	Scope      types.String `tfsdk:"scope"`
}
type ProxyNamespacesServicesWorkloadsModel struct {
	Status              types.Map    `tfsdk:"status"`
	ControlPlaneAddress types.String `tfsdk:"control_plane_address"`
	EnvoyVersion        types.String `tfsdk:"envoy_version"`
	IstioVersion        types.String `tfsdk:"istio_version"`
}
type WorkloadsNamespacesServicesModel struct {
	Proxy   ProxyNamespacesServicesWorkloadsModel `tfsdk:"proxy"`
	Address types.String                          `tfsdk:"address"`
	IsVm    types.Bool                            `tfsdk:"is_vm"`
	Name    types.String                          `tfsdk:"name"`
}
type PortsNamespacesServicesModel struct {
	KubernetesNodePort types.Int64  `tfsdk:"kubernetes_node_port"`
	Name               types.String `tfsdk:"name"`
	Number             types.Int64  `tfsdk:"number"`
}
type ServicesNamespacesModel struct {
	Workloads                   []*WorkloadsNamespacesServicesModel `tfsdk:"workloads"`
	KubernetesServiceFqdn       types.String                        `tfsdk:"kubernetes_service_fqdn"`
	NumHops                     types.Int64                         `tfsdk:"num_hops"`
	NumKubernetesEndpoints      types.Int64                         `tfsdk:"num_kubernetes_endpoints"`
	Namespace                   types.String                        `tfsdk:"namespace"`
	Tier1GatewayHost            types.Bool                          `tfsdk:"tier1_gateway_host"`
	Name                        types.String                        `tfsdk:"name"`
	Hostname                    types.String                        `tfsdk:"hostname"`
	NumVmEndpoints              types.Int64                         `tfsdk:"num_vm_endpoints"`
	SpiffeIds                   types.List                          `tfsdk:"spiffe_ids"`
	MeshExternal                types.Bool                          `tfsdk:"mesh_external"`
	Subsets                     types.List                          `tfsdk:"subsets"`
	KubernetesExternalAddresses types.List                          `tfsdk:"kubernetes_external_addresses"`
	GatewayHost                 types.Bool                          `tfsdk:"gateway_host"`
	Ports                       []*PortsNamespacesServicesModel     `tfsdk:"ports"`
	KubernetesServiceIp         types.String                        `tfsdk:"kubernetes_service_ip"`
	Selector                    types.Map                           `tfsdk:"selector"`
	CanonicalName               types.String                        `tfsdk:"canonical_name"`
}
type NamespacesModel struct {
	Services []*ServicesNamespacesModel `tfsdk:"services"`
	Name     types.String               `tfsdk:"name"`
}
type TokenTtlModel struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}
type LocalityModel struct {
	Region types.String `tfsdk:"region"`
}
type ImageInstallTemplateHelmModel struct {
	Registry types.String `tfsdk:"registry"`
	Tag      types.String `tfsdk:"tag"`
}
type ServiceInstallTemplateHelmOperatorModel struct {
	Annotations types.Map `tfsdk:"annotations"`
}
type ImagePullSecretsInstallTemplateHelmOperatorServiceAccountModel struct {
	TypeUrl types.String `tfsdk:"type_url"`
	Value   types.String `tfsdk:"value"`
}
type ServiceAccountInstallTemplateHelmOperatorModel struct {
	PullSecret       types.String                                                      `tfsdk:"pull_secret"`
	PullUsername     types.String                                                      `tfsdk:"pull_username"`
	Annotations      types.Map                                                         `tfsdk:"annotations"`
	ImagePullSecrets []*ImagePullSecretsInstallTemplateHelmOperatorServiceAccountModel `tfsdk:"image_pull_secrets"`
	PullPassword     types.String                                                      `tfsdk:"pull_password"`
}
type StrategyInstallTemplateHelmOperatorDeploymentModel struct {
	Value   types.String `tfsdk:"value"`
	TypeUrl types.String `tfsdk:"type_url"`
}
type TolerationsInstallTemplateHelmOperatorDeploymentModel struct {
	Value   types.String `tfsdk:"value"`
	TypeUrl types.String `tfsdk:"type_url"`
}
type AffinityInstallTemplateHelmOperatorDeploymentModel struct {
	TypeUrl types.String `tfsdk:"type_url"`
	Value   types.String `tfsdk:"value"`
}
type EnvInstallTemplateHelmOperatorDeploymentModel struct {
	Value   types.String `tfsdk:"value"`
	TypeUrl types.String `tfsdk:"type_url"`
}
type DeploymentInstallTemplateHelmOperatorModel struct {
	ReplicaCount   types.Int64                                              `tfsdk:"replica_count"`
	Strategy       StrategyInstallTemplateHelmOperatorDeploymentModel       `tfsdk:"strategy"`
	Tolerations    []*TolerationsInstallTemplateHelmOperatorDeploymentModel `tfsdk:"tolerations"`
	Affinity       AffinityInstallTemplateHelmOperatorDeploymentModel       `tfsdk:"affinity"`
	Annotations    types.Map                                                `tfsdk:"annotations"`
	Env            []*EnvInstallTemplateHelmOperatorDeploymentModel         `tfsdk:"env"`
	PodAnnotations types.Map                                                `tfsdk:"pod_annotations"`
}
type OperatorInstallTemplateHelmModel struct {
	Service        ServiceInstallTemplateHelmOperatorModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmOperatorModel `tfsdk:"service_account"`
	Deployment     DeploymentInstallTemplateHelmOperatorModel     `tfsdk:"deployment"`
}
type TsbInstallTemplateHelmSecretsModel struct {
	Cacert types.String `tfsdk:"cacert"`
}
type EdgeInstallTemplateHelmSecretsXcpModel struct {
	Cert  types.String `tfsdk:"cert"`
	Key   types.String `tfsdk:"key"`
	Token types.String `tfsdk:"token"`
}
type XcpInstallTemplateHelmSecretsModel struct {
	Rootca            types.String                           `tfsdk:"rootca"`
	Rootcakey         types.String                           `tfsdk:"rootcakey"`
	AutoGenerateCerts types.Bool                             `tfsdk:"auto_generate_certs"`
	Edge              EdgeInstallTemplateHelmSecretsXcpModel `tfsdk:"edge"`
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
type HelmInstallTemplateModel struct {
	Image    ImageInstallTemplateHelmModel    `tfsdk:"image"`
	Operator OperatorInstallTemplateHelmModel `tfsdk:"operator"`
	Secrets  SecretsInstallTemplateHelmModel  `tfsdk:"secrets"`
	Spec     SpecInstallTemplateHelmModel     `tfsdk:"spec"`
}
type InstallTemplateModel struct {
	Helm    HelmInstallTemplateModel `tfsdk:"helm"`
	Message types.String             `tfsdk:"message"`
}
type KeysServiceAccountModel struct {
	PublicKey    types.String `tfsdk:"public_key"`
	DefaultToken types.String `tfsdk:"default_token"`
	Encoding     types.String `tfsdk:"encoding"`
	Id           types.String `tfsdk:"id"`
	PrivateKey   types.String `tfsdk:"private_key"`
}
type ServiceAccountModel struct {
	DisplayName types.String               `tfsdk:"display_name"`
	Keys        []*KeysServiceAccountModel `tfsdk:"keys"`
	Description types.String               `tfsdk:"description"`
}
type Tier1ClusterModel struct {
	Value types.Bool `tfsdk:"value"`
}
type LastSyncTimeStateModel struct {
	Seconds types.Int64 `tfsdk:"seconds"`
	Nanos   types.Int64 `tfsdk:"nanos"`
}
type StateModel struct {
	IstioVersions types.List             `tfsdk:"istio_versions"`
	LastSyncTime  LastSyncTimeStateModel `tfsdk:"last_sync_time"`
	Provider      types.String           `tfsdk:"provider"`
	TsbCpVersion  types.String           `tfsdk:"tsb_cp_version"`
	XcpVersion    types.String           `tfsdk:"xcp_version"`
}
type ClusterModel struct {
	Network         types.String         `tfsdk:"network"`
	TokenTtl        TokenTtlModel        `tfsdk:"token_ttl"`
	DisplayName     types.String         `tfsdk:"display_name"`
	TrustDomain     types.String         `tfsdk:"trust_domain"`
	Labels          types.Map            `tfsdk:"labels"`
	Namespaces      []*NamespacesModel   `tfsdk:"namespaces"`
	Parent          types.String         `tfsdk:"parent"`
	Locality        LocalityModel        `tfsdk:"locality"`
	Name            types.String         `tfsdk:"name"`
	ServiceAccount  ServiceAccountModel  `tfsdk:"service_account"`
	Id              types.String         `tfsdk:"id"`
	Description     types.String         `tfsdk:"description"`
	State           StateModel           `tfsdk:"state"`
	NamespaceScope  NamespaceScopeModel  `tfsdk:"namespace_scope"`
	InstallTemplate InstallTemplateModel `tfsdk:"install_template"`
	Tier1Cluster    Tier1ClusterModel    `tfsdk:"tier1_cluster"`
}
