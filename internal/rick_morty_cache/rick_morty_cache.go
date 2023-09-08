package rickmortycache

import (
	"sync"
	"time"
)


type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	cache map[string]cacheEntry
	mux *sync.Mutex 
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val: value,
		createdAt: time.Now().UTC(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reaploop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	} 
}

func (c *Cache) reap(now time.Time, latest time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k , v := range c.cache {
		if v.createdAt.Before(now.Add(-latest)) {
			delete(c.cache, k)
		}
	}
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux: &sync.Mutex{},
	}

	go c.reaploop(interval)

	return c
}



