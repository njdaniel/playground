package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l6 := &ListNode{6, nil}
	l5 := &ListNode{5, l6}
	l4 := &ListNode{4, l5}
	l3 := &ListNode{3, l4}
	l2 := &ListNode{2, l3}
	l := &ListNode{1, l2}
	nl := oddEvenList(l)
	//fmt.Printf("%d->%d->%d", nl, nl.Next, nl.Next.Next)
	p := nl
	for p != nil {
		fmt.Println(p)
		p = p.Next
	}
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