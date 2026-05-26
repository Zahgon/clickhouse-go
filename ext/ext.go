package ext

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
)

func NewTable(name string, columns ...func(t *Table) error) (*Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Table struct {
	name  string
	block *proto.Block
}

func (tbl *Table) Name() string { _ = "STUB: not implemented"; return "" }

func (tbl *Table) Structure() string { _ = "STUB: not implemented"; return "" }

func (tbl *Table) Block() *proto.Block { _ = "STUB: not implemented"; return nil }

func (tbl *Table) Append(v ...any) error { _ = "STUB: not implemented"; return nil }

func Column(name string, ct column.Type) func(t *Table) error {
	_ = "STUB: not implemented"
	return nil
}
