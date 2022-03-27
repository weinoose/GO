package src

import "fmt"

func Arrays() {

	// Creating variety of arrays.
	var A1 [10]int

	// Declaring our independet variables to put elements into arrays.
	var i, j int

	for i = 0; i < 10; i++ {
		// Setting element at location i to i + 100
		A1[i] = i + 100
	}

	// Printing the output.
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, A1[j])
	}
}

func Slicing() {
	// Creating variety of arrays.
	var A2 = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var A3 = []float32{-92.5, 25.5}

	// Basic indexing of elements.
	fmt.Println(A2[3])
	fmt.Println(A3[1])

	// Series of indexing of elements.
	// [a:b] means a includes, b not.
	fmt.Println(A2[2:4])
	fmt.Println(A2[:4])
	fmt.Println(A2[:1])
	fmt.Println(A3[0:1])
	fmt.Println(A3[0:])
	fmt.Println(A3[:])
}
