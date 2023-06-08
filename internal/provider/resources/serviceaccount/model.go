package serviceaccount

import types "github.com/hashicorp/terraform-plugin-framework/types"

// tfsdk typed model definition
// keys
type Keys_Model struct {
	PrivateKey   types.String `tfsdk:"private_key"`
	PublicKey    types.String `tfsdk:"public_key"`
	DefaultToken types.String `tfsdk:"default_token"`
	Encoding     types.String `tfsdk:"encoding"`
	Id           types.String `tfsdk:"id"`
}

type ServiceAccountModel struct {
	Id          types.String `tfsdk:"id"`
	Keys        types.List   `tfsdk:"keys"` // Keys_Model
	Name        types.String `tfsdk:"name"`
	Parent      types.String `tfsdk:"parent"`
	Description types.String `tfsdk:"description"`
	DisplayName types.String `tfsdk:"display_name"`
}
