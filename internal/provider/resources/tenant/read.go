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
	types "github.com/hashicorp/terraform-plugin-framework/types"
	tsbv2 "github.com/tetrateio/api/tsb/v2"
	"github.com/tetrateio/tetrate/pkg/api"
	"github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

func (r *TenantResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Load state into the model
	var model tenantResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tenant, err := r.client.GetTenant(ctx, &tsbv2.GetTenantRequest{Fqn: model.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Error reading Tenant", "GetTenant request failed: "+err.Error())
		return
	}

	meta, err := helpers.FromFQN(api.TenantKind, model.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Error reading Tenant", "FQN parsing failed: "+err.Error())
		return
	}

	model.Id = types.StringValue(tenant.Fqn)
	model.Name = types.StringValue(meta.Name)
	model.Parent = types.StringValue(helpers.ParentFQN(api.TenantKind, meta))
	model.Description = types.StringValue(tenant.Description)
	model.DisplayName = types.StringValue(tenant.DisplayName)
	model.SecurityDomain = types.StringValue(tenant.SecurityDomain)

	// Save model into Terraform model
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
