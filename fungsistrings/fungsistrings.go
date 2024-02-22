package fungsistrings

import (
	"fmt"
	"strings"
)

func Contains() {
	var isExists = strings.Contains("john wick", "wick")
	fmt.Println(isExists)
}

func HashingPrefix() {
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1) // true

	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2) // false
}

func HashingSuffix() {
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1) // false

	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2) // true
}

func Count() {
	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany) // 2
}

func Indexing() {
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1) // 2

	var index2 = strings.Index("ethan hunt", "n")
	fmt.Println(index2) // 4
}

func Replacing() {
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"

	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"

	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"
}

func Repeats() {
	var str = strings.Repeat("na", 4)
	fmt.Println(str) // "nananana"
}

func Splits() {
	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1) // output: ["the", "dark", "knight"]

	var string2 = strings.Split("batman", "")
	fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]
}

func Joins() {
	var data = []string{"banana", "papaya", "tomato"}
	var str = strings.Join(data, "-")
	fmt.Println(str) // "banana-papaya-tomato"
}

func Lower() {
	var str = strings.ToLower("aLAy")
	fmt.Println(str) // "alay"
}

func Upper() {
	var str = strings.ToUpper("eat!")
	fmt.Println(str) // "EAT!"
}
