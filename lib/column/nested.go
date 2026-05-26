package column

import (
	"github.com/ClickHouse/ch-go/proto"
)

type Nested struct {
	Interface
	name string
}

func (col *Nested) Reset() { _ = "STUB: not implemented"; return }

func asDDL(cols []namedCol) string { _ = "STUB: not implemented"; return "" }

func (col *Nested) parse(t Type, sc *ServerContext) (_ Interface, err error) {
	_ = "STUB: not implemented"
	return *new(Interface), nil
}

func nestedColumns(raw string) (columns []namedCol) { _ = "STUB: not implemented"; return nil }

func (col *Nested) ReadStatePrefix(reader *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *Nested) WriteStatePrefix(buffer *proto.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

var _ Interface = (*Nested)(nil)
