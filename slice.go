package main

import "fmt"

func slice() {
    // var fruits = []string{"apple","grape","banana","melon"} // slice
    // fmt.Println(fruits[0])

    // var fruitsA = []string{"apple","grape"} // slice
    // fmt.Println(fruitsA)
    // var fruitsB = [2]string{"banana","melon"} // array
    // fmt.Println(fruitsB)
    // var fruitsC = [...]string{"grape","banana"} // array
    // fmt.Println(fruitsC)

    // // operasi slice
    // var newFruitsA = fruits[0:2]
    // fmt.Println(newFruitsA)
    // var newFruitsB = fruits[0:4]
    // fmt.Println(newFruitsB)
    // var newFruitsC = fruits[0:0]
    // fmt.Println(newFruitsC)
    // var newFruitsD = fruits[4:4]
    // fmt.Println(newFruitsD)
    // var newFruitsG = fruits[:]
    // fmt.Println(newFruitsG)
    // var newFruitsH = fruits[2:]
    // fmt.Println(newFruitsH)
    // var newFruitsI = fruits[:2]
    // fmt.Println(newFruitsI)

    // slice adalah reference
    // var aFruits = fruits[0:3]
    // var bFruits = fruits[1:4]
    // var aaFruits = aFruits[1:2]
    // var baFruits = bFruits[0:1]
    // fmt.Println(fruits);
    // fmt.Println(aFruits);
    // fmt.Println(bFruits);
    // fmt.Println(aaFruits);
    // fmt.Println(baFruits);
    // baFruits[0] = "pinnaple"
    // fmt.Println(fruits);
    // fmt.Println(aFruits);
    // fmt.Println(bFruits);
    // fmt.Println(aaFruits);
    // fmt.Println(baFruits);

    // fungsi
    // len (total data), cap (total cap arr)
    // fmt.Println(len(fruits))
    // fmt.Println(cap(fruits))
    // fmt.Println(len(aFruits)) // len: 3
    // fmt.Println(cap(aFruits)) // cap: 4
    // fmt.Println(len(bFruits)) // len: 3
    // fmt.Println(cap(bFruits)) // cap: 3

    // append
    // var fruits1 = []string{"apple", "grape", "banana"}
    // var bFruits = fruits1[0:2]
    // fmt.Println(fruits1);
    // fmt.Println(cap(bFruits))
    // fmt.Println(len(bFruits))
    // fmt.Println(fruits1)
    // fmt.Println(bFruits)
    // var cFruits1 = append(bFruits, "papaya")
    // fmt.Println(fruits1)
    // fmt.Println(bFruits)
    // fmt.Println(cFruits1)

    // copy
    dst := make([]string, 3)
    src := []string{"watermelon", "pinnaple", "apple", "orange"}
    n := copy(dst, src)
    fmt.Println(dst)
    fmt.Println(src)
    fmt.Println(n)
    dstA := []string{"potato", "potato", "potato"}
    srcA := []string{"watermelon", "pinnaple"}
    nA := copy(dstA, srcA)
    fmt.Println(dstA)
    fmt.Println(srcA)
    fmt.Println(nA)
}
