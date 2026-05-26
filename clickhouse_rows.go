package clickhouse

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
)

type rows struct {
	err       error
	row       int
	block     *proto.Block
	totals    *proto.Block
	errors    chan error
	stream    chan *proto.Block
	columns   []string
	structMap *structMap
	closed    bool
}

func (r *rows) Next() (result bool) { _ = "STUB: not implemented"; return false }

func (r *rows) Scan(dest ...any) error { _ = "STUB: not implemented"; return nil }

// call without next when result is empty

func (r *rows) ScanStruct(dest any) error { _ = "STUB: not implemented"; return nil }

func (r *rows) Totals(dest ...any) error { _ = "STUB: not implemented"; return nil }

func (r *rows) Columns() []string { _ = "STUB: not implemented"; return nil }

func (r *rows) Close() error { _ = "STUB: not implemented"; return nil }

func (r *rows) Err() error { _ = "STUB: not implemented"; return nil }

func (r *rows) HasData() bool { _ = "STUB: not implemented"; return false }

type row struct {
	err  error
	rows *rows
}

func (r *row) Err() error { _ = "STUB: not implemented"; return nil }

func (r *row) ScanStruct(dest any) error { _ = "STUB: not implemented"; return nil }

func (r *row) Scan(dest ...any) error { _ = "STUB: not implemented"; return nil }
