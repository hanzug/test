package main

import (
	"fmt"
	"sync"
	"time"
)

/*
sync.Cond是一个条件变量，它是用来等待或者宣布事件发生的线程间的通信机制。条件变量总是与一个互斥锁（sync.Mutex或sync.RWMutex）一起使用。

sync.Cond有三个主要的方法：

Wait(): 这个方法会释放与条件变量绑定的互斥锁，然后将调用者（goroutine）阻塞，直到被Signal()或Broadcast()唤醒。唤醒后，Wait()会在返回前重新获取互斥锁。

Signal(): 这个方法会唤醒一个等待（被Wait()阻塞）的goroutine。如果没有等待的goroutine，调用Signal()没有效果。

Broadcast(): 这个方法会唤醒所有等待的goroutine。如果没有等待的goroutine，调用Broadcast()没有效果。
*/

// 一个长度不超过2的队列
func main() {
	var m sync.Mutex                    // 创建一个互斥锁
	c := sync.NewCond(&m)               // 使用互斥锁创建一个新的条件变量
	queue := make([]interface{}, 0, 10) // 创建一个队列

	// removeFromQueue 是一个函数，它会在延迟一段时间后从队列中移除元素
	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)                 // 等待一段时间
		c.L.Lock()                        // 获取互斥锁
		queue = queue[1:]                 // 移除队列中的第一个元素
		fmt.Println("Removed from queue") // 打印消息
		c.L.Unlock()                      // 释放互斥锁
		c.Signal()                        // 发送信号，唤醒一个等待的goroutine
	}

	for i := 0; i < 10; i++ {
		c.L.Lock() // 获取互斥锁
		for len(queue) == 2 {
			c.Wait() // 如果队列的长度等于2，就等待
		}
		fmt.Println("Adding to queue")      // 打印消息
		queue = append(queue, struct{}{})   // 向队列中添加一个元素
		go removeFromQueue(1 * time.Second) // 在一个新的goroutine中移除队列中的元素
		c.L.Unlock()                        // 释放互斥锁
	}
}
