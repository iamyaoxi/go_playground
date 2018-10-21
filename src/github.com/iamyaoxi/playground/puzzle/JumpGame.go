package puzzle

/**
 * Given an array of non-negative integers, you are initially positioned at the first index of the array.
 * Each element in the array represents your maximum jump length at that position.
 * Create a function to determine if you are able to reach the last index.
 * 
 * 
 * Examples
 * Input: [2,3,1,1,4]
 * Output: true
 * Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

 * Input: [3,2,1,0,4]
 * Output: false
 * Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
 */
func Jump(testArray []int) bool {

	//If only one element, return true
	if len(testArray) <= 1 {
		return true
	}

	//Rather than finding each jumps in each possibility
	//We want to directly know maximum jump in each index
	//Start at [0]
	maxJump := testArray[0]

	//Try to iterate each index
	for i:= 0; i<len(testArray); i++ {
		//See how far this index can go
		maxTemp := testArray[i]+i
		
		//If we stuck in certain index, and no longer can go forward
		//Return false
		if maxJump <= i && testArray[i] == 0 {
			return false
		}

		//Compare the maximum jump in the previous index
		//With current one
		if maxJump < maxTemp {
			//Replace current is higher
			maxJump = maxTemp
		}

		//If we can reach the last index
		//Directly return true
		if maxJump >= len(testArray)-1 {
			return true
		}
	}

	//Return false as default
	return false
}