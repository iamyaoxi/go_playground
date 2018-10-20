package playground

import "testing"
import "time"
import "math/rand"

/** Test case:
 * [0] => true
 * [1, 0] => true
 * [2, 3, 1, 1, 4] => true
 * [2,0,5,0,1,1,0,0] => true
 * [3, 3, 5, 0, 0, 1, 0, 2, 5] => true;
 * [0, 1] => false
 * [3, 2, 1, 0, 4] => false
 * [3, 3, 3, 0, 0, 0, 1] => false
 * [1,0,1,0,1,0,1,0,1] => false
 * [10,9,8,7,6,5,4,3,2,1,0,1] => false
 */
func TestJump(t *testing.T) {

	testArray := []int{2,3,1,1,4}
	result := Jump(testArray)
	if !result {
		t.Errorf("Should be true")
	}

	testArray = []int{0}
	result = Jump(testArray)
	if !result {
		t.Errorf("Should be true")
	}

	testArray = []int{2,0,5,0,1,1,0,0}
	result = Jump(testArray)
	if !result {
		t.Errorf("Should be true")
	}
	
	testArray = []int{3, 3, 5, 0, 0, 1, 0, 2, 5}
	result = Jump(testArray)
	if !result {
		t.Errorf("Should be true")
	}

	testArray = []int{0, 1}
	result = Jump(testArray)
	if result {
		t.Errorf("Should be false")
	}

	testArray = []int{3, 2, 1, 0, 4}
	result = Jump(testArray)
	if result {
		t.Errorf("Should be false")
	}

	testArray = []int{3, 3, 3, 0, 0, 0, 1}
	result = Jump(testArray)
	if result {
		t.Errorf("Should be false")
	}

	testArray = []int{1,0,1,0,1,0,1,0,1}
	result = Jump(testArray)
	if result {
		t.Errorf("Should be false")
	}

	testAdvancedArray := make([]int, 1000)
	for i:=1; i<=1000; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		randNumber := rand.New(s1).Intn(8)
		testAdvancedArray[i-1] = randNumber
	}
	result = Jump(testAdvancedArray)
}