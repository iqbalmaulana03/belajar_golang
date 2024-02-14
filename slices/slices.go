package slices

import "fmt"

func InisialisasiSlice() {
	var fruits = []string{"nanas", "apple", "anggur", "banana"}

	fmt.Println(fruits)
}

func ArraySlice() {
	var fruits = []string{"nanas", "apple", "anggur", "banana"}
	var newFruits = fruits[0:2]

	fmt.Println(newFruits)
}

func TypeReferences() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]

	baFruits[0] = "pinnaple"

	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)
}

func Len() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	fmt.Println(len(fruits))
}

func Cap() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	var bFruits = fruits[1:4]
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))
}

func Append() {
	var fruits = []string{"apple", "grape", "banana"}

	var bFruits = fruits[0:2]

	fmt.Println(cap(bFruits)) // 3
	fmt.Println(len(bFruits)) // 2

	fmt.Println(fruits)  // ["apple", "grape", "banana"]
	fmt.Println(bFruits) // ["apple", "grape"]

	var cFruits = append(bFruits, "papaya")

	fmt.Println(fruits)
	fmt.Println(bFruits)
	fmt.Println(cFruits)
}

func Copy() {
	dst := []string{"potato", "potato", "potato"}
	src := []string{"watermelon", "pinnaple"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
}

func SLice3Indexs() {
	var fruits = []string{"apple", "grape", "banana"}

	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(aFruits)      // ["apple", "grape"]
	fmt.Println(len(aFruits)) // len: 2
	fmt.Println(cap(aFruits)) // cap: 3

	fmt.Println(bFruits)      // ["apple", "grape"]
	fmt.Println(len(bFruits)) // len: 2
	fmt.Println(cap(bFruits)) // cap: 2
}
