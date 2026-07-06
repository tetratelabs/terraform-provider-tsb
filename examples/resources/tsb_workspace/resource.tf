# Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

# A Workspace groups a set of namespaces across one or more clusters under a tenant.
resource "tsb_workspace" "example" {
  parent       = "organizations/tetrate/tenants/example"
  name         = "example-workspace"
  display_name = "Example Workspace"
  description  = "Managed by Terraform"

  namespace_selector = {
    names = ["*/example-ns"]
  }
}

# Import an existing workspace by its fully-qualified name:
#   terraform import tsb_workspace.example organizations/tetrate/tenants/example/workspaces/example-workspace
