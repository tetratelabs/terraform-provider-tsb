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
	"github.com/hashicorp/terraform-plugin-framework/types"
	v2 "github.com/tetrateio/api/tsb/types/v2"
	tsbv2 "github.com/tetrateio/api/tsb/v2"
	"github.com/tetrateio/tetrate/pkg/api"
	"github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

func (r *TenantResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Load plan into the model
	var model tenantResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}


	tenant, err := r.client.CreateTenant(ctx, &tsbv2.CreateTenantRequest{
		Parent: helpers.FQN(api.OrganizationKind, &v2.ObjectMeta{Name: model.Organization.ValueString()}),
		Name:   model.Name.ValueString(),
		Tenant: &tsbv2.Tenant{
			DisplayName:    model.DisplayName.ValueString(),
			Description:    model.Description.ValueString(),
			SecurityDomain: model.SecurityDomain.ValueString(),
		},
	})
	if err != nil {
		resp.Diagnostics.AddError("Error creating Tenant", "CreateTenant request failed: "+err.Error())
		return
	}

	model.Id = types.StringValue(tenant.Fqn)
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
