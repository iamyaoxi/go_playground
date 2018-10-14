package playground

import (
	"testing"
	"math/rand"
	"time"
)

//Unit Test to find integer in a Binary Test
func TestSimpleBinarySearch(t *testing.T) {
	
	//Find one random integer between 0 - 1 Trillion
	const maxField = 1000000000000
	s1 := rand.NewSource(time.Now().UnixNano())
	randomNumber := rand.New(s1).Intn(maxField)

	simple := SimpleSearch(maxField)
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

		//It shouldn't take more than 40 steps
		//At most it's only 39 Steps
		if (stepsTaken > 40) {
			t.Errorf("Too much steps taken")
		}
	}
}