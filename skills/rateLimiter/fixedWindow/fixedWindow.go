package main

import (
	"fmt"
	"sync"
	"time"
)

// FixedWindowLimiter 控制请求的频率
type FixedWindowLimiter struct {
	limit      int           // 窗口内允许的最大请求数
	count      int           // 当前窗口的请求数
	mutex      sync.Mutex    // 保护 count 的互斥锁
	start      time.Time     // 窗口开始时间
	windowSize time.Duration // 窗口大小
}

// NewFixedWindowLimiter 创建一个新的 FixedWindowLimiter
func NewFixedWindowLimiter(limit int, windowSize time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		limit:      limit,
		windowSize: windowSize,
		start:      time.Now(),
	}
}

// Allow 检查是否可以处理新的请求
func (l *FixedWindowLimiter) Allow() bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	now := time.Now()
	if now.Sub(l.start) >= l.windowSize {
		// 当前时间超过窗口大小，重置窗口
		l.start = now
		l.count = 0
	}

	if l.count < l.limit {
		// 如果当前窗口内的请求数未达到限制，允许请求
		l.count++
		return true
	}

	return false
}

func main() {
	limiter := NewFixedWindowLimiter(5, 10*time.Second) // 每10秒最多5个请求

	for i := 0; i < 20; i++ {
		if limiter.Allow() {
			fmt.Println("Request allowed:", i)
		} else {
			fmt.Println("Request denied:", i)
		}
		time.Sleep(1 * time.Second)
	}
}
