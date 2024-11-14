package lru_cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	r := NewLRUCache[string](WithCap[string](3))
	assert.Equal(t, 0, r.values.Len())
	assert.Len(t, r.cache, 0)

	r.Add("key1", 1)
	assert.Equal(t, 1, r.values.Len())
	assert.Len(t, r.cache, 1)

	res, ok := r.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, 1, res)

	r.Add("key2", 2)
	r.Add("key3", 3)
	r.Add("key4", 4)

	assert.Equal(t, 3, r.values.Len())
	assert.Len(t, r.cache, 3)

	v, ok := r.Get("key1")
	assert.False(t, ok)
	v, ok = r.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, 2, v)
	r.Get("key3")
	assert.True(t, ok)
	assert.Equal(t, 3, v)
	r.Get("key4")
	assert.True(t, ok)
	assert.Equal(t, 4, v)
}
