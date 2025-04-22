package main

import "fmt"

func seleksiKondisi() {
	// if, else, else if
	var point = 8
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus, nilai Anda %d\n", point)
	}

	// variable temporary
	var point1 = 8840.0
	if percent := point1 / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// switch case
	var point2 = 6
	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("Not Bad")
	}

	// switch case banyak kondisi
	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("Not Bad")
	}

	// switch case kurung kurawal case & default
	switch point2 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		{
			fmt.Println("Not Bad")
			fmt.Println("you can be better!")
		}
	}

	// switch case gaya if - else
	switch {
	case point2 == 8:
		fmt.Println("perfect")
	case (point2 < 8) && (point2 > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// swicth case fallthrough
	switch {
	case point2 == 0:
		fmt.Println("perfect")
	case (point2 < 8) && (point2 > 3):
		fmt.Println("awesome")
		fallthrough
	case point2 < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// seleksi bersarang
	if point2 > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
