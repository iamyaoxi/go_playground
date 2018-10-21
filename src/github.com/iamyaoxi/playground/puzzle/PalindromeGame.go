package puzzle

//Palindrome is if the words is the same from backward to forward
func PalindromeGame(text string) bool {

	//Know the middle length of the string
	var middle = len(text)/2
	var left, right int
	if len(text) % 2 == 0 {
		left = middle
		right = middle - 1
	} else {
		left = middle - 1
		right = middle + 1
	}

	//Then start checking from left to right 
	//Both going to the each end
	for left >=0 && right < len(text) {
		if text[left] != text[right] {
			return false
		}

		left--
		right++
	}

	//Default return true
	return true
}