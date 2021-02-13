package golangbasics

import (
	"fmt"
	"math"
)

func DataTypes() {
	/* Defining a clear variable */
	var x string = "Emir Yarkın Yaman"
	var y int = 17
	var z float32 = 1.7
	var f bool = true

	/* Defining a close variable*/
	t := "Emir Yarkın Yaman"
	n := 17
	m := 1.7
	p := false

	/* String Formatting Type 1 */
	fmt.Println("Variables are:", "*", x, "*", y, "*", z, "*", t, "*", n, "*", m, "*", f, "*", p)

	/* Mathematical Operators */
	fmt.Println(2 + 5)
	fmt.Println(2 * 5)
	fmt.Println(math.Pow(2, 5))
	fmt.Println(8 / 2)

	/*

		Strange Division Information about GoLang

		int / int Returns int even if it is decimal. If you really want to make a split;
		convert numbers to float32 first, then divide.

	*/

}
