package main

import (
	"fmt"
	"os"
)

// 命令行参数解析
// 只获取命令行参数
func main() {
	// os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, value := range os.Args {
			fmt.Printf("args[%d] = %v\n", index, value)
		}
	}
}
