package main

import (
	"log"
)

// 配置日志的前缀
func main() {
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[pprof]")
	log.Println("这是一条很普通的日志。")
}
