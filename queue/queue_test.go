package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	q := NewQueue[int](5)
	assert.Len(t, q.values, 0)

	q.Push(1, 2, 3)
	assert.Len(t, q.values, 3)

	q.Push(2, 7, 10)
	assert.Len(t, q.values, 6)

	assert.Equal(t, []int{1, 2, 3, 2, 7, 10}, q.values)

	allValues := append([]int{}, q.values...)
	for i := 0; i < 6; i++ {
		v, ok := q.Pop()
		assert.True(t, ok)
		assert.Equal(t, allValues[i], v)
	}

	_, ok := q.Pop()
	assert.False(t, ok)

	q.Push(777)
	assert.Len(t, q.values, 1)
	v, ok := q.Pop()
	assert.True(t, ok)
	assert.Equal(t, 777, v)
}
