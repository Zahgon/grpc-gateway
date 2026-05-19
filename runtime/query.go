package runtime

import (
	"net/url"
	"regexp"

	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var valuesKeyRegexp = regexp.MustCompile(`^(.*)\[(.*)\]$`)

var currentQueryParser QueryParameterParser = &DefaultQueryParser{}

// QueryParameterParser defines interface for all query parameter parsers
type QueryParameterParser interface {
	Parse(msg proto.Message, values url.Values, filter *utilities.DoubleArray) error
}

// PopulateQueryParameters parses query parameters
// into "msg" using current query parser
func PopulateQueryParameters(msg proto.Message, values url.Values, filter *utilities.DoubleArray) error {
	_ = "STUB: not implemented"
	return nil
}

// DefaultQueryParser is a QueryParameterParser which implements the default
// query parameters parsing behavior.
//
// See https://github.com/grpc-ecosystem/grpc-gateway/issues/2632 for more context.
type DefaultQueryParser struct{}

// Parse populates "values" into "msg".
// A value is ignored if its key starts with one of the elements in "filter".
func (*DefaultQueryParser) Parse(msg proto.Message, values url.Values, filter *utilities.DoubleArray) error {
	_ = "STUB: not implemented"
	return nil
}

// PopulateFieldFromPath sets a value in a nested Protobuf structure.
func PopulateFieldFromPath(msg proto.Message, fieldPathString string, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func normalizeFieldPath(msgValue protoreflect.Message, fieldPath []string) []string {
	_ = "STUB: not implemented"
	return nil
}

// return initial field path values if no matching  message field was found

// If this is the last element, we're done

// Only singular message fields are allowed

// Get the nested message

func populateFieldValueFromPath(msgValue protoreflect.Message, fieldPath []string, values []string) error {
	_ = "STUB: not implemented"
	return nil
}

// Get field by name

// We're not returning an error here because this could just be
// an extra query parameter that isn't part of the request.

// Check if oneof already set

// If this is the last element, we're done

// Only singular message fields are allowed

// Get the nested message

func populateField(fieldDescriptor protoreflect.FieldDescriptor, msgValue protoreflect.Message, value string) error {
	_ = "STUB: not implemented"
	return nil
}

func populateRepeatedField(fieldDescriptor protoreflect.FieldDescriptor, list protoreflect.List, values []string) error {
	_ = "STUB: not implemented"
	return nil
}

func populateMapField(fieldDescriptor protoreflect.FieldDescriptor, mp protoreflect.Map, values []string) error {
	_ = "STUB: not implemented"
	return nil
}

func parseField(fieldDescriptor protoreflect.FieldDescriptor, value string) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}

// Look for enum by name

// Look for enum by number

func parseMessage(msgDescriptor protoreflect.MessageDescriptor, value string) (protoreflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(protoreflect.Value), nil
}
