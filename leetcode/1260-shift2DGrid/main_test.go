package main

import (
	"reflect"
	"testing"
)

func Test_shiftGrid(t *testing.T) {
	type args struct {
		grid [][]int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "test0: simple shift",
			args:args{grid: [][]int{{1, 2, 3,},
				{3, 4, 5,},
				{7, 8, 9,}}},
				want: [][]int{{9, 1, 2,},
					{3, 4, 5,},
					{6, 7, 8}}},
		{name: "one column", args:args{
			grid: [][]int{{1}, {2}, {3}, {4}, {7}, {6}, {5}},
			k:   23,
		}, want: [][]int{{6}, {5}, {1}, {2}, {3}, {4}, {7}}},
		{name: "large k", args:args{
			grid: [][]int{{1}},
			k:    100,
		}, want: [][]int{{1}}},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftGrid(tt.args.grid, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shiftGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}