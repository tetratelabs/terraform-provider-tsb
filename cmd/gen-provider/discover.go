// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

// resourceInfo is the discovered description of a single manageable TSB resource
// (one that exposes the standard Create/Get/Update/Delete quartet over an
// FQN-named entity).
type resourceInfo struct {
	// Stem is the method-name stem, e.g. "Workspace" or "Settings".
	Stem string
	// Group is the API group segment, e.g. "gateway"; empty for the core
	// tetrateio.api.tsb.v2 package.
	Group string
	// BaseName is the unique Go identifier base for this resource's generated
	// resource/constructor symbols (group-qualified, e.g. "GatewayGroup"; for
	// shared-message resources, service-qualified, e.g. "IstioGatewayVirtualService").
	BaseName string
	// TFName is the Terraform type suffix, e.g. "workspace" or "gateway_group".
	TFName string
	// ServiceName is the gRPC service Go name, e.g. "Workspaces" or "IstioGateway".
	ServiceName string
	// SharedMessage is true when several resources are backed by the same entity
	// message (e.g. IstioObject for the Istio direct-mode resources); such
	// resources share one model and are identified by service+stem instead.
	SharedMessage bool
	// Collection is the FQN collection segment, e.g. "workspaces".
	Collection string

	// MessageGoType is the resource message Go type (pointer-elem struct).
	MessageGoType reflect.Type
	// MessageDesc is the resource message proto descriptor.
	MessageDesc protoreflect.MessageDescriptor

	// ClientCtor is the gRPC client constructor, e.g. "NewWorkspacesClient".
	ClientCtor string
	// ServicePkg is the Go import path the client constructor lives in.
	ServicePkg string

	CreateMethod, GetMethod, UpdateMethod, DeleteMethod string

	// CreateReqGoType is the Go type of the Create request message.
	CreateReqGoType reflect.Type
	// CreateMsgField is the Go field name on the Create request holding the message.
	CreateMsgField string
	// CreateHasParent / CreateHasName indicate whether the Create request carries
	// Parent / Name fields (root resources like Organization have no Parent).
	CreateHasParent, CreateHasName bool
	// GetReqGoType / DeleteReqGoType are the Go request types for Get/Delete.
	GetReqGoType, DeleteReqGoType reflect.Type
	// DeleteHasForce indicates the Delete request carries a Force field.
	DeleteHasForce bool
}

var (
	fqnCollectionRe = regexp.MustCompile(`(?:^|/)([a-z0-9-]+)/\*$`)
	versionSegRe    = regexp.MustCompile(`^v\d`)
)

// tsbPackagePrefix scopes discovery to the TSB API surface. Every registered
// package under it is scanned structurally, so new groups/resources are picked
// up without code changes (the group package only needs to be imported; see
// imports.go).
const tsbPackagePrefix = "tetrateio.api.tsb."

// discoverResources walks all registered TSB file descriptors and returns the
// resources that expose the standard CRUD quartet. Per-resource discovery
// failures are reported as warnings and skipped rather than aborting.
func discoverResources() ([]resourceInfo, error) {
	var out []resourceInfo

	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		if !strings.HasPrefix(string(fd.Package()), tsbPackagePrefix) {
			return true
		}
		svcs := fd.Services()
		for i := 0; i < svcs.Len(); i++ {
			svc := svcs.Get(i)
			methods := indexMethods(svc)
			for stem, upd := range updateMethods(svc) {
				ri, ok, err := buildResource(svc, stem, upd, methods)
				if err != nil {
					fmt.Fprintf(os.Stderr, "gen-provider: skipping candidate %s/%s: %v\n", fd.Package(), stem, err)
					continue
				}
				if ok {
					out = append(out, ri)
				}
			}
		}
		return true
	})

	markSharedMessages(out)
	sort.Slice(out, func(i, j int) bool { return out[i].TFName < out[j].TFName })
	return out, nil
}

// markSharedMessages handles resources whose entity message backs more than one
// CRUD endpoint (e.g. tsb/types/v2.IstioObject, the raw Istio config wrapper
// reused for VirtualService/Gateway/… across the Istio direct-mode services).
// These cannot derive a unique identity from the message, so they are renamed
// after the service + method stem (e.g. tsb_istio_gateway_virtual_service) and
// flagged so the emitter gives them a shared model.
func markSharedMessages(in []resourceInfo) {
	counts := map[reflect.Type]int{}
	for _, r := range in {
		counts[r.MessageGoType]++
	}
	for i := range in {
		if counts[in[i].MessageGoType] <= 1 {
			continue
		}
		in[i].SharedMessage = true
		in[i].BaseName = in[i].ServiceName + in[i].Stem
		in[i].TFName = snakeCase(in[i].ServiceName) + "_" + snakeCase(in[i].Stem)
	}
}

