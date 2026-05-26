package column

import (
	"reflect"
	"time"

	"github.com/ClickHouse/ch-go/proto"
)

var (
	minDate, _ = time.Parse("2006-01-02 15:04:05", "1970-01-01 00:00:00")
	maxDate, _ = time.Parse("2006-01-02 15:04:05", "2106-01-01 00:00:00")
)

const (
	defaultDateFormatNoZone   = "2006-01-02"
	defaultDateFormatWithZone = "2006-01-02 -07:00"
)

type Date struct {
	col      proto.ColDate
	name     string
	location *time.Location
}

func (col *Date) parse(t Type, tz *time.Location) (_ *Date, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Date) Reset() { _ = "STUB: not implemented"; return }

func (col *Date) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Date) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Date) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Date) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Date) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *Date) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *Date) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Date) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func parseDate(value string, minDate time.Time, maxDate time.Time, location *time.Location) (tv time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (col *Date) parseDate(value string) (tv time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (col *Date) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Date) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Date) row(i int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// proto.Date is normalized as time.Time with UTC timezone.
// We make sure Date return from ClickHouse matches server timezone or user defined location.

var _ Interface = (*Date)(nil)
