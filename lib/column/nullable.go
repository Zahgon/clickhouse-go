package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

type Nullable struct {
	base     Interface
	nulls    proto.ColUInt8
	enable   bool
	scanType reflect.Type
	name     string
}

func (col *Nullable) Reset() { _ = "STUB: not implemented"; return }

func (col *Nullable) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Nullable) parse(t Type, sc *ServerContext) (_ *Nullable, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Nullable) Base() Interface { _ = "STUB: not implemented"; return *new(Interface) }

func (col *Nullable) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Nullable) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Nullable) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Nullable) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *Nullable) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *Nullable) Append(v any) ([]uint8, error) { _ = "STUB: not implemented"; return nil, nil }

func (col *Nullable) AppendRow(v any) error {
	_ = "STUB: not implemented"
	// Might receive double pointers like **String, because of how Nullable columns are read
	// Unpack because we can't write double pointers
	return nil
}

// used to detect sql.Null* types

func (col *Nullable) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Nullable) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Nullable) ReadStatePrefix(reader *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Nullable) WriteStatePrefix(buffer *proto.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

var _ Interface = (*Nullable)(nil)
