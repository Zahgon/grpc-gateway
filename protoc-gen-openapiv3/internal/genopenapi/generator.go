package genopenapi

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor"
	"google.golang.org/protobuf/types/pluginpb"
)

// Generate produces one OpenAPI 3.1.0 JSON document per input proto file.
// Files without HTTP-bound services are skipped (no output file is emitted).
func Generate(reg *descriptor.Registry, files []*descriptor.File) ([]*pluginpb.CodeGeneratorResponse_File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// No HTTP annotations found, skip file.

// generateFile builds a Document for a single proto file. The boolean return
// is false when the file has no HTTP-bound operations to emit.
func generateFile(reg *descriptor.Registry, file *descriptor.File) (*Document, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// Prime seenTags with any document-level annotation tags so a service's
// default tag does not clobber the annotation-provided description.
// applyDocumentOverride must run before this point so its tags are
// included; the post-iteration orphan check below depends on this set
// being complete.

// Tracks operationIds across all bindings of all services in this file
// to enforce the OpenAPI 3.1.0 uniqueness rule. Generator-derived ids
// (`<Service>_<Method>`) are naturally unique; collisions here mean
// two annotations (or an annotation colliding with a default) chose
// the same id.

// Collected as we go so we can validate against the final tag set
// after service iteration completes (services may add their default
// tag after their operations have already declared tags).

// Service is hidden by visibility rules, skip.

// Track whether any methods of this service are visible,
// to avoid emitting a tag for a service with only hidden methods.

// Method is hidden by visibility rules, skip.

// At least one method is visible, so the service's tag
// should be emitted.

// Emit the service's tag if it has any visible
// methods and the tag hasn't already been emitted.

// Validate that every tag referenced by an operation is declared
// somewhere — either by a service-derived default or by a
// document-level openapiv3_document.tags annotation. An orphan tag
// would render as a name with no metadata in the document, which is
// almost always a mistake (a typo, or a forgotten doc.tags entry).

// opTagRef remembers a tag reference from an operation, deferred for
// validation after service iteration completes.
type opTagRef struct {
	operationID string
	tag         string
}
