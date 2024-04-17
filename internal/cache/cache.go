package cache

import (
	"log"
	"sync"
	"time"
)

// Cache is a basic in-memory key-value cache implementation.
// Mutex is for controlling concurrent access to the cache.
type Cache[K comparable, V any] struct {
	items map[K]V
	mu    sync.Mutex
}

// New creates a new Cache instance.
func NewCache[K comparable, V any]() *Cache[K, V] {
	log.Println("Using without expiry caching system")
	return &Cache[K, V]{
		items: make(map[K]V),
	}
}

// Set adds or updates a key-value pair in the cache.
func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
}

// Get retrieves the value associated with the given key from the cache
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, found := c.items[key]
	return value, found
}

// Remove deletes the key-value pair with the specified key from the cache.
func (c *Cache[K, V]) Remove(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

// Pop removes and returns the value associated with the specified key from the cache.
func (c *Cache[K, V]) Pop(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, found := c.items[key]
	if found {
		delete(c.items, key)
	}

	return value, found
}
