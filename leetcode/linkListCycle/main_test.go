package main

import "testing"

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	
	ll4 := &ListNode{
		Val: 3,
		Next: ll3,
	}
	ll3 := &ListNode{

	}
	ll2 := &ListNode{

	}
	ll1 := &ListNode{

	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		name: "ex1",
		args: args{
			head: &ListNode{
				 &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 0,
						Next: &ListNode{
							Val: -4,
							Next:
						}
					}
				}
					}
		},
		want: true,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
