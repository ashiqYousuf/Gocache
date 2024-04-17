package cache

import "sync"

// Cache is a basic in-memory key-value cache implementation.
// Mutex is for controlling concurrent access to the cache.
type Cache[K comparable, V any] struct {
	items map[K]V
	mu    sync.Mutex
}

func New[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		items: make(map[K]V),
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
}
