package descriptor

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/httprule"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

// IsWellKnownType returns true if the provided fully qualified type name is considered 'well-known'.
func IsWellKnownType(typeName string) bool { _ = "STUB: not implemented"; return false }

// GoPackage represents a golang package.
type GoPackage struct {
	// Path is the package path to the package.
	Path string
	// Name is the package name of the package
	Name string
	// Alias is an alias of the package unique within the current invocation of gRPC-Gateway generator.
	Alias string
}

// Standard returns whether the import is a golang standard package.
func (p GoPackage) Standard() bool { _ = "STUB: not implemented"; return false }

// String returns a string representation of this package in the form of import line in golang.
func (p GoPackage) String() string { _ = "STUB: not implemented"; return "" }

// ResponseFile wraps pluginpb.CodeGeneratorResponse_File.
type ResponseFile struct {
	*pluginpb.CodeGeneratorResponse_File
	// GoPkg is the Go package of the generated file.
	GoPkg GoPackage
}

// File wraps descriptorpb.FileDescriptorProto for richer features.
type File struct {
	*descriptorpb.FileDescriptorProto
	// GoPkg is the go package of the go file generated from this file.
	GoPkg GoPackage
	// GeneratedFilenamePrefix is used to construct filenames for generated
	// files associated with this source file.
	//
	// For example, the source file "dir/foo.proto" might have a filename prefix
	// of "dir/foo". Appending ".pb.go" produces an output file of "dir/foo.pb.go".
	GeneratedFilenamePrefix string
	// Messages is the list of messages defined in this file.
	Messages []*Message
	// Enums is the list of enums defined in this file.
	Enums []*Enum
	// Services is the list of services defined in this file.
	Services []*Service
}

// Pkg returns package name or alias if it's present
func (f *File) Pkg() string { _ = "STUB: not implemented"; return "" }

// proto2 determines if the syntax of the file is proto2.
func (f *File) proto2() bool { _ = "STUB: not implemented"; return false }

// Message describes a protocol buffer message types.
type Message struct {
	*descriptorpb.DescriptorProto
	// File is the file where the message is defined.
	File *File
	// Outers is a list of outer messages if this message is a nested type.
	Outers []string
	// Fields is a list of message fields.
	Fields []*Field
	// Index is proto path index of this message in File.
	Index int
	// ForcePrefixedName when set to true, prefixes a type with a package prefix.
	ForcePrefixedName bool
}

// FQMN returns a fully qualified message name of this message.
func (m *Message) FQMN() string { _ = "STUB: not implemented"; return "" }

// GoType returns a go type name for the message type.
// It prefixes the type name with the package alias if
// its belonging package is not "currentPackage".
func (m *Message) GoType(currentPackage string) string { _ = "STUB: not implemented"; return "" }

// Enum describes a protocol buffer enum types.
type Enum struct {
	*descriptorpb.EnumDescriptorProto
	// File is the file where the enum is defined
	File *File
	// Outers is a list of outer messages if this enum is a nested type.
	Outers []string
	// Index is a enum index value.
	Index int
	// ForcePrefixedName when set to true, prefixes a type with a package prefix.
	ForcePrefixedName bool
}

// FQEN returns a fully qualified enum name of this enum.
func (e *Enum) FQEN() string { _ = "STUB: not implemented"; return "" }

// GoType returns a go type name for the enum type.
// It prefixes the type name with the package alias if
// its belonging package is not "currentPackage".
func (e *Enum) GoType(currentPackage string) string { _ = "STUB: not implemented"; return "" }

func goTypeName(outers []string, name string) string { _ = "STUB: not implemented"; return "" }

// Service wraps descriptorpb.ServiceDescriptorProto for richer features.
type Service struct {
	*descriptorpb.ServiceDescriptorProto
	// File is the file where this service is defined.
	File *File
	// Methods is the list of methods defined in this service.
	Methods []*Method
	// ForcePrefixedName when set to true, prefixes a type with a package prefix.
	ForcePrefixedName bool
}

// FQSN returns the fully qualified service name of this service.
func (s *Service) FQSN() string { _ = "STUB: not implemented"; return "" }

// InstanceName returns object name of the service with package prefix if needed
func (s *Service) InstanceName() string { _ = "STUB: not implemented"; return "" }

// ClientConstructorName returns name of the Client constructor with package prefix if needed
func (s *Service) ClientConstructorName() string { _ = "STUB: not implemented"; return "" }

// Method wraps descriptorpb.MethodDescriptorProto for richer features.
type Method struct {
	*descriptorpb.MethodDescriptorProto
	// Service is the service which this method belongs to.
	Service *Service
	// RequestType is the message type of requests to this method.
	RequestType *Message
	// ResponseType is the message type of responses from this method.
	ResponseType *Message
	Bindings     []*Binding
}

