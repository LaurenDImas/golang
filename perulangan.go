package main

import "fmt"

func perulangan() {
	//for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// for argumen
	var a = 0
	for a < 5 {
		fmt.Println("angka", a)
		a++
	}

	// for tanpa argumen
	var u = 0
	for {
		fmt.Println("Angka", u)
		u++
		if u == 5 {
			break
		}
	}

	// for break & continue
	for e := 1; e < 10; e++ {
		if e%2 == 1 {
			continue
		}

		if e > 8 {
			break
		}

		// fmt.Println("angka", e)
	}

	// perulangan bersarang
	/*
		       . . . . .
			   . . . .
			   . . .
		       . .
		       .
	*/
	for b := 0; b < 5; b++ {
		for c := b; c < 5; c++ {
			fmt.Print(c, " ")
		}
		fmt.Println()
	}
	fmt.Println("---------")
	/*
		       . . . . .
			     . . . .
			       . . .
			         . .
			           .
	*/
	for d := 0; d < 5; d++ {
		for e := 0; e < d; e++ {
			fmt.Printf("  ")
		}

		for f := 4; f >= d; f-- {
			fmt.Print(f, " ")
		}
		fmt.Println()
	}
	fmt.Println("---------")
	/*
		       . . . . . . . . . .
			   . . . .     . . . .
			   . . .         . . .
		       . .             . .
		       .                 .
	*/
	for g := 0; g < 5; g++ {
		for h := g; h < 5; h++ {
			fmt.Print(h, " ")
		}

		for k := 0; k < g; k++ {
			fmt.Print("    ")
		}

		for h := g; h < 5; h++ {
			fmt.Print(h, " ")
		}

		fmt.Println()

	}
	/*
				       . . . . . . . . . .
					   . . . .     . . . .
					   . . .         . . .
				       . .             . .
				       .                 .
		               . .             . .
		               . . .         . . .
		               . . . .     . . . .
		               . . . . . . . . . .
	*/
	for g := 0; g < 5; g++ {
		for h := g; h < 5; h++ {
			fmt.Print(h, " ")
		}

		for k := 0; k < g; k++ {
			fmt.Print("    ")
		}

		for h := g; h < 5; h++ {
			fmt.Print(h, " ")
		}

		fmt.Println()

	}

	for g := 0; g < 5; g++ {
		for h := 0; h <= g; h++ {
			fmt.Print(h, " ")
		}

		for k := g; k < 4; k++ {
			fmt.Print("    ")
		}

		for h := 0; h <= g; h++ {
			fmt.Print(h, " ")
		}

		fmt.Println()
	}

	// label dalam perulangan
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

}
