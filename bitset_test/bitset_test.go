package bitset_test

import (
	"testing"

	"github.com/umahmood/bitset"
)

func TestNewReturnsValidBitset(t *testing.T) {
	b := bitset.New(100)
	if b == nil {
		t.Errorf("new: error in creating bitset instance.")
	}
}

func TestSet(t *testing.T) {
	in := []uint64{0, 5, 63, 64, 99, 127}
	want := true
	b := bitset.New(100)
	for _, v := range in {
		b.Set(v)
		got := b.Test(v)
		if got != want {
			t.Errorf("set: bit not set b.set(%d) got %t want %t", v, got, want)
		}
	}
}

func TestReset(t *testing.T) {
	in := []uint64{0, 5, 63, 64, 99, 127}
	want := false
	b := bitset.New(100)
	for _, v := range in {
		b.Set(v)
		b.Reset(v)
		got := b.Test(v)
		if got != want {
			t.Errorf("reset: bit not set b.set(%d) got %t want %t", v, got, want)
		}
	}
}

func TestFlip(t *testing.T) {
	in := []uint64{0, 63, 64, 127}
	want := true
	b := bitset.New(100)
	for _, v := range in {
		b.Flip(v)
		got := b.Test(v)
		if got != want {
			t.Errorf("flip: bit not flipped b.flip(%d) got %t want %t", v, got, want)
		}
	}
}

func TestFlipAndThenUnFlip(t *testing.T) {
	in := []uint64{0, 63, 64, 127}
	want := false
	b := bitset.New(100)
	for _, v := range in {
		b.Flip(v) // flip
		b.Flip(v) // un-flip
		got := b.Test(v)
		if got != want {
			t.Errorf("flip: bit not unflipped b.flip(%d) got %t want %t", v, got, want)
		}
	}
}

func TestMethodTestForBitsSetToTrue(t *testing.T) {
	in := []uint64{0, 5, 63, 64, 99, 127}
	want := true
	b := bitset.New(100)
	for _, v := range in {
		b.Set(v)
		got := b.Test(v)
		if got != want {
			t.Errorf("test: bit not set b.test(%d) got %t want %t", v, got, want)
		}
	}
}

func TestMethodTestForBitsSetToFalse(t *testing.T) {
	in := []uint64{0, 5, 63, 64, 99, 127}
	want := false
	b := bitset.New(100)
	for _, v := range in {
		got := b.Test(v)
		if got != want {
			t.Errorf("test: bit not set b.test(%d) got %t want %t", v, got, want)
		}
	}
}

func TestAllSomeBitsSet(t *testing.T) {
	in := []uint64{0, 63, 64, 127}
	want := false
	b := bitset.New(100)
	for _, v := range in {
		b.Set(v)
	}
	got := b.All()
	if got != want {
		t.Errorf("test: all bits set b.All() got %t want %t", got, want)
	}
}

func TestAllNoBitsSet(t *testing.T) {
	want := false
	b := bitset.New(100)
	got := b.All()
	if got != want {
		t.Errorf("test: all bits set b.All() got %t want %t", got, want)
	}
}

func TestAllBitsSet(t *testing.T) {
	want := true
	b := bitset.New(100)
	var i uint64
	for ; i < b.Size(); i++ {
		b.Set(i)
	}
	got := b.All()
	if got != want {
		t.Errorf("test: all bits not set b.All() got %t want %t", got, want)
	}
}

func TestAnyWithBitsSet(t *testing.T) {
	want := true
	b := bitset.New(100)
	b.Set(0)
	b.Set(127)
	got := b.Any()
	if got != want {
		t.Errorf("test: no bits set b.Any() got %t want %t", got, want)
	}
}

func TestAnyNoBitsSet(t *testing.T) {
	want := false
	b := bitset.New(100)
	got := b.Any()
	if got != want {
		t.Errorf("test: bits set b.Any() got %t want %t", got, want)
	}
}

func TestNoneNoBitsSet(t *testing.T) {
	want := true
	b := bitset.New(100)
	got := b.None()
	if got != want {
		t.Errorf("test: bits set b.None() got %t want %t", got, want)
	}
}

func TestNoneSomeBitsSet(t *testing.T) {
	want := false
	b := bitset.New(100)
	b.Set(114)
	b.Set(127)
	got := b.None()
	if got != want {
		t.Errorf("test: no bits set b.None() got %t want %t", got, want)
	}
}

func TestTrueCount(t *testing.T) {
	in := []uint64{0, 10, 63, 64, 110, 127}
	var want uint64 = 5
	b := bitset.New(100)
	for _, v := range in {
		b.Set(v)
	}
	b.Reset(10)
	b.Reset(110)
	b.Reset(42) // 42nd bit not set.
	b.Flip(65)
	got := b.TrueCount()
	if got != want {
		t.Errorf("test: count mis-match b.TrueCount() got %d want %d", got, want)
	}
}

func TestCopySrcAndDstBitsAreExact(t *testing.T) {
	a := bitset.New(100)
	a.Set(0)
	a.Set(63)
	a.Set(64)
	a.Set(127)
	b := bitset.New(200)
	b.Copy(a)
	var i uint64
	for i = 0; i < a.Size(); i++ {
		want := a.Test(i)
		got := b.Test(i)
		if got != want {
			t.Errorf("copy: did not make an exact copy of source bitset, "+
				"position %d got %t want %t", i, got, want)
		}
	}
}

func TestCopySrcAndDstTrueCountAreExact(t *testing.T) {
	a := bitset.New(100)
	a.Set(0)
	a.Set(63)
	a.Set(64)
	a.Set(127)
	b := bitset.New(200)
	b.Copy(a)
	want := a.TrueCount()
	got := b.TrueCount()
	if got != want {
		t.Errorf("copy: true count mis-match got %d want %d", got, want)
	}
}

func TestCopySrcAndDstSizesAreExact(t *testing.T) {
	a := bitset.New(200)
	b := bitset.New(100)
	b.Copy(a)
	want := a.Size()
	got := b.Size()
	if got != want {
		t.Errorf("copy: size mis-match got %d want %d", got, want)
	}
}

func TestStringer(t *testing.T) {
	in := []uint64{0, 63, 64, 127}
	want := "1000000000000000000000000000000000000000000000000000000000000001" +
		"1000000000000000000000000000000000000000000000000000000000000001"
	b := bitset.New(100)
	for _, v := range in {
		b.Set(v)
	}
	got := b.String()
	if got != want {
		t.Errorf("string: bit pattern mis-match got %s want %s", got, want)
	}
}
