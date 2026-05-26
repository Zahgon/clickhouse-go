package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

const DynamicSerializationVersion = 3
const DynamicDeprecatedSerializationVersion = 1
const DynamicNullDiscriminator = -1 // The Null index changes as data is being built, use -1 as placeholder for writes.
const DefaultMaxDynamicTypes = 32

func supportsFlatDynamicJSON(sc *ServerContext) bool {
	_ = "STUB: not implemented"
	// Any CH version more than 25.6
	return false
}

type Dynamic struct {
	chType Type
	sc     *ServerContext
	name   string

	serializationVersion uint64

	totalTypes     int // Null is last type index + 1, so this doubles as the Null type index for reads.
	discriminators []int
	offsets        []int

	columns           []Interface
	columnIndexByType map[string]int

	deprecated deprecatedDynamic
}

func (c *Dynamic) parse(t Type, sc *ServerContext) (_ *Dynamic, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SharedVariant is special, and does not count against totalTypes

// Reset to 0 after adding SharedVariant

func (c *Dynamic) addColumn(col Interface) int { _ = "STUB: not implemented"; return 0 }

func (c *Dynamic) Name() string { _ = "STUB: not implemented"; return "" }

func (c *Dynamic) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (c *Dynamic) Rows() int { _ = "STUB: not implemented"; return 0 }

func (c *Dynamic) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (c *Dynamic) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (c *Dynamic) appendDiscriminatorRow(d int) { _ = "STUB: not implemented"; return }

func (c *Dynamic) appendNullRow() { _ = "STUB: not implemented"; return }

func (c *Dynamic) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Dynamic) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

// If preferred type wasn't provided, try each column

// Do not try to fit into SharedVariant

// If no existing columns match, try matching a ClickHouse type from common Go types

func (c *Dynamic) encodeHeader(buffer *proto.Buffer) error { _ = "STUB: not implemented"; return nil }

func discriminatorWriter(totalTypes uint64, buffer *proto.Buffer) func(uint64) {
	_ = "STUB: not implemented"
	return nil
}

func (c *Dynamic) encodeData(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (c *Dynamic) WriteStatePrefix(buffer *proto.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Dynamic) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (c *Dynamic) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (c *Dynamic) Reset() { _ = "STUB: not implemented"; return }

func (c *Dynamic) decodeHeader(reader *proto.Reader) error { _ = "STUB: not implemented"; return nil }

func discriminatorReader(totalTypes uint64, reader *proto.Reader) func() (uint64, error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *Dynamic) decodeData(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Dynamic) ReadStatePrefix(reader *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Dynamic) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}
