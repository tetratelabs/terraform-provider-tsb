package serviceaccount

import (
	"context"
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	v2 "github.com/tetrateio/api/tsb/v2"
	api "github.com/tetrateio/tetrate/pkg/api"
	helpers "github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

func (r *ServiceAccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var model ServiceAccountModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request := &v2.GetServiceAccountRequest{Fqn: model.Id.ValueString()}
	serviceaccount, err := r.client.GetServiceAccount(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Error reading ServiceAccount", "GetServiceAccountRequest failed: "+err.Error())
		return
	}
	meta, err := helpers.FromFQN(api.ServiceAccountKind, model.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error readingServiceAccount", "FQN parsing failed: "+err.Error())
		return
	}
	model.Id = types.StringValue(serviceaccount.Fqn)
	model.Name = types.StringValue(meta.Name)
	model.Parent = types.StringValue(helpers.ParentFQN(api.ServiceAccountKind, meta))
	model.DisplayName = types.StringValue(serviceaccount.DisplayName)
	model.Keys = func() []*Keys_Model {
		size := len(serviceaccount.Keys)
		tmp := make([]*Keys_Model, size, size)
		for i, keys := range serviceaccount.Keys {
			tmp[i] = &Keys_Model{
				DefaultToken: types.StringValue(keys.DefaultToken),
				Encoding:     types.StringValue(v2.ServiceAccount_KeyPair_Encoding_name[int32(keys.Encoding)]),
				Id:           types.StringValue(keys.Id),
				PrivateKey:   types.StringValue(keys.PrivateKey),
				PublicKey:    types.StringValue(keys.PublicKey),
			}
		}
		return tmp
	}()
	model.Description = types.StringValue(serviceaccount.Description)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
