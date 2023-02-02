resource "tsb_tenant" "example" {
  parent = "organizations/tetrate"
  name             = "tenantone"
  tenant = {
  display_name     = "Tenant One"
  description      = "I'm a tenant"
  }
}