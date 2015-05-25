package bitset

import "strings"

const bitsPerWord uint64 = 6
const numOfBits uint64 = 64

type Bitset struct {
	set       []uint64
	size      uint64
	trueCount uint64
}

// New given the desired number of bits, returns a new instance of a bitset.
func New(bits uint64) *Bitset {
	n := (bits >> bitsPerWord) + 1
	return &Bitset{set: make([]uint64, n), size: n * numOfBits}
}

// wordIndex given a position calculates the index in the bitset.
func (b *Bitset) wordIndex(p uint64) uint64 {
	return (p >> bitsPerWord)
}

// posIndex given a position and a word index, locates the target bit in that
// word index.
func (b *Bitset) posIndex(p, n uint64) uint64 {
	return (p - (n * numOfBits))
}

// checkBounds determines if 'n' is greater than the number of bits in the
// bit set. if it is, the method panics.
func (b *Bitset) checkBounds(n uint64) {
	if n > b.size {
		panic("index out of bounds")
	}
}

// Size returns the number of bits that the bitset can hold. The size is a
// a multiple of 64.
func (b *Bitset) Size() uint64 {
	return b.size
}

// Set sets the bit at a given position to 1 (true), the method panics if the
// postion to set, is greater than the number of bits in the bitset.
func (b *Bitset) Set(p uint64) {
	b.checkBounds(p)
	n := b.wordIndex(p)
	b.set[n] |= (1 << b.posIndex(p, n))
	b.trueCount++
}

// Reset sets the bit at a given position to 0 (false), the method panics if
// the postion to reset, is greater than the number of bits in the bitset.
func (b *Bitset) Reset(p uint64) {
	b.checkBounds(p)
	n := b.wordIndex(p)
	if b.Test(p) {
		b.set[n] &= ^(1 << b.posIndex(p, n))
		b.trueCount--
	}
}

// Test returns the value of a bit at a given position. true if the bit is set
// to 1, or false if the bit is set to 0. The method panics if the postion to
// test, is greater than the number of bits in the bitset.
func (b *Bitset) Test(p uint64) bool {
	b.checkBounds(p)
	n := b.wordIndex(p)
	t := b.set[n] & (1 << b.posIndex(p, n))
	if t == 0 {
		return false
	}
	return true
}

// Flip inverts the value of a bit at a given position. the method panics if
// the postion to flip, is greater than the number of bits in the bitset.
func (b *Bitset) Flip(p uint64) {
	b.checkBounds(p)
	n := b.wordIndex(p)
	b.set[n] ^= (1 << b.posIndex(p, n))
	if b.Test(p) {
		b.trueCount++
	} else {
		b.trueCount--
	}
}

// All tests if all the bits in the bitset are set to true.
func (b *Bitset) All() bool {
	var i uint64 = 0
	for ; i < b.size; i++ {
		if !b.Test(i) {
			return false
		}
	}
	return true
}

// Any tests if any of the bits in the bitset are set to true.
func (b *Bitset) Any() bool {
	var i uint64 = 0
	for ; i < b.size; i++ {
		if b.Test(i) {
			return true
		}
	}
	return false
}

// None tests if none of the bits in the bitset are set to true.
func (b *Bitset) None() bool {
	var i uint64 = 0
	for ; i < b.size; i++ {
		if b.Test(i) {
			return false
		}
	}
	return true
}

// TrueCount returns the number of bits in the bitset set to 1 (true).
func (b *Bitset) TrueCount() uint64 {
	return b.trueCount
}

// Copy makes this bitset an exact independant copy of the argument bitset.
func (b *Bitset) Copy(a *Bitset) {
	b.set = nil
	b.set = make([]uint64, len(a.set))
	b.size = a.size
	b.trueCount = a.trueCount
	copy(b.set, a.set)
}

// String returns a string representation of the bitset. The length of the
// string will be equal to bitset.Size(). Matches the Stringer interface.
func (b *Bitset) String() string {
	n := ((b.size - 1) >> bitsPerWord) + 1
	s := make([]string, n*numOfBits)
	var i uint64
	var j uint64
	for ; i < n; i++ {
		x := b.set[i]
		for j = 0; j < 64; j++ {
			if (x & (1 << j)) > 0 {
				s = append(s, "1")
			} else {
				s = append(s, "0")
			}
		}
	}
	return strings.Join(s, "")
}
