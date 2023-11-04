//表示hello.go所在的包是main，在go中，每个包都必须归属一个包
package main

//表示引入一个包，包名为fmt，引入该包后可以使用fmt包的函数

//func是关键字，表示一个函数
// func main() {
// 	fmt.Println("hello\n") //shift+alt+向下
// }

//gp build -o mehello.exe hello.go

//格式 gofmt -w hello.go

// 栈区、堆区
// 基本数据类型一般分配在栈区，编译器存在一个逃逸分析
// 引用数据类型一般分配在堆区，编译器存在一个逃逸分析

// var XF_VERSION string
// var XF_BUILDDATE string
// var XF_COMMIT string

// type Stu struct{

// }

// func test(i int){
// 	var s Stu
// 	s.test(i)
// }

// func (s *Stu) test(i int){
// 	fmt.Println("Stu test..", i)
// }
