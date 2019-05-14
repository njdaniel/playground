package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	s := "[1, 2, 3]"
	node := Deserialize(s)
	fmt.Println(node)
	fmt.Println(*node.Left)
	fmt.Println(*node.Right)
	
}

// Deserialize func changes a bfs slice to a linked list
func Deserialize(s string) *TreeNode {
	// serialize to string slice
	s = s[1:len(s)-1]
	fmt.Println(s)
	ss := strings.Split(s, ", ")
	fmt.Println(ss)

	// Put in queue
	q := make([]*TreeNode, 0)
	for _, x := range ss {
		fmt.Printf("x = %v\t with type %T\n", x, x)
		val, err := strconv.Atoi(x)
		if err != nil {
			fmt.Printf("cant convert string to int: %v", err)
		}
		fmt.Printf("atoi value = %v\n", val)
		n := TreeNode{Val: val}
		//fmt.Printf("n val: %v\n", n.Val)
		q = append(q, &n)
	}

	// connect
	for k, x := range q {
		//fmt.Println("in q loop")
		//fmt.Printf("val: %v", x.Val)
		var pLeft *TreeNode
		var pRight *TreeNode
		if  k+1 > len(q)-1 {
			pLeft = nil
			pRight = nil
			continue
		}
		if  k+2 > len(q)-1 {
			pRight = nil
			continue
		}
		pLeft = q[k+1]
		pRight = q[k+2]
		x.Left = pLeft
		x.Right = pRight

	}

	return q[0]
}

// Serialize func changes linked list to a serialized string
func Serialize(root *TreeNode) string {
	fmt.Println(root)
	return ""
}
