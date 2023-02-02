package provider

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var (
	testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"tsb": providerserver.NewProtocol6WithError(New("test")()),
	}

	testAccAddress          = os.Getenv("TSB_ADDRESS")
	testAccUsername         = os.Getenv("TSB_USERNAME")
	testAccPassword         = os.Getenv("TSB_PASSWORD")
	testAccOrganizationName = os.Getenv("TSB_ORGANIZATION")
)

func testAccPreCheck(t *testing.T) {
	// Verify all env vars were set
	if testAccAddress == "" {
		t.Fatal("TSB_ADDRESS must be set for acceptance testing")
	}
	if testAccUsername == "" {
		t.Fatal("TSB_USERNAME must be set for acceptance testing")
	}
	if testAccPassword == "" {
		t.Fatal("TSB_PASSWORD must be set for acceptance testing")
	}
	if testAccOrganizationName == "" {
		t.Fatal("TSB_PASSWORD must be set for acceptance testing")
	}
}

// Pass all non-provider blocks in
// Provider is automatically added to the front
func buildConfig(t *testing.T, elems ...string) string {
	p := providerConfig{
		Address:   testAccAddress,
		BasicAuth: basicAuthConfig{Username: testAccUsername, Password: testAccPassword},
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
