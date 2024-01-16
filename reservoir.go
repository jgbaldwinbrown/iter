package iter

import (
	"math/rand"
)

func Sample[T any](it Iter[T], keep int, ra *rand.Rand) ([]T, error) {
	r := make([]T, 0, keep)
	i := 0
	e := it.Iterate(func (val T) error {
		if i < keep {
			r = append(r, val)
			i++
			return nil
		}

		choice := ra.Intn(i)
		if choice < keep {
			r[choice] = val
		}

		i++
		return nil
	}
	return r, e
}
