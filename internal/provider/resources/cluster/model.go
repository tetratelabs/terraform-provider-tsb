package cluster

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type NamespaceScopeModel struct {
	Scope      types.String `tfsdk:"scope"`
	Exceptions types.List   `tfsdk:"exceptions"`
}
type Tier1ClusterModel struct {
	Value types.Bool `tfsdk:"value"`
}
type KeysServiceAccountModel struct{}
type ServiceAccountModel struct {
	DisplayName types.String            `tfsdk:"display_name"`
	Keys        KeysServiceAccountModel `tfsdk:"keys"`
	Description types.String            `tfsdk:"description"`
}
type LocalityModel struct {
	Region types.String `tfsdk:"region"`
}
type TokenTtlModel struct {
	Nanos   types.Number `tfsdk:"nanos"`
	Seconds types.Number `tfsdk:"seconds"`
}
type EnvInstallTemplateHelmOperatorDeploymentModel struct{}
type PodAnnotationsInstallTemplateHelmOperatorDeploymentModel struct{}
type StrategyInstallTemplateHelmOperatorDeploymentModel struct {
	Value   types.String `tfsdk:"value"`
	TypeUrl types.String `tfsdk:"type_url"`
}
type TolerationsInstallTemplateHelmOperatorDeploymentModel struct{}
type AffinityInstallTemplateHelmOperatorDeploymentModel struct {
	TypeUrl types.String `tfsdk:"type_url"`
	Value   types.String `tfsdk:"value"`
}
type AnnotationsInstallTemplateHelmOperatorDeploymentModel struct{}
type DeploymentInstallTemplateHelmOperatorModel struct {
	Tolerations    TolerationsInstallTemplateHelmOperatorDeploymentModel    `tfsdk:"tolerations"`
	Affinity       AffinityInstallTemplateHelmOperatorDeploymentModel       `tfsdk:"affinity"`
	Annotations    AnnotationsInstallTemplateHelmOperatorDeploymentModel    `tfsdk:"annotations"`
	Env            EnvInstallTemplateHelmOperatorDeploymentModel            `tfsdk:"env"`
	PodAnnotations PodAnnotationsInstallTemplateHelmOperatorDeploymentModel `tfsdk:"pod_annotations"`
	ReplicaCount   types.Number                                             `tfsdk:"replica_count"`
	Strategy       StrategyInstallTemplateHelmOperatorDeploymentModel       `tfsdk:"strategy"`
}
type AnnotationsInstallTemplateHelmOperatorServiceModel struct{}
type ServiceInstallTemplateHelmOperatorModel struct {
	Annotations AnnotationsInstallTemplateHelmOperatorServiceModel `tfsdk:"annotations"`
}
type ImagePullSecretsInstallTemplateHelmOperatorServiceAccountModel struct{}
type AnnotationsInstallTemplateHelmOperatorServiceAccountModel struct{}
type ServiceAccountInstallTemplateHelmOperatorModel struct {
	PullSecret       types.String                                                   `tfsdk:"pull_secret"`
	PullUsername     types.String                                                   `tfsdk:"pull_username"`
	Annotations      AnnotationsInstallTemplateHelmOperatorServiceAccountModel      `tfsdk:"annotations"`
	ImagePullSecrets ImagePullSecretsInstallTemplateHelmOperatorServiceAccountModel `tfsdk:"image_pull_secrets"`
	PullPassword     types.String                                                   `tfsdk:"pull_password"`
}
type OperatorInstallTemplateHelmModel struct {
	Deployment     DeploymentInstallTemplateHelmOperatorModel     `tfsdk:"deployment"`
	Service        ServiceInstallTemplateHelmOperatorModel        `tfsdk:"service"`
	ServiceAccount ServiceAccountInstallTemplateHelmOperatorModel `tfsdk:"service_account"`
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
type ElasticsearchInstallTemplateHelmSecretsModel struct {
	Username types.String `tfsdk:"username"`
	Cacert   types.String `tfsdk:"cacert"`
	Password types.String `tfsdk:"password"`
}
type TsbInstallTemplateHelmSecretsModel struct {
	Cacert types.String `tfsdk:"cacert"`
}
type SecretsInstallTemplateHelmModel struct {
	Elasticsearch         ElasticsearchInstallTemplateHelmSecretsModel         `tfsdk:"elasticsearch"`
	Tsb                   TsbInstallTemplateHelmSecretsModel                   `tfsdk:"tsb"`
	Xcp                   XcpInstallTemplateHelmSecretsModel                   `tfsdk:"xcp"`
	ClusterServiceAccount ClusterServiceAccountInstallTemplateHelmSecretsModel `tfsdk:"cluster_service_account"`
}
type SpecInstallTemplateHelmModel struct {
	TypeUrl types.String `tfsdk:"type_url"`
	Value   types.String `tfsdk:"value"`
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
type LabelsModel struct{}
type NamespacesModel struct{}
type LastSyncTimeStateModel struct {
	Nanos   types.Number `tfsdk:"nanos"`
	Seconds types.Number `tfsdk:"seconds"`
}
type StateModel struct {
	TsbCpVersion  types.String           `tfsdk:"tsb_cp_version"`
	XcpVersion    types.String           `tfsdk:"xcp_version"`
	IstioVersions types.List             `tfsdk:"istio_versions"`
	LastSyncTime  LastSyncTimeStateModel `tfsdk:"last_sync_time"`
	Provider      types.String           `tfsdk:"provider"`
}
type ClusterModel struct {
	Description     types.String         `tfsdk:"description"`
	Locality        LocalityModel        `tfsdk:"locality"`
	Network         types.String         `tfsdk:"network"`
	Namespaces      NamespacesModel      `tfsdk:"namespaces"`
	TrustDomain     types.String         `tfsdk:"trust_domain"`
	TokenTtl        TokenTtlModel        `tfsdk:"token_ttl"`
	Labels          LabelsModel          `tfsdk:"labels"`
	ServiceAccount  ServiceAccountModel  `tfsdk:"service_account"`
	DisplayName     types.String         `tfsdk:"display_name"`
	Name            types.String         `tfsdk:"name"`
	Id              types.String         `tfsdk:"id"`
	State           StateModel           `tfsdk:"state"`
	NamespaceScope  NamespaceScopeModel  `tfsdk:"namespace_scope"`
	Tier1Cluster    Tier1ClusterModel    `tfsdk:"tier1_cluster"`
	Parent          types.String         `tfsdk:"parent"`
	InstallTemplate InstallTemplateModel `tfsdk:"install_template"`
}
