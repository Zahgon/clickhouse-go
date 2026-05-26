package cityhash102

import (
	"hash"
)

type City64 struct {
	s []byte
}

var _ hash.Hash64 = (*City64)(nil)
var _ hash.Hash = (*City64)(nil)

func New64() hash.Hash64 { _ = "STUB: not implemented"; return *new(hash.Hash64) }

func (this *City64) Sum(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func (this *City64) Sum64() uint64 { _ = "STUB: not implemented"; return 0 }

func (this *City64) Reset() { _ = "STUB: not implemented"; return }

func (this *City64) BlockSize() int { _ = "STUB: not implemented"; return 0 }

func (this *City64) Write(s []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (this *City64) Size() int { _ = "STUB: not implemented"; return 0 }
