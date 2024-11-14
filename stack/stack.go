package stack

import "sync"

const defaultStackCap = 100

type Stack[T comparable] struct {
	cap    int
	values []T
	mu     *sync.Mutex
}

func NewStack[T comparable](opts ...Option[T]) Stack[T] {
	s := Stack[T]{
		mu:  &sync.Mutex{},
		cap: defaultStackCap,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.values = make([]T, 0, s.cap)
	return s
}

type Option[T comparable] func(stack Stack[T])

func WithCap[T comparable](cap int) Option[T] {
	return func(s Stack[T]) {
		if cap > 0 {
			s.cap = cap
		}
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
