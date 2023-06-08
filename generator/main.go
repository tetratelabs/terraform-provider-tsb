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
	"fmt"
	"os"
	"path/filepath"

	"github.com/dave/jennifer/jen"
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
	*NewResource(
		WithName("ServiceAccount"),
		WithPkgImportPath(TsbV2),
		WithClient("TeamsClient"),
		WithSchema(tsbv2.ServiceAccountSchema()),
	),
}

const (
	Resource  = "github.com/hashicorp/terraform-plugin-framework/resource"
	Types     = "github.com/hashicorp/terraform-plugin-framework/types"
	BaseTypes = "github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	Path      = "github.com/hashicorp/terraform-plugin-framework/path"
	TsbV2     = "github.com/tetrateio/api/tsb/v2"
	Helpers   = "github.com/tetratelabs/terraform-provider-tsb/internal/helpers"
	PkgAPI    = "github.com/tetrateio/tetrate/pkg/api"
	SamberLo  = "github.com/samber/lo"
)

func main() {
	// Must be ran from root
	repoRoot, _ := os.Getwd()

	files := map[string]func(resource) *jen.File{
		"resource.go": genResource,
		"model.go":    genModel,
		"import.go":   genImport,
		"create.go":   genCreate,
		"read.go":     genRead,
		"update.go":   genUpdate,
		"delete.go":   genDelete,
	}

	errors := false

	for _, r := range resources {
		// Make the dir
		resourceDir := filepath.Join(repoRoot, "internal", "provider", "resources", r.lowerName)
		os.MkdirAll(resourceDir, 0755)

		for fname, gen := range files {
			file, err := os.Create(filepath.Join(resourceDir, fname))
			if err != nil {
				fmt.Printf("Error creating file %v. %v", filepath.Join(resourceDir, fname), err)
				errors = true
				continue
			}
			defer file.Close()
			err = gen(r).Render(file)
			if err != nil {
				errors = true
				fmt.Printf("Error rendering file %v. %v", filepath.Join(resourceDir, fname), err)
				err = os.WriteFile(filepath.Join(resourceDir, fname+".log"), []byte(err.Error()), 0666)
				if err != nil {
					fmt.Printf("%v", err)
				}
				continue
			}
		}
	}

	if errors {
		panic("errors generating")
	}

}
