package main

func main() {
	
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node.Val >= L && node.Val <= R {
			sum += node.Val
		}
		if node.Left != nil {
			f(node.Left)
		}
		if node.Right != nil {
			f(node.Right)
		}
	}
	f(root)
	return sum
}
