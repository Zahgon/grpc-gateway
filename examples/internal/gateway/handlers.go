package gateway

import (
	"net/http"

	"google.golang.org/grpc"
)

// openAPIServer returns OpenAPI specification files located under "/openapiv2/"
func openAPIServer(dir string) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// allowCORS allows Cross Origin Resource Sharing from any origin.
// Don't do this without consideration in production systems.
func allowCORS(h http.Handler) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

// healthzServer returns a simple health handler which returns ok.
func healthzServer(conn *grpc.ClientConn) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// Invoke method to move connection from Idle to Ready

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rsp *logResponseWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }

// Unwrap returns the original http.ResponseWriter. This is necessary
// to expose Flush() and Push() on the underlying response writer.
func (rsp *logResponseWriter) Unwrap() http.ResponseWriter {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter)
}

func newLogResponseWriter(w http.ResponseWriter) *logResponseWriter {
	_ = "STUB: not implemented"
	return nil
}

// logRequestBody logs the request body when the response status code is not 200.
// This addresses the issue of being unable to retrieve the request body in the customErrorHandler middleware.
func logRequestBody(h http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}
