package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"

	"github.com/shopspring/decimal"
)

type Decimal struct {
	chType    Type
	scale     int
	precision int
	name      string
	col       proto.Column
}

func (col *Decimal) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Decimal) Reset() { _ = "STUB: not implemented"; return }

func (col *Decimal) parse(t Type) (_ *Decimal, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Decimal) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Decimal) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Decimal) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Decimal) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *Decimal) row(i int) *decimal.Decimal { _ = "STUB: not implemented"; return nil }

func (col *Decimal) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *Decimal) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Decimal) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *Decimal) append(v *decimal.Decimal) { _ = "STUB: not implemented"; return }

func (col *Decimal) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Decimal) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Decimal) Scale() int64 { _ = "STUB: not implemented"; return 0 }

func (col *Decimal) Precision() int64 { _ = "STUB: not implemented"; return 0 }

var _ Interface = (*Decimal)(nil)
