package clickhouse_api

import (
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type FastProductPricing struct {
	Price    int64  `json:",omitempty"`
	Currency string `json:",omitempty"`
}

type FastProduct struct {
	ID        clickhouse.Dynamic `json:"id"`
	Name      string             `json:"name"`
	Tags      []string           `json:"tags"`
	Pricing   FastProductPricing `json:"pricing"`
	Metadata  map[string]any     `json:"metadata"`
	CreatedAt time.Time          `json:"created_at" chType:"DateTime64(3)"`
}

// SerializeClickHouseJSON implements clickhouse.JSONSerializer for faster struct appending
func (p *FastProduct) SerializeClickHouseJSON() (*clickhouse.JSON, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeserializeClickHouseJSON implements clickhouse.JSONDeserializer for faster struct scanning
func (p *FastProduct) DeserializeClickHouseJSON(obj *clickhouse.JSON) error {
	_ = "STUB: not implemented"
	return nil
}

func NewExampleFastProduct() *FastProduct { _ = "STUB: not implemented"; return nil }

func JSONFastStructExample() error { _ = "STUB: not implemented"; return nil }
