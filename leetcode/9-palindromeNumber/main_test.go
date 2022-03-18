package main

import "testing"

func Test_isPalindromeNoString(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "leetcode example1: ", args: args{x: 121}, want: true},
		{name: "leetcode example2: ", args: args{x: -121}, want: false},
		{name: "leetcode example3: ", args: args{x: 10}, want: false},
		{name: "example4: even success", args: args{x: 1221}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeNoString(tt.args.x); got != tt.want {
				t.Errorf("isPalindromeNoString() = %v, want %v", got, tt.want)
			}
		})
	}
}
