package clickhouse

import (
	"context"
)

func (h *httpConnect) asyncInsert(ctx context.Context, query string, wait bool, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:bodyclose // false positive
