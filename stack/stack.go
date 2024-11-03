package stack

import "sync"

const defaultStackCap = 100

type Stack[T comparable] struct {
	values []T
	mu     sync.Mutex
}

func NewStack[T comparable](c int) Stack[T] {
	if c == 0 {
		c = defaultStackCap
	}

	return Stack[T]{
		values: make([]T, 0, c),
	}
}

func (s *Stack[T]) Push(values ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.values = append(s.values, values...)
}

func (s *Stack[T]) Pop() (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.values) == 0 {
		var nilValue T
		return nilValue, false
	}

	lastValue := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]

	return lastValue, true
}
