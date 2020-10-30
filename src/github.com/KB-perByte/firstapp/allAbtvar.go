package main

import (
	"fmt"
)

var (
	actorName    string = "SRK"
	comapanion   string = "Parrot"
	doctorNumber int    = 3
	season       int    = 2
	flag         bool   = false
)

var (
	counter int = 0
)

var i int = 27

func main() { //shadowing variable in the inner scope takes precedience
	x := 1 == 1
	y := 2 == 3
	var myFLAG bool
	var i int = 42
	fmt.Println("%v", myFLAG)
	fmt.Println("%v, %T", x, x)
	fmt.Println("%v, %T", y, y)
	fmt.Println(i)

}
