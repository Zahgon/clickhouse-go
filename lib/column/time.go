package column

import (
	"reflect"
	"time"

	"github.com/ClickHouse/ch-go/proto"
)

type Time struct {
	chType Type
	name   string
	col    proto.ColTime
}

func (col *Time) Reset() { _ = "STUB: not implemented"; return }

func (col *Time) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Time) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Time) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Time) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Time) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

// ScanRow implements column.Interface.
// It is used to read a single column value of the row and store in
// `dest` Go variable.
func (col *Time) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

// Append implements column.Interface.
// It is used for columnar inserts. Insert multiple Go value for
// single ClickHouse Time type.
func (col *Time) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default all zeros, meaning no null values

// AppendRow implements column.Interface.
// It is used to insert column value in a row.
// Converts Go type into ClickHouse type to be inserted.
func (col *Time) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *Time) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Time) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Time) row(i int) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (col *Time) parseTime(value string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}

// helpers

func parseDuration(value string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
