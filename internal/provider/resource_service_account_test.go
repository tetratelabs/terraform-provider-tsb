package provider

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccServiceAccountResource(t *testing.T) {
	name := fmt.Sprintf("tf_service_account_%v", time.Now().Unix())
	original := serviceAccountConfig{
		Name:        name,
		ParentFQN:   fmt.Sprintf("organizations/%s", testAccOrganizationName),
		Description: "I am a test ServiceAccount created during Terraform Provider acceptance testing",
		DisplayName: "Terraform Provider Test Original",
	}
	fqn := fmt.Sprintf("%v/serviceaccounts/%v", original.ParentFQN, original.Name)
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
					resource.TestCheckResourceAttr("tsb_service_account."+name, "name", original.Name),
					resource.TestCheckResourceAttr("tsb_service_account."+name, "parent", original.ParentFQN),
					resource.TestCheckResourceAttr("tsb_service_account."+name, "service_account.fqn", fqn),
					resource.TestCheckResourceAttr("tsb_service_account."+name, "service_account.display_name", original.DisplayName),
					resource.TestCheckResourceAttr("tsb_service_account."+name, "service_account.description", original.Description),
				),
			},
			// ImportState testing
			// {
			// 	ResourceName:      "tsb_service_account." + name,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
			// Update and Read testing
			// {
			// 	Config: buildConfig(t, updated.Block(t)),
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		resource.TestCheckResourceAttr("tsb_service_account."+name, "display_name", updated.DisplayName)),
			// },
			// Delete testing automatically occurs in TestCase
		},
	})
}

type serviceAccountConfig struct {
	Name        string
	ParentFQN   string
	DisplayName string
	Description string
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
	parent = "{{.ParentFQN}}"
	name = "{{.Name}}"
	service_account = {
		display_name = "{{.DisplayName}}"
		description = "{{.Description}}"
	}
}
`
