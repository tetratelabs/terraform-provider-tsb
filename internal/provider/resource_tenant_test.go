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
	"bytes"
	"fmt"
	"testing"
	"text/template"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTenantResource(t *testing.T) {
	name := fmt.Sprintf("tf_tenant_%v", time.Now().Unix())
	original := tenantConfig{
		Name:           name,
		ParentFQN:      fmt.Sprintf("organizations/%s", testAccOrganizationName),
		Description:    "I am a test Tenant created during Terraform Provider acceptance testing",
		DisplayName:    "Terraform Provider Test Original",
		SecurityDomain: fmt.Sprintf("organizations/%s/securitydomains/yolo", testAccOrganizationName),
	}
	fqn := fmt.Sprintf("%v/tenants/%v", original.ParentFQN, original.Name)
	updated := original
	updated.DisplayName = "Terraform Provider Test Updated"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: buildConfig(t, original.Block(t)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tsb_tenant."+name, "name", original.Name),
					resource.TestCheckResourceAttr("tsb_tenant."+name, "parent", original.ParentFQN),
					resource.TestCheckResourceAttr("tsb_tenant."+name, "tenant.fqn", fqn),
					resource.TestCheckResourceAttr("tsb_tenant."+name, "tenant.display_name", original.DisplayName),
					resource.TestCheckResourceAttr("tsb_tenant."+name, "tenant.description", original.Description),
					resource.TestCheckResourceAttr("tsb_tenant."+name, "tenant.security_domain", original.SecurityDomain),
				),
			},
			// ImportState testing
			{
				ResourceName:      "tsb_tenant." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: buildConfig(t, updated.Block(t)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tsb_tenant."+name, "tenant.display_name", updated.DisplayName)),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

type tenantConfig struct {
	Name           string
	ParentFQN      string
	DisplayName    string
	Description    string
	SecurityDomain string
}

func (c tenantConfig) Block(t *testing.T) string {
	tmpl, _ := template.New("tenant").Parse(tenantTmpl)
	w := bytes.NewBuffer(nil)
	if err := tmpl.Execute(w, c); err != nil {
		t.Fatal("unable to build tenant block: %w", err)
	}
	return w.String()
}

const tenantTmpl = `
resource "tsb_tenant" "{{.Name}}" {
	parent = "{{.ParentFQN}}"
	name = "{{.Name}}"
	tenant = {
		display_name = "{{.DisplayName}}"
		description = "{{.Description}}"
		security_domain = "{{.SecurityDomain}}"
	}
}
`
