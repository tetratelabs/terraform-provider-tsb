provider "tsb" {
  address = "tsb.tetrate.com:443"
  basic_auth = {
    username = "someusername"
    password = "supersecretpassword"
  }
}