package genopenapi

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv3/options"
)

// Annotation lookups. Each returns (nil, false) when the extension is not
// present. A returned annotation's non-empty sub-fields replace defaults the
// generator would otherwise derive from proto comments or proto types.

func fileDocumentAnnotation(file *descriptor.File) (*options.Document, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func methodOperationAnnotation(m *descriptor.Method) (*options.Operation, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func messageSchemaAnnotation(msg *descriptor.Message) (*options.Schema, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func fieldSchemaAnnotation(field *descriptor.Field) (*options.Schema, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// applyDocumentOverride applies file-level Document overrides onto the
// generated OpenAPI document. Non-empty fields replace defaults; empty fields
// leave the current value untouched. Returns an error if the annotation is
// invalid — for example, a License with no name, a Server with no url, a Tag
// with no name, or any ExternalDocs without a url. All four are spec-required
// fields per OpenAPI 3.1.0.
func applyDocumentOverride(doc *Document, d *options.Document) error {
	_ = "STUB: not implemented"
	return nil
}

// validateServer enforces the OpenAPI 3.1.0 Server Object's required `url`
// field. Empty url is invalid even though `description` alone may look
// useful in proto.
func validateServer(s *options.Server) error { _ = "STUB: not implemented"; return nil }

// validateExternalDocs enforces the OpenAPI 3.1.0 External Documentation
// Object's required `url` field.
func validateExternalDocs(ed *options.ExternalDocs) error { _ = "STUB: not implemented"; return nil }

// applyOperationOverride applies method-level Operation overrides onto the
// generated operation. Annotation values replace comment-derived summary and
// description; a non-empty tag list replaces the default (the service name);
// servers from the annotation are appended to any defaults. The annotation
// `deprecated` flag is one-way: it can flip deprecation on, but cannot
// clear a flag inherited from the proto cascade.
//
// Returns an error if the annotation contains an external_docs without a url
// or a server without a url, both spec-required per OpenAPI 3.1.0.
func applyOperationOverride(op *Operation, o *options.Operation) error {
	_ = "STUB: not implemented"
	return nil
}

// applySchemaBodyOverride applies the title and deprecated fields from a
// Schema annotation onto a schema body. Used for both message-level and
// field-level annotations. For $ref-typed fields the caller must first
// ensure an allOf wrapper exists, since neither field can sit alongside
// $ref without one. The annotation `deprecated` flag is one-way: it can
// flip deprecation on, but cannot clear a flag inherited from the proto
// cascade.
func applySchemaBodyOverride(s *Schema, o *options.Schema) { _ = "STUB: not implemented"; return }

// applyMessageSchemaOverride applies a message-level Schema annotation onto
// a component schema.
func applyMessageSchemaOverride(s *Schema, o *options.Schema) { _ = "STUB: not implemented"; return }

// annotationNeedsSchemaBody reports whether a field annotation sets anything
// that cannot be expressed as a $ref sibling in OpenAPI 3.1.0. `title` and
// `deprecated` both require a real schema body; `description` can sit as a
// $ref sibling directly. Used by propertySchema to decide when a referenced
// field needs an allOf wrapper.
func annotationNeedsSchemaBody(o *options.Schema) bool { _ = "STUB: not implemented"; return false }
