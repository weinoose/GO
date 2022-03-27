package src

import "fmt"

func Matrixes() {

	// Value Initialized Matrixes

	var a = [3][4]int{
		{0, 1, 2, 3},   // Initializers for row indexed by 0.
		{4, 5, 6, 7},   // Initializers for row indexed by 1.
		{8, 9, 10, 11}} // Initializers for row indexed by 2.

	fmt.Println(a[2])         // 3rf row of matrix.
	fmt.Println(a[2][3])      // 4th element of the 3rd row of matrix.
	fmt.Println(a[0:2])       // 1st and 2nd rows of matrix.
	fmt.Println(a[0:2][1])    // 2nd row of 1st and 2nd rows of matrix.
	fmt.Println(a[0:2][1][3]) // 4th element of 2nd row of 1st and 2nd rows of matrix.

	// Value not Initialized Matrixes

	var b [8][3]float64

	fmt.Println(b[6])         // 7th row of matrix.
	fmt.Println(b[5][1])      // 2nd element of the 6th row of matrix.
	fmt.Println(b[3:6])       // 4th and 6th rows of matrix.
	fmt.Println(b[1:7][3])    // 4th row of 2nd and 7th rows of matrix.
	fmt.Println(b[1:7][3][2]) // 3rd element of 4th row of 2nd and 7th rows of matrix.

}
