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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tsbv2 "github.com/tetrateio/api/tsb/v2"

	"github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
	"github.com/tetratelabs/terraform-provider-tsb/internal/provider/defaults"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &UserResource{}
var _ resource.ResourceWithImportState = &UserResource{}

func NewResource() resource.Resource {
	return &UserResource{}
}

// UserResource defines the resource implementation.
type UserResource struct {
	client tsbv2.TeamsClient
}

type userResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Parent      types.String `tfsdk:"parent"`
	DisplayName types.String `tfsdk:"display_name"`
	LoginName   types.String `tfsdk:"login_name"`
	FirstName   types.String `tfsdk:"first_name"`
	LastName    types.String `tfsdk:"last_name"`
	Email       types.String `tfsdk:"email"`
	SourceType  types.String `tfsdk:"source_type"`
}

func (r *UserResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

// Schema implements resource.Resource
func (*UserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Description:   "The parent ID of the user.",
				Required:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"display_name": schema.StringAttribute{
				Description: "User friendly name for the resource.",
				Optional:    true,
			},
			"email": schema.StringAttribute{
				Description: "Email for the user where alerts and other notifications will be sent.",
				Optional:    true,
			},
			"first_name": schema.StringAttribute{
				Description: "The first name of the user.",
				Optional:    true,
			},
			"last_name": schema.StringAttribute{
				Description: "The last name of the user, if any.",
				Optional:    true,
			},
			"login_name": schema.StringAttribute{
				Description: "The username used in the login credentials.",
				Required:    true,
			},
			"source_type": schema.StringAttribute{
				Computed:    true,
				Description: "Where the user comes from. It can be a local user that exists only in TSB (LOCAL) or it can be a user that has been synchronized from the Identity Provider (LDAP).",
				Default:     defaults.StringDefault(tsbv2.SourceType_name[4]), // MANUAL
			},
		},
	}
}

func (r *UserResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	clients := helpers.BuildClientsResource(ctx, req, resp)
	if resp.Diagnostics.HasError() || clients == nil {
		return
	}
	r.client = clients.Team
}
