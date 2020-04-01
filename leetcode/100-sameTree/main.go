package main

func main() {
	
}

/**
* Definition for a binary tree node.
*/
 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }
// isSameTree check if the 2 binary are the same
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
