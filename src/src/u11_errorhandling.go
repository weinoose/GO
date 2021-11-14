package src

import (
	"fmt"
	"os"
)

func ErrorHandling() {
	BasicErrorHandling()
	WorkWithPanic()
	fmt.Println()
}

func BasicErrorHandling() {

	// Basic Error handling.

	f, err := os.Open("file.txt")

	errornum := 0

	defer fmt.Println("Execution finished with", errornum, "error counts.")

	if err != nil { // Python Except = GoLang err != nil method.
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı...", pErr.Path)
			errornum++
		}
	} else { // Python try = err!=nil's else.
		fmt.Println("Dosya bulundu...")
		fmt.Println(f.Name())
		errornum += 0
	}
}

func WorkWithPanic() {
	// Working with panic keyword.
	/*
		The Panic keyword is used to terminate the program with a special error message.
		When the Panic keyword is encountered, the following sequence of instructions is followed:
		Execution of current function stops
		Any function called with the defer keyword is executed
		Program execution is terminated.
	*/

	num := -6

	if num < 0 {
		panic("This is a Panic situation.")
		// Println statement under panic keyword will not work.
	} else {
		fmt.Println("Number is positive.")
	}
}
