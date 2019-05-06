package main

import (
	"fmt"
)

func main() {
	l1 := &TreeNode{1, nil, nil}
	root := TreeNode{5, l1, nil}
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
	count := 0
	_, count = isUni(root, count)
	return count
}

func isUni(n *TreeNode, count int) (bool, int) {
	if n.Left == nil && n.Right == nil {
		count++
		return true, count
	}
	l, _ := isUni(n.Left, count)
	r, _ := isUni(n.Right, count)
	if ((l == true && n.Val == n.Left.Val) || n.Left == nil) && ((r == true && n.Val == n.Right.Val) || n.Right == nil) {
		count++
		return true, count
	}
	return false, count
}
