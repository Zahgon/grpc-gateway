package runtime

import (
	"io"
)

// JSONBuiltin is a Marshaler which marshals/unmarshals into/from JSON
// with the standard "encoding/json" package of Golang.
// Although it is generally faster for simple proto messages than JSONPb,
// it does not support advanced features of protobuf, e.g. map, oneof, ....
//
// The NewEncoder and NewDecoder types return *json.Encoder and
// *json.Decoder respectively.
type JSONBuiltin struct{}

// ContentType always Returns "application/json".
func (*JSONBuiltin) ContentType(_ interface{}) string { _ = "STUB: not implemented"; return "" }

// Marshal marshals "v" into JSON
func (j *JSONBuiltin) Marshal(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// MarshalIndent is like Marshal but applies Indent to format the output
		nil
}

func (j *JSONBuiltin) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unmarshal unmarshals JSON data into "v".
func (j *JSONBuiltin) Unmarshal(data []byte, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// NewDecoder returns a Decoder which reads JSON stream from "r".
func (j *JSONBuiltin) NewDecoder(r io.Reader) Decoder {
	_ = "STUB: not implemented"
	return *new(Decoder)
}

// NewEncoder returns an Encoder which writes JSON stream into "w".
func (j *JSONBuiltin) NewEncoder(w io.Writer) Encoder {
	_ = "STUB: not implemented"
	return *new(Encoder)
}

// Delimiter for newline encoded JSON streams.
func (j *JSONBuiltin) Delimiter() []byte { _ = "STUB: not implemented"; return nil }
