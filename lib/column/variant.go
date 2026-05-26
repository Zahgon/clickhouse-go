package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

const SupportedVariantSerializationVersion = 0
const VariantNullDiscriminator uint8 = 255

type Variant struct {
	chType Type
	name   string

	discriminators []uint8
	offsets        []int

	columns         []Interface
	columnTypeIndex map[string]uint8
}

func (c *Variant) parse(t Type, sc *ServerContext) (_ *Variant, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Variant) addColumn(col Interface) { _ = "STUB: not implemented"; return }

func (c *Variant) appendDiscriminatorRow(d uint8) { _ = "STUB: not implemented"; return }

func (c *Variant) appendNullRow() { _ = "STUB: not implemented"; return }

func (c *Variant) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Variant) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (c *Variant) Rows() int { _ = "STUB: not implemented"; return 0 }

func (c *Variant) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (c *Variant) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (c *Variant) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Variant) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

// If preferred type wasn't provided, try each column

func (c *Variant) WriteStatePrefix(buffer *proto.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Variant) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (c *Variant) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (c *Variant) Reset() { _ = "STUB: not implemented"; return }

func (c *Variant) ReadStatePrefix(reader *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Variant) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}
