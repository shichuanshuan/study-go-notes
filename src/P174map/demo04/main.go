package main

import "fmt"

func main(){
	// map只能使用for-range循环遍历
	cities := make(map[string]string)
	cities["No.1"] = "juan"
	cities["No.2"] = "lan"
	cities["No.3"] = "shuan"

	for index, value := range cities {
		fmt.Printf("类型为= %v 值为= %v\n", index, value)
	}
	fmt.Println()


	// map的长度
	fmt.Println("cities 有", len(cities), "对 key-value")
	fmt.Println()


	// 使用 for-range 遍历所有信息
	student := make(map[string]map[string]string)
	student["No.1"] = make(map[string]string)
	student["No.1"]["name"] = "shichuanshuan"
	student["No.1"]["sex"] = "boy"

	student["No.2"] = make(map[string]string)
	student["No.2"]["name"] = "Nolike"
	student["No.2"]["sex"] = "giry"

	for k1, v1 := range student {
		fmt.Printf("k1 = %v v1 = %v\n", k1, v1)

		for k2, v2 := range v1{
			fmt.Printf("k2 = %v v2 = %v\n", k2, v2)
		}

		fmt.Println()
	}
}