// FQMN returns a fully qualified rpc method name of this method.
func (m *Method) FQMN() string { _ = "STUB: not implemented"; return "" }

// Binding describes how an HTTP endpoint is bound to a gRPC method.
type Binding struct {
	// Method is the method which the endpoint is bound to.
	Method *Method
	// Index is a zero-origin index of the binding in the target method
	Index int
	// PathTmpl is path template where this method is mapped to.
	PathTmpl httprule.Template
	// HTTPMethod is the HTTP method which this method is mapped to.
	HTTPMethod string
	// PathParams is the list of parameters provided in HTTP request paths.
	PathParams []Parameter
	// Body describes parameters provided in HTTP request body.
	Body *Body
	// ResponseBody describes field in response struct to marshal in HTTP response body.
	ResponseBody *Body
}

// ExplicitParams returns a list of explicitly bound parameters of "b",
// i.e. a union of field path for body and field paths for path parameters.
func (b *Binding) ExplicitParams() []string { _ = "STUB: not implemented"; return nil }

// Field wraps descriptorpb.FieldDescriptorProto for richer features.
type Field struct {
	*descriptorpb.FieldDescriptorProto
	// Message is the message type which this field belongs to.
	Message *Message
	// FieldMessage is the message type of the field.
	FieldMessage *Message
	// ForcePrefixedName when set to true, prefixes a type with a package prefix.
	ForcePrefixedName bool
}

// FQFN returns a fully qualified field name of this field.
func (f *Field) FQFN() string { _ = "STUB: not implemented"; return "" }

// Parameter is a parameter provided in http requests
type Parameter struct {
	// FieldPath is a path to a proto field which this parameter is mapped to.
	FieldPath
	// Target is the proto field which this parameter is mapped to.
	Target *Field
	// Method is the method which this parameter is used for.
	Method *Method
}

// ConvertFuncExpr returns a go expression of a converter function.
// The converter function converts a string into a value for the parameter.
func (p Parameter) ConvertFuncExpr() (string, error) { _ = "STUB: not implemented"; return "", nil }

// IsEnum returns true if the field is an enum type, otherwise false is returned.
func (p Parameter) IsEnum() bool { _ = "STUB: not implemented"; return false }

// IsRepeated returns true if the field is repeated, otherwise false is returned.
func (p Parameter) IsRepeated() bool { _ = "STUB: not implemented"; return false }

// IsProto2 returns true if the field is proto2, otherwise false is returned.
func (p Parameter) IsProto2() bool { _ = "STUB: not implemented"; return false }

// Body describes a http (request|response) body to be sent to the (method|client).
// This is used in body and response_body options in google.api.HttpRule
type Body struct {
	// FieldPath is a path to a proto field which the (request|response) body is mapped to.
	// The (request|response) body is mapped to the (request|response) type itself if FieldPath is empty.
	FieldPath FieldPath
}

// AssignableExpr returns an assignable expression in Go to be used to initialize method request object.
// It starts with "msgExpr", which is the go expression of the method request object.
func (b Body) AssignableExpr(msgExpr string, currentPackage string) string {
	_ = "STUB: not implemented"
	return ""
}

// AssignableExprPrep returns preparatory statements for an assignable expression to initialize the
// method request object.
func (b Body) AssignableExprPrep(msgExpr string, currentPackage string) string {
	_ = "STUB: not implemented"
	return ""
}

// FieldPath is a path to a field from a request message.
type FieldPath []FieldPathComponent

// String returns a string representation of the field path.
func (p FieldPath) String() string { _ = "STUB: not implemented"; return "" }

// IsNestedProto3 indicates whether the FieldPath is a nested Proto3 path.
func (p FieldPath) IsNestedProto3() bool { _ = "STUB: not implemented"; return false }

// IsOptionalProto3 indicates whether the FieldPath is a proto3 optional field.
func (p FieldPath) IsOptionalProto3() bool { _ = "STUB: not implemented"; return false }

// AssignableExpr is an assignable expression in Go to be used to assign a value to the target field.
// It starts with "msgExpr", which is the go expression of the method request object. Before using
// such an expression the prep statements must be emitted first, in case the field path includes
// a oneof. See FieldPath.AssignableExprPrep.
func (p FieldPath) AssignableExpr(msgExpr string, currentPackage string) string {
	_ = "STUB: not implemented"
	return ""
}

// We need to check if the target is not proto3_optional first.
// Under the hood, proto3_optional uses oneof to signal to old proto3 clients
// that presence is tracked for this field. This oneof is known as a "synthetic" oneof.

// AssignableExprPrep returns preparation statements for an assignable expression to assign a value
// to the target field. The Go expression of the method request object is "msgExpr". This is only
// needed for field paths that contain oneofs. Otherwise, an empty string is returned.
func (p FieldPath) AssignableExprPrep(msgExpr string, currentPackage string) string {
	_ = "STUB: not implemented"
	return ""
}

