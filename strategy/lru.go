package strategy

import "fmt"

// Concrete Strategy

type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("LRU strategy is used")
}
