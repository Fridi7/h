package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack[int](WithCap[int](5))
	assert.Len(t, s.values, 0)

	s.Push(1, 2, 3)
	assert.Len(t, s.values, 3)

	s.Push(2, 7, 10)
	assert.Len(t, s.values, 6)

	assert.Equal(t, []int{1, 2, 3, 2, 7, 10}, s.values)

	allValues := append([]int{}, s.values...)
	for i := 1; i <= 6; i++ {
		v, ok := s.Pop()
		assert.True(t, ok)
		assert.Equal(t, allValues[len(allValues)-i], v)
	}

	_, ok := s.Pop()
	assert.False(t, ok)

	s.Push(777)
	assert.Len(t, s.values, 1)
	v, ok := s.Pop()
	assert.True(t, ok)
	assert.Equal(t, 777, v)
}
