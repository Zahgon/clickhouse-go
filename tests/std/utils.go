package std

import (
	"crypto/tls"
	"database/sql"
	"net/url"

	"github.com/ClickHouse/clickhouse-go/v2"
	clickhouse_tests "github.com/ClickHouse/clickhouse-go/v2/tests"
)

func GetStdTestEnvironment() (clickhouse_tests.ClickHouseTestEnvironment, error) {
	_ = "STUB: not implemented"
	return *new(clickhouse_tests.ClickHouseTestEnvironment), nil
}

func CheckMinServerVersion(conn *sql.DB, major, minor, patch uint64) bool {
	_ = "STUB: not implemented"
	return false
}

func GetDSNConnection(environment string, protocol clickhouse.Protocol, secure bool, opts url.Values) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetConnectionFromDSN(dsn string) (*sql.DB, error) { _ = "STUB: not implemented"; return nil, nil }

func GetConnectionFromDSNWithSessionID(dsn string, sessionID string) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Optionally provide session ID after initial version check to prevent locking

func GetConnectionWithOptions(options *clickhouse.Options) *sql.DB {
	_ = "STUB: not implemented"
	return nil
}

func GetOpenDBConnection(environment string, protocol clickhouse.Protocol, settings clickhouse.Settings, tlsConfig *tls.Config, compression *clickhouse.Compression) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetOpenDBConnectionJWT(environment string, protocol clickhouse.Protocol, settings clickhouse.Settings, tlsConfig *tls.Config, jwtFunc clickhouse.GetJWTFunc) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ToJson(obj any) string { _ = "STUB: not implemented"; return "" }
