package ext

import (
	"fmt"
	"iter"
)

type Set[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{m: make(map[T]struct{})}
}

func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.m, v)
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set[T]) Size() int {
	return len(s.m)
}

func (s *Set[T]) Slice() []T {
	out := make([]T, 0, len(s.m))
	for v := range s.m {
		out = append(out, v)
	}
	return out
}

func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s.m {
			if !yield(v) {
				return
			}
		}
	}
}

func (s *Set[T]) FromSlice(values []T) {
	for _, v := range values {
		s.Add(v)
	}
}

func (s *Set[T]) Copy() *Set[T] {
	ns := New[T]()
	for v := range s.m {
		ns.Add(v)
	}
	return ns
}

func (set *Set[T]) Print() {
	fmt.Print("Set contents: [")
	first := true
	for v := range set.All() {
		if !first {
			fmt.Print(", ")
		}
		fmt.Print(v)
		first = false
	}
	fmt.Println("]")
}
