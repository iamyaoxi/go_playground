package playground

import "testing"

//Unit Test for Fibonacci function
func Test(t *testing.T) {
	//Expected value is set
	expected := [11]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	actual := make([]int, 11)

	f := Fibonacci()
	for i:= 0; i < 11; i++ {
		//assigning the actual value
		actual[i] = f()
	}

	for i := range expected {
		//Throw error if expected != actual
		if expected[i] != actual[i] {
			t.Errorf("Expected: %v, Actual: %v", expected[i], actual[i])
		}
	}
}