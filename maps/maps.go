package maps

import "fmt"

func Maps() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])
}

func NilaiMaps() {
	var data map[string]int

	data = map[string]int{}
	data["one"] = 1
}

func ForMaps() {
	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"juni":     67,
	}

	for key, val := range chicken {
		fmt.Println(key, " \t:", val)
	}
}

func MenghapusMaps() {
	var chicken = map[string]int{"januari": 50, "feebruari": 40}

	fmt.Println(len(chicken))
	fmt.Println(chicken)

	delete(chicken, "januari")

	fmt.Println(len(chicken))
	fmt.Println(chicken)
}

func DeteksiMaps() {
	var chicken = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exists")
	}
}

func SliceMaps() {
	// var chickens = []map[string]string{
	// 	{"name": "chicken blue", "gender": "male"},
	// 	{"name": "chicken yellow", "gender": "female"},
	// 	{"name": "chicken green", "gender": "male"},
	// }

	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}

	for _, chiken := range chickens {
		fmt.Println(chiken["id"], chiken["community"])
	}
}
