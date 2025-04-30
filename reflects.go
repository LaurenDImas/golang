package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
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

type User struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
	Age   int    `validate:"min=18"`
}

func validateStruct(s interface{}) []string {
	var errs []string
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)

	if v.Kind() != reflect.Struct {
		return []string{"input must be a struct"}
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		tag := t.Field(i).Tag.Get("validate")
		fieldName := t.Field(i).Name
		// fmt.Println(field, tag, fieldName)
		rules := strings.Split(tag, ",")

		for _, rule := range rules {
			switch {
			case rule == "required":
				if field.Kind() == reflect.String && field.String() == "" {
					errs = append(errs, fieldName+" is required")
				}
			case rule == "email":
				if !strings.Contains(field.String(), "@") {
					errs = append(errs, fieldName+" must be a valid email")
				}
			case strings.HasPrefix(rule, "min="):
				minVal, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))
				if field.Kind() == reflect.Int && int(field.Int()) < minVal {
					errs = append(errs, fmt.Sprintf("%s must be at least %d", fieldName, minVal))
				}
			}
		}
	}

	return errs
}

func reflects() {
	var user = User{
		Name:  "",
		Email: "inisalahemail",
		Age:   12,
	}

	errors := validateStruct(user)

	if len(errors) > 0 {
		fmt.Println("Validation errors:")
		for _, err := range errors {
			fmt.Println("-", err)
		}
	} else {
		fmt.Println("Validation passed")
	}

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
