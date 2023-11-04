package main

import "fmt"


// Golang中的方法是作用在指定的数据类型上的
// 即：和指定的数据类型绑定
type Person struct{
	Name string
}

type Dog struct{
	Name string
}

func (p Person) test(){
	fmt.Println("test() name= ", p.Name)
}

func main(){
	// 调用方法是值传递
	var p Person

	p.Name = "tom"
	p.test()

	// 不能直接调用，必须通过方法才能调用。
	// test()
	// var dog Dog
	// dog.test()

	// 

}