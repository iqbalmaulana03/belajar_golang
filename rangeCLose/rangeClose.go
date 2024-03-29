package rangeclose

import (
	"fmt"
	"runtime"
)

func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func RangeClose() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	go sendMessage(messages)
	printMessage(messages)
}
