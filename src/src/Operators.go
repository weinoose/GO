package src

import (
	"fmt"
	"math"
)

func Operators() {
	x := true
	y := false

	// Logical Operators: and, or, not equal, equal

	statement1 := x && y
	statement2 := x || y
	statement3 := x != y
	statement4 := x == y

	fmt.Println(statement1)
	fmt.Println(statement2)
	fmt.Println(statement3)
	fmt.Println(statement4)

	z := 12
	t := 17

	// Comprasion Operators

	statement5 := z > t
	statement6 := z < t
	statement7 := z != t
	statement8 := z == t

	fmt.Println(statement5)
	fmt.Println(statement6)
	fmt.Println(statement7)
	fmt.Println(statement8)

	// Mathematical Operators

	var b int = 17
	var c float32 = 1.7

	fmt.Println(b + int(c))
	fmt.Println(b * int(c))
	fmt.Println(math.Pow(float64(b), float64(c)))
	fmt.Println(b / int(c))
	fmt.Println(float64(b) / float64(b))

	/*

		Strange Division Information about GoLang

		int / int Returns int even if it is decimal. If you really want to make a split;
		convert numbers to float32 first, then divide.

	*/
}
