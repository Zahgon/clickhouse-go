package column

import (
	"github.com/ClickHouse/ch-go/proto"
)

type deprecatedDynamic struct {
	maxTypes  int
	typeNames []string
}

func (c *Dynamic) sortColumnsForEncoding() { _ = "STUB: not implemented"; return }

func (c *Dynamic) encodeHeader_v1(buffer *proto.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

// SharedVariant is implicitly present in Dynamic, do not append to type names

func (c *Dynamic) encodeData_v1(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (c *Dynamic) decodeHeader_v1(reader *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

// Re-sort after adding SharedVariant

// Reset to server's totalTypes

func (c *Dynamic) decodeData_v1(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}
