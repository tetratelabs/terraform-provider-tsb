# Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

# Install a gateway data-plane deployment via TSB. The InstallGateway resource's
# entity (GatewaySpec) lives in the install/dataplane API group; it is exposed
# under the gateway service and scoped to a gateway group.

terraform {
  required_providers {
    tsb = {
      source = "tetrateio/tsb"
    }
  }
}

provider "tsb" {
  endpoint     = var.endpoint
  token        = var.token
  organization = "tetrate"
}

variable "endpoint" {
  type    = string
  default = "tsb.example.com:443"
}

variable "token" {
  type      = string
  sensitive = true
}

resource "tsb_gateway_install_gateway" "ingress" {
  parent = "organizations/tetrate/tenants/example/workspaces/bookinfo/gatewaygroups/bookinfo-gateways"
  name   = "bookinfo-ingress"

  type             = "INGRESS"
  target_namespace = "bookinfo"
  target_cluster   = "cluster-1"
  concurrency      = 2

  # TSB requires the gateway deployment's `app` label (it is propagated to the
  # generated Pod/Service and is what other resources select on).
  kube_spec = {
    deployment = {
      labels = {
        app = "bookinfo-gateway"
      }
    }
  }
}
