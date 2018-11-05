package main

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome(121) {
		t.Error(`IsPalindrome(121) = false. Should be True`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome(-121) {
		t.Error(`IsPalindrome(-121) = true`)
	}
}
