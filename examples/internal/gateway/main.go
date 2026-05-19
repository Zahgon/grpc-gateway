package gateway

import (
	"context"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// Endpoint describes a gRPC endpoint
type Endpoint struct {
	Network, Addr string
}

// Options is a set of options to be passed to Run
type Options struct {
	// Addr is the address to listen
	Addr string

	// GRPCServer defines an endpoint of a gRPC service
	GRPCServer Endpoint

	// OpenAPIDir is a path to a directory from which the server
	// serves OpenAPI specs.
	OpenAPIDir string

	// Mux is a list of options to be passed to the gRPC-Gateway multiplexer
	Mux []gwruntime.ServeMuxOption
}

// Run starts a HTTP server and blocks while running if successful.
// The server will be shutdown when "ctx" is canceled.
func Run(ctx context.Context, opts Options) error { _ = "STUB: not implemented"; return nil }

// Do not use logRequestBody for ExcessBodyServer because it will perform
// io.ReadAll and mask the issue:
// https://github.com/grpc-ecosystem/grpc-gateway/issues/5236
