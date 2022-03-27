package src

import "fmt"

func Factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * Factorial(i-1)
}
func Recursion() {
	var i int = 15
	fmt.Printf("Factorial of %d is %d\n", i, Factorial(i))
}
