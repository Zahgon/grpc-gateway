package server

import (
	"context"

	examples "github.com/grpc-ecosystem/grpc-gateway/v2/examples/internal/proto/examplepb"
)

// Implements of EchoServiceServer

type echoServer struct{}

func newEchoServer() examples.EchoServiceServer {
	_ = "STUB: not implemented"
	return *new(examples.EchoServiceServer)
}

func (s *echoServer) Echo(ctx context.Context, msg *examples.SimpleMessage) (*examples.SimpleMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *echoServer) EchoBody(ctx context.Context, msg *examples.SimpleMessage) (*examples.SimpleMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *echoServer) EchoDelete(ctx context.Context, msg *examples.SimpleMessage) (*examples.SimpleMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *echoServer) EchoPatch(ctx context.Context, msg *examples.DynamicMessageUpdate) (*examples.DynamicMessageUpdate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *echoServer) EchoUnauthorized(ctx context.Context, msg *examples.SimpleMessage) (*examples.SimpleMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EchoStatus demonstrates handling of two Status messages with the same name from different packages.
func (s *echoServer) EchoStatus(ctx context.Context, req *examples.StatusCheckRequest) (*examples.StatusCheckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
