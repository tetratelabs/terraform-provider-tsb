resource "tsb_team" "example" {
  parent = "organizations/tetrate"
  name             = "teamone"
  team = {
    display_name     = "Team One"
    description      = "I'm a team"
    members          = ["organizations/tetrate/serviceaccounts/someserviceaccount"]
  }
}