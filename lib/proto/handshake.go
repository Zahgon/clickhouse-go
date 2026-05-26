package proto

import (
	"time"

	chproto "github.com/ClickHouse/ch-go/proto"
	"go.yaml.in/yaml/v3"
)

type ClientHandshake struct {
	ProtocolVersion uint64

	ClientName    string
	ClientVersion Version
}

func (h ClientHandshake) Encode(buffer *chproto.Buffer) { _ = "STUB: not implemented"; return }

func (h ClientHandshake) String() string { _ = "STUB: not implemented"; return "" }

type ServerHandshake struct {
	Name        string
	DisplayName string
	Revision    uint64
	Version     Version
	Timezone    *time.Location
}

type Version struct {
	Major uint64
	Minor uint64
	Patch uint64
}

func ParseVersion(v string) (ver Version) { _ = "STUB: not implemented"; return *new(Version) }

func CheckMinVersion(constraint Version, version Version) bool {
	_ = "STUB: not implemented"
	return false
}

func (srv *ServerHandshake) Decode(reader *chproto.Reader) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (srv ServerHandshake) String() string { _ = "STUB: not implemented"; return "" }

func (v Version) String() string { _ = "STUB: not implemented"; return "" }

func (v *Version) UnmarshalYAML(value *yaml.Node) (err error) {
	_ = "STUB: not implemented"
	return nil
}
