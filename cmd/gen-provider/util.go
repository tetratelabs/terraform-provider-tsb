// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package main

import (
	"regexp"
	"strings"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

// snakeCase converts a Go identifier to snake_case, matching the convention used
// by the schema generator (api/protoc-plugins/protoc-gen-terraform) so model
// tfsdk tags line up with the generated schema attribute names.
func snakeCase(s string) string {
	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
