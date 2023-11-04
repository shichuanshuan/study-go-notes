package main

import (
	"log"
	"os"
)

// 自定义logger
// out 设置日志信息写入的目的地， prefix添加前缀， flag定义日志的属性（时间、文件等等）
// func New(out io.Writer, prefix string, flag int) *logger
func main() {
	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Lmicroseconds)
	logger.Println("这是自定义的logger记录日志")
}
