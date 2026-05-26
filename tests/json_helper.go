package tests

import (
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/chcol"
)

var JSONTestDate, _ = time.Parse(time.RFC3339, "2024-12-13T02:09:30.123Z")

type TestStructAddress struct {
	Street  string `chType:"String"`
	City    string `chType:"String"`
	Country string `chType:"String"`
}

type TestStruct struct {
	Name   string
	Age    int64
	Active bool
	Score  float64

	Tags    []string
	Numbers []int64

	Address TestStructAddress

	KeysNumbers map[string]int64
	Metadata    map[string]any

	Timestamp time.Time `chType:"DateTime64(3)"`

	DynamicString chcol.Dynamic
	DynamicInt    chcol.Dynamic
	DynamicMap    chcol.Dynamic
}

// FastTestStruct is a distinctly separate type that implements clickhouse.JSONSerializer and clickhouse.JSONDeserializer
// The struct must be a separate type since the JSON column is unable to ignore the interface implementation.
type FastTestStruct struct {
	ts TestStruct
}

// SerializeClickHouseJSON implements clickhouse.JSONSerializer for faster struct appending
func (fts *FastTestStruct) SerializeClickHouseJSON() (*clickhouse.JSON, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeserializeClickHouseJSON implements clickhouse.JSONDeserializer for faster struct scanning
func (fts *FastTestStruct) DeserializeClickHouseJSON(obj *clickhouse.JSON) error {
	_ = "STUB: not implemented"
	return nil
}

func BuildTestJSONPaths() *chcol.JSON { _ = "STUB: not implemented"; return nil }

func BuildTestJSONStruct() TestStruct { _ = "STUB: not implemented"; return *new(TestStruct) }

func BuildFastTestJSONStruct() FastTestStruct {
	_ = "STUB: not implemented"
	return *new(FastTestStruct)
}
