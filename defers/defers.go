package defers

import (
	"fmt"
	"os"
)

func OrderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, silahkan menunggu")
	if menu == "pizza" {
		fmt.Print("Pilihan tepat!", " ")
		fmt.Print("Pizza ditempat kami paling enak!", "\n")
		return
	}

	fmt.Println("Pesanan anda :", menu)
}

func AnonymousDefere() {
	number := 3

	if number == 3 {
		fmt.Println("Halo 1")
		func() {
			defer fmt.Println("halo 3")
		}()
	}

	fmt.Println("halo 2")
}

func Exits() {
	defer fmt.Println("halo")
	os.Exit(1)
	fmt.Println("selamat datang")
}
