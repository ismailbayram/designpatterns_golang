package strategy

func RunStrategy() {
	fifo := &fifo{}
	lru := &lru{}
	cache := initCache(fifo)

	cache.add("a", "b")
	cache.add("a", "b")
	cache.add("a", "b")

	cache.setEvictionStrategy(lru)
	cache.add("d", "4")
}
