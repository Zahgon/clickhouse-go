package clickhouse

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"

	"github.com/ClickHouse/clickhouse-go/v2/ext"
)

var _contextOptionKey = &QueryOptions{
	settings: Settings{
		"_contextOption": struct{}{},
	},
}

type Settings map[string]any

// CustomSetting is a helper struct to distinguish custom settings from important ones.
// For native protocol, is_important flag is set to value 0x02 (see https://github.com/ClickHouse/ClickHouse/blob/c873560fe7185f45eed56520ec7d033a7beb1551/src/Core/BaseSettings.h#L516-L521)
// Only string value is supported until formatting logic that exists in ClickHouse is implemented in clickhouse-go. (https://github.com/ClickHouse/ClickHouse/blob/master/src/Core/Field.cpp#L312 and https://github.com/ClickHouse/clickhouse-go/issues/992)
type CustomSetting struct {
	Value string
}

// ColumnNameAndType represents a column name and type
type ColumnNameAndType struct {
	Name string
	Type string
}

type Parameters map[string]string
type (
	QueryOption  func(*QueryOptions) error
	AsyncOptions struct {
		ok   bool
		wait bool
	}
	QueryOptions struct {
		span     trace.SpanContext
		async    AsyncOptions
		queryID  string
		quotaKey string
		jwt      string
		events   struct {
			logs          func(*Log)
			progress      func(*Progress)
			profileInfo   func(*ProfileInfo)
			profileEvents func([]ProfileEvent)
		}
		settings            Settings
		parameters          Parameters
		external            []*ext.Table
		blockBufferSize     uint8
		userLocation        *time.Location
		columnNamesAndTypes []ColumnNameAndType
		clientInfo          ClientInfo
	}
)

func WithSpan(span trace.SpanContext) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

func WithQueryID(queryID string) QueryOption { _ = "STUB: not implemented"; return *new(QueryOption) }

func WithBlockBufferSize(size uint8) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

func WithQuotaKey(quotaKey string) QueryOption { _ = "STUB: not implemented"; return *new(QueryOption) }

// WithJWT overrides the existing authentication with the given JWT.
// This only applies for clients connected with HTTPS to ClickHouse Cloud.
func WithJWT(jwt string) QueryOption { _ = "STUB: not implemented"; return *new(QueryOption) }

// WithColumnNamesAndTypes is used to provide a predetermined list of
// column names and types for HTTP inserts.
// Without this, the HTTP implementation will parse the query and run a
// DESCRIBE TABLE request to fetch and validate column names.
func WithColumnNamesAndTypes(columnNamesAndTypes []ColumnNameAndType) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

// WithClientInfo appends client info data to the query, visible in the system.query_log table.
// This does not replace the client info provided in the connection options, it appends to it.
// Can be called multiple times to append more info.
func WithClientInfo(ci ClientInfo) QueryOption { _ = "STUB: not implemented"; return *new(QueryOption) }

func WithSettings(settings Settings) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

func WithParameters(params Parameters) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

func WithLogs(fn func(*Log)) QueryOption { _ = "STUB: not implemented"; return *new(QueryOption) }

func WithProgress(fn func(*Progress)) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

func WithProfileInfo(fn func(*ProfileInfo)) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

func WithProfileEvents(fn func([]ProfileEvent)) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

func WithExternalTable(t ...*ext.Table) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

func WithAsync(wait bool) QueryOption { _ = "STUB: not implemented"; return *new(QueryOption) }

// Deprecated: use `WithAsync` instead.
func WithStdAsync(wait bool) QueryOption { _ = "STUB: not implemented"; return *new(QueryOption) }

func WithUserLocation(location *time.Location) QueryOption {
	_ = "STUB: not implemented"
	return *new(QueryOption)
}

func ignoreExternalTables() QueryOption { _ = "STUB: not implemented"; return *new(QueryOption) }

// Context returns a derived context with the given ClickHouse QueryOptions.
// Existing QueryOptions will be overwritten per option if present.
// The QueryOptions Settings map will be initialized if nil.
func Context(parent context.Context, options ...QueryOption) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// queryOptions returns a mutable copy of the QueryOptions struct within the given context.
// If ClickHouse context was not provided, an empty struct with a valid Settings map is returned.
// If the context has a deadline greater than 1s then max_execution_time setting is appended.
func queryOptions(ctx context.Context) QueryOptions {
	_ = "STUB: not implemented"
	return *new(QueryOptions)
}

// queryOptionsJWT returns the JWT within the given context's QueryOptions.
// Empty string if not present.
func queryOptionsJWT(ctx context.Context) string { _ = "STUB: not implemented"; return "" }

// queryOptionsAsync returns the AsyncOptions struct within the given context's QueryOptions.
func queryOptionsAsync(ctx context.Context) AsyncOptions {
	_ = "STUB: not implemented"
	return *new(AsyncOptions)
}

// queryOptionsUserLocation returns the *time.Location within the given context's QueryOptions.
func queryOptionsUserLocation(ctx context.Context) *time.Location {
	_ = "STUB: not implemented"
	return nil
}

// WithoutProfileEvents instructs the server not to send profile events for this query.
// This is a performance optimization for servers >= 25.11 that support the send_profile_events setting.
// On older servers, the setting is unknown and the server will return an error.
func WithoutProfileEvents() QueryOption { _ = "STUB: not implemented"; return *new(QueryOption) }

func (q *QueryOptions) onProcess() *onProcess { _ = "STUB: not implemented"; return nil }

// clone returns a copy of QueryOptions where Settings and Parameters are safely mutable.
func (q *QueryOptions) clone() QueryOptions { _ = "STUB: not implemented"; return *new(QueryOptions) }
