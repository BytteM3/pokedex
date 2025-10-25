package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheMap map[string]cacheEntry
	interval time.Duration
	mu       sync.Mutex
}

func (c *Cache) Add(key string, value []byte) {
	newCacheEntry := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheMap[key] = newCacheEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.cacheMap[key]
	if !ok {
		return nil, false
	}
	return value.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for k, v := range c.cacheMap {
			if now.Sub(v.createdAt) > c.interval {
				delete(c.cacheMap, k)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(itv time.Duration) *Cache {
	newMap := make(map[string]cacheEntry)
	newCache := Cache{
		cacheMap: newMap,
		interval: itv,
		mu:       sync.Mutex{},
	}
	go newCache.reapLoop()
	return &newCache
}
