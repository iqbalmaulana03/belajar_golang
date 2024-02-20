package randoms

import (
	"fmt"
	"math/rand"
	"time"
)

func Generate() {
	randomizer := rand.New(rand.NewSource(10))
	fmt.Println("random k-1 :", randomizer.Int())
	fmt.Println("random k-2 :", randomizer.Int())
	fmt.Println("random k-3 :", randomizer.Int())
}

func UnixNano() {
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("random k-1 :", randomizer.Int())
	fmt.Println("random k-2 :", randomizer.Int())
	fmt.Println("random k-3 :", randomizer.Int())
}

func RandomDataNumerik() {
	randomizer := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("random int:", randomizer.Int())
	fmt.Println("random float32:", randomizer.Float32())
	fmt.Println("random uint:", randomizer.Uint32())
}

func randomStrings(length int) string {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)

}

func RandomsString() {
	fmt.Println("random string 5 karakter:", randomStrings(5))
}
