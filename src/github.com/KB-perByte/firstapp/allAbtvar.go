package main

import (
	"fmt"
)

var (
	actorName    string = "SRK"
	comapanion   string = "Parrot"
	doctorNumber int    = 3
	season       int    = 2
)

var (
	counter int = 0
)

var i int = 27

func main() { //shadowing variable in the inner scope takes precedience
	var i int = 42

	fmt.Println(i)

}
