package proto

import (
	"time"

	chproto "github.com/ClickHouse/ch-go/proto"
)

type Progress struct {
	Rows       uint64
	Bytes      uint64
	TotalRows  uint64
	WroteRows  uint64
	WroteBytes uint64
	Elapsed    time.Duration
	withClient bool
}

func (p *Progress) Decode(reader *chproto.Reader, revision uint64) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *Progress) String() string { _ = "STUB: not implemented"; return "" }
