package proto

import (
	chproto "github.com/ClickHouse/ch-go/proto"
)

type TableColumns struct {
	First  string
	Second string
}

func (t *TableColumns) Decode(reader *chproto.Reader, revision uint64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (t *TableColumns) String() string { _ = "STUB: not implemented"; return "" }
