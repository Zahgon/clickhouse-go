package clickhouse

import (
	"context"
	"log/slog"
	"net"
	"sync"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"

	"github.com/ClickHouse/ch-go/compress"
	chproto "github.com/ClickHouse/ch-go/proto"

	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
)

func dial(ctx context.Context, addr string, num int, opt *Options) (*connect, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get base logger and enrich with connection-specific context

// warn only on the first connection in the pool

// https://github.com/ClickHouse/ClickHouse/blob/master/src/Client/Connection.cpp
type connect struct {
	id                   int
	opt                  *Options
	conn                 net.Conn
	logger               *slog.Logger
	server               ServerVersion
	closed               bool
	buffer               *chproto.Buffer
	reader               *chproto.Reader
	released             bool
	revision             uint64
	structMap            *structMap
	compression          CompressionMethod
	connectedAt          time.Time
	compressor           *compress.Writer
	readTimeout          time.Duration
	blockBufferSize      uint8
	maxCompressionBuffer int
	readerMutex          sync.Mutex
	closeMutex           sync.Mutex
}

func (c *connect) connID() int { _ = "STUB: not implemented"; return 0 }

func (c *connect) getLogger() *slog.Logger { _ = "STUB: not implemented"; return nil }

func (c *connect) connectedAtTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *connect) serverVersion() (*ServerVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *connect) settings(querySettings Settings) []proto.Setting {
	_ = "STUB: not implemented"
	return nil
}

func (c *connect) isBad() bool { _ = "STUB: not implemented"; return false }

func (c *connect) isReleased() bool { _ = "STUB: not implemented"; return false }

func (c *connect) setReleased(released bool) { _ = "STUB: not implemented"; return }

func (c *connect) isClosed() bool { _ = "STUB: not implemented"; return false }

func (c *connect) setClosed() { _ = "STUB: not implemented"; return }

func (c *connect) close() error { _ = "STUB: not implemented"; return nil }

func (c *connect) progress() (*Progress, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *connect) exception() error { _ = "STUB: not implemented"; return nil }

func (c *connect) compressBuffer(start int) error { _ = "STUB: not implemented"; return nil }

func (c *connect) sendData(block *proto.Block, name string) error {
	_ = "STUB: not implemented"
	return nil
}

func serverVersionToContext(v ServerVersion) column.ServerContext {
	_ = "STUB: not implemented"
	return *new(column.ServerContext)
}

func (c *connect) readData(ctx context.Context, packet byte, compressible bool) (*proto.Block, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *connect) freeBuffer() { _ = "STUB: not implemented"; return }

func (c *connect) flush() error { _ = "STUB: not implemented"; return nil }

// Nothing to flush.

// startReadWriteTimeout applies the configured read timeout to conn.
// If a context deadline is provided, a read and write deadline is set.
// This should be matched with a deferred call to clearReadWriteTimeout.
func (c *connect) startReadWriteTimeout(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// context level deadlines override configured read timeout

// clearReadWriteTimeout removes the read timeout from conn.
// If a context deadline is provided, the read and write timeout is cleared too.
func (c *connect) clearReadWriteTimeout(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// context level deadlines should clear read + write deadlines.
