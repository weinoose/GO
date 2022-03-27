package src

import "fmt"

func ForLoop() {
	// for loop with integer counter.

	i := 0

	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// for loop works with arrays.

	cars := [5]string{"Ferrari", "Mercedes", "Alpine", "McLaren", "Aston Martin"}

	for i := 0; i < len(cars); i++ {
		fmt.Println(cars[i])
	}
}
