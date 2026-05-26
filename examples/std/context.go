package std

func UseContext() error { _ = "STUB: not implemented"; return nil }

// we can use context to pass settings to a specific API call

// queries can be cancelled using the context

// set a deadline for a query - this will cancel the query after the absolute time is reached. Again terminates the connection only,
// queries will continue to completion in ClickHouse

// set a query id to assist tracing queries in logs e.g. see system.query_log

// set a quota key - first create the quota

// queries can be cancelled using the context

// we will get some results before cancel

// query is cancelled using Context, so rows.Err() will return context canceled error

// expected cancellation after col2 == 3
