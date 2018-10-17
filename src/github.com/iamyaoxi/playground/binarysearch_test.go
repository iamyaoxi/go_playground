package playground

import (
	"testing"
	"math/rand"
	"math"
	"time"
)

//Unit Test to find integer in a Binary Test
func TestSimpleBinarySearch(t *testing.T) {
	
	//Find one random integer between 0 - 1 Trillion
	const maxField = 1000000000000
	s1 := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(s1).Intn(maxField)

	simple := SimpleBinarySearch(maxField)
	//starts with true
	smaller := true
	stepsTaken := 0

	//Start searching
	for theGuess := simple(smaller); ; stepsTaken++ {

		//If the guess is correct, quickly break out of the loop
		if theGuess == randomNumber {
			t.Log("GUESS CORRECT! -> Number: ", theGuess, " Steps Taken: ", stepsTaken)
			break
		}

		//Is it smaller?
		smaller = randomNumber < theGuess
		theGuess = simple(smaller)

		//It shouldn't take more than Log2(maxField)
		if stepsTaken > int(math.Log2(maxField)) {
			t.Errorf("Too much steps taken")
		}
	}
}