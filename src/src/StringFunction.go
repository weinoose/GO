package src

import (
	"fmt"
	"strings"
)

func StringFunction() {
	string := "Emir"

	/*
		Makes the whole string cases to lower.
		result := strings.ToLower(name)

		Makes the whole string cases to upper.
		result := strings.ToLower(name)

		Finds how many "E" there are in the expression, otherwise it returns 0.
		result := strings.Count(name, "z")

		Queries how many "E's" are in the expression, returns True or False.
		result := strings.Contains(name, "z")

		Substitute.
		result := strings.Replace(name, "Emir", "Yarkın", 1)

		İstenen karakterin Index numarasını döndürür.
		result := strings.Index(name, "m")

		Returns the Index number of the desired character.
		result := strings.HasPrefix(name, "E")

		Queries whether it ends with the sent parameter.
		result := strings.HasSuffix(name, "r")

	*/

	result := strings.Split(string, "m")
	fmt.Println(string)
	fmt.Println(result)
}
