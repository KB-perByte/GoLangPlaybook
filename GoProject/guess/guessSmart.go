package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	low := 1
	high := 100

	fmt.Println("Think of a number between ", low, " and", high)
	scanner.Scan()
	inp, err := strconv.Atoi(scanner.Text())

	if inp <= high && inp >= low {

		if err != nil {

			println("Fault in code")
			os.Exit(0)

		}

		for { //Binary search

			guess := (low + high) / 2
			fmt.Println("I guess the number is", guess)
			fmt.Println("Is that:")
			fmt.Println("(a) too hight")
			fmt.Println("(b) too low")
			fmt.Println("(c) correct")

			scanner.Scan()

			response := scanner.Text()

			if response == "a" {
				high = guess - 1
			} else if response == "b" {
				low = guess + 1
			} else if response == "c" {
				fmt.Println("I won!")
				break
			} else {
				fmt.Println(" Choose withing the options")
			}
		}

	} else {

		fmt.Println("STOP CHEATING !!")

	}
}
