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

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	tsbv2 "github.com/tetrateio/api/tsb/v2"

	"github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ datasource.DataSource = &OrganizationDataSource{}

func NewOrganizationDataSource() datasource.DataSource {
	return &OrganizationDataSource{}
}

// OrganizationDataSource defines the data source implementation.
type OrganizationDataSource struct {
	client tsbv2.OrganizationsClient
}

// Schema implements datasource.DataSource
func (*OrganizationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	s, d := tsbv2.GenSchemaGetOrganizationRequest(ctx)
	resp.Diagnostics.Append(d...)
	resp.Schema = helpers.ResourceSchemaToDatasourceSchemar(s)
}

func (d *OrganizationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization"
}

func (d *OrganizationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	clients := helpers.BuildClientsDatasource(ctx, req, resp)
	if resp.Diagnostics.HasError() || clients == nil {
		return
	}
	d.client = clients.Organization
}

func (d *OrganizationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	data := tsbv2.NewGetOrganizationRequestModelFromConfig(ctx, req.Config, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	org, err := d.client.GetOrganization(ctx, data.ToGo(ctx))
	if err != nil {
		resp.Diagnostics.AddError("Error reading Organization", "GetOrganization request failed: "+err.Error())
		return
	}

	data.LoadFromResult(ctx, org.Fqn, &tsbv2.GetOrganizationRequest{Fqn: org.Fqn})

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
