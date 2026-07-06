// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

// Command terraform-provider-tsb is the Terraform provider plugin for managing
// Tetrate Service Bridge (TSB) v2 API resources.
package main

import (
	"context"
	"flag"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/tetratelabs/terraform-provider-tsb/internal/provider"
)

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/tetratelabs/tsb",
		Debug:   debug,
	}

	if err := providerserver.Serve(context.Background(), provider.New(), opts); err != nil {
		log.Fatal(err.Error())
	}
}
