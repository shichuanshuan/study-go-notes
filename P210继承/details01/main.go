package main

import "fmt"

type Student struct{
	Name string
	age int
}

func (s *Student) GetName(){
	fmt.Println("A Name =", s.Name)
}

func (s *Student) Getage(){
	fmt.Println("A age =", s.Name)
}


//
type A struct{
	Student
	Name string
}

func (s *A) GetName(){
	fmt.Println("B Name =", s.Name)
}

func main(){

	// 结构体可以使用嵌套匿名结构体所有的字段和方法
	// 首字母大写或者小写的字段、方法，都可以使用
	// var a A
	// a.Name = "Tom"
	// a.age = 20
	// a.Student.GetName()
	// a.Student.Getage()
	// fmt.Println()


	// 匿名结构体字段访问可以简化
	// a.Name = "Mary"
	// a.age = 21
	// a.GetName()
	// a.Getage()


	// 当结构体和匿名结构体有相同的字段或方法时，
	// 编译器采用就近访问原则，如果希望访问匿名结构体的字段和方法，
	// 可以通过匿名结构体名来区分
	var a A

	a.Name = "Jack"
	a.age = 20
	a.GetName()
	a.Getage()
}
