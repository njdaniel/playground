package main

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	//check if
	odd := head
	evenHead := head.Next
	even := odd.Next
	p := evenHead.Next
	counter := 3

	for ;p != nil; counter++{
		switch {
		case counter % 2 == 1 :
			odd.Next = p
			odd = p
			p = p.Next
		case counter % 2 == 0:
			even.Next = p
			even = p
			p = p.Next
		}
	}

	odd.Next = evenHead

	return head
}