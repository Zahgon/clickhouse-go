package clickhouse_api

func UseContext() error { _ = "STUB: not implemented"; return nil }

// we can use context to pass settings to a specific API call

// queries can be cancelled using the context

// set a deadline for a query - this will cancel the query after the absolute time is reached.
// queries will continue to completion in ClickHouse

// set a query id to assist tracing queries in logs e.g. see system.query_log

// set a quota key - first create the quota
