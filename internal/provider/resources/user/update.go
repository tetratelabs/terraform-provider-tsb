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
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	v2 "github.com/tetrateio/api/tsb/v2"
)

func (r *UserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var model UserModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request := &v2.GetUserRequest{Fqn: model.Id.ValueString()}
	user, err := r.client.GetUser(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Error updating User", "GetUserRequest failed: "+err.Error())
		return
	}
	updateTo := &v2.User{
		DisplayName: model.DisplayName.ValueString(),
		Email:       model.Email.ValueString(),
		Etag:        user.Etag,
		FirstName:   model.FirstName.ValueString(),
		Fqn:         model.Id.ValueString(),
		LastName:    model.LastName.ValueString(),
		LoginName:   model.LoginName.ValueString(),
		SourceType:  v2.SourceType(v2.SourceType_value[model.SourceType.ValueString()]),
	}
	_, err = r.client.UpdateUser(ctx, updateTo)
	if err != nil {
		resp.Diagnostics.AddError("Error updating User", "UpdateUser failed: "+err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
