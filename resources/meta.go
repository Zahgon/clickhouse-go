package resources

import (
	_ "embed"

	"go.yaml.in/yaml/v3"

	"github.com/ClickHouse/clickhouse-go/v2/lib/proto"
)

type Meta struct {
	ClickhouseVersions []proto.Version `yaml:"clickhouse_versions"`
	GoVersions         []proto.Version `yaml:"go_versions"`
	hVersion           proto.Version
}

//go:embed meta.yml
var metaFile []byte
var ClientMeta Meta

func init() {
	if err := yaml.Unmarshal(metaFile, &ClientMeta); err != nil {
		panic(err)
	}
	ClientMeta.hVersion = ClientMeta.findGreatestVersion()
}

func (m *Meta) IsSupportedClickHouseVersion(v proto.Version) bool {
	_ = "STUB: not implemented"
	return false
}

// check our patch is greater

func (m *Meta) SupportedVersions() string { _ = "STUB: not implemented"; return "" }

func (m *Meta) findGreatestVersion() proto.Version {
	_ = "STUB: not implemented"
	return *new(proto.Version)
}
