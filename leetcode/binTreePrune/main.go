package main

import "reflect"

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}


func main() {
	
}

func pruneTree(root *TreeNode) *TreeNode {
	return root
}

func serialize(n *TreeNode) []interface{} {
	s := make([]interface{}, 0)
	var f func(n *TreeNode)
	f = func(n *TreeNode) {
		if reflect.ValueOf(n.Val).IsNil() {
		 s = append(s, nil)
			
		}
	}
	f(n)
	return s
}