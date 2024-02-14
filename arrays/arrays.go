package arrays

import "fmt"

func Array() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "D"
	names[2] = "Water"
	names[3] = "Law"

	fmt.Println(names[0], names[1], names[2], names[3])
}

func InisialisasiArray() {
	var fruits = [4]string{"apple", "nanas", "semangka", "jeruk"}

	fmt.Println("Jumalh Elements \t\t", len(fruits))
	fmt.Println("Isi semua elemets \t", fruits)
}

func NoInisialisasiArray() {
	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemt \t:", len(numbers))
}

func MultidimensiArray() {
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1 ", numbers1)
	fmt.Println("numbers2 ", numbers2)
}

func PrulanganArray() {
	var fruits = [4]string{"apple", "nanas", "jeruk", "semangka"}

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("element %d : %s\n", i, fruits[i])
	}
}

func PerulanganRangeArray() {
	var fruits = [4]string{"apple", "nanas", "jeruk", "semangka"}

	for i, fruit := range fruits {
		fmt.Printf("element %d : %s\n", i, fruit)
	}
}

func PerulanganNotUndescore() {
	var fruits = [4]string{"apple", "nanas", "jeruk", "semangka"}

	for _, fruit := range fruits {
		fmt.Printf("nama buah: %s\n", fruit)
	}
}

func MakeArray() {
	var fruits = make([]string, 2)

	fruits[0] = "apple"
	fruits[1] = "jeruk"

	fmt.Println(fruits)
}
