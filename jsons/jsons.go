package jsons

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func DecodeJson() {
	var jsonString = `{"Name": "Ahmad Iqbal", "Age": 23}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age :", data.Age)
}

func DecodeMapJson() {

	var jsonString = `{"Name": "Ahmad Iqbal", "Age": 23}`
	var jsonData = []byte(jsonString)

	var data interface{}
	json.Unmarshal(jsonData, &data)

	var decodeData = data.(map[string]interface{})
	fmt.Println("user	:", decodeData["Name"])
	fmt.Println("age	:", decodeData["Age"])
}

func DecodeArrayJson() {

	var jsonString = `[
		{"Name": "Ahmad Iqbal", "Age": 23},
		{"Name": "Aldi Candra", "Age": 22}
	]`

	var data []User

	var err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1:", data[0].FullName)
	fmt.Println("user 2:", data[1].FullName)
}

func EncodingJson() {

	var objects = []User{{"Ahmad Iqbal", 23}, {"Aldi Candra", 22}}
	var jsonData, err = json.Marshal(objects)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
