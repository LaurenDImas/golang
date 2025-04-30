package main

import (
	"fmt"
	"reflect"
)

type studentReflect struct {
	Name  string
	Grade int
}

func (s *studentReflect) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama       :", reflectType.Field(i).Name)
		fmt.Println("tipe data  :", reflectType.Field(i).Type)
		fmt.Println("nilai      :", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *studentReflect) SetName(name string) {
	s.Name = name
}

func reflects() {
	// var number = 23
	// var reflectValue = reflect.ValueOf(number)
	// fmt.Println("tipe variable :", reflectValue.Type())
	// fmt.Println("tipe kind :", reflectValue.Kind())
	// fmt.Println("nilai variable1 :", reflectValue.Interface())
	// fmt.Println("nilai variable2 :", reflectValue.Interface().(int))
	// if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("nilai variable :", reflectValue.Int())
	// }

	// // property variable object
	// fmt.Println("--------------")
	// var s1 = &studentReflect{Name: "lauren", Grade: 2}
	// s1.getPropertyInfo()

	// var s2 = &studentReflect{Name: "jhon wick", Grade: 2}
	// fmt.Println("nama :", s2.Name)
	// var reflectValue1 = reflect.ValueOf(s2)
	// var method = reflectValue1.MethodByName("SetName")
	// method.Call([]reflect.Value{
	// 	reflect.ValueOf("wick"),
	// })
	// fmt.Println("nama :", s2.Name)
}
