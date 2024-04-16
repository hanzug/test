package main

import "fmt"

func main() {

	// 给出固定容量的是数组
	var a [5]int
	fmt.Println("emp:", a)

	// 通过下标访问
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	//len函数给出长度
	fmt.Println("len:", len(a))

	//直接声明
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//编译器自动检查
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	//可以指定元素在数组中的下标
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	//循环访问
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	//二维数组初始化
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
