package golangbasics

import "fmt"

func Operators() {
	x := true
	y := false

	/* LOGICAL OPERATORS */

	statement1 := x && y /* AND */
	statement2 := x || y /* OR */
	statement3 := x != y /* NOT EQUAL */
	statement4 := x == y /* EQUAL */

	fmt.Println(statement1)
	fmt.Println(statement2)
	fmt.Println(statement3)
	fmt.Println(statement4)

	z := 12
	t := 17

	/* MATHEMATICAL OPERATORS */

	statement5 := z > t  /* GREATER THAN */
	statement6 := z < t  /* LOWER THAN */
	statement7 := z != t /* NOT EQUAL */
	statement8 := z == t /* EQUAL */

	fmt.Println(statement5)
	fmt.Println(statement6)
	fmt.Println(statement7)
	fmt.Println(statement8)
}
