package main

import (
	"reflect"
	"testing"
)

func Test_numSmallerByFrequency(t *testing.T) {
	type args struct {
		queries []string
		words   []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "ex2", args: args{
			queries: []string{"bba","abaaaaaa","aaaaaa","bbabbabaab","aba","aa","baab","bbbbbb","aab","bbabbaabb"},
			words: []string{"aaabbb","aab","babbab","babbbb","b","bbbbbbbbab","a","bbbbbbbbbb","baaabbaab","aa"}},
		want: []int{6,1,1,2,3,3,3,1,3,2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSmallerByFrequency(tt.args.queries, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numSmallerByFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		a      []int
		search int
	}
	tests := []struct {
		name            string
		args            args
		wantResult      int
		wantSearchCount int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotSearchCount := binarySearch(tt.args.a, tt.args.search)
			if gotResult != tt.wantResult {
				t.Errorf("binarySearch() gotResult = %v, want %v", gotResult, tt.wantResult)
			}
			if gotSearchCount != tt.wantSearchCount {
				t.Errorf("binarySearch() gotSearchCount = %v, want %v", gotSearchCount, tt.wantSearchCount)
			}
		})
	}
}