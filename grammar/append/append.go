package main

import "fmt"

/*
如果slice的容量足以容纳新增的元素，那么append会在原地修改切片（即地址不变）；
如果容量不足，append会自动分配一个新的底层数组，并返回一个新的切片指针。这意味着返回的切片可能与原切片不同。
*/

func main() {
	// 创建一个初始容量为4的整型切片
	slice := make([]int, 0, 4)
	fmt.Printf("Initial slice: len=%d, cap=%d, %v\n", len(slice), cap(slice), slice)

	// 向切片中添加元素，直到等于其初始容量
	for i := 1; i <= 4; i++ {
		slice = append(slice, i)
		fmt.Printf("After appending %d: len=%d, cap=%d, %v\n", i, len(slice), cap(slice), slice)
	}

	// 创建slice的副本，然后再次使用append
	newSlice := append(slice, 6)
	fmt.Printf("After appending to newSlice: len=%d, cap=%d, %v\n", len(newSlice), cap(newSlice), newSlice)

	// 修改newSlice的第一个元素，并观察slice的变化
	newSlice[0] = 100
	fmt.Printf("After modifying newSlice: %v\n", newSlice)
	fmt.Printf("Original slice remains unchanged: %v\n", slice)
}
