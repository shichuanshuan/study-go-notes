package main

import "fmt"

func main(){
	// 输出三个学生的name 和sex
	strudy := make(map[string]map[string]string)

	strudy["No.1"] = make(map[string]string)
	strudy["No.1"]["name"] = "Tom"
	strudy["No.1"]["sex"] = "boy"

	strudy["No.2"] = make(map[string]string)
	strudy["No.2"]["name"] = "May"
	strudy["No.2"]["sex"] = "giry"

	fmt.Println(strudy)
}