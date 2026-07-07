# Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

# Istio "direct mode": manage a raw Istio resource (here a VirtualService) through
# TSB. The object's spec is the Istio API proto carried as google.protobuf.Any and
# surfaced as a JSON string — supply it with jsonencode(...) including the `@type`.

terraform {
  required_providers {
    tsb = {
      source = "tetratelabs/tsb"
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

# A VirtualService scoped to a gateway group (direct mode via the IstioGateway
# service). `parent` is the gateway group FQN.
resource "tsb_istio_gateway_virtual_service" "reviews" {
  parent = "organizations/tetrate/tenants/example/workspaces/bookinfo/gatewaygroups/bookinfo-gateways"
  name   = "reviews"

  metadata = {
    api_version = "networking.istio.io/v1alpha3"
    kind        = "VirtualService"
    name        = "reviews"
    namespace   = "bookinfo"
  }

  spec = jsonencode({
    "@type" = "type.googleapis.com/istio.networking.v1alpha3.VirtualService"
    hosts   = ["reviews.bookinfo.svc.cluster.local"]
    http = [{
      route = [{
        destination = {
          host = "reviews.bookinfo.svc.cluster.local"
          port = { number = 9080 }
        }
      }]
    }]
  })
}
