package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	CreatedAt time.Time
	val       []byte
}

type Cache struct {
	CacheMap map[string]cacheEntry
	mu       *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	var cache Cache
	cache.CacheMap = make(map[string]cacheEntry)
	cache.mu = &sync.Mutex{}

	go cache.reapLoop(interval)
	return cache
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case t := <-ticker.C:
			for key, val := range c.CacheMap {
				if t.Sub(val.CreatedAt) >= interval {
					c.mu.Lock()
					delete(c.CacheMap, key)
					c.mu.Unlock()
				}
			}
		}
	}
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	currTime := time.Now()

	c.CacheMap[key] = cacheEntry{
		val:       val,
		CreatedAt: currTime,
	}
	c.mu.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	v, ok := c.CacheMap[key]
	c.mu.Unlock()
	if !ok {
		return nil, false
	}
	return v.val, true
}
