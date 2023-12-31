package main

import "fmt"

// 1. 接口本身不能直接创建实例
type AInterface interface{
	Say()
}

type Stu struct{

}

func (stu Stu)Say(){
	fmt.Println("1. stu Say()")
}


// 5. 只要是自定义类型，都可以实现接口，不仅仅是数据类型
type integer int

func (i integer)Say(){
	fmt.Println("5. integer Say i =", i)
}


// 6. 一个自定义类型可以实现多个接口
type BInterface interface{
	Hello()
}

type CInterface interface{
	Printf()
}


type Monster struct{

}

func (m Monster)Printf(){
	fmt.Println("6. Monster Printf...")
}

func (m Monster)Hello(){
	fmt.Println("6. Monster Hello...")
}

func main(){

	// 1. 接口本身不能直接创建实例
	// var a AInterface
	// a.Say()

	// 但可以指向一个实现了该接口的自定义类型的变量
	var stu Stu   // 结构体变量， 实现了 Say() 也就实现了AInterface
	var a AInterface = stu
	a.Say()
	stu.Say()
	fmt.Println()


	// 5. 只要是自定义类型，都可以实现接口，不仅仅是数据类型
	var i integer = 10
	var b AInterface = i
	b.Say()
	fmt.Println()


	// 6. 一个自定义类型可以实现多个接口
	// Monster 实现了AInterface 和 BInterface 
	var monster Monster
	var binterface BInterface = monster
	var cinterface CInterface = monster

	binterface.Hello()
	cinterface.Printf()
}
