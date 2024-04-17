package cache

import "time"

// Both Cache & TTLCache implement this interface
type Cacher interface {
	Set(key, value any, ttl time.Duration)
	Get(key any) (any, bool)
	Remove(key any)
	Pop(key any) (any, bool)
}

type RenderKeyValue struct {
	Key   any `json:"key"`
	Value any `json:"value"`
}

// item represents a cache item with a value and an expiration time.
type item[V any] struct {
	value  V
	expiry time.Time
}

// isExpired checks if the cache item has expired.
func (i item[V]) isExpired() bool {
	return time.Now().After(i.expiry)
}
