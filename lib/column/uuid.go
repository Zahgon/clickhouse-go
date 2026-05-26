package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"

	"github.com/google/uuid"
)

type UUID struct {
	col  proto.ColUUID
	name string
}

func (col *UUID) Reset() { _ = "STUB: not implemented"; return }

func (col *UUID) Name() string { _ = "STUB: not implemented"; return "" }

func (col *UUID) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *UUID) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *UUID) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *UUID) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *UUID) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *UUID) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *UUID) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *UUID) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *UUID) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *UUID) row(i int) (uuid uuid.UUID) { _ = "STUB: not implemented"; return *new(uuid.UUID) }

var _ Interface = (*UUID)(nil)
