package optimize

import "testing"

func TestMaxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{prices: []int{7,1,5,3,6,4}}, want: 7},
		{name: "test2", args: args{prices: []int{1,2,3,4,5}}, want: 4},
		{name: "test3", args: args{prices: []int{7,6,5,4,3,2}}, want: 0},
		{name: "test4", args: args{prices: []int{6,1,3,2,4,7}}, want: 7},
		{name: "test5", args: args{prices: []int{6,1,3,2,7}}, want: 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.args.prices); got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}