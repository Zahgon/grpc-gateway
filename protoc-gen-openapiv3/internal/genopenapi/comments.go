package genopenapi

import (
	"regexp"

	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor"
	"google.golang.org/protobuf/types/descriptorpb"
)

// Proto path indices into the FileDescriptorProto for source code info.
// These mirror the field numbers in descriptor.proto and are how protoc
// records line ranges per declaration.
var (
	messageProtoPath  = protoFieldNumber[descriptorpb.FileDescriptorProto]("MessageType")
	enumProtoPath     = protoFieldNumber[descriptorpb.FileDescriptorProto]("EnumType")
	serviceProtoPath  = protoFieldNumber[descriptorpb.FileDescriptorProto]("Service")
	methodProtoPath   = protoFieldNumber[descriptorpb.ServiceDescriptorProto]("Method")
	nestedProtoPath   = protoFieldNumber[descriptorpb.DescriptorProto]("NestedType")
	fieldProtoPath    = protoFieldNumber[descriptorpb.DescriptorProto]("Field")
	enumTypeProtoPath = protoFieldNumber[descriptorpb.DescriptorProto]("EnumType")
)

// protoFieldNumber pulls the field number out of a descriptorpb struct's
// protobuf tag. We resolve this once at startup so a future descriptor.proto
// rename is caught immediately rather than silently producing wrong comment
// lookups.
func protoFieldNumber[T any](fieldName string) int32 { _ = "STUB: not implemented"; return 0 }

// Tag format: "bytes,4,rep,name=message_type,..."; the second comma-
// separated component is always the field number.

// internalCommentPattern matches AIP-192 internal comment markers (-- ... --).
var internalCommentPattern = regexp.MustCompile(`(?s)\(--.*?--\)`)

// extractComments returns the leading comment for the given proto path within
// a file. Returns an empty string if no comment is present or the file has no
// source-code info.
func extractComments(file *descriptor.File, path []int32) string {
	_ = "STUB: not implemented"
	return ""
}

// Strip a single leading space from continuation lines. This
// handles the common protoc output shape but not tabs, deeper
// indentation, or doxygen-style "\n *" prefixes; comments using
// those will render with ragged margins.

// splitSummaryDescription splits a comment block into a one-line summary
// (the first paragraph) and a longer description (the remainder), following
// the convention used by Google AIP and the v2 generator.
func splitSummaryDescription(comment string) (summary, description string) {
	_ = "STUB: not implemented"
	return "", ""
}

// serviceComments returns the leading comment on a service declaration.
func serviceComments(svc *descriptor.Service) string { _ = "STUB: not implemented"; return "" }

// methodComments returns the leading comment on a method declaration.
func methodComments(m *descriptor.Method) string { _ = "STUB: not implemented"; return "" }

// messageComments returns the leading comment on a message declaration,
// walking outer messages so nested types resolve correctly.
func messageComments(reg *descriptor.Registry, msg *descriptor.Message) string {
	_ = "STUB: not implemented"
	return ""
}

// fieldComments returns the leading comment on a field declaration.
func fieldComments(reg *descriptor.Registry, field *descriptor.Field) string {
	_ = "STUB: not implemented"
	return ""
}

// enumComments returns the leading comment on an enum declaration.
func enumComments(reg *descriptor.Registry, enum *descriptor.Enum) string {
	_ = "STUB: not implemented"
	return ""
}

// messagePath builds the source-code-info path to a (possibly nested) message.
func messagePath(reg *descriptor.Registry, msg *descriptor.Message) []int32 {
	_ = "STUB: not implemented"
	return nil
}

// enumPath builds the source-code-info path to a (possibly nested) enum.
func enumPath(reg *descriptor.Registry, enum *descriptor.Enum) []int32 {
	_ = "STUB: not implemented"
	return nil
}

// outerPath walks a chain of outer message names and returns their indices,
// separated by the nested-type field number, suitable for appending to a
// source-code-info path.
func outerPath(reg *descriptor.Registry, file *descriptor.File, outers []string) []int32 {
	_ = "STUB: not implemented"
	return nil
}
