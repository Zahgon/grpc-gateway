package httprule

import (
	"fmt"
)

type template struct {
	segments []segment
	verb     string
	template string
}

type segment interface {
	fmt.Stringer
	compile() (ops []op)
}

type wildcard struct{}

type deepWildcard struct{}

type literal string

type variable struct {
	path     string
	segments []segment
}

func (wildcard) String() string { _ = "STUB: not implemented"; return "" }

func (deepWildcard) String() string { _ = "STUB: not implemented"; return "" }

func (l literal) String() string { _ = "STUB: not implemented"; return "" }

func (v variable) String() string { _ = "STUB: not implemented"; return "" }

func (t template) String() string { _ = "STUB: not implemented"; return "" }
