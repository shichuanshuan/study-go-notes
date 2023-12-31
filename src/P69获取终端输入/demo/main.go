// 从终端获取数据
package main

import (
	"fmt"
)

func main(){
	// 要求：从控制台接收数据信息，【姓名，年龄，薪水，是否通过】
	var name string
	var age int
	var sal float32
	var isPass bool

	// 方式1fmt.Scanln,
	// fmt.Println("请输入姓名 ")
	// fmt.Scanln(&name, &age, &sal, &isPass)

	// fmt.Println("请输入年龄")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水")
	// fmt.Scanln(&sal)

	// fmt.Println("请确定是否通过")
	// fmt.Scanln(&isPass)
	// fmt.Printf("姓名：%v 年龄：%v 薪水：%v 是否通过：%v ",name, age, sal, isPass)

	// 方式2 fmt.Scanf
	fmt.Println("请输入姓名， 年龄， 薪水， 是否通过")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("姓名：%v 年龄：%v 薪水：%v 是否通过：%v ",name, age, sal, isPass)
}