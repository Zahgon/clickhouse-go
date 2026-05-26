package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"

	"github.com/paulmach/orb"
)

type Polygon struct {
	set  *Array
	name string
}

func (col *Polygon) Reset() { _ = "STUB: not implemented"; return }

func (col *Polygon) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Polygon) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Polygon) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Polygon) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Polygon) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *Polygon) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *Polygon) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Polygon) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *Polygon) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Polygon) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Polygon) row(i int) orb.Polygon { _ = "STUB: not implemented"; return *new(orb.Polygon) }

var _ Interface = (*Polygon)(nil)
