package descriptor

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor/openapiconfig"
	"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

// Registry is a registry of information extracted from pluginpb.CodeGeneratorRequest.
type Registry struct {
	// msgs is a mapping from fully-qualified message name to descriptor
	msgs map[string]*Message

	// enums is a mapping from fully-qualified enum name to descriptor
	enums map[string]*Enum

	// files is a mapping from file path to descriptor
	files map[string]*File

	// meths is a mapping from fully-qualified method name to descriptor
	meths map[string]*Method

	// prefix is a prefix to be inserted to golang package paths generated from proto package names.
	prefix string

	// pkgMap is a user-specified mapping from file path to proto package.
	pkgMap map[string]string

	// pkgAliases is a mapping from package aliases to package paths in go which are already taken.
	pkgAliases map[string]string

	// allowDeleteBody permits http delete methods to have a body
	allowDeleteBody bool

	// externalHttpRules is a mapping from fully qualified service method names to additional HttpRules applicable besides the ones found in annotations.
	externalHTTPRules map[string][]*annotations.HttpRule

	// allowMerge generation one OpenAPI file out of multiple protos
	allowMerge bool

	// mergeFileName target OpenAPI file name after merge
	mergeFileName string

	// includePackageInTags controls whether the package name defined in the `package` directive
	// in the proto file can be prepended to the gRPC service name in the `Tags` field of every operation.
	includePackageInTags bool

	// repeatedPathParamSeparator specifies how path parameter repeated fields are separated
	repeatedPathParamSeparator repeatedFieldSeparator

	// useJSONNamesForFields if true json tag name is used for generating fields in OpenAPI definitions,
	// otherwise the original proto name is used. It's helpful for synchronizing the OpenAPI definition
	// with gRPC-Gateway response, if it uses json tags for marshaling.
	useJSONNamesForFields bool

	// useProto3FieldSemantics if true proto3 field semantics are used for generating fields in OpenAPI definitions.
	useProto3FieldSemantics bool

	// openAPINamingStrategy is the naming strategy to use for assigning OpenAPI field and parameter names. This can be one of the following:
	// - `legacy`: use the legacy naming strategy from protoc-gen-swagger, that generates unique but not necessarily
	//             maximally concise names. Components are concatenated directly, e.g., `MyOuterMessageMyNestedMessage`.
	// - `simple`: use a simple heuristic for generating unique and concise names. Components are concatenated using
	//             dots as a separator, e.g., `MyOuterMessage.MyNestedMessage` (if `MyNestedMessage` alone is unique,
	//             `MyNestedMessage` will be used as the OpenAPI name).
	// - `fqn`:    always use the fully-qualified name of the proto message (leading dot removed) as the OpenAPI
	//             name.
	openAPINamingStrategy string

	// visibilityRestrictionSelectors is a map of selectors for `google.api.VisibilityRule`s that will be included in the OpenAPI output.
	visibilityRestrictionSelectors map[string]bool

	// useGoTemplate determines whether you want to use GO templates
	// in your protofile comments
	useGoTemplate bool

	// goTemplateArgs specifies a list of key value pair inputs to be displayed in Go templates
	goTemplateArgs map[string]string

	// ignoreComments determines whether all protofile comments should be excluded from output
	ignoreComments bool

	// removeInternalComments determines whether to remove substrings in comments that begin with
	// `(--` and end with `--)` as specified in https://google.aip.dev/192#internal-comments.
	removeInternalComments bool

	// enumsAsInts render enum as integer, as opposed to string
	enumsAsInts bool

	// omitEnumDefaultValue omits default value of enum
	omitEnumDefaultValue bool

	// disableDefaultErrors disables the generation of the default error types.
	// This is useful for users who have defined custom error handling.
	disableDefaultErrors bool

	// simpleOperationIDs removes the service prefix from the generated
	// operationIDs. This risks generating duplicate operationIDs.
	simpleOperationIDs bool

	standalone bool
	// warnOnUnboundMethods causes the registry to emit warning logs if an RPC method
	// has no HttpRule annotation.
	warnOnUnboundMethods bool

	// proto3OptionalNullable specifies whether Proto3 Optional fields should be marked as x-nullable.
	proto3OptionalNullable bool

	// fileOptions is a mapping of file name to additional OpenAPI file options
	fileOptions map[string]*options.Swagger

	// methodOptions is a mapping of fully-qualified method name to additional OpenAPI method options
	methodOptions map[string]*options.Operation

	// messageOptions is a mapping of fully-qualified message name to additional OpenAPI message options
	messageOptions map[string]*options.Schema

	//serviceOptions is a mapping of fully-qualified service name to additional OpenAPI service options
	serviceOptions map[string]*options.Tag

	// fieldOptions is a mapping of the fully-qualified name of the parent message concat
	// field name and a period to additional OpenAPI field options
	fieldOptions map[string]*options.JSONSchema

	// generateUnboundMethods causes the registry to generate proxy methods even for
	// RPC methods that have no HttpRule annotation.
	generateUnboundMethods bool

	// omitPackageDoc, if false, causes a package comment to be included in the generated code.
	omitPackageDoc bool

	// recursiveDepth sets the maximum depth of a field parameter
	recursiveDepth int

	// annotationMap is used to check for duplicate HTTP annotations
	annotationMap map[annotationIdentifier]struct{}

	// disableServiceTags disables the generation of service tags.
	// This is useful if you do not want to expose the names of your backend grpc services.
	disableServiceTags bool

	// disableDefaultResponses disables the generation of default responses.
	// Useful if you have to support custom response codes that are not 200.
	disableDefaultResponses bool

	// useAllOfForRefs, if set, will use allOf as container for $ref to preserve same-level
	// properties
	useAllOfForRefs bool

	// omitArrayItemTypeWhenRefSibling, if set, will omit 'type: object' in array items when $ref is present
	// to avoid no-$ref-siblings violations in OpenAPI v2
	omitArrayItemTypeWhenRefSibling bool

	// allowPatchFeature determines whether to use PATCH feature involving update masks (using google.protobuf.FieldMask).
	allowPatchFeature bool

	// preserveRPCOrder, if true, will ensure the order of paths emitted in openapi swagger files mirror
	// the order of RPC methods found in proto files. If false, emitted paths will be ordered alphabetically.
	preserveRPCOrder bool

	// enableRpcDeprecation whether to process grpc method's deprecated option
	enableRpcDeprecation bool

	// enableFieldDeprecation whether to process proto field's deprecated option
	enableFieldDeprecation bool

	// expandSlashedPathPatterns, if true, for a path parameter carrying a sub-path, described via parameter pattern (i.e.
	// the pattern contains forward slashes), this will expand the _pattern_ into the URI and will _replace_ the parameter
	// with new path parameters inferred from patterns wildcards.
	//
	// Example: a Google AIP style path "/v1/{name=projects/*/locations/*}/datasets/{dataset}" with a "name" parameter
	// containing sub-path will generate "/v1/projects/{project}/locations/{location}/datasets/{dataset}" path in OpenAPI.
	// Note that the original "name" parameter is replaced with "project" and "location" parameters.
	//
	// This leads to more compliant and readable OpenAPI suitable for documentation, but may complicate client
	// implementation if you want to pass the original "name" parameter.
	expandSlashedPathPatterns bool

	// generateXGoType is a global generator option for generating x-go-type annotations
	generateXGoType bool
}

