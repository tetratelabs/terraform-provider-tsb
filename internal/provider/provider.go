// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure tsbProvider satisfies the provider.Provider interface.
var _ provider.Provider = (*tsbProvider)(nil)

// version is set at build time via -ldflags; "dev" by default.
var version = "dev"

// tsbProvider is the Terraform provider for Tetrate Service Bridge (TSB).
type tsbProvider struct{}

// New returns a constructor for the TSB provider, suitable for providerserver.Serve.
func New() func() provider.Provider {
	return func() provider.Provider {
		return &tsbProvider{}
	}
}

// providerModel maps the provider configuration block.
type providerModel struct {
	Endpoint     types.String `tfsdk:"endpoint"`
	Token        types.String `tfsdk:"token"`
	Username     types.String `tfsdk:"username"`
	Password     types.String `tfsdk:"password"`
	Organization types.String `tfsdk:"organization"`
	CustomCAFile types.String `tfsdk:"ca_file"`
	TLSDisabled  types.Bool   `tfsdk:"tls_disabled"`
	TLSInsecure  types.Bool   `tfsdk:"tls_insecure"`
}

func (p *tsbProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "tsb"
	resp.Version = version
}

func (p *tsbProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage Tetrate Service Bridge (TSB) resources via the TSB v2 API.",
		Attributes: map[string]schema.Attribute{
			"endpoint": schema.StringAttribute{
				Optional:    true,
				Description: "TSB gRPC endpoint (host:port). May also be set via the TSB_ENDPOINT environment variable.",
			},
			"token": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "TSB API token sent in the x-tetrate-token header. Takes precedence over username/password. May also be set via the TSB_TOKEN environment variable.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Description: "Username for HTTP Basic authentication, used when `token` is not set. May also be set via the TSB_USERNAME environment variable.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Password for HTTP Basic authentication, used together with `username`. May also be set via the TSB_PASSWORD environment variable.",
			},
			"organization": schema.StringAttribute{
				Optional:    true,
				Description: "TSB organization the resources belong to. May also be set via the TSB_ORG environment variable.",
			},
			"ca_file": schema.StringAttribute{
				Optional:    true,
				Description: "Path to a PEM CA bundle used to verify the TSB server certificate. May also be set via the TSB_CA_FILE environment variable.",
			},
			"tls_disabled": schema.BoolAttribute{
				Optional:    true,
				Description: "Disable transport security entirely (plaintext). May also be set via the TSB_TLS_DISABLED environment variable.",
			},
			"tls_insecure": schema.BoolAttribute{
				Optional:    true,
				Description: "Keep TLS enabled but skip server certificate verification. May also be set via the TSB_TLS_INSECURE environment variable.",
			},
		},
	}
}

func (p *tsbProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data providerModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	cfg := Config{
		Endpoint:     stringWithEnv(data.Endpoint, "TSB_ENDPOINT"),
		Token:        stringWithEnv(data.Token, "TSB_TOKEN"),
		Username:     stringWithEnv(data.Username, "TSB_USERNAME"),
		Password:     stringWithEnv(data.Password, "TSB_PASSWORD"),
		Organization: stringWithEnv(data.Organization, "TSB_ORG"),
		CustomCAFile: stringWithEnv(data.CustomCAFile, "TSB_CA_FILE"),
		TLSDisabled:  boolWithEnv(data.TLSDisabled, "TSB_TLS_DISABLED"),
		TLSInsecure:  boolWithEnv(data.TLSInsecure, "TSB_TLS_INSECURE"),
	}

	if cfg.Endpoint == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("endpoint"),
			"Missing TSB endpoint",
			"The provider requires a TSB endpoint. Set the `endpoint` attribute or the TSB_ENDPOINT environment variable.",
		)
	}
	// Require either a token or a username/password pair.
	if cfg.Token == "" && (cfg.Username == "" || cfg.Password == "") {
		resp.Diagnostics.AddError(
			"Missing TSB credentials",
			"The provider requires either a `token` (TSB_TOKEN) or both `username` and `password` "+
				"(TSB_USERNAME / TSB_PASSWORD) for authentication.",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	conn, err := cfg.dial(ctx)
	if err != nil {
		resp.Diagnostics.AddError("Unable to connect to TSB", err.Error())
		return
	}

	c := &client{conn: conn, config: cfg}
	resp.ResourceData = c
	resp.DataSourceData = c
}

func (p *tsbProvider) Resources(_ context.Context) []func() resource.Resource {
	// Generated resources plus hand-written ones that don't fit the CRUD-quartet
	// generator (e.g. the Get/Set-based access bindings).
	return append(Resources(), NewAccessBindingResource)
}

func (p *tsbProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// stringWithEnv returns the configured value, falling back to the named env var.
func stringWithEnv(v types.String, env string) string {
	if !v.IsNull() && !v.IsUnknown() {
		return v.ValueString()
	}
	return os.Getenv(env)
}

// boolWithEnv returns the configured value, falling back to the named env var
// (parsed as a bool; unparseable or unset env yields false).
func boolWithEnv(v types.Bool, env string) bool {
	if !v.IsNull() && !v.IsUnknown() {
		return v.ValueBool()
	}
	b, _ := strconv.ParseBool(os.Getenv(env))
	return b
}
