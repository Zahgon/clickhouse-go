package clickhouse

import (
	"context"
)

func (h *httpConnect) exec(ctx context.Context, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:bodyclose // false positive
