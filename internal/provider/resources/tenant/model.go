package tenant

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
// config_generation_metadata
type ConfigGenerationMetadata_Model struct {
	Annotations types.Map `tfsdk:"annotations"`
	Labels      types.Map `tfsdk:"labels"`
}

type TenantModel struct {
	Id                        types.String                   `tfsdk:"id"`
	Name                      types.String                   `tfsdk:"name"`
	Parent                    types.String                   `tfsdk:"parent"`
	SecurityDomain            types.String                   `tfsdk:"security_domain"`
	ConfigGenerationMetadata  ConfigGenerationMetadata_Model `tfsdk:"config_generation_metadata"`
	DeletionProtectionEnabled types.Bool                     `tfsdk:"deletion_protection_enabled"`
	Description               types.String                   `tfsdk:"description"`
	DisplayName               types.String                   `tfsdk:"display_name"`
}
