package column

import "github.com/ClickHouse/ch-go/proto"

// ColStrProvider defines provider of proto.ColStr
type ColStrProvider func(name string) proto.ColStr

// colStrProvider provide proto.ColStr for Column() when type is String
var colStrProvider ColStrProvider = defaultColStrProvider

// defaultColStrProvider defines sample provider for proto.ColStr
func defaultColStrProvider(string) proto.ColStr {
	_ = "STUB: not implemented"
	return *

	// issue: https://github.com/ClickHouse/clickhouse-go/issues/1164
	// WithAllocBufferColStrProvider allow pre alloc buffer cap for proto.ColStr
	//
	//	It is more suitable for scenarios where a lot of data is written in batches
	new(proto.ColStr)
}

func WithAllocBufferColStrProvider(cap int) { _ = "STUB: not implemented"; return }

// WithColStrProvider more flexible than WithAllocBufferColStrProvider, such as use sync.Pool
func WithColStrProvider(provider ColStrProvider) { _ = "STUB: not implemented"; return }
