package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entry map[string]CacheEntry
	Mu    sync.Mutex
}

type CacheEntry struct {
	CreatedAt time.Time 
	Val       []byte 
}

func NewCache(interval time.Duration) *Cache {
	c := Cache {
		Entry: make(map[string]CacheEntry),
		Mu:    sync.Mutex{},
	}

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			<-ticker.C
			c.reapLoop(interval)
		}
	}()

	return &c 
}

func (c *Cache) Add(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	c.Entry[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	entry, ok := c.Entry[key]
	if !ok {
		return nil, false
	}

	return entry.Val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	for key, entry := range c.Entry {
		entryAliveTime := time.Since(entry.CreatedAt)
		if entryAliveTime >= interval {
			delete(c.Entry, key)
		}
	}
}