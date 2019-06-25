package palindromeNum

import "testing"

func TestPalindrome(t *testing.T) {
	res := isPalindrome(123)
	if res != false {
		t.Error("Expected false, got ", res)
	}
}

func TestPalindrome2(t *testing.T) {
	res := isPalindrome(10)
	if res != false {
		t.Error("Expected false, got ", res)
	}
}
