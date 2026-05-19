package descriptor

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor/openapiconfig"
)

func loadOpenAPIConfigFromYAML(yamlFileContents []byte, yamlSourceLogName string) (*openapiconfig.OpenAPIConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reject unknown fields because OpenAPIConfig is only used here

func registerOpenAPIOptions(registry *Registry, openAPIConfig *openapiconfig.OpenAPIConfig, yamlSourceLogName string) error {
	_ = "STUB: not implemented"
	return nil
}

// Nothing to do

// LoadOpenAPIConfigFromYAML loads an  OpenAPI Configuration from the given YAML file
// and registers the OpenAPI options the given registry.
// This must be done after loading the proto file.
func (r *Registry) LoadOpenAPIConfigFromYAML(yamlFile string) error {
	_ = "STUB: not implemented"
	return nil
}
