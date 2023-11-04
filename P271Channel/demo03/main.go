package main

import "fmt"

func main() {
	intChan := make(chan int, 100)
	intChan <- 100
	intChan <- 200

	// 管道关闭后，不能写数据，读数据可以
	close(intChan)
	// intChan <- 300  // 不能写数据

	n1 := <-intChan // 可以读数据
	fmt.Println("n1 = ", n1)
}