package bitset

import "testing"

func TestWordIndex(t *testing.T) {
	in := []uint64{0, 64, 128, 192, 256, 320, 384, 448, 512, 576}
	want := []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := New(100)
	n := 0
	for n < len(in) {
		got := b.wordIndex(in[n])
		if got != want[n] {
			t.Errorf("wordIndex: incorrect index got %d want %d", got, want[n])
		}
		n++
	}
}

func TestPosIndex(t *testing.T) {
	check := func(got, want uint64) {
		if got != want {
			t.Errorf("posIndex: incorrect index got %d want %d", got, want)
		}
	}

	b := New(100)

	check(b.posIndex(0, 0), 0)
	check(b.posIndex(42, 0), 42)
	check(b.posIndex(63, 0), 63)

	check(b.posIndex(64, 1), 0)
	check(b.posIndex(110, 1), 46)
	check(b.posIndex(127, 1), 63)

	check(b.posIndex(128, 2), 0)
	check(b.posIndex(145, 2), 17)
	check(b.posIndex(191, 2), 63)

	check(b.posIndex(192, 3), 0)
	check(b.posIndex(242, 3), 50)
	check(b.posIndex(255, 3), 63)

	check(b.posIndex(256, 4), 0)
	check(b.posIndex(308, 4), 52)
	check(b.posIndex(319, 4), 63)
}

func TestCheckBoundsThrowsPanic(t *testing.T) {
	want := "index out of bounds"
	b := New(100)
	defer func() {
		if got := recover(); got != nil {
			switch got.(type) {
			case string:
				if got != want {
					t.Errorf("checkBounds: got \"%s\" want \"%s\"", got, want)
				}
			}
		}
	}()

	b.checkBounds(200)
}

func TestCheckBoundsDoesNotThrowPanic(t *testing.T) {
	var want error
	b := New(100)
	defer func() {
		if got := recover(); got != want {
			t.Errorf("checkBounds: got %T want %T", got, want)
		}
	}()

	b.checkBounds(0)
	b.checkBounds(63)
	b.checkBounds(127)
}

func TestSizeIsRoundedUpToMultipleOf64(t *testing.T) {
	var in uint64 = 100
	var want uint64 = 128

	b := New(in)
	got := b.Size()

	if got != want {
		t.Errorf("size: mis-match bitset(%d) got %d want %d", in, got, want)
	}
}
