package ring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRing(t *testing.T) {
	r := NewRing[int]()

	r.Add(1)

	res := r.Get()
	assert.Len(t, res, 1)
	assert.Equal(t, []int{1}, res)

	r.Add(2)
	r.Add(90)

	res = r.Get()
	assert.Len(t, res, 3)
	assert.Equal(t, []int{1, 2, 90}, res)

	head := r.GetHead()
	assert.Equal(t, 1, head.value)
	assert.Equal(t, 2, head.next.value)
	assert.Equal(t, 90, head.next.next.value)
	assert.Equal(t, 1, head.next.next.next.value)

}
