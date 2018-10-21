package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	if !IsPalindrome("121") {
		t.Error(`IsPalindrome(121) = false. Should be True`)
	}
}
