package main

import "fmt"

func main() {
	defer fmt.Println("执行 defer 1")
	defer fmt.Println("执行 defer 2")
	defer fmt.Println("执行 defer 3")

	fmt.Println("函数主体")
}
