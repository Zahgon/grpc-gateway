package genopenapi

import (
	"reflect"
	"regexp"
	"sync"

	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor"
	openapi_options "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/genproto/googleapis/api/visibility"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/known/structpb"
)

// The OpenAPI specification does not allow for more than one endpoint with the same HTTP method and path.
// This prevents multiple gRPC service methods from sharing the same stripped version of the path and method.
// For example: `GET /v1/{name=organizations/*}/roles` and `GET /v1/{name=users/*}/roles` both get stripped to `GET /v1/{name}/roles`.
// We must make the URL unique by adding a suffix and an incrementing index to each path parameter
// to differentiate the endpoints.
// Since path parameter names do not affect the request contents (i.e. they're replaced in the path)
// this will be hidden from the real grpc gateway consumer.
const pathParamUniqueSuffixDeliminator = "_"

const paragraphDeliminator = "\n\n"

// wktSchemas are the schemas of well-known-types.
// The schemas must match with the behavior of the JSON unmarshaler in
// https://github.com/protocolbuffers/protobuf-go/blob/v1.25.0/encoding/protojson/well_known_types.go
var wktSchemas = map[string]schemaCore{
	".google.protobuf.FieldMask": {
		Type: "string",
	},
	".google.protobuf.Timestamp": {
		Type:   "string",
		Format: "date-time",
	},
	".google.protobuf.Duration": {
		Type: "string",
	},
	".google.protobuf.StringValue": {
		Type: "string",
	},
	".google.protobuf.BytesValue": {
		Type:   "string",
		Format: "byte",
	},
	".google.protobuf.Int32Value": {
		Type:   "integer",
		Format: "int32",
	},
	".google.protobuf.UInt32Value": {
		Type:   "integer",
		Format: "int64",
	},
	".google.protobuf.Int64Value": {
		Type:   "string",
		Format: "int64",
	},
	".google.protobuf.UInt64Value": {
		Type:   "string",
		Format: "uint64",
	},
	".google.protobuf.FloatValue": {
		Type:   "number",
		Format: "float",
	},
	".google.protobuf.DoubleValue": {
		Type:   "number",
		Format: "double",
	},
	".google.protobuf.BoolValue": {
		Type: "boolean",
	},
	".google.protobuf.Empty": {
		Type: "object",
	},
	".google.protobuf.Struct": {
		Type: "object",
	},
	".google.protobuf.Value": {},
	".google.protobuf.ListValue": {
		Type: "array",
		Items: (*openapiItemsObject)(&openapiSchemaObject{
			schemaCore: schemaCore{
				Type: "object",
			},
		}),
	},
	".google.protobuf.NullValue": {
		Type: "string",
	},
}

