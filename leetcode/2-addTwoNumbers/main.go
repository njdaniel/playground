package main

func main() {
	
}

type ListNode struct {
	Val int
	Next *ListNode
}

// addTwoNumbers given two *Listnodes return the addition
// example
//Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//Output: 7 -> 0 -> 8
//Explanation: 342 + 465 = 807.
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	sum := new(ListNode)
	for {
		//get digitsum and remainder
		digit := l1.Val +l2.Val

		//go to next node

		// if both l1 and l2 are nil; break
	}
	return sum
}
