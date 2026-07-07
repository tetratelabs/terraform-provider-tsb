# Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

# Grant RBAC access on a TSB resource via its access bindings (AccessPolicy).
#
# AccessBindings always exist for every resource, so this is a Get/Set resource:
# `tsb_access_binding` is authoritative for the target's bindings — it replaces
# the full set on create/update and clears them on destroy (the policy object
# itself cannot be deleted).

terraform {
  required_providers {
    tsb = {
      source = "tetratelabs/tsb"
    }
  }
}

provider "tsb" {
  endpoint     = var.endpoint
  organization = "tetrate"
  token        = var.token
}

variable "endpoint" {
  type    = string
  default = "tsb.example.com:443"
}

variable "token" {
  type      = string
  sensitive = true
}

resource "tsb_access_binding" "bookinfo_workspace" {
  # The TSB resource the bindings apply to.
  fqn         = "organizations/tetrate/tenants/example/workspaces/bookinfo"
  description = "Access to the bookinfo workspace"

  allow = [
    {
      role = "organizations/tetrate/roles/admin"
      subjects = [
        { user = "alice@example.com" },
        { team = "organizations/tetrate/teams/platform" },
      ]
    },
    {
      role = "organizations/tetrate/roles/editor"
      subjects = [
        { service_account = "organizations/tetrate/serviceaccounts/ci" },
      ]
    },
  ]
}
