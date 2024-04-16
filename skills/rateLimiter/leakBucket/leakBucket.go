package main

import (
	"sync"
	"time"
)

// LeakBucket 漏桶结构
type LeakBucket struct {
	capacity int        // 桶的容量
	water    int        // 当前水量（即数据量）
	rate     int        // 漏水速率（每秒）
	lastLeak time.Time  // 上次漏水时间
	lock     sync.Mutex // 保证并发安全
}

// NewLeakBucket 创建一个新的漏桶
func NewLeakBucket(capacity, rate int) *LeakBucket {
	return &LeakBucket{
		capacity: capacity,
		water:    0,
		rate:     rate,
		lastLeak: time.Now(),
	}
}

// AddWater 向桶中添加水（数据）
func (lb *LeakBucket) AddWater(amount int) bool {
	lb.lock.Lock()
	defer lb.lock.Unlock()

	// 计算自上次漏水以来的漏水量
	now := time.Now()
	leakAmount := int(now.Sub(lb.lastLeak).Seconds()) * lb.rate
	lb.water -= leakAmount
	if lb.water < 0 {
		lb.water = 0
	}
	lb.lastLeak = now

	// 添加水到桶中
	if lb.water+amount > lb.capacity {
		return false // 如果超出容量，则拒绝添加
	}
	lb.water += amount
	return true
}

func main() {
	bucket := NewLeakBucket(10, 1) // 容量为10，每秒漏水速率为1

	// 模拟添加水（数据）
	for i := 0; i < 20; i++ {
		if bucket.AddWater(1) {
			println("Added water", i)
		} else {
			println("Bucket full, water", i, "not added")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
