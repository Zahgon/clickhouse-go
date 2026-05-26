package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

type String struct {
	name string
	col  proto.ColStr
}

func (col *String) Reset() { _ = "STUB: not implemented"; return }

func (col String) Name() string { _ = "STUB: not implemented"; return "" }

func (String) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (String) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *String) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *String) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *String) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *String) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *String) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *String) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *String) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

var _ Interface = (*String)(nil)
