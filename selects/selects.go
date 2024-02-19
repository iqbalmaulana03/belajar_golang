package selects

import (
	"fmt"
	"runtime"
)

func getAverage(numbers []int, ch chan float64) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}

	ch <- float64(sum) / float64(len(numbers))
}

func getMax(numbers []int, ch chan int) {
	var max = numbers[0]
	for _, number := range numbers {
		if max < number {
			max = number
		}
	}

	ch <- max
}

func Selects() {
	runtime.GOMAXPROCS(2)

	var numbers = []int{3, 4, 3, 5, 6, 3, 2, 2, 6, 3, 4, 3, 6}
	fmt.Println("numbers :", numbers)

	var ch1 = make(chan float64)
	go getAverage(numbers, ch1)

	var ch2 = make(chan int)
	go getMax(numbers, ch2)

	for i := 0; i < 2; i++ {
		select {
		case avg := <-ch1:
			fmt.Printf("Avg \t: %.2f \n", avg)
		case max := <-ch2:
			fmt.Printf("Max \t: %d \n", max)
		}
	}
}
