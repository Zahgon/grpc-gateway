package codegenerator

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func supportedCodeGeneratorFeatures() uint64 {
	_ = "STUB: not implemented"
	// Enable support for Protobuf Editions
	return 0
}

func supportedEditions() (descriptorpb.Edition, descriptorpb.Edition) {
	_ = "STUB: not implemented"
	// Declare support up to edition 2024
	return *new(descriptorpb.Edition), *new(descriptorpb.Edition)
}

// SetSupportedFeaturesOnPluginGen sets supported proto3 features
// on protogen.Plugin.
func SetSupportedFeaturesOnPluginGen(gen *protogen.Plugin) { _ = "STUB: not implemented"; return }

// SetSupportedFeaturesOnCodeGeneratorResponse sets supported proto3 features
// on pluginpb.CodeGeneratorResponse.
func SetSupportedFeaturesOnCodeGeneratorResponse(resp *pluginpb.CodeGeneratorResponse) {
	_ = "STUB: not implemented"
	return
}
