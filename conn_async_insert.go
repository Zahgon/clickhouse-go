package clickhouse

import (
	"context"
)

func (c *connect) asyncInsert(ctx context.Context, query string, wait bool, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}
