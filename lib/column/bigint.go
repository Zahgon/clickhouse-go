package column

import (
	"math/big"
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

type BigInt struct {
	size   int
	chType Type
	name   string
	signed bool
	col    proto.Column
}

func (col *BigInt) Reset() { _ = "STUB: not implemented"; return }

func (col *BigInt) Name() string { _ = "STUB: not implemented"; return "" }

func (col *BigInt) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *BigInt) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *BigInt) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *BigInt) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *BigInt) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func (col *BigInt) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *BigInt) AppendRow(v any) error { _ = "STUB: not implemented"; return nil }

func (col *BigInt) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *BigInt) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func (col *BigInt) row(i int) *big.Int { _ = "STUB: not implemented"; return nil }

func (col *BigInt) append(v *big.Int) { _ = "STUB: not implemented"; return }

func bigIntToRaw(dest []byte, v *big.Int) { _ = "STUB: not implemented"; return }

func rawToBigInt(v []byte, signed bool) *big.Int {
	_ = "STUB: not implemented"
	// LittleEndian to BigEndian
	return nil
}

// [0] ^ will +1

// neg ^ will -1

func endianSwap(src []byte, not bool) { _ = "STUB: not implemented"; return }

var _ Interface = (*BigInt)(nil)
