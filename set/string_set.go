package set

import (
	"sync"
)

type StringSet struct {
	elements map[string]void
	mu       sync.Mutex
}

func (s *StringSet) Add(element string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.Contains(element) {
		return false
	}

	s.elements[element] = null
	return true
}

func (s *StringSet) Delete(element string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.Contains(element) {
		delete(s.elements, element)
		return true
	} else {
		return false
	}
}

func (s *StringSet) Contains(element string) bool {
	_, ok := s.elements[element]
	return ok
}

func (s *StringSet) Size() int {
	return len(s.elements)
}

func (s *StringSet) Iterator() <-chan string {
	s.mu.Lock()
	defer s.mu.Unlock()
	iterator := make(chan string, len(s.elements))
	for v, _ := range s.elements {
		iterator <- v
	}
	close(iterator)
	return iterator
}

func (s *StringSet) ToArr() []string {
	arr := make([]string, 0)
	for e := range s.Iterator() {
		arr = append(arr, e)
	}
	return arr
}

func NewString() *StringSet {
	return &StringSet{
		elements: make(map[string]void),
	}
}

func FromStrArr(arr []string) *StringSet {
	set := NewString()
	for _, e := range arr {
		set.Add(e)
	}
	return set
}
