package server

import (
	"context"

	examples "github.com/grpc-ecosystem/grpc-gateway/v2/examples/internal/proto/examplepb"
)

// Implements of UnannotatedEchoServiceServer

type unannotatedEchoServer struct{}

func newUnannotatedEchoServer() examples.UnannotatedEchoServiceServer {
	_ = "STUB: not implemented"
	return *new(examples.UnannotatedEchoServiceServer)
}

func (s *unannotatedEchoServer) Echo(ctx context.Context, msg *examples.UnannotatedSimpleMessage) (*examples.UnannotatedSimpleMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *unannotatedEchoServer) EchoBody(ctx context.Context, msg *examples.UnannotatedSimpleMessage) (*examples.UnannotatedSimpleMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *unannotatedEchoServer) EchoDelete(ctx context.Context, msg *examples.UnannotatedSimpleMessage) (*examples.UnannotatedSimpleMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *unannotatedEchoServer) EchoNested(ctx context.Context, msg *examples.UnannotatedSimpleMessage) (*examples.UnannotatedSimpleMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
