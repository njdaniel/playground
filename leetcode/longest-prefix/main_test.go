package main

import (
	"strings"
	"testing"
)

func TestlongestCommonPrefix(t *testing.T) {
	input := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(input)
	if result != "fl" {
		t.Errorf(`longestCommonPrefix(%s) == %s. Should be "fl" `, strings.Join(input, ","), result)
	}
	input2 := []string{"dog", "racecar", "car"}
	result2 := longestCommonPrefix(input)
	if result2 != "" {
		t.Errorf(`longestCommonPrefix(%s) == %s. Should be "" `, strings.Join(input2, ","), result2)
	}
}
