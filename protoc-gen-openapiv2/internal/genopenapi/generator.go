package genopenapi

import (
	"errors"

	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor"
	gen "github.com/grpc-ecosystem/grpc-gateway/v2/internal/generator"
	"go.yaml.in/yaml/v3"
)

var errNoTargetService = errors.New("no target service defined in the file")

type generator struct {
	reg    *descriptor.Registry
	format Format
}

type wrapper struct {
	fileName string
	swagger  *openapiSwaggerObject
}

type GeneratorOptions struct {
	Registry       *descriptor.Registry
	RecursiveDepth int
}

// New returns a new generator which generates grpc gateway files.
func New(reg *descriptor.Registry, format Format) gen.Generator {
	_ = "STUB: not implemented"
	return *new(gen.Generator)
}

// Merge a lot of OpenAPI file (wrapper) to single one OpenAPI file
func mergeTargetFile(targets []*wrapper, mergeFileName string) *wrapper {
	_ = "STUB: not implemented"
	return nil
}

// Q: What's up with the alias types here?
// A: We don't want to completely override how these structs are marshaled into
// JSON, we only want to add fields (see below, extensionMarshalJSON).
// An infinite recursion would happen if we'd call json.Marshal on the struct
// that has swaggerObject as an embedded field. To avoid that, we'll create
// type aliases, and those don't have the custom MarshalJSON methods defined
// on them. See http://choly.ca/post/go-json-marshalling/ (or, if it ever
// goes away, use
// https://web.archive.org/web/20190806073003/http://choly.ca/post/go-json-marshalling/).
func (so openapiSwaggerObject) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalYAML implements yaml.Marshaler interface.
//
// It is required in order to pass extensions inline.
//
// Example:
//
//	extensions: {x-key: x-value}
//	type: string
//
// It will be rendered as:
//
//	x-key: x-value
//	type: string
//
// Use generics when the project will be upgraded to go 1.18+.
func (so openapiSwaggerObject) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Custom json marshaller for openapiPathsObject. Ensures
// openapiPathsObject is marshalled into expected format in generated
// swagger.json.
func (po openapiPathsObject) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// marshal key

// marshal value

// Custom yaml marshaller for openapiPathsObject. Ensures
// openapiPathsObject is marshalled into expected format in generated
// swagger.yaml.
func (po openapiPathsObject) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We can simplify this implementation once the go-yaml bug is resolved. See: https://github.com/go-yaml/yaml/issues/643.
//
//	func (pio *openapiPathItemObject) toYAMLNode() (*yaml.Node, error) {
//		var node yaml.Node
//		if err := node.Encode(pio); err != nil {
//			return nil, err
//		}
//		return &node, nil
//	}
func (pio *openapiPathItemObject) toYAMLNode() (*yaml.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiInfoObject) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiInfoObject) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiSecuritySchemeObject) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiSecuritySchemeObject) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiOperationObject) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiOperationObject) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiResponseObject) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiResponseObject) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiSchemaObject) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiSchemaObject) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiParameterObject) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiParameterObject) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiTagObject) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (so openapiTagObject) MarshalYAML() (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func extensionMarshalJSON(so interface{}, extensions []extension) ([]byte, error) {
	_ = "STUB: not implemented"
	// To append arbitrary keys to the struct we'll render into json,
	// we're creating another struct that embeds the original one, and
	// its extra fields:
	//
	// The struct will look like
	// struct {
	//   *openapiCore
	//   XGrpcGatewayFoo json.RawMessage `json:"x-grpc-gateway-foo"`
	//   XGrpcGatewayBar json.RawMessage `json:"x-grpc-gateway-bar"`
	// }
	// and thus render into what we want -- the JSON of openapiCore with the
	// extensions appended.
	return nil, nil
}

// embedded

// encodeOpenAPI converts OpenAPI file obj to pluginpb.CodeGeneratorResponse_File
func encodeOpenAPI(file *wrapper, format Format) (*descriptor.ResponseFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func deprecateFieldsAndMethods(file *descriptor.File) { _ = "STUB: not implemented"; return }

func (g *generator) Generate(targets []*descriptor.File) ([]*descriptor.ResponseFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Because of how the generator merges definitions, it is simpler to deprecate field and methods here if the file is deprecated

// try to find proto leader

// merge protos to leader

func (so openapiSwaggerObject) sortPathsAlphabetically() { _ = "STUB: not implemented"; return }

// AddErrorDefs Adds google.rpc.Status and google.protobuf.Any
// to registry (used for error-related API responses)
func AddErrorDefs(reg *descriptor.Registry) error {
	_ = "STUB: not implemented"
	// load internal protos
	return nil
}

func extensionsToMap(extensions []extension) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}
