package main

import (
	"fmt"
	"strings"
)

type studentMethod struct {
	name  string
	grade int
}

func (s studentMethod) sayHello() {
	fmt.Println("halo", s.name)
}

func (s studentMethod) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func (s studentMethod) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *studentMethod) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

func method() {
	var s1 = studentMethod{"jhon wick", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan", name)

	fmt.Println("--------------")

	var s2 = studentMethod{"lauren dimas", 22}
	fmt.Println("s2 before", s2.name)
	// lauren dimas

	s2.changeName1("yogi pratama")
	fmt.Println("s2 after changeName1", s2.name)
	// lauren dimas

	s2.changeName2("yogi pratama")
	fmt.Println("s2 after changeName2", s2.name)
	// yogi pratama

	var s3 = studentMethod{"lauren dimas", 23}
	s3.sayHello()

	var s4 = &studentMethod{"dimas yogi", 24}
	s4.sayHello()
}
