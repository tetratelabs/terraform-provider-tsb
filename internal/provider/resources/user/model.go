package user

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
type UserModel struct {
	DisplayName types.String `tfsdk:"display_name"`
	FirstName   types.String `tfsdk:"first_name"`
	SourceType  types.String `tfsdk:"source_type"`
	Id          types.String `tfsdk:"id"`
	Email       types.String `tfsdk:"email"`
	LoginName   types.String `tfsdk:"login_name"`
	LastName    types.String `tfsdk:"last_name"`
	Name        types.String `tfsdk:"name"`
	Parent      types.String `tfsdk:"parent"`
}
