package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"

	"github.com/paulmach/orb"
)

type MultiPolygon struct {
	set  *Array
	name string
}

func (col *MultiPolygon) Reset() { _ = "STUB: not implemented"; return }

func (col *MultiPolygon) Name() string { _ = "STUB: not implemented"; return "" }

func (col *MultiPolygon) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *MultiPolygon) ScanType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (col *MultiPolygon) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *MultiPolygon) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *MultiPolygon) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *MultiPolygon) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *MultiPolygon) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *MultiPolygon) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *MultiPolygon) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *MultiPolygon) row(i int) orb.MultiPolygon {
	_ = "STUB: not implemented"
	return *new(orb.MultiPolygon)
}

var _ Interface = (*MultiPolygon)(nil)
