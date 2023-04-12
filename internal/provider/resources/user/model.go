package user

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type UserModel struct {
	FirstName   types.String `tfsdk:"first_name"`
	SourceType  types.String `tfsdk:"source_type"`
	DisplayName types.String `tfsdk:"display_name"`
	Email       types.String `tfsdk:"email"`
	LoginName   types.String `tfsdk:"login_name"`
	Name        types.String `tfsdk:"name"`
	Id          types.String `tfsdk:"id"`
	LastName    types.String `tfsdk:"last_name"`
	Parent      types.String `tfsdk:"parent"`
}
