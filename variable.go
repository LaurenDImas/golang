package main

import "fmt"

func variable() {
	var firstName string = "john"
	var middleName = "cornius"
	lastName := "wick"
	lastName = "pratama"
	fmt.Printf("halo %s %s %s!\n", firstName, middleName, lastName)
	fmt.Println("halo", firstName, middleName, lastName+"!")

	// var first, second, third string
	// first, second, third = "satu", "dua", "tiga"

	// var fourth, fifth, sixth string = "empat", "lima", "enam"

	// seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	_ = "belajar Golang"
	_ = "Golang itu mudah"
	// name, _ := "john", "wick"
	// fmt.Println(name)

	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)

}
