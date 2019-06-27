package main

import "testing"

func Test_wordPattern(t *testing.T) {
	type args struct {
		pattern string
		str     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"ex1", args{"abab", "cat dog cat dog"}, true},
		{"ex2", args{"abab", "dog dog dog dog"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordPattern(tt.args.pattern, tt.args.str); got != tt.want {
				t.Errorf("wordPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
