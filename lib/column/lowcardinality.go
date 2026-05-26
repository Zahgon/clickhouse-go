package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

const indexTypeMask = 0b11111111

const (
	keyUInt8  = 0
	keyUInt16 = 1
	keyUInt32 = 2
	keyUInt64 = 3
)

const (
	/// Need to read dictionary if it wasn't.
	needGlobalDictionaryBit = 1 << 8
	/// Need to read additional keys. Additional keys are stored before indexes as value N and N keys after them.
	hasAdditionalKeysBit = 1 << 9
	/// Need to update dictionary. It means that previous granule has different dictionary.
	needUpdateDictionary = 1 << 10

	updateAll = hasAdditionalKeysBit | needUpdateDictionary
)

const sharedDictionariesWithAdditionalKeys = 1

// https://github.com/ClickHouse/ClickHouse/blob/master/src/Columns/ColumnLowCardinality.cpp
// https://github.com/ClickHouse/clickhouse-cpp/blob/master/clickhouse/columns/lowcardinality.cpp
type LowCardinality struct {
	key      byte
	rows     int
	index    Interface
	chType   Type
	nullable bool

	keys8  UInt8
	keys16 UInt16
	keys32 UInt32
	keys64 UInt64

	append struct {
		keys  []int
		index map[any]int
	}
	name string
}

func (col *LowCardinality) Reset() { _ = "STUB: not implemented"; return }

func (col *LowCardinality) Name() string { _ = "STUB: not implemented"; return "" }

func (col *LowCardinality) parse(t Type, sc *ServerContext) (_ *LowCardinality, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *LowCardinality) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *LowCardinality) ScanType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (col *LowCardinality) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *LowCardinality) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *LowCardinality) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *LowCardinality) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *LowCardinality) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

// init

// second check is unfortunate - but we could be passed a *type(nil) e.g. via LowCardinality(Nullable(String))

func (col *LowCardinality) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *LowCardinality) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

// We already have keys, so this column is probably in a block directly decoded from the server, and we should
// not reset them

func (col *LowCardinality) ReadStatePrefix(reader *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *LowCardinality) WriteStatePrefix(buffer *proto.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *LowCardinality) keys() Interface { _ = "STUB: not implemented"; return *new(Interface) }

func (col *LowCardinality) indexRowNum(row int) int { _ = "STUB: not implemented"; return 0 }

var (
	_ Interface           = (*LowCardinality)(nil)
	_ CustomSerialization = (*LowCardinality)(nil)
)
