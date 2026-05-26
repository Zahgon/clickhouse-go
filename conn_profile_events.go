package clickhouse

import (
	"context"
	"time"
)

type ProfileEvent struct {
	Hostname    string
	CurrentTime time.Time
	ThreadID    uint64
	Type        string
	Name        string
	Value       int64
}

func (c *connect) profileEvents(ctx context.Context, scanEvents bool) ([]ProfileEvent, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
