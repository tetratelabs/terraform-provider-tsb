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

package organization_test

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/tetratelabs/terraform-provider-tsb/internal/provider/test"
)

func TestAccOrganizationDataSource(t *testing.T) {
	org := orgConfig{Id: fmt.Sprintf("organizations/%s", test.AccOrganizationName), Name: test.AccOrganizationName}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { test.AccPreCheck(t) },
		ProtoV6ProviderFactories: test.AccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: test.BuildConfig(t, org.Block(t)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.tsb_organization."+test.AccOrganizationName, "id", org.Id),
					resource.TestCheckResourceAttrSet("data.tsb_organization."+test.AccOrganizationName, "display_name"),
				),
			},
		},
	})
}

type orgConfig struct {
	Name        string
	Id          string
	Description string
	DisplayName string
}

func (c orgConfig) Block(t *testing.T) string {
	tmpl, _ := template.New("org").Parse(orgTmpl)
	w := bytes.NewBuffer(nil)
	if err := tmpl.Execute(w, c); err != nil {
		t.Fatal("unable to build org block: %w", err)
	}
	return w.String()
}

const orgTmpl = `
data "tsb_organization" "{{.Name}}" {
	id = "{{.Id}}"
}
`
