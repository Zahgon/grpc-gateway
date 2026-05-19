package runtime

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// JSONPb is a Marshaler which marshals/unmarshals into/from JSON
// with the "google.golang.org/protobuf/encoding/protojson" marshaler.
// It supports the full functionality of protobuf unlike JSONBuiltin.
//
// The NewDecoder method returns a DecoderWrapper, so the underlying
// *json.Decoder methods can be used.
type JSONPb struct {
	protojson.MarshalOptions
	protojson.UnmarshalOptions
}

// ContentType always returns "application/json".
func (*JSONPb) ContentType(_ interface{}) string { _ = "STUB: not implemented"; return "" }

// Marshal marshals "v" into JSON.
func (j *JSONPb) Marshal(v interface{}) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (j *JSONPb) marshalTo(w io.Writer, v interface{}) error { _ = "STUB: not implemented"; return nil }

var (
	// protoMessageType is stored to prevent constant lookup of the same type at runtime.
	protoMessageType = reflect.TypeFor[proto.Message]()
)

// marshalNonProto marshals a non-message field of a protobuf message.
// This function does not correctly marshal arbitrary data structures into JSON,
// it is only capable of marshaling non-message field values of protobuf,
// i.e. primitive types, enums; pointers to primitives or enums; maps from
// integer/string types to primitives/enums/pointers to messages.
func (j *JSONPb) marshalNonProtoField(v interface{}) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Unmarshal unmarshals JSON "data" into "v"
func (j *JSONPb) Unmarshal(data []byte, v interface{}) error { _ = "STUB: not implemented"; return nil }

// NewDecoder returns a Decoder which reads JSON stream from "r".
func (j *JSONPb) NewDecoder(r io.Reader) Decoder { _ = "STUB: not implemented"; return *new(Decoder) }

// DecoderWrapper is a wrapper around a *json.Decoder that adds
// support for protos to the Decode method.
type DecoderWrapper struct {
	*json.Decoder
	protojson.UnmarshalOptions
}

// Decode wraps the embedded decoder's Decode method to support
// protos using a jsonpb.Unmarshaler.
func (d DecoderWrapper) Decode(v interface{}) error { _ = "STUB: not implemented"; return nil }

// NewEncoder returns an Encoder which writes JSON stream into "w".
func (j *JSONPb) NewEncoder(w io.Writer) Encoder { _ = "STUB: not implemented"; return *new(Encoder) }

// mimic json.Encoder by adding a newline (makes output
// easier to read when it contains multiple encoded items)

func unmarshalJSONPb(data []byte, unmarshaler protojson.UnmarshalOptions, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func decodeJSONPb(d *json.Decoder, unmarshaler protojson.UnmarshalOptions, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Decode into bytes for marshalling

func decodeNonProtoField(d *json.Decoder, unmarshaler protojson.UnmarshalOptions, v interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// Decode into bytes for marshalling

// TODO(yugui) Should use proto.StructProperties?

type protoEnum interface {
	fmt.Stringer
	EnumDescriptor() ([]byte, []int)
}

var typeProtoEnum = reflect.TypeFor[protoEnum]()

var typeProtoMessage = reflect.TypeFor[proto.Message]()

// Delimiter for newline encoded JSON streams.
func (j *JSONPb) Delimiter() []byte { _ = "STUB: not implemented"; return nil }

var (
	convFromType = map[reflect.Kind]reflect.Value{
		reflect.String:  reflect.ValueOf(String),
		reflect.Bool:    reflect.ValueOf(Bool),
		reflect.Float64: reflect.ValueOf(Float64),
		reflect.Float32: reflect.ValueOf(Float32),
		reflect.Int64:   reflect.ValueOf(Int64),
		reflect.Int32:   reflect.ValueOf(Int32),
		reflect.Uint64:  reflect.ValueOf(Uint64),
		reflect.Uint32:  reflect.ValueOf(Uint32),
		reflect.Slice:   reflect.ValueOf(Bytes),
	}
)
