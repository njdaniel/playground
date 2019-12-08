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


