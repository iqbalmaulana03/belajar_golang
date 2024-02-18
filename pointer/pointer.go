package pointer

import "fmt"

func Pointer() {

	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA (address) :", &numberA)

	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("numberB (adddress) :", &numberB)
}

func PerubahanPointer() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("numberB (adddress) :", &numberB)

	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("numberB (adddress) :", &numberB)
}

func ParameterPointer() {
	var number = 4
	fmt.Println("before :", number)

	change(&number, 10)
	fmt.Println("after :", number)
}

func change(original *int, value int) {
	*original = value
}