func listEnumNames(reg *descriptor.Registry, enum *descriptor.Enum) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func listEnumNumbers(reg *descriptor.Registry, enum *descriptor.Enum) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func getEnumDefault(reg *descriptor.Registry, enum *descriptor.Enum) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func getEnumDefaultNumber(reg *descriptor.Registry, enum *descriptor.Enum) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// messageToQueryParameters converts a message to a list of OpenAPI query parameters.
func messageToQueryParameters(message *descriptor.Message, reg *descriptor.Registry, pathParams []descriptor.Parameter, body *descriptor.Body, httpMethod string) (params []openapiParameterObject, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When body is set to oneof field, we want to skip other fields in the oneof group.

func isBodySameOneOf(body *descriptor.Body, field *descriptor.Field) bool {
	_ = "STUB: not implemented"
	return false
}

// queryParams converts a field to a list of OpenAPI query parameters recursively through the use of nestedQueryParams.
func queryParams(message *descriptor.Message, field *descriptor.Field, prefix string, reg *descriptor.Registry, pathParams []descriptor.Parameter, body *descriptor.Body, recursiveCount int) (params []openapiParameterObject, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type cycleChecker struct {
	m     map[string]int
	count int
}

func newCycleChecker(recursive int) *cycleChecker { _ = "STUB: not implemented"; return nil }

// Check returns whether name is still within recursion
// toleration
func (c *cycleChecker) Check(name string) bool { _ = "STUB: not implemented"; return false }

// provision map entry if not available

func (c *cycleChecker) Branch() *cycleChecker { _ = "STUB: not implemented"; return nil }

// nestedQueryParams converts a field to a list of OpenAPI query parameters recursively.
// This function is a helper function for queryParams, that keeps track of cyclical message references
// through the use of
//
//	touched map[string]int
//
// If a cycle is discovered, an error is returned, as cyclical data structures are dangerous
// in query parameters.
func nestedQueryParams(message *descriptor.Message, field *descriptor.Field, prefix string, reg *descriptor.Registry, pathParams []descriptor.Parameter, body *descriptor.Body, cycle *cycleChecker) (params []openapiParameterObject, err error) {
	_ = "STUB: not implemented"
	// make sure the parameter is not already listed as a path parameter
	return nil, nil
}

// make sure the parameter is not already listed as a body parameter

// This will generate a query in the format map_name[key_type]

// TODO: currently, mapping object in query parameter is not supported

// verify if the field is required

// verify if the field is required in message options

// Required fields can be field names or json_name values

// verify if the field is deprecated, either via proto or annotation

// array

// nested type, recurse

// Check for cyclical message reference:

// Construct a new map with the message name so a cycle further down the recursive path can be detected.
// Do not keep anything in the original touched reference and do not pass that reference along.  This will
// prevent clobbering adjacent records while recursing.

func getMapParamKey(t descriptorpb.FieldDescriptorProto_Type) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// findServicesMessagesAndEnumerations discovers all messages and enums defined in the RPC methods of the service.
func findServicesMessagesAndEnumerations(s []*descriptor.Service, reg *descriptor.Registry, m messageMap, ms messageMap, e enumMap, refs refMap) {
	_ = "STUB: not implemented"
	return
}

// Request may be fully included in query

// Only process methods with HTTP bindings (exposed via HTTP annotations)
// This prevents unused message definitions from appearing in the OpenAPI document

// findNestedMessagesAndEnumerations those can be generated by the services.
func findNestedMessagesAndEnumerations(message *descriptor.Message, reg *descriptor.Registry, m messageMap, e enumMap) {
	_ = "STUB: not implemented"
	// Iterate over all the fields that
	return
}

// If the type is an empty string then it is a proto primitive

// collectReferencedNamesForCache scans services and messages to collect all
// FQMNs/FQENs that will be referenced, WITHOUT using the naming cache.
// This allows us to build the cache with the correct filtered names BEFORE
// any code tries to use it.
func collectReferencedNamesForCache(services []*descriptor.Service, messages []*descriptor.Message, reg *descriptor.Registry) map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

// Scan services FIRST so collectNestedTypeFQNs fully traverses
// message graphs without being short-circuited by pre-populated entries.

// Add method FQN (needed for body:"*" with path params)

// Add request/response types

// Recursively add nested types

// Add messages from the current file AFTER service scanning.
// This must come after the service loop's collectNestedTypeFQNs calls,
// otherwise pre-populated message entries cause the traversal to
// short-circuit and miss nested types like enums inside referenced messages.
// We also traverse each message's nested types here because
// renderMessagesAsDefinition renders ALL messages from the file, not just
// those reachable from service methods. Without this, cross-package types
// referenced by non-service messages would be missing from the naming cache.

// Add google.rpc.Status if default errors enabled

// Also add nested types of Status

// collectNestedTypeFQNs recursively collects FQMNs/FQENs for all nested types
// of a message. Does NOT use the naming cache.
func collectNestedTypeFQNs(message *descriptor.Message, reg *descriptor.Registry, refs map[string]bool) {
	_ = "STUB: not implemented"
	return
}

// primitive type

// already visited

// If it's a message, recurse

// Enums don't have nested types, no recursion needed

func skipRenderingRef(refName string) bool { _ = "STUB: not implemented"; return false }

func renderMessageAsDefinition(msg *descriptor.Message, reg *descriptor.Registry, customRefs refMap, pathParams []descriptor.Parameter) (openapiSchemaObject, error) {
	_ = "STUB: not implemented"
	return *new(openapiSchemaObject), nil
}

// Warning: Make sure not to overwrite any fields already set on the schema type.

// Only hoist required fields to parent if there are no path params inside this field.

// To avoid populating both the field schema require and message schema require, unset the field schema require.
// See issue #2635.

// When there are path params, we need to separate field-level required from nested required.
// The field name itself (if required) should be in parent's required, but nested field names
// should stay in the nested schema's required.

// Check if the field name is in the fieldSchema.Required (it would be if the field is marked REQUIRED)

// Add the field name to parent's required if the field itself is required

// Keep only the nested required fields in the field schema

// Per the JSON Reference syntax: Any members other than "$ref" in a JSON Reference object SHALL be ignored.
// https://tools.ietf.org/html/draft-pbryan-zyp-json-ref-03#section-3
// However, use allOf to specify Title/Description/Example/readOnly fields.

func renderFieldAsDefinition(f *descriptor.Field, reg *descriptor.Registry, refs refMap, pathParams []descriptor.Parameter) (openapiSchemaObject, error) {
	_ = "STUB: not implemented"
	return *new(openapiSchemaObject), nil
}

// Use title and description from field instead of nested message if present.

// to handle case where path param is present inside the field of descriptorpb.FieldDescriptorProto_TYPE_MESSAGE type
// it still needs to consider the behaviour of the field which was being done by schemaOfField() in case there are no path params

// transformAnyForJSON should be called when the schema object represents a google.protobuf.Any, and will replace the
// Properties slice with a single value for '@type'. We mutate the incorrectly named field so that we inherit the same
// documentation as specified on the original field in the protobuf descriptors.
func transformAnyForJSON(schema *openapiSchemaObject, useJSONNames bool) {
	_ = "STUB: not implemented"
	return
}

func renderMessagesAsDefinition(messages messageMap, d openapiDefinitionsObject, reg *descriptor.Registry, customRefs refMap, pathParams []descriptor.Parameter) error {
	_ = "STUB: not implemented"
	// Sort keys so that when two messages flatten to the same OpenAPI definition
	// name the winner is deterministic (last in sorted order wins) rather than
	// varying with Go's random map iteration order.
	return nil
}

// isVisible checks if a field/RPC is visible based on the visibility restriction
// combined with the `visibility_restriction_selectors`.
// Elements with an overlap on `visibility_restriction_selectors` are visible, those without are not visible.
// Elements without `google.api.VisibilityRule` annotations entirely are always visible.
func isVisible(r *visibility.VisibilityRule, reg *descriptor.Registry) bool {
	_ = "STUB: not implemented"
	return false
}

// No restrictions results in the element always being visible

func shouldExcludeField(name string, excluded []descriptor.Parameter) bool {
	_ = "STUB: not implemented"
	return false
}

func filterOutExcludedFields(fields []string, excluded []descriptor.Parameter) []string {
	_ = "STUB: not implemented"
	return nil
}

// schemaOfFieldBase returns a base Schema Object for a protobuf field.
func schemaOfFieldBase(f *descriptor.Field, reg *descriptor.Registry, refs refMap) openapiSchemaObject {
	_ = "STUB: not implemented"
	return *new(openapiSchemaObject)
}

// Only set core.Type = "object" for MESSAGE types with $ref if the flag is not set.
// When omitArrayItemTypeWhenRefSibling is true, we omit "type: object" to avoid
// no-$ref-siblings violations in OpenAPI v2, since $ref already implies the type is object.

// schemaOfField returns a OpenAPI Schema Object for a protobuf field.
func schemaOfField(f *descriptor.Field, reg *descriptor.Registry, refs refMap) openapiSchemaObject {
	_ = "STUB: not implemented"
	return *new(openapiSchemaObject)
}

// primitiveSchema returns a pair of "Type" and "Format" in JSON Schema for
// the given primitive field type.
// The last return parameter is true iff the field type is actually primitive.
func primitiveSchema(t descriptorpb.FieldDescriptorProto_Type) (ftype, format string, ok bool) {
	_ = "STUB: not implemented"
	return "", "", false
}

// 64bit integer types are marshaled as string in the default JSONPb marshaler.
// TODO(yugui) Add an option to declare 64bit integers as int64.
//
// NOTE: uint64 is not a predefined format of integer type in OpenAPI spec.
// So we cannot expect that uint64 is commonly supported by OpenAPI processor.

// Ditto.

// Ditto.

// NOTE: in OpenAPI specification, format should be empty on boolean type

// NOTE: in OpenAPI specification, can be empty on string type
// see: https://swagger.io/specification/v2/#data-types

// Ditto.

// renderEnumerationsAsDefinition inserts enums into the definitions object.
func renderEnumerationsAsDefinition(enums enumMap, d openapiDefinitionsObject, reg *descriptor.Registry, customRefs refMap) {
	_ = "STUB: not implemented"
	return
}

// it may be necessary to sort the result of the GetValue function.

// Warning: Make sure not to overwrite any fields already set on the schema type.
// This is only a subset of the fields from JsonSchema since most of them only apply to arrays or objects not enums

// Enum comments should go to Description, not Title.
// updateOpenAPIDataFromComments may set Title as a fallback
// when Summary field is not available on schema objects.
// https://github.com/grpc-ecosystem/grpc-gateway/issues/2670

// Take in a FQMN or FQEN and return a OpenAPI safe version of the FQMN and
// a boolean indicating if FQMN was properly resolved.
func fullyQualifiedNameToOpenAPIName(fqn string, reg *descriptor.Registry) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Lookup message type by location.name and return an openapiv2-safe version
// of its FQMN.
func lookupMsgAndOpenAPIName(location, name string, reg *descriptor.Registry) (*descriptor.Message, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// registriesSeen is used to memoise calls to resolveFullyQualifiedNameToOpenAPINames so
// we don't repeat it unnecessarily, since it can take some time.
var (
	registriesSeen      = map[*descriptor.Registry]map[string]string{}
	registriesSeenMutex sync.Mutex
)

// Take the names of every proto message and generate a unique reference for each, according to the given strategy.
func resolveFullyQualifiedNameToOpenAPINames(messages []string, namingStrategy string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

var canRegexp = regexp.MustCompile("{([a-zA-Z][a-zA-Z0-9_.-]*)([^}]*)}")

// templateToParts splits a URL template into path segments for use by `partsToOpenAPIPath` and `partsToRegexpMap`.
//
// Parameters:
//   - path:	The URL template as defined by https://github.com/googleapis/googleapis/blob/master/google/api/http.proto
//   - reg:	The descriptor registry used to read compiler flags
//   - fields:	The fields of the request message, only used when `useJSONNamesForFields` is true
//   - msgs:	The Messages of the service binding, only used when `useJSONNamesForFields` is true
//
// Returns:
//
//	The path segments of the URL template.
func templateToParts(path string, reg *descriptor.Registry, fields []*descriptor.Field, msgs []*descriptor.Message) []string {
	_ = "STUB: not implemented"
	// It seems like the right thing to do here is to just use
	// strings.Split(path, "/") but that breaks badly when you hit a url like
	// /{my_field=prefix/*}/ and end up with 2 sections representing my_field.
	// Instead do the right thing and write a small pushdown (counter) automata
	// for it.
	return nil
}

// Push on the stack

// Pop from the stack

// Since the stack was empty when we hit the '/' we are done with this
// section.

// Only treat this as a verb if we're at the end of the path or
// if there are no more path segments (only more literals after the colon)

// Now append the last element to parts

// processParametersInSegment processes a path segment (like ":verb/{param}") to convert
// parameter names to camelCase while preserving the overall structure
func processParametersInSegment(segment string, fields []*descriptor.Field, msgs []*descriptor.Message) string {
	_ = "STUB: not implemented"
	return ""
}

// partsToOpenAPIPath converts each path part of the form /path/{string_value=strprefix/*} which is defined in
// https://github.com/googleapis/googleapis/blob/master/google/api/http.proto to the OpenAPI expected form /path/{string_value}.
// For example this would replace the path segment of "{foo=bar/*}" with "{foo}" or "prefix{bang=bash/**}" with "prefix{bang}".
// OpenAPI 2 only allows simple path parameters with the constraints on that parameter specified in the OpenAPI
// schema's "pattern" instead of in the path parameter itself.
func partsToOpenAPIPath(parts []string, overrides map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

// Last item is a verb (":" LITERAL).

// partsToRegexpMap returns a map of parameter name to ECMA 262 patterns
// which is what the "pattern" field on an OpenAPI parameter expects.
// See https://swagger.io/specification/v2/ (Parameter Object) and
// https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.2.3.
// The expression is generated based on expressions defined by https://github.com/googleapis/googleapis/blob/master/google/api/http.proto
// "Path Template Syntax" section which allow for a "param_name=foobar/*/bang/**" style expressions inside
// the path parameter placeholders that indicate constraints on the values of those parameters.
// This function will scan the split parts of a path template for parameters and
// outputs a map of the name of the parameter to a ECMA regular expression.  See the http.proto file for descriptions
// of the supported syntax. This function will ignore any path parameters that don't contain a "=" after the
// parameter name.  For supported parameters, we assume "*" represent all characters except "/" as it's
// intended to match a single path element and we assume "**" matches any character as it's intended to match multiple
// path elements.
// For example "{name=organizations/*/roles/*}" would produce the regular expression for the "name" parameter of
// "organizations/[^/]+/roles/[^/]+" or "{bar=bing/*/bang/**}" would produce the regular expression for the "bar"
// parameter of "bing/[^/]+/bang/.+".
//
// Note that OpenAPI does not actually support path parameters with "/", see https://github.com/OAI/OpenAPI-Specification/issues/892
func partsToRegexpMap(parts []string) map[string]string { _ = "STUB: not implemented"; return nil }

// this part matches the standard and should be made into a regular expression
// assume the string's characters other than "**" and "*" are literals (not necessarily a good assumption 100% of the times, but it will support most use cases)

// ** implies any character including "/"
// * implies any character except "/"

func renderServiceTags(services []*descriptor.Service, reg *descriptor.Registry) []openapiTagObject {
	_ = "STUB: not implemented"
	return nil
}

// If no description is set from options, use proto comments

// findServiceIndex finds the index of a service within its file's service list.
func findServiceIndex(svc *descriptor.Service) int { _ = "STUB: not implemented"; return 0 }

// expandPathPatterns searches the URI parts for path parameters with pattern and when the pattern contains a sub-path,
// it expands the pattern into the URI parts and adds the new path parameters to the pathParams slice.
//
// Parameters:
//   - pathParts:	the URI parts parsed from the path template with `templateToParts` function
//   - pathParams: the path parameters of the service binding
//
// Returns:
//
//	The modified pathParts and pathParams slice.
func expandPathPatterns(pathParts []string, pathParams []descriptor.Parameter, reg *descriptor.Registry) ([]string, []descriptor.Parameter) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the new parameter from the pattern replaces the old path parameter

func renderServices(services []*descriptor.Service, paths *openapiPathsObject, reg *descriptor.Registry, requestResponseRefs, customRefs refMap, msgs []*descriptor.Message, defs openapiDefinitionsObject) error {
	_ = "STUB: not implemented"
	// Correctness of svcIdx and methIdx depends on 'services' containing the services in the same order as the 'file.Service' array.
	return nil
}

// Iterate over all the OpenAPI parameters

// split the path template into its parts

// extract any constraints specified in the path placeholders into ECMA regular expressions

// Keep track of path parameter overrides

// If there is no mandatory format based on the field,
// allow it to be overridden by the user

// verify if the parameter is deprecated, either via proto or annotation

// Parameters in gRPC-Gateway can only be strings?

// Now check if there is a body parameter

// Recursively render fields as definitions as long as they contain path parameters.
// Special case for top level body if we don't have a body field.

// No field for body, use type.

// Special workaround for Empty: it's well-known type but wknSchemas only returns schema.schemaCore; but we need to set schema.Properties which is a level higher.

// Body field path is limited to one path component. From google.api.HttpRule.body:
// "NOTE: the referred field must be present at the top-level of the request message type."
// Ref: https://github.com/googleapis/googleapis/blob/b3397f5febbf21dfc69b875ddabaf76bee765058/google/api/http.proto#L350-L352

// Align pathParams with body field path.

// When there are no path parameters, we only need the base schema of the field.
// https://github.com/grpc-ecosystem/grpc-gateway/issues/3058

// renderFieldAsDefinition may add the body field name to the schema's required array
// via updateSwaggerObjectFromFieldBehavior. However, for body parameters, the schema
// represents the field's type, not the containing message. The body field name should
// only be in the schema's required array if it's actually a property of the schema.
// Remove the body field name from required if it's not a property to avoid invalid entries.

// Build a set of property names

// Filter required array: keep field names that are either:
// 1. Not the body field name, OR
// 2. The body field name AND it's actually a property

// It's a property, keep it (but only once)

// else: It's not a property, skip it

// Not the body field name, keep it

// add the parameters to the query string

// handle case where we have an existing mapping for the same path and method

// Without a path parameter, there is nothing to vary to support multiple mappings of the same path/method.
// Previously this did not log an error and only overwrote the mapping, we now log the error but
// still overwrite the mapping

// Iterate until there is not an existing operation that matches the same escaped path.
// Most of the time this will only be a single iteration, but a large API could technically have
// a pretty large amount of these if it used similar patterns for all its functions.

// update the pathItemObject we are adding to with the new path

// Don't link to a full definition for
// empty; it's overly verbose.
// schema.Properties{} renders it as
// well, without a definition

// Special workaround for Empty: it's well-known type but wknSchemas only returns schema.schemaCore; but we need to set schema.Properties which is a level higher.

// This is resolving the value of response_body in the google.api.HttpRule

// Special case HttpBody responses, they will be unformatted bytes

// The error response is still JSON, but technically the full response
// is still unformatted, so don't include the error response structure.

// https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#responses-object

// OperationID must be unique in an OpenAPI v2 definition.

// Fill reference map with referenced request messages

// Set Tag with the user-defined service name

// Merge response data into default response if available.

// TODO(ivucica): add remaining fields of operation object

// Success! return nil on the error object

// Returns the openapiPathItemObject associated with a path. If path is not present, returns
// empty openapiPathItemObject and false.
func getPathItemObject(paths openapiPathsObject, path string) (openapiPathItemObject, bool) {
	_ = "STUB: not implemented"
	return *new(openapiPathItemObject), false
}

// If a path already exists in openapiPathsObject, updates that path's openapiPathItemObject. If not,
// appends a new path and openapiPathItemObject to the openapiPathsObject.
func updatePaths(paths *openapiPathsObject, path string, pathItemObject openapiPathItemObject) {
	_ = "STUB: not implemented"
	return
}

func mergeDescription(schema openapiSchemaObject) string { _ = "STUB: not implemented"; return "" }

// join title because title of parameter object will be ignored

func operationForMethod(httpMethod string) func(*openapiPathItemObject) *openapiOperationObject {
	_ = "STUB: not implemented"
	return nil
}

// This function is called with a param which contains the entire definition of a method.
func applyTemplate(p param) (*openapiSwaggerObject, error) {
	_ = "STUB: not implemented"
	// Create the basic template object. This is the object that everything is
	// defined off of.
	return nil, nil
}

// OpenAPI 2.0 is the version of this document

// IMPORTANT: Initialize the naming cache BEFORE any code that uses fullyQualifiedNameToOpenAPIName.
// This ensures consistent naming between renderServices (which generates $refs) and
// renderMessagesAsDefinition (which generates definitions).
//
// Pre-scan to collect referenced names WITHOUT using the naming cache.
// This allows us to build the cache with the correct filtered names upfront.

// Get all names from the registry

// Filter: EXCLUDE names that are from a DIFFERENT package AND are NOT referenced
// This way we keep all names from the current package, and all referenced names from other packages

// Include if: (1) from current package, OR (2) actually referenced, OR (3) from google.*/grpc.* packages

// Initialize the naming cache BEFORE renderServices so all lookups use consistent naming

// Loops through all the services and their exposed GET/POST/PUT/DELETE definitions
// and create entries for all of them.
// Also adds custom user specified references to second map.
// NOTE: This now uses the naming cache initialized above.

// Add the error type to the message map

// just in case there is an error looking up runtimeError

// Find all the service's messages and enumerations that are defined (recursively)
// and write request, response and other custom (but referenced) types out as definition objects.
// NOTE: This uses the same naming cache that was used by renderServices above.

// File itself might have some comments and metadata.

// There may be additional options in the OpenAPI option in the proto.

// Populate all Paths with Responses set at top level,
// preferring Responses already set over those at the top level.

// Don't overwrite already existing Responses

// Additional fields on the OpenAPI v2 spec's "OpenAPI" object
// should be added here, once supported in the proto.

// Finally add any references added by users that aren't
// otherwise rendered.

func mergeTags(existingTags []openapiTagObject, tags []openapiTagObject) []openapiTagObject {
	_ = "STUB: not implemented"
	return nil
}

func processExtensions(inputExts map[string]*structpb.Value) ([]extension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateHeaderTypeAndFormat(headerType, format string) error {
	_ = "STUB: not implemented"
	// The type of the object. The value MUST be one of "string", "number", "integer", "boolean", or "array"
	// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#headerObject
	// Note: currently not implementing array as we are only implementing this in the operation response context
	return nil
}

// the format property is an open string-valued property, and can have any value to support documentation needs
// primary check for format is to ensure that the number/integer formats are extensions of the specified type
// See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#dataTypeFormat

func validateDefaultValueTypeAndFormat(headerType string, defaultValue string, format string) error {
	_ = "STUB: not implemented"
	return nil
}

func isQuotedString(s string) bool { _ = "STUB: not implemented"; return false }

func isJSONNumber(s string, t string) error { _ = "STUB: not implemented"; return nil }

// Floating point values that cannot be represented as sequences of digits (such as Infinity and NaN) are not permitted.
// See: https://tools.ietf.org/html/rfc4627#section-2.4

func isBool(s string) bool {
	_ = "STUB: not implemented"
	// Unable to use strconv.ParseBool because it returns truthy values https://golang.org/pkg/strconv/#example_ParseBool
	// per https://swagger.io/specification/v2/#data-types
	// type: boolean represents two values: true and false. Note that truthy and falsy values such as "true", "", 0 or null are not considered boolean values.
	return false
}

func processHeaders(inputHdrs map[string]*openapi_options.Header) (openapiHeadersObject, error) {
	_ = "STUB: not implemented"
	return *new(openapiHeadersObject), nil
}

func removeInternalComments(comment string) string { _ = "STUB: not implemented"; return "" }

// Trim only one line prior to all spaces

// updateOpenAPIDataFromComments updates a OpenAPI object based on a comment
// from the proto file.
//
// First paragraph of a comment is used for summary. Remaining paragraphs of
// a comment are used for description. If 'Summary' field is not present on
// the passed swaggerObject, the summary and description are joined by \n\n.
//
// If there is a field named 'Info', its 'Summary' and 'Description' fields
// will be updated instead.
//
// If there is no 'Summary', the same behavior will be attempted on 'Title',
// but only if the last character is not a period.
func updateOpenAPIDataFromComments(reg *descriptor.Registry, swaggerObject interface{}, data interface{}, comment string, isPackageObject bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Checks whether the "ignore_comments" flag is set to true

// Checks whether the "remove_internal_comments" flag is set to true

// Checks whether the "use_go_templates" flag is set to true

// Figure out what to apply changes to.

// No such field? Apply summary and description directly to
// passed object.

// Figure out which properties to update.

// If there is a summary (or summary-equivalent) and it's empty, use the first
// paragraph as summary, and the rest as description.

// overrides the schema value only if it's empty
// keep the comment precedence when updating the package definition

// overrides the schema value only if it's empty
// keep the comment precedence when updating the package definition

// There was no summary field on the swaggerObject. Try to apply the
// whole comment into description if the OpenAPI object description is empty.

func fieldProtoComments(reg *descriptor.Registry, msg *descriptor.Message, field *descriptor.Field) string {
	_ = "STUB: not implemented"
	return ""
}

func enumValueProtoComments(reg *descriptor.Registry, enum *descriptor.Enum) string {
	_ = "STUB: not implemented"
	return ""
}

func protoComments(reg *descriptor.Registry, file *descriptor.File, outers []string, typeName string, typeIndex int32, fieldPaths ...int32) string {
	_ = "STUB: not implemented"
	return ""
}

// TODO(ivucica): this is a hack to fix "// " being interpreted as "//".
// perhaps we should:
// - split by \n
// - determine if every (but first and last) line begins with " "
// - trim every line only if that is the case
// - join by \n

func goTemplateComments(comment string, data interface{}, reg *descriptor.Registry) string {
	_ = "STUB: not implemented"
	return ""
}

// Allows importing documentation from a file

// Runs template over imported file

// Grabs title and description from a field

// If there is an error parsing the templating insert the error as string in the comment
// to make it easier to debug the template error

// If there is an error executing the templating insert the error as string in the comment
// to make it easier to debug the error

var (
	messageProtoPath = protoPathIndex(reflect.TypeOf((*descriptorpb.FileDescriptorProto)(nil)), "MessageType")
	nestedProtoPath  = protoPathIndex(reflect.TypeOf((*descriptorpb.DescriptorProto)(nil)), "NestedType")
	packageProtoPath = protoPathIndex(reflect.TypeOf((*descriptorpb.FileDescriptorProto)(nil)), "Package")
	serviceProtoPath = protoPathIndex(reflect.TypeOf((*descriptorpb.FileDescriptorProto)(nil)), "Service")
	methodProtoPath  = protoPathIndex(reflect.TypeOf((*descriptorpb.ServiceDescriptorProto)(nil)), "Method")
)

func isProtoPathMatches(paths []int32, outerPaths []int32, typeName string, typeIndex int32, fieldPaths []int32) bool {
	_ = "STUB: not implemented"
	return false
}

// path for package comments is just [2], and all the other processing
// is too complex for it.

// protoPathIndex returns a path component for google.protobuf.descriptor.SourceCode_Location.
//
// Specifically, it returns an id as generated from descriptor proto which
// can be used to determine what type the id following it in the path is.
// For example, if we are trying to locate comments related to a field named
// `Address` in a message named `Person`, the path will be:
//
//	[4, a, 2, b]
//
// While `a` gets determined by the order in which the messages appear in
// the proto file, and `b` is the field index specified in the proto
// file itself, the path actually needs to specify that `a` refers to a
// message and not, say, a service; and  that `b` refers to a field and not
// an option.
//
// protoPathIndex figures out the values 4 and 2 in the above example. Because
// messages are top level objects, the value of 4 comes from field id for
// `MessageType` inside `google.protobuf.descriptor.FileDescriptor` message.
// This field has a message type `google.protobuf.descriptor.DescriptorProto`.
// And inside message `DescriptorProto`, there is a field named `Field` with id
// 2.
//
// Some code generators seem to be hardcoding these values; this method instead
// interprets them from `descriptor.proto`-derived Go source as necessary.
func protoPathIndex(descriptorType reflect.Type, what string) int32 {
	_ = "STUB: not implemented"
	return 0
}

// extractOperationOptionFromMethodDescriptor extracts the message of type
// openapi_options.Operation from a given proto method's descriptor.
func extractOperationOptionFromMethodDescriptor(meth *descriptorpb.MethodDescriptorProto) (*openapi_options.Operation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extractSchemaOptionFromMessageDescriptor extracts the message of type
// openapi_options.Schema from a given proto message's descriptor.
func extractSchemaOptionFromMessageDescriptor(msg *descriptorpb.DescriptorProto) (*openapi_options.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extractEnumSchemaOptionFromEnumDescriptor extracts the message of type
// openapi_options.EnumSchema from a given proto enum's descriptor.
func extractEnumSchemaOptionFromEnumDescriptor(enum *descriptorpb.EnumDescriptorProto) (*openapi_options.EnumSchema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extractTagOptionFromServiceDescriptor extracts the tag of type
// openapi_options.Tag from a given proto service's descriptor.
func extractTagOptionFromServiceDescriptor(svc *descriptorpb.ServiceDescriptorProto) (*openapi_options.Tag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// extractOpenAPIOptionFromFileDescriptor extracts the message of type
// openapi_options.OpenAPI from a given proto method's descriptor.
func extractOpenAPIOptionFromFileDescriptor(file *descriptorpb.FileDescriptorProto) (*openapi_options.Swagger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractJSONSchemaFromFieldDescriptor(fd *descriptorpb.FieldDescriptorProto) (*openapi_options.JSONSchema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extractFieldBehaviorFromFieldDescriptor(fd *descriptorpb.FieldDescriptorProto) ([]annotations.FieldBehavior, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFieldVisibilityOption(fd *descriptor.Field) *visibility.VisibilityRule {
	_ = "STUB: not implemented"
	return nil
}

func getServiceVisibilityOption(fd *descriptor.Service) *visibility.VisibilityRule {
	_ = "STUB: not implemented"
	return nil
}

func getMethodVisibilityOption(fd *descriptor.Method) *visibility.VisibilityRule {
	_ = "STUB: not implemented"
	return nil
}

func getEnumValueVisibilityOption(fd *descriptorpb.EnumValueDescriptorProto) *visibility.VisibilityRule {
	_ = "STUB: not implemented"
	return nil
}

func getMethodOpenAPIOption(reg *descriptor.Registry, meth *descriptor.Method) (*openapi_options.Operation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getMessageOpenAPIOption(reg *descriptor.Registry, msg *descriptor.Message) (*openapi_options.Schema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getEnumOpenAPIOption(reg *descriptor.Registry, enum *descriptor.Enum) (*openapi_options.EnumSchema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getServiceOpenAPIOption(reg *descriptor.Registry, svc *descriptor.Service) (*openapi_options.Tag, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFileOpenAPIOption(reg *descriptor.Registry, file *descriptor.File) (*openapi_options.Swagger, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFieldOpenAPIOption(reg *descriptor.Registry, fd *descriptor.Field) (*openapi_options.JSONSchema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFieldBehaviorOption(reg *descriptor.Registry, fd *descriptor.Field) ([]annotations.FieldBehavior, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func protoJSONSchemaToOpenAPISchemaCore(j *openapi_options.JSONSchema, reg *descriptor.Registry, refs refMap) schemaCore {
	_ = "STUB: not implemented"
	return *new(schemaCore)
}

func updateswaggerObjectFromJSONSchema(s *openapiSchemaObject, j *openapi_options.JSONSchema, reg *descriptor.Registry, data interface{}) {
	_ = "STUB: not implemented"
	return
}

func updateSwaggerObjectFromFieldBehavior(s *openapiSchemaObject, j []annotations.FieldBehavior, reg *descriptor.Registry, field *descriptor.Field) {
	_ = "STUB: not implemented"
	return
}

// OpenAPI v3 supports a writeOnly property, but this is not supported in Open API v2

func openapiSchemaFromProtoEnumSchema(s *openapi_options.EnumSchema, reg *descriptor.Registry, refs refMap, data interface{}) openapiSchemaObject {
	_ = "STUB: not implemented"
	return *new(openapiSchemaObject)
}

func openapiSchemaFromProtoSchema(s *openapi_options.Schema, reg *descriptor.Registry, refs refMap, data interface{}) openapiSchemaObject {
	_ = "STUB: not implemented"
	return *new(openapiSchemaObject)
}

func openapiExamplesFromProtoExamples(in map[string]string) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// JSON example objects are rendered raw.

// All other mimetype examples are rendered as strings.

func protoJSONSchemaTypeToFormat(in []openapi_options.JSONSchema_JSONSchemaSimpleTypes) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// Can't support more than 1 type, just return the first element.
// This is due to an inconsistency in the design of the openapiv2 proto
// and that used in schemaCore. schemaCore uses the v3 definition of types,
// which only allows a single string, while the openapiv2 proto uses the OpenAPI v2
// definition, which defers to the JSON schema definition, which allows a string or an array.
// Sources:
// https://swagger.io/specification/#itemsObject
// https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.5.2

// NOTE: in OpenAPI specification, format should be empty on boolean type

// NOTE: in OpenAPI specification, format should be empty on string type

// Maybe panic?

func protoExternalDocumentationToOpenAPIExternalDocumentation(in *openapi_options.ExternalDocumentation, reg *descriptor.Registry, data interface{}) *openapiExternalDocumentationObject {
	_ = "STUB: not implemented"
	return nil
}

func addCustomRefs(d openapiDefinitionsObject, reg *descriptor.Registry, refs refMap) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip already existing definitions

// ?? Should be either enum or msg

// Run again in case any new refs were added

func lowerCamelCase(fieldName string, fields []*descriptor.Field, msgs []*descriptor.Message) string {
	_ = "STUB: not implemented"
	return ""
}

func getReservedJSONName(fieldName string, messageNameToFieldsToJSONName map[string]map[string]string, fieldNameToType map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

func find(a []string, x string) int {
	_ = "STUB: not implemented"
	// This is a linear search but we are dealing with a small number of fields
	return 0
}

// Make a deep copy of the outer parameters that has paramName as the first component,
// but remove the first component of the field path.
func subPathParams(paramName string, outerParams []descriptor.Parameter) []descriptor.Parameter {
	_ = "STUB: not implemented"
	return nil
}

func getFieldConfiguration(reg *descriptor.Registry, fd *descriptor.Field) *openapi_options.JSONSchema_FieldConfiguration {
	_ = "STUB: not implemented"
	return nil
}
