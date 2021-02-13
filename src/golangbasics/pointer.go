package golangbasics

import (
	"fmt"
)

func Pointer() {
	age := 22
	pAge := &age
	fmt.Println("Address of age", pAge)
	fmt.Println("Adress of age (from pointer)", *pAge)
	lage := age + 1
	pLage := &lage
	fmt.Println("Address of age", pLage)
	fmt.Println("Adress of age (from pointer)", *pLage)
}
