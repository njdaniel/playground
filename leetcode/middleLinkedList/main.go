package main

import "fmt"

func main() {
	n := 1
	fmt.Println(n/2)
}

type ListNode struct {
	Val int
	Next *ListNode
}

// middleNode returns the *ListNode in the middle, if two returns the 2nd one
//ex1. [1,2,3] return *2
//ex2. [1,2] return *2
//ex3. [1,2,3,4] return *3
func middleNode(head *ListNode) *ListNode {
	// get the serialized values of the nodes
	s := serialize(head)
	node := head
	// loop through n/2 times and return the node
	for i := 0; i < len(s)/2; i++ {
		if node == nil {
			break
		}
		node = node.Next
	}
	return node
}

func serialize(head *ListNode) []int {
	s := make([]int,0)
	node := head
	for {
		s = append(s, node.Val)
		node = node.Next
		if node == nil {
			break
		}
	}
	return s
}
