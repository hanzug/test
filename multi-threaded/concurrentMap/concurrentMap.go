package concurrentMap

import "sync"

// ConcurrentMap 使用更新粒度的map提高并发
type ConcurrentMap struct {
	mps   []map[string]any
	seg   int
	locks []sync.RWMutex
	seed  uint32
}
