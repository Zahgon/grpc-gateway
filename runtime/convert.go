package runtime

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// String just returns the given string.
// It is just for compatibility to other types.
func String(val string) (string, error) {
	_ = "STUB: not implemented"

	// StringSlice converts 'val' where individual strings are separated by
	// 'sep' into a string slice.
	return "", nil
}

func StringSlice(val, sep string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Bool converts the given string representation of a boolean value into bool.
func Bool(val string) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// BoolSlice converts 'val' where individual booleans are separated by
		// 'sep' into a bool slice.
		nil
}

func BoolSlice(val, sep string) ([]bool, error) { _ = "STUB: not implemented"; return nil, nil }

// Float64 converts the given string representation into representation of a floating point number into float64.
func Float64(val string) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// Float64Slice converts 'val' where individual floating point numbers are separated by
// 'sep' into a float64 slice.
func Float64Slice(val, sep string) ([]float64, error) { _ = "STUB: not implemented"; return nil, nil }

// Float32 converts the given string representation of a floating point number into float32.
func Float32(val string) (float32, error) { _ = "STUB: not implemented"; return 0, nil }

// Float32Slice converts 'val' where individual floating point numbers are separated by
// 'sep' into a float32 slice.
func Float32Slice(val, sep string) ([]float32, error) { _ = "STUB: not implemented"; return nil, nil }

// Int64 converts the given string representation of an integer into int64.
func Int64(val string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// Int64Slice converts 'val' where individual integers are separated by
// 'sep' into an int64 slice.
func Int64Slice(val, sep string) ([]int64, error) { _ = "STUB: not implemented"; return nil, nil }

// Int32 converts the given string representation of an integer into int32.
func Int32(val string) (int32, error) { _ = "STUB: not implemented"; return 0, nil }

// Int32Slice converts 'val' where individual integers are separated by
// 'sep' into an int32 slice.
func Int32Slice(val, sep string) ([]int32, error) { _ = "STUB: not implemented"; return nil, nil }

// Uint64 converts the given string representation of an integer into uint64.
func Uint64(val string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// Uint64Slice converts 'val' where individual integers are separated by
// 'sep' into a uint64 slice.
func Uint64Slice(val, sep string) ([]uint64, error) { _ = "STUB: not implemented"; return nil, nil }

// Uint32 converts the given string representation of an integer into uint32.
func Uint32(val string) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// Uint32Slice converts 'val' where individual integers are separated by
// 'sep' into a uint32 slice.
func Uint32Slice(val, sep string) ([]uint32, error) { _ = "STUB: not implemented"; return nil, nil }

// Bytes converts the given string representation of a byte sequence into a slice of bytes
// A bytes sequence is encoded in URL-safe base64 without padding
func Bytes(val string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// BytesSlice converts 'val' where individual bytes sequences, encoded in URL-safe
// base64 without padding, are separated by 'sep' into a slice of byte slices.
func BytesSlice(val, sep string) ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Timestamp converts the given RFC3339 formatted string into a timestamp.Timestamp.
func Timestamp(val string) (*timestamppb.Timestamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Duration converts the given string into a timestamp.Duration.
func Duration(val string) (*durationpb.Duration, error) { _ = "STUB: not implemented"; return nil, nil }

// Enum converts the given string into an int32 that should be type casted into the
// correct enum proto type.
func Enum(val string, enumValMap map[string]int32) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// EnumSlice converts 'val' where individual enums are separated by 'sep'
// into a int32 slice. Each individual int32 should be type casted into the
// correct enum proto type.
func EnumSlice(val, sep string, enumValMap map[string]int32) ([]int32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Support for google.protobuf.wrappers on top of primitive types

// StringValue well-known type support as wrapper around string type
func StringValue(val string) (*wrapperspb.StringValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FloatValue well-known type support as wrapper around float32 type
func FloatValue(val string) (*wrapperspb.FloatValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DoubleValue well-known type support as wrapper around float64 type
func DoubleValue(val string) (*wrapperspb.DoubleValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BoolValue well-known type support as wrapper around bool type
func BoolValue(val string) (*wrapperspb.BoolValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Int32Value well-known type support as wrapper around int32 type
func Int32Value(val string) (*wrapperspb.Int32Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UInt32Value well-known type support as wrapper around uint32 type
func UInt32Value(val string) (*wrapperspb.UInt32Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Int64Value well-known type support as wrapper around int64 type
func Int64Value(val string) (*wrapperspb.Int64Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UInt64Value well-known type support as wrapper around uint64 type
func UInt64Value(val string) (*wrapperspb.UInt64Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BytesValue well-known type support as wrapper around bytes[] type
func BytesValue(val string) (*wrapperspb.BytesValue, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
