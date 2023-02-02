# [WIP] TSB Terraform Provider

This provider uses the Terraform Plugin Framework (v3) NOT the Terraform Plugin SDK (v1 + v2). The scaffolding repository for this can be found [here](https://github.com/hashicorp/terraform-provider-scaffolding-framework). 

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.13.x
-	[Go](https://golang.org/doc/install) >= 1.18

## Developing the Provider

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make test`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make test
```

## Directory Structure

- `docs` documentation autogenerated from a combination of `examples` and Schemas.
- `examples` example Terraform configuration, used by docs generation code.
- `internal/provider` the code that takes Terraform configuration and creates resources in TSB.
- `tools` ensures the Go modules are present for doc generation.


## TODOs
- [ ] Evaluate whether we can gen code using golangs code gen from the api produced schema
- [ ] Migrate API from deprecated types.String creation
- [ ] Add Users
- [ ] Add Roles
- [ ] Add Org Access Bindings
- [ ] Add Tenant Access Bindings
- [ ] Add Clusters
- [ ] Add new ClusterToken replacement
- [ ] Add Workspace
- [ ] Add GatewayGroup
- [ ] Add IngressGateway
- [ ] Add TrafficGroup
- [ ] Add SecurityGroup



Blocked by https://github.com/hashicorp/terraform-plugin-framework/issues/529:
- [ ] Turn on Team Testing
- [ ] Turn on Service Account Testing
- [ ] Validate Key encoding in Service Accounts
