package column

import (
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

type Nothing struct {
	name string
	col  proto.ColNothing
}

func (col *Nothing) Reset() { _ = "STUB: not implemented"; return }

func (col Nothing) Name() string { _ = "STUB: not implemented"; return "" }

func (Nothing) Type() Type             { _ = "STUB: not implemented"; return *new(Type) }
func (Nothing) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }
func (Nothing) Rows() int              { _ = "STUB: not implemented"; return 0 }
func (Nothing) Row(int, bool) any      { _ = "STUB: not implemented"; return *new(any) }
func (Nothing) ScanRow(any, int) error { _ = "STUB: not implemented"; return nil }

func (Nothing) Append(any) ([]uint8, error) { _ = "STUB: not implemented"; return nil, nil }

func (col Nothing) AppendRow(any) error { _ = "STUB: not implemented"; return nil }

func (col Nothing) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (Nothing) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

var _ Interface = (*Nothing)(nil)
