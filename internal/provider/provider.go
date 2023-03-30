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
	"crypto/tls"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
	"github.com/tetratelabs/terraform-provider-tsb/internal/provider/datasources/organization"
	"github.com/tetratelabs/terraform-provider-tsb/internal/provider/resources/serviceaccount"
	"github.com/tetratelabs/terraform-provider-tsb/internal/provider/resources/tenant"
	"github.com/tetratelabs/terraform-provider-tsb/internal/provider/resources/user"
	"github.com/tetratelabs/terraform-provider-tsb/internal/provider/validators"
)

// Ensure TsbProvider satisfies various provider interfaces.
var _ provider.Provider = &TsbProvider{}

// TsbProvider defines the provider implementation.
type TsbProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Schema implements provider.Provider
func (*TsbProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				MarkdownDescription: "The address that the management plane can be reached at. Must include port.",
				Required:            true,
				Validators:          []validator.String{validators.AddressIncludesPort()},
			},
			"basic_auth": schema.SingleNestedAttribute{
				Description: "The basic auth credentials to communicate with a TSB management plane.",
				Required:    true,
				Attributes: map[string]schema.Attribute{
					"username": schema.StringAttribute{
						Required: true,
					},
					"password": schema.StringAttribute{
						Required:  true,
						Sensitive: true,
					},
				},
			},
		},
	}
}

// TsbProviderModel describes the provider data model.
type TsbProviderModel struct {
	Address   types.String      `tfsdk:"address"`
	BasicAuth helpers.BasicAuth `tfsdk:"basic_auth"`
}

func (p *TsbProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "tsb"
	resp.Version = p.version
}

func (p *TsbProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data TsbProviderModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TODO: support passing in custom ca bundles
	tlsConfig := &tls.Config{}
	opts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
		grpc.WithPerRPCCredentials(data.BasicAuth),
	}

	cc, err := grpc.DialContext(ctx, data.Address.ValueString(), opts...)
	if err != nil {
		panic(err)
	}

	clients := helpers.NewClients(cc)
	resp.DataSourceData = clients
	resp.ResourceData = clients
}

func (p *TsbProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		tenant.NewResource,
		serviceaccount.NewResource,
		user.NewResource,
	}
}

func (p *TsbProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		organization.NewDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &TsbProvider{
			version: version,
		}
	}
}
