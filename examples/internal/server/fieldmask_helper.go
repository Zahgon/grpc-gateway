package server

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	field_mask "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func applyFieldMask(patchee, patcher proto.Message, mask *field_mask.FieldMask) {
	_ = "STUB: not implemented"
	return
}

func getField(msg protoreflect.Message, path string) (field protoreflect.FieldDescriptor, parent protoreflect.Message) {
	_ = "STUB: not implemented"
	return *new(protoreflect.FieldDescriptor), *new(protoreflect.Message)
}
