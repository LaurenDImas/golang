package main

import "fmt"

func array() {
	// array

	var names [4]string
	names[0] = "Lauren"
	names[1] = "Dimas"
	names[2] = "Yogi"
	names[3] = "Pratama"
	fmt.Println(names[0], names[1], names[2], names[3])

    var fruits = [4]string{"apple","grape","banana","melon"}
    fmt.Println("Jumlah element \t\t", len(fruits))
    fmt.Println("Isi semua element \t", fruits)

    var numbers = [...]int{1,2,3,4}
    fmt.Println("Jumlah element \t\t", len(numbers))
    fmt.Println("Isi semua element \t", numbers)

    var numbers1 = [2][3]int{[3]int{3,2,3},[3]int{3,4,5}}
    var numbers2 = [2][3]int{{3,2,3},{3,4,5}}
    fmt.Println("numbers1", numbers1)
    fmt.Println("numbers2", numbers2)

    for i := 0; i < len(fruits); i++ {
        fmt.Printf("elemen %d : %s\n", i , fruits[i])
    }

    for _, fruit := range fruits {
        fmt.Printf("nama buah : %s\n", fruit)
    }

    var fruits1 = make([]string, 2)
    fruits1[0] = "apple"
    fruits1[1] = "manggo"
    fmt.Println(fruits1)
}
