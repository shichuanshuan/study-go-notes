package main

import "fmt"

func main(){
	// map切片的使用
	/*
	要求：使用一个map来记录monster的信息 name 和 age， 也就是说一个
	monster 对应一个map， 并且妖怪的个数可以动态的增加=>map切片
	*/
	// 1.声明一个map切片
	var monster []map[string]string
	monster = make([]map[string]string, 2) // 准备放两个妖怪进去

	// 2.增加一个妖怪信息
	if monster[0] == nil {
		monster[0] = make(map[string]string)
		monster[0]["name"] = "牛魔王"
		monster[0]["age"] = "1000"
		monster[0]["sex"] = "boy"
	}

	if monster[1] == nil {
		monster[1] = make(map[string]string)
		monster[1]["name"] = "玉兔精"
		monster[1]["age"] = "600"
	}


	// 下面越界
	// if monster[2] == nil {
	// 	monster[2] = make(map[string]string)
	// 	monster[2]["name"] = "狐狸精"
	// 	monster[2]["age"] = "300"
	// }
	newMonster := map[string]string{
		"name" : "新妖怪~狐狸精",
		"age" : "300",
	}
	monster = append(monster, newMonster)
	fmt.Println(monster)
}