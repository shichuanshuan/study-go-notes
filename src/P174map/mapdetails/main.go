package main

import "fmt"

func test(map1 map[int]int){
	map1[3] = 60
}

type Stu struct{
	name string
	age int
	sex string
}

func main(){
	// 1. map是引用类型
	num1 := make(map[int]int) 

	num1[1] = 10
	num1[2] = 20
	num1[3] = 30
	test(num1)

	fmt.Println("num1 = ", num1)
	fmt.Println()


	// 2. map 容量达到后，再想 map 增加元素，会自动扩容，并不会发生panic，也就是说map能动态增长 键值对(key - value)


	// 3. map 的 value 经常使用 struct， 更适合管理复杂的数据
	var students map[string]Stu
	students = make(map[string]Stu)

	stu1 := Stu{"tom~", 18, "boy"}
	stu2 := Stu{"Haidi~", 13, "giry"}

	students["student1"] = stu1
	students["student2"] = stu2

	fmt.Printf("%v\n", students)
	fmt.Println()


	// 遍历各个学生
	for index, v := range students{
		fmt.Printf("学生的编号是%v\n", index)
		fmt.Printf("学生的姓名是%v\n", v.name)
		fmt.Printf("学生的年龄是%v\n", v.age)
		fmt.Printf("学生的性别是%v\n", v.sex)
		fmt.Println()
	}
}
