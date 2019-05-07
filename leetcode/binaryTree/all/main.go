package main

import "fmt"

// TreeNode is a struct for binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	l1 := &TreeNode{5, nil, nil}
	r1 := &TreeNode{5, nil, nil}
	root := TreeNode{5, l1, r1}
	c := isUnivalTree(&root)
	fmt.Println(c)
}

func isUnivalTree(root *TreeNode) bool {
	vals := dfs(root)
	return sameVal(vals)
}

func dfs(n *TreeNode) []int {
	val := make([]int, 0)
	var f func(n *TreeNode)
	f = func(n *TreeNode) {
		val = append(val, n.Val)
		if n.Left != nil {
			f(n.Left)
		}
		if n.Right != nil {
			f(n.Right)
		}
	}
	f(n)
	return val
}

func sameVal(vals []int) bool {
	val := vals[0]
	for _, x := range vals {
		if val != x {
			return false
		}
	}
	return true
}
