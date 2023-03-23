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

package organization

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	types "github.com/hashicorp/terraform-plugin-framework/types"
	tsbv2 "github.com/tetrateio/api/tsb/v2"
)

func (d *OrganizationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	// Load request into the model
	var model organizationDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	org, err := d.client.GetOrganization(ctx, &tsbv2.GetOrganizationRequest{Fqn: model.Id.ValueString()})
	if err != nil {
		resp.Diagnostics.AddError("Error reading Organization", "GetOrganization request failed: "+err.Error())
		return
	}

	model.Id = types.StringValue(org.Fqn)
	model.DisplayName = types.StringValue(org.DisplayName)

	// Save model into Terraform model
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
