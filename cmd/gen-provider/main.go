// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

// Command gen-provider generates the Terraform provider's resource, model, and
// registration Go files from the TSB v2 API descriptors. It reuses the schema
// functions emitted by api/protoc-plugins/protoc-gen-terraform.
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	outDir := flag.String("out", "internal/provider", "output directory for generated provider files")
	dryRun := flag.Bool("dry-run", false, "discover and print resources without writing files")
	flag.Parse()

	resources, err := discoverResources()
	if err != nil {
		fmt.Fprintf(os.Stderr, "gen-provider: discovery failed: %v\n", err)
		os.Exit(1)
	}

	if *dryRun {
		fmt.Printf("discovered %d resources:\n", len(resources))
		for _, r := range resources {
			fmt.Printf("  tsb_%-28s msg=%s.%s client=%s collection=%s force=%v\n",
				r.TFName, r.MessageGoType.PkgPath(), r.MessageGoType.Name(),
				r.ClientCtor, r.Collection, r.DeleteHasForce)
		}
		return
	}

	n, err := generate(resources, *outDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "gen-provider: generation failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("gen-provider: wrote %d/%d resources to %s\n", n, len(resources), *outDir)
}
