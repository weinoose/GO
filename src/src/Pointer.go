package src

import (
	"fmt"
)

func Pointer() {
	age := 22
	pointerAge := &age
	fmt.Println("Address of age", pointerAge)
	fmt.Println("Value of age (from pointer)", *pointerAge)
	fmt.Println("Value of age (from variable)", age)
	lage := age + 1
	pointerLage := &lage
	fmt.Println("Address of age", pointerLage)
	fmt.Println("Value of age (from pointer)", *pointerLage)
	fmt.Println("Value of age (from variable)", lage)

	// Chapter II

	var a int = 20
	var ip *int
	ip = &a

	fmt.Printf("Address of a variable: %x\n", &a)

	// Address stored in pointer variable.
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	// Access the value using the pointer.
	fmt.Printf("Value of *ip variable: %d\n", *ip)
}
