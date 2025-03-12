package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt *time.Time
	val       []byte
}

type Cache struct {
	cache map[string]cacheEntry
	mutex sync.Mutex
}

func NewCache(duration time.Duration) *Cache {
	c := &Cache{
		cache: map[string]cacheEntry{},
		mutex: sync.Mutex{},
	}
	go c.reapLoop(duration)
	return c
}

func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)

	for _ = range ticker.C {
		c.mutex.Lock()
		for key := range c.cache {
			if c.cache[key].createdAt.Add(duration).Before(time.Now()) {
				delete(c.cache, key)
			}
		}
		c.mutex.Unlock()
	}

}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	now := time.Now()
	c.cache[key] = cacheEntry{
		createdAt: &now,
		val:       val,
	}
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	ce, ok := c.cache[key]
	if ok {
		val = ce.val
		now := time.Now()
		ce.createdAt = &now
	}
	return val, ok
}
