package stream


func MapTo[T comparable, R comparable](iter Stream[T], f func(T) R) Stream[R] {
	var result []R
	for _, v := range iter.arr {
		result = append(result, f(v))
	}
	return From(result)
}

func FlatMapTo[T comparable, R comparable](iter Stream[T], f func(T) []R) Stream[R] {
	var result []R
	for _, v := range iter.arr {
		result = append(result, f(v)...)
	}
	return From(result)
}

func ToMapStream[T comparable, K comparable, V comparable](iter Stream[T], f func(T) (K,V)) MapStream[K, V] {
	result := make(map[K]V)
	for _, p := range iter.arr {
		k, v := f(p)
		result[k] = v
	}
	return FromMap(result)
}

