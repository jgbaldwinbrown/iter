package iter

type Tuple2[T, U any] {
	V1 T
	V2 U
}

func Zip2[T, U any](i1 Iter[T], i2 Iter[U]) *Iterator[Pair[T, U]] {
	i2 := 
}
