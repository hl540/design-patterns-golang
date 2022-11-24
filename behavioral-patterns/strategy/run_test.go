package strategy

import "testing"

func Test1(t *testing.T) {
	lfu := &Lfu{}

	// 初始化缓存，设置默认清理策略
	cache := initCache(lfu)

	// 提交元素
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	// 更改清理策略
	cache.setEvictionAlgo(&Lru{})
	cache.add("d", "4")

	// 更改清理策略
	cache.setEvictionAlgo(&Fifo{})
	cache.add("e", "5")
}
