package playground

//Binary steps that returns the 
func SimpleBinarySearch(maxRange int) func(bool) int {

	//We define minimum and maximum value
	minValue := 0
	maxValue := maxRange 
	//And also the current guess
	currentGuess := maxValue

	//Search algorithm
	return func(smaller bool) int {
		//If it's smaller
		if (smaller) {
			maxValue = currentGuess-1
		} else { //If it's bigger
			minValue = currentGuess+1
		}

		currentGuess = (maxValue + minValue) / 2
		return currentGuess
	}
}