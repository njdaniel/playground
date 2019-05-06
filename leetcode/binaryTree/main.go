package main

import (
	"fmt"
)

func main() {
	root := TreeNode{5, nil, nil}
	c := countUnivalSubtrees(&root)
	fmt.Println(c)
}

// TreeNode is the defination of a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	var count int

	return count
}

func isUni(n *TreeNode) bool {
	if n.Left == nil && n.Right == nil {
		return true
	}
	if ((isUni(n.Left) && n.Val == n.Left.Val) || n.Left == nil) && ((isUni(n.Right) && n.Val == n.Right.Val) || n.Right == nil) {
		return true
	}
	return false
}
