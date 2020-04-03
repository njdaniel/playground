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
	sumNode := new(ListNode)
	sum := l1.Val +l2.Val
	digit := sum % 10
	remainder := sum - digit
	sumNode.Val = digit
	l1 = l1.Next
	l2 = l2.Next
	for {
		l1digit := 0
		l2digit := 0
		// if both l1 and l2 are nil; break
		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			l1digit = 0
		} else {
			l1digit = l1.Val
		}
		if l2 == nil {
			l2digit = 0
		} else {
			l2digit = l2.Val
		}
		//get digitsum and remainder
		sum = l1digit + l2digit + remainder
		digit = sum % 10
		remainder = sum - digit
		newSumNode := new(ListNode)
		sumNode.Next = newSumNode
		newSumNode.Val = digit
		//go to next node
		l1 = l1.Next
		l2 = l2.Next
	}
	return sumNode
}
