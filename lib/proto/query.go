package proto

import (
	"os"

	chproto "github.com/ClickHouse/ch-go/proto"
	"go.opentelemetry.io/otel/trace"
)

var (
	osUser      = os.Getenv("USER")
	hostname, _ = os.Hostname()
)

type Query struct {
	ID                       string
	ClientName               string
	ClientVersion            Version
	ClientTCPProtocolVersion uint64
	Span                     trace.SpanContext
	Body                     string
	QuotaKey                 string
	Settings                 Settings
	Parameters               Parameters
	Compression              bool
	InitialUser              string
	InitialAddress           string
}

func (q *Query) Encode(buffer *chproto.Buffer, revision uint64) error {
	_ = "STUB: not implemented"
	return nil

	// client_info
}

// settings

/* empty string is a marker of the end of setting */

/* empty string is a marker of the end of parameters */

func swap64(b []byte) { _ = "STUB: not implemented"; return }

func (q *Query) encodeClientInfo(buffer *chproto.Buffer, revision uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// initial_user
// initial_query_id
// initial_address

// initial_query_start_time_microseconds

// interface [tcp - 1, http - 2]

// https://github.com/ClickHouse/ClickHouse/issues/34369

// https://github.com/ClickHouse/ClickHouse/issues/34369

// collaborate_with_initiator
// count_participating_replicas
// number_of_current_replica

type Settings []Setting

type Setting struct {
	Key       string
	Value     any
	Important bool
	Custom    bool
}

const (
	settingFlagImportant = 0x01
	settingFlagCustom    = 0x02
)

func (s Settings) Encode(buffer *chproto.Buffer, revision uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Setting) encode(buffer *chproto.Buffer, revision uint64) error {
	_ = "STUB: not implemented"
	return nil
}

type Parameters []Parameter

type Parameter struct {
	Key   string
	Value string
}

func (s Parameters) Encode(buffer *chproto.Buffer, revision uint64) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Parameter) encode(buffer *chproto.Buffer, revision uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// encodes a field dump with an appropriate type format
// implements the same logic as in ClickHouse Field::restoreFromDump (https://github.com/ClickHouse/ClickHouse/blob/master/src/Core/Field.cpp#L312)
// currently, only string type is supported
func encodeFieldDump(value any) (string, error) { _ = "STUB: not implemented"; return "", nil }
