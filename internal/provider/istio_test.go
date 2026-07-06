// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"strings"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	networking "istio.io/api/networking/v1alpha3"

	typesv2 "github.com/tetrateio/tetrate/api/tsb/types/v2"
)

// TestIstioObjectSpecRoundTrip verifies that the Istio direct-mode spec (a
// google.protobuf.Any holding an istio.io/api proto) round-trips through the
// JSON-string model representation. It also confirms the istio types are
// registered so protojson can resolve the Any.
func TestIstioObjectSpecRoundTrip(t *testing.T) {
	ctx := context.Background()

	vs := &networking.VirtualService{
		Hosts:    []string{"reviews.bookinfo.svc.cluster.local"},
		Gateways: []string{"bookinfo-gateway"},
	}
	spec, err := anypb.New(vs)
	if err != nil {
		t.Fatalf("anypb.New: %v", err)
	}
	in := &typesv2.IstioObject{
		Fqn: "organizations/o/tenants/t/workspaces/w/gatewaygroups/g/virtualservices/reviews",
		Metadata: &typesv2.IstioObject_ConfigMeta{
			ApiVersion: "networking.istio.io/v1alpha3",
			Kind:       "VirtualService",
			Name:       "reviews",
			Namespace:  "bookinfo",
		},
		Spec: spec,
	}

	var m IstioObjectModel
	if d := flattenIstioObject(ctx, in, &m); d.HasError() {
		t.Fatalf("flatten: %v", d)
	}
	if m.Spec.IsNull() {
		t.Fatal("spec should not be null after flatten")
	}
	if !strings.Contains(m.Spec.ValueString(), "VirtualService") {
		t.Fatalf("spec JSON should carry the @type, got %q", m.Spec.ValueString())
	}
	if got := m.Metadata.Kind.ValueString(); got != "VirtualService" {
		t.Errorf("metadata.kind = %q, want VirtualService", got)
	}

	out, d := expandIstioObject(ctx, &m)
	if d.HasError() {
		t.Fatalf("expand: %v", d)
	}
	gotVS := &networking.VirtualService{}
	if err := out.Spec.UnmarshalTo(gotVS); err != nil {
		t.Fatalf("unmarshal round-tripped spec: %v", err)
	}
	if !proto.Equal(vs, gotVS) {
		t.Errorf("spec round-trip mismatch:\n in:  %v\n out: %v", vs, gotVS)
	}
}
