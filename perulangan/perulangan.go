package perulangan

import "fmt"

func For() {
	for i := 0; i < 5; i++ {
		fmt.Println("Angka : ", i)
	}
}

func ForCondition() {
	var i = 0

	for i < 5 {
		fmt.Println("Angka : ", i)
		i++
	}
}

func ForNotCondition() {
	var i = 0

	for {
		fmt.Println("Angka : ", i)

		i++

		if i == 5 {
			break
		}
	}
}

func ForRange() {
	xs := "123"

	for i, v := range xs {
		fmt.Println("Index : ", i, "Value : ", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array

	for _, v := range ys {
		fmt.Println("Value : ", v)
	}

	var zs = ys[0:2] //slice
	for _, v := range zs {
		fmt.Println("Value : ", v)
	}

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} //map
	for k, v := range kvs {
		fmt.Println("Key : ", k, "Value : ", v)
	}

	for range kvs {
		fmt.Println("Done")
	}
}

func BreakCoutinue() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
}

func PerulanganBersarang() {
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
}

func PerulanganLabel() {
outerLoop:
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
