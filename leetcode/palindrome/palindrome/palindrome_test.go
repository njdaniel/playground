package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	type args struct {
		runes []rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "letter test true", args:args{runes: []rune{'a', 'b', 'b', 'a'}}, want:true},
		{name: "letter test false", args:args{runes: []rune{'a', 'a', 'b'}}, want:false},
		{name: "number test true", args:args{runes: []rune{'1', '2', '1'}}, want:true},
		{name: "number test false", args:args{runes: []rune{'1', '2', '3'}}, want:false},
		{name: "one letter char", args: args{runes: []rune{'a'}}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.runes); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}