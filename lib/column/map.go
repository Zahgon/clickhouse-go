package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

// https://github.com/ClickHouse/ClickHouse/blob/master/src/Columns/ColumnMap.cpp
type Map struct {
	keys     Interface
	values   Interface
	chType   Type
	offsets  Int64
	scanType reflect.Type
	name     string
}

type OrderedMap interface {
	Get(key any) (any, bool)
	Put(key any, value any)
	Keys() <-chan any
}

type MapIterator interface {
	Next() bool
	Key() any
	Value() any
}

type IterableOrderedMap interface {
	Put(key any, value any)
	Iterator() MapIterator
}

func (col *Map) Reset() { _ = "STUB: not implemented"; return }

func (col *Map) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Map) parse(t Type, sc *ServerContext) (_ Interface, err error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

func (col *Map) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Map) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Map) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Map) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *Map) ScanRow(dest any, i int) error { _ = "STUB: not implemented"; return nil }

func (col *Map) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Map) AppendRow(v any) error {
	_ = "STUB: not implemented"

	// NOTE: successful Map.parse() make sure we have
	// valid col.scanType
	return nil
}

func (col *Map) Decode(reader *proto.Reader, rows int) error { _ = "STUB: not implemented"; return nil }

func (col *Map) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Map) ReadStatePrefix(reader *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func (col *Map) WriteStatePrefix(encoder *proto.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Map) row(n int) reflect.Value { _ = "STUB: not implemented"; return *new(reflect.Value) }

// Convert interface{} nil to typed nil (such as nil *string) to preserve map element
// https://github.com/ClickHouse/clickhouse-go/issues/1515

func (col *Map) orderedRow(n int) ([]any, []any) { _ = "STUB: not implemented"; return nil, nil }

var (
	_ Interface           = (*Map)(nil)
	_ CustomSerialization = (*Map)(nil)
)
