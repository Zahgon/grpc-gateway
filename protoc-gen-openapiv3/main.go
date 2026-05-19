// Command protoc-gen-openapiv3 implements an OpenAPI 3.1.0 generator for
// proto files annotated with google.api.http rules.
//
// Status: alpha. The emitted JSON shape is not yet stable — encodings for
// oneofs, wrapper types, enums, and path-template expansion may change
// between minor releases while the mapping rules settle in response to
// real-world feedback. For a production-stable OpenAPI pipeline today,
// use protoc-gen-openapiv2.
package main

import (
	"flag"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/protobuf/types/pluginpb"
)

var (
	visibilityRestrictionSelectors = utilities.StringArrayFlag(flag.CommandLine, "visibility_restriction_selectors", "list of `google.api.VisibilityRule` visibility labels to include in the generated output when a visibility annotation is defined. Repeat this option to supply multiple values. Elements without visibility annotations are unaffected by this setting.")
	disableDefaultErrors           = flag.Bool("disable_default_errors", false, "if set, disables generation of default errors. This is useful if you have defined custom error handling")
)

func main() {
	if err := run(); err != nil {
		emitError(err)
		os.Exit(1)
	}
}

func run() error { _ = "STUB: not implemented"; return nil }

func emitFiles(files []*pluginpb.CodeGeneratorResponse_File) { _ = "STUB: not implemented"; return }

func emitError(err error) {
	_ = "STUB: not implemented"
	// Echo to stderr in addition to the proto response
	return
}

func emitResp(resp *pluginpb.CodeGeneratorResponse) { _ = "STUB: not implemented"; return }

// parseReqParam parses the protoc plugin parameter string (comma-separated
// key=value pairs) and sets the corresponding flags. When a flag name is
// present without a value (e.g. "disable_default_errors"), it is treated as
// "true", so that boolean flags can be enabled by name alone.
func parseReqParam(param string, f *flag.FlagSet) error { _ = "STUB: not implemented"; return nil }
