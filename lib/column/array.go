package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

var scanTypeAny = reflect.TypeOf((*any)(nil)).Elem()

type offset struct {
	values   UInt64
	scanType reflect.Type
}

type Array struct {
	depth    int
	chType   Type
	values   Interface
	offsets  []*offset
	scanType reflect.Type
	name     string
}

func (col *Array) Reset() { _ = "STUB: not implemented"; return }

func (col *Array) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Array) parse(t Type, sc *ServerContext) (_ *Array, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Array) Base() Interface { _ = "STUB: not implemented"; return *new(Interface) }

func (col *Array) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Array) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Array) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Array) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *Array) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Array) AppendRow(v any) error {
	_ = "STUB: not implemented"

	// try to use reflection-free method.
	return nil
}

func (col *Array) appendRowDefault(v any) error { _ = "STUB: not implemented"; return nil }

func appendRowPlain[T any](col *Array, arr []T) error { _ = "STUB: not implemented"; return nil }

func appendNullableRowPlain[T any](col *Array, arr []*T) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Array) append(elem reflect.Value, level int) error {
	_ = "STUB: not implemented"
	return nil
}

// allows to traverse pointers to slices and slices cast to `any`

// reflect.Value.Len() & reflect.Value.Index() is called in `append` method which is only valid for
// Slice, Array and String that make sense here.

func (col *Array) appendOffset(level int, offset uint64) { _ = "STUB: not implemented"; return }

func (col *Array) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Array) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Array) ReadStatePrefix(reader *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Array) WriteStatePrefix(buffer *proto.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Array) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *Array) scan(sliceType reflect.Type, row int) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (col *Array) scanSlice(sliceType reflect.Type, row int, level int) (reflect.Value, error) {
	_ = "STUB: not implemented"
	// We could try and set - if it exceeds just return immediately
	return *new(reflect.Value), nil
}

//Array(Nested

//Array(Array

// Array(Tuple possible outside JSON object cases e.g. if the user defines a  Array(Array( Tuple(String, Int64) ))

func (col *Array) scanSliceOfObjects(sliceType reflect.Type, row int) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// catches any - Note this swallows custom interfaces to which maps couldn't conform

// make a slice of the right type - we need this to be a slice of a type capable of taking an object as nested

// tuples can be read as arrays

// catches []any - Note this swallows custom interfaces to which maps could never conform

// the following 2 functions can probably be refactored - the share alot of common code for structs and maps
func (col *Array) scanSliceOfMaps(sliceType reflect.Type, row int) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// Array(Tuple so depth 1 for JSON

func (col *Array) scanSliceOfStructs(sliceType reflect.Type, row int) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// Array(Tuple so depth 1 for JSON

// create a slice of the type from the sliceType - if this might be any as its driven by the target datastructure

var (
	_ Interface           = (*Array)(nil)
	_ CustomSerialization = (*Array)(nil)
)
