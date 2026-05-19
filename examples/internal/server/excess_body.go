package server

import (
	"context"

	examples "github.com/grpc-ecosystem/grpc-gateway/v2/examples/internal/proto/examplepb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	excessBody_contextChRPC    = make(chan context.Context)
	excessBody_contextChStream = make(chan context.Context)
)

func ExcessBodyServer_RetrieveContextRPC() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func ExcessBodyServer_RetrieveContextStream() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type excessBodyServer struct{}

func newExcessBodyServer() examples.ExcessBodyServiceServer {
	_ = "STUB: not implemented"
	return *new(examples.ExcessBodyServiceServer)
}

func (s excessBodyServer) NoBodyRpc(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s excessBodyServer) NoBodyServerStream(req *emptypb.Empty, stream grpc.ServerStreamingServer[emptypb.Empty]) error {
	_ = "STUB: not implemented"
	return nil
}

func (s excessBodyServer) WithBodyRpc(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s excessBodyServer) WithBodyServerStream(req *emptypb.Empty, stream grpc.ServerStreamingServer[emptypb.Empty]) error {
	_ = "STUB: not implemented"
	return nil
}
