package clickhouse_api

import (
	"crypto/tls"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	clickhouse_tests "github.com/ClickHouse/clickhouse-go/v2/tests"
)

const TestSet string = "examples_clickhouse_api"

func GetNativeConnection(settings clickhouse.Settings, tlsConfig *tls.Config, compression *clickhouse.Compression) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func GetHTTPConnection(sessionName string, settings clickhouse.Settings, tlsConfig *tls.Config, compression *clickhouse.Compression) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func GetNativeTestEnvironment() (clickhouse_tests.ClickHouseTestEnvironment, error) {
	_ = "STUB: not implemented"
	return *new(clickhouse_tests.ClickHouseTestEnvironment), nil
}

func GetNativeConnectionWithOptions(settings clickhouse.Settings, tlsConfig *tls.Config, compression *clickhouse.Compression) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func CheckMinServerVersion(conn driver.Conn, major, minor, patch uint64) bool {
	_ = "STUB: not implemented"
	return false
}

var randSeed = time.Now().UnixNano()

func ResetRandSeed() { _ = "STUB: not implemented"; return }
