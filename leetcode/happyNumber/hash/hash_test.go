package hash

import "testing"

func TestIsHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "test1", args: args{n: 19}, want: true},
		//{name: "test2", args: args{n: 18}, want: true},
		{name: "test3", args: args{n: 116}, want: false},
		{name: "test4", args: args{n: 7}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHappy(tt.args.n); got != tt.want {
				t.Errorf("IsHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}