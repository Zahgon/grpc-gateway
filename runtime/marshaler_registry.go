package runtime

import (
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
)

// MIMEWildcard is the fallback MIME type used for requests which do not match
// a registered MIME type.
const MIMEWildcard = "*"

var (
	acceptHeader      = http.CanonicalHeaderKey("Accept")
	contentTypeHeader = http.CanonicalHeaderKey("Content-Type")

	defaultMarshaler = &HTTPBodyMarshaler{
		Marshaler: &JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	}
)

// MarshalerForRequest returns the inbound/outbound marshalers for this request.
// It checks the registry on the ServeMux for the MIME type set by the Content-Type header.
// If it isn't set (or the request Content-Type is empty), checks for "*".
// If there are multiple Content-Type headers set, choose the first one that it can
// exactly match in the registry.
// Otherwise, it follows the above logic for "*"/InboundMarshaler/OutboundMarshaler.
func MarshalerForRequest(mux *ServeMux, r *http.Request) (inbound Marshaler, outbound Marshaler) {
	_ = "STUB: not implemented"
	return *new(Marshaler), *new(Marshaler)
}

// marshalerRegistry is a mapping from MIME types to Marshalers.
type marshalerRegistry struct {
	mimeMap map[string]Marshaler
}

// add adds a marshaler for a case-sensitive MIME type string ("*" to match any
// MIME type).
func (m marshalerRegistry) add(mime string, marshaler Marshaler) error {
	_ = "STUB: not implemented"
	return nil
}

// makeMarshalerMIMERegistry returns a new registry of marshalers.
// It allows for a mapping of case-sensitive Content-Type MIME type string to runtime.Marshaler interfaces.
//
// For example, you could allow the client to specify the use of the runtime.JSONPb marshaler
// with an "application/jsonpb" Content-Type and the use of the runtime.JSONBuiltin marshaler
// with an "application/json" Content-Type.
// "*" can be used to match any Content-Type.
// This can be attached to a ServerMux with the marshaler option.
func makeMarshalerMIMERegistry() marshalerRegistry {
	_ = "STUB: not implemented"
	return *new(marshalerRegistry)
}

// WithMarshalerOption returns a ServeMuxOption which associates inbound and outbound
// Marshalers to a MIME type in mux.
func WithMarshalerOption(mime string, marshaler Marshaler) ServeMuxOption {
	_ = "STUB: not implemented"
	return *new(ServeMuxOption)
}
