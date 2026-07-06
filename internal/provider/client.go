// Copyright (c) Tetrate, Inc 2026 All Rights Reserved.

package provider

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"

	"github.com/tetrateio/tetrate/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	xcredentials "google.golang.org/grpc/experimental/credentials"
)

// Config holds the resolved provider configuration used to dial TSB. It is
// populated by the provider's Configure method from the provider config block
// and/or environment variables, then shared with every resource.
type Config struct {
	// Endpoint is the TSB gRPC address (host:port).
	Endpoint string
	// Token is the TSB API token sent in the x-tetrate-token header. Takes
	// precedence over Username/Password when both are set.
	Token string
	// Username and Password authenticate via HTTP Basic auth (the Authorization
	// header) when no Token is provided.
	Username string
	Password string
	// Organization is the TSB organization the resources live under. It is used
	// to default the `parent` FQN prefix when a resource does not set one.
	Organization string
	// CustomCAFile, when set, is a PEM bundle used as the TLS root CA pool.
	CustomCAFile string
	// TLSDisabled turns off transport security entirely (plaintext).
	TLSDisabled bool
	// TLSInsecure keeps TLS on but skips server certificate verification.
	TLSInsecure bool
}

// client bundles the resolved config and the live gRPC connection. A pointer to
// this is handed to each resource via the provider's ResourceData.
type client struct {
	conn   *grpc.ClientConn
	config Config
}

// dial opens a gRPC connection to TSB using the same transport and auth
// conventions as tctl/teamsync: a per-RPC credential (an x-tetrate-token token,
// or HTTP Basic username/password) and ALPN-disabled TLS (xcredentials) unless
// TLS is turned off.
func (c Config) dial(ctx context.Context) (*grpc.ClientConn, error) {
	credOpt, err := c.transportCredentials()
	if err != nil {
		return nil, err
	}

	// DialContext is deprecated upstream, but the existing TSB tooling (tctl)
	// still relies on it for connection-reliability reasons; we match that here.
	conn, err := grpc.DialContext(ctx, c.Endpoint,
		credOpt,
		grpc.WithPerRPCCredentials(c.perRPCCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to TSB at %q: %w", c.Endpoint, err)
	}
	return conn, nil
}

// perRPCCredentials selects the gRPC call credential: a token takes precedence,
// otherwise HTTP Basic auth from username/password.
func (c Config) perRPCCredentials() credentials.PerRPCCredentials {
	if c.Token != "" {
		return auth.TetrateToken(c.Token)
	}
	return auth.Basic{Username: c.Username, Password: c.Password}.PerRPCCredentials()
}

// transportCredentials mirrors tctl/pkg/bridge/auth.go TransportCredentials.
func (c Config) transportCredentials() (grpc.DialOption, error) {
	if c.TLSDisabled {
		return grpc.WithTransportCredentials(insecure.NewCredentials()), nil
	}

	// #nosec G402 -- InsecureSkipVerify is opt-in via provider configuration.
	tlsConfig := &tls.Config{
		InsecureSkipVerify: c.TLSInsecure,
	}

	if c.CustomCAFile != "" {
		ca, err := os.ReadFile(c.CustomCAFile)
		if err != nil {
			return nil, fmt.Errorf("failed to load custom CA file: %w", err)
		}
		tlsConfig.RootCAs = x509.NewCertPool()
		if ok := tlsConfig.RootCAs.AppendCertsFromPEM(ca); !ok {
			return nil, errors.New("error loading custom CA")
		}
	}

	// ALPN protocols were only added to front envoy in TSB 1.13; disable ALPN to
	// keep compatibility with older TSB instances, matching tctl.
	return grpc.WithTransportCredentials(xcredentials.NewTLSWithALPNDisabled(tlsConfig)), nil
}
