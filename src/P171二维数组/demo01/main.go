package main

import "fmt"

func main(){
	// 二维数组的遍历
	var arr [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("arr[0]", len(arr[0]))

	// for循环遍历
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t", arr[i][j])
		}
		fmt.Println()
	}


	// for-range 遍历二维数组
	for index, value := range arr{
		for j, value2 := range value{
			fmt.Printf("arr[%d][%d] = %d\n", index, j, value2)
		}
		
	}
}