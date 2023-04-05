// Copyright 2023 Tetrate
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"path/filepath"

	tsbv2 "github.com/tetrateio/api/tsb/v2"
)

var resources = []resource{
	*NewResource(
		WithName("Tenant"),
		WithPkgImportPath(TsbV2),
		WithClient("TenantsClient"),
		WithSchema(tsbv2.TenantSchema()),
	),
	*NewResource(
		WithName("User"),
		WithPkgImportPath(TsbV2),
		WithClient("TeamsClient"),
		WithSchema(tsbv2.UserSchema()),
	),
	*NewResource(
		WithName("Cluster"),
		WithPkgImportPath(TsbV2),
		WithClient("ClustersClient"),
		WithSchema(tsbv2.ClusterSchema()),
	),
}

const (
	Resource = "github.com/hashicorp/terraform-plugin-framework/resource"
	Types    = "github.com/hashicorp/terraform-plugin-framework/types"
	Path     = "github.com/hashicorp/terraform-plugin-framework/path"
	TsbV2    = "github.com/tetrateio/api/tsb/v2"
	Helpers  = "github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
	PkgAPI   = "github.com/tetrateio/tetrate/pkg/api"
)

func main() {
	// Must be ran from root
	repoRoot, _ := os.Getwd()

	for _, r := range resources {
		// Make the dir
		resourceDir := filepath.Join(repoRoot, "internal", "provider", "resources", r.lowerName)
		os.MkdirAll(resourceDir, 0755)

		// Write the files
		os.WriteFile(filepath.Join(resourceDir, "resource.go"), []byte(genResource(r).GoString()), 0644)
		os.WriteFile(filepath.Join(resourceDir, "model.go"), []byte(genModel(r).GoString()), 0644)
		os.WriteFile(filepath.Join(resourceDir, "import.go"), []byte(genImport(r).GoString()), 0644)
		os.WriteFile(filepath.Join(resourceDir, "create.go"), []byte(genCreate(r).GoString()), 0644)
		os.WriteFile(filepath.Join(resourceDir, "read.go"), []byte(genRead(r).GoString()), 0644)
		os.WriteFile(filepath.Join(resourceDir, "update.go"), []byte(genUpdate(r).GoString()), 0644)
		os.WriteFile(filepath.Join(resourceDir, "delete.go"), []byte(genDelete(r).GoString()), 0644)
	}

}
