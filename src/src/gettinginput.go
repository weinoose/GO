package src

import "fmt"

func GettingInput() {
	fmt.Print("Insert your name here: ")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Your name is: " + name)
}
