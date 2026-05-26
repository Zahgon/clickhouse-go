package column

import (
	"reflect"
	"time"

	"github.com/ClickHouse/ch-go/proto"
)

var (
	minDateTime, _ = time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
	maxDateTime, _ = time.Parse("2006-01-02 15:04:05", "2105-12-31 23:59:59")
)

const (
	defaultDateTimeFormatNoZone   = "2006-01-02 15:04:05"
	defaultDateTimeFormatWithZone = "2006-01-02 15:04:05 -07:00"
)

type DateTime struct {
	chType Type
	name   string
	col    proto.ColDateTime
}

func (col *DateTime) Reset() { _ = "STUB: not implemented"; return }

func (col *DateTime) Name() string { _ = "STUB: not implemented"; return "" }

func (col *DateTime) parse(t Type, tz *time.Location) (_ *DateTime, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *DateTime) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *DateTime) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *DateTime) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *DateTime) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *DateTime) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *DateTime) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil,

		// we assume int64 is in seconds and don't currently scale to the precision
		nil
}

func (col *DateTime) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

// we assume int64 is in seconds and don't currently scale to the precision

func (col *DateTime) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *DateTime) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *DateTime) row(i int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (col *DateTime) parseDateTime(value string) (tv time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

var _ Interface = (*DateTime)(nil)
