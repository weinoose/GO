package golangbasics

import "fmt"

func Functions() {
	normal()
	parameted(10, 15)
	fmt.Println(returnedd(3, 2))
	fmt.Println(mreturnedd(17, 7))
	fmt.Println(variadicc(10, 23, 27, 4, 87, 45, 71, 6, 17, 7))
	fmt.Println(variadicc(10, 23, 27, 4, 87, 7))
}

// NORMAL FUNCTIONS

func normal() {
	fmt.Println("Hello, Hi!")
}

// PARAMETED FUNCTIONS

func parameted(num1 int, num2 int) {
	fmt.Println(num1 + num2)
}

// RETURNED FUNCTIONS

func returnedd(num1 int, num2 int) int {
	total := num1 * num2
	return total
}

// MULTIPLE RETURNED FUNCTIONS

func mreturnedd(num1 int, num2 int) (int, int, int, float32) {
	total := num1 * num2
	total1 := num1 + num2
	total2 := num1 - num2
	total3 := float32(num1) / float32(num2)
	return total, total1, total2, total3
}

// VARIADIC FUNCTIONS ( PYTHON > *params *kwargs FUNCTION )

func variadicc(liste ...int) int {
	sum := 0
	i := 0
	for i < len(liste) {
		sum = sum + liste[i]
		i = i + 1
	}

	return sum
}
