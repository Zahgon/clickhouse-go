package clickhouse_api

import (
	"time"
)

type row struct {
	Col1       uint64
	Col4       time.Time
	Col2       string
	Col3       []uint8
	ColIgnored string
}

func AppendStruct() error { _ = "STUB: not implemented"; return nil }
