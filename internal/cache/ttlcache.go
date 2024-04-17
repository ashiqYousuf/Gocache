package cache

import (
	"log"
	"sync"
	"time"
)

// TTLCache is a generic cache implementation with support for time-to-live
type TTLCache[K comparable, V any] struct {
	items map[K]item[V]
	mu    sync.Mutex
}

// NewTTL creates a new TTLCache instance and starts a goroutine to periodically
// remove expired items every 5 seconds
func NewTTLCache[K comparable, V any]() *TTLCache[K, V] {
	log.Println("Using expiry caching system")

	c := &TTLCache[K, V]{
		items: make(map[K]item[V]),
	}

	go func() {
		for range time.Tick(5 * time.Second) {
			c.mu.Lock()

			for key, item := range c.items {
				if item.isExpired() {
					delete(c.items, key)
				}
			}

			c.mu.Unlock()
		}
	}()

	return c
}

// Set adds a new item to the cache
func (c *TTLCache[K, V]) Set(key K, value V, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = item[V]{
		value:  value,
		expiry: time.Now().Add(ttl),
	}
}

// Get retrieves the value associated with the given key from the cache.
func (c *TTLCache[K, V]) Get(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]
	if !found {
		return item.value, false
	}

	if item.isExpired() {
		delete(c.items, key)
		return item.value, false
	}

	return item.value, true
}

// Remove removes the item with the specified key from the cache.
func (c *TTLCache[K, V]) Remove(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}

// Pop removes and returns the item with the specified key from the cache.
func (c *TTLCache[K, V]) Pop(key K) (V, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, found := c.items[key]
	if !found {
		return item.value, false
	}

	delete(c.items, key)

	if item.isExpired() {
		return item.value, false
	}

	return item.value, true
}
