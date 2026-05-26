package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

// SharedVariant deprecated. Use Dynamic/JSON serialization version 3.
type SharedVariant struct {
	name       string
	stringData String
}

func (c *SharedVariant) Name() string { _ = "STUB: not implemented"; return "" }

func (c *SharedVariant) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (c *SharedVariant) Rows() int { _ = "STUB: not implemented"; return 0 }

func (c *SharedVariant) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (c *SharedVariant) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (c *SharedVariant) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *SharedVariant) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (c *SharedVariant) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (c *SharedVariant) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *SharedVariant) ScanType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (c *SharedVariant) Reset() { _ = "STUB: not implemented"; return }
