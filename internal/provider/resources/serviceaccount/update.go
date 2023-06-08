package serviceaccount

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	v2 "github.com/tetrateio/api/tsb/v2"
)

func (r *ServiceAccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var model ServiceAccountModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request := &v2.GetServiceAccountRequest{Fqn: model.Id.ValueString()}
	serviceaccount, err := r.client.GetServiceAccount(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Error updating ServiceAccount", "GetServiceAccountRequest failed: "+err.Error())
		return
	}
	updateTo := &v2.ServiceAccount{
		Description: model.Description.ValueString(),
		DisplayName: model.DisplayName.ValueString(),
		Etag:        serviceaccount.Etag,
		Fqn:         model.Id.ValueString(),
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
	}
	_, err = r.client.UpdateServiceAccount(ctx, updateTo)
	if err != nil {
		resp.Diagnostics.AddError("Error updating ServiceAccount", "UpdateServiceAccount failed: "+err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
