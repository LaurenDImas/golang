package main

import (
	"fmt"
	"strings"
)

type personPointer struct {
	name string
	age  int
}

func interfaceKosong() {
	var secret interface{}
	secret = "lauren dimas"
	fmt.Println(secret)
	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)
	secret = 12.4
	fmt.Println(secret)

	var data map[string]interface{}
	data = map[string]interface{}{
		"name":      "Lauren Dimas",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}
	fmt.Println(data)

	fmt.Println("-----------------")
	// casting interface (type assertions)
	var secret1 interface{}
	secret1 = 2
	var number = secret.(float64) * 10
	fmt.Println(secret1, "multiplied by 10 is :", number)
	secret1 = []string{"apple", "manggo", "banana"}
	var fruits = strings.Join(secret1.([]string), ", ")
	fmt.Println(fruits, "is my favorite fruits")

	fmt.Println("-----------------")
	// casting interface object pointer
	// biasa
	var secret2 = personPointer{name: "wick", age: 27}
	var name = secret2.name
	fmt.Println(name)
	// interface
	var secret3 interface{} = &personPointer{name: "wick", age: 27}
	var name1 = secret3.(*personPointer).name
	fmt.Println(name1)

	fmt.Println("-----------------")
	// slice map interface{}
	var person1 = []map[string]interface{}{
		{"name": "Lauren", "age": 22},
		{"name": "Dimas", "age": 23},
		{"name": "Yogi", "age": 24},
	}
	for _, person := range person1 {
		fmt.Println(person["name"], "age is", person["age"])
	}
	var fruits1 = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 20},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}
	for _, fruit := range fruits1 {
		fmt.Print(fruit)
	}
}
