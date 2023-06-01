package cluster

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
// state/discovered_locality
type State_DiscoveredLocality_Model struct {
	Region types.String `tfsdk:"region"`
}

// state/last_sync_time
type State_LastSyncTime_Model struct {
	Seconds types.Int64 `tfsdk:"seconds"`
	Nanos   types.Int64 `tfsdk:"nanos"`
}

// state
type State_Model struct {
	TsbCpVersion       types.String                   `tfsdk:"tsb_cp_version"`
	XcpVersion         types.String                   `tfsdk:"xcp_version"`
	DiscoveredLocality State_DiscoveredLocality_Model `tfsdk:"discovered_locality"`
	IstioVersions      types.List                     `tfsdk:"istio_versions"`
	LastSyncTime       State_LastSyncTime_Model       `tfsdk:"last_sync_time"`
	Provider           types.String                   `tfsdk:"provider"`
}

// tier1_cluster
type Tier1Cluster_Model struct {
	Value types.Bool `tfsdk:"value"`
}

// token_ttl
type TokenTtl_Model struct {
	Seconds types.Int64 `tfsdk:"seconds"`
	Nanos   types.Int64 `tfsdk:"nanos"`
}

// namespace_scope
type NamespaceScope_Model struct {
	Scope      types.String `tfsdk:"scope"`
	Exceptions types.List   `tfsdk:"exceptions"`
}

// service_account/keys
type ServiceAccount_Keys_Model struct {
	Encoding     types.String `tfsdk:"encoding"`
	Id           types.String `tfsdk:"id"`
	PrivateKey   types.String `tfsdk:"private_key"`
	PublicKey    types.String `tfsdk:"public_key"`
	DefaultToken types.String `tfsdk:"default_token"`
}

// service_account
type ServiceAccount_Model struct {
	Description types.String                 `tfsdk:"description"`
	DisplayName types.String                 `tfsdk:"display_name"`
	Keys        []*ServiceAccount_Keys_Model `tfsdk:"keys"`
}

// namespaces/services/ports
type Namespaces_Services_Ports_Model struct {
	Number             types.Int64  `tfsdk:"number"`
	KubernetesNodePort types.Int64  `tfsdk:"kubernetes_node_port"`
	Name               types.String `tfsdk:"name"`
}

// namespaces/services/workloads/proxy
type Namespaces_Services_Workloads_Proxy_Model struct {
	Status              types.Map    `tfsdk:"status"`
	ControlPlaneAddress types.String `tfsdk:"control_plane_address"`
	EnvoyVersion        types.String `tfsdk:"envoy_version"`
	IstioVersion        types.String `tfsdk:"istio_version"`
}

// namespaces/services/workloads
type Namespaces_Services_Workloads_Model struct {
	Name    types.String                              `tfsdk:"name"`
	Proxy   Namespaces_Services_Workloads_Proxy_Model `tfsdk:"proxy"`
	Address types.String                              `tfsdk:"address"`
	IsVm    types.Bool                                `tfsdk:"is_vm"`
}

// namespaces/services
type Namespaces_Services_Model struct {
	Selector                    types.Map                              `tfsdk:"selector"`
	Ports                       []*Namespaces_Services_Ports_Model     `tfsdk:"ports"`
	NumVmEndpoints              types.Int64                            `tfsdk:"num_vm_endpoints"`
	KubernetesServiceIp         types.String                           `tfsdk:"kubernetes_service_ip"`
	Namespace                   types.String                           `tfsdk:"namespace"`
	GatewayHost                 types.Bool                             `tfsdk:"gateway_host"`
	Subsets                     types.List                             `tfsdk:"subsets"`
	MeshExternal                types.Bool                             `tfsdk:"mesh_external"`
	CanonicalName               types.String                           `tfsdk:"canonical_name"`
	KubernetesServiceFqdn       types.String                           `tfsdk:"kubernetes_service_fqdn"`
	Hostname                    types.String                           `tfsdk:"hostname"`
	Tier1GatewayHost            types.Bool                             `tfsdk:"tier1_gateway_host"`
	Workloads                   []*Namespaces_Services_Workloads_Model `tfsdk:"workloads"`
	NumHops                     types.Int64                            `tfsdk:"num_hops"`
	Name                        types.String                           `tfsdk:"name"`
	NumKubernetesEndpoints      types.Int64                            `tfsdk:"num_kubernetes_endpoints"`
	SpiffeIds                   types.List                             `tfsdk:"spiffe_ids"`
	KubernetesExternalAddresses types.List                             `tfsdk:"kubernetes_external_addresses"`
}

// namespaces
type Namespaces_Model struct {
	Name     types.String                 `tfsdk:"name"`
	Services []*Namespaces_Services_Model `tfsdk:"services"`
}

// locality
type Locality_Model struct {
	Region types.String `tfsdk:"region"`
}

type ClusterModel struct {
	Network        types.String         `tfsdk:"network"`
	Tier1Cluster   Tier1Cluster_Model   `tfsdk:"tier1_cluster"`
	Description    types.String         `tfsdk:"description"`
	Name           types.String         `tfsdk:"name"`
	TrustDomain    types.String         `tfsdk:"trust_domain"`
	Id             types.String         `tfsdk:"id"`
	State          State_Model          `tfsdk:"state"`
	Namespaces     []*Namespaces_Model  `tfsdk:"namespaces"`
	NamespaceScope NamespaceScope_Model `tfsdk:"namespace_scope"`
	DisplayName    types.String         `tfsdk:"display_name"`
	Parent         types.String         `tfsdk:"parent"`
	Locality       Locality_Model       `tfsdk:"locality"`
	TokenTtl       TokenTtl_Model       `tfsdk:"token_ttl"`
	ServiceAccount ServiceAccount_Model `tfsdk:"service_account"`
	Labels         types.Map            `tfsdk:"labels"`
}
