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
	tsbv2 "github.com/tetrateio/api/tsb/v2"
	v2 "github.com/tetrateio/api/tsb/v2"
)

func (r *UserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Load plan into the model
	var model userResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// We need to run a get to retreive the latest etag and set fqn as its not going to be in the plan (only state)
	user, err := r.client.GetUser(ctx, &tsbv2.GetUserRequest{Fqn: model.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Error updating User", "GetUser request failed: "+err.Error())
		return
	}

	// Do the actual update
	_, err = r.client.UpdateUser(ctx, &v2.User{
		Fqn:         model.Id.ValueString(),
		Etag:        user.Etag,
		DisplayName: model.DisplayName.ValueString(),
		LoginName:   model.LoginName.ValueString(),
		FirstName:   model.FirstName.ValueString(),
		LastName:    model.LastName.ValueString(),
		Email:       model.Email.ValueString(),
		SourceType:  tsbv2.SourceType(tsbv2.SourceType_value[model.SourceType.ValueString()]),
	})
	if err != nil {
		resp.Diagnostics.AddError("Error updating User", "UpdateUser request failed: "+err.Error())
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
