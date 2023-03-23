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

package test

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/tetratelabs/terraform-provider-tsb/internal/provider"
)

// AccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var (
	AccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"tsb": providerserver.NewProtocol6WithError(provider.New("test")()),
	}

	AccAddress          = os.Getenv("TSB_ADDRESS")
	AccUsername         = os.Getenv("TSB_USERNAME")
	AccPassword         = os.Getenv("TSB_PASSWORD")
	AccOrganizationName = os.Getenv("TSB_ORGANIZATION")
)

func AccPreCheck(t *testing.T) {
	// Verify all env vars were set
	if AccAddress == "" {
		t.Fatal("TSB_ADDRESS must be set for acceptance testing")
	}
	if AccUsername == "" {
		t.Fatal("TSB_USERNAME must be set for acceptance testing")
	}
	if AccPassword == "" {
		t.Fatal("TSB_PASSWORD must be set for acceptance testing")
	}
	if AccOrganizationName == "" {
		t.Fatal("TSB_PASSWORD must be set for acceptance testing")
	}
}

// Pass all non-provider blocks in
// Provider is automatically added to the front
func BuildConfig(t *testing.T, elems ...string) string {
	p := providerConfig{
		Address:   AccAddress,
		BasicAuth: basicAuthConfig{Username: AccUsername, Password: AccPassword},
	}
	return strings.Join(append([]string{p.Block(t)}, elems...), "")
}

type providerConfig struct {
	Address   string
	BasicAuth basicAuthConfig
}

func (c providerConfig) Block(t *testing.T) string {
	tmpl, _ := template.New("provider").Parse(providerTmpl)
	w := bytes.NewBuffer(nil)
	if err := tmpl.Execute(w, c); err != nil {
		t.Fatal("unable to build provider block: %w", err)
	}
	return w.String()
}

type basicAuthConfig struct {
	Username string
	Password string
}

const providerTmpl = `
provider "tsb" {
	address = "{{.Address}}"
	basic_auth = {
		username = "{{.BasicAuth.Username}}"
		password = "{{.BasicAuth.Password}}"
	}
}
`
