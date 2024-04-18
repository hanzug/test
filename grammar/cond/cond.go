package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	queue := make([]struct{}, 0)

	// 生产者
	go func() {
		for {
			// 生产一个元素
			time.Sleep(time.Second)
			mu.Lock()
			queue = append(queue, struct{}{}) // 假设生产的元素是一个空结构体
			fmt.Println("Produced an item")
			mu.Unlock()
			cond.Signal() // 通知消费者有新的元素可消费
		}
	}()

	// 消费者
	go func() {
		for {
			mu.Lock()
			for len(queue) == 0 {
				cond.Wait() // 等待队列中有元素
			}
			queue = queue[1:] // 消费一个元素
			fmt.Println("Consumed an item")
			mu.Unlock()
		}
	}()

	select {} // 阻塞主goroutine，不让程序退出
}
