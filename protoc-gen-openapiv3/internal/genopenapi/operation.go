package genopenapi

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor"
	"google.golang.org/protobuf/types/descriptorpb"
)

// statusSchemaName is the component name we use for the auto-injected
// google.rpc.Status default error response schema.
const statusSchemaName = "google.rpc.Status"

// buildOperation produces an OpenAPI Operation for one HTTP binding of an
// RPC method, registering any referenced schemas with the schema builder.
//
// `bindingIdx` disambiguates operationId when a method has multiple bindings;
// the first binding gets the bare ID, subsequent bindings append `_<idx>`.
// `pathParams` is the synthetic OpenAPI path parameter list produced by
// convertPathTemplate; it may be longer than binding.PathParams when a
// single proto field expanded into multiple wildcards.
func buildOperation(b *schemaBuilder, svc *descriptor.Service, m *descriptor.Method, binding *descriptor.Binding, bindingIdx int, pathParams []pathParam) (*Operation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// buildParameters returns the path and query parameters for an operation.
// Body fields are excluded; path-bound fields are emitted as `in: path` and
// everything else from the request type is emitted as `in: query`.
//
// Path parameters are driven by `pathParams` (the synthetic OpenAPI list from
// convertPathTemplate) rather than by `binding.PathParams` directly, so a
// single proto field that expanded into multiple URL wildcards yields one
// OpenAPI parameter per wildcard. The proto field is looked up via the
// dotted name on each pathParam so the schema, comments, and deprecation
// flag come from the original descriptor.
func buildParameters(b *schemaBuilder, m *descriptor.Method, binding *descriptor.Binding, pathParams []pathParam) []*ParameterRef {
	_ = "STUB: not implemented"
	return nil
}

// Should not happen: the path template referenced a field that
// the binding did not record. Emit an open string parameter so
// the operation is at least syntactically valid.

// Field is hidden by visibility rules, skip.

// Already handled as a path parameter, skip.

// This field is part of the request body, skip.

// queryParameters expands a single request field into one or more query
// parameters:
//
//   - Scalars, enums, well-known types, and repeated scalars/enums/WKTs
//     produce a single parameter using the field's own schema (repeated
//     scalars become an array-typed parameter, matching the runtime's
//     `?tag=a&tag=b` form).
//   - Non-WKT message fields recurse into their own fields with a dotted
//     prefix ("filter.kind", "filter.range.start", ...).
//   - Map fields emit a single parameter named `<name>[<keyType>]`,
//     matching the runtime's `?labels[foo]=bar` form; see mapQueryParameter.
//   - Repeated messages are skipped — they have no natural query-string
//     representation.
//
// Cycles are bounded by the registry's recursion depth: if recursing into a
// message would exceed the limit on the current path, the recursion is
// truncated at that point and a log line is emitted. Truncating beats failing
// because a partially-flattened spec is still useful.
//
// parentDeprecated propagates deprecation from an ancestor field: if a
// message-typed field is deprecated, all flattened child parameters inherit
// the flag even if the nested fields themselves are not marked deprecated.
func (b *schemaBuilder) queryParameters(field *descriptor.Field, prefix string, parentDeprecated bool, cycle *queryCycleChecker) []*ParameterRef {
	_ = "STUB: not implemented"
	return nil
}

// Non-WKT message: distinguish maps, repeated messages, and regular
// nested messages. An unresolved message falls through to the
// single-parameter path so the operation stays syntactically valid.

// Map fields: emit a single parameter using the runtime's
// `name[key]` form, where the bracketed token documents the
// expected key type. The parameter's schema is the map value's
// schema. Unsupported key types (float, double, bytes) are
// dropped with a log line — the runtime can't key URLs by them.

// Repeated messages cannot be naturally represented in a query string.

// Field is hidden by visibility rules, skip.

// mapQueryParameter constructs a single query parameter for a proto map
// field. The parameter name is `<jsonName>[<keyType>]`, matching the
// `name[key]` URL form the runtime parses (see
// runtime.PopulateQueryParameters — any real key is accepted between the
// brackets; the token in the spec just documents the expected key type).
// The parameter schema is the map value's schema.
//
// Unsupported key types — float, double, bytes — are dropped with a log
// line: the runtime can't key URL parameters by them.
func mapQueryParameter(b *schemaBuilder, field *descriptor.Field, name string, deprecated bool, entry *descriptor.Message) []*ParameterRef {
	_ = "STUB: not implemented"
	return nil
}

// queryMapKeyType returns the JSON Schema type name used to document the
// expected key type of a map query parameter. float/double/bytes keys are
// intentionally unsupported: they don't survive round-tripping through a URL
// query string cleanly enough to be useful.
func queryMapKeyType(t descriptorpb.FieldDescriptorProto_Type) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// queryCycleChecker bounds recursion depth when flattening nested message
// fields into query parameters. It tracks how many times each message has
// been entered on the current path; the same message may appear multiple
// times across sibling branches but only up to `limit` times along any one
// chain of recursive calls.
type queryCycleChecker struct {
	depth map[string]int
	limit int
}

func newQueryCycleChecker(limit int) *queryCycleChecker { _ = "STUB: not implemented"; return nil }

// enter records a recursion into the named message. It returns false (and
// does not record) if entering would exceed the configured limit, so the
// caller can stop without an unbalanced leave.
func (c *queryCycleChecker) enter(fqmn string) bool { _ = "STUB: not implemented"; return false }

func (c *queryCycleChecker) leave(fqmn string) {
	_ = "STUB: not implemented"

	// buildRequestBody constructs the requestBody for a binding. There are two
	// shapes:
	//
	//   - body="*": the body is the entire request message minus path parameters.
	//     We synthesize an inline object schema rather than referencing the
	//     request component, because the component still includes the path
	//     fields.
	//   - body="some_field": the body is just that field's type.
	return
}

func buildRequestBody(b *schemaBuilder, m *descriptor.Method, binding *descriptor.Binding) *RequestBodyRef {
	_ = "STUB: not implemented"
	return nil
}

// body="*": synthesize a body-only inline schema.

// Field is hidden by visibility rules, skip.

// body="field_name": the body is the type of that single field.

// buildResponses constructs the responses map for an RPC: a 200 with the
// response message schema, and (unless disabled) a default google.rpc.Status
// error response.
//
// RPCs returning google.protobuf.Empty get a 200 with an empty object
// schema rather than the HTTP-conventional 204 No Content. The grpc-gateway
// runtime writes `{}` on success regardless of the response type, so the
// spec has to match that or generated clients will reject valid responses.
func buildResponses(b *schemaBuilder, m *descriptor.Method) *Responses {
	_ = "STUB: not implemented"
	return nil
}

// ensureStatusSchema makes sure the google.rpc.Status component schema is
// present.
func ensureStatusSchema(doc *Document) { _ = "STUB: not implemented"; return }

// operationID returns the OpenAPI operationId for a method binding. We use
// "<Service>_<Method>" with a numeric suffix on bindings beyond the first.
func operationID(svc *descriptor.Service, m *descriptor.Method, bindingIdx int) string {
	_ = "STUB: not implemented"
	return ""
}

func needsRequestBody(method string) bool { _ = "STUB: not implemented"; return false }

func isPathParam(field *descriptor.Field, params []descriptor.Parameter) bool {
	_ = "STUB: not implemented"
	return false
}

func isBodyField(field *descriptor.Field, body *descriptor.Body) bool {
	_ = "STUB: not implemented"
	return false
}

// body="*"
