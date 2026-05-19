package genopenapi

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor"
	"google.golang.org/genproto/googleapis/api/annotations"
)

// schemaBuilder owns the in-progress component schema set for one document.
// Construction is intentionally lazy: schemas are only added when something
// references them, transitively from RPC request/response types. We never
// emit unreferenced messages.
type schemaBuilder struct {
	reg *descriptor.Registry
	doc *Document
}

func newSchemaBuilder(reg *descriptor.Registry, doc *Document) *schemaBuilder {
	_ = "STUB: not implemented"
	return nil
}

// fieldSchema returns the schema (or $ref) describing the given proto field's
// type. For repeated fields it produces an array; for map entries it produces
// an object with additionalProperties; for messages and enums it produces a
// $ref into components and ensures the referenced schema is generated.
func (b *schemaBuilder) fieldSchema(field *descriptor.Field) *SchemaOrRef {
	_ = "STUB: not implemented"
	return nil
}

// Map fields are repeated synthetic messages with map_entry=true. Resolve
// the message once: a hit on map_entry takes the map path, a hit on a
// regular message hands the resolved descriptor to scalarOrRef so it does
// not have to look it up again, and a miss emits an open-object array
// directly (avoiding a redundant lookup in scalarOrRef).

// messageRef returns a $ref to the given message's component schema, ensuring
// the component is generated. Well-known types are inlined instead.
func (b *schemaBuilder) messageRef(msg *descriptor.Message) *SchemaOrRef {
	_ = "STUB: not implemented"
	return nil
}

// mapSchema returns the additionalProperties schema for a proto map field.
// Map entries are messages with two fields:
// key (1) and value (2); the value field's type drives additionalProperties.
func (b *schemaBuilder) mapSchema(msg *descriptor.Message) *SchemaOrRef {
	_ = "STUB: not implemented"
	return nil
}

// Should not happen per the proto spec, but emit an open object
// rather than nil which would produce invalid OpenAPI.

// scalarOrRef produces the schema for a non-repeated field type. Message and
// enum types resolve to $refs into components; everything else is an inline
// scalar schema.
func (b *schemaBuilder) scalarOrRef(field *descriptor.Field) *SchemaOrRef {
	_ = "STUB: not implemented"
	return nil
}

// 64-bit ints are JSON strings per protojson.

// OpenAPI has no "uint32" format; int64 is the narrowest that
// covers the full uint32 range.

// Well-known types are inlined; everything else is referenced.

// Fall back to a string value on unhandled field types

// ensureMessageSchema generates a component schema for the given message if
// it has not already been emitted. Map-entry messages are skipped.
func (b *schemaBuilder) ensureMessageSchema(msg *descriptor.Message) {
	_ = "STUB: not implemented"
	return
}

// Map entries are handled separately

// Message already in schema, skip

// Reserve the slot up front so cycles terminate: when a nested field
// recurses into a message we're already building, the exists check
// above short-circuits instead of looping.

// Partition fields into oneof groups and regular fields. Synthetic
// proto3-optional oneofs are treated as regular fields.

// Field is hidden by visibility rules, skip.

// Field is hidden by visibility rules, skip.

// Each proto oneof group constrains its fields to "at most one set"
// (proto3 oneof allows zero or one). Encode that as a JSON Schema
// `oneOf` whose options are:
//
//   1. a "none of the fields are set" guard, expressed as
//      {type: object, not: {anyOf: [{required: F1}, ..., {required: Fn}]}}
//   2. one option per field carrying that field's own typed schema:
//      {type: object, properties: {Fi: <fieldSchema>}, required: [Fi]}
//
// `oneOf` requires exactly one sub-schema to match, so:
//   - zero set      → only the guard matches            → passes
//   - exactly one Fi → only the Fi option matches        → passes
//   - two or more   → multiple Fi options match         → fails
//
// Each per-field option embeds the field's full schema (a $ref for
// nested message types, the inlined schema for scalars and WKTs),
// not just `required`.
//
// The same property schemas also live on the parent's top-level
// `properties` map, so consumers that don't walk into `oneOf` still
// see the full set of fields. The duplication is intentional —
// `properties`, `required`, and `oneOf` are orthogonal in JSON
// Schema and tooling tends to consult them independently.
//
// Single-field oneof groups produce a trivially-true constraint
// ("Fi is either present or not present"); we skip those for output
// hygiene. proto3-optional fields use synthetic single-field oneofs
// and are already filtered by splitOneofs.
//
// Multiple groups are independent constraints, encoded as an `allOf`
// of one `oneOf` per group. A single non-trivial group is hoisted
// directly onto the schema's `oneOf` for less verbose output.

// Collect only visible fields for the constraint.

// Single-field (or empty) oneof groups are a no-op constraint, so skip them.

// No oneof groups, nothing to do.

// A single oneof group can be hoisted directly onto the schema for less verbose output.

// Multiple groups are independent constraints, so combine them with allOf.

// addProperty inserts a single field as a property of schema, attaching
// description (via $ref sibling for refs, or directly for inline schemas)
// and updating Required from field_behavior.
func (b *schemaBuilder) addProperty(schema *Schema, field *descriptor.Field) {
	_ = "STUB: not implemented"
	return
}

