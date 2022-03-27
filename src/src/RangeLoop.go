package src

import "fmt"

func RangeLoop() {
	cars := []string{"Ferrari", "Mercedes", "Alpine", "McLaren", "Aston Martin"}

	for i, car := range cars {
		// Shows index number.
		fmt.Println(i)
		// Shows the element.
		fmt.Println(car)
	}
}
