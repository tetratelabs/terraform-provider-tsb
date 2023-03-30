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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tsbv2 "github.com/tetrateio/api/tsb/v2"

	"github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &TenantResource{}
var _ resource.ResourceWithImportState = &TenantResource{}

func NewResource() resource.Resource {
	return &TenantResource{}
}

// TenantResource defines the resource implementation.
type TenantResource struct {
	client tsbv2.TenantsClient
}

type tenantResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Parent         types.String `tfsdk:"parent"`
	Description    types.String `tfsdk:"description"`
	DisplayName    types.String `tfsdk:"display_name"`
	SecurityDomain types.String `tfsdk:"security_domain"`
}

func (r *TenantResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tenant"
}

// Schema implements resource.Resource
func (*TenantResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Fully-qualified name of the resource.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Description:   "The short name for the resource to be created.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"parent": schema.StringAttribute{
				Description:   "The parent ID of the Tenant.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"display_name": schema.StringAttribute{
				Description: "User friendly name for the resource.",
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A description of the resource.",
				Optional:    true,
			},
			"security_domain": schema.StringAttribute{
				Description: "Security domains can be used to group different resources under the same security domain. Although security domain is not resource itself currently, it follows a fqn format `organizations/myorg/securitydomains/mysecuritydomain`, and a child cannot override any ancestor's security domain. Once a security domain is assigned to a _Tenant_, all the children resources will belong to that security domain in the same way a _Workspace_ belongs to a _Tenant_, a _Workspace_ will also belong to the security domain assigned to the _Tenant_. Security domains can also be used to define _Security settings Authorization rules_ in which you can allow or deny request from or to a security domain.",
				Optional:    true,
			},
		},
	}
}

func (r *TenantResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	clients := helpers.BuildClientsResource(ctx, req, resp)
	if resp.Diagnostics.HasError() || clients == nil {
		return
	}
	r.client = clients.Tenant
}
