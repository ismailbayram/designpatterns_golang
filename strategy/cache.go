package strategy

// Context

type cache struct {
	storage          map[string]string
	evictionStrategy evictionStrategy
	capacity         int
	maxCapacity      int
}

func initCache(e evictionStrategy) *cache {
	storage := make(map[string]string)
	return &cache{
		storage:          storage,
		evictionStrategy: e,
		capacity:         0,
		maxCapacity:      2,
	}
}

func (c *cache) setEvictionStrategy(e evictionStrategy) {
	c.evictionStrategy = e
}

func (c *cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) get(key string) {
	delete(c.storage, key)
}

func (c *cache) evict() {
	c.evictionStrategy.evict(c)
	c.capacity--
}
