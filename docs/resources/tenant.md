---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tsb_tenant Resource - terraform-provider-tsb"
subcategory: ""
description: |-
  
---

# tsb_tenant (Resource)



## Example Usage

```terraform
resource "tsb_tenant" "example" {
  parent = "organizations/tetrate"
  name   = "tenantone"
  tenant = {
    display_name = "Tenant One"
    description  = "I'm a tenant"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The short name for the resource to be created.
- `tenant` (Attributes) Details of the tenant to be created. (see [below for nested schema](#nestedatt--tenant))

### Optional

- `parent` (String) Parent resource where the Tenant will be created.

### Read-Only

- `id` (String) The FQN but for Terraform

<a id="nestedatt--tenant"></a>
### Nested Schema for `tenant`

Optional:

- `description` (String) A description of the resource.
- `display_name` (String) User friendly name for the resource.
- `etag` (String) The etag for the resource. This field is automatically computed and must be sent on every update to the resource to prevent concurrent modifications.
- `fqn` (String) Fully-qualified name of the resource. This field is read-only.
- `security_domain` (String) Security domains can be used to group different resources under the same security domain. Although security domain is not resource itself currently, it follows a fqn format `organizations/myorg/securitydomains/mysecuritydomain`, and a child cannot override any ancestor's security domain. Once a security domain is assigned to a _Tenant_, all the children resources will belong to that security domain in the same way a _Workspace_ belongs to a _Tenant_, a _Workspace_ will also belong to the security domain assigned to the _Tenant_. Security domains can also be used to define _Security settings Authorization rules_ in which you can allow or deny request from or to a security domain.


