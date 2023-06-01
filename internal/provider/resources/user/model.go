package user

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type UserModel struct {
	Id          types.String `tfsdk:"id"`
	DisplayName types.String `tfsdk:"display_name"`
	Email       types.String `tfsdk:"email"`
	Parent      types.String `tfsdk:"parent"`
	FirstName   types.String `tfsdk:"first_name"`
	LoginName   types.String `tfsdk:"login_name"`
	SourceType  types.String `tfsdk:"source_type"`
	LastName    types.String `tfsdk:"last_name"`
	Name        types.String `tfsdk:"name"`
}
