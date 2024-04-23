package main

import "fmt"

//func sum(s []int, c chan int) {
//	sum := 0
//	for _, v := range s {
//		sum += v
//	}
//	c <- sum
//}
//
//func ReadAndClose() {}
//
//func main() {
//	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//
//	c := make(chan int)
//
//	go sum(numbers[:len(numbers)/2], c)
//	go sum(numbers[len(numbers)/2:], c)
//
//	x, y := <-c, <-c
//
//	// 打印结果
//	fmt.Println("Sum of first half:", x)
//	fmt.Println("Sum of second half:", y)
//	fmt.Println("Total sum:", x+y)
//}

func main() {
	var ch chan int
	val := <-ch
	fmt.Print(val)
}
