package playground

//Experiment how to get Fibonacci
func Fibonacci() func() int {
	
	//Fibonacci is always based on sum of current and previous value
	//so let's have them both declared
	previousValue := 0
	currentValue := 0

	//Function closures allow references outside the function body
	return func() int {
		//Result is always current value + previous value
		result := currentValue + previousValue
		//But if it's 0, it will added manually
		if result == 0 {
			previousValue = 1
			return 0	
		}
		
		//Now move those value back
		previousValue = currentValue
		currentValue = result
		
		//And return the result
		return result
	}
}

//Experiment how to print Pascal Triangle
func PascalTriangle() func() []int {
	//Pascal Triangle based on the sum of the previous line
	var previousLine []int
	var currentLine []int

	return func() []int {

		//Add the first 1 in each line
		currentLine = make([]int, len(previousLine) + 1)
		currentLine[0] = 1
		for i := range previousLine {
			if i == len(previousLine) - 1 {
				//Add the last 1 in each line
				currentLine[i+1] = 1
				break
			}
			//Sum the previous line to each other
			sum := previousLine[i] + previousLine[i+1]
			currentLine[i+1] = sum
		}

		previousLine = currentLine
		return currentLine
	}
}