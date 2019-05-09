package main

import (
	"reflect"
	"testing"
)

func TestDeserialize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"simple", args{"[1, 2, 3]"},
			&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Deserialize(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
