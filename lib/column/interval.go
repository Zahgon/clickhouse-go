package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

type Interval struct {
	chType Type
	name   string
	col    proto.ColInt64
}

func (col *Interval) Reset() { _ = "STUB: not implemented"; return }

func (col *Interval) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Interval) parse(t Type) (Interface, error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

func (col *Interval) Type() Type              { _ = "STUB: not implemented"; return *new(Type) }
func (col *Interval) ScanType() reflect.Type  { _ = "STUB: not implemented"; return *new(reflect.Type) }
func (col *Interval) Rows() int               { _ = "STUB: not implemented"; return 0 }
func (col *Interval) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *Interval) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (Interval) Append(any) ([]uint8, error) { _ = "STUB: not implemented"; return nil, nil }

func (Interval) AppendRow(any) error { _ = "STUB: not implemented"; return nil }

func (col *Interval) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (Interval) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Interval) row(i int) string { _ = "STUB: not implemented"; return "" }

var _ Interface = (*Interval)(nil)
