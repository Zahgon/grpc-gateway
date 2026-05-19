package server

import (
	"context"
	"sync"

	examples "github.com/grpc-ecosystem/grpc-gateway/v2/examples/internal/proto/examplepb"
	"github.com/grpc-ecosystem/grpc-gateway/v2/examples/internal/proto/oneofenum"
	"github.com/grpc-ecosystem/grpc-gateway/v2/examples/internal/proto/pathenum"
	"github.com/grpc-ecosystem/grpc-gateway/v2/examples/internal/proto/sub"
	"github.com/grpc-ecosystem/grpc-gateway/v2/examples/internal/proto/sub2"
	"github.com/rogpeppe/fastuuid"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// Implements of ABitOfEverythingServiceServer

var uuidgen = fastuuid.MustNewGenerator()

type _ABitOfEverythingServer struct {
	v map[string]*examples.ABitOfEverything
	m sync.Mutex
}

type ABitOfEverythingServer interface {
	examples.ABitOfEverythingServiceServer
	examples.StreamServiceServer
}

func newABitOfEverythingServer() ABitOfEverythingServer {
	_ = "STUB: not implemented"
	return *new(ABitOfEverythingServer)
}

func (s *_ABitOfEverythingServer) Create(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) CreateBody(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) CreateBook(ctx context.Context, req *examples.CreateBookRequest) (*examples.Book, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) UpdateBook(ctx context.Context, req *examples.UpdateBookRequest) (*examples.Book, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) BulkCreate(stream examples.StreamService_BulkCreateServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ABitOfEverythingServer) Lookup(ctx context.Context, msg *sub2.IdMessage) (*examples.ABitOfEverything, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) List(opt *examples.Options, stream examples.StreamService_ListServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ABitOfEverythingServer) Download(opt *examples.Options, stream examples.StreamService_DownloadServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ABitOfEverythingServer) Custom(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) DoubleColon(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) Update(ctx context.Context, msg *examples.ABitOfEverything) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) UpdateV2(ctx context.Context, msg *examples.UpdateV2Request) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"

	// If there is no update mask do a regular update
	return nil, nil
}

func (s *_ABitOfEverythingServer) Delete(ctx context.Context, msg *sub2.IdMessage) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) GetQuery(ctx context.Context, msg *examples.ABitOfEverything) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) GetRepeatedQuery(ctx context.Context, msg *examples.ABitOfEverythingRepeated) (*examples.ABitOfEverythingRepeated, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) Echo(ctx context.Context, msg *sub.StringMessage) (*sub.StringMessage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) BulkEcho(stream examples.StreamService_BulkEchoServer) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ABitOfEverythingServer) BulkEchoDuration(stream examples.StreamService_BulkEchoDurationServer) error {
	_ = "STUB: not implemented"
	return nil
}

// Channel to coordinate between read and write goroutines

// Sleep to mock the delay in receiving the request close.
// Accommodates the integration test client which is not a true
// bidirectional streaming client that supports request streaming.

func (s *_ABitOfEverythingServer) DeepPathEcho(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) NoBindings(ctx context.Context, msg *durationpb.Duration) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) Timeout(ctx context.Context, msg *emptypb.Empty) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) ErrorWithDetails(ctx context.Context, msg *emptypb.Empty) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) GetMessageWithBody(ctx context.Context, msg *examples.MessageWithBody) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) PostWithEmptyBody(ctx context.Context, msg *examples.Body) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) CheckGetQueryParams(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) CheckNestedEnumGetQueryParams(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) CheckPostQueryParams(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) OverwriteRequestContentType(ctx context.Context, msg *examples.Body) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) OverwriteResponseContentType(ctx context.Context, msg *emptypb.Empty) (*wrapperspb.StringValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) CheckExternalPathEnum(ctx context.Context, msg *pathenum.MessageWithPathEnum) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) CheckExternalNestedPathEnum(ctx context.Context, msg *pathenum.MessageWithNestedPathEnum) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) CheckStatus(ctx context.Context, empty *emptypb.Empty) (*examples.CheckStatusResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) Exists(ctx context.Context, msg *examples.ABitOfEverything) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) CustomOptionsRequest(ctx context.Context, msg *examples.ABitOfEverything) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) TraceRequest(ctx context.Context, msg *examples.ABitOfEverything) (*examples.ABitOfEverything, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) PostOneofEnum(ctx context.Context, msg *oneofenum.OneofEnumMessage) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *_ABitOfEverythingServer) PostRequiredMessageType(ctx context.Context, req *examples.RequiredMessageTypeRequest) (*emptypb.Empty, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
