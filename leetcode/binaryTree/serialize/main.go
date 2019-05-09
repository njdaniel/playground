package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	s := "[1, 2, 3]"
	fmt.Println(Deserialize(s))
	
}

// Deserialize func changes a bfs slice to a linked list
func Deserialize(s string) *TreeNode {
	//ss := make([]rune, 0)
	//for _, x := range s {
	//	ss = append(ss, x)
	//}
	//ss = ss[1:len(s)-1]
	s = s[1:len(s)-1]
	fmt.Println(s)
	ss := strings.Split(s, ",")
	fmt.Println(ss)

	return &TreeNode{1, nil, nil}
}

// Serialize func changes linked list to a serialized string
func Serialize(root *TreeNode) string {
	fmt.Println(root)
	return ""
}
