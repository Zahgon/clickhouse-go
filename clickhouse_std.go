package clickhouse

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"log/slog"
	"reflect"

	chdriver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

var globalConnID int64

type stdConnOpener struct {
	err    error
	opt    *Options
	logger *slog.Logger
}

func (o *stdConnOpener) Driver() driver.Driver {
	_ = "STUB: not implemented"
	return *new(driver.Driver)
}

func (o *stdConnOpener) Connect(ctx context.Context) (_ driver.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

// Create a logger with connection-specific context

var _ driver.Connector = (*stdConnOpener)(nil)

func init() {
	sql.Register("clickhouse", &stdDriver{logger: newNoopLogger()})
}

// isConnBrokenError returns true if the error class indicates that the
// db connection is no longer usable and should be marked bad
func isConnBrokenError(err error) bool { _ = "STUB: not implemented"; return false }

func Connector(opt *Options) driver.Connector {
	_ = "STUB: not implemented"
	return *new(driver.Connector)
}

func OpenDB(opt *Options) *sql.DB { _ = "STUB: not implemented"; return nil }

// Ok to set these configs irrespective of values in opt.
// Because opt.setDefaults() would have set some sane values
// for these configs.

type stdConnect interface {
	isBad() bool
	close() error
	query(ctx context.Context, release nativeTransportRelease, query string, args ...any) (*rows, error)
	exec(ctx context.Context, query string, args ...any) error
	ping(ctx context.Context) (err error)
	prepareBatch(ctx context.Context, release nativeTransportRelease, acquire nativeTransportAcquire, query string, options chdriver.PrepareBatchOptions) (chdriver.Batch, error)
	asyncInsert(ctx context.Context, query string, wait bool, args ...any) error
}

type stdDriver struct {
	opt    *Options
	conn   stdConnect
	commit func() error
	logger *slog.Logger
}

var _ driver.Conn = (*stdDriver)(nil)
var _ driver.ConnBeginTx = (*stdDriver)(nil)
var _ driver.ExecerContext = (*stdDriver)(nil)
var _ driver.QueryerContext = (*stdDriver)(nil)
var _ driver.ConnPrepareContext = (*stdDriver)(nil)

func (std *stdDriver) Open(dsn string) (_ driver.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

var _ driver.Driver = (*stdDriver)(nil)

func (std *stdDriver) ResetSession(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

var _ driver.SessionResetter = (*stdDriver)(nil)

func (std *stdDriver) Ping(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

var _ driver.Pinger = (*stdDriver)(nil)

func (std *stdDriver) Begin() (driver.Tx, error) {
	_ = "STUB: not implemented"
	return *new(driver.Tx), nil
}

func (std *stdDriver) BeginTx(ctx context.Context, opts driver.TxOptions) (driver.Tx, error) {
	_ = "STUB: not implemented"
	return *new(driver.Tx), nil
}

func (std *stdDriver) Commit() error { _ = "STUB: not implemented"; return nil }

func (std *stdDriver) Rollback() error { _ = "STUB: not implemented"; return nil }

var _ driver.Tx = (*stdDriver)(nil)

func (std *stdDriver) CheckNamedValue(nv *driver.NamedValue) error {
	_ = "STUB: not implemented"
	return nil
}

var _ driver.NamedValueChecker = (*stdDriver)(nil)

func (std *stdDriver) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

func (std *stdDriver) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	_ = "STUB: not implemented"
	return *new(driver.Rows), nil
}

func (std *stdDriver) Prepare(query string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

func (std *stdDriver) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	_ = "STUB: not implemented"
	return *new(driver.Stmt), nil
}

func (std *stdDriver) Close() error { _ = "STUB: not implemented"; return nil }

type stdBatch struct {
	batch  chdriver.Batch
	logger *slog.Logger
}

func (s *stdBatch) NumInput() int { _ = "STUB: not implemented"; return 0 }
func (s *stdBatch) Exec(args []driver.Value) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

func (s *stdBatch) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	_ = "STUB: not implemented"
	return *new(driver.Result), nil
}

var _ driver.StmtExecContext = (*stdBatch)(nil)

func (s *stdBatch) Query(args []driver.Value) (driver.Rows, error) {
	_ = "STUB: not implemented"
	// Note: not implementing driver.StmtQueryContext accordingly
	return *new(driver.Rows), nil
}

func (s *stdBatch) Close() error { _ = "STUB: not implemented"; return nil }

type stdRows struct {
	rows   *rows
	logger *slog.Logger
}

func (r *stdRows) Columns() []string { _ = "STUB: not implemented"; return nil }

func (r *stdRows) ColumnTypeScanType(idx int) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

var _ driver.RowsColumnTypeScanType = (*stdRows)(nil)

func (r *stdRows) ColumnTypeDatabaseTypeName(idx int) string { _ = "STUB: not implemented"; return "" }

func (r *stdRows) ColumnTypeNullable(idx int) (nullable, ok bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (r *stdRows) ColumnTypePrecisionScale(idx int) (precision, scale int64, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

var _ driver.Rows = (*stdRows)(nil)
var _ driver.RowsNextResultSet = (*stdRows)(nil)
var _ driver.RowsColumnTypeDatabaseTypeName = (*stdRows)(nil)
var _ driver.RowsColumnTypeNullable = (*stdRows)(nil)
var _ driver.RowsColumnTypePrecisionScale = (*stdRows)(nil)

func (r *stdRows) Next(dest []driver.Value) error { _ = "STUB: not implemented"; return nil }

// We don't know what is the destination type at this stage,
// but destination type might be a sql.Null* type that expects to receive a value
// instead of a pointer to a value. ClickHouse-go returns pointers to values for nullable columns.
//
// This is a compatibility layer to make sure that the driver works with the standard library.
// Due to reflection used it has a performance cost.

func (r *stdRows) HasNextResultSet() bool { _ = "STUB: not implemented"; return false }

func (r *stdRows) NextResultSet() error { _ = "STUB: not implemented"; return nil }

var _ driver.RowsNextResultSet = (*stdRows)(nil)

func (r *stdRows) Close() error { _ = "STUB: not implemented"; return nil }

var _ driver.Rows = (*stdRows)(nil)
