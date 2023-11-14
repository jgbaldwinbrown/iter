package iter

import (
	"sort"
	"testing"
	"reflect"
)

func TestChanCat(t *testing.T) {
	in1 := []int{1000, 1000, 1000, 1000}
	in2 := []int{1000, 1000, 1000, 1000}

	it1 := make([]Iter[int], 0, len(in1))
	for _, val := range in1 {
		it1 = append(it1, IntIter(val))
	}
	itit1 := SliceIter(it1)
	c1 := Cat[int](itit1)
	out1, e := Collect[int](c1)
	if e != nil {
		panic(e)
	}
	sort.Slice(out1, func(i, j int) bool {
		return out1[i] < out1[j]
	})

	it2 := make([]Iter[int], 0, len(in2))
	for _, val := range in2 {
		it2 = append(it2, IntIter(val))
	}
	itit2 := SliceIter(it2)
	c2 := ChanCat[int](itit2, 1500, 3)
	out2, e := Collect[int](c2)
	if e != nil {
		panic(e)
	}
	sort.Slice(out2, func(i, j int) bool {
		return out2[i] < out2[j]
	})

	if !reflect.DeepEqual(out1, out2) {
		t.Errorf("out1 %v != out2 %v", out1, out2)
	}
}
