package main

import (
	"encoding/json"
	"fmt"
)

type Hero struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testMap() {
	// var a map[string]interface{}
	// a = make(map[string]interface{})

	// a["str"] = Hero{
	// 	Name:     "shidashuan",
	// 	Age:      25,
	// 	Birthday: "0116",
	// 	Sal:      8000.0,
	// 	Skill:    "love",
	// }
	hero := Hero{
		Name:     "shidashuan",
		Age:      25,
		Birthday: "0116",
		Sal:      8000.0,
		Skill:    "love",
	}

	// fmt.Printf("marshal after = %v\n", a)

	data, err := json.Marshal(hero)
	if err != nil {
		fmt.Printf("err = %v\n", err)
	}

	fmt.Printf("input = %v\n", string(data))

	// ======================================
	b := make(map[string]interface{})
	err = json.Unmarshal(data, &b)
	if err != nil {
		fmt.Printf("err = %v\n", err)
	}
	fmt.Printf("unmarshal = %v\n", b)
	// fmt.Printf("unmarshal = %v %T\n", b["Age"], b["Age"])
	fmt.Printf("unmarshal = %v\n", b["Age"].(float64))
}

func main() {
	testMap()
}

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// func testMap() {
// 	var a map[string]interface{}
// 	a = make(map[string]interface{})

// 	a["name"] = "shi"
// 	a["age"] = 25
// 	a["address"] = "ys"

// 	fmt.Printf("marshal after = %v\n", a)

// 	data, err := json.Marshal(a)
// 	if err != nil {
// 		fmt.Printf("err = %v\n", err)
// 	}

// 	fmt.Printf("input = %v\n", string(data))

// 	// ======================================
// 	b := make(map[string]interface{})
// 	err = json.Unmarshal(data, &b)
// 	if err != nil {
// 		fmt.Printf("err = %v\n", err)
// 	}
// 	fmt.Printf("unmarshal = %v\n", b)
// 	fmt.Printf("unmarshal = %v\n", b["name"])
// }

// func main() {
// 	testMap()
// }
