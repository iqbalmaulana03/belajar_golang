package errorspanics

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Errors() {
	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	var number int
	var err error

	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}
}

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty!")
	}
	return true, nil
}

func CustomerError() {
	var name string
	fmt.Print("Type your name :")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		fmt.Println(err.Error())
	}
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func Panics() {
	var name string
	fmt.Print("Type your name :")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}

func Recovers() {
	defer catch()

	var name string
	fmt.Print("Type your name :")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
		fmt.Println("end")
	}
}

func RecobersAnonymous() {
	datas := []string{"superman", "aquaman", "batman"}

	for _, data := range datas {
		func() {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occured on looping", data, "| message:", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()

			panic("some error happen")
		}()
	}
}
