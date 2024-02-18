package main

import (
	"fmt"
	"strings"
)

func anys() {
	var secret interface{}

	secret = "Ahmad Iqbal"
	fmt.Println(secret)

	secret = []string{"apple", "amnggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)
}

func data() {
	var datas map[string]interface{}

	datas = map[string]interface{}{
		"name":      "Ahmad Iqbal",
		"grade":     2,
		"brakfeast": []string{"apple", "manggo", "banana"},
	}

	fmt.Println(datas)
}

func Alias() {
	var data map[string]any

	data = map[string]any{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}

	fmt.Println(data)
}

func assertions() {
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10

	fmt.Println(secret, "multipled by 10 is :", number)

	secret = []string{"apple", "manggo", "banana"}
	var fruits = strings.Join(secret.([]string), ", ")
	fmt.Println(fruits, "is my favorite fruits")
}

func casting() {

	type person struct {
		name string
		age  int
	}

	var secret interface{} = &person{name: "Iqbal", age: 20}
	var name = secret.(*person).name
	fmt.Println(name)
}

func kombinasi() {
	var persons = []map[string]interface{}{
		{"name": "Ahmad", "age": 23},
		{"name": "Maulana", "age": 23},
		{"name": "Iqbal", "age": 23},
	}

	var fruits = []interface{}{
		map[string]interface{}{"name": "Banana", "total": 29},
		[]string{"manggo", "appled", "orange"},
		"papaya",
	}

	for _, person := range persons {
		fmt.Println(person["name"], "age is", person["age"])
	}

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