type repeatedFieldSeparator struct {
	name string
	sep  rune
}

type annotationIdentifier struct {
	method       string
	pathTemplate string
	service      *Service
}

// NewRegistry returns a new Registry.
func NewRegistry() *Registry { _ = "STUB: not implemented"; return nil }

// Load loads definitions of services, methods, messages, enumerations and fields from "req".
func (r *Registry) Load(req *pluginpb.CodeGeneratorRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// Note: keep in mind that this might be not enough because
// protogen.Plugin is used only to load files here.
// The support for features must be set on the pluginpb.CodeGeneratorResponse.

func (r *Registry) LoadFromPlugin(gen *protogen.Plugin) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Registry) load(gen *protogen.Plugin) error { _ = "STUB: not implemented"; return nil }

// loadFile loads messages, enumerations and fields from "file".
// It does not load services and methods in "file".  You need to call
// loadServices after loadFiles is called for all files to load services and methods.
func (r *Registry) loadFile(filePath string, file *protogen.File) {
	_ = "STUB: not implemented"
	return
}

func (r *Registry) registerMsg(file *File, outerPath []string, msgs []*descriptorpb.DescriptorProto) {
	_ = "STUB: not implemented"
	return
}

func (r *Registry) registerEnum(file *File, outerPath []string, enums []*descriptorpb.EnumDescriptorProto) {
	_ = "STUB: not implemented"
	return
}

