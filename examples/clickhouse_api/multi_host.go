package clickhouse_api

import (
	"github.com/ClickHouse/clickhouse-go/v2"
)

func MultiHostVersion() error { _ = "STUB: not implemented"; return nil }

func MultiHostRoundRobinVersion() error { _ = "STUB: not implemented"; return nil }

func MultiHostRandomVersion() error { _ = "STUB: not implemented"; return nil }

func multiHostVersion(connOpenStrategy *clickhouse.ConnOpenStrategy) error {
	_ = "STUB: not implemented"
	return nil
}
