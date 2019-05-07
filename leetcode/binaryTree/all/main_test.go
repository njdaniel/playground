package main

import "testing"

func Test_sameVal(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"same", args{[]int{1, 1, 1}}, true},
		{"dif", args{[]int{1, 2, 1}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sameVal(tt.args.vals); got != tt.want {
				t.Errorf("sameVal() = %v, want %v", got, tt.want)
			}
		})
	}
}
