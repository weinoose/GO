package src

import "fmt"

func Functions() {
	Function()
	fmt.Println(Paramated(10, 15))
	fmt.Println(Returned(3, 2))
	fmt.Println(MReturned(17, 7))
	fmt.Println(Variadic(10, 23, 27, 4, 87, 45, 71, 6, 17, 7))
	fmt.Println(Variadic(10, 23, 27, 4, 87, 7))
}

// Basic Functions

func Function() {
	fmt.Println("Hello, Hi!")
}

// Paramated Functions

func Paramated(num1 int, num2 int) int {
	return num1 + num2
}

// Returned Functions

func Returned(num1 int, num2 int) int {
	total := num1 * num2
	return total
}

// Multiple Returned Functions

func MReturned(num1 int, num2 int) (int, int, int, float32) {
	total := num1 * num2
	total1 := num1 + num2
	total2 := num1 - num2
	total3 := float32(num1) / float32(num2)
	return total, total1, total2, total3
}

// Variadic Functions like in python's *params *kwargs

func Variadic(liste ...int) int {
	sum := 0
	i := 0
	for i < len(liste) {
		sum = sum + liste[i]
		i = i + 1
	}

	return sum
}
