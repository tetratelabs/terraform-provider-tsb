// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

// The Istio direct-mode resources (tsb_istio_*) carry their spec as a
// google.protobuf.Any holding an istio.io/api proto. Blank-importing the istio
// API packages registers those concrete types in the global proto registry so
// protojson can resolve the Any when converting the spec to/from its JSON-string
// representation (jsontypes.Normalized). New Istio kinds only need their package
// added here.
import (
	_ "istio.io/api/extensions/v1alpha1" // WasmPlugin
	_ "istio.io/api/networking/v1alpha3" // VirtualService, Gateway, DestinationRule, ServiceEntry, Sidecar, EnvoyFilter
	_ "istio.io/api/networking/v1beta1"  // (newer networking kinds)
	_ "istio.io/api/security/v1beta1"    // AuthorizationPolicy, RequestAuthentication, PeerAuthentication
)
