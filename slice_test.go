package iter

import (
	"testing"
	"io"
)

var g1 int

func BenchmarkSlice(b *testing.B) {
	b.StopTimer()
	sl := make([]int, b.N)
	b.StartTimer()

	var val int
	for _, v := range sl {
		val = v
	}
	g1 = val
}

var g2 int

func BenchmarkSliceIter(b *testing.B) {
	b.StopTimer()
	sl := make([]int, b.N)
	si := SliceIter(sl)
	b.StartTimer()

	var val int
	si.Iterate(func(v int) error {
		val = v
		return nil
	})
	g2 = val
}

var g3 int

func BenchmarkSlicePuller(b *testing.B) {
	b.StopTimer()
	sl := make([]int, b.N)
	si := SliceIter(sl)
	sp := Pull[int](si, 256)
	defer sp.Close()
	b.StartTimer()

	var val int
	for v, e := sp.Next(); e != io.EOF; v, e = sp.Next() {
		val = v
	}
	g3 = val
}

var g4 int

func BenchmarkSlicePullerUnbuffered(b *testing.B) {
	b.StopTimer()
	sl := make([]int, b.N)
	si := SliceIter(sl)
	sp := Pull[int](si, 0)
	defer sp.Close()
	b.StartTimer()

	var val int
	for v, e := sp.Next(); e != io.EOF; v, e = sp.Next() {
		val = v
	}
	g4 = val
}
