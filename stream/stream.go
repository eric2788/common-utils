package stream

import (
	mapset "github.com/deckarep/golang-set/v2"
	"math/rand"
	"time"
)

type Stream[T comparable] struct {
	arr    []T
	random rand.Rand
}

// chainable

func (s Stream[T]) Filter(predicate func(T) bool) Stream[T] {
	var result []T
	for _, v := range s.arr {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return From(result)
}

func (s Stream[T]) Distinct() Stream[T] {
	set := mapset.NewSet[T]()
	for _, v := range s.arr {
		set.Add(v)
	}
	return FromSet(set)
}

// Map only works for same type, if you want to change type, use MapTo
func (s Stream[T]) Map(f func(T) T) Stream[T] {
	var result []T
	for _, v := range s.arr {
		result = append(result, f(v))
	}
	return From(result)
}

// FlatMap only works for same type, if you want to change type, use FlatMapTo
func (s Stream[T]) FlatMap(f func(T) []T) Stream[T] {
	var result []T
	for _, v := range s.arr {
		result = append(result, f(v)...)
	}
	return From(result)
}

func (s Stream[T]) Shuffle() Stream[T] {
	var shuffled []T
	copy(shuffled, s.arr)
	rand.Shuffle(len(s.arr), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return From(shuffled)
}

// unchainable

func (s Stream[T]) Reduce(reducer func(T, T) T) T {
	var result T
	for _, v := range s.arr {
		result = reducer(result, v)
	}
	return result
}

func (s Stream[T]) Find(predicate func(T) bool) *T {
	for _, v := range s.arr {
		if predicate(v) {
			return &v
		}
	}
	return nil
}

func (s Stream[T]) FindLast(predicate func(T) bool) *T {
	for i := len(s.arr) - 1; i >= 0; i-- {
		if predicate(s.arr[i]) {
			return &s.arr[i]
		}
	}
	return nil
}

func (s Stream[T]) AllMatch(predicate func(T) bool) bool {
	for _, v := range s.arr {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func (s Stream[T]) AnyMatch(predicate func(T) bool) bool {
	for _, v := range s.arr {
		if predicate(v) {
			return true
		}
	}
	return false
}

func (s Stream[T]) NoneMatch(predicate func(T) bool) bool {
	for _, v := range s.arr {
		if predicate(v) {
			return false
		}
	}
	return true
}

func (s Stream[T]) Count(predicate func(T) bool) int {
	var count int
	for _, v := range s.arr {
		if predicate(v) {
			count++
		}
	}
	return count
}

// does not return anything

func (s Stream[T]) ForEach(consumer func(T)) {
	for _, v := range s.arr {
		consumer(v)
	}
}

func (s Stream[T]) ForEachIndexed(consumer func(int, T)) {
	for i, v := range s.arr {
		consumer(i, v)
	}
}

// to collection

func (s Stream[T]) ToArr() []T {
	return s.arr
}

func (s Stream[T]) ToSet() mapset.Set[T] {
	return mapset.NewSet(s.arr...)
}

func (s Stream[T]) ToMap(keySelector func(T) string) map[string]T {
	result := make(map[string]T)
	for _, v := range s.arr {
		result[keySelector(v)] = v
	}
	return result
}

func From[T comparable](arr []T) Stream[T] {
	return createStream(arr)
}

func FromSet[T comparable](set mapset.Set[T]) Stream[T] {
	return createStream(set.ToSlice())
}

func createStream[T comparable](arr []T) Stream[T] {
	r := rand.Rand{}
	r.Seed(time.Now().UnixNano())
	return Stream[T]{arr, r}
}