// goPackageOf returns the Go import path declared by a proto file's go_package
// option (stripping any ";alias" suffix). Used to qualify a service's gRPC client
// constructor, which lives in the service file's package — not necessarily the
// entity message's package (true for the Istio direct-mode services).
func goPackageOf(fd protoreflect.FileDescriptor) string {
	opts, ok := fd.Options().(*descriptorpb.FileOptions)
	if !ok || opts == nil {
		return ""
	}
	gp := opts.GetGoPackage()
	if i := strings.IndexByte(gp, ';'); i >= 0 {
		gp = gp[:i]
	}
	return gp
}

// schemaPackages are the non-tsb proto package prefixes whose messages have a
// generated XxxSchema() (their protos are listed in buf.gen.terraform.yaml).
var schemaPackages = []string{
	tsbPackagePrefix,
	"tetrateio.api.install.dataplane.",
}

// hasGeneratedSchema reports whether a message package has reusable schema funcs.
func hasGeneratedSchema(pkg protoreflect.FullName) bool {
	for _, p := range schemaPackages {
		if strings.HasPrefix(string(pkg), p) {
			return true
		}
	}
	return false
}

// groupFromPackage returns the API group segment for a proto package full name
// (e.g. "tetrateio.api.tsb.gateway.v2" -> "gateway"); the core
// "tetrateio.api.tsb.v2" package returns "".
func groupFromPackage(pkg protoreflect.FullName) string {
	rest := strings.TrimPrefix(string(pkg), tsbPackagePrefix)
	// rest is like "gateway.v2" or "v2" or "observability.telemetry.v2".
	segs := strings.Split(rest, ".")
	if len(segs) > 0 && versionSegRe.MatchString(segs[len(segs)-1]) {
		segs = segs[:len(segs)-1]
	}
	return strings.Join(segs, "_")
}

// camelGroup converts a group segment ("gateway", "observability_telemetry") to
// a Go identifier prefix ("Gateway", "ObservabilityTelemetry").
func camelGroup(group string) string {
	var b strings.Builder
	for _, part := range strings.Split(group, "_") {
		if part == "" {
			continue
		}
		b.WriteString(strings.ToUpper(part[:1]) + part[1:])
	}
	return b.String()
}

// indexMethods returns the service methods keyed by their Go name.
func indexMethods(svc protoreflect.ServiceDescriptor) map[string]protoreflect.MethodDescriptor {
	m := map[string]protoreflect.MethodDescriptor{}
	ms := svc.Methods()
	for i := 0; i < ms.Len(); i++ {
		md := ms.Get(i)
		m[string(md.Name())] = md
	}
	return m
}

// updateMethods returns Update<Stem> methods whose input == output (the resource
// entity), keyed by stem.
func updateMethods(svc protoreflect.ServiceDescriptor) map[string]protoreflect.MethodDescriptor {
	out := map[string]protoreflect.MethodDescriptor{}
	ms := svc.Methods()
	for i := 0; i < ms.Len(); i++ {
		md := ms.Get(i)
		name := string(md.Name())
		if !strings.HasPrefix(name, "Update") {
			continue
		}
		if md.Input().FullName() != md.Output().FullName() {
			continue
		}
		out[strings.TrimPrefix(name, "Update")] = md
	}
	return out
}

