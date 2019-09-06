package bank

import "testing"

func TestBalance(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Balance(); got != tt.want {
				t.Errorf("Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeposit(t *testing.T) {
	type args struct {
		amount int
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

func TestWithdraw(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Withdraw(tt.args.amount); got != tt.want {
				t.Errorf("Withdraw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_teller(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
