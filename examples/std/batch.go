package std

import (
	_ "github.com/ClickHouse/clickhouse-go/v2"
)

func BatchInsert() error { _ = "STUB: not implemented"; return nil }

// Map(String, UInt8)
// Array(String)
// Tuple(String, UInt8, Array(Map(String, String)))
