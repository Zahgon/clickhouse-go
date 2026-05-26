package clickhouse

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
)

// Connection::sendQuery
// https://github.com/ClickHouse/ClickHouse/blob/master/src/Client/Connection.cpp
func (c *connect) sendQuery(body string, o *QueryOptions) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:govet

func parametersToProtoParameters(parameters Parameters) (s proto.Parameters) {
	_ = "STUB: not implemented"
	return *new(proto.Parameters)
}
