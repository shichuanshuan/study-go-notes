// 切片的遍历

package main

import (
	"fmt"
)

func main() {
	// 使用常规方式遍历
	slice := []int{10, 20, 30}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v ", i, slice[i])
	}

	fmt.Println()
	// 使用for - range方式遍历
	for index, value := range slice {
		fmt.Printf("slice[%v] = %v ", index, value)
	}
}
