package src

import (
	"fmt"
)

func VariablesAndTypes() {
	// Defining an Data Type v1
	var x string = "Emir Yarkın Yaman"
	var y int = 17
	var z float32 = 1.7
	var f bool = true

	// Defining an data type v2
	t := "Emir Yarkın Yaman"
	n := 17
	m := 1.7
	p := false

	// Defining an data type v2
	var q = "Emir Yarkın Yaman"
	var w = 17
	var r = 1.7
	var l = false

	// String Formatting
	fmt.Printf("Variables are: %s, %d, %f, %v\n", x, y, z, f)
	fmt.Printf("Variables are: %s, %d, %f, %v\n", t, n, m, p)
	fmt.Printf("Variables are: %s, %d, %f, %v\n", q, w, r, l)

	/*

		In string formatting;

		d - decimal integer
		o - octal integer
		O - octal integer with 0o prefix
		b - binary integer
		x - hexadecimal integer lowercase
		X - hexadecimal integer uppercase
		f - decimal floating point, lowercase
		F - decimal floating point, uppercase
		e - scientific notation (mantissa/exponent), lowercase
		E - scientific notation (mantissa/exponent), uppercase
		g - the shortest representation of %e or %f
		G - the shortest representation of %E or %F
		c - a character represented by the corresponding Unicode code point
		q - a quoted character
		U - Unicode escape sequence
		t - the word true or false
		s - a string
		v - default format
		#v - Go-syntax representation of the value
		T - a Go-syntax representation of the type of the value
		p - pointer address
		% - a double %% prints a single %

	*/
}
