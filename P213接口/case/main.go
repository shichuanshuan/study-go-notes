package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1. 声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2. 声明一个Hero结构体切片类型
type HeroSlice []Hero

// 3. 实现interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less方法就是决定你使用什么标准进行排序
// 按Hero年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

func main() {

	// 测试看看我们是否可以对结构体切片进行排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		// 定义结构体变量，每循环一次，将内容输入进去一次，输入完后空间被释放
		// 下次循环又重新创建变量，以此类推出现可以打印出十组数据
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		// 将hero append 到 heroes切片
		heroes = append(heroes, hero)
	}

	// 看看排序前的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}
	fmt.Println()

	sort.Sort(heroes)
	fmt.Println("* * * * * * * * 排序后 * * * * * * * * ")
	for _, v := range heroes {
		fmt.Println(v)
	}
}
