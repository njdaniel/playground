package main

import "testing"

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		S string
		T string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{T: "abc", S: "abc"}, want: true},
		{name: "test2", args: args{S: "ab#c", T: "ad#c"}, want: true},
		{name: "test3: failure", args: args{S: "abc", T: "adc"}, want: false},
		{name: "test4: multiple backspaces", args: args{S: "ab##c", T: "c"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backspaceCompare(tt.args.S, tt.args.T); got != tt.want {
				t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
			}
		})
	}
}