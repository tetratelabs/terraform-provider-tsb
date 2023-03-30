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

package serviceaccount_test

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/tetratelabs/terraform-provider-tsb/internal/provider/test"
)

func TestAccServiceAccountResource(t *testing.T) {
	name := fmt.Sprintf("tf_service_account_%v", time.Now().Unix())
	parent := fmt.Sprintf("organizations/%v", test.AccOrganizationName)
	id := fmt.Sprintf("%v/serviceaccounts/%v", parent, name)
	original := serviceAccountConfig{
		Name:        name,
		Parent:      parent,
		Description: "I am a test ServiceAccount created during Terraform Provider acceptance testing",
		DisplayName: "Terraform Provider Test Original",
		KeyEncoding: "JWK",
	}
	updated := original
	updated.DisplayName = "Terraform Provider Test Updated"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { test.AccPreCheck(t) },
		ProtoV6ProviderFactories: test.AccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: test.BuildConfig(t, original.Block(t)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tsb_service_account."+name, "id", id),
					resource.TestCheckResourceAttr("tsb_service_account."+name, "name", original.Name),
					resource.TestCheckResourceAttr("tsb_service_account."+name, "parent", original.Parent),
					resource.TestCheckResourceAttr("tsb_service_account."+name, "display_name", original.DisplayName),
					resource.TestCheckResourceAttr("tsb_service_account."+name, "description", original.Description),
					resource.TestCheckResourceAttr("tsb_service_account."+name, "key_encoding", original.KeyEncoding),
					resource.TestCheckResourceAttrSet("tsb_service_account."+name, "keys.0.id"),
					resource.TestCheckResourceAttrSet("tsb_service_account."+name, "keys.0.private_key"),
					resource.TestCheckResourceAttrSet("tsb_service_account."+name, "keys.0.public_key"),
					resource.TestCheckResourceAttr("tsb_service_account."+name, "keys.0.encoding", original.KeyEncoding),
					resource.TestCheckResourceAttrSet("tsb_service_account."+name, "keys.0.default_token"),
				),
			},
			// ImportState testing
			{
				ResourceName:            "tsb_service_account." + name,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"keys.0.default_token", "keys.0.encoding", "keys.0.id", "keys.0.private_key", "keys.0.public_key", "keys.#", "keys.0.%", "key_encoding"},
			},
			// Update and Read testing
			{
				Config: test.BuildConfig(t, updated.Block(t)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tsb_service_account."+name, "display_name", updated.DisplayName),
					resource.TestCheckResourceAttrSet("tsb_service_account."+name, "keys.0.private_key"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

type serviceAccountConfig struct {
	Name        string
	Parent      string
	Description string
	DisplayName string
	KeyEncoding string
}

func (c serviceAccountConfig) Block(t *testing.T) string {
	tmpl, _ := template.New("serviceAccount").Parse(serviceAccountTmpl)
	w := bytes.NewBuffer(nil)
	if err := tmpl.Execute(w, c); err != nil {
		t.Fatal("unable to build serviceAccount block: %w", err)
	}
	return w.String()
}

const serviceAccountTmpl = `
resource "tsb_service_account" "{{.Name}}" {
	name = "{{.Name}}"
	parent = "{{.Parent}}"
	display_name = "{{.DisplayName}}"
	description = "{{.Description}}"
	key_encoding = "{{.KeyEncoding}}"
}
`
