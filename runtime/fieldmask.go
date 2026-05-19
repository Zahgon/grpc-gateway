package runtime

import (
	"io"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	field_mask "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func getFieldByName(fields protoreflect.FieldDescriptors, name string) protoreflect.FieldDescriptor {
	_ = "STUB: not implemented"
	return *new(protoreflect.FieldDescriptor)
}

// FieldMaskFromRequestBody creates a FieldMask printing all complete paths from the JSON body.
func FieldMaskFromRequestBody(r io.Reader, msg proto.Message) (*field_mask.FieldMask, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// dequeue an item

// if the item is an object, then enqueue all of its children

// As per: https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/field_mask.proto#L85-L86
// Do not recurse into repeated fields. The repeated field goes on the end of the path and we stop.

// otherwise, it's a leaf node so print its path

// Sort for deterministic output in the presence
// of repeated fields.

func isProtobufAnyMessage(md protoreflect.MessageDescriptor) bool {
	_ = "STUB: not implemented"
	return false
}

func isDynamicProtoMessage(md protoreflect.MessageDescriptor) bool {
	_ = "STUB: not implemented"
	return false
}

// buildPathsBlindly does not attempt to match proto field names to the
// json value keys.  Instead it relies completely on the structure of
// the unmarshalled json contained within in.
// Returns a slice containing all subpaths with the root at the
// passed in name and json value.
func buildPathsBlindly(name string, in interface{}) []string { _ = "STUB: not implemented"; return nil }

// This should never happen since we should always check that we only add
// nodes of type map[string]interface{} to the queue.

// This is not a struct, so there are no more levels to descend.

// fieldMaskPathItem stores an in-progress deconstruction of a path for a fieldmask
type fieldMaskPathItem struct {
	// the list of prior fields leading up to node connected by dots
	path string

	// a generic decoded json object the current item to inspect for further path extraction
	node interface{}

	// parent message
	msg protoreflect.Message
}
