package utilities

import (
	"io"
)

// IOReaderFactory takes in an io.Reader and returns a function that will allow you to create a new reader that begins
// at the start of the stream
func IOReaderFactory(r io.Reader) (func() io.Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
