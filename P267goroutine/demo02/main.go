package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获得当前系统cpu核数
	num := runtime.NumCPU()
	fmt.Println("couNum =", num)

	// 这里设置num-1的cpu运行go程序
	// num - 1 保留一个cpu
	runtime.GOMAXPROCS(num - 1)
	fmt.Println("OK")
}
