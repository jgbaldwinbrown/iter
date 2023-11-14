package iter

import (
	"golang.org/x/sync/errgroup"
)

func Cat[T any](its Iter[Iter[T]]) *Iterator[T] {
	return &Iterator[T]{Iteratef: func(yield func(T) error) error {
		return its.Iterate(func(it Iter[T]) error {
			return it.Iterate(yield)
		})
	}}
}

func ChanCat[T any](its Iter[Iter[T]], bufsize, threads int) *Iterator[T] {
	return &Iterator[T]{Iteratef: func(yield func(T) error) error {
		c := make(chan T, bufsize)
		f := func(val T) error {
			c <- val
			return nil
		}

		var g errgroup.Group
		g.SetLimit(threads)
		errc := make(chan error, 1)

		go func() {
			its.Iterate(func(it Iter[T]) error {
				g.Go(func() error {
					return it.Iterate(f)
				})
				return nil
			})
			err := g.Wait()
			close(c)
			errc <- err
		}()

		for val := range c {
			e := yield(val)
			if e != nil {
				return e
			}
		}

		return <-errc
	}}
}
