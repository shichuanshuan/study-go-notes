package main

import "fmt"

type Student struct{

}

// 写一个函数，循环判断函数的类型
func TypeJudeg(items... interface{}){
	for index, value := range items {
		switch value.(type){
			case int8, int, int64 :
				fmt.Printf("第 %d 个数，是 %T 类型, 值是 %v\n", index + 1, value, value)
			case float32, float64 :
				fmt.Printf("第 %d 个数，是 %T 类型, 值是 %v\n", index + 1, value, value)
			case string :
				fmt.Printf("第 %d 个数，是 %T 类型, 值是 %v\n", index + 1, value, value)
			case Student :
				fmt.Printf("第 %d 个数，是 %T 类型, 值是 %v\n", index + 1, value, value)
			case *Student :
				fmt.Printf("第 %d 个数，是 %T 类型, 值是 %v\n", index + 1, value, value)
		}
	}
}

func main(){
	var i32_Ret int
	var s32_Ret string
	var f32_Ret float32
	
	stu1 := Student{}
	stu2 := &Student{}

	TypeJudeg(s32_Ret, i32_Ret, stu1, f32_Ret, stu2)
}