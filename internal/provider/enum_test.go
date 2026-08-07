// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	rbac "github.com/tetrateio/tetrate/api/tsb/rbac/v2"
	tsbtypes "github.com/tetrateio/tetrate/api/tsb/types/v2"
)

// TestExpandEnumRejectsUnrecognisedName covers the core bug: a non-empty string
// that is not a declared enum name must produce an error rather than silently
// becoming the enum's zero value.
func TestExpandEnumRejectsUnrecognisedName(t *testing.T) {
	ctx := context.Background()
	m := &GatewayGroupModel{ConfigMode: types.StringValue("bridged")} // wrong case

	out, d := expandGatewayGroup(ctx, m)

	if !d.HasError() {
		t.Fatalf("expected an error diagnostic for an unrecognised enum name, got none (config_mode became %v)", out.ConfigMode)
	}
	var sb strings.Builder
	for _, e := range d.Errors() {
		sb.WriteString(e.Summary())
		sb.WriteString(" ")
		sb.WriteString(e.Detail())
		sb.WriteString(" ")
	}
	got := sb.String()
	for _, want := range []string{"config_mode", "bridged", "BRIDGED", "DIRECT"} {
		if !strings.Contains(got, want) {
			t.Errorf("diagnostic %q does not mention %q", got, want)
		}
	}
}

// TestExpandEnumAcceptsDeclaredName guards against the fix rejecting valid input.
func TestExpandEnumAcceptsDeclaredName(t *testing.T) {
	ctx := context.Background()
	m := &GatewayGroupModel{ConfigMode: types.StringValue("DIRECT")}

	out, d := expandGatewayGroup(ctx, m)

	if d.HasError() {
		t.Fatalf("unexpected diagnostics for a valid enum name: %v", d)
	}
	if out.ConfigMode != tsbtypes.ConfigMode_DIRECT {
		t.Errorf("config_mode = %v, want DIRECT", out.ConfigMode)
	}
}

// TestExpandEnumLeavesUnsetValueAtZero is the regression guard described in
// issue #20: fixupSchema makes optional leaf attributes Optional+Computed, so an
// attribute the user never set arrives null (or unknown, at create). ValueString
// returns "" for both, and that must keep mapping to the proto zero value
// silently — erroring there would break every config that omits an optional enum.
func TestExpandEnumLeavesUnsetValueAtZero(t *testing.T) {
	ctx := context.Background()
	for name, v := range map[string]types.String{
		"null":    types.StringNull(),
		"unknown": types.StringUnknown(),
	} {
		t.Run(name, func(t *testing.T) {
			out, d := expandGatewayGroup(ctx, &GatewayGroupModel{ConfigMode: v})
			if d.HasError() {
				t.Fatalf("unset enum produced diagnostics: %v", d)
			}
			if out.ConfigMode != tsbtypes.ConfigMode_BRIDGED {
				t.Errorf("config_mode = %v, want the zero value BRIDGED", out.ConfigMode)
			}
		})
	}
}

// TestExpandEnumRejectsNumericName covers the version-skew case: protobuf-go
// renders an enum number the compiled descriptor does not know as its decimal,
// so a newer server's value round-trips through state as e.g. "7". Erroring is
// deliberate — silently sending the zero value would overwrite the server's
// real configuration.
func TestExpandEnumRejectsNumericName(t *testing.T) {
	ctx := context.Background()

	_, d := expandGatewayGroup(ctx, &GatewayGroupModel{ConfigMode: types.StringValue("7")})

	if !d.HasError() {
		t.Fatal("expected an error diagnostic for an unrecognised numeric enum value, got none")
	}
}

// TestExpandEnumRejectsUnrecognisedNameInList covers the repeated-enum path,
// which builds its values in a loop rather than through scalarToProto.
func TestExpandEnumRejectsUnrecognisedNameInList(t *testing.T) {
	ctx := context.Background()
	perms, d := types.ListValueFrom(ctx, types.StringType, []string{"READ", "EXECUTE"})
	if d.HasError() {
		t.Fatalf("building list: %v", d)
	}

	out, d := expandRole_Rule(ctx, &Role_RuleModel{Permissions: perms})

	if !d.HasError() {
		t.Fatalf("expected an error diagnostic for an unrecognised enum name in a list, got none (permissions became %v)", out.Permissions)
	}
	var sb strings.Builder
	for _, e := range d.Errors() {
		sb.WriteString(e.Summary())
		sb.WriteString(" ")
		sb.WriteString(e.Detail())
		sb.WriteString(" ")
	}
	if got := sb.String(); !strings.Contains(got, "EXECUTE") {
		t.Errorf("diagnostic %q does not mention the offending value", got)
	}
}

// TestExpandEnumAcceptsDeclaredNamesInList guards the valid repeated case.
func TestExpandEnumAcceptsDeclaredNamesInList(t *testing.T) {
	ctx := context.Background()
	perms, d := types.ListValueFrom(ctx, types.StringType, []string{"READ", "WRITE"})
	if d.HasError() {
		t.Fatalf("building list: %v", d)
	}

	out, d := expandRole_Rule(ctx, &Role_RuleModel{Permissions: perms})

	if d.HasError() {
		t.Fatalf("unexpected diagnostics for valid enum names: %v", d)
	}
	want := []rbac.Permission{rbac.Permission_READ, rbac.Permission_WRITE}
	if len(out.Permissions) != len(want) {
		t.Fatalf("permissions = %v, want %v", out.Permissions, want)
	}
	for i := range want {
		if out.Permissions[i] != want[i] {
			t.Errorf("permissions[%d] = %v, want %v", i, out.Permissions[i], want[i])
		}
	}
}
