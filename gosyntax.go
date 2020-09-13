package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
	fmt.Println("Stored!!")
}

func main() { //capitalizing a funtion or anything will make go export it
	fmt.Println("A numer from 1 - 100000", rand.Intn(100000))
}

// use godoc to learn
// godoc fmt println
