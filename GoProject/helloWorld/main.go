package main

import "fmt"

func main() {
	var i int //long syntax
	i = 42
	fmt.Println(i)

	//fmt.Println("Hello")

	var f float32 = 3.14 //generic declaration
	fmt.Println(f)

	firstName := "Sagar Paul" //implicit syntax
	fmt.Println(firstName)

	b := true //boolean
	fmt.Println(b)

	cxNum := complex(3, 4) //complex number
	fmt.Println(cxNum)

	r, im := real(cxNum), imag(cxNum)
	fmt.Println(r, im)

	var firstNamePtr *string = new(string)
	*firstNamePtr = "Paul"
	fmt.Println(*firstNamePtr)
}
