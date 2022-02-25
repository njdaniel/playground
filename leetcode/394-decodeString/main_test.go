package main

import (
	"reflect"
	"testing"
)

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "example1", args: args{s: "3[a]2[bc]"}, want: "aaabcbc"},
		{name: "example2", args: args{s: "3[a2[c]]"}, want: "accaccacc"},
		{name: "example3", args: args{s: "2[abc]3[cd]ef"}, want: "abcabccdcdcdef"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func Test_isNumber(t *testing.T) {
//	type args struct {
//		r rune
//	}
//	tests := []struct {
//		name string
//		args args
//		want bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := isNumber(tt.args.r); got != tt.want {
//				t.Errorf("isNumber() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_decode(t *testing.T) {
	type args struct {
		k       int
		encoded []rune
	}
	tests := []struct {
		name string
		args args
		want []rune
	}{
		{name: "ex1", args: args{k: 3, encoded: []rune{'a'}}, want: []rune{'a', 'a', 'a'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.k, tt.args.encoded); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
