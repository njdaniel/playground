package main

import (
	"reflect"
	"testing"
)

func Test_shortestDistanceColor(t *testing.T) {
	type args struct {
		colors  []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestDistanceColor(tt.args.colors, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestDistanceColor() = %v, want %v", got, tt.want)
			}
		})
	}
}
