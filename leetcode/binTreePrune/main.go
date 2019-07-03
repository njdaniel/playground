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

}

func serialize(n *TreeNode) []interface{} {
	s := make([]interface{}, 0)
	var f func(n *TreeNode)
	f = func(n *TreeNode) {
		switch reflect.ValueOf(n) {
		case :
			
		}
	}
	
	return s
}