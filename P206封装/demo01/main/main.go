package main

import (
	"fmt"
	"GoCode/P206封装/demo01/model"
)

func main(){
	p := model.NewPerson("shi")

	// 给薪水和年龄赋值
	(*p).SetAge(19)
	(*p).SetSal(9000)

	// 获取薪水和年龄
	fmt.Printf("Age = %v Sal = %v\n", (*p).GetAge(), (*p).GetSal())

	fmt.Println(*p)
}

