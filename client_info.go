package clickhouse

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
)

const ClientName = "clickhouse-go"

const (
	ClientVersionMajor       = 2
	ClientVersionMinor       = 46
	ClientVersionPatch       = 0
	ClientTCPProtocolVersion = proto.DBMS_TCP_PROTOCOL_VERSION
)

type ClientInfo struct {
	Products []struct {
		Name    string
		Version string
	}

	Comment []string
}

// Append returns a new copy of the combined ClientInfo structs
func (a ClientInfo) Append(b ClientInfo) ClientInfo {
	_ = "STUB: not implemented"
	return *new(ClientInfo)
}

func (o ClientInfo) String() string { _ = "STUB: not implemented"; return "" }

func mapKeysInOrder[V any](m map[string]V) []string { _ = "STUB: not implemented"; return nil }
