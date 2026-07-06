# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this is

The Terraform provider for Tetrate Service Bridge (TSB), managing TSB v2 API resources (`tsb_workspace`, `tsb_gateway_group`, `tsb_istio_gateway_virtual_service`, etc.) via gRPC. Nearly all resource code under `internal/provider/*_gen.go` and `internal/provider/models_gen.go` is machine-generated from the TSB v2 API's proto descriptors — read `cmd/gen-provider` to understand *why* code looks the way it does before editing generated output by hand (you shouldn't; see Architecture below).

This repo was recently split out of a monorepo (branch `migrate-from-monorepo`); `.superpowers/sdd/progress.md` has the migration history if relevant context is missing here.

## Local dev setup gotcha

`go.mod` currently has:

```
replace github.com/tetrateio/tetrate => /Users/chirauki/src/tetrate-dos
```

This is a **local filesystem replace**, not a real dependency — `go build`/`go test` will only work on a machine with that path present. CI (`.github/workflows/go.yml`) instead pulls `github.com/tetrateio/tetrate` as a real (private) module over SSH, using `GOPRIVATE=github.com/tetrateio/*` and an `insteadOf` git rewrite for `https://github.com/` → `git@github.com:`. If you need to touch dependency resolution, be aware the local machine and CI resolve this module differently.

## Common commands

```sh
make build      # compile the plugin into a Terraform filesystem-mirror layout (provider/registry.terraform.io/...)
make test       # go test ./...
make generate   # regenerate internal/provider/*_gen.go from the TSB v2 API descriptors (cmd/gen-provider)
make docs       # regenerate docs/ (tfplugindocs) from schemas + examples/
make format     # gosimports + gofmt + license-header apply (header.txt)
make check      # generate + docs + format, then fails if that produced any git diff (this is what CI's "check" job runs)
make all        # generate format build test
```

- Single test: `go test ./internal/provider/ -run TestName -v`
- After `make build`, run `terraform init -upgrade` in a config pointed at the filesystem mirror so the dependency lock picks up the new binary's checksum (see the `PROVIDER_HOST`/`PROVIDER_NAMESPACE`/`MIRROR_DIR` vars in the Makefile and the matching `provider_installation.filesystem_mirror` block in `~/.terraformrc`).
- `PROVIDER_VERSION` is derived from git (exact tag with `v` stripped, else `0.0.0-<short-sha>`) — don't hardcode versions elsewhere.
- Never hand-edit a `*_gen.go` file or `models_gen.go`; change `cmd/gen-provider` and re-run `make generate` instead. `make check` (run in CI) will catch drift between the generator and committed generated output.

## Architecture

### Code generation pipeline (`cmd/gen-provider`)

`gen-provider` walks the proto registry (via blank imports in `imports.go` — one line per TSB API *group* like `gateway/v2`; new *resources* within an already-imported group need no generator change) and:

1. **`discover.go`** — finds every gRPC service method following the `Create<Stem>`/`Get<Stem>`/`Update<Stem>`/`Delete<Stem>` CRUD-quartet convention, keyed off an `Update<Stem>` method whose input/output message *is* the resource entity. Derives the Terraform type name, FQN collection segment (parsed out of the `google.api.http` `{fqn=.../*}` path template), and the Go client/request types via reflection over the compiled proto registry. Resources backed by a message shared across services (Istio direct-mode: `IstioObject` reused by VirtualService/Gateway/etc.) are detected and disambiguated by service+stem instead of message name.
2. **`emit.go`/`emit_resource.go`/`emit_conv.go`** — using `github.com/dave/jennifer` to generate Go source, walks each resource's message closure recursively, emitting one `<Base>Model` struct (with `tfsdk` tags) plus expand/flatten functions per referenced message type, and a full `resource.Resource` implementation (CRUD methods) per top-level resource into `internal/provider/<tf_name>_resource_gen.go`. Also emits `models_gen.go`, `provider_gen.go` (the `Resources()` registry), and `parity_gen_test.go` (asserts model `tfsdk` tags exactly match the reused schema's attributes — catches model/schema drift at test time rather than a runtime panic).
3. A resource whose model closure hits an unsupported proto shape (oneofs beyond what's handled, message maps, structpb, etc.) is **skipped with a warning**, not a hard failure — the generator prefers a working subset of resources over an all-or-nothing abort. Check generator stderr output after `make generate` if an expected resource is missing.
4. Each generated resource reuses an `XxxSchema()` function assumed already generated elsewhere (by `protoc-gen-terraform`, upstream in the API repo) — `internal/provider/schema_fixup.go`'s `fixupSchema` patches that reused schema at serve time (not codegen time) so it's valid for the plugin framework: drops `Default` from `Required` attributes, and makes `Optional` leaf attributes `Optional+Computed` (TSB always returns a fully materialized object server-side, so unset optional fields must round-trip the server's default rather than causing a "produced inconsistent result" plan error). Nested-object attributes are deliberately left non-Computed since generated models back them with plain Go structs that can't represent "unknown".
5. Fields backed by TSB messages that only return their value at create time (`createOnlyMessages` in `emit.go`, e.g. `ServiceAccount.KeyPair`) are modeled as a computed normalized-JSON string (`jsontypes.Normalized`), captured from prior state before every flatten and restored after, so reads/updates never blank out a secret the server stops echoing back.

### Hand-written resources

Not every resource fits the CRUD-quartet generator. `internal/provider/access_binding_resource.go` (`tsb_access_binding`) is Get/Set-based rather than Create/Delete: an RBAC access policy implicitly exists for every TSB resource and can't be created or destroyed, only read and replaced — Delete clears the bindings via `SetPolicy` with an empty list instead of removing anything. It's registered manually alongside the generated set in `provider.go`'s `Resources()`. Any future resource with non-quartet semantics belongs here, not in the generator.

### Provider runtime (`internal/provider/provider.go`, `client.go`)

- `provider.go` wires provider-level config (`endpoint`/`token`/`username`+`password`/`organization`/`ca_file`/`tls_disabled`/`tls_insecure`, each with a `TSB_*` env var fallback) and dials a single shared gRPC connection, handed to every resource as `*client` via `ResourceData`.
- `client.go`'s `dial`/`transportCredentials`/`perRPCCredentials` intentionally mirror `tctl`/`teamsync`'s connection conventions: token auth takes precedence over HTTP Basic, and TLS has ALPN explicitly disabled (`xcredentials.NewTLSWithALPNDisabled`) for compatibility with TSB instances older than 1.13.

### Testing

- `parity_gen_test.go` (generated) — model/schema attribute parity, see above.
- `*_test.go` files hand-written alongside the generated resources cover conversion helpers (`conv.go`), collections, oneofs, and create-only-field handling — these describe non-obvious edge cases in the expand/flatten logic and are a good reference before changing `conv.go` or `emit_conv.go`.
