package provider

import (
	"bytes"
	"testing"
	"text/template"
)

func TestAccTeamResource(t *testing.T) {
	// orgFQN := fmt.Sprintf("organizations/%s", testAccOrganizationName)
	// serviceAccountName := fmt.Sprintf("tf_service_account_%v", time.Now().Unix())
	// serviceAccount := serviceAccountConfig{
	// 	ParentFQN:   orgFQN,
	// 	Name:        serviceAccountName,
	// 	Description: "I am a test service account created during Terraform Provider acceptance testing",
	// }
	// name := fmt.Sprintf("tf_team_%v", time.Now().Unix())
	// original := teamConfig{
	// 	Name:        name,
	// 	ParentFQN:   fmt.Sprintf("organizations/%s", testAccOrganizationName),
	// 	Description: "I am a test Team created during Terraform Provider acceptance testing",
	// 	DisplayName: "Terraform Provider Test Original",
	// 	Members:     []string{fmt.Sprintf("resource.tsb_service_account.%s.id", serviceAccountName)},
	// }
	// updated := original
	// updated.DisplayName = "Terraform Provider Test Updated"

	// resource.Test(t, resource.TestCase{
	// 	PreCheck:                 func() { testAccPreCheck(t) },
	// 	ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
	// 	Steps: []resource.TestStep{
	// 		// Create and Read testing
	// 		{
	// 			Config: buildConfig(t, serviceAccount.Block(t), original.Block(t)),
	// 			Check: resource.ComposeAggregateTestCheckFunc(
	// 				resource.TestCheckResourceAttr("tsb_team."+name, "name", original.Name),
	// 				resource.TestCheckResourceAttr("tsb_team."+name, "parent", original.ParentFQN),
	// 				resource.TestCheckResourceAttr("tsb_team."+name, "team.display_name", original.DisplayName),
	// 				resource.TestCheckResourceAttr("tsb_team."+name, "team.description", original.Description),
	// 				resource.TestCheckResourceAttr("tsb_team."+name, "team.members.0", fmt.Sprintf("%s/serviceaccounts/%s", orgFQN, serviceAccountName)),
	// 				resource.TestCheckResourceAttr("tsb_team."+name, "team.source_type", strconv.Itoa(int(tsbv2.SourceType_MANUAL))),
	// 			),
	// 		},
	// 		// ImportState testing
	// 		{
	// 			ResourceName:      "tsb_team." + name,
	// 			ImportState:       true,
	// 			ImportStateVerify: true,
	// 		},
	// 		// Update and Read testing
	// 		{
	// 			Config: buildConfig(t, serviceAccount.Block(t), updated.Block(t)),
	// 			Check: resource.ComposeAggregateTestCheckFunc(
	// 				resource.TestCheckResourceAttr("tsb_team."+name, "team.display_name", updated.DisplayName)),
	// 		},
	// 		// Delete testing automatically occurs in TestCase
	// 	},
	// })
}

type teamConfig struct {
	Name        string
	ParentFQN   string
	DisplayName string
	Description string
	Members     []string
}

func (c teamConfig) Block(t *testing.T) string {
	tmpl, _ := template.New("team").Parse(teamTmpl)
	w := bytes.NewBuffer(nil)
	if err := tmpl.Execute(w, c); err != nil {
		t.Fatal("unable to build team block: %w", err)
	}
	return w.String()
}

const teamTmpl = `
resource "tsb_team" "{{.Name}}" {
	parent = "{{.ParentFQN}}"
	name = "{{.Name}}"
	team = {
		display_name = "{{.DisplayName}}"
		description = "{{.Description}}"
		{{if .Members}}members = [{{range .Members}}{{.}}{{end}}{{end}}]
	}
}
`
