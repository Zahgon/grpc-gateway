package genopenapi

// wktSchemas maps fully-qualified protobuf well-known type names (with leading
// dot, as proto descriptors return them) to the inline OpenAPI schema we emit
// when a field references that type.
//
// All WKTs are inlined; this generator does not emit reusable component
// schemas for them. Behavior matches protojson's serialization model:
//   - Timestamp / Duration / FieldMask: string forms
//   - Wrapper types: their underlying primitive
//   - Empty / Struct: object
//   - Value: any (no type constraint)
//   - ListValue: array of any
//   - NullValue: literal null
//   - Any: object with @type plus open additional properties
//
// Wrapper types are intentionally not marked nullable. The strictly-correct
// JSON Schema 2020-12 form is `type: [<primitive>, "null"]`, but no Go
// OpenAPI generator in the ecosystem handles that form yet: openapi-
// generator-cli writes a literal `anyOf<string,integer>` placeholder for
// it, and oapi-codegen errors with "unhandled Schema type: &[string null]"
// on request bodies that reference it. Until tooling catches up we
// describe wrappers as their underlying primitive; the gateway runtime
// still accepts a JSON null on the wire because protojson treats wrappers
// as optional regardless.
//
// Each call should return a fresh schema so callers can mutate it without
// risk of cross-contamination.
func wellKnownTypeSchema(typeName string) *Schema { _ = "STUB: not implemented"; return nil }

// uint32's range (0..4294967295) exceeds int32; OpenAPI has no
// "uint32" format, so int64 is the narrowest that fits.
