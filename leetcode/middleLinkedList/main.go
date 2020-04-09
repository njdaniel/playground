package main

import "fmt"

func main() {
	n := 3
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
	// loop through n/2 times and return the node
	return nil
}
