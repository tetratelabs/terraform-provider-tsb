# Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

terraform {
  required_providers {
    tsb = {
      source = "tetratelabs/tsb"
    }
  }
}

# Configuration can also be supplied via environment variables:
#   TSB_ENDPOINT, TSB_TOKEN, TSB_USERNAME, TSB_PASSWORD, TSB_ORG, TSB_CA_FILE,
#   TSB_TLS_INSECURE, TSB_TLS_DISABLED
provider "tsb" {
  endpoint     = "tsb.example.com:443"
  organization = "tetrate"

  # Authenticate with a token...
  token = var.tsb_token

  # ...or with username/password (used when no token is set):
  # username = "admin"
  # password = var.tsb_password

  # Skip server certificate verification (e.g. self-signed dev certs):
  # tls_insecure = true
}

variable "tsb_token" {
  type      = string
  sensitive = true
  default   = ""
}

variable "tsb_password" {
  type      = string
  sensitive = true
  default   = ""
}
