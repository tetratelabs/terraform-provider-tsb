package cluster

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
// namespaces/services/ports
type Namespaces_Services_Ports_Model struct {
	Number             types.Int64  `tfsdk:"number"`
	KubernetesNodePort types.Int64  `tfsdk:"kubernetes_node_port"`
	Name               types.String `tfsdk:"name"`
}

// namespaces/services/workloads/proxy
type Namespaces_Services_Workloads_Proxy_Model struct {
	EnvoyVersion        types.String `tfsdk:"envoy_version"`
	IstioVersion        types.String `tfsdk:"istio_version"`
	Status              types.Map    `tfsdk:"status"`
	ControlPlaneAddress types.String `tfsdk:"control_plane_address"`
}

// namespaces/services/workloads
type Namespaces_Services_Workloads_Model struct {
	Proxy   Namespaces_Services_Workloads_Proxy_Model `tfsdk:"proxy"`
	Address types.String                              `tfsdk:"address"`
	IsVm    types.Bool                                `tfsdk:"is_vm"`
	Name    types.String                              `tfsdk:"name"`
}

// namespaces/services
type Namespaces_Services_Model struct {
	Subsets                     types.List                             `tfsdk:"subsets"`
	NumKubernetesEndpoints      types.Int64                            `tfsdk:"num_kubernetes_endpoints"`
	KubernetesServiceFqdn       types.String                           `tfsdk:"kubernetes_service_fqdn"`
	Name                        types.String                           `tfsdk:"name"`
	Namespace                   types.String                           `tfsdk:"namespace"`
	Tier1GatewayHost            types.Bool                             `tfsdk:"tier1_gateway_host"`
	NumVmEndpoints              types.Int64                            `tfsdk:"num_vm_endpoints"`
	Hostname                    types.String                           `tfsdk:"hostname"`
	KubernetesExternalAddresses types.List                             `tfsdk:"kubernetes_external_addresses"`
	Selector                    types.Map                              `tfsdk:"selector"`
	SpiffeIds                   types.List                             `tfsdk:"spiffe_ids"`
	GatewayHost                 types.Bool                             `tfsdk:"gateway_host"`
	Workloads                   []*Namespaces_Services_Workloads_Model `tfsdk:"workloads"`
	MeshExternal                types.Bool                             `tfsdk:"mesh_external"`
	Ports                       []*Namespaces_Services_Ports_Model     `tfsdk:"ports"`
	CanonicalName               types.String                           `tfsdk:"canonical_name"`
	NumHops                     types.Int64                            `tfsdk:"num_hops"`
	KubernetesServiceIp         types.String                           `tfsdk:"kubernetes_service_ip"`
}

// namespaces
type Namespaces_Model struct {
	Name     types.String                 `tfsdk:"name"`
	Services []*Namespaces_Services_Model `tfsdk:"services"`
}

// tier1_cluster
type Tier1Cluster_Model struct {
	Value types.Bool `tfsdk:"value"`
}

// namespace_scope
type NamespaceScope_Model struct {
	Scope      types.String `tfsdk:"scope"`
	Exceptions types.List   `tfsdk:"exceptions"`
}

// token_ttl
type TokenTtl_Model struct {
	Seconds types.Int64 `tfsdk:"seconds"`
	Nanos   types.Int64 `tfsdk:"nanos"`
}

// locality
type Locality_Model struct {
	Region types.String `tfsdk:"region"`
}

// service_account/keys
type ServiceAccount_Keys_Model struct {
	Id           types.String `tfsdk:"id"`
	PrivateKey   types.String `tfsdk:"private_key"`
	PublicKey    types.String `tfsdk:"public_key"`
	DefaultToken types.String `tfsdk:"default_token"`
	Encoding     types.String `tfsdk:"encoding"`
}

// service_account
type ServiceAccount_Model struct {
	Keys        []*ServiceAccount_Keys_Model `tfsdk:"keys"`
	Description types.String                 `tfsdk:"description"`
	DisplayName types.String                 `tfsdk:"display_name"`
}

// state/last_sync_time
type State_LastSyncTime_Model struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}

// state
type State_Model struct {
	TsbCpVersion  types.String             `tfsdk:"tsb_cp_version"`
	XcpVersion    types.String             `tfsdk:"xcp_version"`
	IstioVersions types.List               `tfsdk:"istio_versions"`
	LastSyncTime  State_LastSyncTime_Model `tfsdk:"last_sync_time"`
	Provider      types.String             `tfsdk:"provider"`
}

type ClusterModel struct {
	Tier1Cluster   Tier1Cluster_Model   `tfsdk:"tier1_cluster"`
	Id             types.String         `tfsdk:"id"`
	Parent         types.String         `tfsdk:"parent"`
	TokenTtl       TokenTtl_Model       `tfsdk:"token_ttl"`
	Locality       Locality_Model       `tfsdk:"locality"`
	State          State_Model          `tfsdk:"state"`
	Labels         types.Map            `tfsdk:"labels"`
	Network        types.String         `tfsdk:"network"`
	NamespaceScope NamespaceScope_Model `tfsdk:"namespace_scope"`
	TrustDomain    types.String         `tfsdk:"trust_domain"`
	Name           types.String         `tfsdk:"name"`
	Namespaces     []*Namespaces_Model  `tfsdk:"namespaces"`
	ServiceAccount ServiceAccount_Model `tfsdk:"service_account"`
	Description    types.String         `tfsdk:"description"`
	DisplayName    types.String         `tfsdk:"display_name"`
}
