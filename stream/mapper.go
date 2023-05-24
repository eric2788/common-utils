package stream


func Map[T comparable, R comparable](iter Stream[T], f func(T) R) Stream[R] {
	var result []R
	for _, v := range iter.arr {
		result = append(result, f(v))
	}
	return Stream[R]{result}
}

func FlatMap[T comparable, R comparable](iter Stream[T], f func(T) []R) Stream[R] {
	var result []R
	for _, v := range iter.arr {
		result = append(result, f(v)...)
	}
	return Stream[R]{result}
}

func ToMap[T comparable, R comparable](iter Stream[R], f func(R) T) map[T]R {
	result := make(map[T]R)
	for _, v := range iter.arr {
		result[f(v)] = v
	}
	return result
}

