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
