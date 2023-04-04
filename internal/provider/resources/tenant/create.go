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
	resource "github.com/hashicorp/terraform-plugin-framework/resource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	v2 "github.com/tetrateio/api/tsb/v2"
)

func (r *TenantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model TenantModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request := &v2.CreateTenantRequest{
		Name:   model.Name.ValueString(),
		Parent: model.Parent.ValueString(),
		Tenant: &v2.Tenant{
			Description:    model.Description.ValueString(),
			DisplayName:    model.DisplayName.ValueString(),
			SecurityDomain: model.SecurityDomain.ValueString(),
		},
	}
	tenant, err := r.client.CreateTenant(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Error creating Tenant", "CreateTenant request failed: "+err.Error())
		return
	}
	model.Id = types.StringValue(tenant.Fqn)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
