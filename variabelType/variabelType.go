package variabel

import "fmt"

func VariabelType() {
	var firstName string = "Ahmad"
	var lastName string
	lastName = "Iqbal"
	fmt.Printf("Hello %s %s !\n", firstName, lastName)
}

func VariabelNotType() {
	var firstName string = "Ahmad"
	lastName := "Iqbal"
	fmt.Printf("Hello %s %s !\n", firstName, lastName)
}

func MultipleVariabel() {
	one, isFriday, twoPoint, say := 1, true, 2.2, "hello"

	fmt.Printf("%d %t %.2f %s \n", one, isFriday, twoPoint, say)
}