func buildResource(
	svc protoreflect.ServiceDescriptor,
	stem string,
	upd protoreflect.MethodDescriptor,
	methods map[string]protoreflect.MethodDescriptor,
) (resourceInfo, bool, error) {
	create, hasCreate := methods["Create"+stem]
	get, hasGet := methods["Get"+stem]
	del, hasDelete := methods["Delete"+stem]
	if !hasCreate || !hasGet || !hasDelete {
		return resourceInfo{}, false, nil
	}

	// The entity message must have a generated XxxSchema() to reuse. TSB groups
	// always do; install/dataplane is included explicitly (its protos are in
	// buf.gen.terraform.yaml) so the InstallGateway entity (GatewaySpec) is
	// supported. Other foreign entities are skipped.
	msgPkg := upd.Input().ParentFile().Package()
	if !hasGeneratedSchema(msgPkg) {
		return resourceInfo{}, false, nil
	}

	msgType, err := goType(upd.Input().FullName())
	if err != nil {
		return resourceInfo{}, false, err
	}

	collection := collectionFromFqn(httpTemplate(get))
	if collection == "" {
		collection = collectionFromFqn(httpTemplate(upd))
	}
	if collection == "" {
		return resourceInfo{}, false, fmt.Errorf("resource %s: cannot derive collection segment from http template", stem)
	}

	createReqType, err := goType(create.Input().FullName())
	if err != nil {
		return resourceInfo{}, false, err
	}
	msgField, err := messageField(createReqType, msgType)
	if err != nil {
		return resourceInfo{}, false, fmt.Errorf("resource %s: %w", stem, err)
	}

	getReqType, err := goType(get.Input().FullName())
	if err != nil {
		return resourceInfo{}, false, err
	}
	delReqType, err := goType(del.Input().FullName())
	if err != nil {
		return resourceInfo{}, false, err
	}

	// Group and name derive from the SERVICE (where the CRUD lives). When the
	// entity lives in a different package than the service (e.g.
	// install/dataplane.GatewaySpec under the gateway service), the proto message
	// name is not meaningful, so use the method stem ("InstallGateway") instead.
	svcPkg := svc.ParentFile().Package()
	group := groupFromPackage(svcPkg)
	nameToken := string(upd.Input().Name())
	if msgPkg != svcPkg {
		nameToken = stem
	}
	baseName := nameToken
	tfName := snakeCase(nameToken)
	if group != "" {
		baseName = camelGroup(group) + nameToken
		tfName = group + "_" + snakeCase(nameToken)
	}

	// The gRPC client constructor lives in the service file's Go package, which
	// differs from the entity message package for shared-message services.
	servicePkg := goPackageOf(svc.ParentFile())
	if servicePkg == "" {
		servicePkg = msgType.PkgPath()
	}

	return resourceInfo{
		Stem:            stem,
		Group:           group,
		BaseName:        baseName,
		TFName:          tfName,
		ServiceName:     string(svc.Name()),
		Collection:      collection,
		MessageGoType:   msgType,
		MessageDesc:     upd.Input(),
		ClientCtor:      "New" + string(svc.Name()) + "Client",
		ServicePkg:      servicePkg,
		CreateMethod:    "Create" + stem,
		GetMethod:       "Get" + stem,
		UpdateMethod:    string(upd.Name()),
		DeleteMethod:    "Delete" + stem,
		CreateReqGoType: createReqType,
		CreateMsgField:  msgField,
		CreateHasParent: hasField(createReqType, "Parent"),
		CreateHasName:   hasField(createReqType, "Name"),
		GetReqGoType:    getReqType,
		DeleteReqGoType: delReqType,
		DeleteHasForce:  hasField(delReqType, "Force"),
	}, true, nil
}

// httpTemplate returns the path template of a method's google.api.http rule
// (whichever verb is set).
func httpTemplate(md protoreflect.MethodDescriptor) string {
	opts, ok := md.Options().(*descriptorpb.MethodOptions)
	if !ok || opts == nil {
		return ""
	}
	rule, ok := proto.GetExtension(opts, annotations.E_Http).(*annotations.HttpRule)
	if !ok || rule == nil {
		return ""
	}
	switch {
	case rule.GetGet() != "":
		return rule.GetGet()
	case rule.GetPut() != "":
		return rule.GetPut()
	case rule.GetPost() != "":
		return rule.GetPost()
	case rule.GetDelete() != "":
		return rule.GetDelete()
	case rule.GetPatch() != "":
		return rule.GetPatch()
	}
	return ""
}

// collectionFromFqn extracts the collection segment from an fqn path template
// such as /v2/{fqn=organizations/*/tenants/*/workspaces/*} -> "workspaces".
func collectionFromFqn(template string) string {
	start := strings.Index(template, "{fqn=")
	if start < 0 {
		return ""
	}
	inner := template[start+len("{fqn="):]
	if end := strings.Index(inner, "}"); end >= 0 {
		inner = inner[:end]
	}
	m := fqnCollectionRe.FindStringSubmatch(inner)
	if len(m) == 2 {
		return m[1]
	}
	return ""
}

// goType resolves a proto message full name to its Go struct type (pointer-elem).
func goType(name protoreflect.FullName) (reflect.Type, error) {
	mt, err := protoregistry.GlobalTypes.FindMessageByName(name)
	if err != nil {
		return nil, fmt.Errorf("cannot resolve Go type for %s: %w", name, err)
	}
	return reflect.TypeOf(mt.New().Interface()).Elem(), nil
}

// messageField returns the Go field name on req whose type is *msg.
func messageField(req, msg reflect.Type) (string, error) {
	want := reflect.PointerTo(msg)
	for i := 0; i < req.NumField(); i++ {
		if req.Field(i).Type == want {
			return req.Field(i).Name, nil
		}
	}
	return "", fmt.Errorf("create request %s has no field of type *%s", req.Name(), msg.Name())
}

func hasField(t reflect.Type, name string) bool {
	_, ok := t.FieldByName(name)
	return ok
}
