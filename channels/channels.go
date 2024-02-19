package channels

import (
	"fmt"
	"runtime"
)

func Channel() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("Ahmad Iqbal")
	go sayHelloTo("Aldy Candra")
	go sayHelloTo("Mustofa Ali")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}

func printMessage(what chan string) {
	fmt.Println(<-what)
}

func ChannelParameter() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	for _, each := range []string{"Ahmad", "Maulana", "Iqbal"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}
