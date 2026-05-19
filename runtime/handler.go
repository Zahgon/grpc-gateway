package runtime

import (
	"context"
	"net/http"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// ForwardResponseStream forwards the stream from gRPC server to REST client.
func ForwardResponseStream(ctx context.Context, mux *ServeMux, marshaler Marshaler, w http.ResponseWriter, req *http.Request, recv func() (proto.Message, error), opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	_ = "STUB: not implemented"
	return
}

func handleForwardResponseServerMetadata(w http.ResponseWriter, mux *ServeMux, md ServerMetadata) {
	_ = "STUB: not implemented"
	return
}

func handleForwardResponseTrailerHeader(w http.ResponseWriter, mux *ServeMux, md ServerMetadata) {
	_ = "STUB: not implemented"
	return
}

func handleForwardResponseTrailer(w http.ResponseWriter, mux *ServeMux, md ServerMetadata) {
	_ = "STUB: not implemented"
	return
}

// responseBody interface contains method for getting field for marshaling to the response body
// this method is generated for response struct from the value of `response_body` in the `google.api.HttpRule`
type responseBody interface {
	XXX_ResponseBody() interface{}
}

// ForwardResponseMessage forwards the message "resp" from gRPC server to REST client.
func ForwardResponseMessage(ctx context.Context, mux *ServeMux, marshaler Marshaler, w http.ResponseWriter, req *http.Request, resp proto.Message, opts ...func(context.Context, http.ResponseWriter, proto.Message) error) {
	_ = "STUB: not implemented"
	return
}

// RFC 7230 https://tools.ietf.org/html/rfc7230#section-4.1.2
// Unless the request includes a TE header field indicating "trailers"
// is acceptable, as described in Section 4.3, a server SHOULD NOT
// generate trailer fields that it believes are necessary for the user
// agent to receive.

func requestAcceptsTrailers(req *http.Request) bool { _ = "STUB: not implemented"; return false }

func handleForwardResponseOptions(ctx context.Context, w http.ResponseWriter, resp proto.Message, opts []func(context.Context, http.ResponseWriter, proto.Message) error) error {
	_ = "STUB: not implemented"
	return nil
}

func handleForwardResponseStreamError(ctx context.Context, wroteHeader bool, marshaler Marshaler, w http.ResponseWriter, req *http.Request, mux *ServeMux, err error, delimiter []byte) {
	_ = "STUB: not implemented"
	return
}

func errorChunk(st *status.Status) map[string]proto.Message { _ = "STUB: not implemented"; return nil }
