package main

import (
	"fmt"
	"project-golang/anonymous"
	"strings"
)

func main() {
	// variabel.VariabelType()
	// variabel.VariabelNotType()
	// variabel.MultipleVariabel()

	// typedata.Numerik()
	// typedata.Boolean()
	// typedata.String()

	// constanta.Const()

	// operator.Aritmatik()
	// operator.Perbandinga()
	// operator.Logic()

	// selection.Selection()
	// selection.Temporary()
	// selection.Switch()
	// selection.SwitchIf()
	// selection.Fallthrough()
	// selection.Bersarang()

	// perulangan.For()
	// perulangan.ForCondition()
	// perulangan.ForNotCondition()
	// perulangan.ForRange()
	// perulangan.BreakCoutinue()
	// perulangan.PerulanganBersarang()
	// perulangan.PerulanganLabel()

	// arrays.Array()
	// arrays.InisialisasiArray()
	// arrays.NoInisialisasiArray()
	// arrays.MultidimensiArray()
	// arrays.PrulanganArray()
	// arrays.PerulanganRangeArray()
	// arrays.PerulanganNotUndescore()
	// arrays.MakeArray()

	// slices.InisialisasiSlice()
	// slices.ArraySlice()
	// slices.TypeReferences()
	// slices.Len()
	// slices.Cap()
	// slices.Append()
	// slices.Copy()
	// slices.SLice3Indexs()

	// maps.Maps()
	// maps.NilaiMaps()
	// maps.ForMaps()
	// maps.MenghapusMaps()
	// maps.DeteksiMaps()
	// maps.SliceMaps()

	// var diameter float64 = 15

	// var area, circumference = multipelreturn.Calculate(diameter)

	// fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	// fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	anonymous.MinMax()
	anonymous.Iife()

	var max = 3
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var howMany, getNumbers = anonymous.FindMax(numbers, max)
	var theNumbers = getNumbers()

	fmt.Println("numbers \t:", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)
	fmt.Println("value \t\n:", theNumbers)

	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = anonymous.Filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})

	var dataLength5 = anonymous.Filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)

	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)
}
