package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

type FixedString struct {
	name string
	col  proto.ColFixedStr
}

func (col *FixedString) Reset() { _ = "STUB: not implemented"; return }

func (col *FixedString) Name() string { _ = "STUB: not implemented"; return "" }

func (col *FixedString) parse(t Type) (*FixedString, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *FixedString) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *FixedString) ScanType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (col *FixedString) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *FixedString) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *FixedString) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

// handle for *[n]byte

// safeAppendRow appends the value to the underlying column with a length check.
// This re-implements the logic from ch-go but without the panic.
// It also fills unused space with zeros.
func (col *FixedString) safeAppendRow(v []byte) error { _ = "STUB: not implemented"; return nil }

// If unset, use first value's length for the string size

// Fill the unused space of the fixed string with zeros

func (col *FixedString) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handle for [][n]byte

func (col *FixedString) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *FixedString) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *FixedString) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *FixedString) row(i int) string { _ = "STUB: not implemented"; return "" }

func (col *FixedString) rowBytes(i int) []byte { _ = "STUB: not implemented"; return nil }

var _ Interface = (*FixedString)(nil)
