package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

type Enum16 struct {
	iv     map[string]proto.Enum16
	vi     map[proto.Enum16]string
	chType Type
	col    proto.ColEnum16
	name   string

	continuous bool
	minEnum    int16
	maxEnum    int16
}

func (col *Enum16) Reset() { _ = "STUB: not implemented"; return }

func (col *Enum16) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Enum16) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Enum16) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Enum16) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Enum16) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *Enum16) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *Enum16) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Enum16) AppendRow(elem any) error { _ = "STUB: not implemented"; return nil }

func (col *Enum16) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Enum16) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

var _ Interface = (*Enum16)(nil)
