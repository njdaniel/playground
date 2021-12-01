package main

import "testing"

func Test_isLonely(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isLonely(tt.args.node); got != tt.want {
				t.Errorf("isLonely() = %v, want %v", got, tt.want)
			}
		})
	}
}
