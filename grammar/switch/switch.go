package _switch

import (
	"fmt"
	"time"
)

// switch 多分支中执行符合条件的分支

func main() {

	// Here's a basic `switch`.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one") // 默认每个case最后都会break掉。
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	//多个条件可以写在一个case中
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// 可以使用表达式计算的条件
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//甚至可以用变量类型作为条件
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
