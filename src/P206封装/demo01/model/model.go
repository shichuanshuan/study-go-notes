package model

import "fmt"

type person struct{
	Name string
	age int
	sal float64
}

// 写一个工厂模式的函数，相当于构造函数
func NewPerson(name string)*person{
	return &person{
		Name : name,
	}
}

// 为了访问age 和 sal 我们编写一对Setxx的方法和Getxxx的方法
func (p *person) SetAge(age int){
	if age > 0 && age < 120 {
		p.age = age
	} else {
		fmt.Println("age err")
	}
}

// 获取年龄函数
func (p *person) GetAge() int{
	return p.age
}

// 设置薪水函数
func (p *person) SetSal(sal float64){
	if sal >= 3000 && sal <= 30000 {
		p.sal = sal
	} else {
		fmt.Println("sal err")
	}
}

// 获得薪水函数
func (p *person) GetSal()float64{
	return p.sal
}