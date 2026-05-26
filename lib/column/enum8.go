package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

type Enum8 struct {
	iv     map[string]proto.Enum8
	vi     map[proto.Enum8]string
	chType Type
	name   string
	col    proto.ColEnum8

	// Encoding of the enums that have been specified by the user.
	// Using this when appending rows, to validate the enum is valud.
	enumValuesBitset [4]uint64
}

func (col *Enum8) Reset() { _ = "STUB: not implemented"; return }

func (col *Enum8) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Enum8) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Enum8) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Enum8) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Enum8) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *Enum8) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *Enum8) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Enum8) AppendRow(elem any) error { _ = "STUB: not implemented"; return nil }

// Check if the enum value is defined

// Check if the enum value is defined

func (col *Enum8) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Enum8) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

var _ Interface = (*Enum8)(nil)
