package structs

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	age   int
	grade int
}

func Student() {
	var s1 student

	s1.name = "Iqbal"
	s1.grade = 2

	var s2 = student{
		grade: 2,
	}

	s2.person.name = "Ahmad"

	fmt.Println("name :", s1.name)
	fmt.Println("student2 :", s2.name)
}

func ObjectPointer() {
	var s1 = student{
		grade: 2,
	}
	s1.person.name = "Iqbal"

	var s2 *student = &s1

	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 2, name :", s2.name)

	s2.name = "Ahmad"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 2, name :", s2.name)
}

func EmbededStuct() {
	var s1 = student{}
	s1.name = "Iqbal"
	s1.age = 2
	s1.person.age = 23

	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s1.person.age)
}

func SubStruct() {
	var p1 = person{name: "Iqbal", age: 22}
	var s1 = student{person: p1, grade: 2}

	fmt.Println("name :", s1.name)
	fmt.Println("age :", s1.age)
	fmt.Println("grade :", s1.grade)
}

func AnonymousStruct() {
	var s1 = struct {
		person
		grade int
	}{
		person: person{"Iqbal", 21},
		grade:  2,
	}

	fmt.Println("name	:", s1.name)
	fmt.Println("age	:", s1.age)
	fmt.Println("grade	:", s1.grade)
}

func SliceAndStruct() {
	type person struct {
		name string
		age  int
	}

	var allStudents = []person{
		{name: "Ahmad", age: 21},
		{name: "Maulana", age: 22},
		{name: "Iqbal", age: 23},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
}

func AnonymousStructAndSlice() {
	var allStudents = []struct {
		person
		grade int
	}{
		{person: person{"Ahmad", 21}, grade: 2},
		{person: person{"Maulana", 22}, grade: 2},
		{person: person{"Iqbal", 23}, grade: 2},
	}

	for _, student := range allStudents {
		fmt.Println(student)
	}
}
