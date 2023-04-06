package user

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type UserModel struct {
	DisplayName types.String `tfsdk:"display_name"`
	LastName    types.String `tfsdk:"last_name"`
	Parent      types.String `tfsdk:"parent"`
	SourceType  types.String `tfsdk:"source_type"`
	Name        types.String `tfsdk:"name"`
	Id          types.String `tfsdk:"id"`
	LoginName   types.String `tfsdk:"login_name"`
	FirstName   types.String `tfsdk:"first_name"`
	Email       types.String `tfsdk:"email"`
}
