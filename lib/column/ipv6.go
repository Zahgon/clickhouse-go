package column

import (
	"net"
	"net/netip"
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

type IPv6 struct {
	col  proto.ColIPv6
	name string
}

func (col *IPv6) Reset() { _ = "STUB: not implemented"; return }

func (col *IPv6) Name() string { _ = "STUB: not implemented"; return "" }

func (col *IPv6) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *IPv6) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *IPv6) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *IPv6) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *IPv6) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func strToIPV6(strIp string) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

func (col *IPv6) AppendV6IPs(ips []netip.Addr) { _ = "STUB: not implemented"; return }

func (col *IPv6) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *IPv6) AppendRow(v any) (err error) { _ = "STUB: not implemented"; return nil }

func (col *IPv6) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *IPv6) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

func IPv6ToBytes(ip net.IP) [16]byte { _ = "STUB: not implemented"; return nil }

// TODO: This should probably return an netip.Addr
func (col *IPv6) row(i int) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func (col *IPv6) rowAddr(i int) netip.Addr { _ = "STUB: not implemented"; return *new(netip.Addr) }

var _ Interface = (*IPv6)(nil)
