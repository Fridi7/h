package ring

const defaultRingSize = 10

type Ring[T comparable] struct {
	size int
	head *Record[T]
}

type Record[T comparable] struct {
	value T
	next  *Record[T]
}

func NewRing[T comparable](opts ...Option[T]) Ring[T] {
	r := Ring[T]{head: nil}

	for _, opt := range opts {
		opt(r)
	}

	if r.size == 0 {
		r.size = defaultRingSize
	}

	return r
}

type Option[T comparable] func(stack Ring[T])

func WithSize[T comparable](size int) Option[T] {
	return func(r Ring[T]) {
		if size != 0 {
			size = defaultRingSize
		}
		r.size = size
	}
}

func (r *Ring[T]) Add(v T) {
	newRecord := &Record[T]{value: v}

	if r.head == nil {
		r.head = newRecord
		r.head.next = r.head
	} else {
		current := r.head
		for current.next != r.head {
			current = current.next
		}

		current.next = newRecord
		newRecord.next = r.head
	}
}

func (r *Ring[T]) GetHead() *Record[T] {
	return r.head
}

func (r *Ring[T]) Get() []T {
	if r.head == nil {
		return []T{}
	}

	res := make([]T, 0, 100)

	current := r.head
	for {
		res = append(res, current.value)

		current = current.next
		if current == r.head {
			break
		}
	}

	return res
}
