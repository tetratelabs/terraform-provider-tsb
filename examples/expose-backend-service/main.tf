# Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

# End-to-end example: create a Tenant, a Workspace and a Gateway Group, then
# expose a backend service (bookinfo's productpage) through a unified Gateway.
#
# Resource hierarchy (each resource's `parent` is the FQN/`id` of its parent):
#
#   organizations/<org>
#     └─ tenant                      tsb_tenant
#          └─ workspace              tsb_workspace
#               └─ gateway group     tsb_gateway_group
#                    └─ gateway      tsb_gateway_gateway

terraform {
  required_providers {
    tsb = {
      source = "tetratelabs/tsb"
    }
  }
}

# Endpoint and token also read from TSB_ENDPOINT / TSB_TOKEN if omitted.
provider "tsb" {
  endpoint     = var.endpoint
  token        = var.token
  organization = var.organization
}

variable "organization" {
  type    = string
  default = "tetrate"
}

variable "endpoint" {
  type    = string
  default = "tsb.example.com:443"
}

variable "token" {
  type      = string
  sensitive = true
}

# 1. Tenant under the organization.
resource "tsb_tenant" "example" {
  parent       = "organizations/${var.organization}"
  name         = "example"
  display_name = "Example Tenant"
}

# 2. Workspace owning the bookinfo namespaces across any cluster.
resource "tsb_workspace" "example" {
  parent       = tsb_tenant.example.id
  name         = "bookinfo"
  display_name = "Bookinfo Workspace"

  namespace_selector = {
    names = ["*/bookinfo"]
  }
}

# 3. Gateway group: a logical grouping of gateways within the workspace.
resource "tsb_gateway_group" "example" {
  parent       = tsb_workspace.example.id
  name         = "bookinfo-gateways"
  display_name = "Bookinfo Gateways"

  namespace_selector = {
    names = ["*/bookinfo"]
  }
}

# 4. Gateway selecting an Istio ingress-gateway workload and exposing the
#    productpage service on HTTP :80 for host bookinfo.example.com.
resource "tsb_gateway_gateway" "example" {
  parent       = tsb_gateway_group.example.id
  name         = "productpage"
  display_name = "Productpage Gateway"

  # Picks the gateway deployment that will serve this configuration.
  workload_selector = {
    namespace = "bookinfo"
    labels = {
      app = "tsb-gateway-bookinfo"
    }
  }

  http = [
    {
      name     = "productpage"
      port     = 80
      hostname = "bookinfo.example.com"

      routing = {
        rules = [
          {
            # Send all matching traffic to the productpage backend service.
            # `host` is in "<namespace>/<fqdn>" form.
            route = {
              service_destination = {
                host = "bookinfo/productpage.bookinfo.svc.cluster.local"
                port = 9080
              }
            }
          }
        ]
      }
    }
  ]
}

output "gateway_fqn" {
  value = tsb_gateway_gateway.example.id
}
