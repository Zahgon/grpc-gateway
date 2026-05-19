package descriptor

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor/apiconfig"
)

func loadGrpcAPIServiceFromYAML(yamlFileContents []byte, yamlSourceLogName string) (*apiconfig.GrpcAPIService, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// As our GrpcAPIService is incomplete, accept unknown fields.

func registerHTTPRulesFromGrpcAPIService(registry *Registry, service *apiconfig.GrpcAPIService, sourceLogName string) error {
	_ = "STUB: not implemented"
	return nil

	// Nothing to do
}

// LoadGrpcAPIServiceFromYAML loads a gRPC API Configuration from the given YAML file
// and registers the HttpRule descriptions contained in it as externalHTTPRules in
// the given registry. This must be done before loading the proto file.
//
// You can learn more about gRPC API Service descriptions from Google's documentation
// at https://cloud.google.com/endpoints/docs/grpc/grpc-service-config
//
// Note that for the purposes of the gateway generator we only consider a subset of all
// available features google supports in their service descriptions.
func (r *Registry) LoadGrpcAPIServiceFromYAML(yamlFile string) error {
	_ = "STUB: not implemented"
	return nil
}
