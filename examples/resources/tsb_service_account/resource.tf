resource "tsb_service_account" "example" {
  name         = "ciautomation"
  parent       = "organizations/tetrate"
  display_name = "CI Automation"
  description  = "Used by our CI tool to create resources"
  key_encoding = "JWT"
}
