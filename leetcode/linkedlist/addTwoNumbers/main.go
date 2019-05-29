package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	l3 := &ListNode{1, nil}
	l2 := &ListNode{2, l3}
	l := &ListNode{3, l2}
	convertLinkedListToInt(l)
}

//addTwoNumbers adds two linkedlists numbers that store each digit in reverse order.
//
//ex: input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//    output: 7 -> 0 -> 8
//    explanation: 342 + 465 = 807
func addTwoNumbers(l1 *ListNode, l2 *ListNode)  *ListNode {
	// convert listnodes to int
	int1, err := convertLinkedListToInt(l1)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	int2, err := convertLinkedListToInt(l2)
	if err != nil {
		fmt.Printf("error %v", err)
	}
	//add two numbers
	sum := int1 + int2
	// convert to linkedlist


	return &ListNode{0, nil}
}

func convertLinkedListToInt(l *ListNode) (int, error) {
	n := make([]string, 0)
	for p := l; p != nil; {
		valString := strconv.Itoa(p.Val)
		n = append([]string{valString}, n...)
		p = p.Next
	}
	ns := strings.Join(n, "")
	a := fmt.Sprintf("%v", ns)
	fmt.Printf("%v \t as type %T ", a, a)
	ni, err := strconv.Atoi(ns)
	if err != nil {
		fmt.Errorf("could convert string back to int: %v", err)
	}
	return ni, err
}

func convertIntToLinkedList(n int) *ListNode {

}