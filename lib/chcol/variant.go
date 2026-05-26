package chcol

import (
	"database/sql/driver"
)

// Variant represents a ClickHouse Variant type that can hold multiple possible types
type Variant struct {
	value  any
	chType string
}

// NewVariant creates a new Variant with the given value
func NewVariant(v any) Variant { _ = "STUB: not implemented"; return *new(Variant) }

// NewVariantWithType creates a new Variant with the given value and ClickHouse type
func NewVariantWithType(v any, chType string) Variant {
	_ = "STUB: not implemented"
	return *new(Variant)
}

// WithType creates a new Variant with the current value and given ClickHouse type
func (v Variant) WithType(chType string) Variant { _ = "STUB: not implemented"; return *new(Variant) }

// Type returns the ClickHouse type as a string.
func (v Variant) Type() string {
	_ = "STUB: not implemented"

	// HasType returns true if the value has a type ClickHouse included.
	return ""
}

func (v Variant) HasType() bool { _ = "STUB: not implemented"; return false }

// Nil returns true if the underlying value is nil.
func (v Variant) Nil() bool { _ = "STUB: not implemented"; return false }

// Any returns the underlying value as any.
func (v Variant) Any() any {
	_ = "STUB: not implemented"

	// Scan implements the sql.Scanner interface
	return *new(any)
}

func (v *Variant) Scan(value any) error { _ = "STUB: not implemented"; return nil }

// Value implements the driver.Valuer interface
func (v Variant) Value() (driver.Value, error) {
	_ = "STUB: not implemented"

	// MarshalJSON implements the json.Marshaler interface
	return *new(driver.Value), nil
}

func (v Variant) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON implements the json.Unmarshaler interface
func (v *Variant) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalText implements the encoding.TextMarshaler interface
func (v Variant) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalText implements the encoding.TextUnmarshaler interface
func (v *Variant) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }
