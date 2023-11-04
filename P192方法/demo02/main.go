package main

import "fmt"


type Person struct{
	Name string
}

// 方法快速入门
// 1. 给Person结构体添加speak方法，输出 xxx是好人
func (person Person) speak(){
	person.Name = "Tom"
	fmt.Printf("1. %v 是好人\n", person.Name)
}

// 2. 给Person结构体添加jisuan方法，可以计算从1+...+1000的结果
func (person Person) jisuan(){
	result := 0
	for i := 1; i <= 1000; i++{
		result += i
	}

	fmt.Println("2. reslut =", result)
}

// 3. 给Person结构体添加jisuan2方法，该方法可以接收一个数n，计算从1+...+1000的结果
func (person Person) jisuan2(n int){
	result := 0
	for i := 1; i <= n; i++{
		result += i
	}

	fmt.Println("3. reslut =", result)
}

// 4. 给Person结构体添加getSum方法，可以计算从两个数的和，并返回结果。
func (person Person) getSum(n1 int, n2 int)int{
	return n1 + n2
}

func main(){

	// 调用方法
	var p Person
	p.speak()
	fmt.Println()

	p.jisuan()
	fmt.Println()

	p.jisuan2(10)
	fmt.Println()

	res := p.getSum(10, 20)
	fmt.Println("4. res = ", res)
}



