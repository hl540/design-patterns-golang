package strategy

// EvictionAlgo 缓存清理策略接口
type EvictionAlgo interface {
	evict(c *Cache)
}
