package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//s := "[1, 2, 3, null, null, 4, 5]"
	//node := Deserialize(s)
	//fmt.Println(node)
	//fmt.Println(*node.Left)
	//fmt.Println(*node.Right)

	r3 := &TreeNode{5, nil, nil}
	r1 := &TreeNode{5, nil, r3}
	r2 := &TreeNode{5, nil, nil}
	l2 := &TreeNode{5, nil, nil}
	l1 := &TreeNode{1, l2, r2}
	root := TreeNode{5, l1, r1}

	fmt.Println(Serialize(&root))
}

// Deserialize func changes a bfs slice to a linked list
func Deserialize(s string) *TreeNode {
	// serialize to string slice
	s = s[1 : len(s)-1]
	fmt.Println(s)
	ss := strings.Split(s, ", ")
	fmt.Println(ss)

	// Put in queue
	q := make([]*TreeNode, 0)
	for _, x := range ss {
		fmt.Printf("x = %v\t with type %T\n", x, x)
		val, err := strconv.Atoi(x)
		if err != nil {
			fmt.Printf("cant convert string to int: %v", err)
		}
		fmt.Printf("atoi value = %v\n", val)
		n := TreeNode{Val: val}
		//fmt.Printf("n val: %v\n", n.Val)
		q = append(q, &n)
	}

	// connect
	for k, x := range q {
		//fmt.Println("in q loop")
		//fmt.Printf("val: %v", x.Val)
		var pLeft *TreeNode
		var pRight *TreeNode
		if k+1 > len(q)-1 {
			pLeft = nil
			pRight = nil
			continue
		}
		if k+2 > len(q)-1 {
			pRight = nil
			continue
		}
		pLeft = q[k+1]
		pRight = q[k+2]
		x.Left = pLeft
		x.Right = pRight

	}

	return q[0]
}

// Queue is a linked list queue for bfs
type Queue struct {
	size  int
	Head  *TreeNode
	queue []*TreeNode
}

func NewQueue() *Queue {
	return new(Queue)
}

func (q *Queue) Push(node *TreeNode) {
	q.queue = append(q.queue, node)
	q.Head = q.queue[len(q.queue)-1]
	q.size++
}

func (q *Queue) Pop() *TreeNode {
	node := q.queue[len(q.queue)-1]
	q.queue = q.queue[:len(q.queue)-1]
	q.size--
	if !q.IsEmpty() {
		q.Head = q.queue[len(q.queue)-1]
	} else {
		q.Head = nil
	}
	return node
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// Serialize func changes linked list to a serialized string
func Serialize(root *TreeNode) string {
	serial := make([]string, 0)
	q0 := NewQueue()
	var bfs func(q Queue)
	q0.Push(root)
	bfs = func(q Queue) {
		if q.IsEmpty() {
			return
		}
		//create new queue for next level
		qn := NewQueue()
		for !q.IsEmpty() {
			serial = append(serial, strconv.Itoa(q.Head.Val))
			if q.Head.Left != nil {
				qn.Push(q.Head.Left)
			}
			if q.Head.Right != nil {
				qn.Push(q.Head.Right)
			}
			q.Pop()
		}
		bfs(*qn)
	}
	bfs(*q0)
	return strings.Join(serial, ", ")
}
