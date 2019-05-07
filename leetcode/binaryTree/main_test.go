package main

import "testing"

func Test_isU(t *testing.T) {
	type args struct {
		n *TreeNode
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
			if got := isU(tt.args.n); got != tt.want {
				t.Errorf("isU() = %v, want %v", got, tt.want)
			}
		})
	}
}
