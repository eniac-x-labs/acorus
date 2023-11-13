package cache

import lru "github.com/hashicorp/golang-lru"

const (
	l1ToL2CacheSize = 1000
	l2ToL1CacheSize = 1000
)

var (
	l1ToL2Cache *lru.ARCCache
	l2ToL1Cache *lru.ARCCache
)

func init() {
	l1ToL2Cache, _ = lru.NewARC(l1ToL2CacheSize)
	l2ToL1Cache, _ = lru.NewARC(l2ToL1CacheSize)
}

func GetL1ToL2CacheSize() *lru.ARCCache {
	return l1ToL2Cache
}

func GetL2ToL1Cache() *lru.ARCCache {
	return l2ToL1Cache
}
