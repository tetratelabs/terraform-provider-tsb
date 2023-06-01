package user

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type UserModel struct {
	LoginName   types.String `tfsdk:"login_name"`
	DisplayName types.String `tfsdk:"display_name"`
	FirstName   types.String `tfsdk:"first_name"`
	Id          types.String `tfsdk:"id"`
	SourceType  types.String `tfsdk:"source_type"`
	LastName    types.String `tfsdk:"last_name"`
	Name        types.String `tfsdk:"name"`
	Parent      types.String `tfsdk:"parent"`
	Email       types.String `tfsdk:"email"`
}
