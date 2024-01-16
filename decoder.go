package iter

import (
	"io"
)

type Decoder interface {
	Decode(any) error
}

func DecoderIter[T any](d Decoder) *Iterator[T] {
	return &Iterator[T]{Iteratef: func(yield func(T) error) error {
		var x T
		for e := d.Decode(&x); e != io.EOF; e = d.Decode(&x) {
			if e != nil {
				return e
			}
			e = yield(x)
			if e != nil {
				return e
			}
		}
		return nil
	}}
}
