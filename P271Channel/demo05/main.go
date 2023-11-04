package main

import (
	"fmt"
	"time"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		data, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("read data =", data)
		time.Sleep(20 * time.Millisecond)
	}
	exitChan <- true

	close(exitChan)
}

// 综合应用
func main() {
	// 创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	for {
		ok := <-exitChan
		fmt.Println("1.ok", ok)
		if !ok {
			break
		}
	}
}