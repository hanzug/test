package main

import "fmt"

func main() {
	// 定义一个数组
	numbers := []int{1, 2, 3}

	// 安全执行函数
	safeExecute(func() {
		fmt.Println("访问数组元素:", numbers[5]) // 数组越界
	})

	// 安全执行函数
	safeExecute(func() {
		var pointer *int
		fmt.Println("解引用空指针:", *pointer) // 空指针引用
	})

	// 安全执行函数
	safeExecute(func() {
		fmt.Println("执行自定义错误条件")
		panic("自定义错误") // 触发自定义 panic
	})

	fmt.Println("程序继续执行")
}

// safeExecute 封装了错误捕获的逻辑
func safeExecute(f func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获到 panic:", r)
		}
	}()
	f()
}
