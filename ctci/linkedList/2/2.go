package main

import "fmt"

func main() {
	n4 := &Node{
		val: 4,
	}
	n3 := &Node{
		val:  3,
		Next: n4,
	}
	n2 := &Node{
		val:  2,
		Next: n3,
	}
	n1 := &Node{
		val:  1,
		Next: n2,
	}
	root := &Node{
		val:  1,
		Next: n1,
	}
	res := find(root, 1)
	fmt.Println(res)
}

type Node struct {
	val  int
	Next *Node
}

func find(root *Node, k int) int {
	cur := root
	for i := 0; i < k; i++ {
		cur = cur.Next
	}
	res := root
	for cur.Next != nil {
		cur = cur.Next
		res = res.Next
	}
	return res.val
}
