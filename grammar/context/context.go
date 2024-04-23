package main

import (
	"context"
	"fmt"
	"time"
)

// 模拟的服务操作函数
func doWork(ctx context.Context) {
	// 模拟一个长时间运行的操作，比如说，这里我们让它运行10秒
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Operation completed successfully.")
	case <-ctx.Done():
		// 如果context被取消（比如超时），则进入这个分支
		fmt.Println("Operation canceled:", ctx.Err())
	}
}

func main() {
	// 创建一个超时时间为5秒的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // 确保在main函数结束时取消context

	// 运行模拟的服务操作
	go doWork(ctx)

	// 等待足够的时间以观察结果（这里不是必需的，只是为了演示）
	time.Sleep(6 * time.Second) // 让主goroutine等待足够的时间以观察doWork的输出
}
