package src

import (
	"fmt"
	"os"
)

func ErrorHandle() {
	ErrorHandling()
	PanicKeyword()
}

func ErrorHandling() {

	// Basic Error handling phrases.

	f, err := os.Open("file.txt")

	errornum := 0

	defer fmt.Println("Execution finished with", errornum, "error counts.")
	// In python, except keyword equals GoLang err != nil method.
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("File has not found", pErr.Path)
			errornum++
		}
		// In python, try keyword equals GoLang err != nil's else.
	} else {
		fmt.Println("File has found.")
		fmt.Println(f.Name())
		errornum += 0
	}
}

func PanicKeyword() {

	// Working with panic keyword.

	/*
		The Panic keyword is used to terminate the program with a special error message.
		When the Panic keyword is encountered, the following sequence of instructions is followed:
		1 - Execution of current function stops
		2 - Any function called with the defer keyword is executed
		3 - Program execution is terminated.
	*/

	num := -6

	if num < 0 {
		panic("This is a Panic situation.")
		// fmt.Println("This line should not be written under right circumstances.")
	} else {
		fmt.Println("Number is positive.")
	}
}
