package column

import (
	"net"
	"net/netip"
	"reflect"

	"github.com/ClickHouse/ch-go/proto"
)

type IPv4 struct {
	name string
	col  proto.ColIPv4
}

func (col *IPv4) Reset() { _ = "STUB: not implemented"; return }

func (col *IPv4) Name() string { _ = "STUB: not implemented"; return "" }

func (col *IPv4) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (col *IPv4) ScanType() reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

func (col *IPv4) Rows() int { _ = "STUB: not implemented"; return 0 }

func (col *IPv4) Row(i int, ptr bool) any { _ = "STUB: not implemented"; return *new(any) }

func (col *IPv4) ScanRow(dest any, row int) error { _ = "STUB: not implemented"; return nil }

func strToIPV4(strIp string) (netip.Addr, error) {
	_ = "STUB: not implemented"
	return *new(netip.Addr), nil
}

func (col *IPv4) AppendV4IPs(ips []netip.Addr) { _ = "STUB: not implemented"; return }

func (col *IPv4) Append(v any) (nulls []uint8, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (col *IPv4) AppendRow(v any) (err error) { _ = "STUB: not implemented"; return nil }

func (col *IPv4) Decode(reader *proto.Reader, rows int) error {
	_ = "STUB: not implemented"
	return nil
}

func (col *IPv4) Encode(buffer *proto.Buffer) { _ = "STUB: not implemented"; return }

// TODO: This should probably return an netip.Addr
func (col *IPv4) row(i int) net.IP { _ = "STUB: not implemented"; return *new(net.IP) }

func (col *IPv4) rowAddr(i int) netip.Addr { _ = "STUB: not implemented"; return *new(netip.Addr) }

func netIPToNetIPAddr(ip net.IP) netip.Addr { _ = "STUB: not implemented"; return *new(netip.Addr) }

var _ Interface = (*IPv4)(nil)
