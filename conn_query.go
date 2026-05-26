package clickhouse

import (
	"context"
)

func (c *connect) query(ctx context.Context, release nativeTransportRelease, query string, args ...any) (*rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// allow block buffer sze to be overridden per query

func (c *connect) queryRow(ctx context.Context, release nativeTransportRelease, query string, args ...any) *row {
	_ = "STUB: not implemented"
	return nil
}
