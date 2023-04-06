package cluster

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type LocalityModel struct {
	Region types.String `tfsdk:"region"`
}
type NamespacesModel struct{}
type LabelsModel struct{}
type KeysServiceAccountModel struct{}
type ServiceAccountModel struct {
	DisplayName types.String            `tfsdk:"display_name"`
	Keys        KeysServiceAccountModel `tfsdk:"keys"`
	Description types.String            `tfsdk:"description"`
}
type Tier1ClusterModel struct {
	Value types.Bool `tfsdk:"value"`
}
type NamespaceScopeModel struct {
	Scope      types.String `tfsdk:"scope"`
	Exceptions types.List   `tfsdk:"exceptions"`
}
type ImageInstallTemplateHelmModel struct {
	Registry types.String `tfsdk:"registry"`
	Tag      types.String `tfsdk:"tag"`
}
type TolerationsInstallTemplateHelmOperatorDeploymentModel struct{}
type AffinityInstallTemplateHelmOperatorDeploymentModel struct {
	TypeUrl types.String `tfsdk:"type_url"`
	Value   types.String `tfsdk:"value"`
}
type AnnotationsInstallTemplateHelmOperatorDeploymentModel struct{}
type EnvInstallTemplateHelmOperatorDeploymentModel struct{}
type PodAnnotationsInstallTemplateHelmOperatorDeploymentModel struct{}
type StrategyInstallTemplateHelmOperatorDeploymentModel struct {
	TypeUrl types.String `tfsdk:"type_url"`
	Value   types.String `tfsdk:"value"`
}
type DeploymentInstallTemplateHelmOperatorModel struct {
	PodAnnotations PodAnnotationsInstallTemplateHelmOperatorDeploymentModel `tfsdk:"pod_annotations"`
	ReplicaCount   types.Number                                             `tfsdk:"replica_count"`
	Strategy       StrategyInstallTemplateHelmOperatorDeploymentModel       `tfsdk:"strategy"`
	Tolerations    TolerationsInstallTemplateHelmOperatorDeploymentModel    `tfsdk:"tolerations"`
	Affinity       AffinityInstallTemplateHelmOperatorDeploymentModel       `tfsdk:"affinity"`
	Annotations    AnnotationsInstallTemplateHelmOperatorDeploymentModel    `tfsdk:"annotations"`
	Env            EnvInstallTemplateHelmOperatorDeploymentModel            `tfsdk:"env"`
}
type AnnotationsInstallTemplateHelmOperatorServiceModel struct{}
type ServiceInstallTemplateHelmOperatorModel struct {
	Annotations AnnotationsInstallTemplateHelmOperatorServiceModel `tfsdk:"annotations"`
}
type ImagePullSecretsInstallTemplateHelmOperatorServiceAccountModel struct{}
type AnnotationsInstallTemplateHelmOperatorServiceAccountModel struct{}
type ServiceAccountInstallTemplateHelmOperatorModel struct {
	PullUsername     types.String                                                   `tfsdk:"pull_username"`
	Annotations      AnnotationsInstallTemplateHelmOperatorServiceAccountModel      `tfsdk:"annotations"`
	ImagePullSecrets ImagePullSecretsInstallTemplateHelmOperatorServiceAccountModel `tfsdk:"image_pull_secrets"`
	PullPassword     types.String                                                   `tfsdk:"pull_password"`
	PullSecret       types.String                                                   `tfsdk:"pull_secret"`
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
	Key   types.String `tfsdk:"key"`
	Token types.String `tfsdk:"token"`
	Cert  types.String `tfsdk:"cert"`
}
type XcpInstallTemplateHelmSecretsModel struct {
	Rootcakey         types.String                           `tfsdk:"rootcakey"`
	AutoGenerateCerts types.Bool                             `tfsdk:"auto_generate_certs"`
	Edge              EdgeInstallTemplateHelmSecretsXcpModel `tfsdk:"edge"`
	Rootca            types.String                           `tfsdk:"rootca"`
}
type ClusterServiceAccountInstallTemplateHelmSecretsModel struct {
	ClusterFqn types.String `tfsdk:"cluster_fqn"`
	EncodedJwk types.String `tfsdk:"encoded_jwk"`
	Jwk        types.String `tfsdk:"jwk"`
}
type ElasticsearchInstallTemplateHelmSecretsModel struct {
	Cacert   types.String `tfsdk:"cacert"`
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}
type SecretsInstallTemplateHelmModel struct {
	Xcp                   XcpInstallTemplateHelmSecretsModel                   `tfsdk:"xcp"`
	ClusterServiceAccount ClusterServiceAccountInstallTemplateHelmSecretsModel `tfsdk:"cluster_service_account"`
	Elasticsearch         ElasticsearchInstallTemplateHelmSecretsModel         `tfsdk:"elasticsearch"`
	Tsb                   TsbInstallTemplateHelmSecretsModel                   `tfsdk:"tsb"`
}
type SpecInstallTemplateHelmModel struct {
	TypeUrl types.String `tfsdk:"type_url"`
	Value   types.String `tfsdk:"value"`
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
type LastSyncTimeStateModel struct {
	Nanos   types.Number `tfsdk:"nanos"`
	Seconds types.Number `tfsdk:"seconds"`
}
type StateModel struct {
	XcpVersion    types.String           `tfsdk:"xcp_version"`
	IstioVersions types.List             `tfsdk:"istio_versions"`
	LastSyncTime  LastSyncTimeStateModel `tfsdk:"last_sync_time"`
	Provider      types.String           `tfsdk:"provider"`
	TsbCpVersion  types.String           `tfsdk:"tsb_cp_version"`
}
type TokenTtlModel struct {
	Nanos   types.Number `tfsdk:"nanos"`
	Seconds types.Number `tfsdk:"seconds"`
}
type ClusterModel struct {
	Labels          LabelsModel          `tfsdk:"labels"`
	ServiceAccount  ServiceAccountModel  `tfsdk:"service_account"`
	Parent          types.String         `tfsdk:"parent"`
	State           StateModel           `tfsdk:"state"`
	Namespaces      NamespacesModel      `tfsdk:"namespaces"`
	Description     types.String         `tfsdk:"description"`
	Network         types.String         `tfsdk:"network"`
	DisplayName     types.String         `tfsdk:"display_name"`
	Tier1Cluster    Tier1ClusterModel    `tfsdk:"tier1_cluster"`
	NamespaceScope  NamespaceScopeModel  `tfsdk:"namespace_scope"`
	TrustDomain     types.String         `tfsdk:"trust_domain"`
	InstallTemplate InstallTemplateModel `tfsdk:"install_template"`
	TokenTtl        TokenTtlModel        `tfsdk:"token_ttl"`
	Locality        LocalityModel        `tfsdk:"locality"`
	Id              types.String         `tfsdk:"id"`
	Name            types.String         `tfsdk:"name"`
}
