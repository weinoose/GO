package golangbasics

import "fmt"

func ForRange() {
	fmt.Println("You are ready to GO.")
	cars := []string{"Lamborghini", "Mazda", "Porsche", "Ferrari", "Subaru"}

	for i, car := range cars {
		fmt.Println(i)   // Index Number
		fmt.Println(car) // Element
	}
}
