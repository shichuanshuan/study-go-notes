package main

import (
	"fmt"
	"time"
)

func main(){
	// 时间和日期相关函数的使用方法

	// 1. 获取当前时间
	now := time.Now()
	fmt.Printf("now = %v type = %T\n", now, now)

	// 2. 通过now可以获取到年月日，时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", now.Month())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	// 方法1 格式化日期时间
	fmt.Printf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(),
	now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("当前年月日 %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(),
	now.Hour(), now.Minute(), now.Second())

	fmt.Printf("dateStr =%v \n", dateStr)


	// 方法2 格式化日期时间
	fmt.Printf(now.Format("2006.01.02 15:04:05"))
	fmt.Println()

	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()

	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()

	fmt.Println(now.Format("02"))
}
