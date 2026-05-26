package chcol

import (
	"database/sql/driver"
)

// JSONSerializer interface allows a struct to be manually converted to an optimized JSON structure instead of relying
// on recursive reflection.
// Note that the struct must be a pointer in order for the interface to be matched, reflection will be used otherwise.
type JSONSerializer interface {
	SerializeClickHouseJSON() (*JSON, error)
}

// JSONDeserializer interface allows a struct to load its data from an optimized JSON structure instead of relying
// on recursive reflection to set its fields.
type JSONDeserializer interface {
	DeserializeClickHouseJSON(*JSON) error
}

// ExtractJSONPathAs is a convenience function for asserting a path to a specific type.
// The underlying value is also extracted from its Dynamic wrapper if present.
// T cannot be a Dynamic, if you want a Dynamic simply use ExtractJSONPathAsDynamic.
func ExtractJSONPathAs[T any](o *JSON, path string) (T, bool) {
	_ = "STUB: not implemented"
	return *new(T), false
}

// ExtractJSONPathAsDynamic is a convenience function for asserting a path to a Dynamic.
// If the value is not a Dynamic, the value is wrapped in an untyped Dynamic with false returned.
func ExtractJSONPathAsDynamic(o *JSON, path string) (Dynamic, bool) {
	_ = "STUB: not implemented"
	return *new(Dynamic), false
}

// JSON represents a ClickHouse JSON type that can hold multiple possible types
type JSON struct {
	valuesByPath map[string]any
}

// NewJSON creates a new empty JSON value
func NewJSON() *JSON { _ = "STUB: not implemented"; return nil }

func (o *JSON) ValuesByPath() map[string]any { _ = "STUB: not implemented"; return nil }

func (o *JSON) SetValueAtPath(path string, value any) { _ = "STUB: not implemented"; return }

func (o *JSON) ValueAtPath(path string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// NestedMap converts the flattened JSON data into a nested structure
func (o *JSON) NestedMap() map[string]any { _ = "STUB: not implemented"; return nil }

// MarshalJSON implements the json.Marshaler interface
func (o *JSON) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Scan implements the sql.Scanner interface
func (o *JSON) Scan(value any) error { _ = "STUB: not implemented"; return nil }

// Value implements the driver.Valuer interface
func (o *JSON) Value() (driver.Value, error) {
	_ = "STUB: not implemented"
	return *new(driver.Value), nil
}
