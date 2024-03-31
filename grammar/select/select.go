package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "来自通道1的数据"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "来自通道2的数据"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("收到:", msg1)
		case msg2 := <-channel2:
			fmt.Println("收到:", msg2)
			//default:
			//	fmt.Println("default")
		}
	}
}
