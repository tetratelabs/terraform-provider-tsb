default: test

# WARNING!!! THESE CREATE ACTUAL RESOURCES
# Run acceptance tests
# Need to set the following env vars:
#   export TSB_ADDRESS=<mp-address>:<port>
#   export TSB_ORGANIZATION=<org-name>
#   export TSB_USERNAME=admin
#   export TSB_PASSWORD=<admin-password>
# GRPC debugging can be enabled as follows:
# GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info
.PHONY: test
test:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 1m