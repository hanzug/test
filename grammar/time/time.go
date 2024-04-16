package main

import (
	"fmt"
	"time"
)

func main() {
	// 获取当前时间
	now := time.Now()
	fmt.Println("当前时间:", now)

	// 格式化时间展示
	fmt.Println("格式化时间展示:", now.Format("2006-01-02 15:04:05 Monday"))

	// 创建特定时间
	specificTime := time.Date(2024, 3, 31, 20, 34, 58, 0, time.UTC)
	fmt.Println("特定时间:", specificTime)

	// 时间比较
	if now.Before(specificTime) {
		fmt.Println("当前时间早于特定时间")
	} else {
		fmt.Println("当前时间晚于或等于特定时间")
	}

	// 时间增加
	oneWeekLater := now.Add(7 * 24 * time.Hour)
	fmt.Println("一周后:", oneWeekLater)

	// 时间间隔 Duration
	duration := oneWeekLater.Sub(now)
	fmt.Println("时间间隔:", duration)

	// 解析字符串格式的时间
	parsedTime, _ := time.Parse("2006-01-02 15:04:05", "2024-03-31 20:34:58")
	fmt.Println("解析时间:", parsedTime)

	// 使用Ticker定时器
	ticker := time.NewTicker(5 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("Ticker at:", t)
		}
	}()

	// 使用Timer
	timer := time.NewTimer(10 * time.Second)
	<-timer.C
	fmt.Println("Timer expired")

	// 停止Ticker
	ticker.Stop()
	fmt.Println("Ticker stopped")

	// 计算两个时间的差值
	diff := specificTime.Sub(now)
	fmt.Println("两个时间的差值:", diff)

	// 获取一天的开始和结束时间
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24*time.Hour - 1)
	fmt.Println("今天开始时间:", startOfDay)
	fmt.Println("今天结束时间:", endOfDay)
}
