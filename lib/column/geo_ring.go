package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"

	"github.com/paulmach/orb"
)

type Ring struct {
	set  *Array
	name string
}

func (col *Ring) Reset() { _ = "STUB: not implemented"; return }

func (col *Ring) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Ring) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Ring) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Ring) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Ring) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *Ring) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *Ring) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Ring) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *Ring) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Ring) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Ring) row(i int) orb.Ring { _ = "STUB: not implemented"; return *new(orb.Ring) }

var _ Interface = (*Ring)(nil)
