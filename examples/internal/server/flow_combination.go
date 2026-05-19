package server

import (
	"context"

	examples "github.com/grpc-ecosystem/grpc-gateway/v2/examples/internal/proto/examplepb"
)

type flowCombinationServer struct{}

func newFlowCombinationServer() examples.FlowCombinationServer {
	_ = "STUB: not implemented"
	return *new(examples.FlowCombinationServer)
}

func (s flowCombinationServer) RpcEmptyRpc(ctx context.Context, req *examples.EmptyProto) (*examples.EmptyProto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s flowCombinationServer) RpcEmptyStream(req *examples.EmptyProto, stream examples.FlowCombination_RpcEmptyStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s flowCombinationServer) StreamEmptyRpc(stream examples.FlowCombination_StreamEmptyRpcServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s flowCombinationServer) StreamEmptyStream(stream examples.FlowCombination_StreamEmptyStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s flowCombinationServer) RpcBodyRpc(ctx context.Context, req *examples.NonEmptyProto) (*examples.EmptyProto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s flowCombinationServer) RpcPathSingleNestedRpc(ctx context.Context, req *examples.SingleNestedProto) (*examples.EmptyProto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s flowCombinationServer) RpcPathNestedRpc(ctx context.Context, req *examples.NestedProto) (*examples.EmptyProto, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s flowCombinationServer) RpcBodyStream(req *examples.NonEmptyProto, stream examples.FlowCombination_RpcBodyStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s flowCombinationServer) RpcPathSingleNestedStream(req *examples.SingleNestedProto, stream examples.FlowCombination_RpcPathSingleNestedStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s flowCombinationServer) RpcPathNestedStream(req *examples.NestedProto, stream examples.FlowCombination_RpcPathNestedStreamServer) error {
	_ = "STUB: not implemented"
	return nil
}
