package clickhouse

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/column"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
)

func fetchColumnNamesAndTypesForInsert(h *httpConnect, release nativeTransportRelease, ctx context.Context, tableName string, requestedColumnNames []string) ([]ColumnNameAndType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// these column types cannot be specified in INSERT queries

// The order of the columns must match the INSERT list, or the DESC table if no insert list was provided

// Validate requested columns present

// Use all columns

func newBlock(h *httpConnect, release nativeTransportRelease, ctx context.Context, query string) (string, *proto.Block, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// If the user didn't supply known column names/types, do expensive DESC TABLE logic

func (h *httpConnect) prepareBatch(ctx context.Context, release nativeTransportRelease, acquire nativeTransportAcquire, query string, opts driver.PrepareBatchOptions) (driver.Batch, error) {
	_ = "STUB: not implemented"
	// release is not used within newBlock since the connection is held for the batch.
	return *new(driver.Batch), nil
}

type httpBatch struct {
	query       string
	err         error
	ctx         context.Context
	conn        *httpConnect
	released    bool
	connRelease nativeTransportRelease
	structMap   *structMap
	sent        bool
	block       *proto.Block
}

func (b *httpBatch) release(err error) { _ = "STUB: not implemented"; return }

func (b *httpBatch) Flush() error {
	_ = "STUB: not implemented"
	// Flush and Send are effectively the same for HTTP, but users should just use Send until we
	// figure out a way to do proper streaming.
	return nil
}

func (b *httpBatch) Close() error { _ = "STUB: not implemented"; return nil }

func (b *httpBatch) Abort() error { _ = "STUB: not implemented"; return nil }

func (b *httpBatch) Append(v ...any) error { _ = "STUB: not implemented"; return nil }

func (b *httpBatch) AppendStruct(v any) error { _ = "STUB: not implemented"; return nil }

func (b *httpBatch) Column(idx int) driver.BatchColumn {
	_ = "STUB: not implemented"
	return *new(driver.BatchColumn)
}

func (b *httpBatch) IsSent() bool { _ = "STUB: not implemented"; return false }

func (b *httpBatch) Send() (err error) { _ = "STUB: not implemented"; return nil }

//nolint:bodyclose // false positive

func (b *httpBatch) Rows() int { _ = "STUB: not implemented"; return 0 }

func (b *httpBatch) Columns() []column.Interface { _ = "STUB: not implemented"; return nil }

var _ driver.Batch = (*httpBatch)(nil)
