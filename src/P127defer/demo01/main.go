package main

import "fmt"

// 在函数中，程序员经常需要创建资源（比如：数据库链接，文件句柄，锁等）
// 为了执行完毕后，及时释放资源，Go的设计者提供了defer（延时机制）
func sum(n1 int, n2 int)int{

	// 当执行到defer时，暂时不执行，会将defer后面的语句压缩到独立的栈（defer栈）
	// 当函数执行完毕，再从defer栈，按照先入后出的方式出栈，执行
	// defer栈和栈不是同一个
	defer fmt.Println("ok1 n1 =", n1)  // defer 3. ok1 n1 = 10
	defer fmt.Println("ok2 n2 =", n2)  // defer 2. ok2 n2 = 20

	res := n1 + n2
	fmt.Println("ok3 n3 =", res)        // defer 1. ok3 res = 30

	return res
}

func main(){
	res := sum(10, 20)

	fmt.Println("res = ", res)     // 4. res = 30
}