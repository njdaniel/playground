package main

import "fmt"

func main() {
	l1 := &TreeNode{5, nil, nil}
	root := TreeNode{5, l1, nil}
	c := countUnivalSubtrees(&root)
	fmt.Println(c)
}

// TreeNode is the definition of a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countUnivalSubtrees(root *TreeNode) int {
	count := 0
	var f func(n *TreeNode) bool
	f = func(n *TreeNode) bool {
		if n.Left == nil && n.Right == nil {
			count++
			return true
		}
		l := false
		r := false
		if n.Left == nil {
			l = true
		} else {
			l = f(n.Left)
		}
		if n.Right == nil {
			r = true
		} else {
			r = f(n.Right)
		}

		if l == true && r == true {
			count++
			return true
		}
		return false
	}
	f(root)
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

func isU(n *TreeNode) int {
	count := 0
	var f func(n *TreeNode) bool
	f = func(n *TreeNode) bool {
		if n.Left == nil && n.Right == nil {
			count++
			return true
		}
		l := false
		r := false
		if n.Left == nil {
			l = true
		} else {
			l = f(n.Left)
		}
		if n.Right == nil {
			r = true
		} else {
			r = f(n.Right)
		}

		if l == true && r == true {
			count++
			return true
		}
		return false
	}
	return count
}

// Deserialize func changes a bfs slice to a linked list
func Deserialize(s string) *TreeNode {

	return &TreeNode{1, nil, nil}
}