// propertySchema produces a property schema for the given field, with
// description and field-behavior flags applied. Description goes on the
// $ref sibling for referenced types, or directly on inline schemas.
//
// A field-level openapiv3_field annotation overrides the comment-derived
// description and may attach a title. Title cannot sit as a $ref sibling,
// so for referenced fields its presence forces an allOf wrapper (same
// mechanism as deprecated/readOnly/writeOnly).
func (b *schemaBuilder) propertySchema(field *descriptor.Field) *SchemaOrRef {
	_ = "STUB: not implemented"
	return nil
}

// 3.1.0 allows description as a sibling of $ref, but title and
// deprecated must live on a real schema body — so an annotation
// that sets either, like readOnly/writeOnly/deprecated cascade,
// forces an allOf wrapper.

// needsAllOfWrap reports whether a referenced field needs an allOf wrapper to
// carry per-occurrence flags (read-only / write-only / deprecated). A pure
// description does not.
func needsAllOfWrap(field *descriptor.Field) bool { _ = "STUB: not implemented"; return false }

// applyFieldFlags writes per-field schema-level flags (deprecated, readOnly,
// writeOnly) onto an inline schema. Required is handled at the parent.
func applyFieldFlags(s *Schema, field *descriptor.Field) { _ = "STUB: not implemented"; return }

// fieldDeprecated reports whether a field should be flagged deprecated in
// the emitted schema. The flag cascades from the enclosing message and
// file: a field is deprecated if the field itself, its containing message,
// or the proto file that declares it is marked with `deprecated = true`.
func fieldDeprecated(field *descriptor.Field) bool { _ = "STUB: not implemented"; return false }

// methodDeprecated reports whether an RPC method should be flagged
// deprecated. The flag cascades from the service and file: a method is
// deprecated if the method itself, its service, or the proto file is
// marked deprecated. OpenAPI 3.1 has no `deprecated` flag on tags, so
// cascading service-level deprecation into every method of the service
// is the only way to reflect it.
func methodDeprecated(m *descriptor.Method) bool { _ = "STUB: not implemented"; return false }

// messageDeprecated reports whether a message component schema should
// be flagged deprecated. The flag cascades from the file: a message is
// deprecated if the message itself or the proto file is marked
// deprecated. Nested types are left alone — they're separately-named
// component schemas and deprecate independently of any outer message.
func messageDeprecated(msg *descriptor.Message) bool { _ = "STUB: not implemented"; return false }

// enumDeprecated reports whether an enum component schema should be
// flagged deprecated. The flag cascades from the file.
func enumDeprecated(enum *descriptor.Enum) bool { _ = "STUB: not implemented"; return false }

// ensureEnumSchema generates a component schema for the given enum if it
// has not been emitted yet. Enum values are rendered as their string names.
//
// The grpc-gateway runtime (via protojson) actually accepts enum values on
// the wire as either string names or integer numbers, but encoding that
// dual acceptance in the schema breaks consumer tooling: openapi-generator-
// cli's Go target produces an unbuildable wrapper type for both the
// `oneOf`-of-homogeneous-enums form (pulls in gopkg.in/validator.v2 plus a
// painful discriminator wrapper) and the `type: [string, integer]`
// mixed-type form (literally writes `anyOf<string,integer>` as the Go type
// name — an unimplemented codegen branch). Until tooling catches up, we
// document the string form in the spec and rely on protojson's leniency
// to accept the integer form at the gateway boundary.
func (b *schemaBuilder) ensureEnumSchema(enum *descriptor.Enum) { _ = "STUB: not implemented"; return }

// Enum value is hidden by visibility rules, skip.

// Every value is hidden by visibility rules. We can't drop the
// component schema — a visible field may still reference it — so
// fall back to an unconstrained string and warn. The user almost
// certainly wants to either broaden their selectors or hide the
// referring fields as well.

// oneofGroup is a oneof declaration plus its constituent fields.
type oneofGroup struct {
	name   string
	fields []*descriptor.Field
}

// splitOneofs separates regular fields from real (non-synthetic) oneof groups
// in declaration order. proto3-optional fields use synthetic single-field
// oneofs which we treat as regular optional fields.
func splitOneofs(msg *descriptor.Message) (regular []*descriptor.Field, groups []oneofGroup) {
	_ = "STUB: not implemented"
	return nil, nil
}

// fieldBehaviors returns the [(google.api.field_behavior) = ...] entries on a
// field, or nil if there are none.
func fieldBehaviors(field *descriptor.Field) []annotations.FieldBehavior {
	_ = "STUB: not implemented"
	return nil
}

func hasFieldBehavior(field *descriptor.Field, target annotations.FieldBehavior) bool {
	_ = "STUB: not implemented"
	return false
}

// jsonName returns the JSON name we use for a proto field. We always honor
// the proto json_name option since the runtime uses
// JSON names by default and we'd rather match what the wire produces.
func jsonName(field *descriptor.Field) string { _ = "STUB: not implemented"; return "" }
