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

package serviceaccount

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tsbv2 "github.com/tetrateio/api/tsb/v2"

	"github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
	"github.com/tetratelabs/terraform-provider-tsb/internal/provider/defaults"
)

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &ServiceAccountResource{}
var _ resource.ResourceWithImportState = &ServiceAccountResource{}

func NewResource() resource.Resource {
	return &ServiceAccountResource{}
}

// ServiceAccountResource defines the resource implementation.
type ServiceAccountResource struct {
	client tsbv2.TeamsClient
}

type serviceAccountResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Parent      types.String `tfsdk:"parent"`
	Description types.String `tfsdk:"description"`
	DisplayName types.String `tfsdk:"display_name"`
	KeyEncoding types.String `tfsdk:"key_encoding"`
	// types.List is used because keys can be null during read, try to avoid for normal resources.
	Keys types.List `tfsdk:"keys"`
}

func (r *ServiceAccountResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service_account"
}

// Schema implements resource.Resource
func (*ServiceAccountResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Description:   "The parent ID of the Service Account.",
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
			"key_encoding": schema.StringAttribute{
				Description:   "The format in which the generated key pairs will be returned. If not set keys are returned in PEM format.",
				Optional:      true,
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
				Default:       defaults.ProtoEnumDefault(tsbv2.ServiceAccount_KeyPair_Encoding_name),
			},
			"keys": schema.ListNestedAttribute{
				Computed:      true,
				Description:   "Keys associated with the service account. A default key-pair is automatically created when the Service Account is created. Note that TSB does not store the private keys, so it is up to the client to store the returned private keys securely, as they are only returned once after creation. Additional keys can be added (and deleted) by using the corresponding key management APIs. +protoc-gen-terraform:computed",
				PlanModifiers: []planmodifier.List{listplanmodifier.UseStateForUnknown()},
				NestedObject: schema.NestedAttributeObject{Attributes: map[string]schema.Attribute{
					"default_token": schema.StringAttribute{
						Description: "A default access token that can be used to authenticate to TSB on behalf of the service account. TSB does not store this token and it is only returned when a service account key is created, similar to the private key. It is up to the client to store the token for future use or to use the TSB CLI to generate new tokens as explained in: https://docs.tetrate.io/service-bridge/latest/en-us/howto/service-accounts",
						Computed:    true,
						Sensitive:   true,
					},
					"encoding": schema.StringAttribute{
						Description: "Format in which the public and private keys are encoded. By default keys are returned in PEM format.",
						Computed:    true,
					},
					"id": schema.StringAttribute{
						Description: "Unique identifier for this key-pair. This should be used as the `kid` (key id) when generating JWT tokens that are signed with this key-pair.",
						Computed:    true,
					},
					"private_key": schema.StringAttribute{
						Description: "The encoded private key associated with the service account. TSB does not store the private key and it is up to the client to store it safely. The encoding format is determined by the `encoding` field.",
						Computed:    true,
						Sensitive:   true,
					},
					"public_key": schema.StringAttribute{
						Description: "The encoded public key associated with the service account. The encoding format is determined by the `encoding` field.",
						Computed:    true,
					},
				}},
			},
		},
	}
}

func (r *ServiceAccountResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	clients := helpers.BuildClientsResource(ctx, req, resp)
	if resp.Diagnostics.HasError() || clients == nil {
		return
	}
	r.client = clients.Team
}
