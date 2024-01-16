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
	})
	return r, e
}

func SamplePerc[T any](it Iter[T], keep float64, ra *rand.Rand) (*Iterator[T]) {
	return &Iterator[T]{Iteratef: func(yield func(T) error) error {
		return it.Iterate(func(x T) error {
			f := ra.Float64()
			if f <= keep {
				if e := yield(x); e != nil {
					return e
				}
			}
			return nil
		})
	}}
}
