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

func TestQueue_Pop(t *testing.T) {
	root := &TreeNode{2, nil, nil}
	type fields struct {
		size  int
		Head  *TreeNode
		queue []*TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		want   *TreeNode
	}{
		// TODO: Add test cases.
		{"one value", fields{size: 1, Head: root, queue: []*TreeNode{root}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				size:  tt.fields.size,
				Head:  tt.fields.Head,
				queue: tt.fields.queue,
			}
			if got := q.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}
