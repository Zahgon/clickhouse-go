package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

type Tuple struct {
	chType  Type
	columns []Interface
	name    string
	isNamed bool           // true if all columns are named
	index   map[string]int // map from col name to offset in columns
}

func (col *Tuple) Reset() { _ = "STUB: not implemented"; return }

func (col *Tuple) Name() string { _ = "STUB: not implemented"; return "" }

type namedCol struct {
	name    string
	colType Type
}

func (col *Tuple) parse(t Type, sc *ServerContext) (_ Interface, err error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

func (col *Tuple) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col Tuple) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Tuple) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Tuple) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

// if this happens we have an unexplained problem

func setJSONFieldValue(field reflect.Value, value reflect.Value) error {
	_ = "STUB: not implemented"
	return nil
}

// check if our target implements sql.Scanner

func getStructFieldValue(field reflect.Value, name string) (reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

func unescapeColName(colName string) string { _ = "STUB: not implemented"; return "" }

func (col *Tuple) scanMap(targetMap reflect.Value, row int) error {
	_ = "STUB: not implemented"
	return nil
}

// get a typed map

// catches any - Note this swallows custom interfaces to which maps couldn't conform

// this wont work if targetMap is a map[string][]any and we try to set a typed slice

func (col *Tuple) scanStruct(targetStruct reflect.Value, row int) error {
	_ = "STUB: not implemented"
	return nil
}

// the column may be serialized using a different name due to a struct "targetStruct" tag

// test if map

// catches []any -Note this swallows custom interfaces to which maps couldn't conform

func (col *Tuple) scanSlice(targetType reflect.Type, row int) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

func (col *Tuple) scan(targetType reflect.Type, row int) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

//tuples can be scanned into slices - specifically default for unnamed tuples

// catches any -Note this swallows custom interfaces to which maps couldn't conform

func (col *Tuple) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *Tuple) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Tuple) AppendRow(v any) error {
	_ = "STUB: not implemented"
	// allows support of tuples where map or slice is typed and NOT any. Will fail if tuple isn't consistent
	return nil
}

// can't interface - likely not exported so ignore the field

// can't interface - likely not exported so ignore the field

func (col *Tuple) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Tuple) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Tuple) ReadStatePrefix(reader *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Tuple) WriteStatePrefix(buffer *proto.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	_ Interface           = (*Tuple)(nil)
	_ CustomSerialization = (*Tuple)(nil)
)

func getStructFieldName(field reflect.StructField) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// not a standard but we allow - to omit fields

// Some JSON tags contain omitempty after a comma but we don't want those in our field name.

// support ch tag as well as this is used elsewhere

// ensures numeric keys and ` are escaped properly
func getMapFieldName(name string) string { _ = "STUB: not implemented"; return "" }
