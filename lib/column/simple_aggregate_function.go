package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

type SimpleAggregateFunction struct {
	base   Interface
	chType Type
	name   string
}

func (col *SimpleAggregateFunction) Reset() { _ = "STUB: not implemented"; return }

func (col *SimpleAggregateFunction) Name() string { _ = "STUB: not implemented"; return "" }

func (col *SimpleAggregateFunction) parse(t Type, sc *ServerContext) (_ Interface, err error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

func (col *SimpleAggregateFunction) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *SimpleAggregateFunction) ScanType() reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

func (col *SimpleAggregateFunction) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *SimpleAggregateFunction) Row(i int, ptr bool) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func (col *SimpleAggregateFunction) ScanRow(dest any, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *SimpleAggregateFunction) Append(v any) ([]uint8, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *SimpleAggregateFunction) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *SimpleAggregateFunction) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *SimpleAggregateFunction) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *SimpleAggregateFunction) ReadStatePrefix(reader *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *SimpleAggregateFunction) WriteStatePrefix(buffer *proto.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

var _ Interface = (*SimpleAggregateFunction)(nil)
var _ CustomSerialization = (*SimpleAggregateFunction)(nil)
