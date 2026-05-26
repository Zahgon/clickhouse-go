package clickhouse

import (
	"context"
	"time"
)

type Log struct {
	Time      time.Time
	TimeMicro uint32
	Hostname  string
	QueryID   string
	ThreadID  uint64
	Priority  int8
	Source    string
	Text      string
}

func (c *connect) logs(ctx context.Context) ([]Log, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
