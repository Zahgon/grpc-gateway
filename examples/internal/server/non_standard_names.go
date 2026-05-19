package server

import (
	"context"

	examples "github.com/grpc-ecosystem/grpc-gateway/v2/examples/internal/proto/examplepb"
)

// Implements NonStandardServiceServer

type nonStandardServer struct{}

func newNonStandardServer() examples.NonStandardServiceServer {
	_ = "STUB: not implemented"
	return *new(examples.NonStandardServiceServer)
}

func (s *nonStandardServer) Update(ctx context.Context, msg *examples.NonStandardUpdateRequest) (*examples.NonStandardMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The fieldmask_helper doesn't generate nested structs if they are nil

func (s *nonStandardServer) UpdateWithJSONNames(ctx context.Context, msg *examples.NonStandardWithJSONNamesUpdateRequest) (*examples.NonStandardMessageWithJSONNames, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The fieldmask_helper doesn't generate nested structs if they are nil
