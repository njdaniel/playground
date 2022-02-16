package main

import (
	"reflect"
	"testing"
)

func Test_findBuildings(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "example1", args: args{heights: []int{4, 2, 3, 1}}, want: []int{0, 2, 3}},
		{name: "example2", args: args{heights: []int{4, 3, 2, 1}}, want: []int{0, 1, 2, 3}},
		{name: "example3", args: args{heights: []int{1, 3, 2, 4}}, want: []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBuildings(tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findBuildings() = %v, want %v", got, tt.want)
			}
		})
	}
}
