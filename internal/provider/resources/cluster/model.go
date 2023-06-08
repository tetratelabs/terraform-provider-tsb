package cluster

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
// state/last_sync_time
type State_LastSyncTime_Model struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}

// state/discovered_locality
type State_DiscoveredLocality_Model struct {
	Region types.String `tfsdk:"region"`
}

// state
type State_Model struct {
	Provider           types.String                   `tfsdk:"provider"`
	TsbCpVersion       types.String                   `tfsdk:"tsb_cp_version"`
	XcpVersion         types.String                   `tfsdk:"xcp_version"`
	DiscoveredLocality State_DiscoveredLocality_Model `tfsdk:"discovered_locality"`
	IstioVersions      types.List                     `tfsdk:"istio_versions"`
	LastSyncTime       State_LastSyncTime_Model       `tfsdk:"last_sync_time"`
}

// namespaces/services/workloads/proxy
type Namespaces_Services_Workloads_Proxy_Model struct {
	IstioVersion        types.String `tfsdk:"istio_version"`
	Status              types.Map    `tfsdk:"status"`
	ControlPlaneAddress types.String `tfsdk:"control_plane_address"`
	EnvoyVersion        types.String `tfsdk:"envoy_version"`
}

// namespaces/services/workloads
type Namespaces_Services_Workloads_Model struct {
	Name    types.String                              `tfsdk:"name"`
	Proxy   Namespaces_Services_Workloads_Proxy_Model `tfsdk:"proxy"`
	Address types.String                              `tfsdk:"address"`
	IsVm    types.Bool                                `tfsdk:"is_vm"`
}

// namespaces/services/ports
type Namespaces_Services_Ports_Model struct {
	KubernetesNodePort types.Int64  `tfsdk:"kubernetes_node_port"`
	Name               types.String `tfsdk:"name"`
	Number             types.Int64  `tfsdk:"number"`
}

// namespaces/services
type Namespaces_Services_Model struct {
	Hostname                    types.String `tfsdk:"hostname"`
	KubernetesServiceIp         types.String `tfsdk:"kubernetes_service_ip"`
	NumHops                     types.Int64  `tfsdk:"num_hops"`
	Namespace                   types.String `tfsdk:"namespace"`
	Selector                    types.Map    `tfsdk:"selector"`
	CanonicalName               types.String `tfsdk:"canonical_name"`
	KubernetesExternalAddresses types.List   `tfsdk:"kubernetes_external_addresses"`
	NumVmEndpoints              types.Int64  `tfsdk:"num_vm_endpoints"`
	MeshExternal                types.Bool   `tfsdk:"mesh_external"`
	Workloads                   types.List   `tfsdk:"workloads"` // Namespaces_Services_Workloads_Model
	SpiffeIds                   types.List   `tfsdk:"spiffe_ids"`
	KubernetesServiceFqdn       types.String `tfsdk:"kubernetes_service_fqdn"`
	Ports                       types.List   `tfsdk:"ports"` // Namespaces_Services_Ports_Model
	Subsets                     types.List   `tfsdk:"subsets"`
	Tier1GatewayHost            types.Bool   `tfsdk:"tier1_gateway_host"`
	NumKubernetesEndpoints      types.Int64  `tfsdk:"num_kubernetes_endpoints"`
	Name                        types.String `tfsdk:"name"`
	GatewayHost                 types.Bool   `tfsdk:"gateway_host"`
}

// namespaces
type Namespaces_Model struct {
	Services types.List   `tfsdk:"services"` // Namespaces_Services_Model
	Name     types.String `tfsdk:"name"`
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
	DisplayName types.String `tfsdk:"display_name"`
	Keys        types.List   `tfsdk:"keys"` // ServiceAccount_Keys_Model
	Description types.String `tfsdk:"description"`
}

// token_ttl
type TokenTtl_Model struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}

// locality
type Locality_Model struct {
	Region types.String `tfsdk:"region"`
}

// namespace_scope
type NamespaceScope_Model struct {
	Exceptions types.List   `tfsdk:"exceptions"`
	Scope      types.String `tfsdk:"scope"`
}

// tier1_cluster
type Tier1Cluster_Model struct {
	Value types.Bool `tfsdk:"value"`
}

type ClusterModel struct {
	Parent         types.String         `tfsdk:"parent"`
	Name           types.String         `tfsdk:"name"`
	Namespaces     types.List           `tfsdk:"namespaces"` // Namespaces_Model
	TokenTtl       TokenTtl_Model       `tfsdk:"token_ttl"`
	NamespaceScope NamespaceScope_Model `tfsdk:"namespace_scope"`
	DisplayName    types.String         `tfsdk:"display_name"`
	Tier1Cluster   Tier1Cluster_Model   `tfsdk:"tier1_cluster"`
	Description    types.String         `tfsdk:"description"`
	Network        types.String         `tfsdk:"network"`
	Labels         types.Map            `tfsdk:"labels"`
	Locality       Locality_Model       `tfsdk:"locality"`
	State          State_Model          `tfsdk:"state"`
	Id             types.String         `tfsdk:"id"`
	ServiceAccount ServiceAccount_Model `tfsdk:"service_account"`
	TrustDomain    types.String         `tfsdk:"trust_domain"`
}
