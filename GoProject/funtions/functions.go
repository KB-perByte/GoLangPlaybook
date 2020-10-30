package main

import "fmt"

func sayHello() {
	fmt.Println("Hello")
}

func sayHelloName(name string) {
	fmt.Println("Hello", name)
}

func addOne(x int) int {
	return x + 1
}

func  recSayHello(name string, stopper int) int {
	if stopper == 5:
		return 1
	fmt.Println("->")
	recSayHello("Hey", stopper)
}

func main() {
	sayHello()
	sayHelloName("Sagar")

	x := 5
	x = addOne(x)
	x = addOne(addOne(x))
	fmt.Println(x)
}
