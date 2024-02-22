package konvertiondata

import (
	"fmt"
	"strconv"
)

func Atoi() {
	var str = "124"
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num)
	}
}

func Itoa() {
	var num = 124
	var str = strconv.Itoa(num)

	fmt.Println(str)
}

func ParsingInt() {
	var str = "124"
	var num, err = strconv.ParseInt(str, 10, 64)

	if err == nil {
		fmt.Println(num)
	}
}

func ParsingBiner() {
	var str = "1010"
	var num, err = strconv.ParseInt(str, 2, 8)

	if err == nil {
		fmt.Println(num)
	}
}

func FormatInt() {
	var num = int64(24)
	var str = strconv.FormatInt(num, 8)

	fmt.Println(str)
}

func ParsingFloat() {
	var str = "24.12"
	var num, err = strconv.ParseFloat(str, 32)

	if err == nil {
		fmt.Println(num)
	}
}

func FormatingFloat() {
	var num = float64(24.12)
	var str = strconv.FormatFloat(num, 'f', 6, 64)

	fmt.Println(str)
}

func ParsingBool() {
	var str = "true"
	var bul, err = strconv.ParseBool(str)

	if err == nil {
		fmt.Println(bul)
	}
}

func FormatingBool() {
	var bul = true
	var str = strconv.FormatBool(bul)

	fmt.Println(str)
}

func Casting() {
	var a float64 = float64(24)
	fmt.Println(a)

	var b int32 = int32(24.00)
	fmt.Println(b)
}

func CastingString() {
	var text1 = "halo"
	var b = []byte(text1)

	fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])

	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)

	fmt.Printf("%s \n", s)

	var c int64 = int64('h')
	fmt.Println(c)

	var d string = string(104)
	fmt.Println(d)
}

func AssertionInterface() {
	var data = map[string]interface{}{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}

	fmt.Println(data["nama"].(string))
	fmt.Println(data["grade"].(int))
	fmt.Println(data["height"].(float64))
	fmt.Println(data["isMale"].(bool))
	fmt.Println(data["hobbies"].([]string))

	for _, val := range data {
		switch val.(type) {
		case string:
			fmt.Println(val.(string))
		case int:
			fmt.Println(val.(int))
		case float64:
			fmt.Println(val.(float64))
		case bool:
			fmt.Println(val.(bool))
		case []string:
			fmt.Println(val.([]string))
		default:
			fmt.Println(val.(int))
		}
	}
}
