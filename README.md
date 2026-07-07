# Terraform Provider for Tetrate Service Bridge (TSB)

[![Go](https://github.com/tetratelabs/terraform-provider-tsb/actions/workflows/go.yml/badge.svg)](https://github.com/tetratelabs/terraform-provider-tsb/actions/workflows/go.yml)

A Terraform provider for managing [Tetrate Service Bridge (TSB)](https://tetrate.io/tetrate-service-bridge/) resources — workspaces, gateways, Istio configuration, and more — via the TSB v2 gRPC API.

## Using the provider

```hcl
terraform {
  required_providers {
    tsb = {
      source = "tetratelabs/tsb"
    }
  }
}

provider "tsb" {
  endpoint     = "tsb.example.com:443"
  organization = "tetrate"
  token        = var.tsb_token
}
```

All provider settings can also be supplied via environment variables (`TSB_ENDPOINT`, `TSB_TOKEN`, `TSB_USERNAME`, `TSB_PASSWORD`, `TSB_ORG`, `TSB_CA_FILE`, `TSB_TLS_INSECURE`, `TSB_TLS_DISABLED`). See [`examples/provider.tf`](examples/provider.tf) and [`docs/index.md`](docs/index.md) for the full set of options.

Resource documentation lives in [`docs/resources/`](docs/resources), with runnable examples under [`examples/resources/`](examples/resources) and end-to-end scenarios under [`examples/`](examples) (e.g. [`cluster-onboarding`](examples/cluster-onboarding), [`install-gateway`](examples/install-gateway), [`istio-direct`](examples/istio-direct)).

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://go.dev/dl/) (see `go.mod` for the required version) — only needed to build the provider yourself

## Development

```sh
make build      # compile the plugin into a Terraform filesystem-mirror layout
make test       # go test ./...
make generate   # regenerate internal/provider/*_gen.go from the TSB v2 API descriptors
make docs       # regenerate docs/ from schemas + examples/
make format     # gosimports + gofmt + license-header apply
make check      # generate + docs + format, then fails if that produced any git diff
make all        # generate format build test
```

After `make build`, point a Terraform config's [filesystem mirror](https://developer.hashicorp.com/terraform/cli/config/config-file#filesystem_mirror) at the `provider/` output directory and run `terraform init -upgrade`.

Most of `internal/provider/` is machine-generated from the TSB v2 API's proto descriptors by `cmd/gen-provider` — never hand-edit a `*_gen.go` file; change the generator and re-run `make generate` instead. See [`CLAUDE.md`](CLAUDE.md) for a detailed walkthrough of the generation pipeline and provider architecture.

## License

[Apache License 2.0](LICENSE)
