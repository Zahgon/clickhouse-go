package column

import (
	"github.com/ClickHouse/ch-go/proto"
)

func (c *JSON) encodeObjectHeader_v1(buffer *proto.Buffer) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *JSON) encodeObjectData_v1(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

// SharedData per row, empty for now.

func (c *JSON) decodeObjectHeader_v1(reader *proto.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *JSON) decodeObjectData_v1(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

// SharedData per row, ignored for now. May cause stream offset issues if present
// one UInt64 per row
