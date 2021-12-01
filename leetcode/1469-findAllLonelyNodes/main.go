package main

import "fmt"

func main() {

}

// TreeNode defination for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//tags:depth-first-search breadth-first-search tree binary-tree
func getLonelyNodes(root *TreeNode) []int {
	ln := make([]int, 0)
	//is lonely?
	return ln
}

func isLonely(node *TreeNode) bool {
	if node.Left == nil || node.Right == nil {
		fmt.Println("Lonely")
		return true
	}
	return false
}
