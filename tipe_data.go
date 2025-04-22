package main

import "fmt"

func tipeData() {
	// numeric non decimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	// numeric decimal
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// bool
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	// string
	var message string = "halo"
	fmt.Printf("message: %s \n", message)

	message = `Nama saya "John Wick".
        Salam kenal.
        Mari belajar "Golang".`
	fmt.Println(message)

	// nil & zero value
}
