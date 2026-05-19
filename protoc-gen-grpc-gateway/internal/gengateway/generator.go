package gengateway

import (
	"errors"

	"github.com/grpc-ecosystem/grpc-gateway/v2/internal/descriptor"
	gen "github.com/grpc-ecosystem/grpc-gateway/v2/internal/generator"
)

var errNoTargetService = errors.New("no target service defined in the file")

type generator struct {
	reg                *descriptor.Registry
	baseImports        []descriptor.GoPackage
	useRequestContext  bool
	registerFuncSuffix string
	allowPatchFeature  bool
	standalone         bool
	useOpaqueAPI       bool
}

// New returns a new generator which generates grpc gateway files.
func New(reg *descriptor.Registry, useRequestContext bool, registerFuncSuffix string,
	allowPatchFeature, standalone bool, useOpaqueAPI bool) gen.Generator {
	_ = "STUB: not implemented"
	return *new(gen.Generator)
}

func (g *generator) Generate(targets []*descriptor.File) ([]*descriptor.ResponseFile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (g *generator) generate(file *descriptor.File) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// addEnumPathParamImports handles adding import of enum path parameter go packages
func (g *generator) addEnumPathParamImports(file *descriptor.File, m *descriptor.Method, pkgSeen map[string]bool) []descriptor.GoPackage {
	_ = "STUB: not implemented"
	return nil
}

// addBodyFieldImports ensures nested body message types pull in their Go packages.
func (g *generator) addBodyFieldImports(
	file *descriptor.File,
	m *descriptor.Method,
	pkgSeen map[string]bool,
) []descriptor.GoPackage {
	_ = "STUB: not implemented"
	return nil
}
