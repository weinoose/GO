package src

import "fmt"

func Conditionals() {

	// 2-Way Conditional

	fmt.Print("Insert a number: ")
	var number int
	fmt.Scanln(&number)

	if number > 10 {
		fmt.Println("Your number is greater than 10.")
	} else {
		fmt.Println("Your number is lower than 10.")
	}

	// 3-Way Conditional

	if number < 0 {
		fmt.Println("Your number is NEGATIVE.")
	} else if number > 0 {
		fmt.Println("Your number is POSITIVE.")
	} else {
		fmt.Println("0 is 0. DAMNASS")
	}

	// Multiple Conditional

	if number%2 == 0 {
		if number > 10 {
			fmt.Println("Your number is even and greater than 10.")
		} else {
			fmt.Println("Your number is even and lower than 10.")
		}
	} else if number%2 == 1 {
		if number > 10 {
			fmt.Println("Your number is odd and greater than 10.")
		} else {
			fmt.Println("Your number is odd and lower than 10.")
		}
	} else {
		fmt.Println("Your number is zero.")
	}
}
