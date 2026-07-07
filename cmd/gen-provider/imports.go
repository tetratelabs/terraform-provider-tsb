// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package main

// Blank imports register the TSB v2 API groups' proto descriptors and Go types
// so the generator can discover their resources via the proto registries.
//
// Discovery is structural and package-agnostic: it scans every registered
// `tetrateio.api.tsb.*` package for the standard CRUD quartet, so adding a new
// *resource* to any group below requires no generator changes. Only adding a
// brand-new API *group* (a new proto package) needs a line here.
import (
	_ "github.com/tetrateio/tetrate/api/tsb/application/v2"
	_ "github.com/tetrateio/tetrate/api/tsb/extension/v2"
	_ "github.com/tetrateio/tetrate/api/tsb/gateway/v2"
	_ "github.com/tetrateio/tetrate/api/tsb/istiointernal/v2"
	_ "github.com/tetrateio/tetrate/api/tsb/profile/v2"
	_ "github.com/tetrateio/tetrate/api/tsb/rbac/v2"
	_ "github.com/tetrateio/tetrate/api/tsb/security/v2"
	_ "github.com/tetrateio/tetrate/api/tsb/traffic/v2"
	_ "github.com/tetrateio/tetrate/api/tsb/v2"
)
