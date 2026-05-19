package genopenapi

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor"
	"google.golang.org/genproto/googleapis/api/visibility"
	"google.golang.org/protobuf/types/descriptorpb"
)

// isVisible checks whether an element with the given VisibilityRule should be
// included in the generated output. Elements without an annotation are always
// visible. When an annotation is present, at least one of its comma-separated
// restriction labels must appear in the registry's configured selectors.
func isVisible(r *visibility.VisibilityRule, reg *descriptor.Registry) bool {
	_ = "STUB: not implemented"
	return false
}

func fieldVisibility(fd *descriptor.Field) *visibility.VisibilityRule {
	_ = "STUB: not implemented"
	return nil
}

func serviceVisibility(svc *descriptor.Service) *visibility.VisibilityRule {
	_ = "STUB: not implemented"
	return nil
}

func methodVisibility(m *descriptor.Method) *visibility.VisibilityRule {
	_ = "STUB: not implemented"
	return nil
}

func enumValueVisibility(v *descriptorpb.EnumValueDescriptorProto) *visibility.VisibilityRule {
	_ = "STUB: not implemented"
	return nil
}
