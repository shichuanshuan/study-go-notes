package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 要求：随机生成五个数，并将其打印
	// 思路
	// 1. 随机生成五个数，rand.Intn()函数
	// 2. 当我们获得随机数后，就放到一个数组 int数组
	// 3. 反转打印，交换的次数为 len

	var intArr [5]int

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100) + 1
	}

	fmt.Println("交换前 = ", intArr)

	temp := 0

	for i := 0; i < len(intArr)/2; i++ {
		temp = intArr[len(intArr)-1-i]
		intArr[len(intArr)-1-i] = intArr[i]
		intArr[i] = temp
	}

	fmt.Println("交换后 = ", intArr)
}
