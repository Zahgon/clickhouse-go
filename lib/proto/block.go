package proto

import (
	"github.com/ClickHouse/ch-go/proto"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
)

type Block struct {
	names         []string
	Packet        byte
	Columns       []column.Interface
	ServerContext *column.ServerContext
}

func NewBlock() *Block { _ = "STUB: not implemented"; return nil }

func (b *Block) Rows() int { _ = "STUB: not implemented"; return 0 }

func (b *Block) AddColumn(name string, ct column.Type) error { _ = "STUB: not implemented"; return nil }

func (b *Block) Append(v ...any) (err error) { _ = "STUB: not implemented"; return nil }

func (b *Block) ColumnsNames() []string {
	_ = "STUB: not implemented"

	// SortColumns sorts our block according to the requested order - a slice of column names. Names must be identical in requested order and block.
	return nil
}

func (b *Block) SortColumns(columns []string) error { _ = "STUB: not implemented"; return nil }

// no preferred sort order

// we assume both lists have the same

func difference(a, b []string) []string { _ = "STUB: not implemented"; return nil }

func (b *Block) EncodeHeader(buffer *proto.Buffer, revision uint64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *Block) EncodeColumn(buffer *proto.Buffer, revision uint64, i int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *Block) Encode(buffer *proto.Buffer, revision uint64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *Block) Decode(reader *proto.Reader, revision uint64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (b *Block) Reset() { _ = "STUB: not implemented"; return }

func encodeBlockInfo(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func decodeBlockInfo(reader *proto.Reader) error { _ = "STUB: not implemented"; return nil }

type BlockError struct {
	Op         string
	Err        error
	ColumnName string
}

func (e *BlockError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *BlockError) Unwrap() error { _ = "STUB: not implemented"; return nil }
