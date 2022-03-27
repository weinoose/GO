package src

import "fmt"

func ConstantVariables() {
	const PI float64 = 3.14
	const EULER float64 = 2.71
	var area float64

	area = PI * EULER
	fmt.Printf("value of area : %f", area)
}
