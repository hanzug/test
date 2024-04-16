package main

import (
	"fmt"
	"sync"
	"time"
)

// WorkerPool 管理一组可以执行 Task 的工作线程
type WorkerPool struct {
	tasks chan func()
	wg    sync.WaitGroup
}

// NewWorkerPool 创建一个新的 WorkerPool
func NewWorkerPool(numWorkers int) *WorkerPool {
	pool := &WorkerPool{
		tasks: make(chan func(), 128),
	}
	pool.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go pool.worker()
	}
	return pool
}

// worker 是运行在单独 goroutine 中的函数，它从 tasks 通道接收并执行任务
func (p *WorkerPool) worker() {
	defer p.wg.Done()
	for task := range p.tasks {
		task()
	}
}

// Run 提交一个新的任务到线程池
func (p *WorkerPool) Run(task func()) {
	p.tasks <- task
}

// Shutdown 等待所有任务完成并关闭 WorkerPool
func (p *WorkerPool) Shutdown() {
	close(p.tasks)
	p.wg.Wait()
}

func main() {
	pool := NewWorkerPool(4) // 创建一个有 4 个工作线程的线程池

	for i := 0; i < 10; i++ {
		count := i
		pool.Run(func() {
			fmt.Printf("Executing task %d\n", count)
			time.Sleep(time.Second) // 模拟任务执行时间
		})
	}

	pool.Shutdown() // 关闭线程池并等待所有任务完成
}
