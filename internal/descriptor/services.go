package descriptor

import (
	options "google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/types/descriptorpb"
)

// loadServices registers services and their methods from "targetFile" to "r".
// It must be called after loadFile is called for all files so that loadServices
// can resolve names of message types and their fields.
func (r *Registry) loadServices(file *File) error { _ = "STUB: not implemented"; return nil }

func (r *Registry) newMethod(svc *Service, md *descriptorpb.MethodDescriptorProto, optsList []*options.HttpRule) (*Method, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(yugui) Handle query params

func extractAPIOptions(meth *descriptorpb.MethodDescriptorProto) (*options.HttpRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func defaultAPIOptions(svc *Service, md *descriptorpb.MethodDescriptorProto) (*options.HttpRule, error) {
	_ = "STUB: not implemented"
	// FQSN prefixes the service's full name with a '.', e.g.: '.example.ExampleService'
	return nil, nil
}

// This generates an HttpRule that matches the gRPC mapping to HTTP/2 described in
// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md#requests
// i.e.:
//   * method is POST
//   * path is "/<service name>/<method name>"
//   * body should contain the serialized request message

func (r *Registry) newParam(meth *Method, path string) (Parameter, error) {
	_ = "STUB: not implemented"
	return *new(Parameter), nil
}

func (r *Registry) newBody(meth *Method, path string) (*Body, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registry) newResponse(meth *Method, path string) (*Body, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// lookupField looks up a field named "name" within "msg".
// It returns nil if no such field found.
func lookupField(msg *Message, name string) *Field { _ = "STUB: not implemented"; return nil }

// resolveFieldPath resolves "path" into a list of fieldDescriptor, starting from "msg".
func (r *Registry) resolveFieldPath(msg *Message, path string, isPathParam bool) ([]FieldPathComponent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
