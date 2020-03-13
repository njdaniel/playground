package substring

import (
	"playground/leetcode/1332-RemovePalindromicSubseq/substring"
	"testing"
)

func Test_removePalindromeSub(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := main.removePalindromeSub(tt.args.s); got != tt.want {
				t.Errorf("removePalindromeSub() = %v, want %v", got, tt.want)
			}
		})
	}
}