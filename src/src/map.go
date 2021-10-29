package src

import "fmt"

func Map() {
	fmt.Println("You are ready to GO.")

	// MAP == Python Dictionary
	dict := make(map[string]string) // SLICING
	dict["Goalkeeper"] = "Alisson Becker"
	dict["Striker"] = "Roberto Firmino"
	fmt.Println(dict)
	fmt.Println("Dictionary contains", len(dict), "values.")
	delete(dict, "Striker") // DELETE'S
	fmt.Println(dict)
	fmt.Println("Dictionary contains", len(dict), "values.")

	value, isContains := dict["Goalkeeper"]
	fmt.Println(value, isContains)
}
