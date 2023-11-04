package main

import (
	"fmt"
	"time"
)

/*
时间的常量
const{
	Nanosecond Duration = 1 纳秒
	Microsecond = 1000 * Nanosecond   // 微妙
	Millisecond = 1000 * Microsecond  // 毫秒
	Second      = 1000 * Millisecond  // 秒
	Minute      = 60   * Second       // 分钟
	Hour        = 60   * Minute       // 小时
}

常量的作用:在程序中可用于获取指定时间单位的时间，比如想得到0.1毫秒
100 * time.Microsecond
*/

func main(){

	// 休眠
	// 需求  每隔1秒钟打印一个数字，打印到100时就退出
	// 需求2 每隔0.1秒中打印一个数字，打印到100时就退出
	i := 0

	for {
		i++
		// fmt.Println(i)

		// 每个一秒打印数字
		// time.Sleep(time.Second)

		// 每个0.1秒打印数字
		// 不能写成time.Sleep(time.Second * 0.1)
		time.Sleep(time.Millisecond * 100)

		if i == 100 {
			break;
		}
	}



	// Unix和UnixNano时间戳的使用
	// 作用：可产生随机数
	fmt.Printf("Unix =%v NanoUnix = %v\n", time.Now().Unix(), time.Now().UnixNano())
}