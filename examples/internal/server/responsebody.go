package server

import (
	"context"

	examples "github.com/grpc-ecosystem/grpc-gateway/v2/examples/internal/proto/examplepb"
)

// Implements of ResponseBodyServiceServer

type responseBodyServer struct{}

func newResponseBodyServer() examples.ResponseBodyServiceServer {
	_ = "STUB: not implemented"
	return *new(examples.ResponseBodyServiceServer)
}

func (s *responseBodyServer) GetResponseBody(ctx context.Context, req *examples.ResponseBodyIn) (*examples.ResponseBodyOut, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *responseBodyServer) ListResponseBodies(ctx context.Context, req *examples.ResponseBodyIn) (*examples.RepeatedResponseBodyOut, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *responseBodyServer) ListResponseStrings(ctx context.Context, req *examples.ResponseBodyIn) (*examples.RepeatedResponseStrings, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *responseBodyServer) GetResponseBodyStream(req *examples.ResponseBodyIn, stream examples.ResponseBodyService_GetResponseBodyStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *responseBodyServer) GetResponseBodySameName(ctx context.Context, req *examples.ResponseBodyIn) (*examples.ResponseBodyValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
