package clickhouse

import (
	"context"
)

func (c *connect) exec(ctx context.Context, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

// set a read deadline - alternative to context.Read operation will fail if no data is received after deadline.

// context level deadlines override any read deadline
