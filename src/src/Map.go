package src

import "fmt"

func Map() {
	fmt.Println("You are ready to GO.")

	// Map means also dictionaries in python.
	dict := make(map[string]string) // Slicing to creating a dictionary.
	dict["Goalkeeper"] = "Alisson Becker"
	dict["Striker"] = "Roberto Firmino"
	fmt.Println(dict) // Print Method
	fmt.Println("Dictionary contains", len(dict), "values.")
	delete(dict, "Striker") // Delete Method
	fmt.Println(dict)
	fmt.Println("Dictionary contains", len(dict), "values.")

	value, isContains := dict["Goalkeeper"]
	fmt.Println(value, isContains)
}
