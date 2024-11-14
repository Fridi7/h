package lru_cache

import (
	"container/list"
	"sync"
)

const defaultLRUCacheSize = 100

type LRUCache[T comparable] struct {
	values *list.List
	cache  map[T]*list.Element
	size   int
	mu     *sync.RWMutex
}

type item[T comparable] struct {
	key   T
	value any
}

func NewLRUCache[T comparable](opts ...Option[T]) LRUCache[T] {
	c := LRUCache[T]{
		mu:   &sync.RWMutex{},
		size: defaultLRUCacheSize,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.values = list.New()
	c.cache = make(map[T]*list.Element, c.size)

	return c
}

type Option[T comparable] func(stack LRUCache[T])

func WithCap[T comparable](size int) Option[T] {
	return func(s LRUCache[T]) {
		if size > 0 {
			s.size = size
		}
	}
}

func (c *LRUCache[T]) Add(k T, v any) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.cache[k]; ok {
		elem.Value.(*item[T]).value = v
		c.values.MoveToFront(elem)
		return
	}

	if c.values.Len() >= c.size {
		backElem := c.values.Back()

		c.values.Remove(backElem)
		delete(c.cache, backElem.Value.(*item[T]).key)
	}

	newElem := c.values.PushFront(&item[T]{key: k, value: v})
	c.cache[k] = newElem
}

func (c *LRUCache[T]) Get(key T) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, exist := c.cache[key]
	if exist {

		c.values.MoveToFront(v)
		return v.Value.(*item[T]).value, true
	}

	var nilValue T
	return nilValue, false
}