// LookupMsg looks up a message type by "name".
// It tries to resolve "name" from "location" if "name" is a relative message name.
func (r *Registry) LookupMsg(location, name string) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LookupEnum looks up an enum type by "name".
// It tries to resolve "name" from "location" if "name" is a relative enum name.
func (r *Registry) LookupEnum(location, name string) (*Enum, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LookupFile looks up a file by name.
func (r *Registry) LookupFile(name string) (*File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Registry) GetUseProto3FieldSemantics() bool { _ = "STUB: not implemented"; return false }

func (r *Registry) SetUseProto3FieldSemantics(useProto3FieldSemantics bool) {
	_ = "STUB: not implemented"
	return
}

// LookupExternalHTTPRules looks up external http rules by fully qualified service method name
func (r *Registry) LookupExternalHTTPRules(qualifiedMethodName string) []*annotations.HttpRule {
	_ = "STUB: not implemented"
	return nil
}

// AddExternalHTTPRule adds an external http rule for the given fully qualified service method name
func (r *Registry) AddExternalHTTPRule(qualifiedMethodName string, rule *annotations.HttpRule) {
	_ = "STUB: not implemented"
	return
}

// UnboundExternalHTTPRules returns the list of External HTTPRules
// which does not have a matching method in the registry
func (r *Registry) UnboundExternalHTTPRules() []string { _ = "STUB: not implemented"; return nil }

// AddPkgMap adds a mapping from a .proto file to proto package name.
func (r *Registry) AddPkgMap(file, protoPkg string) { _ = "STUB: not implemented"; return }

// SetPrefix registers the prefix to be added to go package paths generated from proto package names.
func (r *Registry) SetPrefix(prefix string) {
	_ = "STUB: not implemented"

	// SetStandalone registers standalone flag to control package prefix
	return
}

func (r *Registry) SetStandalone(standalone bool) { _ = "STUB: not implemented"; return }

// SetRecursiveDepth records the max recursion count
func (r *Registry) SetRecursiveDepth(count int) { _ = "STUB: not implemented"; return }

// GetRecursiveDepth returns the max recursion count
func (r *Registry) GetRecursiveDepth() int { _ = "STUB: not implemented"; return 0 }

// ReserveGoPackageAlias reserves the unique alias of go package.
// If succeeded, the alias will be never used for other packages in generated go files.
// If failed, the alias is already taken by another package, so you need to use another
// alias for the package in your go files.
func (r *Registry) ReserveGoPackageAlias(alias, pkgpath string) error {
	_ = "STUB: not implemented"
	return nil
}

// GetAllFQMNs returns a list of all FQMNs
func (r *Registry) GetAllFQMNs() []string { _ = "STUB: not implemented"; return nil }

// GetAllFQENs returns a list of all FQENs
func (r *Registry) GetAllFQENs() []string { _ = "STUB: not implemented"; return nil }

func (r *Registry) GetAllFQMethNs() []string { _ = "STUB: not implemented"; return nil }

// SetAllowDeleteBody controls whether http delete methods may have a
// body or fail loading if encountered.
func (r *Registry) SetAllowDeleteBody(allow bool) { _ = "STUB: not implemented"; return }

// SetAllowMerge controls whether generation one OpenAPI file out of multiple protos
func (r *Registry) SetAllowMerge(allow bool) { _ = "STUB: not implemented"; return }

// IsAllowMerge whether generation one OpenAPI file out of multiple protos
func (r *Registry) IsAllowMerge() bool { _ = "STUB: not implemented"; return false }

// SetMergeFileName controls the target OpenAPI file name out of multiple protos
func (r *Registry) SetMergeFileName(mergeFileName string) { _ = "STUB: not implemented"; return }

// SetIncludePackageInTags controls whether the package name defined in the `package` directive
// in the proto file can be prepended to the gRPC service name in the `Tags` field of every operation.
func (r *Registry) SetIncludePackageInTags(allow bool) { _ = "STUB: not implemented"; return }

// IsIncludePackageInTags checks whether the package name defined in the `package` directive
// in the proto file can be prepended to the gRPC service name in the `Tags` field of every operation.
func (r *Registry) IsIncludePackageInTags() bool { _ = "STUB: not implemented"; return false }

// GetRepeatedPathParamSeparator returns a rune specifying how
// path parameter repeated fields are separated.
func (r *Registry) GetRepeatedPathParamSeparator() rune { _ = "STUB: not implemented"; return 0 }

// GetRepeatedPathParamSeparatorName returns the name path parameter repeated
// fields repeatedFieldSeparator. I.e. 'csv', 'pipe', 'ssv' or 'tsv'
func (r *Registry) GetRepeatedPathParamSeparatorName() string { _ = "STUB: not implemented"; return "" }

// SetRepeatedPathParamSeparator sets how path parameter repeated fields are
// separated. Allowed names are 'csv', 'pipe', 'ssv' and 'tsv'.
func (r *Registry) SetRepeatedPathParamSeparator(name string) error {
	_ = "STUB: not implemented"
	return nil
}

// SetUseJSONNamesForFields sets useJSONNamesForFields
func (r *Registry) SetUseJSONNamesForFields(use bool) { _ = "STUB: not implemented"; return }

// GetUseJSONNamesForFields returns useJSONNamesForFields
func (r *Registry) GetUseJSONNamesForFields() bool { _ = "STUB: not implemented"; return false }

// SetUseFQNForOpenAPIName sets useFQNForOpenAPIName
// Deprecated: use SetOpenAPINamingStrategy instead.
func (r *Registry) SetUseFQNForOpenAPIName(use bool) { _ = "STUB: not implemented"; return }

// GetUseFQNForOpenAPIName returns useFQNForOpenAPIName
// Deprecated: Use GetOpenAPINamingStrategy().
func (r *Registry) GetUseFQNForOpenAPIName() bool { _ = "STUB: not implemented"; return false }

// GetMergeFileName return the target merge OpenAPI file name
func (r *Registry) GetMergeFileName() string { _ = "STUB: not implemented"; return "" }

// SetOpenAPINamingStrategy sets the naming strategy to be used.
func (r *Registry) SetOpenAPINamingStrategy(strategy string) { _ = "STUB: not implemented"; return }

// GetOpenAPINamingStrategy retrieves the naming strategy that is in use.
func (r *Registry) GetOpenAPINamingStrategy() string { _ = "STUB: not implemented"; return "" }

// SetUseGoTemplate sets useGoTemplate
func (r *Registry) SetUseGoTemplate(use bool) { _ = "STUB: not implemented"; return }

// GetUseGoTemplate returns useGoTemplate
func (r *Registry) GetUseGoTemplate() bool { _ = "STUB: not implemented"; return false }

func (r *Registry) SetGoTemplateArgs(kvs []string) { _ = "STUB: not implemented"; return }

func (r *Registry) GetGoTemplateArgs() map[string]string { _ = "STUB: not implemented"; return nil }

// SetIgnoreComments sets ignoreComments
func (r *Registry) SetIgnoreComments(ignore bool) { _ = "STUB: not implemented"; return }

// GetIgnoreComments returns ignoreComments
func (r *Registry) GetIgnoreComments() bool { _ = "STUB: not implemented"; return false }

// SetRemoveInternalComments sets removeInternalComments
func (r *Registry) SetRemoveInternalComments(remove bool) { _ = "STUB: not implemented"; return }

// GetRemoveInternalComments returns removeInternalComments
func (r *Registry) GetRemoveInternalComments() bool { _ = "STUB: not implemented"; return false }

// SetEnumsAsInts set enumsAsInts
func (r *Registry) SetEnumsAsInts(enumsAsInts bool) { _ = "STUB: not implemented"; return }

// GetEnumsAsInts returns enumsAsInts
func (r *Registry) GetEnumsAsInts() bool { _ = "STUB: not implemented"; return false }

// SetOmitEnumDefaultValue sets omitEnumDefaultValue
func (r *Registry) SetOmitEnumDefaultValue(omit bool) { _ = "STUB: not implemented"; return }

// GetOmitEnumDefaultValue returns omitEnumDefaultValue
func (r *Registry) GetOmitEnumDefaultValue() bool { _ = "STUB: not implemented"; return false }

// SetVisibilityRestrictionSelectors sets the visibility restriction selectors.
func (r *Registry) SetVisibilityRestrictionSelectors(selectors []string) {
	_ = "STUB: not implemented"
	return
}

// GetVisibilityRestrictionSelectors retrieves the visibility restriction selectors.
func (r *Registry) GetVisibilityRestrictionSelectors() map[string]bool {
	_ = "STUB: not implemented"
	return nil
}

// SetDisableDefaultErrors sets disableDefaultErrors
func (r *Registry) SetDisableDefaultErrors(use bool) { _ = "STUB: not implemented"; return }

// GetDisableDefaultErrors returns disableDefaultErrors
func (r *Registry) GetDisableDefaultErrors() bool { _ = "STUB: not implemented"; return false }

// SetSimpleOperationIDs sets simpleOperationIDs
func (r *Registry) SetSimpleOperationIDs(use bool) { _ = "STUB: not implemented"; return }

// GetSimpleOperationIDs returns simpleOperationIDs
func (r *Registry) GetSimpleOperationIDs() bool { _ = "STUB: not implemented"; return false }

// SetWarnOnUnboundMethods sets warnOnUnboundMethods
func (r *Registry) SetWarnOnUnboundMethods(warn bool) { _ = "STUB: not implemented"; return }

// SetGenerateUnboundMethods sets generateUnboundMethods
func (r *Registry) SetGenerateUnboundMethods(generate bool) { _ = "STUB: not implemented"; return }

// SetOmitPackageDoc controls whether the generated code contains a package comment (if set to false, it will contain one)
func (r *Registry) SetOmitPackageDoc(omit bool) { _ = "STUB: not implemented"; return }

// GetOmitPackageDoc returns whether a package comment will be omitted from the generated code
func (r *Registry) GetOmitPackageDoc() bool { _ = "STUB: not implemented"; return false }

// SetProto3OptionalNullable set proto3OptionalNullable
func (r *Registry) SetProto3OptionalNullable(proto3OptionalNullable bool) {
	_ = "STUB: not implemented"
	return
}

// GetProto3OptionalNullable returns proto3OptionalNullable
func (r *Registry) GetProto3OptionalNullable() bool { _ = "STUB: not implemented"; return false }

// RegisterOpenAPIOptions registers OpenAPI options
func (r *Registry) RegisterOpenAPIOptions(opts *openapiconfig.OpenAPIOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// build map of all registered methods

// build map of all registered fields

// GetOpenAPIFileOption returns a registered OpenAPI option for a file
func (r *Registry) GetOpenAPIFileOption(file string) (*options.Swagger, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetOpenAPIMethodOption returns a registered OpenAPI option for a method
func (r *Registry) GetOpenAPIMethodOption(qualifiedMethod string) (*options.Operation, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetOpenAPIMessageOption returns a registered OpenAPI option for a message
func (r *Registry) GetOpenAPIMessageOption(qualifiedMessage string) (*options.Schema, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetOpenAPIServiceOption returns a registered OpenAPI option for a service
func (r *Registry) GetOpenAPIServiceOption(qualifiedService string) (*options.Tag, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetOpenAPIFieldOption returns a registered OpenAPI option for a field
func (r *Registry) GetOpenAPIFieldOption(qualifiedField string) (*options.JSONSchema, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (r *Registry) FieldName(f *Field) string { _ = "STUB: not implemented"; return "" }

func (r *Registry) CheckDuplicateAnnotation(httpMethod string, httpTemplate string, svc *Service) error {
	_ = "STUB: not implemented"
	return nil
}

// SetDisableServiceTags sets disableServiceTags
func (r *Registry) SetDisableServiceTags(use bool) { _ = "STUB: not implemented"; return }

// GetDisableServiceTags returns disableServiceTags
func (r *Registry) GetDisableServiceTags() bool { _ = "STUB: not implemented"; return false }

// SetDisableDefaultResponses sets disableDefaultResponses
func (r *Registry) SetDisableDefaultResponses(use bool) { _ = "STUB: not implemented"; return }

// GetDisableDefaultResponses returns disableDefaultResponses
func (r *Registry) GetDisableDefaultResponses() bool { _ = "STUB: not implemented"; return false }

// SetUseAllOfForRefs sets useAllOfForRefs
func (r *Registry) SetUseAllOfForRefs(use bool) { _ = "STUB: not implemented"; return }

// GetUseAllOfForRefs returns useAllOfForRefs
func (r *Registry) GetUseAllOfForRefs() bool { _ = "STUB: not implemented"; return false }

// SetOmitArrayItemTypeWhenRefSibling sets omitArrayItemTypeWhenRefSibling
func (r *Registry) SetOmitArrayItemTypeWhenRefSibling(omit bool) { _ = "STUB: not implemented"; return }

// GetOmitArrayItemTypeWhenRefSibling returns omitArrayItemTypeWhenRefSibling
func (r *Registry) GetOmitArrayItemTypeWhenRefSibling() bool {
	_ = "STUB: not implemented"
	return false
}

// SetAllowPatchFeature sets allowPatchFeature
func (r *Registry) SetAllowPatchFeature(allow bool) { _ = "STUB: not implemented"; return }

// GetAllowPatchFeature returns allowPatchFeature
func (r *Registry) GetAllowPatchFeature() bool { _ = "STUB: not implemented"; return false }

// SetPreserveRPCOrder sets preserveRPCOrder
func (r *Registry) SetPreserveRPCOrder(preserve bool) { _ = "STUB: not implemented"; return }

// IsPreserveRPCOrder returns preserveRPCOrder
func (r *Registry) IsPreserveRPCOrder() bool { _ = "STUB: not implemented"; return false }

// SetEnableRpcDeprecation sets enableRpcDeprecation
func (r *Registry) SetEnableRpcDeprecation(enable bool) { _ = "STUB: not implemented"; return }

// GetEnableRpcDeprecation returns enableRpcDeprecation
func (r *Registry) GetEnableRpcDeprecation() bool { _ = "STUB: not implemented"; return false }

// SetEnableFieldDeprecation sets enableFieldDeprecation
func (r *Registry) SetEnableFieldDeprecation(enable bool) { _ = "STUB: not implemented"; return }

// GetEnableFieldDeprecation returns enableFieldDeprecation
func (r *Registry) GetEnableFieldDeprecation() bool { _ = "STUB: not implemented"; return false }

func (r *Registry) SetExpandSlashedPathPatterns(expandSlashedPathPatterns bool) {
	_ = "STUB: not implemented"
	return
}

func (r *Registry) GetExpandSlashedPathPatterns() bool { _ = "STUB: not implemented"; return false }

func (r *Registry) SetGenerateXGoType(generateXGoType bool) { _ = "STUB: not implemented"; return }

func (r *Registry) GetGenerateXGoType() bool { _ = "STUB: not implemented"; return false }
