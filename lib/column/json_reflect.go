package column

import (
	"reflect"

	"github.com/ClickHouse/clickhouse-go/v2/lib/chcol"
)

// Decoding (Scanning)

// scanIntoStruct will iterate the provided struct and scan JSON data into the matching fields
func (c *JSON) scanIntoStruct(dest any, row int) error { _ = "STUB: not implemented"; return nil }

// scanIntoMap converts JSON data into a map
func (c *JSON) scanIntoMap(dest any, row int) error { _ = "STUB: not implemented"; return nil }

// fillStruct will iterate the provided struct and scan JSON data into the matching fields recursively
func (c *JSON) fillStruct(val reflect.Value, prefix string, row int) error {
	_ = "STUB: not implemented"
	return nil
}

// fillMap will iterate the provided map and scan JSON data in recursively
func (c *JSON) fillMap(val reflect.Value, prefix string, row int) error {
	_ = "STUB: not implemented"
	return nil
}

// splitter

// Encoding (Append, AppendRow)

// structToJSON converts a struct to JSON data
func structToJSON(v any) (*chcol.JSON, error) { _ = "STUB: not implemented"; return nil, nil }

// mapToJSON converts a map to JSON data
func mapToJSON(v any) (*chcol.JSON, error) { _ = "STUB: not implemented"; return nil, nil }

// iterateStruct recursively iterates through a struct and adds its fields to the JSON data
func iterateStruct(val reflect.Value, prefix string, json *chcol.JSON) error {
	_ = "STUB: not implemented"
	return nil
}

// handle `json:"name,omitempty"`

// iterateStructSkipTypes is a set of struct types that will not be iterated.
// Instead, the value will be assigned directly for use within Dynamic row appending.
var iterateStructSkipTypes = map[reflect.Type]struct{}{
	scanTypeIP:           {},
	scanTypeUUID:         {},
	scanTypeTime:         {},
	scanTypeTime:         {},
	scanTypeRing:         {},
	scanTypePoint:        {},
	scanTypeBigInt:       {},
	scanTypePolygon:      {},
	scanTypeDecimal:      {},
	scanTypeMultiPolygon: {},
	scanTypeVariant:      {},
	scanTypeDynamic:      {},
	scanTypeJSON:         {},
}

// handleValue processes a single value and adds it to the JSON data
func handleValue(val reflect.Value, path string, json *chcol.JSON, forcedType string) error {
	_ = "STUB: not implemented"
	return nil
}

// Only iterate maps if they are map[string]interface{}

const MaxMapPathDepth = 32

// iterateMap recursively iterates through a map and adds its values to the JSON data
func iterateMap(val reflect.Value, prefix string, json *chcol.JSON) error {
	_ = "STUB: not implemented"
	return nil
}
