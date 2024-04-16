package main

import (
	"sync"
	"time"
)

// TokenBucket 令牌桶结构
type TokenBucket struct {
	capacity  int           // 桶的容量
	tokens    int           // 当前令牌数量
	rate      time.Duration // 令牌填充速率
	lastCheck time.Time     // 上次检查时间
	lock      sync.Mutex    // 保证并发安全
}

// NewTokenBucket 创建一个新的令牌桶
func NewTokenBucket(capacity int, rate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:  capacity,
		tokens:    capacity, // 初始时桶满
		rate:      rate,
		lastCheck: time.Now(),
	}
}

// Take 尝试从桶中取出令牌
func (tb *TokenBucket) Take(n int) bool {
	tb.lock.Lock()
	defer tb.lock.Unlock()

	// 通过时间差计算令牌数量
	now := time.Now()
	elapsed := now.Sub(tb.lastCheck)
	tb.lastCheck = now
	newTokens := int(elapsed / tb.rate)
	if newTokens > 0 {
		tb.tokens = min(tb.capacity, tb.tokens+newTokens)
	}

	// 检查是否有足够的令牌
	if tb.tokens >= n {
		tb.tokens -= n
		return true
	}
	return false
}

// min 辅助函数，返回两个整数的最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	bucket := NewTokenBucket(10, time.Second) // 每秒填充一个令牌，容量为10

	// 模拟请求
	for i := 0; i < 20; i++ {
		if bucket.Take(1) {
			println("Request", i, "processed")
		} else {
			println("Request", i, "denied")
		}
		time.Sleep(100 * time.Millisecond)
	}
}
