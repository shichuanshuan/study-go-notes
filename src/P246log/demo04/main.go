package main

import (
	"fmt"
	"log"
	"os"
)

// 配置日志输出位置
// func SetOutput(w io.writer)
func main() {
	logFile, err := os.OpenFile("F:/GOlang/src/GoCode/P246log/demo04/log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Open log file failed err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志")
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")
}