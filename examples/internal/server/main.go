package server

import (
	"context"
	"net"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// Run starts the example gRPC service.
// "network" and "address" are passed to net.Listen.
func Run(ctx context.Context, network, address string) error { _ = "STUB: not implemented"; return nil }

// ServeGRPC registers the example gRPC services on a fresh grpc.Server and
// serves the given listener until ctx is cancelled. Callers retain
// ownership of the listener and are responsible for closing it.
func ServeGRPC(ctx context.Context, l net.Listener) error { _ = "STUB: not implemented"; return nil }

// RunInProcessGateway starts the invoke in process http gateway.
func RunInProcessGateway(ctx context.Context, addr string, opts ...runtime.ServeMuxOption) error {
	_ = "STUB: not implemented"
	return nil
}
