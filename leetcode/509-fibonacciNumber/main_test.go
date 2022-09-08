package main

import (
	"testing"
)

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "0", args: args{n: 0}, want: 0},
		{name: "1", args: args{n: 1}, want: 1},
		{name: "2", args: args{n: 2}, want: 1},
		{name: "3", args: args{n: 3}, want: 2},
		{name: "4", args: args{n: 4}, want: 3},
		{name: "5", args: args{n: 5}, want: 5},
		{name: "6", args: args{n: 6}, want: 8},
		{name: "7", args: args{n: 7}, want: 13},
		{name: "8", args: args{n: 8}, want: 21},
		{name: "9", args: args{n: 9}, want: 34},
		{name: "10", args: args{n: 10}, want: 55},
		{name: "11", args: args{n: 11}, want: 89},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibOpt(t *testing.T) {
	type args struct {
		n int
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
			if got := fibOpt(tt.args.n); got != tt.want {
				t.Errorf("fibOpt() = %v, want %v", got, tt.want)
			}
		})
	}
}
