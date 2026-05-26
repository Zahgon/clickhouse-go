package proto

import (
	chproto "github.com/ClickHouse/ch-go/proto"
)

type ProfileInfo struct {
	Rows                      uint64
	Bytes                     uint64
	Blocks                    uint64
	AppliedLimit              bool
	RowsBeforeLimit           uint64
	CalculatedRowsBeforeLimit bool
}

func (p *ProfileInfo) Decode(reader *chproto.Reader, revision uint64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *ProfileInfo) String() string { _ = "STUB: not implemented"; return "" }
