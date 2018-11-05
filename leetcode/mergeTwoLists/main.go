package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n12 := ListNode{
		Val: 3,
	}
	n11 := ListNode{
		Val:  1,
		Next: &n12,
	}

	n22 := ListNode{
		Val: 4,
	}
	n21 := ListNode{
		Val:  2,
		Next: &n22,
	}
	lst := mergeTwoLists(&n11, &n21)
	for lst != nil {
		fmt.Println(lst.Val)
		lst = lst.Next
	}
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head ListNode
	var p *ListNode
	p = &head
	for l1 != nil && l2 != nil {
		if l1.Val == l2.Val {
			temp := l1.Next
			temp2 := l2.Next
			p.Next = l1
			p = l1
			p.Next = l2
			p = l2
			p.Next = nil
			l1 = temp
			l2 = temp2
		} else if l1.Val > l2.Val {
			temp := l2.Next
			p.Next = l2
			p = l2
			p.Next = nil
			l2 = temp
		} else if l1.Val < l2.Val {
			temp := l1.Next
			p.Next = l1
			p = l1
			p.Next = nil
			l1 = temp
		}
	}
	// append non nil list to p
	if l1 == nil {
		p.Next = l2
	}
	if l2 == nil {
		p.Next = l1
	}
	return head.Next
}
