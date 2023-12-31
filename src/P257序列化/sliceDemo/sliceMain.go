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

func testSlice() {
	var slice map[string]interface{}
	// var m1 map[string]interface{}

	// m1 = make(map[string]interface{})
	// m1["name"] = "dashuan"
	// m1["age"] = 25
	// m1["address"] = "ys"

	// slice = append(slice, m1)

	var m2 map[string]interface{}
	hero := Hero{
		Name:     "shidashuan",
		Age:      25,
		Birthday: "0116",
		Sal:      8000.0,
		Skill:    "love",
	}
	slice = make(map[string]interface{})
	slice["str"] = hero
	// slice = append(slice, m2)

	fmt.Printf("after = %v\n", slice)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("marshal err = %v\n", err)
	}

	fmt.Printf("input = %v\n", string(data))

	// ====================================

	slice2 := make([]map[string]interface{}, 2)
	err = json.Unmarshal(data, &slice2)
	if err != nil {
		fmt.Printf("unmarshal err = %v\n", err)
	}

	fmt.Printf("unmarshal = %v\n", slice2)
	// fmt.Printf("unmarshal = %v\n", slice2["str"])
	for index, value := range slice2[0] {
		fmt.Printf("index = [%v] value = [%v]\n", index, value)
	}
}

func main() {
	testSlice()
}

/*
package main

import (
	"encoding/json"
	"fmt"
)

func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}

	m1 = make(map[string]interface{})
	m1["name"] = "dashuan"
	m1["age"] = 25
	m1["address"] = "ys"

	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "sanfen"
	m2["age"] = 109
	m2["address"] = [2]string{"wudang", "xiawu"}
	slice = append(slice, m2)

	fmt.Printf("after = %v\n", slice)

	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("marshal err = %v\n", err)
	}

	fmt.Printf("input = %v\n", string(data))

	// ====================================

	slice2 := make([]map[string]interface{}, 2)
	err = json.Unmarshal(data, &slice2)
	if err != nil {
		fmt.Printf("unmarshal err = %v\n", err)
	}

	fmt.Printf("unmarshal = %v\n", slice2)
	fmt.Printf("unmarshal = %v\n", slice2[0])
}

func main() {
	testSlice()
}
*/
