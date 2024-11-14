package queue

import "sync"

const defaultQueueCap = 100

type Queue[T comparable] struct {
	cap    int
	values []T
	mu     *sync.Mutex
}

func NewQueue[T comparable](opts ...Option[T]) Queue[T] {
	q := Queue[T]{
		mu:  &sync.Mutex{},
		cap: defaultQueueCap,
	}

	for _, opt := range opts {
		opt(q)
	}

	q.values = make([]T, 0, q.cap)
	return q
}

type Option[T comparable] func(stack Queue[T])

func WithCap[T comparable](cap int) Option[T] {
	return func(q Queue[T]) {
		if cap > 0 {
			q.cap = cap
		}
	}
}

func (q *Queue[T]) Push(values ...T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.values = append(q.values, values...)
}

func (q *Queue[T]) Pop() (T, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.values) == 0 {
		var nilValue T
		return nilValue, false
	}

	firstValue := q.values[0]
	q.values = q.values[1:]

	return firstValue, true
}
