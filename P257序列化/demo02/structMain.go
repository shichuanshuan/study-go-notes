package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Hero struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	hero := Hero{
		Name:     "shidashuan",
		Age:      25,
		Birthday: "0116",
		Sal:      8000.0,
		Skill:    "love",
	}

	hero2 := Hero{}

	data, err := json.Marshal(&hero)
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}

	fmt.Printf("input = %v\n", string(data))

	// =====================================
	err = json.Unmarshal(data, &hero2)
	if err != nil {
		errors.New("unmarshal error")
	}
	fmt.Printf("unmarshal = %v\n", hero2)
	// fmt.Printf("unmarshal = %v\n", hero2["Name"])
}

func main() {
	testStruct()
}
