package main

import (
	"fmt"
	"strings"
)

// type student struct {
// 	name  string
// 	grade int
// }

type person struct {
	name string
	age  int
}
type People = person

type student struct {
	grade int
	age   int
	person
}

// nested struct
type student1 struct {
	person struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}

func structObj() {
	structNested()
	// alias
	var p1 = person{"wick", 21}
	fmt.Println(p1)

	var p2 = People{"wick", 21}
	fmt.Println(p2)
}

// structNested
func structNested() {

	var students = []student1{
		{
			person: struct {
				name string
				age  int
			}{name: "lauren", age: 20},
			grade:   4,
			hobbies: []string{"main", "mandi Bola"},
		},
		{
			person: struct {
				name string
				age  int
			}{name: "lauren", age: 20},
			grade:   4,
			hobbies: []string{"main", "mandi Bola"},
		},
		{
			person: struct {
				name string
				age  int
			}{name: "lauren", age: 20},
			grade:   4,
			hobbies: []string{"main", "mandi Bola"},
		},
	}

	for _, student := range students {
		hobbies := strings.Join(student.hobbies, ", ")
		fmt.Printf("Nama: %s, umur: %d, grade: %d, Hobby: %s \n", student.person.name, student.person.age, student.grade, hobbies)
	}
}

/*

// menggunakan var
func structVar() {
	// hanya deklarasi
	var student struct {
		grade int
	}
	fmt.Println(student)

	// deklarasi sekaligus insialisasi
	var student1 = struct {
		grade int
	}{
		12,
	}
	fmt.Println(student1.grade)
}

// structSliceAnonymous
func structSliceAnonymous() {
	var allStudents = []struct {
		person
		grade int
	}{
		{person: person{"lauren", 21}, grade: 2},
		{person: person{"dimas", 22}, grade: 1},
		{person: person{"pratama", 21}, grade: 3},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
}

// structSliceStruct
func structSlice() {
	var allStudents = []person{
		{"Lauren", 23},
		{name: "Ethan", age: 23},
		{name: "Pratama", age: 23},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
}

// Anonymous Struct
func structAnonymous() {
 	var s1 = struct {
 		person
 		grade int
 	}{}

 	s1.person = person{"wick", 21}
 	s1.grade = 2

 	fmt.Println("name :", s1.person.name)
 	fmt.Println("age  :", s1.person.age)
 	fmt.Println("grade:", s1.grade)

 	var s2 = struct {
 		person
 		grade int
 	}{
 		person: person{"wick", 21},
 		grade:  2,
 	}

 	fmt.Println("name :", s2.person.name)
 	fmt.Println("age  :", s2.person.age)
 	fmt.Println("grade:", s2.grade)
}

// Embeded struct
func structEmbed() {
 	var s1 = student{}
 	s1.name = "wick"
 	s1.age = 21
 	s1.person.age = 22
 	s1.grade = 2

 	var p11 = person{name: "wick", age: 21}
 	var s11 = student{person: p11, grade: 2}

 	 fmt.Println("name  :", s1.name)
 	 fmt.Println("age   :", s1.age)
 	 fmt.Println("age   :", s1.person.age)
 	 fmt.Println("grade :", s1.grade)

 	fmt.Println("name  :", s11.name)
 	fmt.Println("age   :", s11.age)
 	fmt.Println("grade :", s11.grade)
}

// basic
func structBasic() {
 	var s1 student
 	s1.name = "jhon wick"
 	s1.grade = 2

 	fmt.Println("name :", s1.name)
 	fmt.Println("grade :", s1.grade)
 }

// struct object
func structObject() {
 	var s1 = student{}
 	s1.name = "wick"
 	s1.grade = 2

 	var s2 = student{"ethan", 2}

 	var s3 = student{name: "jason"}

 	fmt.Println("student 1 :", s1.name)
 	fmt.Println("student 2 :", s2.name)
 	fmt.Println("student 3 :", s3.name)
 }

// struct pointer
func structPointer() {
 	var s1 = student{name: "wick", grade: 2}

 	var s2 *student = &s1
 	fmt.Println("student 1, name :", s1.name)
 	fmt.Println("student 4, name :", s2.name)

 	s2.name = "ethan"
 	fmt.Println("student 1, name :", s1.name)
 	fmt.Println("student 4, name :", s2.name)
}
*/
