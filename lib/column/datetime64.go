package column

import (
	"reflect"
	"time"

	"github.com/ClickHouse/ch-go/proto"
)

var (
	minDateTime64, _ = time.Parse("2006-01-02 15:04:05", "1900-01-01 00:00:00")
	maxDateTime64, _ = time.Parse("2006-01-02 15:04:05", "2262-04-11 23:47:16")
)

const (
	defaultDateTime64FormatNoZone   = "2006-01-02 15:04:05.999999999"
	defaultDateTime64FormatWithZone = "2006-01-02 15:04:05.999999999 -07:00"
)

type DateTime64 struct {
	chType   Type
	timezone *time.Location
	name     string
	col      proto.ColDateTime64
}

func (col *DateTime64) Reset() { _ = "STUB: not implemented"; return }

func (col *DateTime64) Name() string { _ = "STUB: not implemented"; return "" }

func (col *DateTime64) parse(t Type, tz *time.Location) (_ Interface, err error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

func (col *DateTime64) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *DateTime64) ScanType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (col *DateTime64) Precision() (int64, bool) { _ = "STUB: not implemented"; return 0, false }

func (col *DateTime64) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *DateTime64) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *DateTime64) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *DateTime64) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil,

		// we assume int64 is in milliseconds and don't currently scale to the precision - no tests to indicate intended
		// historical behaviour
		nil
}

func (col *DateTime64) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *DateTime64) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *DateTime64) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *DateTime64) row(i int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (col *DateTime64) timeToInt64(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }

func (col *DateTime64) parseDateTime(value string) (tv time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

var _ Interface = (*DateTime64)(nil)
