package in_memory_cache

import "sync"

type Cache interface {
	Set(k, v string)
	Get(k string) (v string, ok bool)
}

type cache struct {
	cMap map[string]string
	mu   sync.RWMutex
}

func (c *cache) Set(k, v string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cMap[k] = v
}

func (c *cache) Get(k string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.cMap[k]

	return v, ok
}
