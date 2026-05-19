package runtime

import (
	"io"
)

// ProtoMarshaller is a Marshaller which marshals/unmarshals into/from serialize proto bytes
type ProtoMarshaller struct{}

// ContentType always returns "application/octet-stream".
func (*ProtoMarshaller) ContentType(_ interface{}) string { _ = "STUB: not implemented"; return "" }

// Marshal marshals "value" into Proto
func (*ProtoMarshaller) Marshal(value interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unmarshal unmarshals proto "data" into "value"
func (*ProtoMarshaller) Unmarshal(data []byte, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// NewDecoder returns a Decoder which reads proto stream from "reader".
func (marshaller *ProtoMarshaller) NewDecoder(reader io.Reader) Decoder {
	_ = "STUB: not implemented"
	return *new(Decoder)
}

// NewEncoder returns an Encoder which writes proto stream into "writer".
func (marshaller *ProtoMarshaller) NewEncoder(writer io.Writer) Encoder {
	_ = "STUB: not implemented"
	return *new(Encoder)
}
