package reflection

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func Reflection() {

	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}
}

func ReflectionInterface() {

	var number = 20
	var reflectionValue = reflect.ValueOf(number)
	var nilai = reflectionValue.Interface().(int)

	fmt.Println("tipe variabel :", reflectionValue.Type())
	fmt.Println("nilai variabel :", nilai)
}

func (s *student) getPropertyInfo() {
	var reflectionValue = reflect.ValueOf(s)

	if reflectionValue.Kind() == reflect.Ptr {
		reflectionValue = reflectionValue.Elem()
	}

	var relectionType = reflectionValue.Type()

	for i := 0; i < reflectionValue.NumField(); i++ {
		fmt.Println("nama		:", relectionType.Field(i).Name)
		fmt.Println("tipe data	:", relectionType.Field(i).Type)
		fmt.Println("nilai		:", reflectionValue.Field(i).Interface())
		fmt.Println("")
	}
}

func ReflectionProperty() {
	var s1 = &student{Name: "Iqbal", Grade: 2}
	s1.getPropertyInfo()
}

func (s *student) SetName(name string) {
	s.Name = name
}

func AksesMethod() {
	var s1 = &student{Name: "Ahmad Iqbal", Grade: 2}
	fmt.Println("nama	:", s1.Name)

	var reflectionValue = reflect.ValueOf(s1)
	var method = reflectionValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Iqbal"),
	})

	fmt.Println("nama	:", s1.Name)
}
