package golangbasics

import "fmt"

func Conditionals() {
	// Practice 1

	fmt.Print("Insert a number: ")
	var number int
	fmt.Scanln(&number)

	if number > 10 {
		fmt.Println("Your number is greater than 10.")
	} else {
		fmt.Println("Your number is lower than 10.")
	}

	// Practice 2

	fmt.Print("Insert a number: ")
	var number1 int
	fmt.Scanln(&number1)

	if number < 0 {
		fmt.Println("Your number is NEGATIVE.")
	} else if number > 0 {
		fmt.Println("Your number is POSITIVE.")
	} else {
		fmt.Println("0 is 0. DAMNASS")
	}
}
