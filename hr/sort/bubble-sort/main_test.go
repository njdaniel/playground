package main

import "testing"

func Test_countSwap(t *testing.T) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Benchmark_countSwap(b *testing.B) {
	type args struct {
		a []int32
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "b1", args:args{a: []int32{3,2,1}}},
	}
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
		})
	}
}