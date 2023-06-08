package serviceaccount

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	v2 "github.com/tetrateio/api/tsb/v2"
)

func ptrify[T any](v T) *T {
	return &v
}
func (r *ServiceAccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model ServiceAccountModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request := &v2.CreateServiceAccountRequest{
		Name:   model.Name.ValueString(),
		Parent: model.Parent.ValueString(),
		ServiceAccount: &v2.ServiceAccount{
			Description: model.Description.ValueString(),
			DisplayName: model.DisplayName.ValueString(),
			Keys: func() []*v2.ServiceAccount_KeyPair {
				a := make([]*v2.ServiceAccount_KeyPair, 0, len(model.Keys))
				for i, keys := range model.Keys {
					a[i] = &v2.ServiceAccount_KeyPair{
						DefaultToken: keys.DefaultToken.ValueString(),
						Encoding:     v2.ServiceAccount_KeyPair_Encoding(v2.ServiceAccount_KeyPair_Encoding_value[keys.Encoding.ValueString()]),
						Id:           keys.Id.ValueString(),
						PrivateKey:   keys.PrivateKey.ValueString(),
						PublicKey:    keys.PublicKey.ValueString(),
					}
				}
				return a
			}(),
		},
	}
	serviceaccount, err := r.client.CreateServiceAccount(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Error creating ServiceAccount", "CreateServiceAccount request failed: "+err.Error())
		return
	}
	model.Id = types.StringValue(serviceaccount.Fqn)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
