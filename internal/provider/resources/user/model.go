package user

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type UserModel struct {
	Id          types.String `tfsdk:"id"`
	LastName    types.String `tfsdk:"last_name"`
	LoginName   types.String `tfsdk:"login_name"`
	DisplayName types.String `tfsdk:"display_name"`
	Email       types.String `tfsdk:"email"`
	Name        types.String `tfsdk:"name"`
	Parent      types.String `tfsdk:"parent"`
	SourceType  types.String `tfsdk:"source_type"`
	FirstName   types.String `tfsdk:"first_name"`
}
