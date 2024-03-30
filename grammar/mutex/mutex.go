package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 是一个并发安全的计数器
type SafeCounter struct {
	mu    sync.Mutex // 互斥锁
	count int        // 共享资源：计数器
}

// Increment 增加计数器的值
func (s *SafeCounter) Increment() {
	s.mu.Lock()         // 上锁
	defer s.mu.Unlock() // 在函数返回时释放锁
	s.count++           // 对共享资源进行操作
}

// Decrement 减少计数器的值
func (s *SafeCounter) Decrement() {
	s.mu.Lock()         // 上锁
	defer s.mu.Unlock() // 在函数返回时释放锁
	s.count--           // 对共享资源进行操作
}

// Value 返回计数器的当前值
func (s *SafeCounter) Value() int {
	s.mu.Lock()         // 上锁
	defer s.mu.Unlock() // 在函数返回时释放锁
	return s.count      // 读取共享资源的值
}

func main() {
	// 创建一个 SafeCounter 实例
	counter := SafeCounter{}

	// 启动多个 goroutine 来并发地增加计数器的值
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Increment() // 调用增加方法
			}
		}()
	}

	// 启动多个 goroutine 来并发地减少计数器的值
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Decrement() // 调用减少方法
			}
		}()
	}

	// 等待一段时间，以确保所有 goroutine 都已完成
	time.Sleep(time.Second)

	// 输出计数器的当前值
	fmt.Println("Final Counter:", counter.Value()) // 应该输出 0
}
