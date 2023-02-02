package provider

import (
	"bytes"
	"fmt"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccOrganizationDataSource(t *testing.T) {
	org := orgConfig{Name: testAccOrganizationName, Fqn: fmt.Sprintf("organizations/%s", testAccOrganizationName)}
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: buildConfig(t, org.Block(t)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.tsb_organization."+org.Name, "id", org.Fqn),
					resource.TestCheckResourceAttr("data.tsb_organization."+org.Name, "fqn", org.Fqn),
				),
			},
		},
	})
}

type orgConfig struct {
	Name string
	Fqn  string
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
	fqn = "{{.Fqn}}"
}
`
