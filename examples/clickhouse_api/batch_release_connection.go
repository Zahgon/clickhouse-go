package clickhouse_api

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

func BatchWithReleaseConnection() error { _ = "STUB: not implemented"; return nil }

type YourBatch struct {
	ctx context.Context

	insertStatement string

	conn  driver.Conn
	batch driver.Batch
}

func New(ctx context.Context, conn driver.Conn, insertStatement string) (*YourBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *YourBatch) Append(col1 uint64, col2 string) error { _ = "STUB: not implemented"; return nil }

func (b *YourBatch) Send() error { _ = "STUB: not implemented"; return nil }

func (b *YourBatch) reset() error { _ = "STUB: not implemented"; return nil }
