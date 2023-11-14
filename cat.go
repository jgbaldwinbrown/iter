package iter

func Cat[T any](its Iter[Iter[T]]) *Iterator[T] {
	return &Iterator[T]{Iteratef: func(yield func(T) error) error {
		return its.Iterate(func(it Iter[T]) error {
			return it.Iterate(yield)
		})
	}}
}
