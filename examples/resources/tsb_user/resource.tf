resource "tsb_user" "example" {
  parent       = "organizations/tetrate"
  name         = "terryform"
  display_name = "Terry Form"
  login_name   = "terryform"
  first_name   = "Terry"
  last_name    = "Form"
  email        = "terry.form@tetrate.io"
}
