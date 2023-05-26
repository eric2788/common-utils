package stream

type (
	MapStream[K comparable, V comparable] struct {
		m map[K]V
	}

	MapEntry[K comparable, V comparable] struct {
		Key   K
		Value V
	}
)

// chainable

func (s MapStream[K, V]) Filter(f func(K, V) bool) MapStream[K, V] {
	var newMap = make(map[K]V)
	for k, v := range s.m {
		if f(k, v) {
			newMap[k] = v
		}
	}
	return FromMap(newMap)
}

func (s MapStream[K, V]) Entries() Stream[MapEntry[K, V]] {
	var t []MapEntry[K, V]
	for k, v := range s.m {
		t = append(t, MapEntry[K, V]{k, v})
	}
	return From(t)
}

func (s MapStream[K, V]) Values() Stream[V] {
	var t []V
	for _, v := range s.m {
		t = append(t, v)
	}
	return From(t)
}

func (s MapStream[K, V]) Keys() Stream[K] {
	var t []K
	for k := range s.m {
		t = append(t, k)
	}
	return From(t)
}

func (s MapStream[K, V]) ForEach(f func(K, V)) {
	for k, v := range s.m {
		f(k, v)
	}
}

func (s MapStream[K, V]) Map(f func(K, V) (K, V)) MapStream[K, V] {
	var newMap = make(map[K]V)
	for k, v := range s.m {
		newK, newV := f(k, v)
		newMap[newK] = newV
	}
	return FromMap(newMap)
}

func (s MapStream[K, V]) FlatMap(f func(K, V) map[K]V) MapStream[K, V] {
	var newMap = make(map[K]V)
	for k, v := range s.m {
		for newK, newV := range f(k, v) {
			newMap[newK] = newV
		}
	}
	return FromMap(newMap)
}

func (s MapStream[K, V]) ToMap() map[K]V {
	return s.m
}

func FromMap[K comparable, V comparable](m map[K]V) MapStream[K, V] {
	return MapStream[K, V]{m}
}
