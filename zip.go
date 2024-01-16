package iter

import (
	"io"
)

type Tuple2[T, U any] struct {
	V1 T
	V2 U
}

func Zip2[T, U any](i1 Iter[T], i2 Iter[U]) *Iterator[Tuple2[T, U]] {
	return &Iterator[Tuple2[T, U]]{Iteratef: func(yield func(Tuple2[T, U]) error) error {
		p2 := Pull(i2, 0)
		defer p2.Close()
		e := i1.Iterate(func(t T) error {
			u, e := p2.Next()
			if e != nil {
				return e
			}
			return yield(Tuple2[T, U]{t, u})
		})
		if e == io.EOF {
			return nil
		}
		return e
	}}
}
