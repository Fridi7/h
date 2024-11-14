package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	q := NewSet[int](WithCap[int](5))
	assert.Len(t, q.values, 0)

	q.Add(1, 2, 3)
	assert.Len(t, q.values, 3)

	q.Add(2, 7, 10, 11)
	assert.Len(t, q.values, 6)

	assert.Equal(t, map[int]struct{}{1: {}, 2: {}, 3: {}, 7: {}, 10: {}, 11: {}}, q.values)

	assert.True(t, q.IsExist(3))
	assert.False(t, q.IsExist(33))

	assert.False(t, q.Delete(77))
	assert.Equal(t, map[int]struct{}{1: {}, 2: {}, 3: {}, 7: {}, 10: {}, 11: {}}, q.values)

	assert.True(t, q.Delete(7))
	assert.Equal(t, map[int]struct{}{1: {}, 2: {}, 3: {}, 10: {}, 11: {}}, q.values)

	assert.False(t, q.IsExist(7))

	assert.ElementsMatch(t, []int{1, 2, 3, 10, 11}, q.GetAll())
}

func TestGetIntersection(t *testing.T) {
	s1 := NewSet[int](WithCap[int](5))
	s1.Add(1, 2, 3, 4, 5)

	s2 := NewSet[int](WithCap[int](5))
	s2.Add(4, 5, 6, 7, 8)

	intersection := GetIntersection(s1, s2)
	assert.ElementsMatch(t, []int{4, 5}, intersection.GetAll())
}

func TestGetSubtraction(t *testing.T) {
	s1 := NewSet[int](WithCap[int](5))
	s1.Add(1, 2, 3, 4, 5)

	s2 := NewSet[int](WithCap[int](5))
	s2.Add(4, 5, 6, 7, 8)

	subtraction := GetSubtraction(s1, s2)
	assert.ElementsMatch(t, []int{1, 2, 3}, subtraction.GetAll())
}

func TestGetUnion(t *testing.T) {
	s1 := NewSet[int](WithCap[int](5))
	s1.Add(1, 2, 3, 4, 5)

	s2 := NewSet[int](WithCap[int](5))
	s2.Add(4, 5, 6, 7, 8)

	union := GetUnion(s1, s2)
	assert.ElementsMatch(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, union.GetAll())
}
