package puzzle

import "testing"

func TestPalindrome(t *testing.T) {
	
	testCase1 := PalindromeGame("ABBA")
	if !testCase1 {
		t.Errorf("ABBA is Palindromic")
	}

	testCase2 := PalindromeGame("ABCBA")
	if !testCase2 {
		t.Errorf("ABCBA is Palindromic")
	}

	testCase3 := PalindromeGame("ADA")
	if !testCase3 {
		t.Errorf("ADA is Palindromic")
	}

	testCase4 := PalindromeGame("A")
	if !testCase4 {
		t.Errorf("A is Palindromic")
	}

	testCase5 := PalindromeGame("ABAB")
	if testCase5 {
		t.Errorf("ABAB is not Palindromic")
	}
}