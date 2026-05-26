package clickhouse

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"

	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
)

type scanSelectQueryFunc func(ctx context.Context, query string, args ...any) (driver.Rows, error)

func scanSelect(queryFunc scanSelectQueryFunc, ctx context.Context, dest any, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// dest should point to empty slice
// to make select result correct

func (ch *clickhouse) Select(ctx context.Context, dest any, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func scan(block *proto.Block, row int, dest ...any) error { _ = "STUB: not implemented"; return nil }
