package user

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type UserModel struct {
	Name        types.String `tfsdk:"name"`
	SourceType  types.String `tfsdk:"source_type"`
	FirstName   types.String `tfsdk:"first_name"`
	LoginName   types.String `tfsdk:"login_name"`
	Id          types.String `tfsdk:"id"`
	Parent      types.String `tfsdk:"parent"`
	DisplayName types.String `tfsdk:"display_name"`
	Email       types.String `tfsdk:"email"`
	LastName    types.String `tfsdk:"last_name"`
}
