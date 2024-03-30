package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count int32
	var wg sync.WaitGroup

	increment := func() {
		for {
			// 使用 CAS 操作进行安全地增加操作
			oldCount := atomic.LoadInt32(&count) // 获取当前值
			newCount := oldCount + 1
			if atomic.CompareAndSwapInt32(&count, oldCount, newCount) {
				// CAS 操作成功，跳出循环
				break
			}
			// 如果 CAS 操作失败，表示其他 goroutine 已经修改了 count 的值，将自动重试
		}
		wg.Done()
	}

	// 启动 100 个 goroutine 来增加计数器
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go increment()
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	fmt.Println("Count:", count) // 输出最终的计数值
}
