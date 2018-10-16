package playground

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
func jumpImpl(testArray []int, maxJump, currentIndex int) bool {
	target := len(testArray)
	reachTarget := false
	
	for i := 1; i <= maxJump; i++ {
		idx := currentIndex + i
		
		if idx >= target-1 {
			return true
		}

		nextJump := testArray[idx]
		if nextJump+idx >= target {
			return true
		}

		if nextJump == 0 {
			
			continue
		}

		//Keep looping until you find true or you don't have anything left
		reachTarget = jumpImpl(testArray, nextJump, idx)

		if reachTarget {
			return true
		}
	}

	//I guess you don't have anything left, so return false
	return reachTarget
}

func Jump(testArray []int) bool {

	//If only one element, return true
	if len(testArray) == 1 {
		return true
	}

	//What's the maximum jump
	maxFirstJump := testArray[0]

	//Return false since every possiblity can't go to last jump
	//fmt.Println("Start to Jump until ", maxFirstJump)
    return jumpImpl(testArray, maxFirstJump, 0)
}