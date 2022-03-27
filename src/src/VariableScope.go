package src

import "fmt"

func VariableScope() {
	// Local Variable Decleration
	var a, b, c int

	// Actual Initialization
	a = 10
	b = 20
	c = a + b

	fmt.Printf("value of a = %d, b = %d and c = %d\n", a, b, c)
}
