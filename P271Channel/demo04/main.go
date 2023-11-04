package main

import "fmt"

func main() {
	intChan := make(chan int, 100)

	// 遍历管道
	for i := 0; i < 100; i++ {
		intChan <- i * 2
	}

	// 遍历管道不能使用普通的 for
	// for i := 0; i < len(intChan);i++ {

	// }
	close(intChan)
	for v := range intChan {
		fmt.Println("v= ", v)
	}
}