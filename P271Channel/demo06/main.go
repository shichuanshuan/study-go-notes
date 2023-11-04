package main

import (
	"fmt"
)

func writerNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}

	close(intChan)
}

//
func primeNum(intChan, primeChan chan int, exitChan chan bool) {
	var flag bool
	for {
		num, ok := <-intChan
		if !ok { // intChan 收取不到
			break
		}
		flag = true // 假设是素数
		// 判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { // 说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			// 将这个数就放入到primeChan
			primeChan <- num
		}

	}
	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	// 这里我们还不能关闭 primeChan
	// 向 exitChan 写入true
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000)
	exitChan := make(chan bool, 4)

	// 开启一个协程，向 intChan 放入 1-8000个数
	go writerNum(intChan)

	// 开启四个协程，从intChan取出数据，并判断是否为素数，如果是
	// 就放入到primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 这里我们主线程，进行处理
	// 直接
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		// 当我们从exitChan 取出来4个结果， 就可以放心的关闭 primeChan
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("素数=", res)
	}
}
