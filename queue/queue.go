package queue

import "sync"

const defaultQueueCap = 100

type Queue[T comparable] struct {
	values []T
	mu     sync.Mutex
}

func NewQueue[T comparable](c int) Queue[T] {
	if c == 0 {
		c = defaultQueueCap
	}

	return Queue[T]{
		values: make([]T, 0, c),
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
