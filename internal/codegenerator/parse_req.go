package codegenerator

import (
	"io"

	"google.golang.org/protobuf/types/pluginpb"
)

// ParseRequest parses a code generator request from a proto Message.
func ParseRequest(r io.Reader) (*pluginpb.CodeGeneratorRequest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
