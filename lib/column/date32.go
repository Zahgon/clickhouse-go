package column

import (
	"reflect"
	"time"

	"github.com/ClickHouse/ch-go/proto"
)

var (
	minDate32, _ = time.Parse("2006-01-02 15:04:05", "1900-01-01 00:00:00")
	maxDate32, _ = time.Parse("2006-01-02 15:04:05", "2299-12-31 00:00:00")
)

type Date32 struct {
	col      proto.ColDate32
	name     string
	location *time.Location
}

func (col *Date32) Reset() { _ = "STUB: not implemented"; return }

func (col *Date32) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Date32) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Date32) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Date32) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Date32) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *Date32) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *Date32) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *Date32) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *Date32) parseDate(value string) (datetime time.Time, err error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func (col *Date32) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Date32) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Date32) row(i int) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// proto.Date is normalized as time.Time with UTC timezone.
// We make sure Date return from ClickHouse matches server timezone or user defined location.

var _ Interface = (*Date32)(nil)