// We need to check if the target is not proto3_optional first.
// Under the hood, proto3_optional uses oneof to signal to old proto3 clients
// that presence is tracked for this field. This oneof is known as a "synthetic" oneof.

// OpaqueSetterExpr returns the Go expression to invoke the generated setter for
// the final component in the path while respecting nested getters required by
// the opaque API.
func (p FieldPath) OpaqueSetterExpr(msgExpr string) string { _ = "STUB: not implemented"; return "" }

// opaqueOwnerExpr builds the Go expression for the message that owns the final
// component in the path by chaining the generated getters.
func (p FieldPath) opaqueOwnerExpr(msgExpr string) string { _ = "STUB: not implemented"; return "" }

// FieldPathComponent is a path component in FieldPath
type FieldPathComponent struct {
	// Name is a name of the proto field which this component corresponds to.
	// TODO(yugui) is this necessary?
	Name string
	// Target is the proto field which this component corresponds to.
	Target *Field
}

// AssignableExpr returns an assignable expression in go for this field.
func (c FieldPathComponent) AssignableExpr() string { _ = "STUB: not implemented"; return "" }

// ValueExpr returns an expression in go for this field.
func (c FieldPathComponent) ValueExpr() string { _ = "STUB: not implemented"; return "" }

var (
	proto3ConvertFuncs = map[descriptorpb.FieldDescriptorProto_Type]string{
		descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:  "runtime.Float64",
		descriptorpb.FieldDescriptorProto_TYPE_FLOAT:   "runtime.Float32",
		descriptorpb.FieldDescriptorProto_TYPE_INT64:   "runtime.Int64",
		descriptorpb.FieldDescriptorProto_TYPE_UINT64:  "runtime.Uint64",
		descriptorpb.FieldDescriptorProto_TYPE_INT32:   "runtime.Int32",
		descriptorpb.FieldDescriptorProto_TYPE_FIXED64: "runtime.Uint64",
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32: "runtime.Uint32",
		descriptorpb.FieldDescriptorProto_TYPE_BOOL:    "runtime.Bool",
		descriptorpb.FieldDescriptorProto_TYPE_STRING:  "runtime.String",
		// FieldDescriptorProto_TYPE_GROUP
		// FieldDescriptorProto_TYPE_MESSAGE
		descriptorpb.FieldDescriptorProto_TYPE_BYTES:    "runtime.Bytes",
		descriptorpb.FieldDescriptorProto_TYPE_UINT32:   "runtime.Uint32",
		descriptorpb.FieldDescriptorProto_TYPE_ENUM:     "runtime.Enum",
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32: "runtime.Int32",
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64: "runtime.Int64",
		descriptorpb.FieldDescriptorProto_TYPE_SINT32:   "runtime.Int32",
		descriptorpb.FieldDescriptorProto_TYPE_SINT64:   "runtime.Int64",
	}

	proto3OptionalConvertFuncs = func() map[descriptorpb.FieldDescriptorProto_Type]string {
		result := make(map[descriptorpb.FieldDescriptorProto_Type]string)
		for typ, converter := range proto3ConvertFuncs {
			// TODO: this will use convert functions from proto2.
			//       The converters returning pointers should be moved
			//       to a more generic file.
			result[typ] = converter + "P"
		}
		return result
	}()

	// TODO: replace it with a IIFE
	proto3RepeatedConvertFuncs = map[descriptorpb.FieldDescriptorProto_Type]string{
		descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:  "runtime.Float64Slice",
		descriptorpb.FieldDescriptorProto_TYPE_FLOAT:   "runtime.Float32Slice",
		descriptorpb.FieldDescriptorProto_TYPE_INT64:   "runtime.Int64Slice",
		descriptorpb.FieldDescriptorProto_TYPE_UINT64:  "runtime.Uint64Slice",
		descriptorpb.FieldDescriptorProto_TYPE_INT32:   "runtime.Int32Slice",
		descriptorpb.FieldDescriptorProto_TYPE_FIXED64: "runtime.Uint64Slice",
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32: "runtime.Uint32Slice",
		descriptorpb.FieldDescriptorProto_TYPE_BOOL:    "runtime.BoolSlice",
		descriptorpb.FieldDescriptorProto_TYPE_STRING:  "runtime.StringSlice",
		// FieldDescriptorProto_TYPE_GROUP
		// FieldDescriptorProto_TYPE_MESSAGE
		descriptorpb.FieldDescriptorProto_TYPE_BYTES:    "runtime.BytesSlice",
		descriptorpb.FieldDescriptorProto_TYPE_UINT32:   "runtime.Uint32Slice",
		descriptorpb.FieldDescriptorProto_TYPE_ENUM:     "runtime.EnumSlice",
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32: "runtime.Int32Slice",
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64: "runtime.Int64Slice",
		descriptorpb.FieldDescriptorProto_TYPE_SINT32:   "runtime.Int32Slice",
		descriptorpb.FieldDescriptorProto_TYPE_SINT64:   "runtime.Int64Slice",
	}

	proto2ConvertFuncs = map[descriptorpb.FieldDescriptorProto_Type]string{
		descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:  "runtime.Float64P",
		descriptorpb.FieldDescriptorProto_TYPE_FLOAT:   "runtime.Float32P",
		descriptorpb.FieldDescriptorProto_TYPE_INT64:   "runtime.Int64P",
		descriptorpb.FieldDescriptorProto_TYPE_UINT64:  "runtime.Uint64P",
		descriptorpb.FieldDescriptorProto_TYPE_INT32:   "runtime.Int32P",
		descriptorpb.FieldDescriptorProto_TYPE_FIXED64: "runtime.Uint64P",
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32: "runtime.Uint32P",
		descriptorpb.FieldDescriptorProto_TYPE_BOOL:    "runtime.BoolP",
		descriptorpb.FieldDescriptorProto_TYPE_STRING:  "runtime.StringP",
		// FieldDescriptorProto_TYPE_GROUP
		// FieldDescriptorProto_TYPE_MESSAGE
		// FieldDescriptorProto_TYPE_BYTES
		// TODO(yugui) Handle bytes
		descriptorpb.FieldDescriptorProto_TYPE_UINT32:   "runtime.Uint32P",
		descriptorpb.FieldDescriptorProto_TYPE_ENUM:     "runtime.EnumP",
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32: "runtime.Int32P",
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64: "runtime.Int64P",
		descriptorpb.FieldDescriptorProto_TYPE_SINT32:   "runtime.Int32P",
		descriptorpb.FieldDescriptorProto_TYPE_SINT64:   "runtime.Int64P",
	}

	proto2RepeatedConvertFuncs = map[descriptorpb.FieldDescriptorProto_Type]string{
		descriptorpb.FieldDescriptorProto_TYPE_DOUBLE:  "runtime.Float64Slice",
		descriptorpb.FieldDescriptorProto_TYPE_FLOAT:   "runtime.Float32Slice",
		descriptorpb.FieldDescriptorProto_TYPE_INT64:   "runtime.Int64Slice",
		descriptorpb.FieldDescriptorProto_TYPE_UINT64:  "runtime.Uint64Slice",
		descriptorpb.FieldDescriptorProto_TYPE_INT32:   "runtime.Int32Slice",
		descriptorpb.FieldDescriptorProto_TYPE_FIXED64: "runtime.Uint64Slice",
		descriptorpb.FieldDescriptorProto_TYPE_FIXED32: "runtime.Uint32Slice",
		descriptorpb.FieldDescriptorProto_TYPE_BOOL:    "runtime.BoolSlice",
		descriptorpb.FieldDescriptorProto_TYPE_STRING:  "runtime.StringSlice",
		// FieldDescriptorProto_TYPE_GROUP
		// FieldDescriptorProto_TYPE_MESSAGE
		// FieldDescriptorProto_TYPE_BYTES
		// TODO(maros7) Handle bytes
		descriptorpb.FieldDescriptorProto_TYPE_UINT32:   "runtime.Uint32Slice",
		descriptorpb.FieldDescriptorProto_TYPE_ENUM:     "runtime.EnumSlice",
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED32: "runtime.Int32Slice",
		descriptorpb.FieldDescriptorProto_TYPE_SFIXED64: "runtime.Int64Slice",
		descriptorpb.FieldDescriptorProto_TYPE_SINT32:   "runtime.Int32Slice",
		descriptorpb.FieldDescriptorProto_TYPE_SINT64:   "runtime.Int64Slice",
	}

	wellKnownTypeConv = map[string]string{
		".google.protobuf.Timestamp":   "runtime.Timestamp",
		".google.protobuf.Duration":    "runtime.Duration",
		".google.protobuf.StringValue": "runtime.StringValue",
		".google.protobuf.FloatValue":  "runtime.FloatValue",
		".google.protobuf.DoubleValue": "runtime.DoubleValue",
		".google.protobuf.BoolValue":   "runtime.BoolValue",
		".google.protobuf.BytesValue":  "runtime.BytesValue",
		".google.protobuf.Int32Value":  "runtime.Int32Value",
		".google.protobuf.UInt32Value": "runtime.UInt32Value",
		".google.protobuf.Int64Value":  "runtime.Int64Value",
		".google.protobuf.UInt64Value": "runtime.UInt64Value",
	}
)
