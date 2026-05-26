package clickhouse

import (
	"context"
	"regexp"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
)

var insertMatch = regexp.MustCompile(`(?i)(?:(?:--[^\n]*|#![^\n]*|#\s[^\n]*)\n\s*)*(INSERT\s+INTO\s+[^( ]+(?:\s*\([^()]*(?:\([^()]*\)[^()]*)*\))?)(?:\s*VALUES)?`)
var columnMatch = regexp.MustCompile(`INSERT INTO .+\s\((?P<Columns>.+)\)$`)

func (c *connect) prepareBatch(ctx context.Context, release nativeTransportRelease, acquire nativeTransportAcquire, query string, opts driver.PrepareBatchOptions) (driver.Batch, error) {
	_ = "STUB: not implemented"
	return *new(driver.Batch), nil
}

// resort batch to specified columns

type batch struct {
	err          error
	ctx          context.Context
	query        string
	conn         *connect
	sent         bool // sent signalize that batch is send to ClickHouse.
	released     bool // released signalize that conn was returned to pool and can't be used.
	closeOnFlush bool // closeOnFlush signalize that batch should close query and release conn when use Flush
	block        *proto.Block
	connRelease  func(*connect, error)
	connAcquire  func(context.Context) (*connect, error)
	onProcess    *onProcess
}

func (b *batch) release(err error) { _ = "STUB: not implemented"; return }

func (b *batch) Abort() error { _ = "STUB: not implemented"; return nil }

func (b *batch) Append(v ...any) error { _ = "STUB: not implemented"; return nil }

// appendRowsBlocks is an experimental feature that allows rows blocks be appended directly to the batch.
// This API is not stable and may be changed in the future.
// See: tests/batch_block_test.go
func (b *batch) appendRowsBlocks(r *rows) error { _ = "STUB: not implemented"; return nil }

// make sure the first block is logged

// rows.Next() will read the next block from the server only if the current block is empty
// only if new block is available we should flush the current block
// the last block will be handled by the batch.Send() method

func (b *batch) AppendStruct(v any) error { _ = "STUB: not implemented"; return nil }

func (b *batch) IsSent() bool { _ = "STUB: not implemented"; return false }

func (b *batch) Column(idx int) driver.BatchColumn {
	_ = "STUB: not implemented"
	return *new(driver.BatchColumn)
}

func (b *batch) Send() (err error) { _ = "STUB: not implemented"; return nil }

// close TCP connection on context cancel. There is no other way simple way to interrupt underlying operations.
// as verified in the test, this is safe to do and cleanups resources later on

// there might be an error caused by context cancellation
// in this case we should return context error instead of net.OpError

func (b *batch) resetConnection() (err error) {
	_ = "STUB: not implemented"
	// acquire a new conn
	return nil
}

func (b *batch) Flush() error { _ = "STUB: not implemented"; return nil }

// broken pipe/conn reset aren't generally recoverable on retry

func (b *batch) Rows() int { _ = "STUB: not implemented"; return 0 }

func (b *batch) Columns() []column.Interface { _ = "STUB: not implemented"; return nil }

func (b *batch) closeQuery() error { _ = "STUB: not implemented"; return nil }

// Close will end the current INSERT without sending the currently buffered rows, and release the connection.
// This may result in zero row inserts if no rows were appended.
// If a batch was already sent this does nothing.
// This should be called via defer after a batch is opened to prevent
// batches from falling out of scope and timing out.
func (b *batch) Close() error { _ = "STUB: not implemented"; return nil }

type batchColumn struct {
	err     error
	batch   driver.Batch
	column  column.Interface
	release func(error)
}

func (b *batchColumn) Append(v any) (err error) { _ = "STUB: not implemented"; return nil }

func (b *batchColumn) AppendRow(v any) (err error) { _ = "STUB: not implemented"; return nil }

var (
	_ (driver.Batch)       = (*batch)(nil)
	_ (driver.BatchColumn) = (*batchColumn)(nil)
)
