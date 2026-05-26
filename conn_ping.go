package clickhouse

import (
	"context"
)

// Connection::ping
// https://github.com/ClickHouse/ClickHouse/blob/master/src/Client/Connection.cpp
func (c *connect) ping(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	// set a read deadline - alternative to context.Read operation will fail if no data is received after deadline.
	return nil
}

// context level deadlines override any read deadline
