package main

import (
	"project-golang/regexs"
)

func main() {
	// penggunaan variabel

	// variabel.VariabelType()
	// variabel.VariabelNotType()
	// variabel.MultipleVariabel()

	// penggunaan type data
	// typedata.Numerik()
	// typedata.Boolean()
	// typedata.String()

	//penggunaan constanta
	// constanta.Const()

	//penggunaan operatort matematika
	// operator.Aritmatik()
	// operator.Perbandinga()
	// operator.Logic()

	//penggunaan selection
	// selection.Selection()
	// selection.Temporary()
	// selection.Switch()
	// selection.SwitchIf()
	// selection.Fallthrough()
	// selection.Bersarang()

	//penggunaan perulangan
	// perulangan.For()
	// perulangan.ForCondition()
	// perulangan.ForNotCondition()
	// perulangan.ForRange()
	// perulangan.BreakCoutinue()
	// perulangan.PerulanganBersarang()
	// perulangan.PerulanganLabel()

	//penggunaan array
	// arrays.Array()
	// arrays.InisialisasiArray()
	// arrays.NoInisialisasiArray()
	// arrays.MultidimensiArray()
	// arrays.PrulanganArray()
	// arrays.PerulanganRangeArray()
	// arrays.PerulanganNotUndescore()
	// arrays.MakeArray()

	//penggunaan slice
	// slices.InisialisasiSlice()
	// slices.ArraySlice()
	// slices.TypeReferences()
	// slices.Len()
	// slices.Cap()
	// slices.Append()
	// slices.Copy()
	// slices.SLice3Indexs()

	//penggunaan map
	// maps.Maps()
	// maps.NilaiMaps()
	// maps.ForMaps()
	// maps.MenghapusMaps()
	// maps.DeteksiMaps()
	// maps.SliceMaps()

	//penggunaan multiple return
	// var diameter float64 = 15
	// var area, circumference = multipelreturn.Calculate(diameter)
	// fmt.Printf("luas lingkaran\t\t: %.2f \n", area)
	// fmt.Printf("keliling lingkaran\t: %.2f \n", circumference)

	// penggunaan anonymous function
	// anonymous.MinMax()
	// anonymous.Iife()
	// var max = 3
	// var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	// var howMany, getNumbers = anonymous.FindMax(numbers, max)
	// var theNumbers = getNumbers()
	// fmt.Println("numbers \t:", numbers)
	// fmt.Printf("find \t: %d\n\n", max)
	// fmt.Println("found \t:", howMany)
	// fmt.Println("value \t\n:", theNumbers)
	// var data = []string{"wick", "jason", "ethan"}
	// var dataContainsO = anonymous.Filter(data, func(each string) bool {
	// 	return strings.Contains(each, "o")
	// })
	// var dataLength5 = anonymous.Filter(data, func(each string) bool {
	// 	return len(each) == 5
	// })
	// fmt.Println("data asli \t\t:", data)
	// fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	// fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)

	//penggunaan pointer
	// pointer.Pointer()
	// pointer.PerubahanPointer()
	// pointer.ParameterPointer()

	// penggunaan struct
	// structs.Student()
	// structs.ObjectPointer()
	// structs.EmbededStuct()
	// structs.SubStruct()
	// structs.AnonymousStruct()
	// structs.SliceAndStruct()
	// structs.AnonymousStructAndSlice()

	// penggunaan method
	// method.Methods()
	// method.MethodPointer()

	// penggunaan private or public dari sebuah method atau function
	// f.Printf("Name  : %s\n", library.Student.Name)
	// f.Printf("Grade : %d\n", library.Student.Grade)
	// sayHello("Iqbal")

	// penggunaan interface
	// var bangunDatar interfaces.Hitung
	// bangunDatar = interfaces.Persegi{10.0}
	// fmt.Println("==== persegi")
	// fmt.Println("luas	:", bangunDatar.Luas())
	// fmt.Println("keliling :", bangunDatar.Keliling())
	// bangunDatar = interfaces.Lingkaran{14.0}
	// var bangunLingkaran interfaces.Lingkaran = bangunDatar.(interfaces.Lingkaran)
	// fmt.Println("==== lingkaran")
	// fmt.Println("luas	:", bangunDatar.Luas())
	// fmt.Println("keliling :", bangunDatar.Keliling())
	// fmt.Println("jari-jari :", bangunLingkaran.JariJari())
	// var bangunRuang = &interfaces.Kubus{4}
	// fmt.Println("==== kubus")
	// fmt.Println("luas	:", bangunRuang.Luas())
	// fmt.Println("keliling :", bangunRuang.Keliling())
	// fmt.Println("volume	:", bangunRuang.Volume())

	// penggunaan any interface
	// anys()
	// data()
	// Alias()
	// assertions()
	// casting()
	// kombinasi()

	// penggunaa reflection
	// reflection.Reflection()
	// reflection.ReflectionInterface()
	// reflection.ReflectionProperty()
	// reflection.AksesMethod()

	// penggunaan go routine
	// routines.Rountines()

	// penggunaan channel
	// channels.Channel()
	// channels.ChannelParameter()

	// penggunaan buffered channel
	// buferedchannel.BuffredCahnnel()

	// penggunaan select
	// selects.Selects()

	// penggunaan range close
	// rangeclose.RangeClose()

	// penggunaan channel time out
	// channeltimeout.TimeOut()

	// penggunaan defer
	// defers.OrderSomeFood("pizza")
	// defers.OrderSomeFood("burger")
	// defers.AnonymousDefere()
	// defers.Exits()

	// penggunaan error, panic, recover
	// errorspanics.Errors()
	// errorspanics.CustomerError()
	// errorspanics.Panics()
	// errorspanics.Recovers()
	// errorspanics.RecobersAnonymous()

	// penggunaan layout format string
	// layoutformat.LayOut()

	// penggunaan random math
	// randoms.Generate()
	// randoms.UnixNano()
	// randoms.RandomDataNumerik()
	// randoms.RandomsString()

	// cara penggunaan time
	// times.Times()
	// times.ParsingTime()
	// times.PredenfinedTime()
	// times.TimeToString()
	// times.HandleErrorTime()
	// times.TimesSlepp()
	// times.ScedhulerSleep()
	// times.NewTimers()
	// times.AftersTime()
	// times.SchedulerTicker()
	// times.TimerGoruotine()
	// times.Duration()
	// times.KalkulasiSelisih()

	// penggunaan konversi tipe data
	// konvertiondata.Atoi()
	// konvertiondata.Itoa()
	// konvertiondata.ParsingInt()
	// konvertiondata.ParsingFloat()
	// konvertiondata.ParsingBool()
	// konvertiondata.ParsingBiner()
	// konvertiondata.FormatingFloat()
	// konvertiondata.FormatingBool()
	// konvertiondata.FormatInt()
	// konvertiondata.Casting()
	// konvertiondata.CastingString()
	// konvertiondata.AssertionInterface()

	// penggunaan package strings
	// fungsistrings.Contains()
	// fungsistrings.Count()
	// fungsistrings.HashingPrefix()
	// fungsistrings.HashingSuffix()
	// fungsistrings.Indexing()
	// fungsistrings.Joins()
	// fungsistrings.Lower()
	// fungsistrings.Repeats()
	// fungsistrings.Replacing()
	// fungsistrings.Splits()
	// fungsistrings.Upper()

	// penggunaan regex
	regexs.FindingAllStrings()
	regexs.FindingString()
	regexs.FindingStringIndexs()
	regexs.MatchingString()
	regexs.Regexs()
	regexs.ReplacingAllStrings()
	regexs.RepleacingAllString()
	regexs.Splits()
}
