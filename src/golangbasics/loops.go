package golangbasics

import "fmt"

func Loops() {
	// FOR LOOP - INTEGER

	i := 0

	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// FOR LOOP - ARRAYS

	cars := [5]string{"Ferrari", "Mazda", "Subaru", "Porsche", "Lamborghini"}

	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}
}
