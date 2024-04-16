package main

import (
	"fmt"
	"sync"
	"time"
)

// SlidingWindowLimiter 控制请求的频率
type SlidingWindowLimiter struct {
	mutex      sync.Mutex    // 保护请求时间戳列表的互斥锁
	requests   []time.Time   // 时间窗口内的所有请求时间戳
	limit      int           // 时间窗口内允许的最大请求数
	windowSize time.Duration // 时间窗口大小
}

// NewSlidingWindowLimiter 创建一个新的 SlidingWindowLimiter
func NewSlidingWindowLimiter(limit int, windowSize time.Duration) *SlidingWindowLimiter {
	return &SlidingWindowLimiter{
		limit:      limit,
		windowSize: windowSize,
	}
}

// Allow 检查是否可以处理新的请求
func (l *SlidingWindowLimiter) Allow() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	now := time.Now()
	// 清除窗口之前的请求
	cutoff := now.Add(-l.windowSize)
	startIndex := 0
	for i, t := range l.requests {
		if t.After(cutoff) {
			startIndex = i
			break
		}
	}
	l.requests = l.requests[startIndex:]

	if len(l.requests) < l.limit {
		// 如果当前窗口内的请求数未达到限制，允许请求
		l.requests = append(l.requests, now)
		return true
	}

	return false
}

func main() {
	limiter := NewSlidingWindowLimiter(5, 10*time.Second) // 每10秒最多5个请求

	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Println("Request allowed:", i)
		} else {
			fmt.Println("Request denied:", i)
		}
		time.Sleep(1 * time.Second)
	}
}
