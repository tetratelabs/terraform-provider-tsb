resource "tsb_tenant" "example" {
  organization    = "tetrate"
  name            = "tenantone"
  display_name    = "Tenant One"
  description     = "I'm a tenant"
  security_domain = "production"
}