package clickhouse

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	_ "time/tzdata"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
)

type Conn = driver.Conn

type (
	Progress      = proto.Progress
	Exception     = proto.Exception
	ProfileInfo   = proto.ProfileInfo
	ServerVersion = proto.ServerHandshake
)

var (
	ErrBatchInvalid              = errors.New("clickhouse: batch is invalid. check appended data is correct")
	ErrBatchAlreadySent          = errors.New("clickhouse: batch has already been sent")
	ErrBatchNotSent              = errors.New("clickhouse: invalid retry, batch not sent yet")
	ErrAcquireConnTimeout        = errors.New("clickhouse: acquire conn timeout. you can increase the number of max open conn or the dial timeout")
	ErrUnsupportedServerRevision = errors.New("clickhouse: unsupported server revision")
	ErrBindMixedParamsFormats    = errors.New("clickhouse [bind]: mixed named, numeric or positional parameters")
	ErrAcquireConnNoAddress      = errors.New("clickhouse: no valid address supplied")
	ErrServerUnexpectedData      = errors.New("code: 101, message: Unexpected packet Data received from client")
	ErrConnectionClosed          = errors.New("clickhouse: connection is closed")
)

type OpError struct {
	Op         string
	ColumnName string
	Err        error
}

func (e *OpError) Error() string { _ = "STUB: not implemented"; return "" }

func Open(opt *Options) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

// nativeTransport represents an implementation (TCP or HTTP) that can be pooled by the main clickhouse struct.
// Implementations are not expected to be thread safe, which is why we provide acquire/release functions.
type nativeTransport interface {
	serverVersion() (*ServerVersion, error)
	query(ctx context.Context, release nativeTransportRelease, query string, args ...any) (*rows, error)
	queryRow(ctx context.Context, release nativeTransportRelease, query string, args ...any) *row
	prepareBatch(ctx context.Context, release nativeTransportRelease, acquire nativeTransportAcquire, query string, opts driver.PrepareBatchOptions) (driver.Batch, error)
	exec(ctx context.Context, query string, args ...any) error
	asyncInsert(ctx context.Context, query string, wait bool, args ...any) error
	ping(context.Context) error
	isBad() bool
	connID() int
	connectedAtTime() time.Time
	isReleased() bool
	setReleased(released bool)
	getLogger() *slog.Logger
	// freeBuffer is called if Options.FreeBufOnConnRelease is set
	freeBuffer()
	close() error
}
type nativeTransportAcquire func(context.Context) (nativeTransport, error)
type nativeTransportRelease func(nativeTransport, error)

// connectionPooler is an connection pool maintain
// idle connections.
type connectionPooler interface {
	Get(ctx context.Context) (nativeTransport, error)
	Put(conn nativeTransport)
	Len() int
	Cap() int
	Close() error
}

type clickhouse struct {
	opt    *Options
	connID int64

	idle connectionPooler
	open chan struct{}

	closeOnce *sync.Once
	closed    *atomic.Bool
}

func (clickhouse) Contributors() []string { _ = "STUB: not implemented"; return nil }

func (ch *clickhouse) ServerVersion() (*driver.ServerVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ch *clickhouse) Query(ctx context.Context, query string, args ...any) (rows driver.Rows, err error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

func (ch *clickhouse) QueryRow(ctx context.Context, query string, args ...any) driver.Row {
	_ = "STUB: not implemented"
	return *new(driver.Row)
}

func (ch *clickhouse) Exec(ctx context.Context, query string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *clickhouse) PrepareBatch(ctx context.Context, query string, opts ...driver.PrepareBatchOption) (driver.Batch, error) {
	_ = "STUB: not implemented"
	return *new(driver.Batch), nil
}

func getPrepareBatchOptions(opts ...driver.PrepareBatchOption) driver.PrepareBatchOptions {
	_ = "STUB: not implemented"
	return *new(driver.PrepareBatchOptions)
}

// Deprecated: use context aware `WithAsync()` for any async operations
func (ch *clickhouse) AsyncInsert(ctx context.Context, query string, wait bool, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *clickhouse) Ping(ctx context.Context) (err error) { _ = "STUB: not implemented"; return nil }

func (ch *clickhouse) Stats() driver.Stats { _ = "STUB: not implemented"; return *new(driver.Stats) }

func (ch *clickhouse) dial(ctx context.Context) (conn nativeTransport, err error) {
	_ = "STUB: not implemented"
	return *new(nativeTransport), nil
}

func DefaultDialStrategy(ctx context.Context, connID int, opt *Options, dial Dial) (r DialResult, err error) {
	_ = "STUB: not implemented"
	return *new(DialResult), nil
}

func (ch *clickhouse) acquire(ctx context.Context) (conn nativeTransport, err error) {
	_ = "STUB: not implemented"
	return *new(nativeTransport), nil
}

// If context is already cancelled, just return without any work
// done this way with single case with default. Otherwise if both ctx is cancelled and ch.open is ready,
// Go would choose one of those at random, thus missing to return deterministically when context is cancelled
// at this point in time.
// Known pattern: https://go.dev/ref/spec#Select_statements

func (ch *clickhouse) release(conn nativeTransport, err error) { _ = "STUB: not implemented"; return }

func (ch *clickhouse) Close() (err error) { _ = "STUB: not implemented"; return nil }
