---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tsb_service_account Resource - terraform-provider-tsb"
subcategory: ""
description: |-
  
---

# tsb_service_account (Resource)



## Example Usage

```terraform
resource "tsb_service_account" "example" {
  name         = "ciautomation"
  organization = "tetrate"
  display_name = "CI Automation"
  description  = "Used by our CI tool to create resources"
  key_encoding = "JWT"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The short name for the resource to be created.
- `organization` (String) The organization the Service Account will belong to.

### Optional

- `description` (String) A description of the resource.
- `display_name` (String) User friendly name for the resource.
- `key_encoding` (String) The format in which the generated key pairs will be returned. If not set keys are returned in PEM format.

### Read-Only

- `id` (String) Fully-qualified name of the resource.
- `keys` (Attributes List) Keys associated with the service account. A default key-pair is automatically created when the Service Account is created. Note that TSB does not store the private keys, so it is up to the client to store the returned private keys securely, as they are only returned once after creation. Additional keys can be added (and deleted) by using the corresponding key management APIs. +protoc-gen-terraform:computed (see [below for nested schema](#nestedatt--keys))

<a id="nestedatt--keys"></a>
### Nested Schema for `keys`

Read-Only:

- `default_token` (String) A default access token that can be used to authenticate to TSB on behalf of the service account. TSB does not store this token and it is only returned when a service account key is created, similar to the private key. It is up to the client to store the token for future use or to use the TSB CLI to generate new tokens as explained in: https://docs.tetrate.io/service-bridge/latest/en-us/howto/service-accounts
- `encoding` (String) Format in which the public and private keys are encoded. By default keys are returned in PEM format.
- `id` (String) Unique identifier for this key-pair. This should be used as the `kid` (key id) when generating JWT tokens that are signed with this key-pair.
- `private_key` (String) The encoded private key associated with the service account. TSB does not store the private key and it is up to the client to store it safely. The encoding format is determined by the `encoding` field.
- `public_key` (String) The encoded public key associated with the service account. The encoding format is determined by the `encoding` field.

