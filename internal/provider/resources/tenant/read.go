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
	api "github.com/tetrateio/tetrate/pkg/api"
	helpers "github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

func (r *TenantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var model TenantModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	request := &v2.GetTenantRequest{Fqn: model.Id.ValueString()}
	tenant, err := r.client.GetTenant(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("Error reading Tenant", "GetTenantRequest failed: "+err.Error())
		return
	}
	meta, err := helpers.FromFQN(api.TenantKind, model.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error readingTenant", "FQN parsing failed: "+err.Error())
		return
	}
	model.Id = types.StringValue(tenant.Fqn)
	model.Name = types.StringValue(meta.Name)
	model.Parent = types.StringValue(helpers.ParentFQN(api.TenantKind, meta))
	model.DisplayName = types.StringValue(tenant.DisplayName)
	model.SecurityDomain = types.StringValue(tenant.SecurityDomain)
	model.Description = types.StringValue(tenant.Description)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
