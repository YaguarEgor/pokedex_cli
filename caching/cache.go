package caching

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	caches map[string]CacheEntry
	mu     *sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		caches: map[string]CacheEntry{},
		mu:     &sync.Mutex{},
	}
	go c.reapLoop(interval)
	return &c
}

func (c Cache) reapLoop(interval time.Duration) {
	for {
		time.Sleep(interval)
		c.mu.Lock()
		for key, val := range c.caches {
			if time.Since(val.createdAt) >= interval {
				delete(c.caches, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.caches[key] = CacheEntry{
		createdAt: time.Now(),
		value:     val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.caches[key]
	if !ok {
		return []byte{}, false
	}
	return val.value, true
}
