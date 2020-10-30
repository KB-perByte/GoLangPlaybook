package main

import "fmt"

func main() {

	// var a  int8 // - 127 , 127
	// var a  int8 // - 127 , 127
	// var b  uint8 // 0 255
	// var c  int16
	// var d  uint16

	//math is not real
	// a := 2.0
	// b := 3.0
	// r := a / b
	// fmt.Printf("%.20f\n", r)

	i := 0
	//infinite loop
	// for { //killed  my ram :D need to see why
	// 	i++
	// 	fmt.Println(i)
	// }

	for i < 10 {
		fmt.Println("Yey ", i)
		i++
	}

	for i := 0; i < 10; i = i + 1 {
		fmt.Println("Nan", i)
	}

	x := 5
	if x > 5 {
		fmt.Println("Hi")
	} else if x == 6 {
		fmt.Println("EH")
	} else {
		fmt.Println("bye")
	}
}
