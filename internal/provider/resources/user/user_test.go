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

package user_test

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/tetratelabs/terraform-provider-tsb/internal/provider/test"
)

func TestAccUserResource(t *testing.T) {
	name := fmt.Sprintf("tf_user_%v", time.Now().Unix())
	parent := fmt.Sprintf("organizations/%v", test.AccOrganizationName)
	id := fmt.Sprintf("%v/users/%v", parent, name)
	original := userConfig{
		Name:        name,
		Parent:      parent,
		DisplayName: "Terry Form",
		LoginName:   "terryfrom",
		FirstName:   "Terry",
		LastName:    "Form",
		Email:       "terry.form@tetrate.io",
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
					resource.TestCheckResourceAttr("tsb_user."+name, "id", id),
					resource.TestCheckResourceAttr("tsb_user."+name, "name", original.Name),
					resource.TestCheckResourceAttr("tsb_user."+name, "parent", original.Parent),
					resource.TestCheckResourceAttr("tsb_user."+name, "display_name", original.DisplayName),
					resource.TestCheckResourceAttr("tsb_user."+name, "login_name", original.LoginName),
					resource.TestCheckResourceAttr("tsb_user."+name, "first_name", original.FirstName),
					resource.TestCheckResourceAttr("tsb_user."+name, "last_name", original.LastName),
					resource.TestCheckResourceAttr("tsb_user."+name, "email", original.Email),
				),
			},
			// ImportState testing
			{
				ResourceName:      "tsb_user." + name,
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Update and Read testing
			{
				Config: test.BuildConfig(t, updated.Block(t)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("tsb_user."+name, "display_name", updated.DisplayName)),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

type userConfig struct {
	Name        string
	Parent      string
	DisplayName string
	LoginName   string
	FirstName   string
	LastName    string
	Email       string
}

func (c userConfig) Block(t *testing.T) string {
	tmpl, _ := template.New("user").Parse(userTmpl)
	w := bytes.NewBuffer(nil)
	if err := tmpl.Execute(w, c); err != nil {
		t.Fatal("unable to build user block: %w", err)
	}
	return w.String()
}

const userTmpl = `
resource "tsb_user" "{{.Name}}" {
	name = "{{.Name}}"
	parent = "{{.Parent}}"
	display_name = "{{.DisplayName}}"
	login_name = "{{.LoginName}}"
	first_name = "{{.FirstName}}"
	last_name = "{{.LastName}}"
	email = "{{.Email}}"
}
`
