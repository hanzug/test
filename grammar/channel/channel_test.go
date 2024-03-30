package main

import (
	"fmt"
	"testing"
)

func TestReadAndClose(t *testing.T) {
	// 创建一个带缓冲区的 Channel，缓冲区大小为 2
	ch := make(chan int, 2)

	// 写入数据到 Channel
	ch <- 1
	ch <- 2

	// 读取 Channel 中的数据
	fmt.Println(<-ch) // 输出：1
	fmt.Println(<-ch) // 输出：2

	// 尝试再次读取数据，会导致程序阻塞，因为 Channel 已经空了

	// 关闭 Channel
	close(ch)

	// 尝试向已关闭的 Channel 写入数据，会引发 panic
	// ch <- 3 // uncomment this line to see panic

	// 尝试从已关闭的 Channel 读取数据，会得到该类型的零值，并且不会引发 panic
	fmt.Println(<-ch) // 输出：0
}
