package main

import "fmt"

type Goods struct{
	Name string
	Price float64
}

type Brand struct{
	Name string
	Address string
}

type TV struct{
	Goods
	Brand
}


type TV2 struct{
	*Goods
	*Brand
}

func main(){
	// 6. 嵌套匿名结构体后，可以在创建结构体变量时，
	// 直接指定各个匿名结构体字段
	tv := TV{ Goods{"电视机001", 1000.9}, Brand{"海尔", "山东"}}
	fmt.Println("tv", tv)
	fmt.Println()

	// 方式2 可以不按照默认顺序
	tv2 := TV{ 
		Goods{
			Name : "电视机002", 
			Price : 2000.9,
			}, 
		Brand{
				Name : "夏普", 
				Address : "北京",
			},
		}
	fmt.Println("tv2", tv2)
	fmt.Println()

	// 方式3
	tv3 := TV2{ &Goods{"电视机003", 1000.9}, &Brand{"创维", "河南"}}
	fmt.Println("tv3", *tv3.Goods, *tv3.Brand)  // 此处不能写成(*tv3).Goods  
	fmt.Println()

	// 方式4
	tv4 := TV2{ 
		&Goods{
			Name : "电视机004", 
			Price : 2000.9,
			}, 
		&Brand{
				Name : "长虹", 
				Address : "四川",
			},
		}
		fmt.Println("tv4", *tv4.Goods, *tv4.Brand)
}