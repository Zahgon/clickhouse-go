package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
	"github.com/paulmach/orb"
)

type LineString struct {
	set  *Array
	name string
}

func (col *LineString) Reset() { _ = "STUB: not implemented"; return }

func (col *LineString) Name() string { _ = "STUB: not implemented"; return "" }

func (col *LineString) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *LineString) ScanType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (col *LineString) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *LineString) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *LineString) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *LineString) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *LineString) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *LineString) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *LineString) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *LineString) row(i int) orb.LineString {
	_ = "STUB: not implemented"
	return *new(orb.LineString)
}

var _ Interface = (*LineString)(nil)
