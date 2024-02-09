package typeData

import "fmt"

func Numerik() {
	var decimalNumber = 2.64

	fmt.Printf("Bilang decimal : %f\n", decimalNumber)
	fmt.Printf("Bilang decimal : %.3f\n", decimalNumber)
}

func Boolean() {
	var exist bool = true

	fmt.Printf("exist? %t \n", exist)
}

func String() {
	var message string = `Nama Saya "Ahmad Iqbal".
	Salam Kenal.
	Mari belajar "golang".`

	fmt.Println(message)
}
