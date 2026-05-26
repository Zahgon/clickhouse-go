package std

import (
	"crypto/tls"
	"database/sql"

	"github.com/ClickHouse/clickhouse-go/v2"
	clickhouse_tests "github.com/ClickHouse/clickhouse-go/v2/tests"
)

const TestSet string = "examples_std_api"

func GetStdDSNConnection(protocol clickhouse.Protocol, secure bool, compress string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetStdOpenDBConnection(protocol clickhouse.Protocol, settings clickhouse.Settings, tlsConfig *tls.Config, compression *clickhouse.Compression) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetStdTestEnvironment() (clickhouse_tests.ClickHouseTestEnvironment, error) {
	_ = "STUB: not implemented"
	return *new(clickhouse_tests.ClickHouseTestEnvironment), nil
}

func CheckMinServerVersion(conn *sql.DB, major, minor, patch uint64) bool {
	_ = "STUB: not implemented"
	return false
}
