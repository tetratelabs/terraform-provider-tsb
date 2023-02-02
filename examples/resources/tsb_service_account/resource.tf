resource "tsb_service_account" "example" {
  parent = "organizations/tetrate"
  name             = "serviceaccountone"
  service_account = {
    display_name     = "Service Account One"
    description      = "I'm a service account"
  }
}