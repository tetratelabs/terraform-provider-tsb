package cluster

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
// tier1_cluster
type Tier1Cluster_Model struct {
	Value types.Bool `tfsdk:"value"`
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
	Description types.String                 `tfsdk:"description"`
	DisplayName types.String                 `tfsdk:"display_name"`
	Keys        []*ServiceAccount_Keys_Model `tfsdk:"keys"`
}

// state/discovered_locality
type State_DiscoveredLocality_Model struct {
	Region types.String `tfsdk:"region"`
}

// state/last_sync_time
type State_LastSyncTime_Model struct {
	Nanos   types.Int64 `tfsdk:"nanos"`
	Seconds types.Int64 `tfsdk:"seconds"`
}

// state
type State_Model struct {
	DiscoveredLocality State_DiscoveredLocality_Model `tfsdk:"discovered_locality"`
	IstioVersions      types.List                     `tfsdk:"istio_versions"`
	LastSyncTime       State_LastSyncTime_Model       `tfsdk:"last_sync_time"`
	Provider           types.String                   `tfsdk:"provider"`
	TsbCpVersion       types.String                   `tfsdk:"tsb_cp_version"`
	XcpVersion         types.String                   `tfsdk:"xcp_version"`
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
	IsVm    types.Bool                                `tfsdk:"is_vm"`
	Name    types.String                              `tfsdk:"name"`
	Proxy   Namespaces_Services_Workloads_Proxy_Model `tfsdk:"proxy"`
	Address types.String                              `tfsdk:"address"`
}

// namespaces/services/ports
type Namespaces_Services_Ports_Model struct {
	KubernetesNodePort types.Int64  `tfsdk:"kubernetes_node_port"`
	Name               types.String `tfsdk:"name"`
	Number             types.Int64  `tfsdk:"number"`
}

// namespaces/services
type Namespaces_Services_Model struct {
	CanonicalName               types.String                           `tfsdk:"canonical_name"`
	MeshExternal                types.Bool                             `tfsdk:"mesh_external"`
	GatewayHost                 types.Bool                             `tfsdk:"gateway_host"`
	Workloads                   []*Namespaces_Services_Workloads_Model `tfsdk:"workloads"`
	KubernetesServiceIp         types.String                           `tfsdk:"kubernetes_service_ip"`
	Ports                       []*Namespaces_Services_Ports_Model     `tfsdk:"ports"`
	KubernetesServiceFqdn       types.String                           `tfsdk:"kubernetes_service_fqdn"`
	Selector                    types.Map                              `tfsdk:"selector"`
	Hostname                    types.String                           `tfsdk:"hostname"`
	NumVmEndpoints              types.Int64                            `tfsdk:"num_vm_endpoints"`
	Namespace                   types.String                           `tfsdk:"namespace"`
	SpiffeIds                   types.List                             `tfsdk:"spiffe_ids"`
	Tier1GatewayHost            types.Bool                             `tfsdk:"tier1_gateway_host"`
	NumKubernetesEndpoints      types.Int64                            `tfsdk:"num_kubernetes_endpoints"`
	Name                        types.String                           `tfsdk:"name"`
	Subsets                     types.List                             `tfsdk:"subsets"`
	NumHops                     types.Int64                            `tfsdk:"num_hops"`
	KubernetesExternalAddresses types.List                             `tfsdk:"kubernetes_external_addresses"`
}

// namespaces
type Namespaces_Model struct {
	Services []*Namespaces_Services_Model `tfsdk:"services"`
	Name     types.String                 `tfsdk:"name"`
}

// token_ttl
type TokenTtl_Model struct {
	Seconds types.Int64 `tfsdk:"seconds"`
	Nanos   types.Int64 `tfsdk:"nanos"`
}

// namespace_scope
type NamespaceScope_Model struct {
	Exceptions types.List   `tfsdk:"exceptions"`
	Scope      types.String `tfsdk:"scope"`
}

type ClusterModel struct {
	Locality       Locality_Model       `tfsdk:"locality"`
	Labels         types.Map            `tfsdk:"labels"`
	NamespaceScope NamespaceScope_Model `tfsdk:"namespace_scope"`
	Parent         types.String         `tfsdk:"parent"`
	ServiceAccount ServiceAccount_Model `tfsdk:"service_account"`
	State          State_Model          `tfsdk:"state"`
	Name           types.String         `tfsdk:"name"`
	TrustDomain    types.String         `tfsdk:"trust_domain"`
	Network        types.String         `tfsdk:"network"`
	Id             types.String         `tfsdk:"id"`
	Description    types.String         `tfsdk:"description"`
	Tier1Cluster   Tier1Cluster_Model   `tfsdk:"tier1_cluster"`
	Namespaces     []*Namespaces_Model  `tfsdk:"namespaces"`
	TokenTtl       TokenTtl_Model       `tfsdk:"token_ttl"`
	DisplayName    types.String         `tfsdk:"display_name"`
}
