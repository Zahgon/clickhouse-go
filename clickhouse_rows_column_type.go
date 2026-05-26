package clickhouse

import (
	"reflect"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type columnType struct {
	name     string
	chType   string
	nullable bool
	scanType reflect.Type
}

func (c *columnType) Name() string { _ = "STUB: not implemented"; return "" }

func (c *columnType) Nullable() bool { _ = "STUB: not implemented"; return false }

func (c *columnType) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (c *columnType) DatabaseTypeName() string { _ = "STUB: not implemented"; return "" }

func (r *rows) ColumnTypes() []driver.ColumnType { _ = "STUB: not implemented"; return nil }
