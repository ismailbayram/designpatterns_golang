package strategy

import "fmt"

// Concrete Strategy

type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("FIFO strategy is used.")
}
