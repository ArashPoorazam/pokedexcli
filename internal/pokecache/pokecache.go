package pokecache

import (
	"sync"
	"time"
)

// Cache -
type Cache struct {
	cacheMap map[string]cacheEntry
	mu       *sync.Mutex
}

// CacheEntry -
type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

// NewCache -
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cacheMap: make(map[string]cacheEntry),
		mu:       &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

// Add
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheMap[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		value:     val,
	}
}

// Get
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.cacheMap[key]
	return val.value, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, val := range c.cacheMap {
		if val.createdAt.Before(now.Add(-last)) {
			delete(c.cacheMap, key)
		}
	}
}
