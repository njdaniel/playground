package main

import "testing"

func Test_longestPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "lc ex2: ", args: args{s: "a"}, want: 1},
		{name: "lc ex3: ", args: args{s: "bb"}, want: 2},
		{name: "lc ex1: ", args: args{s: "abccccdd"}, want: 7},
		{name: "lc ex1: ", args: args{s: "dbccccde"}, want: 7},
		{name: "lc ex1: ", args: args{s: "aaaa"}, want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
