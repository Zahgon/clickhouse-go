package proto

import (
	"github.com/ClickHouse/ch-go/proto"
)

type Exception struct {
	Code       int32
	Name       string
	Message    string
	StackTrace string
	Nested     []Exception
	nested     bool
}

func (e *Exception) Error() string { _ = "STUB: not implemented"; return "" }

func (e *Exception) Decode(reader *proto.Reader) (err error) { _ = "STUB: not implemented"; return nil }

func (e *Exception) decode(reader *proto.Reader) (err error) { _ = "STUB: not implemented"; return nil }
