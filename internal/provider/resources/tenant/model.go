package tenant

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type TenantModel struct {
	Id             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Parent         types.String `tfsdk:"parent"`
	SecurityDomain types.String `tfsdk:"security_domain"`
	Description    types.String `tfsdk:"description"`
	DisplayName    types.String `tfsdk:"display_name"`
}
