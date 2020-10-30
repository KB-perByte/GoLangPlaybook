package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Think of a number between 1 and 100:")
	scanner.Scan()

	guess := 50

	for {
		fmt.Println("I guess the number is", guess)
		fmt.Println("Is that:")
		fmt.Println("(a) too hight")
		fmt.Println("(b) too low")
		fmt.Println("(c) correct")

		scanner.Scan()

		response := scanner.Text()

		if response == "a" {
			guess--
		} else if response == "b" {
			guess++
		} else if response == "c" {
			fmt.Println("I won!")
			break
		} else {
			fmt.Println(" Choose withing the options")
		}
	}
}
