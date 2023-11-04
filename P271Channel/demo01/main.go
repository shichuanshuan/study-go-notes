package main

import "fmt"

// channel 快速入门案例
// 特点：
// 1.channel 本质就是一个数据结构-队列
// 2.数据是先进先出【FIFO】first in first out
// 3.线程安全，多goroutine访问时，不需要加锁，也就是说channel本身就是线程安全的
// 4.channel有类型的，一个string的channel只能存放string类型数据

func main() {

	// 1.创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	// 2.看看intChan是什么
	// 它的值是一个地址，地址指向管道队列
	fmt.Printf("1. 值=%v 地址=%v", intChan, &intChan)
	fmt.Println()

	// 3.向管道写入数据
	intChan <- 10
	var num int = 12
	intChan <- num

	// 注意我们给管道写数据的时候，不能超过容量,超过容量就会报错
	intChan <- 20
	// intChan <- 30

	// 4.看看管道的长度
	fmt.Printf("4. len = %d, cap = %d\n", len(intChan), cap(intChan))

	// 5.从管道中读取数据
	// 管道是先进先出
	var num2 int
	num2 = <-intChan
	fmt.Println("5. ", num2)

	// 6. 在没有使用协程的情况下，如果我们的管道数据已经全部取出来，再取就会报错；
}