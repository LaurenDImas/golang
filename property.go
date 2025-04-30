package main

import (
	"fmt"
	f "fmt"
	"hello-world/library"
	. "hello-world/library"
)

func property() {
	SayHello()
	var s1 = library.Student{"lauren", 21}
	f.Println("name", s1.Name)
	f.Println("grade", s1.Grade)

	fmt.Println("-----------")

	fmt.Printf("Name  : %s\n", library.Student1.Name)
	fmt.Printf("Grade : %d\n", library.Student1.Grade)
}
