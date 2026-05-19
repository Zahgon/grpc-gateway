package genopenapi

import (
	"io"
)

type Format string

const (
	FormatJSON Format = "json"
	FormatYAML Format = "yaml"
)

type ContentEncoder interface {
	Encode(v interface{}) (err error)
}

func (f Format) Validate() error { _ = "STUB: not implemented"; return nil }

func (f Format) NewEncoder(w io.Writer) (ContentEncoder, error) {
	_ = "STUB: not implemented"
	return *new(ContentEncoder), nil
}
