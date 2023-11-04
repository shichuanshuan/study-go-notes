package main

import "fmt"

// 5. 如果一个类型实现了String()这个方法
// 那么fmt.Println默认会调用这个变量的String()进行输出
type Student struct{
	Name string
	Age int
}

// 给*Student实现方法String
func (student Student) String()string{
	str := fmt.Sprintf("Name = %v Age = %v", student.Name, student.Age)
	return str
}

func main(){
	var stu Student

	stu.Name = "tom"
	stu.Age = 20
	fmt.Printf("%s", stu)
}