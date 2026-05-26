package clickhouse_api

import (
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

type ProductPricing struct {
	Price    int64  `json:",omitempty"`
	Currency string `json:",omitempty"`
}

type Product struct {
	ID        clickhouse.Dynamic     `json:"id"`
	Name      string                 `json:"name"`
	Tags      []string               `json:"tags"`
	Pricing   ProductPricing         `json:"pricing"`
	Metadata  map[string]interface{} `json:"metadata"`
	CreatedAt time.Time              `json:"created_at" chType:"DateTime64(3)"`
}

func NewExampleProduct() *Product { _ = "STUB: not implemented"; return nil }

func JSONStructExample() error { _ = "STUB: not implemented"; return nil }
