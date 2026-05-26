package column

import (
	"reflect"
	"time"

	"github.com/ClickHouse/ch-go/proto"
)

type Time64 struct {
	chType Type
	name   string
	col    proto.ColTime64
}

func (col *Time64) Reset() { _ = "STUB: not implemented"; return }

func (col *Time64) Name() string { _ = "STUB: not implemented"; return "" }

func (col *Time64) parse(t Type) (_ Interface, err error) {
	_ = "STUB: not implemented"

	// if no precision is given say just Time64 (instead of Time64(3|6|9))
	// it is treated as 3 (milliseconds)
	return *new(Interface), nil
}

func (col *Time64) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *Time64) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *Time64) Precision() (int64, bool) { _ = "STUB: not implemented"; return 0, false }

func (col *Time64) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *Time64) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

// ScanRow implements column.Interface.
// It is used to read a single column value of the row and store in
// `dest` Go variable.
func (col *Time64) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

// Append implements column.Interface.
// It is used for columnar inserts. Insert multiple Go value for
// single ClickHouse Time64 type.
func (col *Time64) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default all zeros, meaning no null values

// AppendRow implements column.Interface.
// It is used to insert column value in a row.
// Converts Go type into ClickHouse type to be inserted.
func (col *Time64) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *Time64) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Time64) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *Time64) row(i int) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

func (col *Time64) parseTime(value string) (time.Duration, error) {
	_ = "STUB: not implemented"
	return *new(time.Duration), nil
}
