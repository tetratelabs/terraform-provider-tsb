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

package user

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	tsbv2 "github.com/tetrateio/api/tsb/v2"
	"github.com/tetrateio/tetrate/pkg/api"
	"github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

func (r *UserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Load state into the model
	var model userResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	user, err := r.client.GetUser(ctx, &tsbv2.GetUserRequest{Fqn: model.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Error reading User", "GetUser request failed: "+err.Error())
		return
	}

	meta, err := helpers.FromFQN(api.UserKind, model.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading User", "FQN parsing failed: "+err.Error())
		return
	}

	model.Id = types.StringValue(user.Fqn)
	model.Name = types.StringValue(meta.Name)
	model.Organization = types.StringValue(meta.Organization)
	model.DisplayName = types.StringValue(user.DisplayName)
	model.Email = types.StringValue(user.Email)
	model.FirstName = types.StringValue(user.FirstName)
	model.LastName = types.StringValue(user.LastName)
	model.LoginName = types.StringValue(user.LoginName)
	model.SourceType = types.StringValue(tsbv2.SourceType_name[int32(user.SourceType)])

	// Save model into Terraform model
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
