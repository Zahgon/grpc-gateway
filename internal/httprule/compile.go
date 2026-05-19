package httprule

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
)

const (
	opcodeVersion = 1
)

// Template is a compiled representation of path templates.
type Template struct {
	// Version is the version number of the format.
	Version int
	// OpCodes is a sequence of operations.
	OpCodes []int
	// Pool is a constant pool
	Pool []string
	// Verb is a VERB part in the template.
	Verb string
	// Fields is a list of field paths bound in this template.
	Fields []string
	// Original template (example: /v1/a_bit_of_everything)
	Template string
}

// Compiler compiles utilities representation of path templates into marshallable operations.
// They can be unmarshalled by runtime.NewPattern.
type Compiler interface {
	Compile() Template
}

type op struct {
	// code is the opcode of the operation
	code utilities.OpCode

	// str is a string operand of the code.
	// num is ignored if str is not empty.
	str string

	// num is a numeric operand of the code.
	num int
}

func (w wildcard) compile() []op { _ = "STUB: not implemented"; return nil }

func (w deepWildcard) compile() []op { _ = "STUB: not implemented"; return nil }

func (l literal) compile() []op { _ = "STUB: not implemented"; return nil }

func (v variable) compile() []op { _ = "STUB: not implemented"; return nil }

func (t template) Compile() Template { _ = "STUB: not implemented"; return *new(Template) }

// eof segment literal represents the "/" path pattern
