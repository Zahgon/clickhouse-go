/*
 * Go implementation of Google city hash (MIT license)
 * https://code.google.com/p/cityhash/
 *
 * MIT License http://www.opensource.org/licenses/mit-license.php
 *
 * I don't even want to pretend to understand the details of city hash.
 * I am only reproducing the logic in Go as faithfully as I can.
 *
 */

package cityhash102

const (
	k0 uint64 = 0xc3a5c85c97cb3127
	k1 uint64 = 0xb492b66fbe98f273
	k2 uint64 = 0x9ae16a3b2f90404f
	k3 uint64 = 0xc949d7c7509e6557

	kMul uint64 = 0x9ddfea08eb382d69
)

func fetch64(p []byte) uint64 { _ = "STUB: not implemented"; return 0 }

//return uint64InExpectedOrder(unalignedLoad64(p))

func fetch32(p []byte) uint32 { _ = "STUB: not implemented"; return 0 }

//return uint32InExpectedOrder(unalignedLoad32(p))

func rotate64(val uint64, shift uint32) uint64 { _ = "STUB: not implemented"; return 0 }

func rotate32(val uint32, shift uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func swap64(a, b *uint64) { _ = "STUB: not implemented"; return }

func swap32(a, b *uint32) { _ = "STUB: not implemented"; return }

func permute3(a, b, c *uint32) { _ = "STUB: not implemented"; return }

func rotate64ByAtLeast1(val uint64, shift uint32) uint64 { _ = "STUB: not implemented"; return 0 }

func shiftMix(val uint64) uint64 { _ = "STUB: not implemented"; return 0 }

type Uint128 [2]uint64

func (this *Uint128) setLower64(l uint64) { _ = "STUB: not implemented"; return }

func (this *Uint128) setHigher64(h uint64) { _ = "STUB: not implemented"; return }

func (this Uint128) Lower64() uint64 { _ = "STUB: not implemented"; return 0 }

func (this Uint128) Higher64() uint64 { _ = "STUB: not implemented"; return 0 }

func (this Uint128) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func hash128to64(x Uint128) uint64 {
	_ = "STUB: not implemented"
	// Murmur-inspired hashing.
	return 0
}

func hashLen16(u, v uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func hashLen16_3(u, v, mul uint64) uint64 {
	_ = "STUB: not implemented"
	// Murmur-inspired hashing.
	return 0
}

func hashLen0to16(s []byte, length uint32) uint64 { _ = "STUB: not implemented"; return 0 }

// This probably works well for 16-byte strings as well, but it may be overkill
func hashLen17to32(s []byte, length uint32) uint64 { _ = "STUB: not implemented"; return 0 }

func weakHashLen32WithSeeds(w, x, y, z, a, b uint64) Uint128 {
	_ = "STUB: not implemented"
	return *new(Uint128)
}

func weakHashLen32WithSeeds_3(s []byte, a, b uint64) Uint128 {
	_ = "STUB: not implemented"
	return *new(Uint128)
}

func hashLen33to64(s []byte, length uint32) uint64 { _ = "STUB: not implemented"; return 0 }

func CityHash64(s []byte, length uint32) uint64 { _ = "STUB: not implemented"; return 0 }

func CityHash64WithSeed(s []byte, length uint32, seed uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func CityHash64WithSeeds(s []byte, length uint32, seed0, seed1 uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func cityMurmur(s []byte, length uint32, seed Uint128) Uint128 {
	_ = "STUB: not implemented"
	return *new(Uint128)
}

// len <= 16

// len > 16

func CityHash128WithSeed(s []byte, length uint32, seed Uint128) Uint128 {
	_ = "STUB: not implemented"
	return *new(Uint128)
}

// We expect length >= 128 to be the common case.  Keep 56 bytes of state:
// v, w, x, y, and z.

// This is the same inner loop as CityHash64(), manually unrolled.

// If 0 < length < 128, hash up to 4 chunks of 32 bytes each from the end of s.

//TODO why not use origin_len ?

// At this point our 48 bytes of state should contain more than
// enough information for a strong 128-bit hash.  We use two
// different 48-byte-to-8-byte hashes to get a 16-byte final result.

func CityHash128(s []byte, length uint32) Uint128 { _ = "STUB: not implemented"; return *new(Uint128) }
