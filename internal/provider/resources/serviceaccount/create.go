// Copyright 2023 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package serviceaccount

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
	tsbv2 "github.com/tetrateio/api/tsb/v2"
)

func (r *ServiceAccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Load plan into the model
	var model serviceAccountResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	serviceAccount, err := r.client.CreateServiceAccount(ctx, &tsbv2.CreateServiceAccountRequest{
		Parent:      model.Parent.ValueString(),
		Name:        model.Name.ValueString(),
		KeyEncoding: tsbv2.ServiceAccount_KeyPair_Encoding(tsbv2.ServiceAccount_KeyPair_Encoding_value[model.KeyEncoding.ValueString()]),
		ServiceAccount: &tsbv2.ServiceAccount{
			DisplayName: model.DisplayName.ValueString(),
			Description: model.Description.ValueString(),
		},
	})
	if err != nil {
		resp.Diagnostics.AddError("Error creating ServiceAccount", "CreateServiceAccount request failed: "+err.Error())
		return
	}

	model.Id = types.StringValue(serviceAccount.Fqn)

	// We have to handle keys differently because they can be null (during a read)
	//
	keys, diagErr := types.ListValue(
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"id":            types.StringType,
				"public_key":    types.StringType,
				"private_key":   types.StringType,
				"encoding":      types.StringType,
				"default_token": types.StringType,
			},
		},
		lo.Map(serviceAccount.Keys, func(elem *tsbv2.ServiceAccount_KeyPair, _ int) attr.Value {
			return types.ObjectValueMust(
				map[string]attr.Type{
					"id":            types.StringType,
					"public_key":    types.StringType,
					"private_key":   types.StringType,
					"encoding":      types.StringType,
					"default_token": types.StringType,
				},
				map[string]attr.Value{
					"id":            types.StringValue(elem.Id),
					"public_key":    types.StringValue(elem.PublicKey),
					"private_key":   types.StringValue(elem.PrivateKey),
					"encoding":      types.StringValue(tsbv2.ServiceAccount_KeyPair_Encoding_name[int32(elem.Encoding)]),
					"default_token": types.StringValue(elem.DefaultToken),
				},
			)
		}),
	)
	resp.Diagnostics.Append(diagErr...)
	if resp.Diagnostics.HasError() {
		return
	}

	model.Keys = keys

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
