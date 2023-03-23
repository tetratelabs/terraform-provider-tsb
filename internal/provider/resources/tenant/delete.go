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

package tenant

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	tsbv2 "github.com/tetrateio/api/tsb/v2"
)

func (r *TenantResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Load state into the model
	var model tenantResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if _, err := r.client.DeleteTenant(ctx, &tsbv2.DeleteTenantRequest{Fqn: model.Id.ValueString()}); err != nil {
		resp.Diagnostics.AddError("Error deleting Tenant", "DeleteTenant request failed: "+err.Error())
		return
	}
}
