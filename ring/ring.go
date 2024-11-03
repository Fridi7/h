package ring

type Ring[T comparable] struct {
	head *Record[T]
}

type Record[T comparable] struct {
	value T
	next  *Record[T]
}

func NewRing[T comparable]() Ring[T] {
	return Ring[T]{head: nil}
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
