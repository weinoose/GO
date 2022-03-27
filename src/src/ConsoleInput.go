package src

import "fmt"

func ConsoleInput() {
	fmt.Print("Insert your name here: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is: " + name)

	fmt.Print("Insert your age here: ")
	var age string
	fmt.Scanln(&age)
	fmt.Println("Your age is: " + name)
}
