package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
	"github.com/paulmach/orb"
)

type MultiLineString struct {
	set  *Array
	name string
}

func (col *MultiLineString) Reset() { _ = "STUB: not implemented"; return }

func (col *MultiLineString) Name() string { _ = "STUB: not implemented"; return "" }

func (col *MultiLineString) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *MultiLineString) ScanType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (col *MultiLineString) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *MultiLineString) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *MultiLineString) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *MultiLineString) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *MultiLineString) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *MultiLineString) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *MultiLineString) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *MultiLineString) row(i int) orb.MultiLineString {
	_ = "STUB: not implemented"
	return *new(orb.MultiLineString)
}

var _ Interface = (*MultiLineString)(nil)
