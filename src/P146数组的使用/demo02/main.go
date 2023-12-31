package main

import (
	"fmt"
)

func main(){

	// for - range遍历数组
	herose := [...]string{"松江", "闵行", "徐汇"}

	for index, value := range herose {
		fmt.Printf("index = %v, value = %v\n", index, value)
	}
}
