package clickhouse

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
)

type onProcess struct {
	data          func(*proto.Block)
	logs          func([]Log)
	progress      func(*Progress)
	profileInfo   func(*ProfileInfo)
	profileEvents func([]ProfileEvent)
}

func (c *connect) firstBlock(ctx context.Context, on *onProcess) (*proto.Block, error) {
	_ = "STUB: not implemented"
	// if context is already timedout/cancelled — we're done
	return nil, nil
}

// do reads in background

// select on context or read channels (results/errors)

func (c *connect) firstBlockImpl(ctx context.Context, on *onProcess) (*proto.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// handling error, return

// handled okay, read next byte

func (c *connect) process(ctx context.Context, on *onProcess) error {
	_ = "STUB: not implemented"
	// if context is already timedout/cancelled — we're done
	return nil
}

// do reads in background

// select on context or read channel (errors)

func (c *connect) processImpl(ctx context.Context, on *onProcess) error {
	_ = "STUB: not implemented"
	return nil
}

// handling error, return

// handled okay, read next byte

func (c *connect) handle(ctx context.Context, packet byte, on *onProcess) error {
	_ = "STUB: not implemented"
	return nil
}

// Progress is already logged in c.progress()

func (c *connect) cancel() error { _ = "STUB: not implemented"; return nil }

// don't reuse a cancelled query as we don't drain the connection
