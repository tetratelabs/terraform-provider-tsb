# Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

# Terraform provider for Tetrate Service Bridge (TSB).
#
# internal/provider is generated from the TSB v2 API descriptors vendored in
# via the github.com/tetrateio/tetrate dependency (see go.mod). Bump that
# dependency, then re-run `make generate`, whenever the upstream API changes.

MODULE_PATH := $(shell go list -m)
GOSIMPORTS  := github.com/rinchsan/gosimports/cmd/gosimports@v0.3.8
LICENSER    := github.com/liamawhite/licenser@169f8d8c2e886454b19cb9e8170981b9ee3eba9e
TFPLUGINDOCS := github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.19.4

# Provider coordinates and the local filesystem mirror to install into. These
# match a Terraform CLI config (~/.terraformrc) such as:
#
#   provider_installation {
#     filesystem_mirror { path = "<MIRROR_DIR>" }
#   }
#
# so that `make build` drops the plugin where `terraform init` discovers it.
PROVIDER_HOST      ?= registry.terraform.io
PROVIDER_NAMESPACE ?= tetratelabs
PROVIDER_TYPE      ?= tsb
MIRROR_DIR         ?= provider

# PROVIDER_VERSION is derived from git: the exact tag when HEAD is tagged
# (leading "v" stripped), otherwise "0.0.0-<short-sha>".
PROVIDER_VERSION ?= $(shell t=$$(git describe --tags --exact-match 2>/dev/null); if [ -n "$$t" ]; then echo "$$t" | sed 's/^v//'; else echo "0.0.0-$$(git rev-parse --short HEAD)"; fi)

VERSION_PKG := github.com/tetratelabs/terraform-provider-tsb/internal/provider

OS_ARCH    := $(shell go env GOOS)_$(shell go env GOARCH)
MIRROR_PKG := $(MIRROR_DIR)/$(PROVIDER_HOST)/$(PROVIDER_NAMESPACE)/$(PROVIDER_TYPE)/$(PROVIDER_VERSION)/$(OS_ARCH)
BIN        := $(MIRROR_PKG)/terraform-provider-$(PROVIDER_TYPE)

.PHONY: all
all: generate format build test

# generate (re)generates the provider model, resource, registration, and parity
# test files from the TSB v2 API descriptors.
.PHONY: generate
generate:
	@echo "Generating provider code..."
	@go run ./cmd/gen-provider -out internal/provider

# docs (re)generates the registry documentation from the provider schema and
# the examples/ directory.
.PHONY: docs
docs:
	@echo "Generating registry docs..."
	@terraform fmt -recursive ./examples/
	@go run $(TFPLUGINDOCS)

# build compiles the provider plugin binary into the filesystem mirror layout.
# After building, run `terraform init -upgrade` so the dependency lock picks up
# the new binary's checksum.
.PHONY: build
build:
	@echo "Building $(BIN) (version $(PROVIDER_VERSION))..."
	@mkdir -p $(MIRROR_PKG)
	@go build -ldflags "-X $(VERSION_PKG).version=$(PROVIDER_VERSION)" -o $(BIN) .

.PHONY: test
test:
	@go test ./...

.PHONY: format
format: ## Format all Go code and apply license headers
	@echo "formatting code..."
	@go run $(GOSIMPORTS) -w -local $(MODULE_PATH) .
	@gofmt -w .
	@echo "applying license headers..."
	@go run $(LICENSER) apply -r -t header.txt -m "Copyright (c) Tetrate" "Tetrate, Inc"

# check re-runs the generators and formatters and fails if that produced any
# diff, i.e. it verifies the committed output is up to date.
.PHONY: check
check: generate docs format
	@[ -z "`git status -uno --porcelain`" ] || (git status && echo 'Check failed: generate/docs/format produced a diff, or the working tree was already dirty.'; exit 1)

.PHONY: clean
clean:
	@rm -f $(BIN)
