package set

import (
	"sync"
)

const defaultSetCap = 100

type Set[T comparable] struct {
	cap    int
	values map[T]struct{}
	mu     *sync.RWMutex
}

func NewSet[T comparable](opts ...Option[T]) Set[T] {
	s := Set[T]{
		mu:  &sync.RWMutex{},
		cap: defaultSetCap,
	}

	for _, opt := range opts {
		opt(s)
	}

	s.values = make(map[T]struct{}, s.cap)
	return s
}

type Option[T comparable] func(stack Set[T])

func WithCap[T comparable](cap int) Option[T] {
	return func(s Set[T]) {
		if cap > 0 {
			s.cap = cap
		}
	}
}

func (s *Set[T]) Add(values ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, v := range values {
		s.values[v] = struct{}{}
	}
}

func (s *Set[T]) Delete(v T) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exist := s.values[v]; exist {
		delete(s.values, v)
		return true
	}

	return false
}

func (s *Set[T]) IsExist(key T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, exist := s.values[key]
	return exist
}

func (s *Set[T]) GetAll() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]T, 0, len(s.values))
	for v := range s.values {
		out = append(out, v)
	}

	return out
}

func GetIntersection[T comparable](s1, s2 Set[T]) Set[T] {
	values1 := s1.GetAll()
	res := NewSet[T](WithCap[T](len(values1)))

	for _, v := range values1 {
		if s2.IsExist(v) {
			res.Add(v)
		}
	}

	return res
}

func GetSubtraction[T comparable](s1, s2 Set[T]) Set[T] {
	values1 := s1.GetAll()
	res := NewSet[T](WithCap[T](len(values1)))

	for _, v := range values1 {
		if !s2.IsExist(v) {
			res.Add(v)
		}
	}

	return res
}

func GetUnion[T comparable](s1, s2 Set[T]) Set[T] {
	values1 := s1.GetAll()
	values2 := s2.GetAll()

	s := NewSet[T](WithCap[T](len(values1) + len(values2)))

	s.Add(values1...)
	s.Add(values2...)
	return s
}
