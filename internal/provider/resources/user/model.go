package user

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
// Schema path
type UserModel struct {
	SourceType  types.String `tfsdk:"source_type"`
	LastName    types.String `tfsdk:"last_name"`
	Parent      types.String `tfsdk:"parent"`
	Email       types.String `tfsdk:"email"`
	FirstName   types.String `tfsdk:"first_name"`
	Id          types.String `tfsdk:"id"`
	DisplayName types.String `tfsdk:"display_name"`
	Name        types.String `tfsdk:"name"`
	LoginName   types.String `tfsdk:"login_name"`
}
