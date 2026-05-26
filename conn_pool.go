package clickhouse

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/internal/circular"
)

var errQueueEmpty = errors.New("clickhouse: connection pool queue is empty")

type connPool struct {
	mu    sync.RWMutex
	conns *circular.Queue[nativeTransport]

	ticker   *time.Ticker
	finish   chan struct{}
	finished chan struct{}

	maxConnLifetime time.Duration
}

func newConnPool(lifetime time.Duration, capacity int) *connPool {
	_ = "STUB: not implemented"
	return nil
}

func (i *connPool) Len() int { _ = "STUB: not implemented"; return 0 }

func (i *connPool) Cap() int { _ = "STUB: not implemented"; return 0 }

func (i *connPool) Get(ctx context.Context) (nativeTransport, error) {
	_ = "STUB: not implemented"
	return *new(nativeTransport), nil
}

// check if pool was closed while we waited on the lock
// return early if pool already closed
// otherwise, pool wont close again while we hold lock
// and so we continue

// this loop continues until either:
// a) the provided context is cancelled
// b) the underlying circular queue is empty
// c) it finds a non-expired connection

// context has been cancelled

// Try to pull a connection

// queue is empty

func (i *connPool) Put(conn nativeTransport) { _ = "STUB: not implemented"; return }

// Try to push the connection

// Buffer is full, close the connection

func (i *connPool) Close() error { _ = "STUB: not implemented"; return nil }

// Drain all remaining connections from the pool

func (i *connPool) closed() bool { _ = "STUB: not implemented"; return false }

func (i *connPool) runDrainPool() { _ = "STUB: not implemented"; return }

// drainPool removes connections from the pool.
// If the pool is closed, it removes all connections.
// Otherwise, it only removes expired connections.
// Must be called with i.mu held.
func (i *connPool) drainPool() {
	_ = "STUB: not implemented"

	// Close all connections
	return
}

// Remove only expired connections

func (i *connPool) isExpired(conn nativeTransport) bool { _ = "STUB: not implemented"; return false }

func (i *connPool) expires(conn nativeTransport) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
