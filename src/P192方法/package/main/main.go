package main

import (
	"fmt"
	"GoCode/P192方法/package/model"
)

func main(){
	// 定student结构体是首字母小写，我们可以通过工厂模式来解决
	var stu = model.NewStudent("Tom", 20.0)

	fmt.Println(*stu)
	fmt.Println("name =", stu.Name, "score =", stu.GetScore())
}