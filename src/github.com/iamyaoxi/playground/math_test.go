package playground

import "testing"

//Unit Test for Fibonacci function
func TestFibonacci(t *testing.T) {

	//Expected value is set
	expectedFibo := [11]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	actualFibo := make([]int, 11)

	fibo := Fibonacci()
	for i:= 0; i < 11; i++ {
		//assigning the actual value
		actualFibo[i] = fibo()
	}

	for i := range expectedFibo {
		//Throw error if expected != actual
		if expectedFibo[i] != actualFibo[i] {
			t.Errorf("FIBONACCI FAILED -> Expected: %v, Actual: %v", expectedFibo[i], actualFibo[i])
		}
	}
}

//Unit Test for Pascal Triangle Function
func TestPascal(t *testing.T) {
	
	//Expected value
	expectedPascal := [][]int {
		[]int{1},
		[]int{1, 1},
		[]int{1, 2, 1},
		[]int{1, 3, 3, 1},
		[]int{1, 4, 6, 4, 1},
		[]int{1, 5, 10, 10, 5, 1},
		[]int{1, 6, 15, 20, 15, 6, 1},
	}
	var actualPascal [][]int = make([][]int, 14, 20)

	//Start assigning value from PascalTriangle
	pascal := PascalTriangle()
	for i:= 0; i < 7; i++ {
		actualPascal[i] = pascal()
	}

	//Start testing, line by line
	for i := range expectedPascal {
		//Value by value
		for j := range expectedPascal[i] {
			//If not the same, throw an error
			if expectedPascal[i][j] != actualPascal[i][j] {
				t.Errorf("PASCAL FAILED -> On Line: %v. Expected: %v, Actual: %v", i, expectedPascal[i], actualPascal[i])
				break
			}
		}
	}
}