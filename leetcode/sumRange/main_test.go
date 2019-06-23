package main

import (
	"reflect"
	"testing"
)

func Test_summaryRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "ex1", args:args{nums: []int{0,1,2,4,5,7}}, want: []string{"0->2", "4->5", "7"}},
		{name: "ex2", args:args{nums:[]int{0,2,3,4,6,8,9}}, want:[]string{"0","2->4","6","8->9"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRanges(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intToStringRange(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "ex1", args:args{nums: []int{1,2,3}}, want: "1->3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToStringRange(tt.args.nums); got != tt.want {
				t.Errorf("intToStringRange() = %v, want %v", got, tt.want)
			}
		})
	}
}