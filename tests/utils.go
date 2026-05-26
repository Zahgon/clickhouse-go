package tests

import (
	"crypto/tls"
	"database/sql"
	"math/rand"
	"net"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
)

var testUUID = uuid.NewString()[0:12]
var testTimestamp = time.Now().UnixMilli()
var randSeed = time.Now().UnixNano()

const defaultClickHouseVersion = "latest"

func GetClickHouseTestVersion() string { _ = "STUB: not implemented"; return "" }

type ClickHouseTestEnvironment struct {
	ContainerID string
	Port        int
	HttpPort    int
	SslPort     int
	HttpsPort   int
	Host        string
	Username    string
	Password    string
	JWT         string
	Database    string
	Version     proto.Version
	ContainerIP string
	Container   testcontainers.Container `json:"-"`
}

func (env *ClickHouseTestEnvironment) setVersion() { _ = "STUB: not implemented"; return }

func CheckMinServerServerVersion(conn driver.Conn, major, minor, patch uint64) bool {
	_ = "STUB: not implemented"
	return false
}

func CreateClickHouseTestEnvironment(testSet string) (ClickHouseTestEnvironment, error) {
	_ = "STUB: not implemented"
	// create a ClickHouse Container
	return *new(ClickHouseTestEnvironment), nil
}

// attempt use docker for CI

// retry

func SetTestEnvironment(testSet string, environment ClickHouseTestEnvironment) {
	_ = "STUB: not implemented"
	return
}

func GetTestEnvironment(testSet string) (ClickHouseTestEnvironment, error) {
	_ = "STUB: not implemented"
	return *new(ClickHouseTestEnvironment), nil
}

func GetExternalTestEnvironment(testSet string) (ClickHouseTestEnvironment, error) {
	_ = "STUB: not implemented"
	return *new(ClickHouseTestEnvironment), nil
}

func ClientOptionsFromEnv(env ClickHouseTestEnvironment, settings clickhouse.Settings, useHTTP bool) clickhouse.Options {
	_ = "STUB: not implemented"
	return *new(clickhouse.Options)
}

func TestClientWithDefaultOptions(env ClickHouseTestEnvironment, settings clickhouse.Settings) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func TestClientDefaultSettings(env ClickHouseTestEnvironment) clickhouse.Settings {
	_ = "STUB: not implemented"
	return *new(clickhouse.Settings)
}

func TestClientWithDefaultSettings(env ClickHouseTestEnvironment) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func TestDatabaseSQLClientWithDefaultOptions(env ClickHouseTestEnvironment, settings clickhouse.Settings) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func TestDatabaseSQLClientWithDefaultSettings(env ClickHouseTestEnvironment) (*sql.DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetConnection(testSet string, t *testing.T, protocol clickhouse.Protocol, settings clickhouse.Settings, tlsConfig *tls.Config, compression *clickhouse.Compression) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func GetConnectionTCP(testSet string, settings clickhouse.Settings, tlsConfig *tls.Config, compression *clickhouse.Compression) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func GetConnectionHTTP(testSet string, sessionName string, settings clickhouse.Settings, tlsConfig *tls.Config, compression *clickhouse.Compression) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func GetJWTConnection(testSet string, settings clickhouse.Settings, tlsConfig *tls.Config, maxConnLifetime time.Duration, jwtFunc clickhouse.GetJWTFunc) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func GetConnectionWithOptions(options *clickhouse.Options) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func getConnection(env ClickHouseTestEnvironment, database string, settings clickhouse.Settings, tlsConfig *tls.Config, compression *clickhouse.Compression) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func getHTTPConnection(env ClickHouseTestEnvironment, sessionName string, database string, settings clickhouse.Settings, tlsConfig *tls.Config, compression *clickhouse.Compression) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

// Each test uses its own session ID.
// This may be problematic on some tests, but overall it is more consistent.

func getJWTConnection(env ClickHouseTestEnvironment, database string, settings clickhouse.Settings, tlsConfig *tls.Config, maxConnLifetime time.Duration, jwtFunc clickhouse.GetJWTFunc) (driver.Conn, error) {
	_ = "STUB: not implemented"
	return *new(driver.Conn), nil
}

func CreateDatabase(testSet string) error { _ = "STUB: not implemented"; return nil }

const (
	readOnlyReadWriteChangeSettings = 0
	readOnlyRead                    = 1
	readOnlyReadChangeSettings      = 2
)

func createUserWithReadOnlySetting(conn driver.Conn, defaultDatabase string, readOnlyType int) (username, password string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func dropUser(conn driver.Conn, username string) error { _ = "STUB: not implemented"; return nil }

func createSimpleTable(client driver.Conn, table string) error {
	_ = "STUB: not implemented"
	return nil
}

func dropTable(client driver.Conn, table string) error { _ = "STUB: not implemented"; return nil }

func getDatabaseName(testSet string) string { _ = "STUB: not implemented"; return "" }

func getRowsCount(t *testing.T, conn driver.Conn, table string) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func deduplicateTable(t *testing.T, conn driver.Conn, table string) {
	_ = "STUB: not implemented"
	return
}

func GetEnv(key, fallback string) string { _ = "STUB: not implemented"; return "" }

func IsSetInEnv(key string) bool { _ = "STUB: not implemented"; return false }

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberBytes = "123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func RandAsciiString(n int) string { _ = "STUB: not implemented"; return "" }

func RandIntString(n int) string { _ = "STUB: not implemented"; return "" }

func RandIPv4() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func RandIPv6() net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func randString(n int, bytes string) string { _ = "STUB: not implemented"; return "" }

// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garbage collection cycles completed.
// thanks to https://golangcode.com/print-the-current-memory-usage/
func PrintMemUsage() { _ = "STUB: not implemented"; return }

// For info on each, see: https://golang.org/pkg/runtime/#MemStats

func bToMb(b uint64) uint64 { _ = "STUB: not implemented"; return 0 }

type NginxReverseHTTPProxyTestEnvironment struct {
	HttpPort       int
	NginxContainer testcontainers.Container `json:"-"`
}

func CreateNginxReverseProxyTestEnvironment(clickhouseEnv ClickHouseTestEnvironment) (NginxReverseHTTPProxyTestEnvironment, error) {
	_ = "STUB: not implemented"
	// create a nginx Container as a reverse proxy
	return *new(NginxReverseHTTPProxyTestEnvironment), nil
}

// replace upstream clickhouse endpoint

// reload new nginx.conf and set http proxy upstream

type TinyProxyTestEnvironment struct {
	HttpPort  int
	Container testcontainers.Container `json:"-"`
}

func (e TinyProxyTestEnvironment) ProxyUrl(t *testing.T) string {
	_ = "STUB: not implemented"
	return ""
}

func CreateTinyProxyTestEnvironment(t *testing.T) (TinyProxyTestEnvironment, error) {
	_ = "STUB: not implemented"
	return *new(TinyProxyTestEnvironment), nil
}

func TestProtocols(rootT *testing.T, testFunc func(t *testing.T, protocol clickhouse.Protocol)) {
	_ = "STUB: not implemented"
	return
}

func CleanupNativeConn(t *testing.T, conn driver.Conn) { _ = "STUB: not implemented"; return }

func OptionsToDSN(o *clickhouse.Options) string { _ = "STUB: not implemented"; return "" }

func ResetRandSeed() { _ = "STUB: not implemented"; return }

func Runtime(m *testing.M, ts string) (exitCode int) { _ = "STUB: not implemented"; return 0 }

//nolint
