package user

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type UserModel struct {
	Id          types.String `tfsdk:"id"`
	SourceType  types.String `tfsdk:"source_type"`
	Email       types.String `tfsdk:"email"`
	Name        types.String `tfsdk:"name"`
	Parent      types.String `tfsdk:"parent"`
	DisplayName types.String `tfsdk:"display_name"`
	FirstName   types.String `tfsdk:"first_name"`
	LastName    types.String `tfsdk:"last_name"`
	LoginName   types.String `tfsdk:"login_name"`
}
