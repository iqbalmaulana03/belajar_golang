package regexs

import (
	"fmt"
	"regexp"
)

func Regexs() {
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var res1 = regex.FindAllString(text, 2)
	fmt.Printf("%#v \n", res1)
	// []string{"banana", "burger"}

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)
	// []string{"banana", "burger", "soup"}
}

func MatchingString() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)
	// true
}

func FindingString() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str = regex.FindString(text)
	fmt.Println(str)
	// "banana"
}

func FindingStringIndexs() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)
	// [0, 6]

	var str = text[0:6]
	fmt.Println(str)
	// "banana"
}

func FindingAllStrings() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str1 = regex.FindAllString(text, -1)
	fmt.Println(str1)
	// ["banana", "burger", "soup"]

	var str2 = regex.FindAllString(text, 1)
	fmt.Println(str2)
	// ["banana"]
}

func RepleacingAllString() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str = regex.ReplaceAllString(text, "potato")
	fmt.Println(str)
	// "potato potato potato"
}

func ReplacingAllStrings() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-z]+`)

	var str = regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	})
	fmt.Println(str)
	// "banana potato soup"
}

func Splits() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-b]+`) // split dengan separator adalah karakter "a" dan/atau "b"

	var str = regex.Split(text, -1)
	fmt.Printf("%#v \n", str)
	// []string{"", "n", "n", " ", "urger soup"}
}
