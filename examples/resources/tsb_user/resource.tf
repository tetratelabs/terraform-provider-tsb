resource "tsb_user" "example" {
  organization = "tetrate"
  name         = "terryform"
  display_name = "Terry Form"
  login_name   = "terryform"
  first_name   = "Terry"
  last_name    = "Form"
  email        = "terry.form@tetrate.io"
}
