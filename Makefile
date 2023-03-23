# Copyright 2023 Tetrate
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

default: build

.PHONY: docs
docs:
	go install 
	go generate ./...

build:
	go build -v ./...

check: docs
	licenser verify Tetrate -r
	[ -z "`git status -uno --porcelain`" ] || (git status && echo 'Check failed. This could be a failed check or dirty git state.'; exit 1)

lint:
	licenser apply Tetrate -r .

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
	source .env && TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 1m
