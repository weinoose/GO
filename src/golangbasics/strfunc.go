package golangbasics

import (
	"fmt"
	// "strings"
)

func StrFunc() {
	name := "Emir"
	// result := strings.ToLower(name) // Makes the whole string cases to lower.
	// result := strings.ToLower(name) // Makes the whole string cases to upper.
	// result := strings.Count(name, "z")     // Finds how many "E" there are in the expression, otherwise it returns 0.
	// result := strings.Contains(name, "z") // Queries how many "E's" are in the expression, returns True or False.
	// result := strings.Replace(name, "Emir", "Yarkın", 1) // Substitute.
	// result := strings.Index(name, "m") // İstenen karakterin Index numarasını döndürür.
	// result := strings.HasPrefix(name, "E") // Returns the Index number of the desired character.
	// result := strings.HasSuffix(name, "r") // Queries whether it ends with the sent parameter.
	// result := strings.Split(name, " ")
	fmt.Println(name)
}
