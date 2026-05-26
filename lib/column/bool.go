package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

type Bool struct {
	col  proto.ColBool
	name string
}

func (col *Bool) Reset() { _ = "STUB: not implemented"; return }

func (col *Bool) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Bool) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Bool) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Bool) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Bool) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *Bool) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *Bool) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Bool) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *Bool) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Bool) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Bool) row(i int) bool { _ = "STUB: not implemented"; return false }

var _ Interface = (*Bool)(nil)
