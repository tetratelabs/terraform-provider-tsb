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

	"github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	tsbv2 "github.com/tetrateio/api/tsb/v2"
	"github.com/tetrateio/tetrate/pkg/api"
	"github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

func (r *ServiceAccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Load state into the model
	var model serviceAccountResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	serviceAccount, err := r.client.GetServiceAccount(
		ctx,
		&tsbv2.GetServiceAccountRequest{
			Fqn:         model.Id.ValueString(),
			KeyEncoding: tsbv2.ServiceAccount_KeyPair_Encoding(tsbv2.ServiceAccount_KeyPair_Encoding_value[model.KeyEncoding.ValueString()]),
		},
	)
	if err != nil {
		resp.Diagnostics.AddError("Error reading ServiceAccount", "GetServiceAccount request failed: "+err.Error())
		return
	}

	meta, err := helpers.FromFQN(api.ServiceAccountKind, model.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading ServiceAccount", "FQN parsing failed: "+err.Error())
		return
	}

	model.Id = types.StringValue(serviceAccount.Fqn)
	model.Name = types.StringValue(meta.Name)
	model.Parent = types.StringValue(helpers.ParentFQN(api.ServiceAccountKind, meta))
	model.Description = types.StringValue(serviceAccount.Description)
	model.DisplayName = types.StringValue(serviceAccount.DisplayName)
	// Keys cannot be read, they are only returned on create so don't touch them!

	// Save model into Terraform model
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
