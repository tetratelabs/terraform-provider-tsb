package helpers

import (
	"context"
	"encoding/base64"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

// BasicAuth is an HTTP Basic credentials provider for gRPC
type BasicAuth struct {
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

// GetRequestMetadata implements credentials.PerRPCCredentials
func (b BasicAuth) GetRequestMetadata(_ context.Context, _ ...string) (map[string]string, error) {
	auth := b.Username.ValueString() + ":" + b.Password.ValueString()
	enc := base64.StdEncoding.EncodeToString([]byte(auth))
	return map[string]string{"authorization": "Basic " + enc}, nil
}

// RequireTransportSecurity implements credentials.PerRPCCredentials
func (BasicAuth) RequireTransportSecurity() bool {
	return false
}
