package runtime

import (
	"context"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ErrorHandlerFunc is the signature used to configure error handling.
type ErrorHandlerFunc func(context.Context, *ServeMux, Marshaler, http.ResponseWriter, *http.Request, error)

// StreamErrorHandlerFunc is the signature used to configure stream error handling.
type StreamErrorHandlerFunc func(context.Context, error) *status.Status

// RoutingErrorHandlerFunc is the signature used to configure error handling for routing errors.
type RoutingErrorHandlerFunc func(context.Context, *ServeMux, Marshaler, http.ResponseWriter, *http.Request, int)

// HTTPStatusError is the error to use when needing to provide a different HTTP status code for an error
// passed to the DefaultRoutingErrorHandler.
type HTTPStatusError struct {
	HTTPStatus int
	Err        error
}

func (e *HTTPStatusError) Error() string { _ = "STUB: not implemented"; return "" }

// HTTPStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func HTTPStatusFromCode(code codes.Code) int { _ = "STUB: not implemented"; return 0 }

// Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.

// HTTPError uses the mux-configured error handler.
func HTTPError(ctx context.Context, mux *ServeMux, marshaler Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	_ = "STUB: not implemented"
	return
}

// HTTPStreamError uses the mux-configured stream error handler to notify error to the client without closing the connection.
func HTTPStreamError(ctx context.Context, mux *ServeMux, marshaler Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	_ = "STUB: not implemented"
	return
}

// DefaultHTTPErrorHandler is the default error handler.
// If "err" is a gRPC Status, the function replies with the status code mapped by HTTPStatusFromCode.
// If "err" is a HTTPStatusError, the function replies with the status code provide by that struct. This is
// intended to allow passing through of specific statuses via the function set via WithRoutingErrorHandler
// for the ServeMux constructor to handle edge cases which the standard mappings in HTTPStatusFromCode
// are insufficient for.
// If otherwise, it replies with http.StatusInternalServerError.
//
// The response body written by this function is a Status message marshaled by the Marshaler.
func DefaultHTTPErrorHandler(ctx context.Context, mux *ServeMux, marshaler Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	_ = "STUB: not implemented"
	// return Internal when Marshal failed
	return
}

// RFC 7230 https://tools.ietf.org/html/rfc7230#section-4.1.2
// Unless the request includes a TE header field indicating "trailers"
// is acceptable, as described in Section 4.3, a server SHOULD NOT
// generate trailer fields that it believes are necessary for the user
// agent to receive.

func DefaultStreamErrorHandler(_ context.Context, err error) *status.Status {
	_ = "STUB: not implemented"
	return nil

	// DefaultRoutingErrorHandler is our default handler for routing errors.
	// By default http error codes mapped on the following error codes:
	//
	//	NotFound -> grpc.NotFound
	//	StatusBadRequest -> grpc.InvalidArgument
	//	MethodNotAllowed -> grpc.Unimplemented
	//	Other -> grpc.Internal, method is not expecting to be called for anything else
}

func DefaultRoutingErrorHandler(ctx context.Context, mux *ServeMux, marshaler Marshaler, w http.ResponseWriter, r *http.Request, httpStatus int) {
	_ = "STUB: not implemented"
	return
}
