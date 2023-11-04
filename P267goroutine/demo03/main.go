package main

import (
	"fmt"
	"sync"
	"time"
)

// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中
// 最后显示出来，要求使用goroutine完成

// 思路
// 1. 编写一个函数，来计算各个数的阶乘，并放入到map中。
// 2. 我们启动的协程多个，统计的将结果放入到map中
// 3. map 应该做出一个全局的。
var (
	myMay = make(map[int]int, 10)
	// lock 是一个全局的互斥锁
	// sync 是包： synchornized 同步
	// Mutex: 是互斥
	lock sync.Mutex
)

func test(n int) {
	res := 1

	// 加互斥锁
	lock.Lock()
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 解锁
	lock.Unlock()

	myMay[n] = res
}

func main() {

	for i := 1; i <= 6; i++ {
		go test(i)
	}

	// 要让主函数等待3秒，否则可能主函数执行完，
	// 而协程没有执行完, 从而输出空
	time.Sleep(2 * time.Second)

	for index, value := range myMay {
		fmt.Printf("map[%d] = %d\n", index, value)
	}
}
