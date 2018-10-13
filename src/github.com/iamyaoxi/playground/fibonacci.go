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