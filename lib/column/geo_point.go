package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"

	"github.com/paulmach/orb"
)

type Point struct {
	name string
	col  proto.ColPoint
}

func (col *Point) Reset() { _ = "STUB: not implemented"; return }

func (col *Point) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Point) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Point) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Point) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Point) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *Point) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *Point) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Point) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *Point) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Point) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Point) row(i int) orb.Point { _ = "STUB: not implemented"; return *new(orb.Point) }

var _ Interface = (*Point)(nil)
