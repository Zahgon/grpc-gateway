package gateway

import (
	"context"
	"net/http"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
)

// newGateway returns a new gateway server which translates HTTP into gRPC.
func newGateway(ctx context.Context, conn *grpc.ClientConn, opts []gwruntime.ServeMuxOption) (http.Handler, error) {
	_ = "STUB: not implemented"
	return *new(http.Handler), nil
}

func dial(network, addr string) (*grpc.ClientConn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// dialTCP creates a client connection via TCP.
// "addr" must be a valid TCP address with a port number.
func dialTCP(addr string) (*grpc.ClientConn, error) { _ = "STUB: not implemented"; return nil, nil }

// dialUnix creates a client connection via a unix domain socket.
// "addr" must be a valid path to the socket.
func dialUnix(addr string) (*grpc.ClientConn, error) { _ = "STUB: not implemented"; return nil, nil }
