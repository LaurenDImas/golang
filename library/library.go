package library

import "fmt"

type Student struct {
	Name  string
	Grade int
}

var Student1 = struct {
	Name  string
	Grade int
}{}

func SayHello() {
	fmt.Println("Hello")
	introduce("lauren")
}

func introduce(name string) {
	fmt.Println("nama saya", name)
}

func init() {
	Student1.Name = "Lauren Dimas"
	Student1.Grade = 2
	fmt.Println("---> library/library.go imported")
}
