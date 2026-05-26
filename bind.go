package clickhouse

import (
	std_driver "database/sql/driver"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

var (
	ErrInvalidTimezone = errors.New("invalid timezone value")
)

func Named(name string, value any) driver.NamedValue {
	_ = "STUB: not implemented"
	return *new(driver.NamedValue)
}

type TimeUnit uint8

const (
	Seconds TimeUnit = iota
	MilliSeconds
	MicroSeconds
	NanoSeconds
)

type GroupSet struct {
	Value []any
}

type ArraySet []any

func DateNamed(name string, value time.Time, scale TimeUnit) driver.NamedDateValue {
	_ = "STUB: not implemented"
	return *new(driver.NamedDateValue)
}

var (
	bindNumericRe    = regexp.MustCompile(`\$[0-9]+`)
	bindPositionalRe = regexp.MustCompile(`[^\\][?]`)
)

func bind(tz *time.Location, query string, args ...any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func checkAllNamedArguments(args ...any) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func bindPositional(tz *time.Location, query string, args ...any) (_ string, err error) {
	_ = "STUB: not implemented"
	return "",

		// Position of previous match for copying
		nil
}

// Index for the argument at current position

// Number of positional arguments that couldn't be matched

// It's fine looping through the query string as bytes, because the (fixed) characters we're looking for
// are in the ASCII range to won't take up more than one byte.

// Copy all previous index to here characters

// Copy all previous index to here characters

// Append the argument value

// If there were no replacements, quick return without copying the string

// Append the remainder

func bindNumeric(tz *time.Location, query string, args ...any) (_ string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

var bindNamedRe = regexp.MustCompile(`@[a-zA-Z0-9\_]+`)

func bindNamed(tz *time.Location, query string, args ...any) (_ string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func formatTime(tz *time.Location, scale TimeUnit, value time.Time) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// It's required to pass timestamp as string due to decimal overflow for higher precision,
// but zero-value string "toDateTime('0')" will be not parsed by ClickHouse.

// Escape the timezone string (timezone may contain malicious SQL query)

var stringQuoteReplacer = strings.NewReplacer(`\`, `\\`, `'`, `\'`)

func format(tz *time.Location, scale TimeUnit, v any) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// map

func join[E any](tz *time.Location, scale TimeUnit, values []E) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func rebind(in []std_driver.NamedValue) []any { _ = "STUB: not implemented"; return nil }
