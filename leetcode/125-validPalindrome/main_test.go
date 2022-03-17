package main

import "testing"

func Test_removeNonAlphaNumeric(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "leetcode: example1 ", args: args{s: "A man, a plan, a canal: Panama"}, want: "AmanaplanacanalPanama", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeNonAlphaNumeric(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("removeNonAlphaNumeric() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("removeNonAlphaNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
