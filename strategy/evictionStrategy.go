package strategy

// Strategy Interface

type evictionStrategy interface {
	evict(c *cache)
}
