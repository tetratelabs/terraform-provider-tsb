# Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

# Cluster onboarding: register a Cluster in TSB, then store the cluster service
# account key as a Kubernetes secret in the cluster being onboarded — the TSB
# control plane reads that secret to authenticate to the management plane.
#
# When the Cluster is created, TSB generates a service account with a JWK key
# pair. The private key is returned ONLY in the create response (it is not stored
# server-side), so the provider captures it into the computed `service_account`
# attribute and preserves it across later reads/updates.

terraform {
  required_providers {
    tsb = {
      source = "tetrateio/tsb"
    }
    kubernetes = {
      source = "hashicorp/kubernetes"
    }
  }
}

provider "tsb" {
  endpoint     = var.tsb_endpoint
  organization = var.tsb_organization
  token        = var.tsb_token
}

# Points at the physical cluster being onboarded (where the TSB control plane runs).
provider "kubernetes" {
  config_path = var.kubeconfig
}

variable "tsb_endpoint" {
  type = string
}

variable "tsb_organization" {
  type    = string
  default = "tetrate"
}

variable "tsb_token" {
  type      = string
  sensitive = true
}

variable "kubeconfig" {
  type    = string
  default = "~/.kube/config"
}

variable "control_plane_namespace" {
  type    = string
  default = "istio-system"
}

# 1. Register the cluster in TSB. The server returns a service account whose
#    private (JWK) key is only present in this create response.
resource "tsb_cluster" "onboarded" {
  parent       = "organizations/${var.tsb_organization}"
  name         = "physical-cluster-1"
  display_name = "Physical Cluster 1"
  network      = "network-1"

  labels = {
    region = "us-east-1"
  }
}

# Decode the JSON-encoded service_account attribute and pull out the JWK key.
locals {
  service_account = jsondecode(tsb_cluster.onboarded.service_account)
  cluster_jwk     = local.service_account.keys[0].privateKey
}

# 2. Store the key as the `cluster-service-account` secret the control plane
#    expects: the cluster FQN plus the JWK private key. The kubernetes provider
#    base64-encodes `data` values automatically.
resource "kubernetes_secret" "cluster_service_account" {
  metadata {
    name      = "cluster-service-account"
    namespace = var.control_plane_namespace
  }

  type = "Opaque"

  data = {
    clusterFqn = tsb_cluster.onboarded.id
    jwk        = local.cluster_jwk
  }
}

output "cluster_fqn" {
  value = tsb_cluster.onboarded.id
}
